// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: purchaseService.proto

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

//充值下单请求参数
type RechargeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string  `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel"` //充值通道: vip_capital余额充值, phone_fee话费充值, deposit交保证金
	Money   float32 `protobuf:"fixed32,2,opt,name=money,proto3" json:"money"`   //充值金额
	Content string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content"` //内容：比如话费充值，这里就是手机号
}

func (x *RechargeRequest) Reset() {
	*x = RechargeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchaseService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RechargeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RechargeRequest) ProtoMessage() {}

func (x *RechargeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_purchaseService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RechargeRequest.ProtoReflect.Descriptor instead.
func (*RechargeRequest) Descriptor() ([]byte, []int) {
	return file_purchaseService_proto_rawDescGZIP(), []int{0}
}

func (x *RechargeRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *RechargeRequest) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *RechargeRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

//下单请求参数
type PurchaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`                                                                                            //订单类型（默认为普通订单）
	AddressId  int64             `protobuf:"varint,2,opt,name=address_id,json=addressId,proto3" json:"address_id"`                                                                //收货地址ID
	Message    string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message"`                                                                                      //买家留言(50字以内)
	PayMethod  int32             `protobuf:"varint,4,opt,name=pay_method,json=payMethod,proto3" json:"pay_method"`                                                                //选中的支付方式
	Source     string            `protobuf:"bytes,5,opt,name=source,proto3" json:"source"`                                                                                        //下单来源
	TicketId   int64             `protobuf:"varint,6,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id"`                                                                   //选中的优惠劵凭证ID
	Points     int32             `protobuf:"varint,7,opt,name=points,proto3" json:"points"`                                                                                       //使用的积分数
	VipCardId  int64             `protobuf:"varint,8,opt,name=vip_card_id,json=vipCardId,proto3" json:"vip_card_id"`                                                              //选中的会员卡
	CustomerId int64             `protobuf:"varint,9,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`                                                             //客户ID（代客下单）
	GoodsList  []*PurchaseItem   `protobuf:"bytes,10,rep,name=goods_list,json=goodsList,proto3" json:"goods_list"`                                                                //选中的商品清单
	Metas      map[string]string `protobuf:"bytes,11,rep,name=metas,proto3" json:"metas" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //订单扩展信息
}

func (x *PurchaseRequest) Reset() {
	*x = PurchaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchaseService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseRequest) ProtoMessage() {}

func (x *PurchaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_purchaseService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseRequest.ProtoReflect.Descriptor instead.
func (*PurchaseRequest) Descriptor() ([]byte, []int) {
	return file_purchaseService_proto_rawDescGZIP(), []int{1}
}

func (x *PurchaseRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PurchaseRequest) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *PurchaseRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PurchaseRequest) GetPayMethod() int32 {
	if x != nil {
		return x.PayMethod
	}
	return 0
}

func (x *PurchaseRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *PurchaseRequest) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *PurchaseRequest) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *PurchaseRequest) GetVipCardId() int64 {
	if x != nil {
		return x.VipCardId
	}
	return 0
}

func (x *PurchaseRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *PurchaseRequest) GetGoodsList() []*PurchaseItem {
	if x != nil {
		return x.GoodsList
	}
	return nil
}

func (x *PurchaseRequest) GetMetas() map[string]string {
	if x != nil {
		return x.Metas
	}
	return nil
}

//下单商品项
type PurchaseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId    int64  `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	SkuId     int64  `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Num       int32  `protobuf:"varint,3,opt,name=num,proto3" json:"num"`
	CartRowId string `protobuf:"bytes,4,opt,name=cart_row_id,json=cartRowId,proto3" json:"cart_row_id"`
}

func (x *PurchaseItem) Reset() {
	*x = PurchaseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchaseService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseItem) ProtoMessage() {}

func (x *PurchaseItem) ProtoReflect() protoreflect.Message {
	mi := &file_purchaseService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseItem.ProtoReflect.Descriptor instead.
func (*PurchaseItem) Descriptor() ([]byte, []int) {
	return file_purchaseService_proto_rawDescGZIP(), []int{2}
}

func (x *PurchaseItem) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *PurchaseItem) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *PurchaseItem) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *PurchaseItem) GetCartRowId() string {
	if x != nil {
		return x.CartRowId
	}
	return ""
}

//付款清单（开具账单）
type Billing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       int32       `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	Total       float32     `protobuf:"fixed32,2,opt,name=total,proto3" json:"total"`
	TotalWeight float32     `protobuf:"fixed32,3,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight"`
	Discount    float32     `protobuf:"fixed32,4,opt,name=discount,proto3" json:"discount"`
	Freight     float32     `protobuf:"fixed32,5,opt,name=freight,proto3" json:"freight"`
	Amount      float32     `protobuf:"fixed32,6,opt,name=amount,proto3" json:"amount"`
	AddressId   int64       `protobuf:"varint,7,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	CustomerId  int64       `protobuf:"varint,8,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	Message     string      `protobuf:"bytes,9,opt,name=message,proto3" json:"message"`                               //买家留言(50字以内)
	Source      string      `protobuf:"bytes,10,opt,name=source,proto3" json:"source"`                                //下单来源
	TicketId    int64       `protobuf:"varint,11,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id"`           //正在使用的优惠劵ID
	TicketIds   []int64     `protobuf:"varint,12,rep,packed,name=ticket_ids,json=ticketIds,proto3" json:"ticket_ids"` //可使用的优惠劵ID集合
	Items       []*BillItem `protobuf:"bytes,13,rep,name=items,proto3" json:"items"`                                  //选中的商品清单
}

