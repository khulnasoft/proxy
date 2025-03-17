// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/config/route/v3/route.proto

package routev3

import (
	v3 "github.com/khulnasoft/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// [#next-free-field: 18]
type RouteConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the route configuration. For example, it might match
	// :ref:`route_config_name
	// <envoy_v3_api_field_extensions.filters.network.http_connection_manager.v3.Rds.route_config_name>` in
	// :ref:`envoy_v3_api_msg_extensions.filters.network.http_connection_manager.v3.Rds`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An array of virtual hosts that make up the route table.
	VirtualHosts []*VirtualHost `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	// An array of virtual hosts will be dynamically loaded via the VHDS API.
	// Both “virtual_hosts“ and “vhds“ fields will be used when present. “virtual_hosts“ can be used
	// for a base routing table or for infrequently changing virtual hosts. “vhds“ is used for
	// on-demand discovery of virtual hosts. The contents of these two fields will be merged to
	// generate a routing table for a given RouteConfiguration, with “vhds“ derived configuration
	// taking precedence.
	Vhds *Vhds `protobuf:"bytes,9,opt,name=vhds,proto3" json:"vhds,omitempty"`
	// Optionally specifies a list of HTTP headers that the connection manager
	// will consider to be internal only. If they are found on external requests they will be cleaned
	// prior to filter invocation. See :ref:`config_http_conn_man_headers_x-envoy-internal` for more
	// information.
	InternalOnlyHeaders []string `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response that
	// the connection manager encodes. Headers specified at this level are applied
	// after headers from any enclosed :ref:`envoy_v3_api_msg_config.route.v3.VirtualHost` or
	// :ref:`envoy_v3_api_msg_config.route.v3.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	ResponseHeadersToAdd []*v3.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each response
	// that the connection manager encodes.
	ResponseHeadersToRemove []string `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request
	// routed by the HTTP connection manager. Headers specified at this level are
	// applied after headers from any enclosed :ref:`envoy_v3_api_msg_config.route.v3.VirtualHost` or
	// :ref:`envoy_v3_api_msg_config.route.v3.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	RequestHeadersToAdd []*v3.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each request
	// routed by the HTTP connection manager.
	RequestHeadersToRemove []string `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	// Headers mutations at all levels are evaluated, if specified. By default, the order is from most
	// specific (i.e. route entry level) to least specific (i.e. route configuration level). Later header
	// mutations may override earlier mutations.
	// This order can be reversed by setting this field to true. In other words, most specific level mutation
	// is evaluated last.
	MostSpecificHeaderMutationsWins bool `protobuf:"varint,10,opt,name=most_specific_header_mutations_wins,json=mostSpecificHeaderMutationsWins,proto3" json:"most_specific_header_mutations_wins,omitempty"`
	// An optional boolean that specifies whether the clusters that the route
	// table refers to will be validated by the cluster manager. If set to true
	// and a route refers to a non-existent cluster, the route table will not
	// load. If set to false and a route refers to a non-existent cluster, the
	// route table will load and the router filter will return a 404 if the route
	// is selected at runtime. This setting defaults to true if the route table
	// is statically defined via the :ref:`route_config
	// <envoy_v3_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.route_config>`
	// option. This setting default to false if the route table is loaded dynamically via the
	// :ref:`rds
	// <envoy_v3_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.rds>`
	// option. Users may wish to override the default behavior in certain cases (for example when
	// using CDS with a static route table).
	ValidateClusters *wrapperspb.BoolValue `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
	// The maximum bytes of the response :ref:`direct response body
	// <envoy_v3_api_field_config.route.v3.DirectResponseAction.body>` size. If not specified the default
	// is 4096.
	//
	// .. warning::
	//
	//	Envoy currently holds the content of :ref:`direct response body
	//	<envoy_v3_api_field_config.route.v3.DirectResponseAction.body>` in memory. Be careful setting
	//	this to be larger than the default 4KB, since the allocated memory for direct response body
	//	is not subject to data plane buffering controls.
	MaxDirectResponseBodySizeBytes *wrapperspb.UInt32Value `protobuf:"bytes,11,opt,name=max_direct_response_body_size_bytes,json=maxDirectResponseBodySizeBytes,proto3" json:"max_direct_response_body_size_bytes,omitempty"`
	// A list of plugins and their configurations which may be used by a
	// :ref:`cluster specifier plugin name <envoy_v3_api_field_config.route.v3.RouteAction.cluster_specifier_plugin>`
	// within the route. All “extension.name“ fields in this list must be unique.
	ClusterSpecifierPlugins []*ClusterSpecifierPlugin `protobuf:"bytes,12,rep,name=cluster_specifier_plugins,json=clusterSpecifierPlugins,proto3" json:"cluster_specifier_plugins,omitempty"`
	// Specify a set of default request mirroring policies which apply to all routes under its virtual hosts.
	// Note that policies are not merged, the most specific non-empty one becomes the mirror policies.
	RequestMirrorPolicies []*RouteAction_RequestMirrorPolicy `protobuf:"bytes,13,rep,name=request_mirror_policies,json=requestMirrorPolicies,proto3" json:"request_mirror_policies,omitempty"`
	// By default, port in :authority header (if any) is used in host matching.
	// With this option enabled, Envoy will ignore the port number in the :authority header (if any) when picking VirtualHost.
	// NOTE: this option will not strip the port number (if any) contained in route config
	// :ref:`envoy_v3_api_msg_config.route.v3.VirtualHost`.domains field.
	IgnorePortInHostMatching bool `protobuf:"varint,14,opt,name=ignore_port_in_host_matching,json=ignorePortInHostMatching,proto3" json:"ignore_port_in_host_matching,omitempty"`
	// Ignore path-parameters in path-matching.
	// Before RFC3986, URI were like(RFC1808): <scheme>://<net_loc>/<path>;<params>?<query>#<fragment>
	// Envoy by default takes ":path" as "<path>;<params>".
	// For users who want to only match path on the "<path>" portion, this option should be true.
	IgnorePathParametersInPathMatching bool `protobuf:"varint,15,opt,name=ignore_path_parameters_in_path_matching,json=ignorePathParametersInPathMatching,proto3" json:"ignore_path_parameters_in_path_matching,omitempty"`
	// This field can be used to provide RouteConfiguration level per filter config. The key should match the
	// :ref:`filter config name
	// <envoy_v3_api_field_extensions.filters.network.http_connection_manager.v3.HttpFilter.name>`.
	// See :ref:`Http filter route specific config <arch_overview_http_filters_per_filter_config>`
	// for details.
	// [#comment: An entry's value may be wrapped in a
	// :ref:`FilterConfig<envoy_v3_api_msg_config.route.v3.FilterConfig>`
	// message to specify additional options.]
	TypedPerFilterConfig map[string]*anypb.Any `protobuf:"bytes,16,rep,name=typed_per_filter_config,json=typedPerFilterConfig,proto3" json:"typed_per_filter_config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The metadata field can be used to provide additional information
	// about the route configuration. It can be used for configuration, stats, and logging.
	// The metadata should go under the filter namespace that will need it.
	// For instance, if the metadata is intended for the Router filter,
	// the filter name should be specified as “envoy.filters.http.router“.
	Metadata *v3.Metadata `protobuf:"bytes,17,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *RouteConfiguration) Reset() {
	*x = RouteConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_route_v3_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteConfiguration) ProtoMessage() {}

