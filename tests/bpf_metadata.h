#pragma once

#include <memory>
#include <string>
#include <utility>
#include <vector>

#include "envoy/network/address.h"
#include "envoy/network/listen_socket.h"
#include "envoy/server/factory_context.h"

#include "absl/types/optional.h"
#include "khulnasoft/bpf_metadata.h"
#include "khulnasoft/host_map.h"
#include "khulnasoft/network_policy.h"
#include "tests/bpf_metadata.pb.h"

namespace Envoy {

extern std::string host_map_config;
extern std::shared_ptr<const Khulnasoft::PolicyHostMap> hostmap;

extern Network::Address::InstanceConstSharedPtr original_dst_address;
extern std::shared_ptr<const Khulnasoft::NetworkPolicyMap> npmap;

extern std::string policy_config;
extern std::string policy_path;
extern std::vector<std::pair<std::string, std::string>> sds_configs;

extern void initTestMaps(Server::Configuration::ListenerFactoryContext& context);

namespace Khulnasoft {
namespace BpfMetadata {

class TestConfig : public Config {
public:
  TestConfig(const ::khulnasoft::TestBpfMetadata& config,
             Server::Configuration::ListenerFactoryContext& context);
  ~TestConfig();

  absl::optional<Khulnasoft::BpfMetadata::SocketMetadata>
  extractSocketMetadata(Network::ConnectionSocket& socket) override;

  // Prevent socket options that require NET_ADMIN privileges from being applied during test
  // execution.
  bool addPrivilegedSocketOptions() override { return false; };
};

} // namespace BpfMetadata
} // namespace Khulnasoft
} // namespace Envoy
