// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/extensions/transport_sockets/internal_upstream/v3/internal_upstream.proto

package internal_upstreamv3

import (
	v3 "github.com/khulnasoft/proxy/go/envoy/config/core/v3"
	v31 "github.com/khulnasoft/proxy/go/envoy/type/metadata/v3"
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

// Configuration for the internal upstream address. An internal address defines
// a loopback user space socket residing in the same proxy instance. This
// extension allows passing additional structured state across the user space
// socket in addition to the regular byte stream. The purpose is to facilitate
// communication between filters on the downstream and the upstream internal
// connections. All filter state objects that are shared with the upstream
// connection are also shared with the downstream internal connection using
// this transport socket.
type InternalUpstreamTransport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the metadata namespaces and values to insert into the downstream
	// internal connection dynamic metadata when an internal address is used as a
	// host. If the destination name is repeated across two metadata source
	// locations, and both locations contain the metadata with the given name,
	// then the latter in the list overrides the former.
	PassthroughMetadata []*InternalUpstreamTransport_MetadataValueSource `protobuf:"bytes,1,rep,name=passthrough_metadata,json=passthroughMetadata,proto3" json:"passthrough_metadata,omitempty"`
	// The underlying transport socket being wrapped.
	TransportSocket *v3.TransportSocket `protobuf:"bytes,3,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
}

func (x *InternalUpstreamTransport) Reset() {
	*x = InternalUpstreamTransport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalUpstreamTransport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalUpstreamTransport) ProtoMessage() {}

func (x *InternalUpstreamTransport) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalUpstreamTransport.ProtoReflect.Descriptor instead.
func (*InternalUpstreamTransport) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDescGZIP(), []int{0}
}

func (x *InternalUpstreamTransport) GetPassthroughMetadata() []*InternalUpstreamTransport_MetadataValueSource {
	if x != nil {
		return x.PassthroughMetadata
	}
	return nil
}

func (x *InternalUpstreamTransport) GetTransportSocket() *v3.TransportSocket {
	if x != nil {
		return x.TransportSocket
	}
	return nil
}

// Describes the location of the imported metadata value.
// If the metadata with the given name is not present at the source location,
// then no metadata is passed through for this particular instance.
type InternalUpstreamTransport_MetadataValueSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies what kind of metadata.
	Kind *v31.MetadataKind `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Name is the filter namespace used in the dynamic metadata.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *InternalUpstreamTransport_MetadataValueSource) Reset() {
	*x = InternalUpstreamTransport_MetadataValueSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalUpstreamTransport_MetadataValueSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalUpstreamTransport_MetadataValueSource) ProtoMessage() {}

func (x *InternalUpstreamTransport_MetadataValueSource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalUpstreamTransport_MetadataValueSource.ProtoReflect.Descriptor instead.
func (*InternalUpstreamTransport_MetadataValueSource) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDescGZIP(), []int{0, 0}
}

func (x *InternalUpstreamTransport_MetadataValueSource) GetKind() *v31.MetadataKind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *InternalUpstreamTransport_MetadataValueSource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto protoreflect.FileDescriptor

var file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x37, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03, 0x0a, 0x19, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x99, 0x01, 0x0a, 0x14, 0x70, 0x61, 0x73,
	0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x66, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76,
	0x33, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x13, 0x70, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x5a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x1a, 0x76, 0x0a, 0x13, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x69, 0x6e, 0x64, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xdc, 0x01, 0x0a, 0x45, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x76, 0x33, 0x42, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x70, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x72, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x33, 0x3b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x76, 0x33, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDescOnce sync.Once
	file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDescData = file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDesc
)

func file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDescGZIP() []byte {
	file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDescData)
	})
	return file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDescData
}

var file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_goTypes = []interface{}{
	(*InternalUpstreamTransport)(nil),                     // 0: envoy.extensions.transport_sockets.internal_upstream.v3.InternalUpstreamTransport
	(*InternalUpstreamTransport_MetadataValueSource)(nil), // 1: envoy.extensions.transport_sockets.internal_upstream.v3.InternalUpstreamTransport.MetadataValueSource
	(*v3.TransportSocket)(nil),                            // 2: envoy.config.core.v3.TransportSocket
	(*v31.MetadataKind)(nil),                              // 3: envoy.type.metadata.v3.MetadataKind
}
var file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.transport_sockets.internal_upstream.v3.InternalUpstreamTransport.passthrough_metadata:type_name -> envoy.extensions.transport_sockets.internal_upstream.v3.InternalUpstreamTransport.MetadataValueSource
	2, // 1: envoy.extensions.transport_sockets.internal_upstream.v3.InternalUpstreamTransport.transport_socket:type_name -> envoy.config.core.v3.TransportSocket
	3, // 2: envoy.extensions.transport_sockets.internal_upstream.v3.InternalUpstreamTransport.MetadataValueSource.kind:type_name -> envoy.type.metadata.v3.MetadataKind
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_init()
}
func file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_init() {
	if File_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalUpstreamTransport); i {
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
		file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalUpstreamTransport_MetadataValueSource); i {
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
			RawDescriptor: file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_msgTypes,
	}.Build()
	File_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto = out.File
	file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_rawDesc = nil
	file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_goTypes = nil
	file_envoy_extensions_transport_sockets_internal_upstream_v3_internal_upstream_proto_depIdxs = nil
}