func (x *RouteConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_route_v3_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteConfiguration.ProtoReflect.Descriptor instead.
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return file_envoy_config_route_v3_route_proto_rawDescGZIP(), []int{0}
}

func (x *RouteConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteConfiguration) GetVirtualHosts() []*VirtualHost {
	if x != nil {
		return x.VirtualHosts
	}
	return nil
}

func (x *RouteConfiguration) GetVhds() *Vhds {
	if x != nil {
		return x.Vhds
	}
	return nil
}

func (x *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if x != nil {
		return x.InternalOnlyHeaders
	}
	return nil
}

func (x *RouteConfiguration) GetResponseHeadersToAdd() []*v3.HeaderValueOption {
	if x != nil {
		return x.ResponseHeadersToAdd
	}
	return nil
}

func (x *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if x != nil {
		return x.ResponseHeadersToRemove
	}
	return nil
}

func (x *RouteConfiguration) GetRequestHeadersToAdd() []*v3.HeaderValueOption {
	if x != nil {
		return x.RequestHeadersToAdd
	}
	return nil
}

func (x *RouteConfiguration) GetRequestHeadersToRemove() []string {
	if x != nil {
		return x.RequestHeadersToRemove
	}
	return nil
}

func (x *RouteConfiguration) GetMostSpecificHeaderMutationsWins() bool {
	if x != nil {
		return x.MostSpecificHeaderMutationsWins
	}
	return false
}

func (x *RouteConfiguration) GetValidateClusters() *wrapperspb.BoolValue {
	if x != nil {
		return x.ValidateClusters
	}
	return nil
}

