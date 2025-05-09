#include "ipcache.h"

#include <arpa/inet.h>
#include <netinet/in.h>

#include <cstdint>
#include <cstring>
#include <memory>
#include <string>

#include "envoy/common/platform.h"
#include "envoy/network/address.h"
#include "envoy/server/factory_context.h"
#include "envoy/singleton/manager.h"

#include "source/common/common/logger.h"
#include "source/common/common/utility.h"

#include "absl/numeric/int128.h"
#include "khulnasoft/bpf.h"
#include "linux/bpf.h"
#include "linux/type_mapper.h"

namespace Envoy {
namespace Khulnasoft {

// IP cache names to look for, in the order of prefefrence
static const char* ipcache_names[] = {"khulnasoft_ipcache_v2", "khulnasoft_ipcache"};

// These must be kept in sync with Khulnasoft source code, should refactor
// them to a separate include file we can include here instead of
// copying them!

typedef uint32_t __be32; // Beware of the byte order!
typedef uint64_t __u64;
typedef uint32_t __u32;
typedef uint16_t __u16;
typedef uint8_t __u8;

PACKED_STRUCT(struct ipcache_key {
  struct bpf_lpm_trie_key lpm_key;
  __u16 pad1;
  __u8 pad2;
  __u8 family;
  union {
    struct {
      __u32 ip4;
      __u32 pad4;
      __u32 pad5;
      __u32 pad6;
    };
    __u32 ip6[4];
  };
});

struct remote_endpoint_info {
  using SecLabelType = __u32;
  SecLabelType sec_label;
  char buf[60]; // Enough space for all fields after the 'sec_label'
};

#define ENDPOINT_KEY_IPV4 1
#define ENDPOINT_KEY_IPV6 2

SINGLETON_MANAGER_REGISTRATION(khulnasoft_ipcache);

IPCacheSharedPtr IPCache::NewIPCache(Server::Configuration::ServerFactoryContext& context,
                                     const std::string& bpf_root) {
  return context.singletonManager().getTyped<Khulnasoft::IPCache>(
      SINGLETON_MANAGER_REGISTERED_NAME(khulnasoft_ipcache), [&bpf_root] {
        auto ipcache = std::make_shared<Khulnasoft::IPCache>(bpf_root);
        if (!ipcache->Open()) {
          ipcache.reset();
        }
        return ipcache;
      });
}

IPCacheSharedPtr IPCache::GetIPCache(Server::Configuration::ServerFactoryContext& context) {
  return context.singletonManager().getTyped<Khulnasoft::IPCache>(
      SINGLETON_MANAGER_REGISTERED_NAME(khulnasoft_ipcache));
}

IPCache::IPCache(const std::string& bpf_root)
    : Bpf(BPF_MAP_TYPE_LPM_TRIE, sizeof(struct ipcache_key),
          sizeof(remote_endpoint_info::SecLabelType), sizeof(struct remote_endpoint_info)),
      bpf_root_(bpf_root) {}

bool IPCache::Open() {
  // Open the bpf maps from Khulnasoft specific paths
  std::string tried_paths;

  for (const char* name : ipcache_names) {
    std::string path(bpf_root_ + "/tc/globals/" + name);
    if (!Bpf::open(path)) {
      if (tried_paths.length() > 0) {
        tried_paths += ", ";
      }
      tried_paths += path;
      continue;
    }
    ENVOY_LOG(debug, "khulnasoft.ipcache: Opened ipcache at {}", path);
    return true;
  }
  ENVOY_LOG(info, "khulnasoft.ipcache: Cannot open ipcache at any of {}", tried_paths);
  return false;
}

uint32_t IPCache::resolve(const Network::Address::Ip* ip) {
  struct ipcache_key key {};
  struct remote_endpoint_info value {};

  if (ip->version() == Network::Address::IpVersion::v4) {
    key.lpm_key = {32 + 32, {}};
    key.family = ENDPOINT_KEY_IPV4;
    key.ip4 = ip->ipv4()->address();
  } else {
    key.lpm_key = {32 + 128, {}};
    key.family = ENDPOINT_KEY_IPV6;
    absl::uint128 ip6 = ip->ipv6()->address();
    memcpy(&key.ip6, &ip6, sizeof key.ip6); // NOLINT(safe-memcpy)
  }

  if (key.family == ENDPOINT_KEY_IPV4) {
    ENVOY_LOG(trace, "khulnasoft.ipcache: Looking up key: {:x}, prefixlen: {}", ntohl(key.ip4),
              key.lpm_key.prefixlen - 32);
  } else if (key.family == ENDPOINT_KEY_IPV6) {
    ENVOY_LOG(trace, "khulnasoft.ipcache: Looking up key: {:x}:{:x}:{:x}:{:x}, prefixlen {}",
              ntohl(key.ip6[0]), ntohl(key.ip6[1]), ntohl(key.ip6[2]), ntohl(key.ip6[3]),
              key.lpm_key.prefixlen - 32);
  }

  if (lookup(&key, &value)) {
    ENVOY_LOG(debug, "khulnasoft.ipcache: {} has ID {}", ip->addressAsString(), value.sec_label);
    return value.sec_label;
  }
  ENVOY_LOG(info, "khulnasoft.ipcache: bpf map lookup failed: {}", Envoy::errorDetails(errno));
  return 0;
}

} // namespace Khulnasoft
} // namespace Envoy
