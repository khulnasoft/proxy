// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/extensions/filters/common/fault/v3/fault.proto

package faultv3

import (
	v3 "github.com/khulnasoft/proxy/go/envoy/type/v3"
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

type FaultDelay_FaultDelayType int32

const (
	// Unused and deprecated.
	FaultDelay_FIXED FaultDelay_FaultDelayType = 0
)

// Enum value maps for FaultDelay_FaultDelayType.
var (
	FaultDelay_FaultDelayType_name = map[int32]string{
		0: "FIXED",
	}
	FaultDelay_FaultDelayType_value = map[string]int32{
		"FIXED": 0,
	}
)

func (x FaultDelay_FaultDelayType) Enum() *FaultDelay_FaultDelayType {
	p := new(FaultDelay_FaultDelayType)
	*p = x
	return p
}

func (x FaultDelay_FaultDelayType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FaultDelay_FaultDelayType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_common_fault_v3_fault_proto_enumTypes[0].Descriptor()
}

func (FaultDelay_FaultDelayType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_common_fault_v3_fault_proto_enumTypes[0]
}

func (x FaultDelay_FaultDelayType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FaultDelay_FaultDelayType.Descriptor instead.
func (FaultDelay_FaultDelayType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescGZIP(), []int{0, 0}
}

// Delay specification is used to inject latency into the
// HTTP/Mongo operation.
// [#next-free-field: 6]
type FaultDelay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FaultDelaySecifier:
	//
	//	*FaultDelay_FixedDelay
	//	*FaultDelay_HeaderDelay_
	FaultDelaySecifier isFaultDelay_FaultDelaySecifier `protobuf_oneof:"fault_delay_secifier"`
	// The percentage of operations/connections/requests on which the delay will be injected.
	Percentage *v3.FractionalPercent `protobuf:"bytes,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *FaultDelay) Reset() {
	*x = FaultDelay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultDelay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultDelay) ProtoMessage() {}

func (x *FaultDelay) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultDelay.ProtoReflect.Descriptor instead.
func (*FaultDelay) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescGZIP(), []int{0}
}

func (m *FaultDelay) GetFaultDelaySecifier() isFaultDelay_FaultDelaySecifier {
	if m != nil {
		return m.FaultDelaySecifier
	}
	return nil
}

func (x *FaultDelay) GetFixedDelay() *durationpb.Duration {
	if x, ok := x.GetFaultDelaySecifier().(*FaultDelay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

func (x *FaultDelay) GetHeaderDelay() *FaultDelay_HeaderDelay {
	if x, ok := x.GetFaultDelaySecifier().(*FaultDelay_HeaderDelay_); ok {
		return x.HeaderDelay
	}
	return nil
}

func (x *FaultDelay) GetPercentage() *v3.FractionalPercent {
	if x != nil {
		return x.Percentage
	}
	return nil
}

type isFaultDelay_FaultDelaySecifier interface {
	isFaultDelay_FaultDelaySecifier()
}

type FaultDelay_FixedDelay struct {
	// Add a fixed delay before forwarding the operation upstream. See
	// https://developers.google.com/protocol-buffers/docs/proto3#json for
	// the JSON/YAML Duration mapping. For HTTP/Mongo, the specified
	// delay will be injected before a new request/operation.
	// This is required if type is FIXED.
	FixedDelay *durationpb.Duration `protobuf:"bytes,3,opt,name=fixed_delay,json=fixedDelay,proto3,oneof"`
}

type FaultDelay_HeaderDelay_ struct {
	// Fault delays are controlled via an HTTP header (if applicable).
	HeaderDelay *FaultDelay_HeaderDelay `protobuf:"bytes,5,opt,name=header_delay,json=headerDelay,proto3,oneof"`
}

func (*FaultDelay_FixedDelay) isFaultDelay_FaultDelaySecifier() {}

func (*FaultDelay_HeaderDelay_) isFaultDelay_FaultDelaySecifier() {}

// Describes a rate limit to be applied.
type FaultRateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to LimitType:
	//
	//	*FaultRateLimit_FixedLimit_
	//	*FaultRateLimit_HeaderLimit_
	LimitType isFaultRateLimit_LimitType `protobuf_oneof:"limit_type"`
	// The percentage of operations/connections/requests on which the rate limit will be injected.
	Percentage *v3.FractionalPercent `protobuf:"bytes,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *FaultRateLimit) Reset() {
	*x = FaultRateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultRateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultRateLimit) ProtoMessage() {}

func (x *FaultRateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultRateLimit.ProtoReflect.Descriptor instead.
func (*FaultRateLimit) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescGZIP(), []int{1}
}

func (m *FaultRateLimit) GetLimitType() isFaultRateLimit_LimitType {
	if m != nil {
		return m.LimitType
	}
	return nil
}

func (x *FaultRateLimit) GetFixedLimit() *FaultRateLimit_FixedLimit {
	if x, ok := x.GetLimitType().(*FaultRateLimit_FixedLimit_); ok {
		return x.FixedLimit
	}
	return nil
}

func (x *FaultRateLimit) GetHeaderLimit() *FaultRateLimit_HeaderLimit {
	if x, ok := x.GetLimitType().(*FaultRateLimit_HeaderLimit_); ok {
		return x.HeaderLimit
	}
	return nil
}

func (x *FaultRateLimit) GetPercentage() *v3.FractionalPercent {
	if x != nil {
		return x.Percentage
	}
	return nil
}

type isFaultRateLimit_LimitType interface {
	isFaultRateLimit_LimitType()
}

type FaultRateLimit_FixedLimit_ struct {
	// A fixed rate limit.
	FixedLimit *FaultRateLimit_FixedLimit `protobuf:"bytes,1,opt,name=fixed_limit,json=fixedLimit,proto3,oneof"`
}

type FaultRateLimit_HeaderLimit_ struct {
	// Rate limits are controlled via an HTTP header (if applicable).
	HeaderLimit *FaultRateLimit_HeaderLimit `protobuf:"bytes,3,opt,name=header_limit,json=headerLimit,proto3,oneof"`
}

func (*FaultRateLimit_FixedLimit_) isFaultRateLimit_LimitType() {}

func (*FaultRateLimit_HeaderLimit_) isFaultRateLimit_LimitType() {}

// Fault delays are controlled via an HTTP header (if applicable). See the
// :ref:`HTTP fault filter <config_http_filters_fault_injection_http_header>` documentation for
// more information.
type FaultDelay_HeaderDelay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FaultDelay_HeaderDelay) Reset() {
	*x = FaultDelay_HeaderDelay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultDelay_HeaderDelay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultDelay_HeaderDelay) ProtoMessage() {}

