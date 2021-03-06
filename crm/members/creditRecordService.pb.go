// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: creditRecordService.proto

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

type CreditRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CreditId   int64   `protobuf:"varint,2,opt,name=credit_id,json=creditId,proto3" json:"credit_id"`
	CustomerId int64   `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	Direction  string  `protobuf:"bytes,4,opt,name=direction,proto3" json:"direction"`
	Amount     float32 `protobuf:"fixed32,5,opt,name=amount,proto3" json:"amount"`
	Type       int32   `protobuf:"varint,6,opt,name=type,proto3" json:"type"`
	OrderId    int64   `protobuf:"varint,7,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	OrderSn    string  `protobuf:"bytes,8,opt,name=order_sn,json=orderSn,proto3" json:"order_sn"`
	Balance    float32 `protobuf:"fixed32,9,opt,name=balance,proto3" json:"balance"`
	Memo       string  `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo"`
	CreatedAt  string  `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string  `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *CreditRecord) Reset() {
	*x = CreditRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditRecordService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditRecord) ProtoMessage() {}

func (x *CreditRecord) ProtoReflect() protoreflect.Message {
	mi := &file_creditRecordService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditRecord.ProtoReflect.Descriptor instead.
func (*CreditRecord) Descriptor() ([]byte, []int) {
	return file_creditRecordService_proto_rawDescGZIP(), []int{0}
}

func (x *CreditRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreditRecord) GetCreditId() int64 {
	if x != nil {
		return x.CreditId
	}
	return 0
}

func (x *CreditRecord) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CreditRecord) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *CreditRecord) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreditRecord) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CreditRecord) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *CreditRecord) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *CreditRecord) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *CreditRecord) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *CreditRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CreditRecord) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//????????????
type CreditRecordWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	//????????????????????????
	CustomerId int64  `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	CustomerSn string `protobuf:"bytes,4,opt,name=customer_sn,json=customerSn,proto3" json:"customer_sn"`
	Name       string `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Mobile     string `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile"`
}

func (x *CreditRecordWhere) Reset() {
	*x = CreditRecordWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditRecordService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditRecordWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditRecordWhere) ProtoMessage() {}

func (x *CreditRecordWhere) ProtoReflect() protoreflect.Message {
	mi := &file_creditRecordService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditRecordWhere.ProtoReflect.Descriptor instead.
func (*CreditRecordWhere) Descriptor() ([]byte, []int) {
	return file_creditRecordService_proto_rawDescGZIP(), []int{1}
}

func (x *CreditRecordWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CreditRecordWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CreditRecordWhere) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CreditRecordWhere) GetCustomerSn() string {
	if x != nil {
		return x.CustomerSn
	}
	return ""
}

func (x *CreditRecordWhere) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreditRecordWhere) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

//
type CreditRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *CreditRecord   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*CreditRecord `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *CreditRecordResponse) Reset() {
	*x = CreditRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditRecordService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditRecordResponse) ProtoMessage() {}

func (x *CreditRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_creditRecordService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditRecordResponse.ProtoReflect.Descriptor instead.
func (*CreditRecordResponse) Descriptor() ([]byte, []int) {
	return file_creditRecordService_proto_rawDescGZIP(), []int{2}
}

func (x *CreditRecordResponse) GetEntity() *CreditRecord {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CreditRecordResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CreditRecordResponse) GetItems() []*CreditRecord {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreditRecordResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CreditRecordResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_creditRecordService_proto protoreflect.FileDescriptor

var file_creditRecordService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xb4, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0xe0, 0x01, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xa7, 0x02,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x45, 0x78, 0x70,
	0x65, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creditRecordService_proto_rawDescOnce sync.Once
	file_creditRecordService_proto_rawDescData = file_creditRecordService_proto_rawDesc
)

func file_creditRecordService_proto_rawDescGZIP() []byte {
	file_creditRecordService_proto_rawDescOnce.Do(func() {
		file_creditRecordService_proto_rawDescData = protoimpl.X.CompressGZIP(file_creditRecordService_proto_rawDescData)
	})
	return file_creditRecordService_proto_rawDescData
}

var file_creditRecordService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_creditRecordService_proto_goTypes = []interface{}{
	(*CreditRecord)(nil),         // 0: services.CreditRecord
	(*CreditRecordWhere)(nil),    // 1: services.CreditRecordWhere
	(*CreditRecordResponse)(nil), // 2: services.CreditRecordResponse
	(*common.Pager)(nil),         // 3: common.Pager
	(*common.Error)(nil),         // 4: common.Error
	(*common.Info)(nil),          // 5: common.Info
}
var file_creditRecordService_proto_depIdxs = []int32{
	0, // 0: services.CreditRecordResponse.entity:type_name -> services.CreditRecord
	3, // 1: services.CreditRecordResponse.pager:type_name -> common.Pager
	0, // 2: services.CreditRecordResponse.items:type_name -> services.CreditRecord
	4, // 3: services.CreditRecordResponse.error:type_name -> common.Error
	5, // 4: services.CreditRecordResponse.info:type_name -> common.Info
	0, // 5: services.CreditRecordService.Income:input_type -> services.CreditRecord
	0, // 6: services.CreditRecordService.Expend:input_type -> services.CreditRecord
	0, // 7: services.CreditRecordService.Get:input_type -> services.CreditRecord
	1, // 8: services.CreditRecordService.Search:input_type -> services.CreditRecordWhere
	2, // 9: services.CreditRecordService.Income:output_type -> services.CreditRecordResponse
	2, // 10: services.CreditRecordService.Expend:output_type -> services.CreditRecordResponse
	2, // 11: services.CreditRecordService.Get:output_type -> services.CreditRecordResponse
	2, // 12: services.CreditRecordService.Search:output_type -> services.CreditRecordResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_creditRecordService_proto_init() }
func file_creditRecordService_proto_init() {
	if File_creditRecordService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_creditRecordService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditRecord); i {
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
		file_creditRecordService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditRecordWhere); i {
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
		file_creditRecordService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditRecordResponse); i {
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
			RawDescriptor: file_creditRecordService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_creditRecordService_proto_goTypes,
		DependencyIndexes: file_creditRecordService_proto_depIdxs,
		MessageInfos:      file_creditRecordService_proto_msgTypes,
	}.Build()
	File_creditRecordService_proto = out.File
	file_creditRecordService_proto_rawDesc = nil
	file_creditRecordService_proto_goTypes = nil
	file_creditRecordService_proto_depIdxs = nil
}
