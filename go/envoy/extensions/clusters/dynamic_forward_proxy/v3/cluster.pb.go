// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/extensions/clusters/dynamic_forward_proxy/v3/cluster.proto

package dynamic_forward_proxyv3

import (
	v31 "github.com/khulnasoft/proxy/go/envoy/config/cluster/v3"
	v32 "github.com/khulnasoft/proxy/go/envoy/config/core/v3"
	v3 "github.com/khulnasoft/proxy/go/envoy/extensions/common/dynamic_forward_proxy/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for the dynamic forward proxy cluster. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#extension: envoy.clusters.dynamic_forward_proxy]
type ClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ClusterImplementationSpecifier:
	//
	//	*ClusterConfig_DnsCacheConfig
	//	*ClusterConfig_SubClustersConfig
	ClusterImplementationSpecifier isClusterConfig_ClusterImplementationSpecifier `protobuf_oneof:"cluster_implementation_specifier"`
	// If true allow the cluster configuration to disable the auto_sni and auto_san_validation options
	// in the :ref:`cluster's upstream_http_protocol_options
	// <envoy_v3_api_field_config.cluster.v3.Cluster.upstream_http_protocol_options>`
	AllowInsecureClusterOptions bool `protobuf:"varint,2,opt,name=allow_insecure_cluster_options,json=allowInsecureClusterOptions,proto3" json:"allow_insecure_cluster_options,omitempty"`
	// If true allow HTTP/2 and HTTP/3 connections to be reused for requests to different
	// origins than the connection was initially created for. This will only happen when the
	// resolved address for the new connection matches the peer address of the connection and
	// the TLS certificate is also valid for the new hostname. For example, if a connection
	// has previously been established to foo.example.com at IP 1.2.3.4 with a certificate
	// that is valid for “*.example.com“, then this connection could be used for requests to
	// bar.example.com if that also resolved to 1.2.3.4.
	//
	// .. note::
	//
	//	By design, this feature will maximize reuse of connections. This means that instead
	//	opening a new connection when an existing connection reaches the maximum number of
	//	concurrent streams, requests will instead be sent to the existing connection.
	//
	// .. note::
	//
	//	The coalesced connections might be to upstreams that would not be otherwise
	//	selected by Envoy. See the section `Connection Reuse in RFC 7540
	//	<https://datatracker.ietf.org/doc/html/rfc7540#section-9.1.1>`_
	AllowCoalescedConnections bool `protobuf:"varint,3,opt,name=allow_coalesced_connections,json=allowCoalescedConnections,proto3" json:"allow_coalesced_connections,omitempty"`
}

func (x *ClusterConfig) Reset() {
	*x = ClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterConfig) ProtoMessage() {}

func (x *ClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterConfig.ProtoReflect.Descriptor instead.
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescGZIP(), []int{0}
}

func (m *ClusterConfig) GetClusterImplementationSpecifier() isClusterConfig_ClusterImplementationSpecifier {
	if m != nil {
		return m.ClusterImplementationSpecifier
	}
	return nil
}

func (x *ClusterConfig) GetDnsCacheConfig() *v3.DnsCacheConfig {
	if x, ok := x.GetClusterImplementationSpecifier().(*ClusterConfig_DnsCacheConfig); ok {
		return x.DnsCacheConfig
	}
	return nil
}

func (x *ClusterConfig) GetSubClustersConfig() *SubClustersConfig {
	if x, ok := x.GetClusterImplementationSpecifier().(*ClusterConfig_SubClustersConfig); ok {
		return x.SubClustersConfig
	}
	return nil
}

func (x *ClusterConfig) GetAllowInsecureClusterOptions() bool {
	if x != nil {
		return x.AllowInsecureClusterOptions
	}
	return false
}

func (x *ClusterConfig) GetAllowCoalescedConnections() bool {
	if x != nil {
		return x.AllowCoalescedConnections
	}
	return false
}

type isClusterConfig_ClusterImplementationSpecifier interface {
	isClusterConfig_ClusterImplementationSpecifier()
}

type ClusterConfig_DnsCacheConfig struct {
	// The DNS cache configuration that the cluster will attach to. Note this configuration must
	// match that of associated :ref:`dynamic forward proxy HTTP filter configuration
	// <envoy_v3_api_field_extensions.filters.http.dynamic_forward_proxy.v3.FilterConfig.dns_cache_config>`.
	DnsCacheConfig *v3.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3,oneof"`
}

type ClusterConfig_SubClustersConfig struct {
	// Configuration for sub clusters, when this configuration is enabled,
	// Envoy will create an independent sub cluster dynamically for each host:port.
	// Most of the configuration of a sub cluster is inherited from the current cluster,
	// i.e. health_checks, dns_resolvers and etc.
	// And the load_assignment will be set to the only one endpoint, host:port.
	//
	// Compared to the dns_cache_config, it has the following advantages:
	//
	//  1. sub clusters will be created with the STRICT_DNS DiscoveryType,
	//     so that Envoy will use all of the IPs resolved from the host.
	//
	// 2. each sub cluster is full featured cluster, with lb_policy and health check and etc enabled.
	SubClustersConfig *SubClustersConfig `protobuf:"bytes,4,opt,name=sub_clusters_config,json=subClustersConfig,proto3,oneof"`
}

func (*ClusterConfig_DnsCacheConfig) isClusterConfig_ClusterImplementationSpecifier() {}

func (*ClusterConfig_SubClustersConfig) isClusterConfig_ClusterImplementationSpecifier() {}

