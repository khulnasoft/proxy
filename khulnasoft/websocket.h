#pragma once

#include "envoy/buffer/buffer.h"
#include "envoy/event/schedulable_cb.h"
#include "envoy/http/header_map.h"
#include "envoy/network/address.h"
#include "envoy/network/filter.h"
#include "envoy/stats/stats_macros.h" // IWYU pragma: keep

#include "source/common/common/logger.h"

#include "khulnasoft/accesslog.h"
#include "khulnasoft/api/accesslog.pb.h"
#include "khulnasoft/websocket_codec.h"
#include "khulnasoft/websocket_config.h"

namespace Envoy {
namespace Khulnasoft {
namespace WebSocket {

class Instance : public Network::Filter,
                 public CodecCallbacks,
                 Logger::Loggable<Logger::Id::filter> {
public:
  Instance(const ConfigSharedPtr& config) : config_(config) {}

  // Network::ReadFilter
  void initializeReadFilterCallbacks(Network::ReadFilterCallbacks& callbacks) override;
  Network::FilterStatus onNewConnection() override;
  Network::FilterStatus onData(Buffer::Instance&, bool end_stream) override;

  // Network::WriteFilter
  void initializeWriteFilterCallbacks(Network::WriteFilterCallbacks& callbacks) override {
    write_callbacks_ = &callbacks;
  }
  Network::FilterStatus onWrite(Buffer::Instance&, bool end_stream) override;

  // WebSocket::CodecCallbacks
  const ConfigSharedPtr& config() override { return config_; }
  void onHandshakeCreated(const Http::RequestHeaderMap& headers) override {
    log_entry_.UpdateFromRequest(0, nullptr, headers);
  }
  void onHandshakeSent() override { config_->Log(log_entry_, ::khulnasoft::EntryType::Request); }
  void onHandshakeRequest(const Http::RequestHeaderMap& headers) override;
  void onHandshakeResponse(const Http::ResponseHeaderMap& headers) override {
    log_entry_.UpdateFromResponse(headers, config_->time_source_);
    config_->Log(log_entry_, ::khulnasoft::EntryType::Response);
  }
  void onHandshakeResponseSent(const Http::ResponseHeaderMap& headers) override {
    bool accepted = headers.Status() && headers.getStatusValue() == "101";
    if (accepted) {
      config_->Log(log_entry_, ::khulnasoft::EntryType::Request);
    } else {
      config_->Log(log_entry_, ::khulnasoft::EntryType::Denied);
      config_->stats_.access_denied_.inc();
    }
    log_entry_.UpdateFromResponse(headers, config_->time_source_);
    config_->Log(log_entry_, ::khulnasoft::EntryType::Response);
  }

  void injectEncoded(Buffer::Instance& data, bool end_stream) override;
  void injectDecoded(Buffer::Instance& data, bool end_stream) override;

  void
  setOriginalDestinationAddress(const Network::Address::InstanceConstSharedPtr& orig_dst) override {
    callbacks_->connection().connectionInfoSetter().restoreLocalAddress(orig_dst);
  }

private:
  const ConfigSharedPtr config_;

  Network::ReadFilterCallbacks* callbacks_{nullptr};
  Network::WriteFilterCallbacks* write_callbacks_{nullptr};
  CodecPtr codec_{nullptr};
  Event::SchedulableCallbackPtr client_handshake_cb_{nullptr};
  Khulnasoft::AccessLog::Entry log_entry_{};
};

} // namespace WebSocket
} // namespace Khulnasoft
} // namespace Envoy
