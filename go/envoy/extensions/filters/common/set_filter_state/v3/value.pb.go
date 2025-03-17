// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/extensions/filters/common/set_filter_state/v3/value.proto

package set_filter_statev3

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

type FilterStateValue_SharedWithUpstream int32

const (
	// Object is not shared with the upstream internal connections.
	FilterStateValue_NONE FilterStateValue_SharedWithUpstream = 0
	// Object is shared with the upstream internal connection.
	FilterStateValue_ONCE FilterStateValue_SharedWithUpstream = 1
	// Object is shared with the upstream internal connection and any internal connection upstream from it.
	FilterStateValue_TRANSITIVE FilterStateValue_SharedWithUpstream = 2
)

// Enum value maps for FilterStateValue_SharedWithUpstream.
var (
	FilterStateValue_SharedWithUpstream_name = map[int32]string{
		0: "NONE",
		1: "ONCE",
		2: "TRANSITIVE",
	}
	FilterStateValue_SharedWithUpstream_value = map[string]int32{
		"NONE":       0,
		"ONCE":       1,
		"TRANSITIVE": 2,
	}
)

func (x FilterStateValue_SharedWithUpstream) Enum() *FilterStateValue_SharedWithUpstream {
	p := new(FilterStateValue_SharedWithUpstream)
	*p = x
	return p
}

func (x FilterStateValue_SharedWithUpstream) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterStateValue_SharedWithUpstream) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_enumTypes[0].Descriptor()
}

func (FilterStateValue_SharedWithUpstream) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_enumTypes[0]
}

func (x FilterStateValue_SharedWithUpstream) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterStateValue_SharedWithUpstream.Descriptor instead.
func (FilterStateValue_SharedWithUpstream) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDescGZIP(), []int{0, 0}
}

// A filter state key and value pair.
// [#next-free-field: 7]
type FilterStateValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Key:
	//
	//	*FilterStateValue_ObjectKey
	Key isFilterStateValue_Key `protobuf_oneof:"key"`
	// Optional filter object factory lookup key. See :ref:`the well-known filter state keys <well_known_filter_state>`
	// for a list of valid factory keys.
	FactoryKey string `protobuf:"bytes,6,opt,name=factory_key,json=factoryKey,proto3" json:"factory_key,omitempty"`
	// Types that are assignable to Value:
	//
	//	*FilterStateValue_FormatString
	Value isFilterStateValue_Value `protobuf_oneof:"value"`
	// If marked as read-only, the filter state key value is locked, and cannot
	// be overridden by any filter, including this filter.
	ReadOnly bool `protobuf:"varint,3,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	// Configures the object to be shared with the upstream internal connections. See :ref:`internal upstream
	// transport <config_internal_upstream_transport>` for more details on the filter state sharing with
	// the internal connections.
	SharedWithUpstream FilterStateValue_SharedWithUpstream `protobuf:"varint,4,opt,name=shared_with_upstream,json=sharedWithUpstream,proto3,enum=envoy.extensions.filters.common.set_filter_state.v3.FilterStateValue_SharedWithUpstream" json:"shared_with_upstream,omitempty"`
	// Skip the update if the value evaluates to an empty string.
	// This option can be used to supply multiple alternatives for the same filter state object key.
	SkipIfEmpty bool `protobuf:"varint,5,opt,name=skip_if_empty,json=skipIfEmpty,proto3" json:"skip_if_empty,omitempty"`
}

func (x *FilterStateValue) Reset() {
	*x = FilterStateValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterStateValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterStateValue) ProtoMessage() {}

func (x *FilterStateValue) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterStateValue.ProtoReflect.Descriptor instead.
func (*FilterStateValue) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDescGZIP(), []int{0}
}

func (m *FilterStateValue) GetKey() isFilterStateValue_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (x *FilterStateValue) GetObjectKey() string {
	if x, ok := x.GetKey().(*FilterStateValue_ObjectKey); ok {
		return x.ObjectKey
	}
	return ""
}