// Configuration for sub clusters. Hard code STRICT_DNS cluster type now.
type SubClustersConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The :ref:`load balancer type <arch_overview_load_balancing_types>` to use
	// when picking a host in a sub cluster. Note that CLUSTER_PROVIDED is not allowed here.
	LbPolicy v31.Cluster_LbPolicy `protobuf:"varint,1,opt,name=lb_policy,json=lbPolicy,proto3,enum=envoy.config.cluster.v3.Cluster_LbPolicy" json:"lb_policy,omitempty"`
	// The maximum number of sub clusters that the DFP cluster will hold. If not specified defaults to 1024.
	MaxSubClusters *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=max_sub_clusters,json=maxSubClusters,proto3" json:"max_sub_clusters,omitempty"`
	// The TTL for sub clusters that are unused. Sub clusters that have not been used in the configured time
	// interval will be purged. If not specified defaults to 5m.
	SubClusterTtl *durationpb.Duration `protobuf:"bytes,3,opt,name=sub_cluster_ttl,json=subClusterTtl,proto3" json:"sub_cluster_ttl,omitempty"`
	// Sub clusters that should be created & warmed upon creation. This might provide a
	// performance improvement, in the form of cache hits, for sub clusters that are going to be
	// warmed during steady state and are known at config load time.
	PreresolveClusters []*v32.SocketAddress `protobuf:"bytes,4,rep,name=preresolve_clusters,json=preresolveClusters,proto3" json:"preresolve_clusters,omitempty"`
}

func (x *SubClustersConfig) Reset() {
	*x = SubClustersConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubClustersConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubClustersConfig) ProtoMessage() {}

func (x *SubClustersConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubClustersConfig.ProtoReflect.Descriptor instead.
func (*SubClustersConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *SubClustersConfig) GetLbPolicy() v31.Cluster_LbPolicy {
	if x != nil {
		return x.LbPolicy
	}
	return v31.Cluster_ROUND_ROBIN
}

func (x *SubClustersConfig) GetMaxSubClusters() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxSubClusters
	}
	return nil
}

func (x *SubClustersConfig) GetSubClusterTtl() *durationpb.Duration {
	if x != nil {
		return x.SubClusterTtl
	}
	return nil
}

func (x *SubClustersConfig) GetPreresolveClusters() []*v32.SocketAddress {
	if x != nil {
		return x.PreresolveClusters
	}
	return nil
}

var File_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto protoreflect.FileDescriptor

var file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x76, 0x33, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8,
	0x03, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x6c, 0x0a, 0x10, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e,
	0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0e,
	0x64, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x77,
	0x0a, 0x13, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33,
	0x2e, 0x53, 0x75, 0x62, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x48, 0x00, 0x52, 0x11, 0x73, 0x75, 0x62, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x1e, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x1b,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x65, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x19, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x47, 0x9a, 0xc5,
	0x88, 0x1e, 0x42, 0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x22, 0x0a, 0x20, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xd9, 0x02, 0x0a, 0x11, 0x53, 0x75,
	0x62, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x50, 0x0a, 0x09, 0x6c, 0x62, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x62, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6c, 0x62, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x4f, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02,
	0x20, 0x00, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x62, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x4b, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00,
	0x52, 0x0d, 0x73, 0x75, 0x62, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x54, 0x74, 0x6c, 0x12,
	0x54, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x12, 0x70, 0x72, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x42, 0xcd, 0x01, 0x0a, 0x40, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x71, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x76, 0x33, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescOnce sync.Once
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescData = file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDesc
)

func file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescGZIP() []byte {
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescData)
	})
	return file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescData
}

var file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_goTypes = []interface{}{
	(*ClusterConfig)(nil),          // 0: envoy.extensions.clusters.dynamic_forward_proxy.v3.ClusterConfig
	(*SubClustersConfig)(nil),      // 1: envoy.extensions.clusters.dynamic_forward_proxy.v3.SubClustersConfig
	(*v3.DnsCacheConfig)(nil),      // 2: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig
	(v31.Cluster_LbPolicy)(0),      // 3: envoy.config.cluster.v3.Cluster.LbPolicy
	(*wrapperspb.UInt32Value)(nil), // 4: google.protobuf.UInt32Value
	(*durationpb.Duration)(nil),    // 5: google.protobuf.Duration
	(*v32.SocketAddress)(nil),      // 6: envoy.config.core.v3.SocketAddress
}
var file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.clusters.dynamic_forward_proxy.v3.ClusterConfig.dns_cache_config:type_name -> envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig
	1, // 1: envoy.extensions.clusters.dynamic_forward_proxy.v3.ClusterConfig.sub_clusters_config:type_name -> envoy.extensions.clusters.dynamic_forward_proxy.v3.SubClustersConfig
	3, // 2: envoy.extensions.clusters.dynamic_forward_proxy.v3.SubClustersConfig.lb_policy:type_name -> envoy.config.cluster.v3.Cluster.LbPolicy
	4, // 3: envoy.extensions.clusters.dynamic_forward_proxy.v3.SubClustersConfig.max_sub_clusters:type_name -> google.protobuf.UInt32Value
	5, // 4: envoy.extensions.clusters.dynamic_forward_proxy.v3.SubClustersConfig.sub_cluster_ttl:type_name -> google.protobuf.Duration
	6, // 5: envoy.extensions.clusters.dynamic_forward_proxy.v3.SubClustersConfig.preresolve_clusters:type_name -> envoy.config.core.v3.SocketAddress
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_init() }
func file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_init() {
	if File_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubClustersConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ClusterConfig_DnsCacheConfig)(nil),
		(*ClusterConfig_SubClustersConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes,
	}.Build()
	File_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto = out.File
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDesc = nil
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_goTypes = nil
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_depIdxs = nil
}
