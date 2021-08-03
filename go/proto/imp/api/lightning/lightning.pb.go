/// Allows for p2p messaging between Impervious nodes

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/imp/api/lightning/lightning.proto

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
// Represents an invoice creation request from your lightning node.
type GenerateInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"` // The amount of satoshis you want to receive
	Memo   string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`      // The human readable memo you want the sender to see
}

func (x *GenerateInvoiceRequest) Reset() {
	*x = GenerateInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateInvoiceRequest) ProtoMessage() {}

func (x *GenerateInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateInvoiceRequest.ProtoReflect.Descriptor instead.
func (*GenerateInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_lightning_lightning_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateInvoiceRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *GenerateInvoiceRequest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

//*
// Represents a response back from an invoice generation request.
type GenerateInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invoice string `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"` // The invoice from your lightning node
}

func (x *GenerateInvoiceResponse) Reset() {
	*x = GenerateInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateInvoiceResponse) ProtoMessage() {}

func (x *GenerateInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateInvoiceResponse.ProtoReflect.Descriptor instead.
func (*GenerateInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_lightning_lightning_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateInvoiceResponse) GetInvoice() string {
	if x != nil {
		return x.Invoice
	}
	return ""
}

//*
// Represents an invoice that will be paid by your lightning node.
type PayInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invoice string `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"` // The invoice to pay
}

func (x *PayInvoiceRequest) Reset() {
	*x = PayInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayInvoiceRequest) ProtoMessage() {}

func (x *PayInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayInvoiceRequest.ProtoReflect.Descriptor instead.
func (*PayInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_lightning_lightning_proto_rawDescGZIP(), []int{2}
}

func (x *PayInvoiceRequest) GetInvoice() string {
	if x != nil {
		return x.Invoice
	}
	return ""
}

//*
// Represents a response back from the payment result.
type PayInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Preimage string `protobuf:"bytes,1,opt,name=preimage,proto3" json:"preimage,omitempty"` // The preimage from the payment result, if successful
}

func (x *PayInvoiceResponse) Reset() {
	*x = PayInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayInvoiceResponse) ProtoMessage() {}

func (x *PayInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayInvoiceResponse.ProtoReflect.Descriptor instead.
func (*PayInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_lightning_lightning_proto_rawDescGZIP(), []int{3}
}

func (x *PayInvoiceResponse) GetPreimage() string {
	if x != nil {
		return x.Preimage
	}
	return ""
}

//*
// Represents an request to check an invoice.
type CheckInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invoice string `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"` // The invoice to check
}

func (x *CheckInvoiceRequest) Reset() {
	*x = CheckInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInvoiceRequest) ProtoMessage() {}

func (x *CheckInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInvoiceRequest.ProtoReflect.Descriptor instead.
func (*CheckInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_lightning_lightning_proto_rawDescGZIP(), []int{4}
}

func (x *CheckInvoiceRequest) GetInvoice() string {
	if x != nil {
		return x.Invoice
	}
	return ""
}

//*
// Represents a response back from the invoice check.
type CheckInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paid bool `protobuf:"varint,1,opt,name=paid,proto3" json:"paid,omitempty"` // The boolean result representing whether or not the invoice was paid
}

func (x *CheckInvoiceResponse) Reset() {
	*x = CheckInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInvoiceResponse) ProtoMessage() {}

