// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: vipService.proto

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

type VipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	//base params
	Id         int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name       string `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	ItemSn     string `protobuf:"bytes,6,opt,name=item_sn,json=itemSn,proto3" json:"item_sn"`
	VipLevelId int32  `protobuf:"varint,7,opt,name=vip_level_id,json=vipLevelId,proto3" json:"vip_level_id"`
}

func (x *VipRequest) Reset() {
	*x = VipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipRequest) ProtoMessage() {}

func (x *VipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vipService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipRequest.ProtoReflect.Descriptor instead.
func (*VipRequest) Descriptor() ([]byte, []int) {
	return file_vipService_proto_rawDescGZIP(), []int{0}
}

func (x *VipRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *VipRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *VipRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *VipRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VipRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VipRequest) GetItemSn() string {
	if x != nil {
		return x.ItemSn
	}
	return ""
}

func (x *VipRequest) GetVipLevelId() int32 {
	if x != nil {
		return x.VipLevelId
	}
	return 0
}

type VipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Item         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Item       `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *VipResponse) Reset() {
	*x = VipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipResponse) ProtoMessage() {}

func (x *VipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vipService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipResponse.ProtoReflect.Descriptor instead.
func (*VipResponse) Descriptor() ([]byte, []int) {
	return file_vipService_proto_rawDescGZIP(), []int{1}
}

func (x *VipResponse) GetEntity() *Item {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *VipResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *VipResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *VipResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *VipResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_vipService_proto protoreflect.FileDescriptor

var file_vipService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x0a, 0x56, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x73, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x53, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x76, 0x69, 0x70, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x76, 0x69, 0x70, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x0b, 0x56, 0x69, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32,
	0x82, 0x03, 0x0a, 0x0a, 0x56, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x56, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x56, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vipService_proto_rawDescOnce sync.Once
	file_vipService_proto_rawDescData = file_vipService_proto_rawDesc
)

func file_vipService_proto_rawDescGZIP() []byte {
	file_vipService_proto_rawDescOnce.Do(func() {
		file_vipService_proto_rawDescData = protoimpl.X.CompressGZIP(file_vipService_proto_rawDescData)
	})
	return file_vipService_proto_rawDescData
}

var file_vipService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vipService_proto_goTypes = []interface{}{
	(*VipRequest)(nil),   // 0: services.VipRequest
	(*VipResponse)(nil),  // 1: services.VipResponse
	(*Item)(nil),         // 2: services.Item
	(*common.Pager)(nil), // 3: common.Pager
	(*common.Error)(nil), // 4: common.Error
	(*common.Info)(nil),  // 5: common.Info
	(*ItemResponse)(nil), // 6: services.ItemResponse
}
var file_vipService_proto_depIdxs = []int32{
	2,  // 0: services.VipResponse.entity:type_name -> services.Item
	3,  // 1: services.VipResponse.pager:type_name -> common.Pager
	2,  // 2: services.VipResponse.items:type_name -> services.Item
	4,  // 3: services.VipResponse.error:type_name -> common.Error
	5,  // 4: services.VipResponse.info:type_name -> common.Info
	2,  // 5: services.VipService.Create:input_type -> services.Item
	2,  // 6: services.VipService.Update:input_type -> services.Item
	2,  // 7: services.VipService.Delete:input_type -> services.Item
	2,  // 8: services.VipService.Get:input_type -> services.Item
	2,  // 9: services.VipService.Display:input_type -> services.Item
	0,  // 10: services.VipService.Search:input_type -> services.VipRequest
	0,  // 11: services.VipService.FrontSearch:input_type -> services.VipRequest
	1,  // 12: services.VipService.Create:output_type -> services.VipResponse
	1,  // 13: services.VipService.Update:output_type -> services.VipResponse
	1,  // 14: services.VipService.Delete:output_type -> services.VipResponse
	1,  // 15: services.VipService.Get:output_type -> services.VipResponse
	1,  // 16: services.VipService.Display:output_type -> services.VipResponse
	6,  // 17: services.VipService.Search:output_type -> services.ItemResponse
	6,  // 18: services.VipService.FrontSearch:output_type -> services.ItemResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_vipService_proto_init() }
func file_vipService_proto_init() {
	if File_vipService_proto != nil {
		return
	}
	file_itemService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vipService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipRequest); i {
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
		file_vipService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipResponse); i {
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
			RawDescriptor: file_vipService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vipService_proto_goTypes,
		DependencyIndexes: file_vipService_proto_depIdxs,
		MessageInfos:      file_vipService_proto_msgTypes,
	}.Build()
	File_vipService_proto = out.File
	file_vipService_proto_rawDesc = nil
	file_vipService_proto_goTypes = nil
	file_vipService_proto_depIdxs = nil
}