func (x *FaultDelay_HeaderDelay) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultDelay_HeaderDelay.ProtoReflect.Descriptor instead.
func (*FaultDelay_HeaderDelay) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescGZIP(), []int{0, 0}
}

// Describes a fixed/constant rate limit.
type FaultRateLimit_FixedLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The limit supplied in KiB/s.
	LimitKbps uint64 `protobuf:"varint,1,opt,name=limit_kbps,json=limitKbps,proto3" json:"limit_kbps,omitempty"`
}

func (x *FaultRateLimit_FixedLimit) Reset() {
	*x = FaultRateLimit_FixedLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultRateLimit_FixedLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultRateLimit_FixedLimit) ProtoMessage() {}

func (x *FaultRateLimit_FixedLimit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultRateLimit_FixedLimit.ProtoReflect.Descriptor instead.
func (*FaultRateLimit_FixedLimit) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FaultRateLimit_FixedLimit) GetLimitKbps() uint64 {
	if x != nil {
		return x.LimitKbps
	}
	return 0
}

// Rate limits are controlled via an HTTP header (if applicable). See the
// :ref:`HTTP fault filter <config_http_filters_fault_injection_http_header>` documentation for
// more information.
type FaultRateLimit_HeaderLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FaultRateLimit_HeaderLimit) Reset() {
	*x = FaultRateLimit_HeaderLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultRateLimit_HeaderLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultRateLimit_HeaderLimit) ProtoMessage() {}

func (x *FaultRateLimit_HeaderLimit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultRateLimit_HeaderLimit.ProtoReflect.Descriptor instead.
func (*FaultRateLimit_HeaderLimit) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescGZIP(), []int{1, 1}
}

