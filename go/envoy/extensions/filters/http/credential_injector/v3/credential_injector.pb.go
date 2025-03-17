// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/extensions/filters/http/credential_injector/v3/credential_injector.proto

package credential_injectorv3

import (
	v3 "github.com/khulnasoft/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
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

// Credential Injector injects credentials into outgoing HTTP requests. The filter configuration is used to retrieve the credentials, or
// they can be requested through the OAuth2 client credential grant. The credentials obtained are then injected into the Authorization header
// of the proxied HTTP requests, utilizing either the Basic or Bearer scheme.
//
// If the credential is not present or there was a failure injecting the credential, the request will fail with “401 Unauthorized“ unless
// “allow_request_without_credential“ is set to “true“.
//
// Notice: This filter is intended to be used for workload authentication, which means that the identity associated with the inserted credential
// is considered as the identity of the workload behind the envoy proxy(in this case, envoy is typically deployed as a sidecar alongside that
// workload). Please note that this filter does not handle end user authentication. Its purpose is solely to authenticate the workload itself.
//
// Here is an example of CredentialInjector configuration with Generic credential, which injects an HTTP Basic Auth credential into the proxied requests.
//
// .. code-block:: yaml
//
//	overwrite: true
//	credential:
//	  name: generic_credential
//	  typed_config:
//	    "@type": type.googleapis.com/envoy.extensions.http.injected_credentials.generic.v3.Generic
//	    credential:
//	      name: credential
//	      sds_config:
//	        path_config_source:
//	          path: credential.yaml
//	    header: Authorization
//
// credential.yaml for Basic Auth:
//
// .. code-block:: yaml
//
//	resources:
//	- "@type": "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.Secret"
//	  name: credential
//	  generic_secret:
//	    secret:
//	      inline_string: "Basic base64EncodedUsernamePassword"
//
// It can also be configured to inject a Bearer token into the proxied requests.
//
// credential.yaml for Bearer Token:
//
// .. code-block:: yaml
//
//	resources:
//	- "@type": "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.Secret"
//	  name: credential
//	  generic_secret:
//	    secret:
//	      inline_string: "Bearer myToken"
type CredentialInjector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether to overwrite the value or not if the injected headers already exist.
	// Value defaults to false.
	Overwrite bool `protobuf:"varint,1,opt,name=overwrite,proto3" json:"overwrite,omitempty"`
	// Whether to send the request to upstream if the credential is not present or if the credential injection
	// to the request fails.
	//
	// By default, a request will fail with “401 Unauthorized“ if the
	// credential is not present or the injection of the credential to the request fails.
	// If set to true, the request will be sent to upstream without the credential.
	AllowRequestWithoutCredential bool `protobuf:"varint,2,opt,name=allow_request_without_credential,json=allowRequestWithoutCredential,proto3" json:"allow_request_without_credential,omitempty"`
	// The credential to inject into the proxied requests
	// [#extension-category: envoy.http.injected_credentials]
	Credential *v3.TypedExtensionConfig `protobuf:"bytes,3,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *CredentialInjector) Reset() {
	*x = CredentialInjector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialInjector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialInjector) ProtoMessage() {}

func (x *CredentialInjector) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialInjector.ProtoReflect.Descriptor instead.
func (*CredentialInjector) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDescGZIP(), []int{0}
}

func (x *CredentialInjector) GetOverwrite() bool {
	if x != nil {
		return x.Overwrite
	}
	return false
}

func (x *CredentialInjector) GetAllowRequestWithoutCredential() bool {
	if x != nil {
		return x.AllowRequestWithoutCredential
	}
	return false
}

func (x *CredentialInjector) GetCredential() *v3.TypedExtensionConfig {
	if x != nil {
		return x.Credential
	}
	return nil
}

var File_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x47, 0x0a, 0x20, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x77, 0x69, 0x74, 0x68,
	0x6f, 0x75, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x54, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x42, 0xe2, 0x01, 0x0a, 0x42, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x33,
	0x42, 0x17, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x6a, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x71, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x76, 0x33, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDescData = file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDesc
)

func file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDescData
}

var file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_goTypes = []interface{}{
	(*CredentialInjector)(nil),      // 0: envoy.extensions.filters.http.credential_injector.v3.CredentialInjector
	(*v3.TypedExtensionConfig)(nil), // 1: envoy.config.core.v3.TypedExtensionConfig
}
var file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.credential_injector.v3.CredentialInjector.credential:type_name -> envoy.config.core.v3.TypedExtensionConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_init()
}
func file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_init() {
	if File_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialInjector); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto = out.File
	file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_rawDesc = nil
	file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_goTypes = nil
	file_envoy_extensions_filters_http_credential_injector_v3_credential_injector_proto_depIdxs = nil
}