func (x *CheckInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imp_api_lightning_lightning_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInvoiceResponse.ProtoReflect.Descriptor instead.
func (*CheckInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_imp_api_lightning_lightning_proto_rawDescGZIP(), []int{5}
}

func (x *CheckInvoiceResponse) GetPaid() bool {
	if x != nil {
		return x.Paid
	}
	return false
}

var File_proto_imp_api_lightning_lightning_proto protoreflect.FileDescriptor

var file_proto_imp_api_lightning_lightning_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x44, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x33, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x2d, 0x0a,
	0x11, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x30, 0x0a, 0x12,
	0x50, 0x61, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x2f,
	0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22,
	0x2a, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x70, 0x61, 0x69, 0x64, 0x32, 0xf8, 0x02, 0x0a, 0x09,
	0x4c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x82, 0x01, 0x0a, 0x0f, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x21, 0x2e,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6e,
	0x0a, 0x0a, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x76,
	0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1e,
	0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0xdb, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x61, 0x69, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x69, 0x6d, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x92, 0x41,
	0xb2, 0x01, 0x12, 0x41, 0x0a, 0x12, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x20,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x65,
	0x72, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x20, 0x41, 0x49, 0x12, 0x15, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x2e, 0x61, 0x69,
	0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x03, 0x01, 0x02, 0x04, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x44,
	0x0a, 0x14, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x6f, 0x6e, 0x20, 0x49, 0x4d, 0x50, 0x12, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x61, 0x69, 0x2f, 0x69, 0x6d, 0x70, 0x2d, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_imp_api_lightning_lightning_proto_rawDescOnce sync.Once
	file_proto_imp_api_lightning_lightning_proto_rawDescData = file_proto_imp_api_lightning_lightning_proto_rawDesc
)

func file_proto_imp_api_lightning_lightning_proto_rawDescGZIP() []byte {
	file_proto_imp_api_lightning_lightning_proto_rawDescOnce.Do(func() {
		file_proto_imp_api_lightning_lightning_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_imp_api_lightning_lightning_proto_rawDescData)
	})
	return file_proto_imp_api_lightning_lightning_proto_rawDescData
}

