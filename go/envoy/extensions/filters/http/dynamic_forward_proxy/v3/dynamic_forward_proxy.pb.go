// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/extensions/filters/http/dynamic_forward_proxy/v3/dynamic_forward_proxy.proto

package dynamic_forward_proxyv3

import (
	v3 "github.com/khulnasoft/proxy/go/envoy/extensions/common/dynamic_forward_proxy/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// Configuration for the dynamic forward proxy HTTP filter. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#extension: envoy.filters.http.dynamic_forward_proxy]
type FilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ImplementationSpecifier:
	//
	//	*FilterConfig_DnsCacheConfig
	//	*FilterConfig_SubClusterConfig
	ImplementationSpecifier isFilterConfig_ImplementationSpecifier `protobuf_oneof:"implementation_specifier"`
	// When this flag is set, the filter will add the resolved upstream address in the filter
	// state. The state should be saved with key
	// “envoy.stream.upstream_address“ (See
	// :repo:`upstream_address.h<source/common/stream_info/upstream_address.h>`).
	SaveUpstreamAddress bool `protobuf:"varint,2,opt,name=save_upstream_address,json=saveUpstreamAddress,proto3" json:"save_upstream_address,omitempty"`
}

func (x *FilterConfig) Reset() {
	*x = FilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig) ProtoMessage() {}

func (x *FilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig.ProtoReflect.Descriptor instead.
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescGZIP(), []int{0}
}

func (m *FilterConfig) GetImplementationSpecifier() isFilterConfig_ImplementationSpecifier {
	if m != nil {
		return m.ImplementationSpecifier
	}
	return nil
}

func (x *FilterConfig) GetDnsCacheConfig() *v3.DnsCacheConfig {
	if x, ok := x.GetImplementationSpecifier().(*FilterConfig_DnsCacheConfig); ok {
		return x.DnsCacheConfig
	}
	return nil
}

func (x *FilterConfig) GetSubClusterConfig() *SubClusterConfig {
	if x, ok := x.GetImplementationSpecifier().(*FilterConfig_SubClusterConfig); ok {
		return x.SubClusterConfig
	}
	return nil
}

func (x *FilterConfig) GetSaveUpstreamAddress() bool {
	if x != nil {
		return x.SaveUpstreamAddress
	}
	return false
}

type isFilterConfig_ImplementationSpecifier interface {
	isFilterConfig_ImplementationSpecifier()
}

type FilterConfig_DnsCacheConfig struct {
	// The DNS cache configuration that the filter will attach to. Note this configuration must
	// match that of associated :ref:`dynamic forward proxy cluster configuration
	// <envoy_v3_api_field_extensions.clusters.dynamic_forward_proxy.v3.ClusterConfig.dns_cache_config>`.
	DnsCacheConfig *v3.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3,oneof"`
}

type FilterConfig_SubClusterConfig struct {
	// The configuration that the filter will use, when the related dynamic forward proxy cluster enabled
	// sub clusters.
	SubClusterConfig *SubClusterConfig `protobuf:"bytes,3,opt,name=sub_cluster_config,json=subClusterConfig,proto3,oneof"`
}

func (*FilterConfig_DnsCacheConfig) isFilterConfig_ImplementationSpecifier() {}

func (*FilterConfig_SubClusterConfig) isFilterConfig_ImplementationSpecifier() {}

// Per route Configuration for the dynamic forward proxy HTTP filter.
type PerRouteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to HostRewriteSpecifier:
	//
	//	*PerRouteConfig_HostRewriteLiteral
	//	*PerRouteConfig_HostRewriteHeader
	HostRewriteSpecifier isPerRouteConfig_HostRewriteSpecifier `protobuf_oneof:"host_rewrite_specifier"`
}

func (x *PerRouteConfig) Reset() {
	*x = PerRouteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerRouteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerRouteConfig) ProtoMessage() {}

func (x *PerRouteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerRouteConfig.ProtoReflect.Descriptor instead.
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescGZIP(), []int{1}
}

func (m *PerRouteConfig) GetHostRewriteSpecifier() isPerRouteConfig_HostRewriteSpecifier {
	if m != nil {
		return m.HostRewriteSpecifier
	}
	return nil
}

func (x *PerRouteConfig) GetHostRewriteLiteral() string {
	if x, ok := x.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewriteLiteral); ok {
		return x.HostRewriteLiteral
	}
	return ""
}

func (x *PerRouteConfig) GetHostRewriteHeader() string {
	if x, ok := x.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewriteHeader); ok {
		return x.HostRewriteHeader
	}
	return ""
}

type isPerRouteConfig_HostRewriteSpecifier interface {
	isPerRouteConfig_HostRewriteSpecifier()
}

