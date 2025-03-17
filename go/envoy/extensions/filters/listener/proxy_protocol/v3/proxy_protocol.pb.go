// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/extensions/filters/listener/proxy_protocol/v3/proxy_protocol.proto

package proxy_protocolv3

import (
	v3 "github.com/khulnasoft/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// [#next-free-field: 6]
type ProxyProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of rules to apply to requests.
	Rules []*ProxyProtocol_Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	// Allow requests through that don't use proxy protocol. Defaults to false.
	//
	// .. attention::
	//
	//	This breaks conformance with the specification.
	//	Only enable if ALL traffic to the listener comes from a trusted source.
	//	For more information on the security implications of this feature, see
	//	https://www.haproxy.org/download/2.1/doc/proxy-protocol.txt
	//
	// .. attention::
	//
	//	Requests of 12 or fewer bytes that match the proxy protocol v2 signature
	//	and requests of 6 or fewer bytes that match the proxy protocol v1
	//	signature will timeout (Envoy is unable to differentiate these requests
	//	from incomplete proxy protocol requests).
	AllowRequestsWithoutProxyProtocol bool `protobuf:"varint,2,opt,name=allow_requests_without_proxy_protocol,json=allowRequestsWithoutProxyProtocol,proto3" json:"allow_requests_without_proxy_protocol,omitempty"`
	// This config controls which TLVs can be passed to filter state if it is Proxy Protocol
	// V2 header. If there is no setting for this field, no TLVs will be passed through.
	//
	// .. note::
	//
	//	If this is configured, you likely also want to set
	//	:ref:`core.v3.ProxyProtocolConfig.pass_through_tlvs <envoy_v3_api_field_config.core.v3.ProxyProtocolConfig.pass_through_tlvs>`,
	//	which controls pass-through for the upstream.
	PassThroughTlvs *v3.ProxyProtocolPassThroughTLVs `protobuf:"bytes,3,opt,name=pass_through_tlvs,json=passThroughTlvs,proto3" json:"pass_through_tlvs,omitempty"`
	// The PROXY protocol versions that won't be matched. Useful to limit the scope and attack surface of the filter.
	//
	// When the filter receives PROXY protocol data that is disallowed, it will reject the connection.
	// By default, the filter will match all PROXY protocol versions.
	// See https://www.haproxy.org/download/2.1/doc/proxy-protocol.txt for details.
	//
	// .. attention::
	//
	//	When used in conjunction with the :ref:`allow_requests_without_proxy_protocol <envoy_v3_api_field_extensions.filters.listener.proxy_protocol.v3.ProxyProtocol.allow_requests_without_proxy_protocol>`,
	//	the filter will not attempt to match signatures for the disallowed versions.
	//	For example, when ``disallowed_versions=V2``, ``allow_requests_without_proxy_protocol=true``,
	//	and an incoming request matches the V2 signature, the filter will allow the request through without any modification.
	//	The filter treats this request as if it did not have any PROXY protocol information.
	DisallowedVersions []v3.ProxyProtocolConfig_Version `protobuf:"varint,4,rep,packed,name=disallowed_versions,json=disallowedVersions,proto3,enum=envoy.config.core.v3.ProxyProtocolConfig_Version" json:"disallowed_versions,omitempty"`
	// The human readable prefix to use when emitting statistics for the filter.
	// If not configured, statistics will be emitted without the prefix segment.
	// See the :ref:`filter's statistics documentation <config_listener_filters_proxy_protocol>` for
	// more information.
	StatPrefix string `protobuf:"bytes,5,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
}

func (x *ProxyProtocol) Reset() {
	*x = ProxyProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocol) ProtoMessage() {}

func (x *ProxyProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocol.ProtoReflect.Descriptor instead.
func (*ProxyProtocol) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyProtocol) GetRules() []*ProxyProtocol_Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *ProxyProtocol) GetAllowRequestsWithoutProxyProtocol() bool {
	if x != nil {
		return x.AllowRequestsWithoutProxyProtocol
	}
	return false
}

func (x *ProxyProtocol) GetPassThroughTlvs() *v3.ProxyProtocolPassThroughTLVs {
	if x != nil {
		return x.PassThroughTlvs
	}
	return nil
}

func (x *ProxyProtocol) GetDisallowedVersions() []v3.ProxyProtocolConfig_Version {
	if x != nil {
		return x.DisallowedVersions
	}
	return nil
}

func (x *ProxyProtocol) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

type ProxyProtocol_KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace — if this is empty, the filter's namespace will be used.
	MetadataNamespace string `protobuf:"bytes,1,opt,name=metadata_namespace,json=metadataNamespace,proto3" json:"metadata_namespace,omitempty"`
	// The key to use within the namespace.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ProxyProtocol_KeyValuePair) Reset() {
	*x = ProxyProtocol_KeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocol_KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocol_KeyValuePair) ProtoMessage() {}

func (x *ProxyProtocol_KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocol_KeyValuePair.ProtoReflect.Descriptor instead.
func (*ProxyProtocol_KeyValuePair) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProxyProtocol_KeyValuePair) GetMetadataNamespace() string {
	if x != nil {
		return x.MetadataNamespace
	}
	return ""
}

func (x *ProxyProtocol_KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// A Rule defines what metadata to apply when a header is present or missing.
type ProxyProtocol_Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type that triggers the rule - required
	// TLV type is defined as uint8_t in proxy protocol. See `the spec
	// <https://www.haproxy.org/download/2.1/doc/proxy-protocol.txt>`_ for details.
	TlvType uint32 `protobuf:"varint,1,opt,name=tlv_type,json=tlvType,proto3" json:"tlv_type,omitempty"`
	// If the TLV type is present, apply this metadata KeyValuePair.
	OnTlvPresent *ProxyProtocol_KeyValuePair `protobuf:"bytes,2,opt,name=on_tlv_present,json=onTlvPresent,proto3" json:"on_tlv_present,omitempty"`
}

