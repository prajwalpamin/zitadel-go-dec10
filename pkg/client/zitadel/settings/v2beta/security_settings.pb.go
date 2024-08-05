// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: zitadel/settings/v2beta/security_settings.proto

package settings

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type SecuritySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmbeddedIframe      *EmbeddedIframeSettings `protobuf:"bytes,1,opt,name=embedded_iframe,json=embeddedIframe,proto3" json:"embedded_iframe,omitempty"`
	EnableImpersonation bool                    `protobuf:"varint,2,opt,name=enable_impersonation,json=enableImpersonation,proto3" json:"enable_impersonation,omitempty"`
}

func (x *SecuritySettings) Reset() {
	*x = SecuritySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_settings_v2beta_security_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecuritySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecuritySettings) ProtoMessage() {}

func (x *SecuritySettings) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_settings_v2beta_security_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecuritySettings.ProtoReflect.Descriptor instead.
func (*SecuritySettings) Descriptor() ([]byte, []int) {
	return file_zitadel_settings_v2beta_security_settings_proto_rawDescGZIP(), []int{0}
}

func (x *SecuritySettings) GetEmbeddedIframe() *EmbeddedIframeSettings {
	if x != nil {
		return x.EmbeddedIframe
	}
	return nil
}

func (x *SecuritySettings) GetEnableImpersonation() bool {
	if x != nil {
		return x.EnableImpersonation
	}
	return false
}

type EmbeddedIframeSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled        bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	AllowedOrigins []string `protobuf:"bytes,2,rep,name=allowed_origins,json=allowedOrigins,proto3" json:"allowed_origins,omitempty"`
}

func (x *EmbeddedIframeSettings) Reset() {
	*x = EmbeddedIframeSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_settings_v2beta_security_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbeddedIframeSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbeddedIframeSettings) ProtoMessage() {}

func (x *EmbeddedIframeSettings) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_settings_v2beta_security_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbeddedIframeSettings.ProtoReflect.Descriptor instead.
func (*EmbeddedIframeSettings) Descriptor() ([]byte, []int) {
	return file_zitadel_settings_v2beta_security_settings_proto_rawDescGZIP(), []int{1}
}

func (x *EmbeddedIframeSettings) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *EmbeddedIframeSettings) GetAllowedOrigins() []string {
	if x != nil {
		return x.AllowedOrigins
	}
	return nil
}

var File_zitadel_settings_v2beta_security_settings_proto protoreflect.FileDescriptor

var file_zitadel_settings_v2beta_security_settings_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x10, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x58, 0x0a, 0x0f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x69, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x49, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x49, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x66, 0x0a, 0x14, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x33, 0x92, 0x41, 0x30, 0x32, 0x28, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4a, 0x04, 0x22, 0x65, 0x6e, 0x22, 0x52, 0x13, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x49, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xf6, 0x01, 0x0a, 0x16, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x49, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x50, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x36, 0x92,
	0x41, 0x33, 0x32, 0x31, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x20, 0x69, 0x66, 0x20, 0x69, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x20, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x69,
	0x73, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x89,
	0x01, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x60, 0x92, 0x41, 0x5d, 0x32, 0x38, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x20, 0x6c,
	0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x5a, 0x49, 0x54, 0x41, 0x44, 0x45, 0x4c, 0x20, 0x69,
	0x6e, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x20, 0x69, 0x66, 0x20, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x2e, 0x4a, 0x21, 0x5b, 0x22, 0x66, 0x6f, 0x6f, 0x2e, 0x62,
	0x61, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x2c, 0x20, 0x22, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68,
	0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x22, 0x5d, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c,
	0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x3b, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_zitadel_settings_v2beta_security_settings_proto_rawDescOnce sync.Once
	file_zitadel_settings_v2beta_security_settings_proto_rawDescData = file_zitadel_settings_v2beta_security_settings_proto_rawDesc
)

func file_zitadel_settings_v2beta_security_settings_proto_rawDescGZIP() []byte {
	file_zitadel_settings_v2beta_security_settings_proto_rawDescOnce.Do(func() {
		file_zitadel_settings_v2beta_security_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_settings_v2beta_security_settings_proto_rawDescData)
	})
	return file_zitadel_settings_v2beta_security_settings_proto_rawDescData
}

var file_zitadel_settings_v2beta_security_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_zitadel_settings_v2beta_security_settings_proto_goTypes = []interface{}{
	(*SecuritySettings)(nil),       // 0: zitadel.settings.v2beta.SecuritySettings
	(*EmbeddedIframeSettings)(nil), // 1: zitadel.settings.v2beta.EmbeddedIframeSettings
}
var file_zitadel_settings_v2beta_security_settings_proto_depIdxs = []int32{
	1, // 0: zitadel.settings.v2beta.SecuritySettings.embedded_iframe:type_name -> zitadel.settings.v2beta.EmbeddedIframeSettings
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_zitadel_settings_v2beta_security_settings_proto_init() }
func file_zitadel_settings_v2beta_security_settings_proto_init() {
	if File_zitadel_settings_v2beta_security_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zitadel_settings_v2beta_security_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecuritySettings); i {
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
		file_zitadel_settings_v2beta_security_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbeddedIframeSettings); i {
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
			RawDescriptor: file_zitadel_settings_v2beta_security_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_settings_v2beta_security_settings_proto_goTypes,
		DependencyIndexes: file_zitadel_settings_v2beta_security_settings_proto_depIdxs,
		MessageInfos:      file_zitadel_settings_v2beta_security_settings_proto_msgTypes,
	}.Build()
	File_zitadel_settings_v2beta_security_settings_proto = out.File
	file_zitadel_settings_v2beta_security_settings_proto_rawDesc = nil
	file_zitadel_settings_v2beta_security_settings_proto_goTypes = nil
	file_zitadel_settings_v2beta_security_settings_proto_depIdxs = nil
}
