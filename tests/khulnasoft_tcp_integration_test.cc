#include <fmt/base.h>
#include <fmt/format.h>
#include <gtest/gtest-param-test.h>
#include <gtest/gtest.h>

#include <chrono>
#include <cstdint>
#include <cstring>
#include <string>

#include "test/integration/fake_upstream.h"
#include "test/integration/integration_tcp_client.h"
#include "test/test_common/environment.h"
#include "test/test_common/utility.h"

#include "absl/time/clock.h"
#include "absl/time/time.h"
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
          proxylib: "proxylib/libkhulnasoft.so"
      - name: envoy.tcp_proxy
        typed_config:
          "@type": type.googleapis.com/envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy
          stat_prefix: tcp_stats
          cluster: cluster1
)EOF";

class KhulnasoftTcpProxyIntegrationTest : public KhulnasoftTcpIntegrationTest {
public:
  KhulnasoftTcpProxyIntegrationTest()
      : KhulnasoftTcpIntegrationTest(fmt::format(
            fmt::runtime(TestEnvironment::substitute(khulnasoft_tcp_proxy_config_fmt, GetParam())),
            "true")) {}
};

INSTANTIATE_TEST_SUITE_P(IpVersions, KhulnasoftTcpProxyIntegrationTest,
                         testing::ValuesIn(TestEnvironment::getIpVersionsForTest()),
                         TestUtility::ipTestParamsToString);

// Test upstream writing before downstream downstream does.
TEST_P(KhulnasoftTcpProxyIntegrationTest, KhulnasoftTcpProxyUpstreamWritesFirst) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->write("hello"));
  tcp_client->waitForData("hello");

  ASSERT_TRUE(tcp_client->write("hello"));
  ASSERT_TRUE(fake_upstream_connection->waitForData(5));

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

// Test proxying data in both directions, and that all data is flushed properly
// when there is an upstream disconnect.
TEST_P(KhulnasoftTcpProxyIntegrationTest, KhulnasoftTcpProxyUpstreamDisconnect) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  ASSERT_TRUE(tcp_client->write("hello"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->waitForData(5));
  ASSERT_TRUE(fake_upstream_connection->write("world"));
  ASSERT_TRUE(fake_upstream_connection->close());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
  tcp_client->waitForHalfClose();
  tcp_client->close();

  EXPECT_EQ("world", tcp_client->data());
}

// Test proxying data in both directions, and that all data is flushed properly
// when the client disconnects.
TEST_P(KhulnasoftTcpProxyIntegrationTest, KhulnasoftTcpProxyDownstreamDisconnect) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  ASSERT_TRUE(tcp_client->write("hello"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->waitForData(5));
  ASSERT_TRUE(fake_upstream_connection->write("world"));
  tcp_client->waitForData("world");
  ASSERT_TRUE(tcp_client->write("hello", true));
  ASSERT_TRUE(fake_upstream_connection->waitForData(10));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
  tcp_client->waitForDisconnect();
}

TEST_P(KhulnasoftTcpProxyIntegrationTest, KhulnasoftTcpProxyLargeWrite) {
  config_helper_.setBufferLimits(1024, 1024);
  initialize();

  std::string data(1024 * 16, 'a');
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  ASSERT_TRUE(tcp_client->write(data));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->waitForData(data.size()));
  ASSERT_TRUE(fake_upstream_connection->write(data));
  tcp_client->waitForData(data);
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
TEST_P(KhulnasoftTcpProxyIntegrationTest, KhulnasoftTcpProxyDownstreamFlush) {
  // Use a very large size to make sure it is larger than the kernel socket read
  // buffer.
  const uint32_t size = 50 * 1024 * 1024;
  config_helper_.setBufferLimits(size / 4, size / 4);
  initialize();

  std::string data(size, 'a');
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

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
TEST_P(KhulnasoftTcpProxyIntegrationTest, KhulnasoftTcpProxyUpstreamFlush) {
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

  ASSERT_TRUE(tcp_client->write(data, true, true, std::chrono::milliseconds(30000)));

  test_server_->waitForGaugeEq("tcp.tcp_stats.upstream_flush_active", 1);
  ASSERT_TRUE(fake_upstream_connection->readDisable(false));
  ASSERT_TRUE(fake_upstream_connection->waitForData(data.size()));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
  tcp_client->waitForHalfClose();

  EXPECT_EQ(test_server_->counter("tcp.tcp_stats.upstream_flush_total")->value(), 1);
  test_server_->waitForGaugeEq("tcp.tcp_stats.upstream_flush_active", 0);
}

// Test that Envoy doesn't crash or assert when shutting down with an upstream
// flush active
TEST_P(KhulnasoftTcpProxyIntegrationTest, KhulnasoftTcpProxyUpstreamFlushEnvoyExit) {
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

  ASSERT_TRUE(tcp_client->write(data, true));

  test_server_->waitForGaugeEq("tcp.tcp_stats.upstream_flush_active", 1);
  test_server_.reset();
  ASSERT_TRUE(fake_upstream_connection->close());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());

  // Success criteria is that no ASSERTs fire and there are no leaks.
}

