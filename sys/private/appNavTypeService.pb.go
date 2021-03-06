// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: appNavTypeService.proto

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

type AppNavTypeWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged         int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize      int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords      string  `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	Id            int32   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Ids           []int32 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids"`
	ApplicationId int32   `protobuf:"varint,6,opt,name=application_id,json=applicationId,proto3" json:"application_id"`
}

func (x *AppNavTypeWhere) Reset() {
	*x = AppNavTypeWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appNavTypeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppNavTypeWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppNavTypeWhere) ProtoMessage() {}

func (x *AppNavTypeWhere) ProtoReflect() protoreflect.Message {
	mi := &file_appNavTypeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppNavTypeWhere.ProtoReflect.Descriptor instead.
func (*AppNavTypeWhere) Descriptor() ([]byte, []int) {
	return file_appNavTypeService_proto_rawDescGZIP(), []int{0}
}

func (x *AppNavTypeWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *AppNavTypeWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AppNavTypeWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *AppNavTypeWhere) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppNavTypeWhere) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *AppNavTypeWhere) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

// ????????????
type AppNavType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name      string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Memo      string    `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo"`
	Sorting   int32     `protobuf:"varint,4,opt,name=sorting,proto3" json:"sorting"`
	Navs      []*AppNav `protobuf:"bytes,5,rep,name=navs,proto3" json:"navs"`
	CreatedAt string    `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string    `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *AppNavType) Reset() {
	*x = AppNavType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appNavTypeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppNavType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppNavType) ProtoMessage() {}

func (x *AppNavType) ProtoReflect() protoreflect.Message {
	mi := &file_appNavTypeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppNavType.ProtoReflect.Descriptor instead.
func (*AppNavType) Descriptor() ([]byte, []int) {
	return file_appNavTypeService_proto_rawDescGZIP(), []int{1}
}

func (x *AppNavType) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppNavType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppNavType) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *AppNavType) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *AppNavType) GetNavs() []*AppNav {
	if x != nil {
		return x.Navs
	}
	return nil
}

func (x *AppNavType) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AppNavType) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type AppNavTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *AppNavType   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*AppNavType `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *AppNavTypeResponse) Reset() {
	*x = AppNavTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appNavTypeService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppNavTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppNavTypeResponse) ProtoMessage() {}

func (x *AppNavTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appNavTypeService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppNavTypeResponse.ProtoReflect.Descriptor instead.
func (*AppNavTypeResponse) Descriptor() ([]byte, []int) {
	return file_appNavTypeService_proto_rawDescGZIP(), []int{2}
}

func (x *AppNavTypeResponse) GetEntity() *AppNavType {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *AppNavTypeResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *AppNavTypeResponse) GetItems() []*AppNavType {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *AppNavTypeResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *AppNavTypeResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_appNavTypeService_proto protoreflect.FileDescriptor

var file_appNavTypeService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0f,
	0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x4e,
	0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x76, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x52, 0x04, 0x6e, 0x61, 0x76, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xda, 0x01, 0x0a,
	0x12, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xa4, 0x04, 0x0a, 0x11, 0x41, 0x70,
	0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x1a,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61,
	0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x1a,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61,
	0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70,
	0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76,
	0x54, 0x79, 0x70, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0a, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x04, 0x54, 0x72, 0x65, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e,
	0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appNavTypeService_proto_rawDescOnce sync.Once
	file_appNavTypeService_proto_rawDescData = file_appNavTypeService_proto_rawDesc
)

func file_appNavTypeService_proto_rawDescGZIP() []byte {
	file_appNavTypeService_proto_rawDescOnce.Do(func() {
		file_appNavTypeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_appNavTypeService_proto_rawDescData)
	})
	return file_appNavTypeService_proto_rawDescData
}

var file_appNavTypeService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_appNavTypeService_proto_goTypes = []interface{}{
	(*AppNavTypeWhere)(nil),    // 0: services.AppNavTypeWhere
	(*AppNavType)(nil),         // 1: services.AppNavType
	(*AppNavTypeResponse)(nil), // 2: services.AppNavTypeResponse
	(*AppNav)(nil),             // 3: services.AppNav
	(*common.Pager)(nil),       // 4: common.Pager
	(*common.Error)(nil),       // 5: common.Error
	(*common.Info)(nil),        // 6: common.Info
}
var file_appNavTypeService_proto_depIdxs = []int32{
	3,  // 0: services.AppNavType.navs:type_name -> services.AppNav
	1,  // 1: services.AppNavTypeResponse.entity:type_name -> services.AppNavType
	4,  // 2: services.AppNavTypeResponse.pager:type_name -> common.Pager
	1,  // 3: services.AppNavTypeResponse.items:type_name -> services.AppNavType
	5,  // 4: services.AppNavTypeResponse.error:type_name -> common.Error
	6,  // 5: services.AppNavTypeResponse.info:type_name -> common.Info
	1,  // 6: services.AppNavTypeService.Create:input_type -> services.AppNavType
	1,  // 7: services.AppNavTypeService.Update:input_type -> services.AppNavType
	0,  // 8: services.AppNavTypeService.Delete:input_type -> services.AppNavTypeWhere
	1,  // 9: services.AppNavTypeService.Get:input_type -> services.AppNavType
	0,  // 10: services.AppNavTypeService.List:input_type -> services.AppNavTypeWhere
	0,  // 11: services.AppNavTypeService.Search:input_type -> services.AppNavTypeWhere
	1,  // 12: services.AppNavTypeService.ChangeSort:input_type -> services.AppNavType
	0,  // 13: services.AppNavTypeService.Tree:input_type -> services.AppNavTypeWhere
	2,  // 14: services.AppNavTypeService.Create:output_type -> services.AppNavTypeResponse
	2,  // 15: services.AppNavTypeService.Update:output_type -> services.AppNavTypeResponse
	2,  // 16: services.AppNavTypeService.Delete:output_type -> services.AppNavTypeResponse
	2,  // 17: services.AppNavTypeService.Get:output_type -> services.AppNavTypeResponse
	2,  // 18: services.AppNavTypeService.List:output_type -> services.AppNavTypeResponse
	2,  // 19: services.AppNavTypeService.Search:output_type -> services.AppNavTypeResponse
	2,  // 20: services.AppNavTypeService.ChangeSort:output_type -> services.AppNavTypeResponse
	2,  // 21: services.AppNavTypeService.Tree:output_type -> services.AppNavTypeResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_appNavTypeService_proto_init() }
func file_appNavTypeService_proto_init() {
	if File_appNavTypeService_proto != nil {
		return
	}
	file_appNavService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_appNavTypeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppNavTypeWhere); i {
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
		file_appNavTypeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppNavType); i {
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
		file_appNavTypeService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppNavTypeResponse); i {
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
			RawDescriptor: file_appNavTypeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appNavTypeService_proto_goTypes,
		DependencyIndexes: file_appNavTypeService_proto_depIdxs,
		MessageInfos:      file_appNavTypeService_proto_msgTypes,
	}.Build()
	File_appNavTypeService_proto = out.File
	file_appNavTypeService_proto_rawDesc = nil
	file_appNavTypeService_proto_goTypes = nil
	file_appNavTypeService_proto_depIdxs = nil
}
