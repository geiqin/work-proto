// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: mobileService.proto

package services

import (
	common "github.com/geiqin/microkit/protobuf/common"
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

type Mobile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile     string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	CustomerId int64  `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Type       string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Code       string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	Operator   string `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	ReadNum    int32  `protobuf:"varint,7,opt,name=read_num,json=readNum,proto3" json:"read_num,omitempty"`
	Checked    bool   `protobuf:"varint,8,opt,name=checked,proto3" json:"checked,omitempty"`
	CheckedAt  string `protobuf:"bytes,9,opt,name=checked_at,json=checkedAt,proto3" json:"checked_at,omitempty"`
	CreatedAt  string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Mobile) Reset() {
	*x = Mobile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mobileService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mobile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mobile) ProtoMessage() {}

func (x *Mobile) ProtoReflect() protoreflect.Message {
	mi := &file_mobileService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mobile.ProtoReflect.Descriptor instead.
func (*Mobile) Descriptor() ([]byte, []int) {
	return file_mobileService_proto_rawDescGZIP(), []int{0}
}

func (x *Mobile) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Mobile) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Mobile) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Mobile) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Mobile) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Mobile) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *Mobile) GetReadNum() int32 {
	if x != nil {
		return x.ReadNum
	}
	return 0
}

func (x *Mobile) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *Mobile) GetCheckedAt() string {
	if x != nil {
		return x.CheckedAt
	}
	return ""
}

func (x *Mobile) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Mobile) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//查询参数
type MobileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	//以下为自定义参数
	Mobile     string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	CustomerId int64  `protobuf:"varint,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Operator   int32  `protobuf:"varint,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Checked    bool   `protobuf:"varint,7,opt,name=checked,proto3" json:"checked,omitempty"`
}

func (x *MobileRequest) Reset() {
	*x = MobileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mobileService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MobileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileRequest) ProtoMessage() {}

func (x *MobileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mobileService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileRequest.ProtoReflect.Descriptor instead.
func (*MobileRequest) Descriptor() ([]byte, []int) {
	return file_mobileService_proto_rawDescGZIP(), []int{1}
}

func (x *MobileRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MobileRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MobileRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *MobileRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *MobileRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *MobileRequest) GetOperator() int32 {
	if x != nil {
		return x.Operator
	}
	return 0
}

func (x *MobileRequest) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

//
type MobileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Mobile       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Mobile     `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *MobileResponse) Reset() {
	*x = MobileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mobileService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MobileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileResponse) ProtoMessage() {}

func (x *MobileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mobileService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileResponse.ProtoReflect.Descriptor instead.
func (*MobileResponse) Descriptor() ([]byte, []int) {
	return file_mobileService_proto_rawDescGZIP(), []int{2}
}

func (x *MobileResponse) GetEntity() *Mobile {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MobileResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MobileResponse) GetItems() []*Mobile {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MobileResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *MobileResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_mobileService_proto protoreflect.FileDescriptor

var file_mobileService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65,
	0x61, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x61, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xcb, 0x01, 0x0a,
	0x0d, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x22, 0xce, 0x01, 0x0a, 0x0e, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x83, 0x01, 0x0a, 0x0d,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mobileService_proto_rawDescOnce sync.Once
	file_mobileService_proto_rawDescData = file_mobileService_proto_rawDesc
)

func file_mobileService_proto_rawDescGZIP() []byte {
	file_mobileService_proto_rawDescOnce.Do(func() {
		file_mobileService_proto_rawDescData = protoimpl.X.CompressGZIP(file_mobileService_proto_rawDescData)
	})
	return file_mobileService_proto_rawDescData
}

var file_mobileService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mobileService_proto_goTypes = []interface{}{
	(*Mobile)(nil),         // 0: services.Mobile
	(*MobileRequest)(nil),  // 1: services.MobileRequest
	(*MobileResponse)(nil), // 2: services.MobileResponse
	(*common.Pager)(nil),   // 3: common.Pager
	(*common.Error)(nil),   // 4: common.Error
	(*common.Info)(nil),    // 5: common.Info
}
var file_mobileService_proto_depIdxs = []int32{
	0, // 0: services.MobileResponse.entity:type_name -> services.Mobile
	3, // 1: services.MobileResponse.pager:type_name -> common.Pager
	0, // 2: services.MobileResponse.items:type_name -> services.Mobile
	4, // 3: services.MobileResponse.error:type_name -> common.Error
	5, // 4: services.MobileResponse.info:type_name -> common.Info
	0, // 5: services.MobileService.Get:input_type -> services.Mobile
	1, // 6: services.MobileService.Search:input_type -> services.MobileRequest
	2, // 7: services.MobileService.Get:output_type -> services.MobileResponse
	2, // 8: services.MobileService.Search:output_type -> services.MobileResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mobileService_proto_init() }
func file_mobileService_proto_init() {
	if File_mobileService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mobileService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mobile); i {
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
		file_mobileService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MobileRequest); i {
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
		file_mobileService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MobileResponse); i {
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
			RawDescriptor: file_mobileService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mobileService_proto_goTypes,
		DependencyIndexes: file_mobileService_proto_depIdxs,
		MessageInfos:      file_mobileService_proto_msgTypes,
	}.Build()
	File_mobileService_proto = out.File
	file_mobileService_proto_rawDesc = nil
	file_mobileService_proto_goTypes = nil
	file_mobileService_proto_depIdxs = nil
}