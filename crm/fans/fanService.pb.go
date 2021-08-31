// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: fanService.proto

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

type Fan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	FanSn         string `protobuf:"bytes,2,opt,name=fan_sn,json=fanSn,proto3" json:"fan_sn"`
	CustomerId    int64  `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	OpenId        string `protobuf:"bytes,4,opt,name=open_id,json=openId,proto3" json:"open_id"`
	UnionId       string `protobuf:"bytes,5,opt,name=union_id,json=unionId,proto3" json:"union_id"`
	Source        int32  `protobuf:"varint,6,opt,name=source,proto3" json:"source"`
	NickName      string `protobuf:"bytes,7,opt,name=nick_name,json=nickName,proto3" json:"nick_name"`
	AvatarUrl     string `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url"`
	Gender        int32  `protobuf:"varint,9,opt,name=gender,proto3" json:"gender"`
	Province      string `protobuf:"bytes,10,opt,name=province,proto3" json:"province"`
	City          string `protobuf:"bytes,11,opt,name=city,proto3" json:"city"`
	Country       string `protobuf:"bytes,12,opt,name=country,proto3" json:"country"`
	Mobile        string `protobuf:"bytes,13,opt,name=mobile,proto3" json:"mobile"`
	Email         string `protobuf:"bytes,14,opt,name=email,proto3" json:"email"`
	Remark        string `protobuf:"bytes,15,opt,name=remark,proto3" json:"remark"`
	RecommenderId int64  `protobuf:"varint,16,opt,name=recommender_id,json=recommenderId,proto3" json:"recommender_id"`
	Status        int32  `protobuf:"varint,17,opt,name=status,proto3" json:"status"`
	RegisteredAt  string `protobuf:"bytes,18,opt,name=registered_at,json=registeredAt,proto3" json:"registered_at"`
	CreatedAt     string `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Fan) Reset() {
	*x = Fan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fanService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fan) ProtoMessage() {}

func (x *Fan) ProtoReflect() protoreflect.Message {
	mi := &file_fanService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fan.ProtoReflect.Descriptor instead.
func (*Fan) Descriptor() ([]byte, []int) {
	return file_fanService_proto_rawDescGZIP(), []int{0}
}

func (x *Fan) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Fan) GetFanSn() string {
	if x != nil {
		return x.FanSn
	}
	return ""
}

func (x *Fan) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Fan) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *Fan) GetUnionId() string {
	if x != nil {
		return x.UnionId
	}
	return ""
}

func (x *Fan) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *Fan) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *Fan) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Fan) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Fan) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Fan) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Fan) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Fan) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Fan) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Fan) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Fan) GetRecommenderId() int64 {
	if x != nil {
		return x.RecommenderId
	}
	return 0
}

func (x *Fan) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Fan) GetRegisteredAt() string {
	if x != nil {
		return x.RegisteredAt
	}
	return ""
}

func (x *Fan) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Fan) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//粉丝查询参数
type FanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//以下为自定义参数
	Id            int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	NickName      string  `protobuf:"bytes,5,opt,name=nick_name,json=nickName,proto3" json:"nick_name"`
	Gender        int32   `protobuf:"varint,6,opt,name=gender,proto3" json:"gender"`
	Mobile        string  `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile"`
	Source        int32   `protobuf:"varint,8,opt,name=source,proto3" json:"source"`
	Status        int32   `protobuf:"varint,9,opt,name=status,proto3" json:"status"`
	OpenId        string  `protobuf:"bytes,10,opt,name=open_id,json=openId,proto3" json:"open_id"`
	CustomerId    int64   `protobuf:"varint,11,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	RecommenderId int64   `protobuf:"varint,12,opt,name=recommender_id,json=recommenderId,proto3" json:"recommender_id"`
	Ids           []int64 `protobuf:"varint,13,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *FanRequest) Reset() {
	*x = FanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fanService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanRequest) ProtoMessage() {}

