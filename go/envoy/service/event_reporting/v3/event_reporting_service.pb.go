// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/service/event_reporting/v3/event_reporting_service.proto

package event_reportingv3

import (
	context "context"
	v3 "github.com/khulnasoft/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// [#not-implemented-hide:]
// An events envoy sends to the management server.
type StreamEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier data that will only be sent in the first message on the stream. This is effectively
	// structured metadata and is a performance optimization.
	Identifier *StreamEventsRequest_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Batch of events. When the stream is already active, it will be the events occurred
	// since the last message had been sent. If the server receives unknown event type, it should
	// silently ignore it.
	//
	// The following events are supported:
	//
	// * :ref:`HealthCheckEvent <envoy_v3_api_msg_data.core.v3.HealthCheckEvent>`
	// * :ref:`OutlierDetectionEvent <envoy_v3_api_msg_data.cluster.v3.OutlierDetectionEvent>`
	Events []*anypb.Any `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *StreamEventsRequest) Reset() {
	*x = StreamEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEventsRequest) ProtoMessage() {}

func (x *StreamEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEventsRequest.ProtoReflect.Descriptor instead.
func (*StreamEventsRequest) Descriptor() ([]byte, []int) {
	return file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDescGZIP(), []int{0}
}

func (x *StreamEventsRequest) GetIdentifier() *StreamEventsRequest_Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *StreamEventsRequest) GetEvents() []*anypb.Any {
	if x != nil {
		return x.Events
	}
	return nil
}

// [#not-implemented-hide:]
// The management server may send envoy a StreamEventsResponse to tell which events the server
// is interested in. In future, with aggregated event reporting service, this message will
// contain, for example, clusters the envoy should send events for, or event types the server
// wants to process.
type StreamEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamEventsResponse) Reset() {
	*x = StreamEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEventsResponse) ProtoMessage() {}

func (x *StreamEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEventsResponse.ProtoReflect.Descriptor instead.
func (*StreamEventsResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDescGZIP(), []int{1}
}

type StreamEventsRequest_Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node sending the event messages over the stream.
	Node *v3.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *StreamEventsRequest_Identifier) Reset() {
	*x = StreamEventsRequest_Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEventsRequest_Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEventsRequest_Identifier) ProtoMessage() {}

func (x *StreamEventsRequest_Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEventsRequest_Identifier.ProtoReflect.Descriptor instead.
func (*StreamEventsRequest_Identifier) Descriptor() ([]byte, []int) {
	return file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StreamEventsRequest_Identifier) GetNode() *v3.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

var File_envoy_service_event_reporting_v3_event_reporting_service_proto protoreflect.FileDescriptor

var file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x33, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x03, 0x0a, 0x13, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x60, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x93, 0x01, 0x0a, 0x0a,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x3a, 0x4b, 0x9a, 0xc5, 0x88, 0x1e, 0x46, 0x0a, 0x44, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x3a, 0x40, 0x9a, 0xc5, 0x88, 0x1e, 0x3b, 0x0a, 0x39, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x41, 0x9a, 0xc5, 0x88,
	0x1e, 0x3c, 0x0a, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9d,
	0x01, 0x0a, 0x15, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0xb1,
	0x01, 0x0a, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x42, 0x1a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDescOnce sync.Once
	file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDescData = file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDesc
)

func file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDescGZIP() []byte {
	file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDescOnce.Do(func() {
		file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDescData)
	})
	return file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDescData
}

var file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_service_event_reporting_v3_event_reporting_service_proto_goTypes = []interface{}{
	(*StreamEventsRequest)(nil),            // 0: envoy.service.event_reporting.v3.StreamEventsRequest
	(*StreamEventsResponse)(nil),           // 1: envoy.service.event_reporting.v3.StreamEventsResponse
	(*StreamEventsRequest_Identifier)(nil), // 2: envoy.service.event_reporting.v3.StreamEventsRequest.Identifier
	(*anypb.Any)(nil),                      // 3: google.protobuf.Any
	(*v3.Node)(nil),                        // 4: envoy.config.core.v3.Node
}
var file_envoy_service_event_reporting_v3_event_reporting_service_proto_depIdxs = []int32{
	2, // 0: envoy.service.event_reporting.v3.StreamEventsRequest.identifier:type_name -> envoy.service.event_reporting.v3.StreamEventsRequest.Identifier
	3, // 1: envoy.service.event_reporting.v3.StreamEventsRequest.events:type_name -> google.protobuf.Any
	4, // 2: envoy.service.event_reporting.v3.StreamEventsRequest.Identifier.node:type_name -> envoy.config.core.v3.Node
	0, // 3: envoy.service.event_reporting.v3.EventReportingService.StreamEvents:input_type -> envoy.service.event_reporting.v3.StreamEventsRequest
	1, // 4: envoy.service.event_reporting.v3.EventReportingService.StreamEvents:output_type -> envoy.service.event_reporting.v3.StreamEventsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_service_event_reporting_v3_event_reporting_service_proto_init() }
func file_envoy_service_event_reporting_v3_event_reporting_service_proto_init() {
	if File_envoy_service_event_reporting_v3_event_reporting_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEventsRequest); i {
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
		file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEventsResponse); i {
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
		file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEventsRequest_Identifier); i {
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
			RawDescriptor: file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_event_reporting_v3_event_reporting_service_proto_goTypes,
		DependencyIndexes: file_envoy_service_event_reporting_v3_event_reporting_service_proto_depIdxs,
		MessageInfos:      file_envoy_service_event_reporting_v3_event_reporting_service_proto_msgTypes,
	}.Build()
	File_envoy_service_event_reporting_v3_event_reporting_service_proto = out.File
	file_envoy_service_event_reporting_v3_event_reporting_service_proto_rawDesc = nil
	file_envoy_service_event_reporting_v3_event_reporting_service_proto_goTypes = nil
	file_envoy_service_event_reporting_v3_event_reporting_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventReportingServiceClient is the client API for EventReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventReportingServiceClient interface {
	// Envoy will connect and send StreamEventsRequest messages forever.
	// The management server may send StreamEventsResponse to configure event stream. See below.
	// This API is designed for high throughput with the expectation that it might be lossy.
	StreamEvents(ctx context.Context, opts ...grpc.CallOption) (EventReportingService_StreamEventsClient, error)
}

type eventReportingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventReportingServiceClient(cc grpc.ClientConnInterface) EventReportingServiceClient {
	return &eventReportingServiceClient{cc}
}

func (c *eventReportingServiceClient) StreamEvents(ctx context.Context, opts ...grpc.CallOption) (EventReportingService_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventReportingService_serviceDesc.Streams[0], "/envoy.service.event_reporting.v3.EventReportingService/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventReportingServiceStreamEventsClient{stream}
	return x, nil
}

type EventReportingService_StreamEventsClient interface {
	Send(*StreamEventsRequest) error
	Recv() (*StreamEventsResponse, error)
	grpc.ClientStream
}

type eventReportingServiceStreamEventsClient struct {
	grpc.ClientStream
}

func (x *eventReportingServiceStreamEventsClient) Send(m *StreamEventsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventReportingServiceStreamEventsClient) Recv() (*StreamEventsResponse, error) {
	m := new(StreamEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventReportingServiceServer is the server API for EventReportingService service.
type EventReportingServiceServer interface {
	// Envoy will connect and send StreamEventsRequest messages forever.
	// The management server may send StreamEventsResponse to configure event stream. See below.
	// This API is designed for high throughput with the expectation that it might be lossy.
	StreamEvents(EventReportingService_StreamEventsServer) error
}

// UnimplementedEventReportingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventReportingServiceServer struct {
}

func (*UnimplementedEventReportingServiceServer) StreamEvents(EventReportingService_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}

func RegisterEventReportingServiceServer(s *grpc.Server, srv EventReportingServiceServer) {
	s.RegisterService(&_EventReportingService_serviceDesc, srv)
}

func _EventReportingService_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventReportingServiceServer).StreamEvents(&eventReportingServiceStreamEventsServer{stream})
}

type EventReportingService_StreamEventsServer interface {
	Send(*StreamEventsResponse) error
	Recv() (*StreamEventsRequest, error)
	grpc.ServerStream
}

type eventReportingServiceStreamEventsServer struct {
	grpc.ServerStream
}

func (x *eventReportingServiceStreamEventsServer) Send(m *StreamEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventReportingServiceStreamEventsServer) Recv() (*StreamEventsRequest, error) {
	m := new(StreamEventsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EventReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.event_reporting.v3.EventReportingService",
	HandlerType: (*EventReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _EventReportingService_StreamEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/event_reporting/v3/event_reporting_service.proto",
}
