// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: integralSettingService.proto

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

type IntegralSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Limit             int32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
	EffectiveType     int32   `protobuf:"varint,3,opt,name=effective_type,json=effectiveType,proto3" json:"effective_type"`
	ExpireYearNum     int32   `protobuf:"varint,4,opt,name=expire_year_num,json=expireYearNum,proto3" json:"expire_year_num"`
	ExpireAt          string  `protobuf:"bytes,6,opt,name=expire_at,json=expireAt,proto3" json:"expire_at"`
	ProtectDays       int32   `protobuf:"varint,7,opt,name=protect_days,json=protectDays,proto3" json:"protect_days"`
	IntegralName      string  `protobuf:"bytes,8,opt,name=integral_name,json=integralName,proto3" json:"integral_name"`
	IntegralRatio     int32   `protobuf:"varint,9,opt,name=integral_ratio,json=integralRatio,proto3" json:"integral_ratio"`
	IsCashExchange    bool    `protobuf:"varint,10,opt,name=is_cash_exchange,json=isCashExchange,proto3" json:"is_cash_exchange"`
	CashExchangeMoney float32 `protobuf:"fixed32,11,opt,name=cash_exchange_money,json=cashExchangeMoney,proto3" json:"cash_exchange_money"`
	CashExchangeLimit float32 `protobuf:"fixed32,12,opt,name=cash_exchange_limit,json=cashExchangeLimit,proto3" json:"cash_exchange_limit"`
	CreatedAt         string  `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt         string  `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	CashExchangeType  int32   `protobuf:"varint,15,opt,name=cash_exchange_type,json=cashExchangeType,proto3" json:"cash_exchange_type"`
	CashExchangeRate  float32 `protobuf:"fixed32,16,opt,name=cash_exchange_rate,json=cashExchangeRate,proto3" json:"cash_exchange_rate"`
}

func (x *IntegralSetting) Reset() {
	*x = IntegralSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralSettingService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralSetting) ProtoMessage() {}

func (x *IntegralSetting) ProtoReflect() protoreflect.Message {
	mi := &file_integralSettingService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegralSetting.ProtoReflect.Descriptor instead.
func (*IntegralSetting) Descriptor() ([]byte, []int) {
	return file_integralSettingService_proto_rawDescGZIP(), []int{0}
}

func (x *IntegralSetting) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IntegralSetting) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *IntegralSetting) GetEffectiveType() int32 {
	if x != nil {
		return x.EffectiveType
	}
	return 0
}

func (x *IntegralSetting) GetExpireYearNum() int32 {
	if x != nil {
		return x.ExpireYearNum
	}
	return 0
}

func (x *IntegralSetting) GetExpireAt() string {
	if x != nil {
		return x.ExpireAt
	}
	return ""
}

func (x *IntegralSetting) GetProtectDays() int32 {
	if x != nil {
		return x.ProtectDays
	}
	return 0
}

func (x *IntegralSetting) GetIntegralName() string {
	if x != nil {
		return x.IntegralName
	}
	return ""
}

func (x *IntegralSetting) GetIntegralRatio() int32 {
	if x != nil {
		return x.IntegralRatio
	}
	return 0
}

func (x *IntegralSetting) GetIsCashExchange() bool {
	if x != nil {
		return x.IsCashExchange
	}
	return false
}

func (x *IntegralSetting) GetCashExchangeMoney() float32 {
	if x != nil {
		return x.CashExchangeMoney
	}
	return 0
}

func (x *IntegralSetting) GetCashExchangeLimit() float32 {
	if x != nil {
		return x.CashExchangeLimit
	}
	return 0
}

func (x *IntegralSetting) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *IntegralSetting) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *IntegralSetting) GetCashExchangeType() int32 {
	if x != nil {
		return x.CashExchangeType
	}
	return 0
}

func (x *IntegralSetting) GetCashExchangeRate() float32 {
	if x != nil {
		return x.CashExchangeRate
	}
	return 0
}

type IntegralSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *IntegralSetting   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*IntegralSetting `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *IntegralSettingResponse) Reset() {
	*x = IntegralSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralSettingService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralSettingResponse) ProtoMessage() {}

func (x *IntegralSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integralSettingService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegralSettingResponse.ProtoReflect.Descriptor instead.
func (*IntegralSettingResponse) Descriptor() ([]byte, []int) {
	return file_integralSettingService_proto_rawDescGZIP(), []int{1}
}

func (x *IntegralSettingResponse) GetEntity() *IntegralSetting {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *IntegralSettingResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *IntegralSettingResponse) GetItems() []*IntegralSetting {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *IntegralSettingResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *IntegralSettingResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_integralSettingService_proto protoreflect.FileDescriptor

var file_integralSettingService_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x04, 0x0a, 0x0f,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x59, 0x65, 0x61,
	0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x61, 0x79,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x44, 0x61, 0x79, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x73, 0x68, 0x5f, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x43, 0x61,
	0x73, 0x68, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x61,
	0x73, 0x68, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x63, 0x61, 0x73, 0x68, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x61,
	0x73, 0x68, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x63, 0x61, 0x73, 0x68, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x61, 0x73, 0x68,
	0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x63, 0x61, 0x73, 0x68, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x61, 0x73, 0x68, 0x5f, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x10, 0x63, 0x61, 0x73, 0x68, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x22, 0xe9, 0x01, 0x0a, 0x17, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x32, 0xb3, 0x04, 0x0a, 0x16, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x21, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x21, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4a, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x0a, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x07,
	0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x52, 0x61,
	0x74, 0x69, 0x6f, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x21,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x43, 0x61, 0x73, 0x68, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integralSettingService_proto_rawDescOnce sync.Once
	file_integralSettingService_proto_rawDescData = file_integralSettingService_proto_rawDesc
)

func file_integralSettingService_proto_rawDescGZIP() []byte {
	file_integralSettingService_proto_rawDescOnce.Do(func() {
		file_integralSettingService_proto_rawDescData = protoimpl.X.CompressGZIP(file_integralSettingService_proto_rawDescData)
	})
	return file_integralSettingService_proto_rawDescData
}

var file_integralSettingService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_integralSettingService_proto_goTypes = []interface{}{
	(*IntegralSetting)(nil),         // 0: services.IntegralSetting
	(*IntegralSettingResponse)(nil), // 1: services.IntegralSettingResponse
	(*common.Pager)(nil),            // 2: common.Pager
	(*common.Error)(nil),            // 3: common.Error
	(*common.Info)(nil),             // 4: common.Info
}
var file_integralSettingService_proto_depIdxs = []int32{
	0,  // 0: services.IntegralSettingResponse.entity:type_name -> services.IntegralSetting
	2,  // 1: services.IntegralSettingResponse.pager:type_name -> common.Pager
	0,  // 2: services.IntegralSettingResponse.items:type_name -> services.IntegralSetting
	3,  // 3: services.IntegralSettingResponse.error:type_name -> common.Error
	4,  // 4: services.IntegralSettingResponse.info:type_name -> common.Info
	0,  // 5: services.IntegralSettingService.Get:input_type -> services.IntegralSetting
	0,  // 6: services.IntegralSettingService.SetEffective:input_type -> services.IntegralSetting
	0,  // 7: services.IntegralSettingService.SetLimit:input_type -> services.IntegralSetting
	0,  // 8: services.IntegralSettingService.SetProtect:input_type -> services.IntegralSetting
	0,  // 9: services.IntegralSettingService.SetName:input_type -> services.IntegralSetting
	0,  // 10: services.IntegralSettingService.SetRatio:input_type -> services.IntegralSetting
	0,  // 11: services.IntegralSettingService.SetCashExchange:input_type -> services.IntegralSetting
	1,  // 12: services.IntegralSettingService.Get:output_type -> services.IntegralSettingResponse
	1,  // 13: services.IntegralSettingService.SetEffective:output_type -> services.IntegralSettingResponse
	1,  // 14: services.IntegralSettingService.SetLimit:output_type -> services.IntegralSettingResponse
	1,  // 15: services.IntegralSettingService.SetProtect:output_type -> services.IntegralSettingResponse
	1,  // 16: services.IntegralSettingService.SetName:output_type -> services.IntegralSettingResponse
	1,  // 17: services.IntegralSettingService.SetRatio:output_type -> services.IntegralSettingResponse
	1,  // 18: services.IntegralSettingService.SetCashExchange:output_type -> services.IntegralSettingResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_integralSettingService_proto_init() }
func file_integralSettingService_proto_init() {
	if File_integralSettingService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integralSettingService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegralSetting); i {
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
		file_integralSettingService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegralSettingResponse); i {
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
			RawDescriptor: file_integralSettingService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_integralSettingService_proto_goTypes,
		DependencyIndexes: file_integralSettingService_proto_depIdxs,
		MessageInfos:      file_integralSettingService_proto_msgTypes,
	}.Build()
	File_integralSettingService_proto = out.File
	file_integralSettingService_proto_rawDesc = nil
	file_integralSettingService_proto_goTypes = nil
	file_integralSettingService_proto_depIdxs = nil
}