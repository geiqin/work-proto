// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: sheetService.proto

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

type SheetWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id       int32 `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids      []int32 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	Title    string  `protobuf:"bytes,5,opt,name=title,proto3" json:"title"`
	Keywords string  `protobuf:"bytes,6,opt,name=keywords,proto3" json:"keywords"`
	Slug     string  `protobuf:"bytes,7,opt,name=slug,proto3" json:"slug"`
}

func (x *SheetWhere) Reset() {
	*x = SheetWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheetService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheetWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetWhere) ProtoMessage() {}

func (x *SheetWhere) ProtoReflect() protoreflect.Message {
	mi := &file_sheetService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetWhere.ProtoReflect.Descriptor instead.
func (*SheetWhere) Descriptor() ([]byte, []int) {
	return file_sheetService_proto_rawDescGZIP(), []int{0}
}

func (x *SheetWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *SheetWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SheetWhere) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SheetWhere) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *SheetWhere) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SheetWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *SheetWhere) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type Sheet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Slug           string       `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug"`
	Title          string       `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	Memo           string       `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo"`
	Content        string       `protobuf:"bytes,5,opt,name=content,proto3" json:"content"`
	Keywords       string       `protobuf:"bytes,6,opt,name=keywords,proto3" json:"keywords"`
	ExpiryDate     string       `protobuf:"bytes,7,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date"`
	Picture        string       `protobuf:"bytes,8,opt,name=picture,proto3" json:"picture"`
	SucceedTip     string       `protobuf:"bytes,9,opt,name=succeed_tip,json=succeedTip,proto3" json:"succeed_tip"`
	FailedTip      string       `protobuf:"bytes,10,opt,name=failed_tip,json=failedTip,proto3" json:"failed_tip"`
	CreatedAt      string       `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string       `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	SheetAttr      []*SheetAttr `protobuf:"bytes,13,rep,name=sheet_attr,json=sheetAttr,proto3" json:"sheet_attr"`
	CanSubmitUser  bool         `protobuf:"varint,14,opt,name=can_submit_user,json=canSubmitUser,proto3" json:"can_submit_user"`
	CanSubmitAdmin bool         `protobuf:"varint,15,opt,name=can_submit_admin,json=canSubmitAdmin,proto3" json:"can_submit_admin"`
}

func (x *Sheet) Reset() {
	*x = Sheet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheetService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sheet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sheet) ProtoMessage() {}

func (x *Sheet) ProtoReflect() protoreflect.Message {
	mi := &file_sheetService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sheet.ProtoReflect.Descriptor instead.
func (*Sheet) Descriptor() ([]byte, []int) {
	return file_sheetService_proto_rawDescGZIP(), []int{1}
}

func (x *Sheet) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sheet) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Sheet) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Sheet) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Sheet) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Sheet) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *Sheet) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *Sheet) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Sheet) GetSucceedTip() string {
	if x != nil {
		return x.SucceedTip
	}
	return ""
}

func (x *Sheet) GetFailedTip() string {
	if x != nil {
		return x.FailedTip
	}
	return ""
}

func (x *Sheet) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Sheet) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Sheet) GetSheetAttr() []*SheetAttr {
	if x != nil {
		return x.SheetAttr
	}
	return nil
}

func (x *Sheet) GetCanSubmitUser() bool {
	if x != nil {
		return x.CanSubmitUser
	}
	return false
}

func (x *Sheet) GetCanSubmitAdmin() bool {
	if x != nil {
		return x.CanSubmitAdmin
	}
	return false
}

type SheetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *Sheet        `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*Sheet      `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *SheetResponse) Reset() {
	*x = SheetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheetService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetResponse) ProtoMessage() {}

func (x *SheetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sheetService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetResponse.ProtoReflect.Descriptor instead.
func (*SheetResponse) Descriptor() ([]byte, []int) {
	return file_sheetService_proto_rawDescGZIP(), []int{2}
}

func (x *SheetResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SheetResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *SheetResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SheetResponse) GetEntity() *Sheet {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SheetResponse) GetItems() []*Sheet {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_sheetService_proto protoreflect.FileDescriptor

var file_sheetService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x68, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x73, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0a, 0x53, 0x68,
	0x65, 0x65, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x22, 0xca, 0x03, 0x0a, 0x05, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x54, 0x69, 0x70, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x70, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x52, 0x09, 0x73, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x26,
	0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x63, 0x61, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x22, 0xcb, 0x01, 0x0a, 0x0d, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xe1,
	0x02, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65,
	0x65, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sheetService_proto_rawDescOnce sync.Once
	file_sheetService_proto_rawDescData = file_sheetService_proto_rawDesc
)

func file_sheetService_proto_rawDescGZIP() []byte {
	file_sheetService_proto_rawDescOnce.Do(func() {
		file_sheetService_proto_rawDescData = protoimpl.X.CompressGZIP(file_sheetService_proto_rawDescData)
	})
	return file_sheetService_proto_rawDescData
}

var file_sheetService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sheetService_proto_goTypes = []interface{}{
	(*SheetWhere)(nil),    // 0: services.SheetWhere
	(*Sheet)(nil),         // 1: services.Sheet
	(*SheetResponse)(nil), // 2: services.SheetResponse
	(*SheetAttr)(nil),     // 3: services.SheetAttr
	(*common.Error)(nil),  // 4: common.Error
	(*common.Info)(nil),   // 5: common.Info
	(*common.Pager)(nil),  // 6: common.Pager
}
var file_sheetService_proto_depIdxs = []int32{
	3,  // 0: services.Sheet.sheet_attr:type_name -> services.SheetAttr
	4,  // 1: services.SheetResponse.error:type_name -> common.Error
	5,  // 2: services.SheetResponse.info:type_name -> common.Info
	6,  // 3: services.SheetResponse.pager:type_name -> common.Pager
	1,  // 4: services.SheetResponse.entity:type_name -> services.Sheet
	1,  // 5: services.SheetResponse.items:type_name -> services.Sheet
	1,  // 6: services.SheetService.Create:input_type -> services.Sheet
	1,  // 7: services.SheetService.Update:input_type -> services.Sheet
	0,  // 8: services.SheetService.Delete:input_type -> services.SheetWhere
	0,  // 9: services.SheetService.Get:input_type -> services.SheetWhere
	0,  // 10: services.SheetService.Search:input_type -> services.SheetWhere
	0,  // 11: services.SheetService.List:input_type -> services.SheetWhere
	2,  // 12: services.SheetService.Create:output_type -> services.SheetResponse
	2,  // 13: services.SheetService.Update:output_type -> services.SheetResponse
	2,  // 14: services.SheetService.Delete:output_type -> services.SheetResponse
	2,  // 15: services.SheetService.Get:output_type -> services.SheetResponse
	2,  // 16: services.SheetService.Search:output_type -> services.SheetResponse
	2,  // 17: services.SheetService.List:output_type -> services.SheetResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_sheetService_proto_init() }
func file_sheetService_proto_init() {
	if File_sheetService_proto != nil {
		return
	}
	file_sheetAttrService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sheetService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheetWhere); i {
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
		file_sheetService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sheet); i {
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
		file_sheetService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheetResponse); i {
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
			RawDescriptor: file_sheetService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sheetService_proto_goTypes,
		DependencyIndexes: file_sheetService_proto_depIdxs,
		MessageInfos:      file_sheetService_proto_msgTypes,
	}.Build()
	File_sheetService_proto = out.File
	file_sheetService_proto_rawDesc = nil
	file_sheetService_proto_goTypes = nil
	file_sheetService_proto_depIdxs = nil
}