func (x *RouteConfiguration) GetMaxDirectResponseBodySizeBytes() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxDirectResponseBodySizeBytes
	}
	return nil
}

func (x *RouteConfiguration) GetClusterSpecifierPlugins() []*ClusterSpecifierPlugin {
	if x != nil {
		return x.ClusterSpecifierPlugins
	}
	return nil
}

func (x *RouteConfiguration) GetRequestMirrorPolicies() []*RouteAction_RequestMirrorPolicy {
	if x != nil {
		return x.RequestMirrorPolicies
	}
	return nil
}

func (x *RouteConfiguration) GetIgnorePortInHostMatching() bool {
	if x != nil {
		return x.IgnorePortInHostMatching
	}
	return false
}

func (x *RouteConfiguration) GetIgnorePathParametersInPathMatching() bool {
	if x != nil {
		return x.IgnorePathParametersInPathMatching
	}
	return false
}

func (x *RouteConfiguration) GetTypedPerFilterConfig() map[string]*anypb.Any {
	if x != nil {
		return x.TypedPerFilterConfig
	}
	return nil
}

func (x *RouteConfiguration) GetMetadata() *v3.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Vhds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration source specifier for VHDS.
	ConfigSource *v3.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
}

func (x *Vhds) Reset() {
	*x = Vhds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_route_v3_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vhds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vhds) ProtoMessage() {}

func (x *Vhds) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_route_v3_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vhds.ProtoReflect.Descriptor instead.
func (*Vhds) Descriptor() ([]byte, []int) {
	return file_envoy_config_route_v3_route_proto_rawDescGZIP(), []int{1}
}

func (x *Vhds) GetConfigSource() *v3.ConfigSource {
	if x != nil {
		return x.ConfigSource
	}
	return nil
}

var File_envoy_config_route_v3_route_proto protoreflect.FileDescriptor

var file_envoy_config_route_v3_route_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x0c, 0x0a, 0x12, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x0c, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x2f, 0x0a,
	0x04, 0x76, 0x68, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x56, 0x68, 0x64, 0x73, 0x52, 0x04, 0x76, 0x68, 0x64, 0x73, 0x12, 0x44,
	0x0a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x10, 0xfa,
	0x42, 0x0d, 0x92, 0x01, 0x0a, 0x22, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xc8, 0x01, 0x00, 0x52,
	0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x6e, 0x6c, 0x79, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x69, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x92, 0x01, 0x03, 0x10, 0xe8, 0x07, 0x52, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12,
	0x4d, 0x0a, 0x1a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x92, 0x01, 0x0a, 0x22, 0x08, 0x72, 0x06, 0xc0,
	0x01, 0x01, 0xc8, 0x01, 0x00, 0x52, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x67,
	0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x92, 0x01, 0x03, 0x10,
	0xe8, 0x07, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x4b, 0x0a, 0x19, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x92,
	0x01, 0x0a, 0x22, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xc8, 0x01, 0x00, 0x52, 0x16, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x12, 0x4c, 0x0a, 0x23, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x1f, 0x6d, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x69,
	0x6e, 0x73, 0x12, 0x47, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x69, 0x0a, 0x23, 0x6d,
	0x61, 0x78, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1e, 0x6d, 0x61, 0x78, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x69, 0x7a,
	0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x69, 0x0a, 0x19, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x17, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x12, 0x6e, 0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x69, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x69,
	0x72, 0x72, 0x6f, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x3e, 0x0a, 0x1c, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x69, 0x6e, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x49, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x12, 0x53, 0x0a, 0x27, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x22, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x7a, 0x0a, 0x17, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x50, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x74, 0x79,
	0x70, 0x65, 0x64, 0x50, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x3a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x5d,
	0x0a, 0x19, 0x54, 0x79, 0x70, 0x65, 0x64, 0x50, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x26, 0x9a,
	0xc5, 0x88, 0x1e, 0x21, 0x0a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x04, 0x56, 0x68, 0x64, 0x73, 0x12, 0x51, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x3a, 0x18, 0x9a, 0xc5, 0x88, 0x1e, 0x13, 0x0a, 0x11, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x68, 0x64, 0x73, 0x42, 0x81, 0x01, 0x0a, 0x23, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x76, 0x33, 0x42, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_route_v3_route_proto_rawDescOnce sync.Once
	file_envoy_config_route_v3_route_proto_rawDescData = file_envoy_config_route_v3_route_proto_rawDesc
)

