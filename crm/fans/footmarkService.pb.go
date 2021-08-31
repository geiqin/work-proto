// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: footmarkService.proto

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

type Footmark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	OpenId        string `protobuf:"bytes,4,opt,name=open_id,json=openId,proto3" json:"open_id"`
	UnionId       string `protobuf:"bytes,5,opt,name=union_id,json=unionId,proto3" json:"union_id"`
	Source        string `protobuf:"bytes,6,opt,name=source,proto3" json:"source"`
	RecommenderId int64  `protobuf:"varint,16,opt,name=recommender_id,json=recommenderId,proto3" json:"recommender_id"`
	Status        int32  `protobuf:"varint,17,opt,name=status,proto3" json:"status"`
	RegisteredAt  string `protobuf:"bytes,18,opt,name=registered_at,json=registeredAt,proto3" json:"registered_at"`
	CreatedAt     string `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Footmark) Reset() {
	*x = Footmark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_footmarkService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Footmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Footmark) ProtoMessage() {}

func (x *Footmark) ProtoReflect() protoreflect.Message {
	mi := &file_footmarkService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Footmark.ProtoReflect.Descriptor instead.
func (*Footmark) Descriptor() ([]byte, []int) {
	return file_footmarkService_proto_rawDescGZIP(), []int{0}
}

func (x *Footmark) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Footmark) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *Footmark) GetUnionId() string {
	if x != nil {
		return x.UnionId
	}
	return ""
}

func (x *Footmark) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Footmark) GetRecommenderId() int64 {
	if x != nil {
		return x.RecommenderId
	}
	return 0
}

func (x *Footmark) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Footmark) GetRegisteredAt() string {
	if x != nil {
		return x.RegisteredAt
	}
	return ""
}

func (x *Footmark) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Footmark) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//粉丝查询参数
type FootmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//以下为自定义参数
	Id            int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	OpenId        string `protobuf:"bytes,5,opt,name=open_id,json=openId,proto3" json:"open_id"`
	UnionId       string `protobuf:"bytes,6,opt,name=union_id,json=unionId,proto3" json:"union_id"`
	Source        string `protobuf:"bytes,7,opt,name=source,proto3" json:"source"`
	RecommenderId int64  `protobuf:"varint,8,opt,name=recommender_id,json=recommenderId,proto3" json:"recommender_id"`
	Status        int32  `protobuf:"varint,9,opt,name=status,proto3" json:"status"`
	Code          string `protobuf:"bytes,10,opt,name=code,proto3" json:"code"`
}

func (x *FootmarkRequest) Reset() {
	*x = FootmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_footmarkService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootmarkRequest) ProtoMessage() {}

func (x *FootmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_footmarkService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootmarkRequest.ProtoReflect.Descriptor instead.
func (*FootmarkRequest) Descriptor() ([]byte, []int) {
	return file_footmarkService_proto_rawDescGZIP(), []int{1}
}

func (x *FootmarkRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *FootmarkRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FootmarkRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *FootmarkRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootmarkRequest) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *FootmarkRequest) GetUnionId() string {
	if x != nil {
		return x.UnionId
	}
	return ""
}

func (x *FootmarkRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *FootmarkRequest) GetRecommenderId() int64 {
	if x != nil {
		return x.RecommenderId
	}
	return 0
}

func (x *FootmarkRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FootmarkRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

//
type FootmarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Footmark     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Footmark   `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *FootmarkResponse) Reset() {
	*x = FootmarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_footmarkService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootmarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootmarkResponse) ProtoMessage() {}

func (x *FootmarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_footmarkService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootmarkResponse.ProtoReflect.Descriptor instead.
func (*FootmarkResponse) Descriptor() ([]byte, []int) {
	return file_footmarkService_proto_rawDescGZIP(), []int{2}
}

func (x *FootmarkResponse) GetEntity() *Footmark {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *FootmarkResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *FootmarkResponse) GetItems() []*Footmark {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FootmarkResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *FootmarkResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_footmarkService_proto protoreflect.FileDescriptor

var file_footmarkService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x66, 0x6f, 0x6f, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x74, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x6e,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x8d, 0x02, 0x0a, 0x0f, 0x46, 0x6f, 0x6f, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x6e, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0xd4, 0x01, 0x0a, 0x10, 0x46, 0x6f, 0x6f, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x46, 0x6f, 0x6f, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x46, 0x6f, 0x6f, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x50, 0x0a, 0x0f, 0x46, 0x6f, 0x6f, 0x74, 0x6d, 0x61,
	0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x02, 0x57, 0x78, 0x12,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x6d,
	0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_footmarkService_proto_rawDescOnce sync.Once
	file_footmarkService_proto_rawDescData = file_footmarkService_proto_rawDesc
)

func file_footmarkService_proto_rawDescGZIP() []byte {
	file_footmarkService_proto_rawDescOnce.Do(func() {
		file_footmarkService_proto_rawDescData = protoimpl.X.CompressGZIP(file_footmarkService_proto_rawDescData)
	})
	return file_footmarkService_proto_rawDescData
}

var file_footmarkService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_footmarkService_proto_goTypes = []interface{}{
	(*Footmark)(nil),         // 0: services.Footmark
	(*FootmarkRequest)(nil),  // 1: services.FootmarkRequest
	(*FootmarkResponse)(nil), // 2: services.FootmarkResponse
	(*common.Pager)(nil),     // 3: common.Pager
	(*common.Error)(nil),     // 4: common.Error
	(*common.Info)(nil),      // 5: common.Info
}
var file_footmarkService_proto_depIdxs = []int32{
	0, // 0: services.FootmarkResponse.entity:type_name -> services.Footmark
	3, // 1: services.FootmarkResponse.pager:type_name -> common.Pager
	0, // 2: services.FootmarkResponse.items:type_name -> services.Footmark
	4, // 3: services.FootmarkResponse.error:type_name -> common.Error
	5, // 4: services.FootmarkResponse.info:type_name -> common.Info
	1, // 5: services.FootmarkService.Wx:input_type -> services.FootmarkRequest
	2, // 6: services.FootmarkService.Wx:output_type -> services.FootmarkResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_footmarkService_proto_init() }
func file_footmarkService_proto_init() {
	if File_footmarkService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_footmarkService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Footmark); i {
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
		file_footmarkService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootmarkRequest); i {
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
		file_footmarkService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootmarkResponse); i {
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
			RawDescriptor: file_footmarkService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_footmarkService_proto_goTypes,
		DependencyIndexes: file_footmarkService_proto_depIdxs,
		MessageInfos:      file_footmarkService_proto_msgTypes,
	}.Build()
	File_footmarkService_proto = out.File
	file_footmarkService_proto_rawDesc = nil
	file_footmarkService_proto_goTypes = nil
	file_footmarkService_proto_depIdxs = nil
}
