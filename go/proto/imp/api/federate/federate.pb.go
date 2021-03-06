/// Allows for p2p federation between Impervious nodes

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/imp/api/federate/federate.proto

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
// Represents a request to federate with a far end node
type RequestFederateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pubkey string `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"` // The public key of the far end LND node running IMP
}

func (x *RequestFederateRequest) Reset() {
	*x = RequestFederateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_federate_federate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestFederateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestFederateRequest) ProtoMessage() {}

func (x *RequestFederateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_federate_federate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestFederateRequest.ProtoReflect.Descriptor instead.
func (*RequestFederateRequest) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_federate_federate_proto_rawDescGZIP(), []int{0}
}

func (x *RequestFederateRequest) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

//*
// Represents a response back from a Federation Request
type RequestFederateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // returned message ID
}

func (x *RequestFederateResponse) Reset() {
	*x = RequestFederateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_federate_federate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestFederateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestFederateResponse) ProtoMessage() {}

func (x *RequestFederateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_federate_federate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestFederateResponse.ProtoReflect.Descriptor instead.
func (*RequestFederateResponse) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_federate_federate_proto_rawDescGZIP(), []int{1}
}

func (x *RequestFederateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

//*
// Represents a request to leave a federation from a far end node
type LeaveFederationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pubkey string `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"` // The public key of the far end LND node running IMP
}

func (x *LeaveFederationRequest) Reset() {
	*x = LeaveFederationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_federate_federate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveFederationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveFederationRequest) ProtoMessage() {}

func (x *LeaveFederationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_federate_federate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveFederationRequest.ProtoReflect.Descriptor instead.
func (*LeaveFederationRequest) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_federate_federate_proto_rawDescGZIP(), []int{2}
}

func (x *LeaveFederationRequest) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

//*
// Represents a response back from a Leave Federation Request
type LeaveFederationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // returned message ID
}

func (x *LeaveFederationResponse) Reset() {
	*x = LeaveFederationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_federate_federate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveFederationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveFederationResponse) ProtoMessage() {}

func (x *LeaveFederationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_federate_federate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveFederationResponse.ProtoReflect.Descriptor instead.
func (*LeaveFederationResponse) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_federate_federate_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveFederationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_imp_api_federate_federate_proto protoreflect.FileDescriptor

var file_proto_imp_api_federate_federate_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x30, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65,
	0x79, 0x22, 0x29, 0x0a, 0x17, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x16,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x22, 0x29,
	0x0a, 0x17, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xfa, 0x01, 0x0a, 0x08, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x77, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12,
	0x75, 0x0a, 0x0f, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x65,
	0x61, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0xc8, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x61, 0x69, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x69, 0x6d, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x92, 0x41,
	0x9f, 0x01, 0x12, 0x40, 0x0a, 0x11, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x20, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x20, 0x41, 0x49, 0x12, 0x15, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x2e, 0x61, 0x69, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x2a, 0x03, 0x01, 0x02, 0x04, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x32, 0x0a,
	0x14, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f,
	0x6e, 0x20, 0x49, 0x4d, 0x50, 0x12, 0x1a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x64,
	0x6f, 0x63, 0x73, 0x2e, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x2e, 0x61,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_imp_api_federate_federate_proto_rawDescOnce sync.Once
	file_proto_imp_api_federate_federate_proto_rawDescData = file_proto_imp_api_federate_federate_proto_rawDesc
)

func file_proto_imp_api_federate_federate_proto_rawDescGZIP() []byte {
	file_proto_imp_api_federate_federate_proto_rawDescOnce.Do(func() {
		file_proto_imp_api_federate_federate_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_imp_api_federate_federate_proto_rawDescData)
	})
	return file_proto_imp_api_federate_federate_proto_rawDescData
}

