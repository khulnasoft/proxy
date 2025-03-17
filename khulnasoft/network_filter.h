#pragma once

#include <cstdint>
#include <memory>
#include <string>

#include "envoy/buffer/buffer.h"
#include "envoy/common/time.h"
#include "envoy/json/json_object.h"
#include "envoy/network/filter.h"
#include "envoy/server/factory_context.h"

#include "source/common/buffer/buffer_impl.h"
#include "source/common/common/logger.h"

#include "khulnasoft/accesslog.h"
#include "khulnasoft/api/accesslog.pb.h"
#include "khulnasoft/api/network_filter.pb.h"
#include "khulnasoft/proxylib.h"

namespace Envoy {
namespace Filter {
namespace KhulnasoftL3 {

/**
 * Shared configuration for Khulnasoft network filter worker
 * Instances. Each new network connection (on each worker thread)
 * get's their own Instance, but they all share a common Config for
 * any given filter chain.
 */
class Config : Logger::Loggable<Logger::Id::config> {
public:
  Config(const ::khulnasoft::NetworkFilter& config, Server::Configuration::FactoryContext& context);
  Config(const Json::Object& config, Server::Configuration::FactoryContext& context);

  void Log(Khulnasoft::AccessLog::Entry&, ::khulnasoft::EntryType);

  Khulnasoft::GoFilterSharedPtr proxylib_;
  TimeSource& time_source_;

private:
  Khulnasoft::AccessLogSharedPtr access_log_;
};

typedef std::shared_ptr<Config> ConfigSharedPtr;

/**
 * Implementation of a Khulnasoft network filter.
 */
class Instance : public Network::Filter, Logger::Loggable<Logger::Id::filter> {
public:
  Instance(const ConfigSharedPtr& config) : config_(config) {}

  // Network::ReadFilter
  Network::FilterStatus onData(Buffer::Instance&, bool end_stream) override;
  Network::FilterStatus onNewConnection() override;
  void initializeReadFilterCallbacks(Network::ReadFilterCallbacks& callbacks) override {
    callbacks_ = &callbacks;
  }

  // Network::WriteFilter
  Network::FilterStatus onWrite(Buffer::Instance&, bool end_stream) override;

private:
  const ConfigSharedPtr config_;
  Network::ReadFilterCallbacks* callbacks_ = nullptr;
  uint32_t remote_id_ = 0;
  uint16_t destination_port_ = 0;
  std::string l7proto_{};
  bool should_buffer_ = false;
  Buffer::OwnedImpl buffer_; // Buffer for initial connection data
  Khulnasoft::GoFilter::InstancePtr go_parser_{};
  Khulnasoft::AccessLog::Entry log_entry_{};
};

} // namespace KhulnasoftL3
} // namespace Filter
} // namespace Envoy
