#include "khulnasoft/network_filter.h"

#include <dlfcn.h>

#include <cstdint>
#include <memory>
#include <string>
#include <vector>

#include "envoy/buffer/buffer.h"
#include "envoy/network/address.h"
#include "envoy/network/connection.h"
#include "envoy/network/filter.h"
#include "envoy/registry/registry.h"
#include "envoy/server/factory_context.h"
#include "envoy/server/filter_config.h"
#include "envoy/stream_info/filter_state.h"
#include "envoy/stream_info/stream_info.h"
#include "envoy/upstream/host_description.h"

#include "source/common/buffer/buffer_impl.h"
#include "source/common/common/logger.h"
#include "source/common/network/upstream_server_name.h"
#include "source/common/network/upstream_subject_alt_names.h"
#include "source/common/protobuf/protobuf.h" // IWYU pragma: keep
#include "source/common/protobuf/utility.h"

#include "absl/status/statusor.h"
#include "khulnasoft/accesslog.h"
#include "khulnasoft/api/accesslog.pb.h"
#include "khulnasoft/api/network_filter.pb.h"
#include "khulnasoft/api/network_filter.pb.validate.h" // IWYU pragma: keep
#include "khulnasoft/filter_state_khulnasoft_policy.h"
#include "khulnasoft/proxylib.h"
#include "proxylib/types.h"

namespace Envoy {
namespace Server {
namespace Configuration {

/**
 * Config registration for the bpf metadata filter. @see
 * NamedNetworkFilterConfigFactory.
 */
class KhulnasoftNetworkConfigFactory : public NamedNetworkFilterConfigFactory {
public:
  // NamedNetworkFilterConfigFactory
  absl::StatusOr<Network::FilterFactoryCb>
  createFilterFactoryFromProto(const Protobuf::Message& proto_config,
                               FactoryContext& context) override {
    auto config = std::make_shared<Filter::KhulnasoftL3::Config>(
        MessageUtil::downcastAndValidate<const ::khulnasoft::NetworkFilter&>(
            proto_config, context.messageValidationVisitor()),
        context);
    return [config](Network::FilterManager& filter_manager) mutable -> void {
      filter_manager.addFilter(std::make_shared<Filter::KhulnasoftL3::Instance>(config));
    };
  }

  ProtobufTypes::MessagePtr createEmptyConfigProto() override {
    return std::make_unique<::khulnasoft::NetworkFilter>();
  }