func (x *ProxyProtocol_Rule) Reset() {
	*x = ProxyProtocol_Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocol_Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocol_Rule) ProtoMessage() {}

func (x *ProxyProtocol_Rule) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocol_Rule.ProtoReflect.Descriptor instead.
func (*ProxyProtocol_Rule) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ProxyProtocol_Rule) GetTlvType() uint32 {
	if x != nil {
		return x.TlvType
	}
	return 0
}

func (x *ProxyProtocol_Rule) GetOnTlvPresent() *ProxyProtocol_KeyValuePair {
	if x != nil {
		return x.OnTlvPresent
	}
	return nil
}

var File_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDesc = []byte{
	0x0a, 0x48, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x1a,
	0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x05, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x5d, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x25, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x21, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x5e, 0x0a, 0x11, 0x70, 0x61, 0x73, 0x73,
	0x5f, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x5f, 0x74, 0x6c, 0x76, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x54, 0x68, 0x72, 0x6f,
	0x75, 0x67, 0x68, 0x54, 0x4c, 0x56, 0x73, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x54, 0x68, 0x72,
	0x6f, 0x75, 0x67, 0x68, 0x54, 0x6c, 0x76, 0x73, 0x12, 0x62, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x1a, 0x58, 0x0a,
	0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x2d, 0x0a,
	0x12, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x1a, 0xa2, 0x01, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x23, 0x0a, 0x08, 0x74, 0x6c, 0x76, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x2a, 0x03, 0x10, 0x80, 0x02, 0x52, 0x07, 0x74, 0x6c,
	0x76, 0x54, 0x79, 0x70, 0x65, 0x12, 0x75, 0x0a, 0x0e, 0x6f, 0x6e, 0x5f, 0x74, 0x6c, 0x76, 0x5f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0c,
	0x6f, 0x6e, 0x54, 0x6c, 0x76, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x3a, 0x43, 0x9a, 0xc5,
	0x88, 0x1e, 0x3e, 0x0a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x42, 0xce, 0x01, 0x0a, 0x41, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x42, 0x12, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x33, 0x3b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDescData = file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDesc
)

func file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDescData)
	})
	return file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDescData
}

var file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_goTypes = []interface{}{
	(*ProxyProtocol)(nil),                   // 0: envoy.extensions.filters.listener.proxy_protocol.v3.ProxyProtocol
	(*ProxyProtocol_KeyValuePair)(nil),      // 1: envoy.extensions.filters.listener.proxy_protocol.v3.ProxyProtocol.KeyValuePair
	(*ProxyProtocol_Rule)(nil),              // 2: envoy.extensions.filters.listener.proxy_protocol.v3.ProxyProtocol.Rule
	(*v3.ProxyProtocolPassThroughTLVs)(nil), // 3: envoy.config.core.v3.ProxyProtocolPassThroughTLVs
	(v3.ProxyProtocolConfig_Version)(0),     // 4: envoy.config.core.v3.ProxyProtocolConfig.Version
}
var file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.listener.proxy_protocol.v3.ProxyProtocol.rules:type_name -> envoy.extensions.filters.listener.proxy_protocol.v3.ProxyProtocol.Rule
	3, // 1: envoy.extensions.filters.listener.proxy_protocol.v3.ProxyProtocol.pass_through_tlvs:type_name -> envoy.config.core.v3.ProxyProtocolPassThroughTLVs
	4, // 2: envoy.extensions.filters.listener.proxy_protocol.v3.ProxyProtocol.disallowed_versions:type_name -> envoy.config.core.v3.ProxyProtocolConfig.Version
	1, // 3: envoy.extensions.filters.listener.proxy_protocol.v3.ProxyProtocol.Rule.on_tlv_present:type_name -> envoy.extensions.filters.listener.proxy_protocol.v3.ProxyProtocol.KeyValuePair
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_init() }
func file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_init() {
	if File_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocol); i {
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
		file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocol_KeyValuePair); i {
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
		file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocol_Rule); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto = out.File
	file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_rawDesc = nil
	file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_goTypes = nil
	file_envoy_extensions_filters_listener_proxy_protocol_v3_proxy_protocol_proto_depIdxs = nil
}
