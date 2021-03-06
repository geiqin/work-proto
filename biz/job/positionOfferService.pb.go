// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: positionOfferService.proto

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

type PositionOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	PositionId int64     `protobuf:"varint,2,opt,name=position_id,json=positionId,proto3" json:"position_id"`
	CustomerId int64     `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	SeekerId   int64     `protobuf:"varint,4,opt,name=seeker_id,json=seekerId,proto3" json:"seeker_id"`
	RealName   string    `protobuf:"bytes,5,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard     string    `protobuf:"bytes,6,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	Gender     int32     `protobuf:"varint,7,opt,name=gender,proto3" json:"gender"`
	Nation     int32     `protobuf:"varint,8,opt,name=nation,proto3" json:"nation"`
	Birthday   string    `protobuf:"bytes,9,opt,name=birthday,proto3" json:"birthday"`
	Mobile     string    `protobuf:"bytes,10,opt,name=mobile,proto3" json:"mobile"`
	AreaId     int64     `protobuf:"varint,11,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	IsSelf     bool      `protobuf:"varint,12,opt,name=is_self,json=isSelf,proto3" json:"is_self"`
	Reason     string    `protobuf:"bytes,13,opt,name=reason,proto3" json:"reason"`
	StartType  int32     `protobuf:"varint,14,opt,name=start_type,json=startType,proto3" json:"start_type"`
	StartDate  string    `protobuf:"bytes,15,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	WorkPeriod int32     `protobuf:"varint,16,opt,name=work_period,json=workPeriod,proto3" json:"work_period"`
	Memo       string    `protobuf:"bytes,17,opt,name=memo,proto3" json:"memo"`
	State      int32     `protobuf:"varint,18,opt,name=state,proto3" json:"state"`
	CreatedAt  string    `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string    `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Position   *Position `protobuf:"bytes,21,opt,name=position,proto3" json:"position"`
}

func (x *PositionOffer) Reset() {
	*x = PositionOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_positionOfferService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionOffer) ProtoMessage() {}

func (x *PositionOffer) ProtoReflect() protoreflect.Message {
	mi := &file_positionOfferService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionOffer.ProtoReflect.Descriptor instead.
func (*PositionOffer) Descriptor() ([]byte, []int) {
	return file_positionOfferService_proto_rawDescGZIP(), []int{0}
}

func (x *PositionOffer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PositionOffer) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *PositionOffer) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *PositionOffer) GetSeekerId() int64 {
	if x != nil {
		return x.SeekerId
	}
	return 0
}

func (x *PositionOffer) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *PositionOffer) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *PositionOffer) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *PositionOffer) GetNation() int32 {
	if x != nil {
		return x.Nation
	}
	return 0
}

func (x *PositionOffer) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *PositionOffer) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *PositionOffer) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *PositionOffer) GetIsSelf() bool {
	if x != nil {
		return x.IsSelf
	}
	return false
}

func (x *PositionOffer) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *PositionOffer) GetStartType() int32 {
	if x != nil {
		return x.StartType
	}
	return 0
}

func (x *PositionOffer) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *PositionOffer) GetWorkPeriod() int32 {
	if x != nil {
		return x.WorkPeriod
	}
	return 0
}

func (x *PositionOffer) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *PositionOffer) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *PositionOffer) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PositionOffer) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PositionOffer) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