func (x *FanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fanService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanRequest.ProtoReflect.Descriptor instead.
func (*FanRequest) Descriptor() ([]byte, []int) {
	return file_fanService_proto_rawDescGZIP(), []int{1}
}

func (x *FanRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *FanRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FanRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *FanRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FanRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *FanRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *FanRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *FanRequest) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *FanRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FanRequest) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *FanRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *FanRequest) GetRecommenderId() int64 {
	if x != nil {
		return x.RecommenderId
	}
	return 0
}

func (x *FanRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//
type FanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Fan          `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Fan        `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *FanResponse) Reset() {
	*x = FanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fanService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanResponse) ProtoMessage() {}

func (x *FanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fanService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanResponse.ProtoReflect.Descriptor instead.
func (*FanResponse) Descriptor() ([]byte, []int) {
	return file_fanService_proto_rawDescGZIP(), []int{2}
}

func (x *FanResponse) GetEntity() *Fan {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *FanResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *FanResponse) GetItems() []*Fan {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FanResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *FanResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_fanService_proto protoreflect.FileDescriptor

var file_fanService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9f, 0x04, 0x0a, 0x03, 0x46, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x66, 0x61, 0x6e, 0x5f, 0x73,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x61, 0x6e, 0x53, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xd9, 0x02, 0x0a, 0x0a, 0x46, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65,
	0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xc5, 0x01,
	0x0a, 0x0b, 0x46, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x45, 0x0a, 0x0c, 0x4d, 0x79, 0x46, 0x61, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xd8, 0x01, 0x0a,
	0x0a, 0x46, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x46, 0x61, 0x6e, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x46, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x46, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x46, 0x61, 0x6e, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x46, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fanService_proto_rawDescOnce sync.Once
	file_fanService_proto_rawDescData = file_fanService_proto_rawDesc
)

func file_fanService_proto_rawDescGZIP() []byte {
	file_fanService_proto_rawDescOnce.Do(func() {
		file_fanService_proto_rawDescData = protoimpl.X.CompressGZIP(file_fanService_proto_rawDescData)
	})
	return file_fanService_proto_rawDescData
}

var file_fanService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_fanService_proto_goTypes = []interface{}{
	(*Fan)(nil),          // 0: services.Fan
	(*FanRequest)(nil),   // 1: services.FanRequest
	(*FanResponse)(nil),  // 2: services.FanResponse
	(*common.Pager)(nil), // 3: common.Pager
	(*common.Error)(nil), // 4: common.Error
	(*common.Info)(nil),  // 5: common.Info
}
var file_fanService_proto_depIdxs = []int32{
	0,  // 0: services.FanResponse.entity:type_name -> services.Fan
	3,  // 1: services.FanResponse.pager:type_name -> common.Pager
	0,  // 2: services.FanResponse.items:type_name -> services.Fan
	4,  // 3: services.FanResponse.error:type_name -> common.Error
	5,  // 4: services.FanResponse.info:type_name -> common.Info
	1,  // 5: services.MyFanService.Info:input_type -> services.FanRequest
	0,  // 6: services.FanService.Update:input_type -> services.Fan
	0,  // 7: services.FanService.Delete:input_type -> services.Fan
	0,  // 8: services.FanService.Get:input_type -> services.Fan
	1,  // 9: services.FanService.Search:input_type -> services.FanRequest
	2,  // 10: services.MyFanService.Info:output_type -> services.FanResponse
	2,  // 11: services.FanService.Update:output_type -> services.FanResponse
	2,  // 12: services.FanService.Delete:output_type -> services.FanResponse
	2,  // 13: services.FanService.Get:output_type -> services.FanResponse
	2,  // 14: services.FanService.Search:output_type -> services.FanResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_fanService_proto_init() }
func file_fanService_proto_init() {
	if File_fanService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fanService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fan); i {
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
		file_fanService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanRequest); i {
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
		file_fanService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanResponse); i {
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
			RawDescriptor: file_fanService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_fanService_proto_goTypes,
		DependencyIndexes: file_fanService_proto_depIdxs,
		MessageInfos:      file_fanService_proto_msgTypes,
	}.Build()
	File_fanService_proto = out.File
	file_fanService_proto_rawDesc = nil
	file_fanService_proto_goTypes = nil
	file_fanService_proto_depIdxs = nil
}
