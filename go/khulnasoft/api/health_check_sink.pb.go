// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: khulnasoft/api/health_check_sink.proto

package khulnasoft

import (
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

// Health check event pipe sink.
// The health check event will be streamed as binary protobufs.
type HealthCheckEventPipeSink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix domain socket path where to connect to send health check events to.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *HealthCheckEventPipeSink) Reset() {
	*x = HealthCheckEventPipeSink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_khulnasoft_api_health_check_sink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckEventPipeSink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckEventPipeSink) ProtoMessage() {}

func (x *HealthCheckEventPipeSink) ProtoReflect() protoreflect.Message {
	mi := &file_khulnasoft_api_health_check_sink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckEventPipeSink.ProtoReflect.Descriptor instead.
func (*HealthCheckEventPipeSink) Descriptor() ([]byte, []int) {
	return file_khulnasoft_api_health_check_sink_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheckEventPipeSink) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_khulnasoft_api_health_check_sink_proto protoreflect.FileDescriptor

var file_khulnasoft_api_health_check_sink_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x18, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x69, 0x70, 0x65, 0x53, 0x69, 0x6e,
	0x6b, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x2e,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x6c,
	0x69, 0x75, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x6c,
	0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_khulnasoft_api_health_check_sink_proto_rawDescOnce sync.Once
	file_khulnasoft_api_health_check_sink_proto_rawDescData = file_khulnasoft_api_health_check_sink_proto_rawDesc
)

func file_khulnasoft_api_health_check_sink_proto_rawDescGZIP() []byte {
	file_khulnasoft_api_health_check_sink_proto_rawDescOnce.Do(func() {
		file_khulnasoft_api_health_check_sink_proto_rawDescData = protoimpl.X.CompressGZIP(file_khulnasoft_api_health_check_sink_proto_rawDescData)
	})
	return file_khulnasoft_api_health_check_sink_proto_rawDescData
}

var file_khulnasoft_api_health_check_sink_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_khulnasoft_api_health_check_sink_proto_goTypes = []interface{}{
	(*HealthCheckEventPipeSink)(nil), // 0: khulnasoft.HealthCheckEventPipeSink
}
var file_khulnasoft_api_health_check_sink_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_khulnasoft_api_health_check_sink_proto_init() }
func file_khulnasoft_api_health_check_sink_proto_init() {
	if File_khulnasoft_api_health_check_sink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_khulnasoft_api_health_check_sink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckEventPipeSink); i {
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
			RawDescriptor: file_khulnasoft_api_health_check_sink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_khulnasoft_api_health_check_sink_proto_goTypes,
		DependencyIndexes: file_khulnasoft_api_health_check_sink_proto_depIdxs,
		MessageInfos:      file_khulnasoft_api_health_check_sink_proto_msgTypes,
	}.Build()
	File_khulnasoft_api_health_check_sink_proto = out.File
	file_khulnasoft_api_health_check_sink_proto_rawDesc = nil
	file_khulnasoft_api_health_check_sink_proto_goTypes = nil
	file_khulnasoft_api_health_check_sink_proto_depIdxs = nil
}