func (x *FilterStateValue) GetFactoryKey() string {
	if x != nil {
		return x.FactoryKey
	}
	return ""
}

func (m *FilterStateValue) GetValue() isFilterStateValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *FilterStateValue) GetFormatString() *v3.SubstitutionFormatString {
	if x, ok := x.GetValue().(*FilterStateValue_FormatString); ok {
		return x.FormatString
	}
	return nil
}

func (x *FilterStateValue) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

func (x *FilterStateValue) GetSharedWithUpstream() FilterStateValue_SharedWithUpstream {
	if x != nil {
		return x.SharedWithUpstream
	}
	return FilterStateValue_NONE
}

func (x *FilterStateValue) GetSkipIfEmpty() bool {
	if x != nil {
		return x.SkipIfEmpty
	}
	return false
}

type isFilterStateValue_Key interface {
	isFilterStateValue_Key()
}

type FilterStateValue_ObjectKey struct {
	// Filter state object key. The key is used to lookup the object factory, unless :ref:`factory_key
	// <envoy_v3_api_field_extensions.filters.common.set_filter_state.v3.FilterStateValue.factory_key>` is set. See
	// :ref:`the well-known filter state keys <well_known_filter_state>` for a list of valid object keys.
	ObjectKey string `protobuf:"bytes,1,opt,name=object_key,json=objectKey,proto3,oneof"`
}

func (*FilterStateValue_ObjectKey) isFilterStateValue_Key() {}

type isFilterStateValue_Value interface {
	isFilterStateValue_Value()
}

type FilterStateValue_FormatString struct {
	// Uses the :ref:`format string <config_access_log_format_strings>` to
	// instantiate the filter state object value.
	FormatString *v3.SubstitutionFormatString `protobuf:"bytes,2,opt,name=format_string,json=formatString,proto3,oneof"`
}

func (*FilterStateValue_FormatString) isFilterStateValue_Value() {}

var File_envoy_extensions_filters_common_set_filter_state_v3_value_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x03, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x55, 0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x01, 0x52, 0x0c,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x8a, 0x01, 0x0a, 0x14, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x58, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x74, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x12, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x55, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x69,
	0x66, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73,
	0x6b, 0x69, 0x70, 0x49, 0x66, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x38, 0x0a, 0x12, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x4e,
	0x43, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x02, 0x42, 0x0a, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x42, 0x0c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0xc8,
	0x01, 0x0a, 0x41, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x6d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x74, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x73,
	0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x76,
	0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDescData = file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDesc
)

func file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDescData)
	})
	return file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDescData
}

var file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_goTypes = []interface{}{
	(FilterStateValue_SharedWithUpstream)(0), // 0: envoy.extensions.filters.common.set_filter_state.v3.FilterStateValue.SharedWithUpstream
	(*FilterStateValue)(nil),                 // 1: envoy.extensions.filters.common.set_filter_state.v3.FilterStateValue
	(*v3.SubstitutionFormatString)(nil),      // 2: envoy.config.core.v3.SubstitutionFormatString
}
var file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.common.set_filter_state.v3.FilterStateValue.format_string:type_name -> envoy.config.core.v3.SubstitutionFormatString
	0, // 1: envoy.extensions.filters.common.set_filter_state.v3.FilterStateValue.shared_with_upstream:type_name -> envoy.extensions.filters.common.set_filter_state.v3.FilterStateValue.SharedWithUpstream
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_init() }
func file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_init() {
	if File_envoy_extensions_filters_common_set_filter_state_v3_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterStateValue); i {
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
	file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FilterStateValue_ObjectKey)(nil),
		(*FilterStateValue_FormatString)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_common_set_filter_state_v3_value_proto = out.File
	file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_rawDesc = nil
	file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_goTypes = nil
	file_envoy_extensions_filters_common_set_filter_state_v3_value_proto_depIdxs = nil
}
