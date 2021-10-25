// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: inductionService.proto

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

type Induction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	EnterpriseId        int64       `protobuf:"varint,2,opt,name=enterprise_id,json=enterpriseId,proto3" json:"enterprise_id"`
	DeptName            string      `protobuf:"bytes,3,opt,name=dept_name,json=deptName,proto3" json:"dept_name"`
	JobNo               string      `protobuf:"bytes,4,opt,name=job_no,json=jobNo,proto3" json:"job_no"`
	CustomerId          int64       `protobuf:"varint,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	RecommenderId       int64       `protobuf:"varint,6,opt,name=recommender_id,json=recommenderId,proto3" json:"recommender_id"`
	RealName            string      `protobuf:"bytes,7,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard              string      `protobuf:"bytes,8,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	Gender              int32       `protobuf:"varint,9,opt,name=gender,proto3" json:"gender"`
	Nation              int32       `protobuf:"varint,10,opt,name=nation,proto3" json:"nation"`
	Birthday            string      `protobuf:"bytes,11,opt,name=birthday,proto3" json:"birthday"`
	Mobile              string      `protobuf:"bytes,12,opt,name=mobile,proto3" json:"mobile"`
	WorkType            int32       `protobuf:"varint,13,opt,name=work_type,json=workType,proto3" json:"work_type"`
	BasicSalary         float32     `protobuf:"fixed32,14,opt,name=basic_salary,json=basicSalary,proto3" json:"basic_salary"`
	HourPrice           float32     `protobuf:"fixed32,15,opt,name=hour_price,json=hourPrice,proto3" json:"hour_price"`
	CostPrice           float32     `protobuf:"fixed32,16,opt,name=cost_price,json=costPrice,proto3" json:"cost_price"`
	SourcePositionId    int64       `protobuf:"varint,17,opt,name=source_position_id,json=sourcePositionId,proto3" json:"source_position_id"`
	SourcePositionTitle string      `protobuf:"bytes,18,opt,name=source_position_title,json=sourcePositionTitle,proto3" json:"source_position_title"`
	EntryDate           string      `protobuf:"bytes,19,opt,name=entry_date,json=entryDate,proto3" json:"entry_date"`
	QuitDate            string      `protobuf:"bytes,20,opt,name=quit_date,json=quitDate,proto3" json:"quit_date"`
	Demo                string      `protobuf:"bytes,21,opt,name=demo,proto3" json:"demo"`
	State               int32       `protobuf:"varint,22,opt,name=state,proto3" json:"state"`
	CreatedAt           string      `protobuf:"bytes,23,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt           string      `protobuf:"bytes,24,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Enterprise          *Enterprise `protobuf:"bytes,25,opt,name=enterprise,proto3" json:"enterprise"`
}

func (x *Induction) Reset() {
	*x = Induction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inductionService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Induction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Induction) ProtoMessage() {}

func (x *Induction) ProtoReflect() protoreflect.Message {
	mi := &file_inductionService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Induction.ProtoReflect.Descriptor instead.
func (*Induction) Descriptor() ([]byte, []int) {
	return file_inductionService_proto_rawDescGZIP(), []int{0}
}

func (x *Induction) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Induction) GetEnterpriseId() int64 {
	if x != nil {
		return x.EnterpriseId
	}
	return 0
}

func (x *Induction) GetDeptName() string {
	if x != nil {
		return x.DeptName
	}
	return ""
}

func (x *Induction) GetJobNo() string {
	if x != nil {
		return x.JobNo
	}
	return ""
}

func (x *Induction) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Induction) GetRecommenderId() int64 {
	if x != nil {
		return x.RecommenderId
	}
	return 0
}

func (x *Induction) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *Induction) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *Induction) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Induction) GetNation() int32 {
	if x != nil {
		return x.Nation
	}
	return 0
}

func (x *Induction) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *Induction) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Induction) GetWorkType() int32 {
	if x != nil {
		return x.WorkType
	}
	return 0
}

func (x *Induction) GetBasicSalary() float32 {
	if x != nil {
		return x.BasicSalary
	}
	return 0
}

func (x *Induction) GetHourPrice() float32 {
	if x != nil {
		return x.HourPrice
	}
	return 0
}

func (x *Induction) GetCostPrice() float32 {
	if x != nil {
		return x.CostPrice
	}
	return 0
}

func (x *Induction) GetSourcePositionId() int64 {
	if x != nil {
		return x.SourcePositionId
	}
	return 0
}

