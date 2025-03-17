#pragma once

#include <string>
#include <vector>

#include "envoy/network/transport_socket.h"
#include "envoy/registry/registry.h"
#include "envoy/server/transport_socket_config.h"
#include "envoy/stats/stats_macros.h" // IWYU pragma: keep

#include "source/common/protobuf/protobuf.h" // IWYU pragma: keep

#include "absl/status/statusor.h"

namespace Envoy {
namespace Khulnasoft {

/**
 * Config registration for the Khulnasoft TLS wrapper transport socket factory.
 * @see TransportSocketConfigFactory.
 */
class TlsWrapperConfigFactory : public virtual Server::Configuration::TransportSocketConfigFactory {
public:
  ~TlsWrapperConfigFactory() override = default;
  std::string name() const override { return name_; }

  const std::string name_ = "khulnasoft.tls_wrapper";
};

class UpstreamTlsWrapperFactory
    : public Server::Configuration::UpstreamTransportSocketConfigFactory,
      public TlsWrapperConfigFactory {
public:
  absl::StatusOr<Network::UpstreamTransportSocketFactoryPtr> createTransportSocketFactory(
      const Protobuf::Message& config,
      Server::Configuration::TransportSocketFactoryContext& context) override;
  ProtobufTypes::MessagePtr createEmptyConfigProto() override;
};

DECLARE_FACTORY(UpstreamTlsWrapperFactory);

class DownstreamTlsWrapperFactory
    : public Server::Configuration::DownstreamTransportSocketConfigFactory,
      public TlsWrapperConfigFactory {
public:
  absl::StatusOr<Network::DownstreamTransportSocketFactoryPtr>
  createTransportSocketFactory(const Protobuf::Message& config,
                               Server::Configuration::TransportSocketFactoryContext& context,
                               const std::vector<std::string>& server_names) override;
  ProtobufTypes::MessagePtr createEmptyConfigProto() override;
};

DECLARE_FACTORY(DownstreamTlsWrapperFactory);
} // namespace Khulnasoft
} // namespace Envoy
