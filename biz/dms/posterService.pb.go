// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: posterService.proto

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

type Poster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalCfgId        int64            `protobuf:"varint,1,opt,name=global_cfg_id,json=globalCfgId,proto3" json:"global_cfg_id"`
	IsOpenPoster       bool             `protobuf:"varint,2,opt,name=is_open_poster,json=isOpenPoster,proto3" json:"is_open_poster"`
	PosterShareUrl     string           `protobuf:"bytes,3,opt,name=poster_share_url,json=posterShareUrl,proto3" json:"poster_share_url"`
	PosterType         int32            `protobuf:"varint,4,opt,name=poster_type,json=posterType,proto3" json:"poster_type"`
	PosterImgId        int64            `protobuf:"varint,5,opt,name=poster_img_id,json=posterImgId,proto3" json:"poster_img_id"`
	PosterImgUrl       string           `protobuf:"bytes,6,opt,name=poster_img_url,json=posterImgUrl,proto3" json:"poster_img_url"`
	PosterShareContent string           `protobuf:"bytes,7,opt,name=poster_share_content,json=posterShareContent,proto3" json:"poster_share_content"`
	PosterGalleries    []*PosterGallery `protobuf:"bytes,24,rep,name=poster_galleries,json=posterGalleries,proto3" json:"poster_galleries"`
}

func (x *Poster) Reset() {
	*x = Poster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posterService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Poster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Poster) ProtoMessage() {}

func (x *Poster) ProtoReflect() protoreflect.Message {
	mi := &file_posterService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Poster.ProtoReflect.Descriptor instead.
func (*Poster) Descriptor() ([]byte, []int) {
	return file_posterService_proto_rawDescGZIP(), []int{0}
}

func (x *Poster) GetGlobalCfgId() int64 {
	if x != nil {
		return x.GlobalCfgId
	}
	return 0
}

func (x *Poster) GetIsOpenPoster() bool {
	if x != nil {
		return x.IsOpenPoster
	}
	return false
}

func (x *Poster) GetPosterShareUrl() string {
	if x != nil {
		return x.PosterShareUrl
	}
	return ""
}

func (x *Poster) GetPosterType() int32 {
	if x != nil {
		return x.PosterType
	}
	return 0
}

func (x *Poster) GetPosterImgId() int64 {
	if x != nil {
		return x.PosterImgId
	}
	return 0
}

func (x *Poster) GetPosterImgUrl() string {
	if x != nil {
		return x.PosterImgUrl
	}
	return ""
}

func (x *Poster) GetPosterShareContent() string {
	if x != nil {
		return x.PosterShareContent
	}
	return ""
}

func (x *Poster) GetPosterGalleries() []*PosterGallery {
	if x != nil {
		return x.PosterGalleries
	}
	return nil
}

type PosterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Poster       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *PosterResponse) Reset() {
	*x = PosterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posterService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosterResponse) ProtoMessage() {}

func (x *PosterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_posterService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosterResponse.ProtoReflect.Descriptor instead.
func (*PosterResponse) Descriptor() ([]byte, []int) {
	return file_posterService_proto_rawDescGZIP(), []int{1}
}

func (x *PosterResponse) GetEntity() *Poster {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PosterResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PosterResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_posterService_proto protoreflect.FileDescriptor

var file_posterService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd,
	0x02, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x63, 0x66, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x66, 0x67, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22,
	0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6d, 0x67,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x67,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x10, 0x70, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x18,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x0f, 0x70,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x81,
	0x01, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x32, 0x7a, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65,
	0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x49,
	0x0a, 0x12, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_posterService_proto_rawDescOnce sync.Once
	file_posterService_proto_rawDescData = file_posterService_proto_rawDesc
)

func file_posterService_proto_rawDescGZIP() []byte {
	file_posterService_proto_rawDescOnce.Do(func() {
		file_posterService_proto_rawDescData = protoimpl.X.CompressGZIP(file_posterService_proto_rawDescData)
	})
	return file_posterService_proto_rawDescData
}

var file_posterService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_posterService_proto_goTypes = []interface{}{
	(*Poster)(nil),         // 0: services.Poster
	(*PosterResponse)(nil), // 1: services.PosterResponse
	(*PosterGallery)(nil),  // 2: services.PosterGallery
	(*common.Error)(nil),   // 3: common.Error
	(*common.Info)(nil),    // 4: common.Info
}
var file_posterService_proto_depIdxs = []int32{
	2, // 0: services.Poster.poster_galleries:type_name -> services.PosterGallery
	0, // 1: services.PosterResponse.entity:type_name -> services.Poster
	3, // 2: services.PosterResponse.error:type_name -> common.Error
	4, // 3: services.PosterResponse.info:type_name -> common.Info
	0, // 4: services.PosterService.Get:input_type -> services.Poster
	0, // 5: services.PosterService.Save:input_type -> services.Poster
	0, // 6: services.FrontPosterService.Get:input_type -> services.Poster
	1, // 7: services.PosterService.Get:output_type -> services.PosterResponse
	1, // 8: services.PosterService.Save:output_type -> services.PosterResponse
	1, // 9: services.FrontPosterService.Get:output_type -> services.PosterResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_posterService_proto_init() }
func file_posterService_proto_init() {
	if File_posterService_proto != nil {
		return
	}
	file_posterGalleryService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_posterService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Poster); i {
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
		file_posterService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosterResponse); i {
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
			RawDescriptor: file_posterService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_posterService_proto_goTypes,
		DependencyIndexes: file_posterService_proto_depIdxs,
		MessageInfos:      file_posterService_proto_msgTypes,
	}.Build()
	File_posterService_proto = out.File
	file_posterService_proto_rawDesc = nil
	file_posterService_proto_goTypes = nil
	file_posterService_proto_depIdxs = nil
}