var file_proto_imp_api_lightning_lightning_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_imp_api_lightning_lightning_proto_goTypes = []interface{}{
	(*GenerateInvoiceRequest)(nil),  // 0: lightning.GenerateInvoiceRequest
	(*GenerateInvoiceResponse)(nil), // 1: lightning.GenerateInvoiceResponse
	(*PayInvoiceRequest)(nil),       // 2: lightning.PayInvoiceRequest
	(*PayInvoiceResponse)(nil),      // 3: lightning.PayInvoiceResponse
	(*CheckInvoiceRequest)(nil),     // 4: lightning.CheckInvoiceRequest
	(*CheckInvoiceResponse)(nil),    // 5: lightning.CheckInvoiceResponse
}
var file_proto_imp_api_lightning_lightning_proto_depIdxs = []int32{
	0, // 0: lightning.Lightning.GenerateInvoice:input_type -> lightning.GenerateInvoiceRequest
	2, // 1: lightning.Lightning.PayInvoice:input_type -> lightning.PayInvoiceRequest
	4, // 2: lightning.Lightning.CheckInvoice:input_type -> lightning.CheckInvoiceRequest
	1, // 3: lightning.Lightning.GenerateInvoice:output_type -> lightning.GenerateInvoiceResponse
	3, // 4: lightning.Lightning.PayInvoice:output_type -> lightning.PayInvoiceResponse
	5, // 5: lightning.Lightning.CheckInvoice:output_type -> lightning.CheckInvoiceResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_imp_api_lightning_lightning_proto_init() }
func file_proto_imp_api_lightning_lightning_proto_init() {
	if File_proto_imp_api_lightning_lightning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_imp_api_lightning_lightning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateInvoiceRequest); i {
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
		file_proto_imp_api_lightning_lightning_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateInvoiceResponse); i {
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
		file_proto_imp_api_lightning_lightning_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayInvoiceRequest); i {
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
		file_proto_imp_api_lightning_lightning_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayInvoiceResponse); i {
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
		file_proto_imp_api_lightning_lightning_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInvoiceRequest); i {
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
		file_proto_imp_api_lightning_lightning_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInvoiceResponse); i {
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
			RawDescriptor: file_proto_imp_api_lightning_lightning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_imp_api_lightning_lightning_proto_goTypes,
		DependencyIndexes: file_proto_imp_api_lightning_lightning_proto_depIdxs,
		MessageInfos:      file_proto_imp_api_lightning_lightning_proto_msgTypes,
	}.Build()
	File_proto_imp_api_lightning_lightning_proto = out.File
	file_proto_imp_api_lightning_lightning_proto_rawDesc = nil
	file_proto_imp_api_lightning_lightning_proto_goTypes = nil
	file_proto_imp_api_lightning_lightning_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LightningClient is the client API for Lightning service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LightningClient interface {
	GenerateInvoice(ctx context.Context, in *GenerateInvoiceRequest, opts ...grpc.CallOption) (*GenerateInvoiceResponse, error)
	PayInvoice(ctx context.Context, in *PayInvoiceRequest, opts ...grpc.CallOption) (*PayInvoiceResponse, error)
	CheckInvoice(ctx context.Context, in *CheckInvoiceRequest, opts ...grpc.CallOption) (*CheckInvoiceResponse, error)
}

type lightningClient struct {
	cc grpc.ClientConnInterface
}

func NewLightningClient(cc grpc.ClientConnInterface) LightningClient {
	return &lightningClient{cc}
}

func (c *lightningClient) GenerateInvoice(ctx context.Context, in *GenerateInvoiceRequest, opts ...grpc.CallOption) (*GenerateInvoiceResponse, error) {
	out := new(GenerateInvoiceResponse)
	err := c.cc.Invoke(ctx, "/lightning.Lightning/GenerateInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningClient) PayInvoice(ctx context.Context, in *PayInvoiceRequest, opts ...grpc.CallOption) (*PayInvoiceResponse, error) {
	out := new(PayInvoiceResponse)
	err := c.cc.Invoke(ctx, "/lightning.Lightning/PayInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningClient) CheckInvoice(ctx context.Context, in *CheckInvoiceRequest, opts ...grpc.CallOption) (*CheckInvoiceResponse, error) {
	out := new(CheckInvoiceResponse)
	err := c.cc.Invoke(ctx, "/lightning.Lightning/CheckInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LightningServer is the server API for Lightning service.
type LightningServer interface {
	GenerateInvoice(context.Context, *GenerateInvoiceRequest) (*GenerateInvoiceResponse, error)
	PayInvoice(context.Context, *PayInvoiceRequest) (*PayInvoiceResponse, error)
	CheckInvoice(context.Context, *CheckInvoiceRequest) (*CheckInvoiceResponse, error)
}

// UnimplementedLightningServer can be embedded to have forward compatible implementations.
type UnimplementedLightningServer struct {
}

func (*UnimplementedLightningServer) GenerateInvoice(context.Context, *GenerateInvoiceRequest) (*GenerateInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInvoice not implemented")
}
func (*UnimplementedLightningServer) PayInvoice(context.Context, *PayInvoiceRequest) (*PayInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayInvoice not implemented")
}
func (*UnimplementedLightningServer) CheckInvoice(context.Context, *CheckInvoiceRequest) (*CheckInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckInvoice not implemented")
}

func RegisterLightningServer(s *grpc.Server, srv LightningServer) {
	s.RegisterService(&_Lightning_serviceDesc, srv)
}

func _Lightning_GenerateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServer).GenerateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lightning.Lightning/GenerateInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServer).GenerateInvoice(ctx, req.(*GenerateInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lightning_PayInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServer).PayInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lightning.Lightning/PayInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServer).PayInvoice(ctx, req.(*PayInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lightning_CheckInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServer).CheckInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lightning.Lightning/CheckInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServer).CheckInvoice(ctx, req.(*CheckInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lightning_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lightning.Lightning",
	HandlerType: (*LightningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateInvoice",
			Handler:    _Lightning_GenerateInvoice_Handler,
		},
		{
			MethodName: "PayInvoice",
			Handler:    _Lightning_PayInvoice_Handler,
		},
		{
			MethodName: "CheckInvoice",
			Handler:    _Lightning_CheckInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/imp/api/lightning/lightning.proto",
}