var File_envoy_extensions_filters_common_fault_v3_fault_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDesc = []byte{
	0x0a, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33,
	0x1a, 0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x03, 0x0a, 0x0a, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x46, 0x0a, 0x0b, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a,
	0x00, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x69, 0x78, 0x65, 0x64, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12,
	0x65, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33,
	0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x1a, 0x49, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x3a, 0x3a, 0x9a, 0xc5, 0x88, 0x1e, 0x35, 0x0a, 0x33,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x61, 0x79, 0x22, 0x1b, 0x0a, 0x0e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x49, 0x58, 0x45, 0x44, 0x10, 0x00,
	0x3a, 0x2e, 0x9a, 0xc5, 0x88, 0x1e, 0x29, 0x0a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79,
	0x42, 0x1b, 0x0a, 0x14, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f,
	0x73, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x4a, 0x04, 0x08,
	0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0xb0, 0x04, 0x0a, 0x0e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x66, 0x0a, 0x0b, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x76, 0x33, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x00, 0x52, 0x0a,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x69, 0x0a, 0x0c, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x44, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x1a, 0x73, 0x0a, 0x0a, 0x46, 0x69, 0x78, 0x65, 0x64,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6b,
	0x62, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x28, 0x01, 0x52, 0x09, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4b, 0x62, 0x70, 0x73, 0x3a, 0x3d, 0x9a,
	0xc5, 0x88, 0x1e, 0x38, 0x0a, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x4d, 0x0a, 0x0b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x3e, 0x9a, 0xc5, 0x88,
	0x1e, 0x39, 0x0a, 0x37, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x32, 0x9a, 0xc5, 0x88,
	0x1e, 0x2d, 0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42,
	0x11, 0x0a, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8,
	0x42, 0x01, 0x42, 0xa7, 0x01, 0x0a, 0x36, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x46,
	0x61, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x33, 0x3b, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescData = file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDesc
)

func file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescData)
	})
	return file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDescData
}

var file_envoy_extensions_filters_common_fault_v3_fault_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_extensions_filters_common_fault_v3_fault_proto_goTypes = []interface{}{
	(FaultDelay_FaultDelayType)(0),     // 0: envoy.extensions.filters.common.fault.v3.FaultDelay.FaultDelayType
	(*FaultDelay)(nil),                 // 1: envoy.extensions.filters.common.fault.v3.FaultDelay
	(*FaultRateLimit)(nil),             // 2: envoy.extensions.filters.common.fault.v3.FaultRateLimit
	(*FaultDelay_HeaderDelay)(nil),     // 3: envoy.extensions.filters.common.fault.v3.FaultDelay.HeaderDelay
	(*FaultRateLimit_FixedLimit)(nil),  // 4: envoy.extensions.filters.common.fault.v3.FaultRateLimit.FixedLimit
	(*FaultRateLimit_HeaderLimit)(nil), // 5: envoy.extensions.filters.common.fault.v3.FaultRateLimit.HeaderLimit
	(*durationpb.Duration)(nil),        // 6: google.protobuf.Duration
	(*v3.FractionalPercent)(nil),       // 7: envoy.type.v3.FractionalPercent
}
var file_envoy_extensions_filters_common_fault_v3_fault_proto_depIdxs = []int32{
	6, // 0: envoy.extensions.filters.common.fault.v3.FaultDelay.fixed_delay:type_name -> google.protobuf.Duration
	3, // 1: envoy.extensions.filters.common.fault.v3.FaultDelay.header_delay:type_name -> envoy.extensions.filters.common.fault.v3.FaultDelay.HeaderDelay
	7, // 2: envoy.extensions.filters.common.fault.v3.FaultDelay.percentage:type_name -> envoy.type.v3.FractionalPercent
	4, // 3: envoy.extensions.filters.common.fault.v3.FaultRateLimit.fixed_limit:type_name -> envoy.extensions.filters.common.fault.v3.FaultRateLimit.FixedLimit
	5, // 4: envoy.extensions.filters.common.fault.v3.FaultRateLimit.header_limit:type_name -> envoy.extensions.filters.common.fault.v3.FaultRateLimit.HeaderLimit
	7, // 5: envoy.extensions.filters.common.fault.v3.FaultRateLimit.percentage:type_name -> envoy.type.v3.FractionalPercent
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_common_fault_v3_fault_proto_init() }
func file_envoy_extensions_filters_common_fault_v3_fault_proto_init() {
	if File_envoy_extensions_filters_common_fault_v3_fault_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultDelay); i {
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
		file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultRateLimit); i {
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
		file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultDelay_HeaderDelay); i {
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
		file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultRateLimit_FixedLimit); i {
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
		file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultRateLimit_HeaderLimit); i {
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
	file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FaultDelay_FixedDelay)(nil),
		(*FaultDelay_HeaderDelay_)(nil),
	}
	file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*FaultRateLimit_FixedLimit_)(nil),
		(*FaultRateLimit_HeaderLimit_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_common_fault_v3_fault_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_common_fault_v3_fault_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_common_fault_v3_fault_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_common_fault_v3_fault_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_common_fault_v3_fault_proto = out.File
	file_envoy_extensions_filters_common_fault_v3_fault_proto_rawDesc = nil
	file_envoy_extensions_filters_common_fault_v3_fault_proto_goTypes = nil
	file_envoy_extensions_filters_common_fault_v3_fault_proto_depIdxs = nil
}
