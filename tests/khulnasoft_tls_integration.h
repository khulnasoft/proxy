#pragma once

#include "envoy/api/api.h"
#include "envoy/network/transport_socket.h"
#include "envoy/ssl/context_manager.h"

namespace Envoy {
namespace Khulnasoft {

Network::UpstreamTransportSocketFactoryPtr
createClientSslTransportSocketFactory(Ssl::ContextManager& context_manager, Api::Api& api);

} // namespace Khulnasoft
} // namespace Envoy