//????????????
type PositionOfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//????????????????????????
	Name       string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	PositionId int64   `protobuf:"varint,5,opt,name=position_id,json=positionId,proto3" json:"position_id"`
	CustomerId int64   `protobuf:"varint,6,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	RealName   string  `protobuf:"bytes,7,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard     string  `protobuf:"bytes,8,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	Gender     int32   `protobuf:"varint,9,opt,name=gender,proto3" json:"gender"`
	Mobile     string  `protobuf:"bytes,10,opt,name=mobile,proto3" json:"mobile"`
	Nation     int32   `protobuf:"varint,11,opt,name=nation,proto3" json:"nation"`
	IsPoor     bool    `protobuf:"varint,12,opt,name=is_poor,json=isPoor,proto3" json:"is_poor"`
	StartType  int32   `protobuf:"varint,13,opt,name=start_type,json=startType,proto3" json:"start_type"`
	StartDate  string  `protobuf:"bytes,14,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	StartPlace int32   `protobuf:"varint,15,opt,name=start_place,json=startPlace,proto3" json:"start_place"`
	WorkPeriod int32   `protobuf:"varint,16,opt,name=work_period,json=workPeriod,proto3" json:"work_period"`
	State      int32   `protobuf:"varint,17,opt,name=state,proto3" json:"state"`
	Memo       string  `protobuf:"bytes,18,opt,name=memo,proto3" json:"memo"`
	IsSelf     bool    `protobuf:"varint,19,opt,name=is_self,json=isSelf,proto3" json:"is_self"`
	SeekerIds  []int64 `protobuf:"varint,20,rep,packed,name=seeker_ids,json=seekerIds,proto3" json:"seeker_ids"`
}

func (x *PositionOfferRequest) Reset() {
	*x = PositionOfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_positionOfferService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionOfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionOfferRequest) ProtoMessage() {}

func (x *PositionOfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_positionOfferService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionOfferRequest.ProtoReflect.Descriptor instead.
func (*PositionOfferRequest) Descriptor() ([]byte, []int) {
	return file_positionOfferService_proto_rawDescGZIP(), []int{1}
}

func (x *PositionOfferRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PositionOfferRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PositionOfferRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *PositionOfferRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PositionOfferRequest) GetPositionId() int64 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *PositionOfferRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *PositionOfferRequest) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *PositionOfferRequest) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *PositionOfferRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *PositionOfferRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *PositionOfferRequest) GetNation() int32 {
	if x != nil {
		return x.Nation
	}
	return 0
}

func (x *PositionOfferRequest) GetIsPoor() bool {
	if x != nil {
		return x.IsPoor
	}
	return false
}

func (x *PositionOfferRequest) GetStartType() int32 {
	if x != nil {
		return x.StartType
	}
	return 0
}

func (x *PositionOfferRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *PositionOfferRequest) GetStartPlace() int32 {
	if x != nil {
		return x.StartPlace
	}
	return 0
}

func (x *PositionOfferRequest) GetWorkPeriod() int32 {
	if x != nil {
		return x.WorkPeriod
	}
	return 0
}

func (x *PositionOfferRequest) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *PositionOfferRequest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *PositionOfferRequest) GetIsSelf() bool {
	if x != nil {
		return x.IsSelf
	}
	return false
}

func (x *PositionOfferRequest) GetSeekerIds() []int64 {
	if x != nil {
		return x.SeekerIds
	}
	return nil
}

//
type PositionOfferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *PositionOffer    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*PositionOffer  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
	Params map[string]string `protobuf:"bytes,6,rep,name=params,proto3" json:"params" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PositionOfferResponse) Reset() {
	*x = PositionOfferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_positionOfferService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionOfferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionOfferResponse) ProtoMessage() {}

func (x *PositionOfferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_positionOfferService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionOfferResponse.ProtoReflect.Descriptor instead.
func (*PositionOfferResponse) Descriptor() ([]byte, []int) {
	return file_positionOfferService_proto_rawDescGZIP(), []int{2}
}

func (x *PositionOfferResponse) GetEntity() *PositionOffer {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PositionOfferResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PositionOfferResponse) GetItems() []*PositionOffer {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PositionOfferResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PositionOfferResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *PositionOfferResponse) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_positionOfferService_proto protoreflect.FileDescriptor

var file_positionOfferService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd9, 0x04, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61,
	0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x72,
	0x65, 0x61, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x53, 0x65, 0x6c, 0x66, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb2, 0x04, 0x0a,
	0x14, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x6f,
	0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x6f, 0x6f, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f,
	0x73, 0x65, 0x6c, 0x66, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x53, 0x65,
	0x6c, 0x66, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x22, 0xe3, 0x02, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x43, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xc4, 0x03, 0x0a, 0x14, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x1a, 0x1f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x1a, 0x1f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_positionOfferService_proto_rawDescOnce sync.Once
	file_positionOfferService_proto_rawDescData = file_positionOfferService_proto_rawDesc
)

