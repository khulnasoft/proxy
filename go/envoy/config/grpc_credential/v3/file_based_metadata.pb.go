// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/config/grpc_credential/v3/file_based_metadata.proto

package grpc_credentialv3

import (
	v3 "github.com/khulnasoft/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
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

type FileBasedMetadataConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Location or inline data of secret to use for authentication of the Google gRPC connection
	// this secret will be attached to a header of the gRPC connection
	SecretData *v3.DataSource `protobuf:"bytes,1,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// Metadata header key to use for sending the secret data
	// if no header key is set, "authorization" header will be used
	HeaderKey string `protobuf:"bytes,2,opt,name=header_key,json=headerKey,proto3" json:"header_key,omitempty"`
	// Prefix to prepend to the secret in the metadata header
	// if no prefix is set, the default is to use no prefix
	HeaderPrefix string `protobuf:"bytes,3,opt,name=header_prefix,json=headerPrefix,proto3" json:"header_prefix,omitempty"`
}

func (x *FileBasedMetadataConfig) Reset() {
	*x = FileBasedMetadataConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_grpc_credential_v3_file_based_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileBasedMetadataConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileBasedMetadataConfig) ProtoMessage() {}

func (x *FileBasedMetadataConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_grpc_credential_v3_file_based_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileBasedMetadataConfig.ProtoReflect.Descriptor instead.
func (*FileBasedMetadataConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *FileBasedMetadataConfig) GetSecretData() *v3.DataSource {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *FileBasedMetadataConfig) GetHeaderKey() string {
	if x != nil {
		return x.HeaderKey
	}
	return ""
}

func (x *FileBasedMetadataConfig) GetHeaderPrefix() string {
	if x != nil {
		return x.HeaderPrefix
	}
	return ""
}

var File_envoy_config_grpc_credential_v3_file_based_metadata_proto protoreflect.FileDescriptor

var file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDesc = []byte{
	0x0a, 0x39, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x76,
	0x33, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xed, 0x01, 0x0a, 0x17, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x64, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x49, 0x0a,
	0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x42, 0x06, 0xb8, 0xb7, 0x8b, 0xa4, 0x02, 0x01, 0x52, 0x0a, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x3a, 0x43, 0x9a, 0xc5,
	0x88, 0x1e, 0x3e, 0x0a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x61,
	0x73, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0xab, 0x01, 0x0a, 0x2d, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x2e, 0x76, 0x33, 0x42, 0x16, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x64, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x2f, 0x76, 0x33, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDescOnce sync.Once
	file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDescData = file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDesc
)

func file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDescGZIP() []byte {
	file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDescOnce.Do(func() {
		file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDescData)
	})
	return file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDescData
}

var file_envoy_config_grpc_credential_v3_file_based_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_grpc_credential_v3_file_based_metadata_proto_goTypes = []interface{}{
	(*FileBasedMetadataConfig)(nil), // 0: envoy.config.grpc_credential.v3.FileBasedMetadataConfig
	(*v3.DataSource)(nil),           // 1: envoy.config.core.v3.DataSource
}
var file_envoy_config_grpc_credential_v3_file_based_metadata_proto_depIdxs = []int32{
	1, // 0: envoy.config.grpc_credential.v3.FileBasedMetadataConfig.secret_data:type_name -> envoy.config.core.v3.DataSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_grpc_credential_v3_file_based_metadata_proto_init() }
func file_envoy_config_grpc_credential_v3_file_based_metadata_proto_init() {
	if File_envoy_config_grpc_credential_v3_file_based_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_grpc_credential_v3_file_based_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileBasedMetadataConfig); i {
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
			RawDescriptor: file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_grpc_credential_v3_file_based_metadata_proto_goTypes,
		DependencyIndexes: file_envoy_config_grpc_credential_v3_file_based_metadata_proto_depIdxs,
		MessageInfos:      file_envoy_config_grpc_credential_v3_file_based_metadata_proto_msgTypes,
	}.Build()
	File_envoy_config_grpc_credential_v3_file_based_metadata_proto = out.File
	file_envoy_config_grpc_credential_v3_file_based_metadata_proto_rawDesc = nil
	file_envoy_config_grpc_credential_v3_file_based_metadata_proto_goTypes = nil
	file_envoy_config_grpc_credential_v3_file_based_metadata_proto_depIdxs = nil
}
