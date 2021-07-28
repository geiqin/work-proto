// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: generalService.proto

package services

import (
	_ "github.com/geiqin/micro-kit/protobuf/common"
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

type GeneralRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize   int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting    string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Keywords   string  `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Name       string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	ItemSn     string  `protobuf:"bytes,6,opt,name=item_sn,json=itemSn,proto3" json:"item_sn,omitempty"`
	BrandId    int32   `protobuf:"varint,7,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	TagId      int32   `protobuf:"varint,8,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	CatId      int32   `protobuf:"varint,9,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"`
	TaxonomyId int64   `protobuf:"varint,10,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id,omitempty"`
	MinPrice   float32 `protobuf:"fixed32,11,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice   float32 `protobuf:"fixed32,12,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	Type       string  `protobuf:"bytes,13,opt,name=type,proto3" json:"type,omitempty"`
	Status     string  `protobuf:"bytes,14,opt,name=status,proto3" json:"status,omitempty"`
	CouponId   int64   `protobuf:"varint,15,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	Valid      bool    `protobuf:"varint,16,opt,name=valid,proto3" json:"valid,omitempty"`  // 是否有效商品
	Invite     string  `protobuf:"bytes,17,opt,name=invite,proto3" json:"invite,omitempty"` // 加密后的邀请码
	Id         int64   `protobuf:"varint,18,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: gorm:"-"
	GoodsSkuId int64  `protobuf:"varint,19,opt,name=goods_sku_id,json=goodsSkuId,proto3" json:"goods_sku_id,omitempty"` //(特殊专用)
	SkuId      int64  `protobuf:"varint,20,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`                  //单一规格商品输入参数（在获取规格商品时使用）
	CatSlug    string `protobuf:"bytes,21,opt,name=cat_slug,json=catSlug,proto3" json:"cat_slug,omitempty"`
	IsSku      bool   `protobuf:"varint,22,opt,name=is_sku,json=isSku,proto3" json:"is_sku,omitempty"`
	Sort       int32  `protobuf:"varint,23,opt,name=sort,proto3" json:"sort,omitempty"`
	Num        int32  `protobuf:"varint,24,opt,name=num,proto3" json:"num,omitempty"`
	// @inject_tag: gorm:"-"
	Ids []int64 `protobuf:"varint,25,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GeneralRequest) Reset() {
	*x = GeneralRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generalService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralRequest) ProtoMessage() {}

func (x *GeneralRequest) ProtoReflect() protoreflect.Message {
	mi := &file_generalService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralRequest.ProtoReflect.Descriptor instead.
func (*GeneralRequest) Descriptor() ([]byte, []int) {
	return file_generalService_proto_rawDescGZIP(), []int{0}
}

func (x *GeneralRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *GeneralRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GeneralRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *GeneralRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *GeneralRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GeneralRequest) GetItemSn() string {
	if x != nil {
		return x.ItemSn
	}
	return ""
}

func (x *GeneralRequest) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *GeneralRequest) GetTagId() int32 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *GeneralRequest) GetCatId() int32 {
	if x != nil {
		return x.CatId
	}
	return 0
}

func (x *GeneralRequest) GetTaxonomyId() int64 {
	if x != nil {
		return x.TaxonomyId
	}
	return 0
}

func (x *GeneralRequest) GetMinPrice() float32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *GeneralRequest) GetMaxPrice() float32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *GeneralRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GeneralRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GeneralRequest) GetCouponId() int64 {
	if x != nil {
		return x.CouponId
	}
	return 0
}

func (x *GeneralRequest) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *GeneralRequest) GetInvite() string {
	if x != nil {
		return x.Invite
	}
	return ""
}

func (x *GeneralRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GeneralRequest) GetGoodsSkuId() int64 {
	if x != nil {
		return x.GoodsSkuId
	}
	return 0
}

func (x *GeneralRequest) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *GeneralRequest) GetCatSlug() string {
	if x != nil {
		return x.CatSlug
	}
	return ""
}

func (x *GeneralRequest) GetIsSku() bool {
	if x != nil {
		return x.IsSku
	}
	return false
}

func (x *GeneralRequest) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *GeneralRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *GeneralRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_generalService_proto protoreflect.FileDescriptor

var file_generalService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x04, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x78, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74,
	0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x69,
	0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b,
	0x75, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x74, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x15,
	0x0a, 0x06, 0x69, 0x73, 0x5f, 0x73, 0x6b, 0x75, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x69, 0x73, 0x53, 0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x32, 0xb0, 0x04,
	0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x09, 0x53, 0x6b, 0x75, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x07, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x0b, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_generalService_proto_rawDescOnce sync.Once
	file_generalService_proto_rawDescData = file_generalService_proto_rawDesc
)

func file_generalService_proto_rawDescGZIP() []byte {
	file_generalService_proto_rawDescOnce.Do(func() {
		file_generalService_proto_rawDescData = protoimpl.X.CompressGZIP(file_generalService_proto_rawDescData)
	})
	return file_generalService_proto_rawDescData
}

var file_generalService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_generalService_proto_goTypes = []interface{}{
	(*GeneralRequest)(nil), // 0: services.GeneralRequest
	(*Item)(nil),           // 1: services.Item
	(*ItemResponse)(nil),   // 2: services.ItemResponse
}
var file_generalService_proto_depIdxs = []int32{
	1, // 0: services.GeneralService.Create:input_type -> services.Item
	1, // 1: services.GeneralService.Update:input_type -> services.Item
	0, // 2: services.GeneralService.Delete:input_type -> services.GeneralRequest
	0, // 3: services.GeneralService.Get:input_type -> services.GeneralRequest
	0, // 4: services.GeneralService.Detail:input_type -> services.GeneralRequest
	0, // 5: services.GeneralService.SkuDetail:input_type -> services.GeneralRequest
	0, // 6: services.GeneralService.Search:input_type -> services.GeneralRequest
	0, // 7: services.GeneralService.Display:input_type -> services.GeneralRequest
	0, // 8: services.GeneralService.FrontSearch:input_type -> services.GeneralRequest
	2, // 9: services.GeneralService.Create:output_type -> services.ItemResponse
	2, // 10: services.GeneralService.Update:output_type -> services.ItemResponse
	2, // 11: services.GeneralService.Delete:output_type -> services.ItemResponse
	2, // 12: services.GeneralService.Get:output_type -> services.ItemResponse
	2, // 13: services.GeneralService.Detail:output_type -> services.ItemResponse
	2, // 14: services.GeneralService.SkuDetail:output_type -> services.ItemResponse
	2, // 15: services.GeneralService.Search:output_type -> services.ItemResponse
	2, // 16: services.GeneralService.Display:output_type -> services.ItemResponse
	2, // 17: services.GeneralService.FrontSearch:output_type -> services.ItemResponse
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_generalService_proto_init() }
func file_generalService_proto_init() {
	if File_generalService_proto != nil {
		return
	}
	file_itemService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_generalService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralRequest); i {
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
			RawDescriptor: file_generalService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_generalService_proto_goTypes,
		DependencyIndexes: file_generalService_proto_depIdxs,
		MessageInfos:      file_generalService_proto_msgTypes,
	}.Build()
	File_generalService_proto = out.File
	file_generalService_proto_rawDesc = nil
	file_generalService_proto_goTypes = nil
	file_generalService_proto_depIdxs = nil
}