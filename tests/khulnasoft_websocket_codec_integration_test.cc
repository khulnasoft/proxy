#include <fmt/base.h>
#include <fmt/format.h>
#include <gtest/gtest-param-test.h>
#include <gtest/gtest.h>

#include <chrono>
#include <cstdint>
#include <string>

#include "test/integration/fake_upstream.h"
#include "test/integration/integration_tcp_client.h"
#include "test/test_common/environment.h"
#include "test/test_common/utility.h"

#include "tests/khulnasoft_tcp_integration.h"

namespace Envoy {

//
// Khulnasoft filters with TCP proxy
//

// params: is_ingress ("true", "false")
const std::string khulnasoft_tcp_proxy_config_fmt = R"EOF(
admin:
  address:
    socket_address:
      address: 127.0.0.1
      port_value: 0
static_resources:
  clusters:
  - name: cluster1
    type: ORIGINAL_DST
    lb_policy: CLUSTER_PROVIDED
    connect_timeout:
      seconds: 1
  - name: xds-grpc-khulnasoft
    connect_timeout:
      seconds: 5
    type: STATIC
    lb_policy: ROUND_ROBIN
    http2_protocol_options:
    load_assignment:
      cluster_name: xds-grpc-khulnasoft
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              pipe:
                path: /var/run/khulnasoft/xds.sock
  listeners:
    name: listener_0
    address:
      socket_address:
        address: 127.0.0.1
        port_value: 0
    listener_filters:
      name: test_bpf_metadata
      typed_config:
        "@type": type.googleapis.com/khulnasoft.TestBpfMetadata
        is_ingress: {0}
    filter_chains:
      filters:
      - name: khulnasoft.network
        typed_config:
          "@type": type.googleapis.com/khulnasoft.NetworkFilter
      - name: khulnasoft.network.websocket.client
        typed_config:
          "@type": type.googleapis.com/khulnasoft.WebSocketClient
          access_log_path: "{{ test_udsdir }}/access_log.sock"
          origin: "jarno.khulnasoft.rocks"
          host: "jarno.khulnasoft.rocks"
          ping_interval:
            nanos: 1000000
          ping_when_idle: true
      - name: khulnasoft.network.websocket.server
        typed_config:
          "@type": type.googleapis.com/khulnasoft.WebSocketServer
          access_log_path: "{{ test_udsdir }}/access_log.sock"
          origin: "jarno.khulnasoft.rocks"
      - name: envoy.tcp_proxy
        typed_config:
          "@type": type.googleapis.com/envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy
          stat_prefix: tcp_stats
          cluster: cluster1
)EOF";

class KhulnasoftWebSocketIntegrationTest : public KhulnasoftTcpIntegrationTest {
public:
  KhulnasoftWebSocketIntegrationTest()
      : KhulnasoftTcpIntegrationTest(fmt::format(
            fmt::runtime(TestEnvironment::substitute(khulnasoft_tcp_proxy_config_fmt, GetParam())),
            "true")) {}

  std::string testPolicyFmt() override {
    return TestEnvironment::substitute(R"EOF(version_info: "0"
resources:
- "@type": type.googleapis.com/khulnasoft.NetworkPolicy
  endpoint_ips:
  - '{{ ntop_ip_loopback_address }}'
  policy: 3
  ingress_per_port_policies:
  - port: {0}
    rules:
    - remote_policies: [ 1 ]
  egress_per_port_policies:
  - port: {0}
    rules:
    - remote_policies: [ 1 ]
)EOF",
                                       GetParam());
  }
};

INSTANTIATE_TEST_SUITE_P(IpVersions, KhulnasoftWebSocketIntegrationTest,
                         testing::ValuesIn(TestEnvironment::getIpVersionsForTest()),
                         TestUtility::ipTestParamsToString);