var file_proto_imp_api_federate_federate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_imp_api_federate_federate_proto_goTypes = []interface{}{
	(*RequestFederateRequest)(nil),  // 0: federate.RequestFederateRequest
	(*RequestFederateResponse)(nil), // 1: federate.RequestFederateResponse
	(*LeaveFederationRequest)(nil),  // 2: federate.LeaveFederationRequest
	(*LeaveFederationResponse)(nil), // 3: federate.LeaveFederationResponse
}
var file_proto_imp_api_federate_federate_proto_depIdxs = []int32{
	0, // 0: federate.Federate.RequestFederate:input_type -> federate.RequestFederateRequest
	2, // 1: federate.Federate.LeaveFederation:input_type -> federate.LeaveFederationRequest
	1, // 2: federate.Federate.RequestFederate:output_type -> federate.RequestFederateResponse
	3, // 3: federate.Federate.LeaveFederation:output_type -> federate.LeaveFederationResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_imp_api_federate_federate_proto_init() }
func file_proto_imp_api_federate_federate_proto_init() {
	if File_proto_imp_api_federate_federate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_imp_api_federate_federate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestFederateRequest); i {
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
		file_proto_imp_api_federate_federate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestFederateResponse); i {
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
		file_proto_imp_api_federate_federate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveFederationRequest); i {
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
		file_proto_imp_api_federate_federate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveFederationResponse); i {
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
			RawDescriptor: file_proto_imp_api_federate_federate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_imp_api_federate_federate_proto_goTypes,
		DependencyIndexes: file_proto_imp_api_federate_federate_proto_depIdxs,
		MessageInfos:      file_proto_imp_api_federate_federate_proto_msgTypes,
	}.Build()
	File_proto_imp_api_federate_federate_proto = out.File
	file_proto_imp_api_federate_federate_proto_rawDesc = nil
	file_proto_imp_api_federate_federate_proto_goTypes = nil
	file_proto_imp_api_federate_federate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FederateClient is the client API for Federate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FederateClient interface {
	//*
	// RequestFederation performs the federation request to a specific peer.
	RequestFederate(ctx context.Context, in *RequestFederateRequest, opts ...grpc.CallOption) (*RequestFederateResponse, error)
	//*
	// LeaveFederation performs the removal of a federated peer (upon message receipt).
	LeaveFederation(ctx context.Context, in *LeaveFederationRequest, opts ...grpc.CallOption) (*LeaveFederationResponse, error)
}

type federateClient struct {
	cc grpc.ClientConnInterface
}

func NewFederateClient(cc grpc.ClientConnInterface) FederateClient {
	return &federateClient{cc}
}

func (c *federateClient) RequestFederate(ctx context.Context, in *RequestFederateRequest, opts ...grpc.CallOption) (*RequestFederateResponse, error) {
	out := new(RequestFederateResponse)
	err := c.cc.Invoke(ctx, "/federate.Federate/RequestFederate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federateClient) LeaveFederation(ctx context.Context, in *LeaveFederationRequest, opts ...grpc.CallOption) (*LeaveFederationResponse, error) {
	out := new(LeaveFederationResponse)
	err := c.cc.Invoke(ctx, "/federate.Federate/LeaveFederation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederateServer is the server API for Federate service.
type FederateServer interface {
	//*
	// RequestFederation performs the federation request to a specific peer.
	RequestFederate(context.Context, *RequestFederateRequest) (*RequestFederateResponse, error)
	//*
	// LeaveFederation performs the removal of a federated peer (upon message receipt).
	LeaveFederation(context.Context, *LeaveFederationRequest) (*LeaveFederationResponse, error)
}

// UnimplementedFederateServer can be embedded to have forward compatible implementations.
type UnimplementedFederateServer struct {
}

func (*UnimplementedFederateServer) RequestFederate(context.Context, *RequestFederateRequest) (*RequestFederateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestFederate not implemented")
}
func (*UnimplementedFederateServer) LeaveFederation(context.Context, *LeaveFederationRequest) (*LeaveFederationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveFederation not implemented")
}

func RegisterFederateServer(s *grpc.Server, srv FederateServer) {
	s.RegisterService(&_Federate_serviceDesc, srv)
}

func _Federate_RequestFederate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFederateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederateServer).RequestFederate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/federate.Federate/RequestFederate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederateServer).RequestFederate(ctx, req.(*RequestFederateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Federate_LeaveFederation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederateServer).LeaveFederation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/federate.Federate/LeaveFederation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederateServer).LeaveFederation(ctx, req.(*LeaveFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Federate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "federate.Federate",
	HandlerType: (*FederateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestFederate",
			Handler:    _Federate_RequestFederate_Handler,
		},
		{
			MethodName: "LeaveFederation",
			Handler:    _Federate_LeaveFederation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/imp/api/federate/federate.proto",
}