//
// Khulnasoft Go test parser "linetester" with TCP proxy
//

// params: is_ingress ("true", "false")
const std::string khulnasoft_linetester_config_fmt = R"EOF(
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
          proxylib: "proxylib/libkhulnasoft.so"
      - name: envoy.tcp_proxy
        typed_config:
          "@type": type.googleapis.com/envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy
          stat_prefix: tcp_stats
          cluster: cluster1
)EOF";

const std::string TCP_POLICY_LINEPARSER_fmt = R"EOF(version_info: "0"
resources:
- "@type": type.googleapis.com/khulnasoft.NetworkPolicy
  endpoint_ips:
  - '{{ ntop_ip_loopback_address }}'
  policy: 3
  ingress_per_port_policies:
  - port: 1
    end_port: {0}
    rules:
    - remote_policies: [ 1 ]
      l7_proto: "test.lineparser"
  egress_per_port_policies:
  - port: 1
    end_port: {0}
    rules:
    - remote_policies: [ 1 ]
      l7_proto: "test.lineparser"
)EOF";

class KhulnasoftGoLinetesterIntegrationTest : public KhulnasoftTcpIntegrationTest {
public:
  KhulnasoftGoLinetesterIntegrationTest()
      : KhulnasoftTcpIntegrationTest(fmt::format(
            fmt::runtime(TestEnvironment::substitute(khulnasoft_linetester_config_fmt, GetParam())),
            "true")) {}

  std::string testPolicyFmt() override {
    return TestEnvironment::substitute(TCP_POLICY_LINEPARSER_fmt, GetParam());
  }
};

INSTANTIATE_TEST_SUITE_P(IpVersions, KhulnasoftGoLinetesterIntegrationTest,
                         testing::ValuesIn(TestEnvironment::getIpVersionsForTest()),
                         TestUtility::ipTestParamsToString);

static FakeRawConnection::ValidatorFunction noMatch(const char* data_to_not_match) {
  return [data_to_not_match](const std::string& data) -> bool {
    auto found = data.find(data_to_not_match);
    return found == std::string::npos;
  };
}

TEST_P(KhulnasoftGoLinetesterIntegrationTest, KhulnasoftGoLineParserUpstreamWritesFirst) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->write("DROP reply direction\n"));
  ASSERT_TRUE(fake_upstream_connection->write("PASS reply direction\n"));
  tcp_client->waitForData("PASS reply direction\n");

  ASSERT_TRUE(tcp_client->write("PASS original direction\n"));
  ASSERT_TRUE(
      fake_upstream_connection->waitForData(FakeRawConnection::waitForInexactMatch("PASS")));

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

TEST_P(KhulnasoftGoLinetesterIntegrationTest, KhulnasoftGoLineParserPartialLines) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->write("DROP reply "));
  absl::SleepFor(absl::Milliseconds(10));
  ASSERT_TRUE(fake_upstream_connection->write("direction\nPASS"));
  absl::SleepFor(absl::Milliseconds(10));
  ASSERT_TRUE(fake_upstream_connection->write(" reply direction\n"));
  tcp_client->waitForData("PASS reply direction\n");

  ASSERT_TRUE(tcp_client->write("PASS original direction\n"));
  ASSERT_TRUE(fake_upstream_connection->waitForData(
      FakeRawConnection::waitForInexactMatch("PASS original direction\n")));

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

TEST_P(KhulnasoftGoLinetesterIntegrationTest, KhulnasoftGoLineParserInject) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(tcp_client->write("INJECT reply direction\n"));
  ASSERT_TRUE(tcp_client->write("PASS original direction\n"));
  ASSERT_TRUE(fake_upstream_connection->write("PASS reply direction\n"));

  // These can in principle arrive in either order
  tcp_client->waitForData("PASS reply direction\n", false);
  tcp_client->waitForData("INJECT reply direction\n", false);

  ASSERT_TRUE(fake_upstream_connection->waitForData(
      FakeRawConnection::waitForInexactMatch("PASS original direction\n")));

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

TEST_P(KhulnasoftGoLinetesterIntegrationTest, KhulnasoftGoLineParserInjectPartial) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->write("PASS reply"));
  ASSERT_TRUE(tcp_client->write("INJECT reply direction\n"));
  ASSERT_TRUE(tcp_client->write("PASS original direction\n"));

  ASSERT_TRUE(fake_upstream_connection->write(" direction\n"));

  // These can in principle arrive in either order
  tcp_client->waitForData("PASS reply direction\n", false);
  tcp_client->waitForData("INJECT reply direction\n", false);

  ASSERT_TRUE(fake_upstream_connection->waitForData(
      FakeRawConnection::waitForInexactMatch("PASS original direction\n")));

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