  std::string name() const override { return "khulnasoft.network"; }
};

/**
 * Static registration for the bpf metadata filter. @see RegisterFactory.
 */
REGISTER_FACTORY(KhulnasoftNetworkConfigFactory, NamedNetworkFilterConfigFactory);

} // namespace Configuration
} // namespace Server

namespace Filter {
namespace KhulnasoftL3 {

Config::Config(const ::khulnasoft::NetworkFilter& config,
               Server::Configuration::FactoryContext& context)
    : time_source_(context.serverFactoryContext().timeSource()), access_log_(nullptr) {
  const auto& access_log_path = config.access_log_path();
  if (access_log_path.length()) {
    access_log_ = Khulnasoft::AccessLog::Open(access_log_path, time_source_);
  }
  if (config.proxylib().length() > 0) {
    proxylib_ = std::make_shared<Khulnasoft::GoFilter>(config.proxylib(), config.proxylib_params());
  }
}

void Config::Log(Khulnasoft::AccessLog::Entry& entry, ::khulnasoft::EntryType type) {
  if (access_log_) {
    access_log_->Log(entry, type);
  }
}

Network::FilterStatus Instance::onNewConnection() {
  auto& conn = callbacks_->connection();
  ENVOY_CONN_LOG(debug, "khulnasoft.network: onNewConnection", conn);

  // Buffer data until proxylib policy is available, if configured with proxylib
  if (config_->proxylib_.get() != nullptr) {
    should_buffer_ = true;
  }

  const auto policy_fs =
      conn.streamInfo().filterState()->getDataReadOnly<Khulnasoft::KhulnasoftPolicyFilterState>(
          Khulnasoft::KhulnasoftPolicyFilterState::key());

  if (!policy_fs) {
    ENVOY_CONN_LOG(warn, "khulnasoft.network: Khulnasoft policy filter state not found", conn);
    return Network::FilterStatus::StopIteration;
  }

  // Default to incoming destination port, may be changed for L7 LB
  destination_port_ = policy_fs->port_;

  // Pass SNI before the upstream callback so that it is available when upstream connection is
  // initialized.
  const auto sni = conn.requestedServerName();
  if (sni != "") {
    ENVOY_CONN_LOG(trace, "khulnasoft.network: SNI: {}", conn, sni);
  }

  // Pass metadata from tls_inspector to the filterstate, if any & not already
  // set via upstream cluster config.
  if (sni != "") {
    auto filterState = conn.streamInfo().filterState();
    auto have_sni =
        filterState->hasData<Network::UpstreamServerName>(Network::UpstreamServerName::key());
    auto have_san = filterState->hasData<Network::UpstreamSubjectAltNames>(
        Network::UpstreamSubjectAltNames::key());
    if (!have_sni && !have_san) {
      filterState->setData(Network::UpstreamServerName::key(),
                           std::make_unique<Network::UpstreamServerName>(sni),
                           StreamInfo::FilterState::StateType::Mutable);
      filterState->setData(Network::UpstreamSubjectAltNames::key(),
                           std::make_unique<Network::UpstreamSubjectAltNames>(
                               std::vector<std::string>{std::string(sni)}),
                           StreamInfo::FilterState::StateType::Mutable);
    }
  }

  callbacks_->addUpstreamCallback([this, policy_fs,
                                   sni](Upstream::HostDescriptionConstSharedPtr host,
                                        StreamInfo::StreamInfo& stream_info) -> bool {
    // Skip enforcement or logging on shadows
    if (stream_info.isShadow()) {
      return true;
    }

    auto& conn = callbacks_->connection();
    ENVOY_CONN_LOG(trace, "khulnasoft.network: in upstream callback", conn);

    // Resolve the destination security ID and port
    uint32_t destination_identity = 0;

    Network::Address::InstanceConstSharedPtr dst_address =
        policy_fs->policyUseUpstreamDestinationAddress()
            ? host->address()
            : stream_info.downstreamAddressProvider().localAddress();
    if (nullptr == dst_address) {
      ENVOY_CONN_LOG(warn, "khulnasoft.network (egress): No destination address", conn);
      return false;
    }
    const auto dip = dst_address->ip();
    if (!dip) {
      ENVOY_CONN_LOG(warn, "khulnasoft.network: Non-IP destination address: {}", conn,
                     dst_address->asString());
      return false;
    }

    if (policy_fs->ingress_) {
      remote_id_ = policy_fs->source_identity_;
    } else {
      remote_id_ = destination_identity;
      destination_port_ = dip->port();
      destination_identity = policy_fs->resolvePolicyId(dip);
    }

    log_entry_.InitFromConnection(policy_fs->pod_ip_, policy_fs->proxy_id_, policy_fs->ingress_,
                                  policy_fs->source_identity_,
                                  stream_info.downstreamAddressProvider().remoteAddress(),
                                  destination_identity, dst_address, &config_->time_source_);

    bool useProxyLib;
    if (!policy_fs->enforceNetworkPolicy(conn, destination_identity, destination_port_, sni,
                                         useProxyLib, l7proto_, log_entry_)) {
      ENVOY_CONN_LOG(debug, "khulnasoft.network: policy DENY on id: {} port: {} sni: \"{}\"", conn,
                     remote_id_, destination_port_, sni);
      config_->Log(log_entry_, ::khulnasoft::EntryType::Denied);
      return false;
    }
    // Emit accesslog if north/south l7 lb, as in that case the traffic is not going back to bpf
    // datapath for policy enforcement
    if (log_entry_.entry_.policy_name() != policy_fs->pod_ip_) {
      config_->Log(log_entry_, ::khulnasoft::EntryType::Request);
    }
    ENVOY_LOG(debug, "khulnasoft.network: policy ALLOW on id: {} port: {} sni: \"{}\"", remote_id_,
              destination_port_, sni);

    if (useProxyLib) {
      const std::string& policy_name = policy_fs->pod_ip_;

      // Initialize Go parser if requested
      if (config_->proxylib_.get() != nullptr) {
        go_parser_ = config_->proxylib_->NewInstance(
            conn, l7proto_, policy_fs->ingress_, policy_fs->source_identity_, destination_identity,
            stream_info.downstreamAddressProvider().remoteAddress()->asString(),
            dst_address->asString(), policy_name);
        if (go_parser_.get() == nullptr) {
          ENVOY_CONN_LOG(warn, "khulnasoft.network: Go parser \"{}\" not found", conn, l7proto_);
          return false;
        }
      }
    }
    should_buffer_ = false;
    return true;
  });

  return Network::FilterStatus::Continue;
}

Network::FilterStatus Instance::onData(Buffer::Instance& data, bool end_stream) {
  auto& conn = callbacks_->connection();
  ENVOY_CONN_LOG(trace, "khulnasoft.network: onData {} bytes, end_stream: {}", conn, data.length(),
                 end_stream);
  const char* reason;

  if (should_buffer_) {
    // Buffer data until upstream is selected and policy resolved
    buffer_.move(data);
    return Network::FilterStatus::Continue;
  }
  // Prepend buffered data if any
  if (buffer_.length() > 0) {
    data.prepend(buffer_);
  }
  if (go_parser_) {
    FilterResult res =
        go_parser_->OnIO(false, data, end_stream); // 'false' marks original direction data
    ENVOY_CONN_LOG(trace, "khulnasoft.network::onData: \'GoFilter::OnIO\' returned {}", conn,
                   Envoy::Khulnasoft::toString(res));

    if (res != FILTER_OK) {
      // Drop the connection due to an error
      go_parser_->Close();
      reason = "proxylib error";
      goto drop_close;
    }

    if (go_parser_->WantReplyInject()) {
      ENVOY_CONN_LOG(trace, "khulnasoft.network::onData: calling write() on an empty buffer", conn);

      // We have no idea when, if ever new data will be received on the
      // reverse direction. Connection write on an empty buffer will cause
      // write filter chain to be called, and gives our write path the
      // opportunity to inject data.
      Buffer::OwnedImpl empty;
      conn.write(empty, false);
    }

    go_parser_->SetOrigEndStream(end_stream);
  } else if (!l7proto_.empty()) {
    const auto& metadata = conn.streamInfo().dynamicMetadata();
    bool changed = log_entry_.UpdateFromMetadata(l7proto_, metadata.filter_metadata().at(l7proto_));

    // Policy may have changed since the connection was established, get fresh policy
    const auto policy_fs =
        conn.streamInfo().filterState()->getDataReadOnly<Khulnasoft::KhulnasoftPolicyFilterState>(
            Khulnasoft::KhulnasoftPolicyFilterState::key());

    if (!policy_fs) {
      ENVOY_CONN_LOG(warn,
                     "khulnasoft.network: Khulnasoft policy filter state not found for pod {}, "
                     "defaulting to DENY",
                     conn, policy_fs->pod_ip_);
      reason = "Khulnasoft metadata lost";
      goto drop_close;
    }
    const auto& policy = policy_fs->getPolicy();
    auto port_policy = policy.findPortPolicy(policy_fs->ingress_, destination_port_);
    if (!port_policy.allowed(remote_id_, metadata)) {
      config_->Log(log_entry_, ::khulnasoft::EntryType::Denied);
      reason = "metadata policy drop";
      goto drop_close;
    } else {
      // accesslog only if metadata has changed
      if (changed) {
        config_->Log(log_entry_, ::khulnasoft::EntryType::Request);
      }
    }
  }

  return Network::FilterStatus::Continue;

drop_close:
  conn.close(Network::ConnectionCloseType::NoFlush, reason);
  return Network::FilterStatus::StopIteration;
}

Network::FilterStatus Instance::onWrite(Buffer::Instance& data, bool end_stream) {
  if (go_parser_) {
    FilterResult res =
        go_parser_->OnIO(true, data, end_stream); // 'true' marks reverse direction data
    ENVOY_CONN_LOG(trace, "khulnasoft.network::OnWrite: \'GoFilter::OnIO\' returned {}",
                   callbacks_->connection(), Envoy::Khulnasoft::toString(res));

    if (res != FILTER_OK) {
      // Drop the connection due to an error
      go_parser_->Close();
      return Network::FilterStatus::StopIteration;
    }

    // XXX: Unfortunately continueReading() continues from the next filter, and
    // there seems to be no way to trigger the whole filter chain to be called.

    go_parser_->SetReplyEndStream(end_stream);
  }

  return Network::FilterStatus::Continue;
}

} // namespace KhulnasoftL3
} // namespace Filter
} // namespace Envoy
