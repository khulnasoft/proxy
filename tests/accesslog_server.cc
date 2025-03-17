#include "tests/accesslog_server.h"

#include <unistd.h>

#include <chrono>
#include <functional>
#include <string>

#include "source/common/common/logger.h"

#include "absl/base/thread_annotations.h"
#include "absl/synchronization/mutex.h"
#include "absl/time/time.h"
#include "absl/types/optional.h"
#include "khulnasoft/api/accesslog.pb.h"
#include "tests/uds_server.h"

namespace Envoy {

AccessLogServer::AccessLogServer(const std::string path)
    : UDSServer(path, std::bind(&AccessLogServer::msgCallback, this, std::placeholders::_1)) {}

AccessLogServer::~AccessLogServer() {}

void AccessLogServer::clear() {
  absl::MutexLock lock(&mutex_);
  messages_.clear();
}

absl::optional<::khulnasoft::LogEntry>
AccessLogServer::waitForMessage(::khulnasoft::EntryType entry_type, std::chrono::milliseconds timeout) {
  absl::MutexLock lock(&mutex_);
  absl::optional<::khulnasoft::LogEntry> entry;
  auto predicate = [this, &entry, entry_type]() ABSL_SHARED_LOCKS_REQUIRED(mutex_) {
    mutex_.AssertHeld();
    for (auto& msg : messages_) {
      if (msg.entry_type() == entry_type) {
        entry = msg;
        return true;
      }
    }
    return false;
  };
  mutex_.AwaitWithTimeout(absl::Condition(&predicate), absl::Milliseconds(timeout.count()));
  return entry;
}

void AccessLogServer::msgCallback(const std::string& data) {
  ::khulnasoft::LogEntry entry;
  if (!entry.ParseFromString(data)) {
    ENVOY_LOG(warn, "Access log parse failed!");
  } else {
    ENVOY_LOG(info, "Access log entry: {}", entry.DebugString());
    absl::MutexLock lock(&mutex_);
    messages_.emplace_back(entry);
  }
}

} // namespace Envoy
