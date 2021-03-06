// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: posterGalleryService.proto

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

type PosterGalleryWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Id       int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: gorm:"-"
	Ids []int64 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *PosterGalleryWhere) Reset() {
	*x = PosterGalleryWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posterGalleryService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosterGalleryWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosterGalleryWhere) ProtoMessage() {}

func (x *PosterGalleryWhere) ProtoReflect() protoreflect.Message {
	mi := &file_posterGalleryService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosterGalleryWhere.ProtoReflect.Descriptor instead.
func (*PosterGalleryWhere) Descriptor() ([]byte, []int) {
	return file_posterGalleryService_proto_rawDescGZIP(), []int{0}
}

func (x *PosterGalleryWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PosterGalleryWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PosterGalleryWhere) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *PosterGalleryWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PosterGalleryWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type PosterGallery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MediaId   int64  `protobuf:"varint,2,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	MediaUrl  string `protobuf:"bytes,3,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	Sorting   int32  `protobuf:"varint,4,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Defaulted bool   `protobuf:"varint,5,opt,name=defaulted,proto3" json:"defaulted,omitempty"`
}

func (x *PosterGallery) Reset() {
	*x = PosterGallery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posterGalleryService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosterGallery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosterGallery) ProtoMessage() {}

func (x *PosterGallery) ProtoReflect() protoreflect.Message {
	mi := &file_posterGalleryService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosterGallery.ProtoReflect.Descriptor instead.
func (*PosterGallery) Descriptor() ([]byte, []int) {
	return file_posterGalleryService_proto_rawDescGZIP(), []int{1}
}

func (x *PosterGallery) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PosterGallery) GetMediaId() int64 {
	if x != nil {
		return x.MediaId
	}
	return 0
}

func (x *PosterGallery) GetMediaUrl() string {
	if x != nil {
		return x.MediaUrl
	}
	return ""
}

func (x *PosterGallery) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *PosterGallery) GetDefaulted() bool {
	if x != nil {
		return x.Defaulted
	}
	return false
}

type PosterGalleryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *PosterGallery   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*PosterGallery `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error    `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info     `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PosterGalleryResponse) Reset() {
	*x = PosterGalleryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posterGalleryService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosterGalleryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosterGalleryResponse) ProtoMessage() {}

func (x *PosterGalleryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_posterGalleryService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosterGalleryResponse.ProtoReflect.Descriptor instead.
func (*PosterGalleryResponse) Descriptor() ([]byte, []int) {
	return file_posterGalleryService_proto_rawDescGZIP(), []int{2}
}

func (x *PosterGalleryResponse) GetEntity() *PosterGallery {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PosterGalleryResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PosterGalleryResponse) GetItems() []*PosterGallery {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PosterGalleryResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PosterGalleryResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_posterGalleryService_proto protoreflect.FileDescriptor

var file_posterGalleryService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x50, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22,
	0x8f, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x65,
	0x64, 0x22, 0xe3, 0x01, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x86, 0x03, 0x0a, 0x14, 0x50, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x41, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x32, 0xf8, 0x01, 0x0a, 0x19, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a,
	0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x04, 0x4d, 0x61, 0x6b, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_posterGalleryService_proto_rawDescOnce sync.Once
	file_posterGalleryService_proto_rawDescData = file_posterGalleryService_proto_rawDesc
)

func file_posterGalleryService_proto_rawDescGZIP() []byte {
	file_posterGalleryService_proto_rawDescOnce.Do(func() {
		file_posterGalleryService_proto_rawDescData = protoimpl.X.CompressGZIP(file_posterGalleryService_proto_rawDescData)
	})
	return file_posterGalleryService_proto_rawDescData
}

var file_posterGalleryService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_posterGalleryService_proto_goTypes = []interface{}{
	(*PosterGalleryWhere)(nil),    // 0: services.PosterGalleryWhere
	(*PosterGallery)(nil),         // 1: services.PosterGallery
	(*PosterGalleryResponse)(nil), // 2: services.PosterGalleryResponse
	(*common.Pager)(nil),          // 3: common.Pager
	(*common.Error)(nil),          // 4: common.Error
	(*common.Info)(nil),           // 5: common.Info
}
var file_posterGalleryService_proto_depIdxs = []int32{
	1,  // 0: services.PosterGalleryResponse.entity:type_name -> services.PosterGallery
	3,  // 1: services.PosterGalleryResponse.pager:type_name -> common.Pager
	1,  // 2: services.PosterGalleryResponse.items:type_name -> services.PosterGallery
	4,  // 3: services.PosterGalleryResponse.error:type_name -> common.Error
	5,  // 4: services.PosterGalleryResponse.info:type_name -> common.Info
	1,  // 5: services.PosterGalleryService.Add:input_type -> services.PosterGallery
	0,  // 6: services.PosterGalleryService.Delete:input_type -> services.PosterGalleryWhere
	0,  // 7: services.PosterGalleryService.DeleteAll:input_type -> services.PosterGalleryWhere
	0,  // 8: services.PosterGalleryService.Search:input_type -> services.PosterGalleryWhere
	0,  // 9: services.PosterGalleryService.List:input_type -> services.PosterGalleryWhere
	0,  // 10: services.FrontPosterGalleryService.Search:input_type -> services.PosterGalleryWhere
	0,  // 11: services.FrontPosterGalleryService.List:input_type -> services.PosterGalleryWhere
	0,  // 12: services.FrontPosterGalleryService.Make:input_type -> services.PosterGalleryWhere
	2,  // 13: services.PosterGalleryService.Add:output_type -> services.PosterGalleryResponse
	2,  // 14: services.PosterGalleryService.Delete:output_type -> services.PosterGalleryResponse
	2,  // 15: services.PosterGalleryService.DeleteAll:output_type -> services.PosterGalleryResponse
	2,  // 16: services.PosterGalleryService.Search:output_type -> services.PosterGalleryResponse
	2,  // 17: services.PosterGalleryService.List:output_type -> services.PosterGalleryResponse
	2,  // 18: services.FrontPosterGalleryService.Search:output_type -> services.PosterGalleryResponse
	2,  // 19: services.FrontPosterGalleryService.List:output_type -> services.PosterGalleryResponse
	2,  // 20: services.FrontPosterGalleryService.Make:output_type -> services.PosterGalleryResponse
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_posterGalleryService_proto_init() }
func file_posterGalleryService_proto_init() {
	if File_posterGalleryService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_posterGalleryService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosterGalleryWhere); i {
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
		file_posterGalleryService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosterGallery); i {
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
		file_posterGalleryService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosterGalleryResponse); i {
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
			RawDescriptor: file_posterGalleryService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_posterGalleryService_proto_goTypes,
		DependencyIndexes: file_posterGalleryService_proto_depIdxs,
		MessageInfos:      file_posterGalleryService_proto_msgTypes,
	}.Build()
	File_posterGalleryService_proto = out.File
	file_posterGalleryService_proto_rawDesc = nil
	file_posterGalleryService_proto_goTypes = nil
	file_posterGalleryService_proto_depIdxs = nil
}