// Test upstream writing before downstream downstream does.
TEST_P(KhulnasoftWebSocketIntegrationTest, KhulnasoftWebSocketUpstreamWritesFirst) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  test_server_->waitForCounterGe("websocket.ping_sent_count", 1);

  ASSERT_TRUE(fake_upstream_connection->write("hello"));
  tcp_client->waitForData("hello");

  ASSERT_TRUE(tcp_client->write("hello"));
  std::string received;
  ASSERT_TRUE(fake_upstream_connection->waitForData(5, &received));
  ASSERT_EQ(received, "hello");

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

// Test proxying data in both directions, and that all data is flushed properly
// when there is an upstream disconnect.
TEST_P(KhulnasoftWebSocketIntegrationTest, KhulnasoftWebSocketUpstreamDisconnect) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  ASSERT_TRUE(tcp_client->write("hello"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  std::string received;
  ASSERT_TRUE(fake_upstream_connection->waitForData(5, &received));
  ASSERT_EQ(received, "hello");

  test_server_->waitForCounterGe("websocket.ping_sent_count", 1);

  ASSERT_TRUE(fake_upstream_connection->write("world"));
  ASSERT_TRUE(fake_upstream_connection->close());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
  tcp_client->waitForHalfClose();
  tcp_client->close();

  EXPECT_EQ("world", tcp_client->data());
}

// Test proxying data in both directions, and that all data is flushed properly
// when the client disconnects.
TEST_P(KhulnasoftWebSocketIntegrationTest, KhulnasoftWebSocketDownstreamDisconnect) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  ASSERT_TRUE(tcp_client->write("hello"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  std::string received;
  ASSERT_TRUE(fake_upstream_connection->waitForData(5, &received));
  ASSERT_EQ(received, "hello");
  ASSERT_TRUE(fake_upstream_connection->write("world"));
  tcp_client->waitForData("world");

  test_server_->waitForCounterGe("websocket.ping_sent_count", 1);

  ASSERT_TRUE(tcp_client->write("hello", true));
  ASSERT_TRUE(fake_upstream_connection->waitForData(10, &received));
  ASSERT_EQ(received, "hellohello");
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
  tcp_client->waitForDisconnect();
}

TEST_P(KhulnasoftWebSocketIntegrationTest, KhulnasoftWebSocketLargeWrite) {
  config_helper_.setBufferLimits(1024, 1024);
  initialize();

  std::string data(1024 * 16, 'a');
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  ASSERT_TRUE(tcp_client->write(data));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  std::string received;
  ASSERT_TRUE(fake_upstream_connection->waitForData(data.size(), &received));
  ASSERT_EQ(received, data);
  ASSERT_TRUE(fake_upstream_connection->write(data));
  tcp_client->waitForData(data);

  test_server_->waitForCounterGe("websocket.ping_sent_count", 1);

  tcp_client->close();
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->close());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());

  uint32_t upstream_pauses =
      test_server_->counter("cluster.cluster1.upstream_flow_control_paused_reading_total")->value();
  uint32_t upstream_resumes =
      test_server_->counter("cluster.cluster1.upstream_flow_control_resumed_reading_total")
          ->value();
  EXPECT_EQ(upstream_pauses, upstream_resumes);
  uint32_t downstream_pauses =
      test_server_->counter("tcp.tcp_stats.downstream_flow_control_paused_reading_total")->value();
  uint32_t downstream_resumes =
      test_server_->counter("tcp.tcp_stats.downstream_flow_control_resumed_reading_total")->value();
  EXPECT_EQ(downstream_pauses, downstream_resumes);
}

