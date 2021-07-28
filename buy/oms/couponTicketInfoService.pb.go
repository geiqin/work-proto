// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: couponTicketInfoService.proto

package services

import (
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

// 优惠劵凭证
type CouponTicket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	TicketSn string `protobuf:"bytes,2,opt,name=ticket_sn,json=ticketSn,proto3" json:"ticket_sn"`
	StartAt  string `protobuf:"bytes,3,opt,name=start_at,json=startAt,proto3" json:"start_at"`
	EndAt    string `protobuf:"bytes,4,opt,name=end_at,json=endAt,proto3" json:"end_at"`
	CouponId int64  `protobuf:"varint,5,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id"`
	// @inject_tag: gorm:"-"
	Coupon *CouponInfo `protobuf:"bytes,6,opt,name=coupon,proto3" json:"coupon" gorm:"-"`
}

func (x *CouponTicket) Reset() {
	*x = CouponTicket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_couponTicketInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponTicket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponTicket) ProtoMessage() {}

func (x *CouponTicket) ProtoReflect() protoreflect.Message {
	mi := &file_couponTicketInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponTicket.ProtoReflect.Descriptor instead.
func (*CouponTicket) Descriptor() ([]byte, []int) {
	return file_couponTicketInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *CouponTicket) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CouponTicket) GetTicketSn() string {
	if x != nil {
		return x.TicketSn
	}
	return ""
}

func (x *CouponTicket) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *CouponTicket) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

func (x *CouponTicket) GetCouponId() int64 {
	if x != nil {
		return x.CouponId
	}
	return 0
}

func (x *CouponTicket) GetCoupon() *CouponInfo {
	if x != nil {
		return x.Coupon
	}
	return nil
}

// 优惠券
type CouponInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CouponSn             string  `protobuf:"bytes,2,opt,name=coupon_sn,json=couponSn,proto3" json:"coupon_sn"`
	Title                string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	IsAtLeast            bool    `protobuf:"varint,4,opt,name=is_at_least,json=isAtLeast,proto3" json:"is_at_least"`
	AtLeast              float32 `protobuf:"fixed32,5,opt,name=at_least,json=atLeast,proto3" json:"at_least"`
	PreferentialType     int32   `protobuf:"varint,6,opt,name=preferential_type,json=preferentialType,proto3" json:"preferential_type"`
	PreferentialMoney    float32 `protobuf:"fixed32,7,opt,name=preferential_money,json=preferentialMoney,proto3" json:"preferential_money"`
	PreferentialDiscount float32 `protobuf:"fixed32,8,opt,name=preferential_discount,json=preferentialDiscount,proto3" json:"preferential_discount"`
	ExchangeNum          int32   `protobuf:"varint,9,opt,name=exchange_num,json=exchangeNum,proto3" json:"exchange_num"`
	RangeType            string  `protobuf:"bytes,10,opt,name=range_type,json=rangeType,proto3" json:"range_type"`
	Description          string  `protobuf:"bytes,11,opt,name=description,proto3" json:"description"`
}

func (x *CouponInfo) Reset() {
	*x = CouponInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_couponTicketInfoService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponInfo) ProtoMessage() {}

func (x *CouponInfo) ProtoReflect() protoreflect.Message {
	mi := &file_couponTicketInfoService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponInfo.ProtoReflect.Descriptor instead.
func (*CouponInfo) Descriptor() ([]byte, []int) {
	return file_couponTicketInfoService_proto_rawDescGZIP(), []int{1}
}

func (x *CouponInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CouponInfo) GetCouponSn() string {
	if x != nil {
		return x.CouponSn
	}
	return ""
}

func (x *CouponInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CouponInfo) GetIsAtLeast() bool {
	if x != nil {
		return x.IsAtLeast
	}
	return false
}

func (x *CouponInfo) GetAtLeast() float32 {
	if x != nil {
		return x.AtLeast
	}
	return 0
}

func (x *CouponInfo) GetPreferentialType() int32 {
	if x != nil {
		return x.PreferentialType
	}
	return 0
}

func (x *CouponInfo) GetPreferentialMoney() float32 {
	if x != nil {
		return x.PreferentialMoney
	}
	return 0
}

func (x *CouponInfo) GetPreferentialDiscount() float32 {
	if x != nil {
		return x.PreferentialDiscount
	}
	return 0
}

func (x *CouponInfo) GetExchangeNum() int32 {
	if x != nil {
		return x.ExchangeNum
	}
	return 0
}

func (x *CouponInfo) GetRangeType() string {
	if x != nil {
		return x.RangeType
	}
	return ""
}

func (x *CouponInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_couponTicketInfoService_proto protoreflect.FileDescriptor

var file_couponTicketInfoService_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xb8, 0x01, 0x0a, 0x0c, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x22, 0xff, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x73, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x61, 0x74, 0x5f,
	0x6c, 0x65, 0x61, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x41,
	0x74, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x5f, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x61, 0x74, 0x4c, 0x65, 0x61, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d,
	0x0a, 0x12, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x33, 0x0a,
	0x15, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x14, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_couponTicketInfoService_proto_rawDescOnce sync.Once
	file_couponTicketInfoService_proto_rawDescData = file_couponTicketInfoService_proto_rawDesc
)

func file_couponTicketInfoService_proto_rawDescGZIP() []byte {
	file_couponTicketInfoService_proto_rawDescOnce.Do(func() {
		file_couponTicketInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_couponTicketInfoService_proto_rawDescData)
	})
	return file_couponTicketInfoService_proto_rawDescData
}

var file_couponTicketInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_couponTicketInfoService_proto_goTypes = []interface{}{
	(*CouponTicket)(nil), // 0: services.CouponTicket
	(*CouponInfo)(nil),   // 1: services.CouponInfo
}
var file_couponTicketInfoService_proto_depIdxs = []int32{
	1, // 0: services.CouponTicket.coupon:type_name -> services.CouponInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_couponTicketInfoService_proto_init() }
func file_couponTicketInfoService_proto_init() {
	if File_couponTicketInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_couponTicketInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponTicket); i {
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
		file_couponTicketInfoService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponInfo); i {
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
			RawDescriptor: file_couponTicketInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_couponTicketInfoService_proto_goTypes,
		DependencyIndexes: file_couponTicketInfoService_proto_depIdxs,
		MessageInfos:      file_couponTicketInfoService_proto_msgTypes,
	}.Build()
	File_couponTicketInfoService_proto = out.File
	file_couponTicketInfoService_proto_rawDesc = nil
	file_couponTicketInfoService_proto_goTypes = nil
	file_couponTicketInfoService_proto_depIdxs = nil
}