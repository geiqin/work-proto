// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: commissionDetailService.proto

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

type CommissionDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CommissionId    int64   `protobuf:"varint,2,opt,name=commission_id,json=commissionId,proto3" json:"commission_id"`
	OrderDetailId   int64   `protobuf:"varint,3,opt,name=order_detail_id,json=orderDetailId,proto3" json:"order_detail_id"`
	ItemId          int64   `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	SkuId           int64   `protobuf:"varint,5,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Name            string  `protobuf:"bytes,6,opt,name=name,proto3" json:"name"`
	ThumbUrl        string  `protobuf:"bytes,7,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url"`
	Num             int32   `protobuf:"varint,8,opt,name=num,proto3" json:"num"`
	Price           float32 `protobuf:"fixed32,9,opt,name=price,proto3" json:"price"`
	CommissionRate  float32 `protobuf:"fixed32,10,opt,name=commission_rate,json=commissionRate,proto3" json:"commission_rate"`
	CommissionMoney float32 `protobuf:"fixed32,11,opt,name=commission_money,json=commissionMoney,proto3" json:"commission_money"`
	CreatedAt       string  `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string  `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *CommissionDetail) Reset() {
	*x = CommissionDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commissionDetailService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommissionDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommissionDetail) ProtoMessage() {}

func (x *CommissionDetail) ProtoReflect() protoreflect.Message {
	mi := &file_commissionDetailService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommissionDetail.ProtoReflect.Descriptor instead.
func (*CommissionDetail) Descriptor() ([]byte, []int) {
	return file_commissionDetailService_proto_rawDescGZIP(), []int{0}
}

func (x *CommissionDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommissionDetail) GetCommissionId() int64 {
	if x != nil {
		return x.CommissionId
	}
	return 0
}

func (x *CommissionDetail) GetOrderDetailId() int64 {
	if x != nil {
		return x.OrderDetailId
	}
	return 0
}

func (x *CommissionDetail) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *CommissionDetail) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *CommissionDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommissionDetail) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

func (x *CommissionDetail) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *CommissionDetail) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CommissionDetail) GetCommissionRate() float32 {
	if x != nil {
		return x.CommissionRate
	}
	return 0
}

func (x *CommissionDetail) GetCommissionMoney() float32 {
	if x != nil {
		return x.CommissionMoney
	}
	return 0
}

func (x *CommissionDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CommissionDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CommissionDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *CommissionDetail   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*CommissionDetail `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error       `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *CommissionDetailResponse) Reset() {
	*x = CommissionDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commissionDetailService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommissionDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommissionDetailResponse) ProtoMessage() {}

func (x *CommissionDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commissionDetailService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommissionDetailResponse.ProtoReflect.Descriptor instead.
func (*CommissionDetailResponse) Descriptor() ([]byte, []int) {
	return file_commissionDetailService_proto_rawDescGZIP(), []int{1}
}

func (x *CommissionDetailResponse) GetEntity() *CommissionDetail {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CommissionDetailResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CommissionDetailResponse) GetItems() []*CommissionDetail {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CommissionDetailResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CommissionDetailResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_commissionDetailService_proto protoreflect.FileDescriptor

var file_commissionDetailService_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xec, 0x01, 0x0a, 0x18, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x89, 0x03, 0x0a, 0x17, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x22, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x22,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commissionDetailService_proto_rawDescOnce sync.Once
	file_commissionDetailService_proto_rawDescData = file_commissionDetailService_proto_rawDesc
)

func file_commissionDetailService_proto_rawDescGZIP() []byte {
	file_commissionDetailService_proto_rawDescOnce.Do(func() {
		file_commissionDetailService_proto_rawDescData = protoimpl.X.CompressGZIP(file_commissionDetailService_proto_rawDescData)
	})
	return file_commissionDetailService_proto_rawDescData
}

var file_commissionDetailService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commissionDetailService_proto_goTypes = []interface{}{
	(*CommissionDetail)(nil),         // 0: services.CommissionDetail
	(*CommissionDetailResponse)(nil), // 1: services.CommissionDetailResponse
	(*common.Pager)(nil),             // 2: common.Pager
	(*common.Error)(nil),             // 3: common.Error
	(*common.Info)(nil),              // 4: common.Info
	(*common.BaseWhere)(nil),         // 5: common.BaseWhere
}
var file_commissionDetailService_proto_depIdxs = []int32{
	0,  // 0: services.CommissionDetailResponse.entity:type_name -> services.CommissionDetail
	2,  // 1: services.CommissionDetailResponse.pager:type_name -> common.Pager
	0,  // 2: services.CommissionDetailResponse.items:type_name -> services.CommissionDetail
	3,  // 3: services.CommissionDetailResponse.error:type_name -> common.Error
	4,  // 4: services.CommissionDetailResponse.info:type_name -> common.Info
	0,  // 5: services.CommissionDetailService.Create:input_type -> services.CommissionDetail
	0,  // 6: services.CommissionDetailService.Update:input_type -> services.CommissionDetail
	0,  // 7: services.CommissionDetailService.Delete:input_type -> services.CommissionDetail
	0,  // 8: services.CommissionDetailService.Get:input_type -> services.CommissionDetail
	5,  // 9: services.CommissionDetailService.Search:input_type -> common.BaseWhere
	1,  // 10: services.CommissionDetailService.Create:output_type -> services.CommissionDetailResponse
	1,  // 11: services.CommissionDetailService.Update:output_type -> services.CommissionDetailResponse
	1,  // 12: services.CommissionDetailService.Delete:output_type -> services.CommissionDetailResponse
	1,  // 13: services.CommissionDetailService.Get:output_type -> services.CommissionDetailResponse
	1,  // 14: services.CommissionDetailService.Search:output_type -> services.CommissionDetailResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_commissionDetailService_proto_init() }
func file_commissionDetailService_proto_init() {
	if File_commissionDetailService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commissionDetailService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommissionDetail); i {
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
		file_commissionDetailService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommissionDetailResponse); i {
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
			RawDescriptor: file_commissionDetailService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commissionDetailService_proto_goTypes,
		DependencyIndexes: file_commissionDetailService_proto_depIdxs,
		MessageInfos:      file_commissionDetailService_proto_msgTypes,
	}.Build()
	File_commissionDetailService_proto = out.File
	file_commissionDetailService_proto_rawDesc = nil
	file_commissionDetailService_proto_goTypes = nil
	file_commissionDetailService_proto_depIdxs = nil
}
