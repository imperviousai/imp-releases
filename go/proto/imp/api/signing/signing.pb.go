/// Allows an Imp node to sign and verify messaging with the connected LND

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/imp/api/signing/signing.proto

package gen

import (
	context "context"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//*
// Represents a request to sign a message
type SignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"` // message to be signed
}

func (x *SignRequest) Reset() {
	*x = SignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_signing_signing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRequest) ProtoMessage() {}

func (x *SignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_signing_signing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRequest.ProtoReflect.Descriptor instead.
func (*SignRequest) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_signing_signing_proto_rawDescGZIP(), []int{0}
}

func (x *SignRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

//*
// Represents a response from a signature request
type SignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"` // signature of signed message
}

func (x *SignResponse) Reset() {
	*x = SignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_signing_signing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignResponse) ProtoMessage() {}

func (x *SignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_signing_signing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignResponse.ProtoReflect.Descriptor instead.
func (*SignResponse) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_signing_signing_proto_rawDescGZIP(), []int{1}
}

func (x *SignResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

//*
// Represents a request to verify a signature and message
type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg       string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`             // message to be verified
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"` // signature of message
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_signing_signing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_signing_signing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_signing_signing_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *VerifyRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

//*
// Represents a response back from a verification request
type VerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"` // result of signature verification
}

func (x *VerifyResponse) Reset() {
	*x = VerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_signing_signing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResponse) ProtoMessage() {}

func (x *VerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_signing_signing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResponse.ProtoReflect.Descriptor instead.
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_signing_signing_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_proto_imp_api_signing_signing_proto protoreflect.FileDescriptor

var file_proto_imp_api_signing_signing_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0b,
	0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2c, 0x0a,
	0x0c, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x3f, 0x0a, 0x0d, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x28, 0x0a, 0x0e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xb5, 0x01, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x4f, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e,
	0x3a, 0x01, 0x2a, 0x12, 0x59, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22,
	0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0xc7,
	0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d,
	0x70, 0x65, 0x72, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x61, 0x69, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x69,
	0x6d, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x92, 0x41, 0x9e, 0x01, 0x12, 0x3f, 0x0a, 0x10, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x26,
	0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x65, 0x72, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x20, 0x41, 0x49, 0x12,
	0x15, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x2e, 0x61, 0x69, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x03, 0x01, 0x02, 0x04,
	0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x32, 0x0a, 0x14, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x6e, 0x20, 0x49, 0x4d, 0x50, 0x12, 0x1a, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x69, 0x6d, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x2e, 0x61, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_imp_api_signing_signing_proto_rawDescOnce sync.Once
	file_proto_imp_api_signing_signing_proto_rawDescData = file_proto_imp_api_signing_signing_proto_rawDesc
)

func file_proto_imp_api_signing_signing_proto_rawDescGZIP() []byte {
	file_proto_imp_api_signing_signing_proto_rawDescOnce.Do(func() {
		file_proto_imp_api_signing_signing_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_imp_api_signing_signing_proto_rawDescData)
	})
	return file_proto_imp_api_signing_signing_proto_rawDescData
}

var file_proto_imp_api_signing_signing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_imp_api_signing_signing_proto_goTypes = []interface{}{
	(*SignRequest)(nil),    // 0: signing.SignRequest
	(*SignResponse)(nil),   // 1: signing.SignResponse
	(*VerifyRequest)(nil),  // 2: signing.VerifyRequest
	(*VerifyResponse)(nil), // 3: signing.VerifyResponse
}
var file_proto_imp_api_signing_signing_proto_depIdxs = []int32{
	0, // 0: signing.Signing.SignMessage:input_type -> signing.SignRequest
	2, // 1: signing.Signing.VerifySignature:input_type -> signing.VerifyRequest
	1, // 2: signing.Signing.SignMessage:output_type -> signing.SignResponse
	3, // 3: signing.Signing.VerifySignature:output_type -> signing.VerifyResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_imp_api_signing_signing_proto_init() }
func file_proto_imp_api_signing_signing_proto_init() {
	if File_proto_imp_api_signing_signing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_imp_api_signing_signing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRequest); i {
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
		file_proto_imp_api_signing_signing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignResponse); i {
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
		file_proto_imp_api_signing_signing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_proto_imp_api_signing_signing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResponse); i {
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
			RawDescriptor: file_proto_imp_api_signing_signing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_imp_api_signing_signing_proto_goTypes,
		DependencyIndexes: file_proto_imp_api_signing_signing_proto_depIdxs,
		MessageInfos:      file_proto_imp_api_signing_signing_proto_msgTypes,
	}.Build()
	File_proto_imp_api_signing_signing_proto = out.File
	file_proto_imp_api_signing_signing_proto_rawDesc = nil
	file_proto_imp_api_signing_signing_proto_goTypes = nil
	file_proto_imp_api_signing_signing_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SigningClient is the client API for Signing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SigningClient interface {
	//*
	// SignMessage signs a message with your node's private key.
	SignMessage(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
	//*
	// Verifymessage verifies a message was signed from another node.
	VerifySignature(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
}

type signingClient struct {
	cc grpc.ClientConnInterface
}

func NewSigningClient(cc grpc.ClientConnInterface) SigningClient {
	return &signingClient{cc}
}

func (c *signingClient) SignMessage(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/signing.Signing/SignMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) VerifySignature(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/signing.Signing/VerifySignature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SigningServer is the server API for Signing service.
type SigningServer interface {
	//*
	// SignMessage signs a message with your node's private key.
	SignMessage(context.Context, *SignRequest) (*SignResponse, error)
	//*
	// Verifymessage verifies a message was signed from another node.
	VerifySignature(context.Context, *VerifyRequest) (*VerifyResponse, error)
}

// UnimplementedSigningServer can be embedded to have forward compatible implementations.
type UnimplementedSigningServer struct {
}

func (*UnimplementedSigningServer) SignMessage(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignMessage not implemented")
}
func (*UnimplementedSigningServer) VerifySignature(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySignature not implemented")
}

func RegisterSigningServer(s *grpc.Server, srv SigningServer) {
	s.RegisterService(&_Signing_serviceDesc, srv)
}

func _Signing_SignMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).SignMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signing.Signing/SignMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).SignMessage(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_VerifySignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).VerifySignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signing.Signing/VerifySignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).VerifySignature(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Signing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "signing.Signing",
	HandlerType: (*SigningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignMessage",
			Handler:    _Signing_SignMessage_Handler,
		},
		{
			MethodName: "VerifySignature",
			Handler:    _Signing_VerifySignature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/imp/api/signing/signing.proto",
}
