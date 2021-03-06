// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: customerBankService.proto

package services

import (
	common "github.com/geiqin/micro-kit/protobuf/common"
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

type CustomerBank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CustomerId    int64  `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	Cardholder    string `protobuf:"bytes,3,opt,name=cardholder,proto3" json:"cardholder"`
	CardNo        string `protobuf:"bytes,4,opt,name=card_no,json=cardNo,proto3" json:"card_no"`
	BankName      string `protobuf:"bytes,5,opt,name=bank_name,json=bankName,proto3" json:"bank_name"`
	ReserveMobile string `protobuf:"bytes,6,opt,name=reserve_mobile,json=reserveMobile,proto3" json:"reserve_mobile"`
	Status        int32  `protobuf:"varint,7,opt,name=status,proto3" json:"status"`
}

func (x *CustomerBank) Reset() {
	*x = CustomerBank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerBankService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerBank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerBank) ProtoMessage() {}

func (x *CustomerBank) ProtoReflect() protoreflect.Message {
	mi := &file_customerBankService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerBank.ProtoReflect.Descriptor instead.
func (*CustomerBank) Descriptor() ([]byte, []int) {
	return file_customerBankService_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerBank) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomerBank) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CustomerBank) GetCardholder() string {
	if x != nil {
		return x.Cardholder
	}
	return ""
}

func (x *CustomerBank) GetCardNo() string {
	if x != nil {
		return x.CardNo
	}
	return ""
}

func (x *CustomerBank) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *CustomerBank) GetReserveMobile() string {
	if x != nil {
		return x.ReserveMobile
	}
	return ""
}

func (x *CustomerBank) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CustomerBankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize   int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id         int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	CustomerId int64  `protobuf:"varint,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	Cardholder string `protobuf:"bytes,5,opt,name=cardholder,proto3" json:"cardholder"`
	CardNo     string `protobuf:"bytes,6,opt,name=card_no,json=cardNo,proto3" json:"card_no"`
	BankName   string `protobuf:"bytes,7,opt,name=bank_name,json=bankName,proto3" json:"bank_name"`
	Status     int32  `protobuf:"varint,8,opt,name=status,proto3" json:"status"`
}

func (x *CustomerBankRequest) Reset() {
	*x = CustomerBankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerBankService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerBankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerBankRequest) ProtoMessage() {}

func (x *CustomerBankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customerBankService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerBankRequest.ProtoReflect.Descriptor instead.
func (*CustomerBankRequest) Descriptor() ([]byte, []int) {
	return file_customerBankService_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerBankRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CustomerBankRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CustomerBankRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomerBankRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CustomerBankRequest) GetCardholder() string {
	if x != nil {
		return x.Cardholder
	}
	return ""
}

func (x *CustomerBankRequest) GetCardNo() string {
	if x != nil {
		return x.CardNo
	}
	return ""
}

func (x *CustomerBankRequest) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *CustomerBankRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

//
type CustomerBankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *CustomerBank   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*CustomerBank `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *CustomerBankResponse) Reset() {
	*x = CustomerBankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerBankService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerBankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerBankResponse) ProtoMessage() {}

func (x *CustomerBankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customerBankService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerBankResponse.ProtoReflect.Descriptor instead.
func (*CustomerBankResponse) Descriptor() ([]byte, []int) {
	return file_customerBankService_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerBankResponse) GetEntity() *CustomerBank {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CustomerBankResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CustomerBankResponse) GetItems() []*CustomerBank {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CustomerBankResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CustomerBankResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_customerBankService_proto protoreflect.FileDescriptor

var file_customerBankService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x0c, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61,
	0x72, 0x64, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x72, 0x64, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72,
	0x64, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xe7, 0x01, 0x0a, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x61, 0x72, 0x64, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x72, 0x64, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x6e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61,
	0x72, 0x64, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x14, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xa9, 0x02, 0x0a,
	0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x1a, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customerBankService_proto_rawDescOnce sync.Once
	file_customerBankService_proto_rawDescData = file_customerBankService_proto_rawDesc
)

func file_customerBankService_proto_rawDescGZIP() []byte {
	file_customerBankService_proto_rawDescOnce.Do(func() {
		file_customerBankService_proto_rawDescData = protoimpl.X.CompressGZIP(file_customerBankService_proto_rawDescData)
	})
	return file_customerBankService_proto_rawDescData
}

var file_customerBankService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_customerBankService_proto_goTypes = []interface{}{
	(*CustomerBank)(nil),         // 0: services.CustomerBank
	(*CustomerBankRequest)(nil),  // 1: services.CustomerBankRequest
	(*CustomerBankResponse)(nil), // 2: services.CustomerBankResponse
	(*common.Pager)(nil),         // 3: common.Pager
	(*common.Error)(nil),         // 4: common.Error
	(*common.Info)(nil),          // 5: common.Info
}
var file_customerBankService_proto_depIdxs = []int32{
	0, // 0: services.CustomerBankResponse.entity:type_name -> services.CustomerBank
	3, // 1: services.CustomerBankResponse.pager:type_name -> common.Pager
	0, // 2: services.CustomerBankResponse.items:type_name -> services.CustomerBank
	4, // 3: services.CustomerBankResponse.error:type_name -> common.Error
	5, // 4: services.CustomerBankResponse.info:type_name -> common.Info
	0, // 5: services.CustomerBankService.Create:input_type -> services.CustomerBank
	0, // 6: services.CustomerBankService.Delete:input_type -> services.CustomerBank
	0, // 7: services.CustomerBankService.Get:input_type -> services.CustomerBank
	1, // 8: services.CustomerBankService.Search:input_type -> services.CustomerBankRequest
	2, // 9: services.CustomerBankService.Create:output_type -> services.CustomerBankResponse
	2, // 10: services.CustomerBankService.Delete:output_type -> services.CustomerBankResponse
	2, // 11: services.CustomerBankService.Get:output_type -> services.CustomerBankResponse
	2, // 12: services.CustomerBankService.Search:output_type -> services.CustomerBankResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_customerBankService_proto_init() }
func file_customerBankService_proto_init() {
	if File_customerBankService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customerBankService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerBank); i {
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
		file_customerBankService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerBankRequest); i {
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
		file_customerBankService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerBankResponse); i {
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
			RawDescriptor: file_customerBankService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customerBankService_proto_goTypes,
		DependencyIndexes: file_customerBankService_proto_depIdxs,
		MessageInfos:      file_customerBankService_proto_msgTypes,
	}.Build()
	File_customerBankService_proto = out.File
	file_customerBankService_proto_rawDesc = nil
	file_customerBankService_proto_goTypes = nil
	file_customerBankService_proto_depIdxs = nil
}