// Test that a downstream flush works correctly (all data is flushed)
TEST_P(KhulnasoftWebSocketIntegrationTest, KhulnasoftWebSocketDownstreamFlush) {
  // Use a very large size to make sure it is larger than the kernel socket read
  // buffer.
  const uint32_t size = 50 * 1024 * 1024;
  config_helper_.setBufferLimits(size / 4, size / 4);
  enableHalfClose(true);
  initialize();

  std::string data(size, 'a');
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  test_server_->waitForCounterGe("websocket.ping_sent_count", 1);

  tcp_client->readDisable(true);
  ASSERT_TRUE(tcp_client->write("", true));

  // This ensures that readDisable(true) has been run on it's thread
  // before tcp_client starts writing.
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());

  ASSERT_TRUE(fake_upstream_connection->write(data, true));

  test_server_->waitForCounterGe("cluster.cluster1.upstream_flow_control_paused_reading_total", 1);
  EXPECT_EQ(test_server_->counter("cluster.cluster1.upstream_flow_control_resumed_reading_total")
                ->value(),
            0);
  tcp_client->readDisable(false);
  tcp_client->waitForData(data);
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());

  uint32_t upstream_pauses =
      test_server_->counter("cluster.cluster1.upstream_flow_control_paused_reading_total")->value();
  uint32_t upstream_resumes =
      test_server_->counter("cluster.cluster1.upstream_flow_control_resumed_reading_total")
          ->value();
  EXPECT_GE(upstream_pauses, upstream_resumes);
  EXPECT_GT(upstream_resumes, 0);
}

// Test that an upstream flush works correctly (all data is flushed)
TEST_P(KhulnasoftWebSocketIntegrationTest, KhulnasoftWebSocketUpstreamFlush) {
  // Use a very large size to make sure it is larger than the kernel socket read
  // buffer.
  const uint32_t size = 50 * 1024 * 1024;
  config_helper_.setBufferLimits(size, size);
  enableHalfClose(true);
  initialize();

  std::string data(size, 'a');
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  test_server_->waitForCounterGe("websocket.ping_sent_count", 1);

  ASSERT_TRUE(fake_upstream_connection->readDisable(true));
  ASSERT_TRUE(fake_upstream_connection->write("", true));

  // This ensures that fake_upstream_connection->readDisable has been run on
  // it's thread before tcp_client starts writing.
  tcp_client->waitForHalfClose();

  ASSERT_TRUE(tcp_client->write(data, true, true, std::chrono::milliseconds(30000)));

  test_server_->waitForGaugeEq("tcp.tcp_stats.upstream_flush_active", 1);

  ASSERT_TRUE(fake_upstream_connection->readDisable(false));
  std::string received;
  ASSERT_TRUE(fake_upstream_connection->waitForData(data.size(), &received));
  ASSERT_EQ(received, data);
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
  tcp_client->waitForHalfClose();

  EXPECT_EQ(test_server_->counter("tcp.tcp_stats.upstream_flush_total")->value(), 1);
  test_server_->waitForGaugeEq("tcp.tcp_stats.upstream_flush_active", 0);
}

// Test that Envoy doesn't crash or assert when shutting down with an upstream
// flush active
TEST_P(KhulnasoftWebSocketIntegrationTest, KhulnasoftWebSocketUpstreamFlushEnvoyExit) {
  // Use a very large size to make sure it is larger than the kernel socket read
  // buffer.
  const uint32_t size = 50 * 1024 * 1024;
  config_helper_.setBufferLimits(size, size);
  initialize();

  std::string data(size, 'a');
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->readDisable(true));
  ASSERT_TRUE(fake_upstream_connection->write("", true));

  // This ensures that fake_upstream_connection->readDisable has been run on
  // it's thread before tcp_client starts writing.
  tcp_client->waitForHalfClose();

  test_server_->waitForCounterGe("websocket.ping_sent_count", 1);

  ASSERT_TRUE(tcp_client->write(data, true));

  test_server_->waitForGaugeEq("tcp.tcp_stats.upstream_flush_active", 1);
  test_server_.reset();
  ASSERT_TRUE(fake_upstream_connection->close());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());

  // Success criteria is that no ASSERTs fire and there are no leaks.
}

} // namespace Envoy
