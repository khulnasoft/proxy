load(
    "@envoy//bazel:envoy_build_system.bzl",
    "envoy_cc_binary",
    "envoy_package",
)

licenses(["notice"])  # Apache 2

envoy_package()

exports_files([
    "linux/bpf.h",
    "linux/bpf_common.h",
    "linux/type_mapper.h",
])

cc_binary(
    name = "khulnasoft-envoy-starter",
    deps = [
        "//starter:main_entry_lib",
    ],
)

envoy_cc_binary(
    name = "khulnasoft-envoy",
    repository = "@envoy",
    deps = [
        # Khulnasoft filters.
        "//khulnasoft:health_check_sink_lib",
        "//khulnasoft:bpf_metadata_lib",
        "//khulnasoft:network_filter_lib",
        "//khulnasoft:l7policy_lib",
        "//khulnasoft:websocket_lib",
        "//khulnasoft:tls_wrapper_lib",
        "@envoy//source/exe:envoy_main_entry_lib",
    ],
)

sh_test(
    name = "envoy_binary_test",
    srcs = ["envoy_binary_test.sh"],
    data = [":khulnasoft-envoy"],
)

sh_binary(
    name = "check_format.py",
    srcs = ["@envoy//tools/code_format:check_format.py"],
    data = [
        ":envoy_build_fixer.py",
        ":header_order.py",
    ],
)

sh_binary(
    name = "header_order.py",
    srcs = ["@envoy//tools/code_format:header_order.py"],
)

sh_binary(
    name = "envoy_build_fixer.py",
    srcs = ["@envoy//tools/code_format:envoy_build_fixer.py"],
)