TEST_P(KhulnasoftGoLinetesterIntegrationTest, KhulnasoftGoLineParserInjectPartialMultiple) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->write("PASS reply"));
  ASSERT_TRUE(tcp_client->write("INJECT reply direction\n"));
  ASSERT_TRUE(tcp_client->write("DROP original direction\n"));
  ASSERT_TRUE(tcp_client->write("INSERT original direction\n"));

  ASSERT_TRUE(fake_upstream_connection->write(" direction\n"));

  // These can in principle arrive in either order
  absl::SleepFor(absl::Milliseconds(10));
  tcp_client->waitForData("PASS reply direction\n", false);
  absl::SleepFor(absl::Milliseconds(10));
  tcp_client->waitForData("INJECT reply direction\n", false);

  ASSERT_TRUE(fake_upstream_connection->waitForData(
      FakeRawConnection::waitForInexactMatch("INSERT original direction\n")));
  ASSERT_TRUE(fake_upstream_connection->waitForData(noMatch("DROP")));

  ASSERT_TRUE(fake_upstream_connection->write("DROP reply direction\n"));
  ASSERT_TRUE(fake_upstream_connection->write("PASS2 reply direction\n"));
  tcp_client->waitForData("PASS2 reply direction\n", false);

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

//
// Khulnasoft Go test parser "blocktester" with TCP proxy
//

// params: is_ingress ("true", "false")
const std::string khulnasoft_blocktester_config_fmt = R"EOF(
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
          proxylib: "proxylib/libkhulnasoft.so"
          proxylib_params:
            access-log-path: "{{ test_udsdir }}/access_log.sock"
      - name: envoy.tcp_proxy
        typed_config:
          "@type": type.googleapis.com/envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy
          stat_prefix: tcp_stats
          cluster: cluster1
)EOF";

const std::string TCP_POLICY_BLOCKPARSER_fmt = R"EOF(version_info: "0"
resources:
- "@type": type.googleapis.com/khulnasoft.NetworkPolicy
  endpoint_ips:
  - '{{ ntop_ip_loopback_address }}'
  policy: 3
  ingress_per_port_policies:
  - port: 1
    end_port: 65535
    rules:
    - remote_policies: [ 1 ]
      l7_proto: "test.blockparser"
  egress_per_port_policies:
  - port: 1
    end_port: 65535
    rules:
    - remote_policies: [ 1 ]
      l7_proto: "test.blockparser"
)EOF";

class KhulnasoftGoBlocktesterIntegrationTest : public KhulnasoftTcpIntegrationTest {
public:
  KhulnasoftGoBlocktesterIntegrationTest()
      : KhulnasoftTcpIntegrationTest(fmt::format(
            fmt::runtime(TestEnvironment::substitute(khulnasoft_blocktester_config_fmt, GetParam())),
            "true")) {}

  std::string testPolicyFmt() override {
    return TestEnvironment::substitute(TCP_POLICY_BLOCKPARSER_fmt, GetParam());
  }
};

INSTANTIATE_TEST_SUITE_P(IpVersions, KhulnasoftGoBlocktesterIntegrationTest,
                         testing::ValuesIn(TestEnvironment::getIpVersionsForTest()),
                         TestUtility::ipTestParamsToString);

TEST_P(KhulnasoftGoBlocktesterIntegrationTest, KhulnasoftGoBlockParserUpstreamWritesFirst) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->write("24:DROP reply direction\n"));
  ASSERT_TRUE(fake_upstream_connection->write("24:PASS reply direction\n"));
  tcp_client->waitForData("24:PASS reply direction\n");

  ASSERT_TRUE(tcp_client->write("27:PASS original direction\n"));
  ASSERT_TRUE(
      fake_upstream_connection->waitForData(FakeRawConnection::waitForInexactMatch("PASS")));

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

TEST_P(KhulnasoftGoBlocktesterIntegrationTest, KhulnasoftGoBlockParserPartialBlocks) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->write("24:DROP reply "));
  ASSERT_TRUE(fake_upstream_connection->write("direction\n24:PASS"));
  ASSERT_TRUE(fake_upstream_connection->write(" reply direction\n"));
  tcp_client->waitForData("24:PASS reply direction\n");

  ASSERT_TRUE(tcp_client->write("27:PASS original direction\n"));
  ASSERT_TRUE(fake_upstream_connection->waitForData(
      FakeRawConnection::waitForInexactMatch("27:PASS original direction\n")));

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