func file_envoy_config_route_v3_route_proto_rawDescGZIP() []byte {
	file_envoy_config_route_v3_route_proto_rawDescOnce.Do(func() {
		file_envoy_config_route_v3_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_route_v3_route_proto_rawDescData)
	})
	return file_envoy_config_route_v3_route_proto_rawDescData
}

var file_envoy_config_route_v3_route_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_config_route_v3_route_proto_goTypes = []interface{}{
	(*RouteConfiguration)(nil),              // 0: envoy.config.route.v3.RouteConfiguration
	(*Vhds)(nil),                            // 1: envoy.config.route.v3.Vhds
	nil,                                     // 2: envoy.config.route.v3.RouteConfiguration.TypedPerFilterConfigEntry
	(*VirtualHost)(nil),                     // 3: envoy.config.route.v3.VirtualHost
	(*v3.HeaderValueOption)(nil),            // 4: envoy.config.core.v3.HeaderValueOption
	(*wrapperspb.BoolValue)(nil),            // 5: google.protobuf.BoolValue
	(*wrapperspb.UInt32Value)(nil),          // 6: google.protobuf.UInt32Value
	(*ClusterSpecifierPlugin)(nil),          // 7: envoy.config.route.v3.ClusterSpecifierPlugin
	(*RouteAction_RequestMirrorPolicy)(nil), // 8: envoy.config.route.v3.RouteAction.RequestMirrorPolicy
	(*v3.Metadata)(nil),                     // 9: envoy.config.core.v3.Metadata
	(*v3.ConfigSource)(nil),                 // 10: envoy.config.core.v3.ConfigSource
	(*anypb.Any)(nil),                       // 11: google.protobuf.Any
}
var file_envoy_config_route_v3_route_proto_depIdxs = []int32{
	3,  // 0: envoy.config.route.v3.RouteConfiguration.virtual_hosts:type_name -> envoy.config.route.v3.VirtualHost
	1,  // 1: envoy.config.route.v3.RouteConfiguration.vhds:type_name -> envoy.config.route.v3.Vhds
	4,  // 2: envoy.config.route.v3.RouteConfiguration.response_headers_to_add:type_name -> envoy.config.core.v3.HeaderValueOption
	4,  // 3: envoy.config.route.v3.RouteConfiguration.request_headers_to_add:type_name -> envoy.config.core.v3.HeaderValueOption
	5,  // 4: envoy.config.route.v3.RouteConfiguration.validate_clusters:type_name -> google.protobuf.BoolValue
	6,  // 5: envoy.config.route.v3.RouteConfiguration.max_direct_response_body_size_bytes:type_name -> google.protobuf.UInt32Value
	7,  // 6: envoy.config.route.v3.RouteConfiguration.cluster_specifier_plugins:type_name -> envoy.config.route.v3.ClusterSpecifierPlugin
	8,  // 7: envoy.config.route.v3.RouteConfiguration.request_mirror_policies:type_name -> envoy.config.route.v3.RouteAction.RequestMirrorPolicy
	2,  // 8: envoy.config.route.v3.RouteConfiguration.typed_per_filter_config:type_name -> envoy.config.route.v3.RouteConfiguration.TypedPerFilterConfigEntry
	9,  // 9: envoy.config.route.v3.RouteConfiguration.metadata:type_name -> envoy.config.core.v3.Metadata
	10, // 10: envoy.config.route.v3.Vhds.config_source:type_name -> envoy.config.core.v3.ConfigSource
	11, // 11: envoy.config.route.v3.RouteConfiguration.TypedPerFilterConfigEntry.value:type_name -> google.protobuf.Any
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_envoy_config_route_v3_route_proto_init() }
func file_envoy_config_route_v3_route_proto_init() {
	if File_envoy_config_route_v3_route_proto != nil {
		return
	}
	file_envoy_config_route_v3_route_components_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_route_v3_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteConfiguration); i {
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
		file_envoy_config_route_v3_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vhds); i {
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
			RawDescriptor: file_envoy_config_route_v3_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_route_v3_route_proto_goTypes,
		DependencyIndexes: file_envoy_config_route_v3_route_proto_depIdxs,
		MessageInfos:      file_envoy_config_route_v3_route_proto_msgTypes,
	}.Build()
	File_envoy_config_route_v3_route_proto = out.File
	file_envoy_config_route_v3_route_proto_rawDesc = nil
	file_envoy_config_route_v3_route_proto_goTypes = nil
	file_envoy_config_route_v3_route_proto_depIdxs = nil
}