func file_positionOfferService_proto_rawDescGZIP() []byte {
	file_positionOfferService_proto_rawDescOnce.Do(func() {
		file_positionOfferService_proto_rawDescData = protoimpl.X.CompressGZIP(file_positionOfferService_proto_rawDescData)
	})
	return file_positionOfferService_proto_rawDescData
}

var file_positionOfferService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_positionOfferService_proto_goTypes = []interface{}{
	(*PositionOffer)(nil),         // 0: services.PositionOffer
	(*PositionOfferRequest)(nil),  // 1: services.PositionOfferRequest
	(*PositionOfferResponse)(nil), // 2: services.PositionOfferResponse
	nil,                           // 3: services.PositionOfferResponse.ParamsEntry
	(*Position)(nil),              // 4: services.Position
	(*common.Pager)(nil),          // 5: common.Pager
	(*common.Error)(nil),          // 6: common.Error
	(*common.Info)(nil),           // 7: common.Info
}
var file_positionOfferService_proto_depIdxs = []int32{
	4,  // 0: services.PositionOffer.position:type_name -> services.Position
	0,  // 1: services.PositionOfferResponse.entity:type_name -> services.PositionOffer
	5,  // 2: services.PositionOfferResponse.pager:type_name -> common.Pager
	0,  // 3: services.PositionOfferResponse.items:type_name -> services.PositionOffer
	6,  // 4: services.PositionOfferResponse.error:type_name -> common.Error
	7,  // 5: services.PositionOfferResponse.info:type_name -> common.Info
	3,  // 6: services.PositionOfferResponse.params:type_name -> services.PositionOfferResponse.ParamsEntry
	1,  // 7: services.PositionOfferService.Create:input_type -> services.PositionOfferRequest
	0,  // 8: services.PositionOfferService.Update:input_type -> services.PositionOffer
	0,  // 9: services.PositionOfferService.Check:input_type -> services.PositionOffer
	0,  // 10: services.PositionOfferService.Delete:input_type -> services.PositionOffer
	0,  // 11: services.PositionOfferService.Get:input_type -> services.PositionOffer
	1,  // 12: services.PositionOfferService.Search:input_type -> services.PositionOfferRequest
	2,  // 13: services.PositionOfferService.Create:output_type -> services.PositionOfferResponse
	2,  // 14: services.PositionOfferService.Update:output_type -> services.PositionOfferResponse
	2,  // 15: services.PositionOfferService.Check:output_type -> services.PositionOfferResponse
	2,  // 16: services.PositionOfferService.Delete:output_type -> services.PositionOfferResponse
	2,  // 17: services.PositionOfferService.Get:output_type -> services.PositionOfferResponse
	2,  // 18: services.PositionOfferService.Search:output_type -> services.PositionOfferResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_positionOfferService_proto_init() }
func file_positionOfferService_proto_init() {
	if File_positionOfferService_proto != nil {
		return
	}
	file_positionService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_positionOfferService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionOffer); i {
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
		file_positionOfferService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionOfferRequest); i {
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
		file_positionOfferService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionOfferResponse); i {
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
			RawDescriptor: file_positionOfferService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_positionOfferService_proto_goTypes,
		DependencyIndexes: file_positionOfferService_proto_depIdxs,
		MessageInfos:      file_positionOfferService_proto_msgTypes,
	}.Build()
	File_positionOfferService_proto = out.File
	file_positionOfferService_proto_rawDesc = nil
	file_positionOfferService_proto_goTypes = nil
	file_positionOfferService_proto_depIdxs = nil
}
