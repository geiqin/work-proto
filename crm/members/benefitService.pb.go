// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: benefitService.proto

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

type Benefit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Slug      string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	IconId    int64  `protobuf:"varint,4,opt,name=icon_id,json=iconId,proto3" json:"icon_id"`
	IconUrl   string `protobuf:"bytes,5,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url"`
	Type      string `protobuf:"bytes,6,opt,name=type,proto3" json:"type"`
	Selected  bool   `protobuf:"varint,7,opt,name=selected,proto3" json:"selected"`
	Package   bool   `protobuf:"varint,8,opt,name=package,proto3" json:"package"`
	Params    string `protobuf:"bytes,9,opt,name=params,proto3" json:"params"`
	Memo      string `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo"`
	CreatedAt string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: gorm:"-"
	Ids []int32 `protobuf:"varint,13,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
}

func (x *Benefit) Reset() {
	*x = Benefit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benefitService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Benefit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Benefit) ProtoMessage() {}

func (x *Benefit) ProtoReflect() protoreflect.Message {
	mi := &file_benefitService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Benefit.ProtoReflect.Descriptor instead.
func (*Benefit) Descriptor() ([]byte, []int) {
	return file_benefitService_proto_rawDescGZIP(), []int{0}
}

func (x *Benefit) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Benefit) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Benefit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Benefit) GetIconId() int64 {
	if x != nil {
		return x.IconId
	}
	return 0
}

func (x *Benefit) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *Benefit) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Benefit) GetSelected() bool {
	if x != nil {
		return x.Selected
	}
	return false
}

func (x *Benefit) GetPackage() bool {
	if x != nil {
		return x.Package
	}
	return false
}

func (x *Benefit) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *Benefit) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Benefit) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Benefit) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Benefit) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//查询参数
type BenefitWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	//以下为自定义参数
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Type     string `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
	Selected bool   `protobuf:"varint,5,opt,name=selected,proto3" json:"selected"`
}

func (x *BenefitWhere) Reset() {
	*x = BenefitWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benefitService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenefitWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenefitWhere) ProtoMessage() {}

func (x *BenefitWhere) ProtoReflect() protoreflect.Message {
	mi := &file_benefitService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenefitWhere.ProtoReflect.Descriptor instead.
func (*BenefitWhere) Descriptor() ([]byte, []int) {
	return file_benefitService_proto_rawDescGZIP(), []int{1}
}

func (x *BenefitWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *BenefitWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *BenefitWhere) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BenefitWhere) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BenefitWhere) GetSelected() bool {
	if x != nil {
		return x.Selected
	}
	return false
}

//
type BenefitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Benefit      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Benefit    `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *BenefitResponse) Reset() {
	*x = BenefitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benefitService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenefitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenefitResponse) ProtoMessage() {}

func (x *BenefitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_benefitService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenefitResponse.ProtoReflect.Descriptor instead.
func (*BenefitResponse) Descriptor() ([]byte, []int) {
	return file_benefitService_proto_rawDescGZIP(), []int{2}
}

func (x *BenefitResponse) GetEntity() *Benefit {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *BenefitResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *BenefitResponse) GetItems() []*Benefit {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *BenefitResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *BenefitResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_benefitService_proto protoreflect.FileDescriptor

var file_benefitService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x07, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x63, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x22, 0x85, 0x01, 0x0a, 0x0c, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0xd1, 0x01, 0x0a, 0x0f, 0x42, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xad, 0x03,
	0x0a, 0x0e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66,
	0x69, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_benefitService_proto_rawDescOnce sync.Once
	file_benefitService_proto_rawDescData = file_benefitService_proto_rawDesc
)

func file_benefitService_proto_rawDescGZIP() []byte {
	file_benefitService_proto_rawDescOnce.Do(func() {
		file_benefitService_proto_rawDescData = protoimpl.X.CompressGZIP(file_benefitService_proto_rawDescData)
	})
	return file_benefitService_proto_rawDescData
}

var file_benefitService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_benefitService_proto_goTypes = []interface{}{
	(*Benefit)(nil),         // 0: services.Benefit
	(*BenefitWhere)(nil),    // 1: services.BenefitWhere
	(*BenefitResponse)(nil), // 2: services.BenefitResponse
	(*common.Pager)(nil),    // 3: common.Pager
	(*common.Error)(nil),    // 4: common.Error
	(*common.Info)(nil),     // 5: common.Info
}
var file_benefitService_proto_depIdxs = []int32{
	0,  // 0: services.BenefitResponse.entity:type_name -> services.Benefit
	3,  // 1: services.BenefitResponse.pager:type_name -> common.Pager
	0,  // 2: services.BenefitResponse.items:type_name -> services.Benefit
	4,  // 3: services.BenefitResponse.error:type_name -> common.Error
	5,  // 4: services.BenefitResponse.info:type_name -> common.Info
	0,  // 5: services.BenefitService.Create:input_type -> services.Benefit
	0,  // 6: services.BenefitService.Update:input_type -> services.Benefit
	0,  // 7: services.BenefitService.Delete:input_type -> services.Benefit
	0,  // 8: services.BenefitService.Get:input_type -> services.Benefit
	1,  // 9: services.BenefitService.Search:input_type -> services.BenefitWhere
	1,  // 10: services.BenefitService.List:input_type -> services.BenefitWhere
	0,  // 11: services.BenefitService.Selected:input_type -> services.Benefit
	2,  // 12: services.BenefitService.Create:output_type -> services.BenefitResponse
	2,  // 13: services.BenefitService.Update:output_type -> services.BenefitResponse
	2,  // 14: services.BenefitService.Delete:output_type -> services.BenefitResponse
	2,  // 15: services.BenefitService.Get:output_type -> services.BenefitResponse
	2,  // 16: services.BenefitService.Search:output_type -> services.BenefitResponse
	2,  // 17: services.BenefitService.List:output_type -> services.BenefitResponse
	2,  // 18: services.BenefitService.Selected:output_type -> services.BenefitResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_benefitService_proto_init() }
func file_benefitService_proto_init() {
	if File_benefitService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_benefitService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Benefit); i {
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
		file_benefitService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenefitWhere); i {
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
		file_benefitService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenefitResponse); i {
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
			RawDescriptor: file_benefitService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_benefitService_proto_goTypes,
		DependencyIndexes: file_benefitService_proto_depIdxs,
		MessageInfos:      file_benefitService_proto_msgTypes,
	}.Build()
	File_benefitService_proto = out.File
	file_benefitService_proto_rawDesc = nil
	file_benefitService_proto_goTypes = nil
	file_benefitService_proto_depIdxs = nil
}