func (x *Billing) Reset() {
	*x = Billing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchaseService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Billing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Billing) ProtoMessage() {}

func (x *Billing) ProtoReflect() protoreflect.Message {
	mi := &file_purchaseService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Billing.ProtoReflect.Descriptor instead.
func (*Billing) Descriptor() ([]byte, []int) {
	return file_purchaseService_proto_rawDescGZIP(), []int{3}
}

func (x *Billing) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Billing) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Billing) GetTotalWeight() float32 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

func (x *Billing) GetDiscount() float32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Billing) GetFreight() float32 {
	if x != nil {
		return x.Freight
	}
	return 0
}

func (x *Billing) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Billing) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *Billing) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Billing) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Billing) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Billing) GetTicketId() int64 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *Billing) GetTicketIds() []int64 {
	if x != nil {
		return x.TicketIds
	}
	return nil
}

func (x *Billing) GetItems() []*BillItem {
	if x != nil {
		return x.Items
	}
	return nil
}

//账单条目
type BillItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId             int64      `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	SkuId              int64      `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Num                int32      `protobuf:"varint,3,opt,name=num,proto3" json:"num"`
	Price              float32    `protobuf:"fixed32,4,opt,name=price,proto3" json:"price"`
	OriginPrice        float32    `protobuf:"fixed32,5,opt,name=origin_price,json=originPrice,proto3" json:"origin_price"`
	SubTotal           float32    `protobuf:"fixed32,6,opt,name=sub_total,json=subTotal,proto3" json:"sub_total"`
	IsGift             bool       `protobuf:"varint,7,opt,name=is_gift,json=isGift,proto3" json:"is_gift"`
	PromotionId        int64      `protobuf:"varint,8,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id"`
	MemberPrice        float32    `protobuf:"fixed32,9,opt,name=member_price,json=memberPrice,proto3" json:"member_price"`                         // 会员价
	LimitDiscountPrice float32    `protobuf:"fixed32,10,opt,name=limit_discount_price,json=limitDiscountPrice,proto3" json:"limit_discount_price"` // 限时折扣价
	LimitDiscountNum   int32      `protobuf:"varint,11,opt,name=limit_discount_num,json=limitDiscountNum,proto3" json:"limit_discount_num"`        // 享受限时折扣的商品数量
	ExchangeNum        int32      `protobuf:"varint,12,opt,name=exchange_num,json=exchangeNum,proto3" json:"exchange_num"`                         // 兑换券兑换商品的数量
	Goods              *GoodsInfo `protobuf:"bytes,13,opt,name=goods,proto3" json:"goods"`
	CartRowId          string     `protobuf:"bytes,14,opt,name=cart_row_id,json=cartRowId,proto3" json:"cart_row_id"`
}

