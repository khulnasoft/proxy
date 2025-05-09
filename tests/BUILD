load(
    "@envoy//bazel:envoy_build_system.bzl",
    "envoy_cc_test",
    "envoy_cc_test_library",
    "envoy_package",
)
load(
    "@envoy_api//bazel:api_build_system.bzl",
    "api_cc_py_proto_library",
)

licenses(["notice"])  # Apache 2

envoy_package()

api_cc_py_proto_library(
    name = "bpf_metadata",
    srcs = ["bpf_metadata.proto"],
)

envoy_cc_test_library(
    name = "accesslog_server_lib",
    srcs = ["accesslog_server.cc"],
    hdrs = ["accesslog_server.h"],
    repository = "@envoy",
    deps = [
        ":uds_server_lib",
        "//khulnasoft/api:accesslog_proto_cc_proto",
        "@envoy//test/test_common:utility_lib",
    ],
)

envoy_cc_test_library(
    name = "uds_server_lib",
    srcs = ["uds_server.cc"],
    hdrs = ["uds_server.h"],
    repository = "@envoy",
    deps = [
        "@envoy//source/common/common:logger_lib",
        "@envoy//source/common/common:thread_lib",
        "@envoy//source/common/network:address_lib",
        "@envoy//test/test_common:thread_factory_for_test_lib",
    ],
)

envoy_cc_test_library(
    name = "bpf_metadata_lib",
    srcs = ["bpf_metadata.cc"],
    hdrs = ["bpf_metadata.h"],
    repository = "@envoy",
    deps = [
        ":bpf_metadata_cc_proto",
        "//khulnasoft:bpf_metadata_lib",
        "//khulnasoft:network_policy_lib",
        "//khulnasoft:socket_option_lib",
        "@envoy//source/extensions/config_subscription/filesystem:filesystem_subscription_lib",
        "@envoy//test/test_common:environment_lib",
    ],
)

envoy_cc_test_library(
    name = "khulnasoft_http_integration_lib",
    srcs = ["khulnasoft_http_integration.cc"],
    hdrs = ["khulnasoft_http_integration.h"],
    repository = "@envoy",
    deps = [
        ":accesslog_server_lib",
        ":bpf_metadata_lib",
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
        "@envoy//test/integration:http_integration_lib",
    ],
)

envoy_cc_test_library(
    name = "khulnasoft_tls_integration_lib",
    srcs = ["khulnasoft_tls_integration.cc"],
    hdrs = ["khulnasoft_tls_integration.h"],
    data = [
        "@envoy//test/config/integration/certs",
    ],
    repository = "@envoy",
    deps = [
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
        "@envoy//source/extensions/filters/listener/tls_inspector:config",
        "@envoy//source/extensions/filters/network/tcp_proxy:config",
        "@envoy//test/integration:integration_lib",
    ],
)

envoy_cc_test_library(
    name = "khulnasoft_tcp_integration_lib",
    srcs = ["khulnasoft_tcp_integration.cc"],
    hdrs = ["khulnasoft_tcp_integration.h"],
    repository = "@envoy",
    deps = [
        ":accesslog_server_lib",
        ":bpf_metadata_lib",
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
        "@envoy//source/extensions/filters/network/tcp_proxy:config",
        "@envoy//test/integration:integration_lib",
        "@envoy//test/integration/clusters:custom_static_cluster",
    ],
)

envoy_cc_test(
    name = "khulnasoft_network_policy_test",
    srcs = ["khulnasoft_network_policy_test.cc"],
    repository = "@envoy",
    deps = [
        "//khulnasoft:network_policy_lib",
        "@envoy//test/mocks/server:factory_context_mocks",
    ],
)

envoy_cc_test(
    name = "metadata_config_test",
    srcs = ["metadata_config_test.cc"],
    repository = "@envoy",
    deps = [
        ":bpf_metadata_lib",
        "@envoy//envoy/network:filter_interface",
        "@envoy//test/mocks/server:listener_factory_context_mocks",
    ],
)

envoy_cc_test(
    name = "accesslog_test",
    srcs = ["accesslog_test.cc"],
    repository = "@envoy",
    deps = [
        "//khulnasoft:accesslog_lib",
        "@envoy//test/mocks/network:connection_mocks",
        "@envoy//test/mocks/stream_info:stream_info_mocks",
        "@envoy//test/test_common:utility_lib",
    ],
)

envoy_cc_test(
    name = "khulnasoft_tcp_integration_test",
    srcs = ["khulnasoft_tcp_integration_test.cc"],
    data = [
        "//proxylib:libkhulnasoft.so",
    ],
    repository = "@envoy",
    deps = [
        ":bpf_metadata_lib",
        ":khulnasoft_tcp_integration_lib",
        "//khulnasoft:network_filter_lib",
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
        "@envoy//test/integration/clusters:custom_static_cluster",
    ],
)