type PerRouteConfig_HostRewriteLiteral struct {
	// Indicates that before DNS lookup, the host header will be swapped with
	// this value. If not set or empty, the original host header value
	// will be used and no rewrite will happen.
	//
	// Note: this rewrite affects both DNS lookup and host header forwarding. However, this
	// option shouldn't be used with
	// :ref:`HCM host rewrite <envoy_v3_api_field_config.route.v3.RouteAction.host_rewrite_literal>` given that the
	// value set here would be used for DNS lookups whereas the value set in the HCM would be used
	// for host header forwarding which is not the desired outcome.
	HostRewriteLiteral string `protobuf:"bytes,1,opt,name=host_rewrite_literal,json=hostRewriteLiteral,proto3,oneof"`
}

type PerRouteConfig_HostRewriteHeader struct {
	// Indicates that before DNS lookup, the host header will be swapped with
	// the value of this header. If not set or empty, the original host header
	// value will be used and no rewrite will happen.
	//
	// Note: this rewrite affects both DNS lookup and host header forwarding. However, this
	// option shouldn't be used with
	// :ref:`HCM host rewrite header <envoy_v3_api_field_config.route.v3.RouteAction.auto_host_rewrite>`
	// given that the value set here would be used for DNS lookups whereas the value set in the HCM
	// would be used for host header forwarding which is not the desired outcome.
	//
	// .. note::
	//
	//	If the header appears multiple times only the first value is used.
	HostRewriteHeader string `protobuf:"bytes,2,opt,name=host_rewrite_header,json=hostRewriteHeader,proto3,oneof"`
}

func (*PerRouteConfig_HostRewriteLiteral) isPerRouteConfig_HostRewriteSpecifier() {}

func (*PerRouteConfig_HostRewriteHeader) isPerRouteConfig_HostRewriteSpecifier() {}

type SubClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The timeout used for sub cluster initialization. Defaults to 5s if not set.
	ClusterInitTimeout *durationpb.Duration `protobuf:"bytes,3,opt,name=cluster_init_timeout,json=clusterInitTimeout,proto3" json:"cluster_init_timeout,omitempty"`
}

func (x *SubClusterConfig) Reset() {
	*x = SubClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubClusterConfig) ProtoMessage() {}

func (x *SubClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubClusterConfig.ProtoReflect.Descriptor instead.
func (*SubClusterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *SubClusterConfig) GetClusterInitTimeout() *durationpb.Duration {
	if x != nil {
		return x.ClusterInitTimeout
	}
	return nil
}

var File_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDesc = []byte{
	0x0a, 0x52, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x40, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64,
	0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x03, 0x0a, 0x0c, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x6c, 0x0a, 0x10, 0x64, 0x6e,
	0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x6e, 0x73, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x78, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x5f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x75,
	0x62, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00,
	0x52, 0x10, 0x73, 0x75, 0x62, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x73, 0x61, 0x76, 0x65, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x4a, 0x9a, 0xc5, 0x88, 0x1e, 0x45, 0x0a, 0x43, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x1a, 0x0a, 0x18, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xde,
	0x01, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x32, 0x0a, 0x14, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x12, 0x68, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69,
	0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x13, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x68, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x3a, 0x4c, 0x9a, 0xc5, 0x88, 0x1e, 0x47, 0x0a, 0x45,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x18, 0x0a, 0x16, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22,
	0x69, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x55, 0x0a, 0x14, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x12, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0xe1, 0x01, 0x0a, 0x44, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x33, 0x42, 0x18, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x75, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescData = file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDesc
)

func file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescData
}

var file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_goTypes = []interface{}{
	(*FilterConfig)(nil),        // 0: envoy.extensions.filters.http.dynamic_forward_proxy.v3.FilterConfig
	(*PerRouteConfig)(nil),      // 1: envoy.extensions.filters.http.dynamic_forward_proxy.v3.PerRouteConfig
	(*SubClusterConfig)(nil),    // 2: envoy.extensions.filters.http.dynamic_forward_proxy.v3.SubClusterConfig
	(*v3.DnsCacheConfig)(nil),   // 3: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
}
var file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.filters.http.dynamic_forward_proxy.v3.FilterConfig.dns_cache_config:type_name -> envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig
	2, // 1: envoy.extensions.filters.http.dynamic_forward_proxy.v3.FilterConfig.sub_cluster_config:type_name -> envoy.extensions.filters.http.dynamic_forward_proxy.v3.SubClusterConfig
	4, // 2: envoy.extensions.filters.http.dynamic_forward_proxy.v3.SubClusterConfig.cluster_init_timeout:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_init()
}
func file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_init() {
	if File_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig); i {
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
		file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerRouteConfig); i {
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
		file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubClusterConfig); i {
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
	file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FilterConfig_DnsCacheConfig)(nil),
		(*FilterConfig_SubClusterConfig)(nil),
	}
	file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PerRouteConfig_HostRewriteLiteral)(nil),
		(*PerRouteConfig_HostRewriteHeader)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto = out.File
	file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDesc = nil
	file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_goTypes = nil
	file_envoy_extensions_filters_http_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_depIdxs = nil
}