func (x *BillItem) Reset() {
	*x = BillItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchaseService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillItem) ProtoMessage() {}

func (x *BillItem) ProtoReflect() protoreflect.Message {
	mi := &file_purchaseService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillItem.ProtoReflect.Descriptor instead.
func (*BillItem) Descriptor() ([]byte, []int) {
	return file_purchaseService_proto_rawDescGZIP(), []int{4}
}

func (x *BillItem) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *BillItem) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *BillItem) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *BillItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BillItem) GetOriginPrice() float32 {
	if x != nil {
		return x.OriginPrice
	}
	return 0
}

func (x *BillItem) GetSubTotal() float32 {
	if x != nil {
		return x.SubTotal
	}
	return 0
}

func (x *BillItem) GetIsGift() bool {
	if x != nil {
		return x.IsGift
	}
	return false
}

func (x *BillItem) GetPromotionId() int64 {
	if x != nil {
		return x.PromotionId
	}
	return 0
}

func (x *BillItem) GetMemberPrice() float32 {
	if x != nil {
		return x.MemberPrice
	}
	return 0
}

func (x *BillItem) GetLimitDiscountPrice() float32 {
	if x != nil {
		return x.LimitDiscountPrice
	}
	return 0
}

func (x *BillItem) GetLimitDiscountNum() int32 {
	if x != nil {
		return x.LimitDiscountNum
	}
	return 0
}

func (x *BillItem) GetExchangeNum() int32 {
	if x != nil {
		return x.ExchangeNum
	}
	return 0
}

func (x *BillItem) GetGoods() *GoodsInfo {
	if x != nil {
		return x.Goods
	}
	return nil
}

func (x *BillItem) GetCartRowId() string {
	if x != nil {
		return x.CartRowId
	}
	return ""
}

type Shopping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *PurchaseRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request"`
	Billing *Billing         `protobuf:"bytes,2,opt,name=billing,proto3" json:"billing"`
}

func (x *Shopping) Reset() {
	*x = Shopping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchaseService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shopping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shopping) ProtoMessage() {}

func (x *Shopping) ProtoReflect() protoreflect.Message {
	mi := &file_purchaseService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shopping.ProtoReflect.Descriptor instead.
func (*Shopping) Descriptor() ([]byte, []int) {
	return file_purchaseService_proto_rawDescGZIP(), []int{5}
}

func (x *Shopping) GetRequest() *PurchaseRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Shopping) GetBilling() *Billing {
	if x != nil {
		return x.Billing
	}
	return nil
}

type PurchaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Shopping `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	//map<string,string> params   = 2;
	//repeated PayMethod methods = 3;
	Error *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info  *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *PurchaseResponse) Reset() {
	*x = PurchaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchaseService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseResponse) ProtoMessage() {}

func (x *PurchaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_purchaseService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseResponse.ProtoReflect.Descriptor instead.
func (*PurchaseResponse) Descriptor() ([]byte, []int) {
	return file_purchaseService_proto_rawDescGZIP(), []int{6}
}

func (x *PurchaseResponse) GetEntity() *Shopping {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PurchaseResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PurchaseResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_purchaseService_proto protoreflect.FileDescriptor

var file_purchaseService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5b, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xb8, 0x03,
	0x0a, 0x0f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x76, 0x69,
	0x70, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x76, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0a, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x1a, 0x38,
	0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x70, 0x0a, 0x0c, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0b, 0x63, 0x61,
	0x72, 0x74, 0x5f, 0x72, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x77, 0x49, 0x64, 0x22, 0xfe, 0x02, 0x0a, 0x07, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x0c, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x73, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xcf, 0x03, 0x0a, 0x08,
	0x42, 0x69, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x67, 0x69, 0x66, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x47, 0x69, 0x66, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x14, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x29, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1e, 0x0a,
	0x0b, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x6c, 0x0a,
	0x08, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x07, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x52, 0x07, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x85, 0x01, 0x0a, 0x10,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x32, 0xd5, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x08, 0x52, 0x65, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_purchaseService_proto_rawDescOnce sync.Once
	file_purchaseService_proto_rawDescData = file_purchaseService_proto_rawDesc
)

func file_purchaseService_proto_rawDescGZIP() []byte {
	file_purchaseService_proto_rawDescOnce.Do(func() {
		file_purchaseService_proto_rawDescData = protoimpl.X.CompressGZIP(file_purchaseService_proto_rawDescData)
	})
	return file_purchaseService_proto_rawDescData
}

var file_purchaseService_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_purchaseService_proto_goTypes = []interface{}{
	(*RechargeRequest)(nil),  // 0: services.RechargeRequest
	(*PurchaseRequest)(nil),  // 1: services.PurchaseRequest
	(*PurchaseItem)(nil),     // 2: services.PurchaseItem
	(*Billing)(nil),          // 3: services.Billing
	(*BillItem)(nil),         // 4: services.BillItem
	(*Shopping)(nil),         // 5: services.Shopping
	(*PurchaseResponse)(nil), // 6: services.PurchaseResponse
	nil,                      // 7: services.PurchaseRequest.MetasEntry
	(*GoodsInfo)(nil),        // 8: services.GoodsInfo
	(*common.Error)(nil),     // 9: common.Error
	(*common.Info)(nil),      // 10: common.Info
	(*OrderResponse)(nil),    // 11: services.OrderResponse
}
var file_purchaseService_proto_depIdxs = []int32{
	2,  // 0: services.PurchaseRequest.goods_list:type_name -> services.PurchaseItem
	7,  // 1: services.PurchaseRequest.metas:type_name -> services.PurchaseRequest.MetasEntry
	4,  // 2: services.Billing.items:type_name -> services.BillItem
	8,  // 3: services.BillItem.goods:type_name -> services.GoodsInfo
	1,  // 4: services.Shopping.request:type_name -> services.PurchaseRequest
	3,  // 5: services.Shopping.billing:type_name -> services.Billing
	5,  // 6: services.PurchaseResponse.entity:type_name -> services.Shopping
	9,  // 7: services.PurchaseResponse.error:type_name -> common.Error
	10, // 8: services.PurchaseResponse.info:type_name -> common.Info
	1,  // 9: services.PurchaseService.Write:input_type -> services.PurchaseRequest
	1,  // 10: services.PurchaseService.Submit:input_type -> services.PurchaseRequest
	0,  // 11: services.PurchaseService.Recharge:input_type -> services.RechargeRequest
	6,  // 12: services.PurchaseService.Write:output_type -> services.PurchaseResponse
	11, // 13: services.PurchaseService.Submit:output_type -> services.OrderResponse
	11, // 14: services.PurchaseService.Recharge:output_type -> services.OrderResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_purchaseService_proto_init() }
func file_purchaseService_proto_init() {
	if File_purchaseService_proto != nil {
		return
	}
	file_goodsInfoService_proto_init()
	file_orderService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_purchaseService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RechargeRequest); i {
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
		file_purchaseService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseRequest); i {
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
		file_purchaseService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseItem); i {
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
		file_purchaseService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Billing); i {
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
		file_purchaseService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillItem); i {
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
		file_purchaseService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shopping); i {
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
		file_purchaseService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseResponse); i {
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
			RawDescriptor: file_purchaseService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_purchaseService_proto_goTypes,
		DependencyIndexes: file_purchaseService_proto_depIdxs,
		MessageInfos:      file_purchaseService_proto_msgTypes,
	}.Build()
	File_purchaseService_proto = out.File
	file_purchaseService_proto_rawDesc = nil
	file_purchaseService_proto_goTypes = nil
	file_purchaseService_proto_depIdxs = nil
}