envoy_cc_test(
    name = "khulnasoft_tls_http_integration_test",
    srcs = ["khulnasoft_tls_http_integration_test.cc"],
    data = [
        "@envoy//test/config/integration/certs",
    ],
    repository = "@envoy",
    deps = [
        ":bpf_metadata_lib",
        ":khulnasoft_http_integration_lib",
        ":khulnasoft_tls_integration_lib",
        "//khulnasoft:l7policy_lib",
        "//khulnasoft:network_filter_lib",
        "//khulnasoft:tls_wrapper_lib",
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
    ],
)

envoy_cc_test(
    name = "khulnasoft_tls_tcp_integration_test",
    srcs = ["khulnasoft_tls_tcp_integration_test.cc"],
    data = [
        "//proxylib:libkhulnasoft.so",
        "@envoy//test/config/integration/certs",
    ],
    repository = "@envoy",
    deps = [
        ":bpf_metadata_lib",
        ":khulnasoft_tcp_integration_lib",
        ":khulnasoft_tls_integration_lib",
        "//khulnasoft:l7policy_lib",
        "//khulnasoft:network_filter_lib",
        "//khulnasoft:tls_wrapper_lib",
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
    ],
)

envoy_cc_test(
    name = "khulnasoft_http_integration_test",
    srcs = ["khulnasoft_http_integration_test.cc"],
    data = [
        "//proxylib:libkhulnasoft.so",
    ],
    repository = "@envoy",
    deps = [
        ":bpf_metadata_lib",
        ":khulnasoft_http_integration_lib",
        "//khulnasoft:l7policy_lib",
        "//khulnasoft:network_filter_lib",
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
        "@envoy//test/integration/clusters:custom_static_cluster",
    ],
)

envoy_cc_test(
    name = "khulnasoft_http_upstream_integration_test",
    srcs = ["khulnasoft_http_upstream_integration_test.cc"],
    data = [
        "//proxylib:libkhulnasoft.so",
    ],
    repository = "@envoy",
    deps = [
        ":bpf_metadata_lib",
        ":khulnasoft_http_integration_lib",
        "//khulnasoft:l7policy_lib",
        "//khulnasoft:network_filter_lib",
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
        "@envoy//source/extensions/upstreams/http/http:config",
        "@envoy//source/extensions/upstreams/http/http:upstream_request_lib",
        "@envoy//test/integration/clusters:custom_static_cluster",
    ],
)

envoy_cc_test(
    name = "khulnasoft_websocket_decap_integration_test",
    srcs = ["khulnasoft_websocket_decap_integration_test.cc"],
    data = [
        "//proxylib:libkhulnasoft.so",
    ],
    repository = "@envoy",
    deps = [
        ":bpf_metadata_lib",
        ":khulnasoft_http_integration_lib",
        "//khulnasoft:network_filter_lib",
        "//khulnasoft:websocket_lib",
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
        "@envoy//source/extensions/filters/network/tcp_proxy:config",
    ],
)

envoy_cc_test(
    name = "khulnasoft_websocket_codec_integration_test",
    srcs = ["khulnasoft_websocket_codec_integration_test.cc"],
    data = [
        "//proxylib:libkhulnasoft.so",
    ],
    repository = "@envoy",
    deps = [
        ":bpf_metadata_lib",
        ":khulnasoft_tcp_integration_lib",
        "//khulnasoft:network_filter_lib",
        "//khulnasoft:websocket_lib",
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
    ],
)

envoy_cc_test(
    name = "khulnasoft_websocket_encap_integration_test",
    srcs = ["khulnasoft_websocket_encap_integration_test.cc"],
    data = [
        "//proxylib:libkhulnasoft.so",
    ],
    repository = "@envoy",
    deps = [
        ":bpf_metadata_lib",
        ":khulnasoft_tcp_integration_lib",
        "//khulnasoft:network_filter_lib",
        "//khulnasoft:websocket_lib",
        "@envoy//source/extensions/clusters/logical_dns:logical_dns_cluster_lib",
        "@envoy//source/extensions/clusters/original_dst:original_dst_cluster_lib",
        "@envoy//source/extensions/clusters/static:static_cluster_lib",
        "@envoy//source/extensions/clusters/strict_dns:strict_dns_cluster_lib",
    ],
)

envoy_cc_test(
    name = "health_check_sink_test",
    srcs = [
        "health_check_sink_server.cc",
        "health_check_sink_server.h",
        "health_check_sink_test.cc",
    ],
    repository = "@envoy",
    deps = [
        ":uds_server_lib",
        "//khulnasoft:health_check_sink_lib",
        "@envoy//test/mocks/access_log:access_log_mocks",
        "@envoy//test/mocks/server:health_checker_factory_context_mocks",
        "@envoy//test/test_common:environment_lib",
        "@envoy//test/test_common:utility_lib",
    ],
)
