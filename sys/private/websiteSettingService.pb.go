// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: websiteSettingService.proto

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

// 网站设置
type WebsiteSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title"`
	SiteUrl     string `protobuf:"bytes,2,opt,name=site_url,json=siteUrl,proto3" json:"site_url"`
	LogoId      int64  `protobuf:"varint,3,opt,name=logo_id,json=logoId,proto3" json:"logo_id"`
	LogoUrl     string `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url"`
	Keywords    string `protobuf:"bytes,5,opt,name=keywords,proto3" json:"keywords"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description"`
	Hotline     string `protobuf:"bytes,7,opt,name=hotline,proto3" json:"hotline"`
	CompanyName string `protobuf:"bytes,8,opt,name=company_name,json=companyName,proto3" json:"company_name"`
	AreaId      int64  `protobuf:"varint,9,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	Address     string `protobuf:"bytes,10,opt,name=address,proto3" json:"address"`
	Lng         string `protobuf:"bytes,11,opt,name=lng,proto3" json:"lng"`
	Lat         string `protobuf:"bytes,12,opt,name=lat,proto3" json:"lat"`
	Email       string `protobuf:"bytes,13,opt,name=email,proto3" json:"email"`
	Qq          string `protobuf:"bytes,14,opt,name=qq,proto3" json:"qq"`
	Icp         string `protobuf:"bytes,15,opt,name=icp,proto3" json:"icp"`
	PoliceIcp   string `protobuf:"bytes,16,opt,name=police_icp,json=policeIcp,proto3" json:"police_icp"`
	Copyright   string `protobuf:"bytes,17,opt,name=copyright,proto3" json:"copyright"`
	ContactMan  string `protobuf:"bytes,18,opt,name=contact_man,json=contactMan,proto3" json:"contact_man"`
	ContactTel  string `protobuf:"bytes,19,opt,name=contact_tel,json=contactTel,proto3" json:"contact_tel"`
}

func (x *WebsiteSetting) Reset() {
	*x = WebsiteSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websiteSettingService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebsiteSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebsiteSetting) ProtoMessage() {}

func (x *WebsiteSetting) ProtoReflect() protoreflect.Message {
	mi := &file_websiteSettingService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebsiteSetting.ProtoReflect.Descriptor instead.
func (*WebsiteSetting) Descriptor() ([]byte, []int) {
	return file_websiteSettingService_proto_rawDescGZIP(), []int{0}
}

func (x *WebsiteSetting) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *WebsiteSetting) GetSiteUrl() string {
	if x != nil {
		return x.SiteUrl
	}
	return ""
}

func (x *WebsiteSetting) GetLogoId() int64 {
	if x != nil {
		return x.LogoId
	}
	return 0
}

func (x *WebsiteSetting) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *WebsiteSetting) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *WebsiteSetting) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WebsiteSetting) GetHotline() string {
	if x != nil {
		return x.Hotline
	}
	return ""
}

func (x *WebsiteSetting) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *WebsiteSetting) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *WebsiteSetting) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *WebsiteSetting) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *WebsiteSetting) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *WebsiteSetting) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *WebsiteSetting) GetQq() string {
	if x != nil {
		return x.Qq
	}
	return ""
}

func (x *WebsiteSetting) GetIcp() string {
	if x != nil {
		return x.Icp
	}
	return ""
}

func (x *WebsiteSetting) GetPoliceIcp() string {
	if x != nil {
		return x.PoliceIcp
	}
	return ""
}

func (x *WebsiteSetting) GetCopyright() string {
	if x != nil {
		return x.Copyright
	}
	return ""
}

func (x *WebsiteSetting) GetContactMan() string {
	if x != nil {
		return x.ContactMan
	}
	return ""
}

func (x *WebsiteSetting) GetContactTel() string {
	if x != nil {
		return x.ContactTel
	}
	return ""
}

//
type WebsiteSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *WebsiteSetting `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Error  *common.Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *WebsiteSettingResponse) Reset() {
	*x = WebsiteSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websiteSettingService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebsiteSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebsiteSettingResponse) ProtoMessage() {}

func (x *WebsiteSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_websiteSettingService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebsiteSettingResponse.ProtoReflect.Descriptor instead.
func (*WebsiteSettingResponse) Descriptor() ([]byte, []int) {
	return file_websiteSettingService_proto_rawDescGZIP(), []int{1}
}

func (x *WebsiteSettingResponse) GetEntity() *WebsiteSetting {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *WebsiteSettingResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *WebsiteSettingResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_websiteSettingService_proto protoreflect.FileDescriptor

var file_websiteSettingService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x03, 0x0a, 0x0e, 0x57,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x17,
	0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6c, 0x6f, 0x67, 0x6f, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55,
	0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c,
	0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6c, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x71, 0x71,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x71, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x63,
	0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x63, 0x70, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x63, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x49, 0x63, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x65, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x65, 0x6c, 0x22, 0x91, 0x01, 0x0a, 0x16,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32,
	0x96, 0x01, 0x0a, 0x15, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x03, 0x53, 0x65, 0x74,
	0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_websiteSettingService_proto_rawDescOnce sync.Once
	file_websiteSettingService_proto_rawDescData = file_websiteSettingService_proto_rawDesc
)

func file_websiteSettingService_proto_rawDescGZIP() []byte {
	file_websiteSettingService_proto_rawDescOnce.Do(func() {
		file_websiteSettingService_proto_rawDescData = protoimpl.X.CompressGZIP(file_websiteSettingService_proto_rawDescData)
	})
	return file_websiteSettingService_proto_rawDescData
}

var file_websiteSettingService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_websiteSettingService_proto_goTypes = []interface{}{
	(*WebsiteSetting)(nil),         // 0: services.WebsiteSetting
	(*WebsiteSettingResponse)(nil), // 1: services.WebsiteSettingResponse
	(*common.Error)(nil),           // 2: common.Error
	(*common.Info)(nil),            // 3: common.Info
	(*common.Empty)(nil),           // 4: common.Empty
}
var file_websiteSettingService_proto_depIdxs = []int32{
	0, // 0: services.WebsiteSettingResponse.entity:type_name -> services.WebsiteSetting
	2, // 1: services.WebsiteSettingResponse.error:type_name -> common.Error
	3, // 2: services.WebsiteSettingResponse.info:type_name -> common.Info
	0, // 3: services.WebsiteSettingService.Set:input_type -> services.WebsiteSetting
	4, // 4: services.WebsiteSettingService.Get:input_type -> common.Empty
	1, // 5: services.WebsiteSettingService.Set:output_type -> services.WebsiteSettingResponse
	1, // 6: services.WebsiteSettingService.Get:output_type -> services.WebsiteSettingResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_websiteSettingService_proto_init() }
func file_websiteSettingService_proto_init() {
	if File_websiteSettingService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_websiteSettingService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebsiteSetting); i {
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
		file_websiteSettingService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebsiteSettingResponse); i {
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
			RawDescriptor: file_websiteSettingService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_websiteSettingService_proto_goTypes,
		DependencyIndexes: file_websiteSettingService_proto_depIdxs,
		MessageInfos:      file_websiteSettingService_proto_msgTypes,
	}.Build()
	File_websiteSettingService_proto = out.File
	file_websiteSettingService_proto_rawDesc = nil
	file_websiteSettingService_proto_goTypes = nil
	file_websiteSettingService_proto_depIdxs = nil
}