TEST_P(KhulnasoftGoBlocktesterIntegrationTest, KhulnasoftGoBlockParserInject) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(tcp_client->write("26:INJECT reply direction\n"));
  ASSERT_TRUE(tcp_client->write("27:PASS original direction\n"));
  ASSERT_TRUE(fake_upstream_connection->write("24:PASS reply direction\n"));

  // These can in principle arrive in either order
  absl::SleepFor(absl::Milliseconds(10));
  tcp_client->waitForData("24:PASS reply direction\n", false);
  absl::SleepFor(absl::Milliseconds(10));
  tcp_client->waitForData("26:INJECT reply direction\n", false);

  ASSERT_TRUE(fake_upstream_connection->waitForData(
      FakeRawConnection::waitForInexactMatch("27:PASS original direction\n")));

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

TEST_P(KhulnasoftGoBlocktesterIntegrationTest, KhulnasoftGoBlockParserInjectPartial) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->write("24:PASS reply"));
  ASSERT_TRUE(tcp_client->write("26:INJECT reply direction\n"));
  ASSERT_TRUE(tcp_client->write("27:PASS original direction\n"));

  ASSERT_TRUE(fake_upstream_connection->write(" direction\n"));

  // These can in principle arrive in either order
  tcp_client->waitForData("24:PASS reply direction\n", false);
  tcp_client->waitForData("26:INJECT reply direction\n", false);

  ASSERT_TRUE(fake_upstream_connection->waitForData(
      FakeRawConnection::waitForInexactMatch("27:PASS original direction\n")));

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

TEST_P(KhulnasoftGoBlocktesterIntegrationTest, KhulnasoftGoBlockParserInjectPartialMultiple) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(fake_upstream_connection->write("24:PASS reply"));
  ASSERT_TRUE(tcp_client->write("26:INJECT reply direction\n"));
  ASSERT_TRUE(tcp_client->write("27:DROP original direction\n"));
  ASSERT_TRUE(tcp_client->write("29:INSERT original direction\n"));

  absl::SleepFor(absl::Milliseconds(100));
  ASSERT_TRUE(fake_upstream_connection->write(" dire"));

  absl::SleepFor(absl::Milliseconds(100));
  ASSERT_TRUE(fake_upstream_connection->write("ction\n"));

  // These can in principle arrive in either order
  tcp_client->waitForData("24:PASS reply direction\n", false);
  tcp_client->waitForData("26:INJECT reply direction\n", false);

  ASSERT_TRUE(fake_upstream_connection->waitForData(
      FakeRawConnection::waitForInexactMatch("29:INSERT original direction\n")));
  ASSERT_TRUE(fake_upstream_connection->waitForData(noMatch("DROP")));

  ASSERT_TRUE(fake_upstream_connection->write("24:DROP reply direction\n"));
  ASSERT_TRUE(fake_upstream_connection->write("25:PASS2 reply direction\n"));
  tcp_client->waitForData("25:PASS2 reply direction\n", false);

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

TEST_P(KhulnasoftGoBlocktesterIntegrationTest, KhulnasoftGoBlockParserInjectBufferOverflow) {
  initialize();
  IntegrationTcpClientPtr tcp_client = makeTcpConnection(lookupPort("tcp_proxy"));
  FakeRawConnectionPtr fake_upstream_connection;
  ASSERT_TRUE(fake_upstreams_[0]->waitForRawConnection(fake_upstream_connection));

  ASSERT_TRUE(tcp_client->write("26:INJECT reply direction\n"));
  ASSERT_TRUE(tcp_client->write("27:DROP original direction\n"));

  char buf[5000];
  memset(buf, 'A', sizeof buf);
  strncpy(buf, "5000:INSERT original direction", 30);
  buf[sizeof buf - 1] = '\n';

  ASSERT_TRUE(tcp_client->write(buf));
  tcp_client->waitForData("26:INJECT reply direction\n", false);

  ASSERT_TRUE(fake_upstream_connection->waitForData(
      FakeRawConnection::waitForInexactMatch("INSERT original direction")));
  ASSERT_TRUE(fake_upstream_connection->waitForData(noMatch("DROP")));

  ASSERT_TRUE(fake_upstream_connection->write("24:DROP reply direction\n"));
  ASSERT_TRUE(fake_upstream_connection->write("25:PASS2 reply direction\n"));
  tcp_client->waitForData("25:PASS2 reply direction\n", false);

  ASSERT_TRUE(fake_upstream_connection->write("", true));
  tcp_client->waitForHalfClose();
  ASSERT_TRUE(tcp_client->write("", true));
  ASSERT_TRUE(fake_upstream_connection->waitForHalfClose());
  ASSERT_TRUE(fake_upstream_connection->waitForDisconnect());
}

} // namespace Envoy
