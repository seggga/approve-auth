// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/auth-v2.proto

package proto

import (
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

type CheckTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *CheckTokenRequest) Reset() {
	*x = CheckTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTokenRequest) ProtoMessage() {}

func (x *CheckTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTokenRequest.ProtoReflect.Descriptor instead.
func (*CheckTokenRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_v2_proto_rawDescGZIP(), []int{0}
}

func (x *CheckTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CheckTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Error uint32 `protobuf:"varint,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *CheckTokenResponse) Reset() {
	*x = CheckTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTokenResponse) ProtoMessage() {}

func (x *CheckTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTokenResponse.ProtoReflect.Descriptor instead.
func (*CheckTokenResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_v2_proto_rawDescGZIP(), []int{1}
}

func (x *CheckTokenResponse) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CheckTokenResponse) GetError() uint32 {
	if x != nil {
		return x.Error
	}
	return 0
}

type RefreshTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *RefreshTokensRequest) Reset() {
	*x = RefreshTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokensRequest) ProtoMessage() {}

func (x *RefreshTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokensRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokensRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_v2_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshTokensRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RefreshTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	Error        uint32 `protobuf:"varint,3,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *RefreshTokensResponse) Reset() {
	*x = RefreshTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_v2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokensResponse) ProtoMessage() {}

func (x *RefreshTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokensResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokensResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_v2_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshTokensResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RefreshTokensResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RefreshTokensResponse) GetError() uint32 {
	if x != nil {
		return x.Error
	}
	return 0
}

var File_proto_auth_v2_proto protoreflect.FileDescriptor

var file_proto_auth_v2_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x76, 0x32, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x32, 0x22, 0x29, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x40, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x2c, 0x0a, 0x14, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x73, 0x0a, 0x15, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xc0, 0x01, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x41, 0x50,
	0x49, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_v2_proto_rawDescOnce sync.Once
	file_proto_auth_v2_proto_rawDescData = file_proto_auth_v2_proto_rawDesc
)

func file_proto_auth_v2_proto_rawDescGZIP() []byte {
	file_proto_auth_v2_proto_rawDescOnce.Do(func() {
		file_proto_auth_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_v2_proto_rawDescData)
	})
	return file_proto_auth_v2_proto_rawDescData
}

var file_proto_auth_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_auth_v2_proto_goTypes = []interface{}{
	(*CheckTokenRequest)(nil),     // 0: authservice.v2.CheckTokenRequest
	(*CheckTokenResponse)(nil),    // 1: authservice.v2.CheckTokenResponse
	(*RefreshTokensRequest)(nil),  // 2: authservice.v2.RefreshTokensRequest
	(*RefreshTokensResponse)(nil), // 3: authservice.v2.RefreshTokensResponse
}
var file_proto_auth_v2_proto_depIdxs = []int32{
	0, // 0: authservice.v2.AuthAPI.CheckToken:input_type -> authservice.v2.CheckTokenRequest
	2, // 1: authservice.v2.AuthAPI.RefreshTokens:input_type -> authservice.v2.RefreshTokensRequest
	1, // 2: authservice.v2.AuthAPI.CheckToken:output_type -> authservice.v2.CheckTokenResponse
	3, // 3: authservice.v2.AuthAPI.RefreshTokens:output_type -> authservice.v2.RefreshTokensResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_v2_proto_init() }
func file_proto_auth_v2_proto_init() {
	if File_proto_auth_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auth_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTokenRequest); i {
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
		file_proto_auth_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTokenResponse); i {
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
		file_proto_auth_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokensRequest); i {
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
		file_proto_auth_v2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokensResponse); i {
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
			RawDescriptor: file_proto_auth_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_v2_proto_goTypes,
		DependencyIndexes: file_proto_auth_v2_proto_depIdxs,
		MessageInfos:      file_proto_auth_v2_proto_msgTypes,
	}.Build()
	File_proto_auth_v2_proto = out.File
	file_proto_auth_v2_proto_rawDesc = nil
	file_proto_auth_v2_proto_goTypes = nil
	file_proto_auth_v2_proto_depIdxs = nil
}