func (x *Induction) GetSourcePositionTitle() string {
	if x != nil {
		return x.SourcePositionTitle
	}
	return ""
}

func (x *Induction) GetEntryDate() string {
	if x != nil {
		return x.EntryDate
	}
	return ""
}

func (x *Induction) GetQuitDate() string {
	if x != nil {
		return x.QuitDate
	}
	return ""
}

func (x *Induction) GetDemo() string {
	if x != nil {
		return x.Demo
	}
	return ""
}

func (x *Induction) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *Induction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Induction) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Induction) GetEnterprise() *Enterprise {
	if x != nil {
		return x.Enterprise
	}
	return nil
}

//标签查询参数
type InductionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//以下为自定义参数
	Id              int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	EnterpriseId    int64  `protobuf:"varint,5,opt,name=enterprise_id,json=enterpriseId,proto3" json:"enterprise_id"`
	JobNo           string `protobuf:"bytes,6,opt,name=job_no,json=jobNo,proto3" json:"job_no"`
	RecommenderId   int64  `protobuf:"varint,7,opt,name=recommender_id,json=recommenderId,proto3" json:"recommender_id"`
	RealName        string `protobuf:"bytes,8,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard          string `protobuf:"bytes,9,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	Gender          int32  `protobuf:"varint,10,opt,name=gender,proto3" json:"gender"`
	Nation          int32  `protobuf:"varint,11,opt,name=nation,proto3" json:"nation"`
	Birthday        string `protobuf:"bytes,12,opt,name=birthday,proto3" json:"birthday"`
	Mobile          string `protobuf:"bytes,13,opt,name=mobile,proto3" json:"mobile"`
	WorkType        int32  `protobuf:"varint,14,opt,name=work_type,json=workType,proto3" json:"work_type"`
	EntryDate       string `protobuf:"bytes,15,opt,name=entry_date,json=entryDate,proto3" json:"entry_date"`
	QuitDate        string `protobuf:"bytes,16,opt,name=quit_date,json=quitDate,proto3" json:"quit_date"`
	Demo            string `protobuf:"bytes,17,opt,name=demo,proto3" json:"demo"`
	State           int32  `protobuf:"varint,18,opt,name=state,proto3" json:"state"`
	CustomerId      int64  `protobuf:"varint,19,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	PositionOfferId int64  `protobuf:"varint,20,opt,name=position_offer_id,json=positionOfferId,proto3" json:"position_offer_id"`
}

func (x *InductionRequest) Reset() {
	*x = InductionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inductionService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InductionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InductionRequest) ProtoMessage() {}

func (x *InductionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inductionService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InductionRequest.ProtoReflect.Descriptor instead.
func (*InductionRequest) Descriptor() ([]byte, []int) {
	return file_inductionService_proto_rawDescGZIP(), []int{1}
}

func (x *InductionRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *InductionRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *InductionRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *InductionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InductionRequest) GetEnterpriseId() int64 {
	if x != nil {
		return x.EnterpriseId
	}
	return 0
}

func (x *InductionRequest) GetJobNo() string {
	if x != nil {
		return x.JobNo
	}
	return ""
}

func (x *InductionRequest) GetRecommenderId() int64 {
	if x != nil {
		return x.RecommenderId
	}
	return 0
}

func (x *InductionRequest) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *InductionRequest) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *InductionRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *InductionRequest) GetNation() int32 {
	if x != nil {
		return x.Nation
	}
	return 0
}

func (x *InductionRequest) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *InductionRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *InductionRequest) GetWorkType() int32 {
	if x != nil {
		return x.WorkType
	}
	return 0
}

func (x *InductionRequest) GetEntryDate() string {
	if x != nil {
		return x.EntryDate
	}
	return ""
}

func (x *InductionRequest) GetQuitDate() string {
	if x != nil {
		return x.QuitDate
	}
	return ""
}

func (x *InductionRequest) GetDemo() string {
	if x != nil {
		return x.Demo
	}
	return ""
}

func (x *InductionRequest) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *InductionRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *InductionRequest) GetPositionOfferId() int64 {
	if x != nil {
		return x.PositionOfferId
	}
	return 0
}

//
type InductionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Induction    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Induction  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *InductionResponse) Reset() {
	*x = InductionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inductionService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InductionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InductionResponse) ProtoMessage() {}

func (x *InductionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inductionService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InductionResponse.ProtoReflect.Descriptor instead.
func (*InductionResponse) Descriptor() ([]byte, []int) {
	return file_inductionService_proto_rawDescGZIP(), []int{2}
}

func (x *InductionResponse) GetEntity() *Induction {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *InductionResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *InductionResponse) GetItems() []*Induction {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *InductionResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *InductionResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_inductionService_proto protoreflect.FileDescriptor

var file_inductionService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90,
	0x06, 0x0a, 0x09, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x70, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6a, 0x6f, 0x62, 0x4e, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x75, 0x72, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x68, 0x6f, 0x75,
	0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x71, 0x75, 0x69, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x69, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x0a, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x22, 0xbc, 0x04, 0x0a, 0x10, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f,
	0x6e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x4e, 0x6f, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x71, 0x75, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x71, 0x75, 0x69, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x6d,
	0x6f, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xd7, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x94, 0x04, 0x0a, 0x10, 0x49,
	0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49,
	0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x49, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inductionService_proto_rawDescOnce sync.Once
	file_inductionService_proto_rawDescData = file_inductionService_proto_rawDesc
)

func file_inductionService_proto_rawDescGZIP() []byte {
	file_inductionService_proto_rawDescOnce.Do(func() {
		file_inductionService_proto_rawDescData = protoimpl.X.CompressGZIP(file_inductionService_proto_rawDescData)
	})
	return file_inductionService_proto_rawDescData
}

var file_inductionService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_inductionService_proto_goTypes = []interface{}{
	(*Induction)(nil),         // 0: services.Induction
	(*InductionRequest)(nil),  // 1: services.InductionRequest
	(*InductionResponse)(nil), // 2: services.InductionResponse
	(*Enterprise)(nil),        // 3: services.Enterprise
	(*common.Pager)(nil),      // 4: common.Pager
	(*common.Error)(nil),      // 5: common.Error
	(*common.Info)(nil),       // 6: common.Info
	(*common.Empty)(nil),      // 7: common.Empty
}
var file_inductionService_proto_depIdxs = []int32{
	3,  // 0: services.Induction.enterprise:type_name -> services.Enterprise
	0,  // 1: services.InductionResponse.entity:type_name -> services.Induction
	4,  // 2: services.InductionResponse.pager:type_name -> common.Pager
	0,  // 3: services.InductionResponse.items:type_name -> services.Induction
	5,  // 4: services.InductionResponse.error:type_name -> common.Error
	6,  // 5: services.InductionResponse.info:type_name -> common.Info
	0,  // 6: services.InductionService.Create:input_type -> services.Induction
	0,  // 7: services.InductionService.Update:input_type -> services.Induction
	1,  // 8: services.InductionService.Entry:input_type -> services.InductionRequest
	1,  // 9: services.InductionService.Quit:input_type -> services.InductionRequest
	0,  // 10: services.InductionService.Delete:input_type -> services.Induction
	0,  // 11: services.InductionService.Get:input_type -> services.Induction
	1,  // 12: services.InductionService.Search:input_type -> services.InductionRequest
	7,  // 13: services.InductionService.GetByCustomerId:input_type -> common.Empty
	2,  // 14: services.InductionService.Create:output_type -> services.InductionResponse
	2,  // 15: services.InductionService.Update:output_type -> services.InductionResponse
	2,  // 16: services.InductionService.Entry:output_type -> services.InductionResponse
	2,  // 17: services.InductionService.Quit:output_type -> services.InductionResponse
	2,  // 18: services.InductionService.Delete:output_type -> services.InductionResponse
	2,  // 19: services.InductionService.Get:output_type -> services.InductionResponse
	2,  // 20: services.InductionService.Search:output_type -> services.InductionResponse
	2,  // 21: services.InductionService.GetByCustomerId:output_type -> services.InductionResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_inductionService_proto_init() }
func file_inductionService_proto_init() {
	if File_inductionService_proto != nil {
		return
	}
	file_enterpriseService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_inductionService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Induction); i {
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
		file_inductionService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InductionRequest); i {
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
		file_inductionService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InductionResponse); i {
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
			RawDescriptor: file_inductionService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inductionService_proto_goTypes,
		DependencyIndexes: file_inductionService_proto_depIdxs,
		MessageInfos:      file_inductionService_proto_msgTypes,
	}.Build()
	File_inductionService_proto = out.File
	file_inductionService_proto_rawDesc = nil
	file_inductionService_proto_goTypes = nil
	file_inductionService_proto_depIdxs = nil
}