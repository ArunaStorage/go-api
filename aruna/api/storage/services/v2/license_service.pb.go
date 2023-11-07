// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: aruna/api/storage/services/v2/license_service.proto

package v2

import (
	v2 "github.com/ArunaStorage/go-api/v2/aruna/api/storage/models/v2"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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

// CreateLicenseRequest
//
// Request object for CreateLicense
type CreateLicenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag  string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`   // CC-BY-SA-4.0
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Creative Commons Attribution-ShareAlike 4.0 International
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// url is optional
	Url string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"` // https://creativecommons.org/licenses/by-sa/4.0/
}

func (x *CreateLicenseRequest) Reset() {
	*x = CreateLicenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLicenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLicenseRequest) ProtoMessage() {}

func (x *CreateLicenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLicenseRequest.ProtoReflect.Descriptor instead.
func (*CreateLicenseRequest) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_license_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLicenseRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *CreateLicenseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLicenseRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateLicenseRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// CreateLicenseResponse
//
// Response object for CreateLicense
type CreateLicenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"` // CC-BY-SA-4.0
}

func (x *CreateLicenseResponse) Reset() {
	*x = CreateLicenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLicenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLicenseResponse) ProtoMessage() {}

func (x *CreateLicenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLicenseResponse.ProtoReflect.Descriptor instead.
func (*CreateLicenseResponse) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_license_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLicenseResponse) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

// GetLicenseRequest
//
// Request object for GetLicense
type GetLicenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"` // 1234567890abcdef
}

func (x *GetLicenseRequest) Reset() {
	*x = GetLicenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLicenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLicenseRequest) ProtoMessage() {}

func (x *GetLicenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLicenseRequest.ProtoReflect.Descriptor instead.
func (*GetLicenseRequest) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_license_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLicenseRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

// GetLicenseResponse
//
// Response object for GetLicense
type GetLicenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	License *v2.License `protobuf:"bytes,1,opt,name=license,proto3" json:"license,omitempty"`
}

func (x *GetLicenseResponse) Reset() {
	*x = GetLicenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLicenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLicenseResponse) ProtoMessage() {}

func (x *GetLicenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLicenseResponse.ProtoReflect.Descriptor instead.
func (*GetLicenseResponse) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_license_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetLicenseResponse) GetLicense() *v2.License {
	if x != nil {
		return x.License
	}
	return nil
}

// ListLicensesRequest
//
// Request object for ListLicense
type ListLicensesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListLicensesRequest) Reset() {
	*x = ListLicensesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLicensesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLicensesRequest) ProtoMessage() {}

func (x *ListLicensesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLicensesRequest.ProtoReflect.Descriptor instead.
func (*ListLicensesRequest) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_license_service_proto_rawDescGZIP(), []int{4}
}

// ListLicenseResponse
//
// Response object for ListLicense
type ListLicensesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Licenses []*v2.License `protobuf:"bytes,1,rep,name=licenses,proto3" json:"licenses,omitempty"`
}

func (x *ListLicensesResponse) Reset() {
	*x = ListLicensesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLicensesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLicensesResponse) ProtoMessage() {}

func (x *ListLicensesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_license_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLicensesResponse.ProtoReflect.Descriptor instead.
func (*ListLicensesResponse) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_license_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListLicensesResponse) GetLicenses() []*v2.License {
	if x != nil {
		return x.Licenses
	}
	return nil
}

var File_aruna_api_storage_services_v2_license_service_proto protoreflect.FileDescriptor

var file_aruna_api_storage_services_v2_license_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x2f,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x32, 0x1a, 0x28, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x29, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22,
	0x54, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x32, 0xd4, 0x03, 0x0a, 0x0e, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x92, 0x01, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x2e, 0x61, 0x72,
	0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01,
	0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x8c,
	0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x2e,
	0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x32, 0x2f,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x7b, 0x74, 0x61, 0x67, 0x7d, 0x12, 0x8d, 0x01,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x32,
	0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12,
	0x0c, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x1a, 0x0e, 0xfa,
	0xd2, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x42, 0x93, 0x01,
	0x0a, 0x3e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x41, 0x72, 0x75,
	0x6e, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32,
	0x42, 0x0e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x72, 0x75, 0x6e, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aruna_api_storage_services_v2_license_service_proto_rawDescOnce sync.Once
	file_aruna_api_storage_services_v2_license_service_proto_rawDescData = file_aruna_api_storage_services_v2_license_service_proto_rawDesc
)

func file_aruna_api_storage_services_v2_license_service_proto_rawDescGZIP() []byte {
	file_aruna_api_storage_services_v2_license_service_proto_rawDescOnce.Do(func() {
		file_aruna_api_storage_services_v2_license_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_aruna_api_storage_services_v2_license_service_proto_rawDescData)
	})
	return file_aruna_api_storage_services_v2_license_service_proto_rawDescData
}

var file_aruna_api_storage_services_v2_license_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aruna_api_storage_services_v2_license_service_proto_goTypes = []interface{}{
	(*CreateLicenseRequest)(nil),  // 0: aruna.api.storage.services.v2.CreateLicenseRequest
	(*CreateLicenseResponse)(nil), // 1: aruna.api.storage.services.v2.CreateLicenseResponse
	(*GetLicenseRequest)(nil),     // 2: aruna.api.storage.services.v2.GetLicenseRequest
	(*GetLicenseResponse)(nil),    // 3: aruna.api.storage.services.v2.GetLicenseResponse
	(*ListLicensesRequest)(nil),   // 4: aruna.api.storage.services.v2.ListLicensesRequest
	(*ListLicensesResponse)(nil),  // 5: aruna.api.storage.services.v2.ListLicensesResponse
	(*v2.License)(nil),            // 6: aruna.api.storage.models.v2.License
}
var file_aruna_api_storage_services_v2_license_service_proto_depIdxs = []int32{
	6, // 0: aruna.api.storage.services.v2.GetLicenseResponse.license:type_name -> aruna.api.storage.models.v2.License
	6, // 1: aruna.api.storage.services.v2.ListLicensesResponse.licenses:type_name -> aruna.api.storage.models.v2.License
	0, // 2: aruna.api.storage.services.v2.LicenseService.CreateLicense:input_type -> aruna.api.storage.services.v2.CreateLicenseRequest
	2, // 3: aruna.api.storage.services.v2.LicenseService.GetLicense:input_type -> aruna.api.storage.services.v2.GetLicenseRequest
	4, // 4: aruna.api.storage.services.v2.LicenseService.ListLicenses:input_type -> aruna.api.storage.services.v2.ListLicensesRequest
	1, // 5: aruna.api.storage.services.v2.LicenseService.CreateLicense:output_type -> aruna.api.storage.services.v2.CreateLicenseResponse
	3, // 6: aruna.api.storage.services.v2.LicenseService.GetLicense:output_type -> aruna.api.storage.services.v2.GetLicenseResponse
	5, // 7: aruna.api.storage.services.v2.LicenseService.ListLicenses:output_type -> aruna.api.storage.services.v2.ListLicensesResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_aruna_api_storage_services_v2_license_service_proto_init() }
func file_aruna_api_storage_services_v2_license_service_proto_init() {
	if File_aruna_api_storage_services_v2_license_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aruna_api_storage_services_v2_license_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLicenseRequest); i {
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
		file_aruna_api_storage_services_v2_license_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLicenseResponse); i {
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
		file_aruna_api_storage_services_v2_license_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLicenseRequest); i {
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
		file_aruna_api_storage_services_v2_license_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLicenseResponse); i {
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
		file_aruna_api_storage_services_v2_license_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLicensesRequest); i {
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
		file_aruna_api_storage_services_v2_license_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLicensesResponse); i {
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
			RawDescriptor: file_aruna_api_storage_services_v2_license_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aruna_api_storage_services_v2_license_service_proto_goTypes,
		DependencyIndexes: file_aruna_api_storage_services_v2_license_service_proto_depIdxs,
		MessageInfos:      file_aruna_api_storage_services_v2_license_service_proto_msgTypes,
	}.Build()
	File_aruna_api_storage_services_v2_license_service_proto = out.File
	file_aruna_api_storage_services_v2_license_service_proto_rawDesc = nil
	file_aruna_api_storage_services_v2_license_service_proto_goTypes = nil
	file_aruna_api_storage_services_v2_license_service_proto_depIdxs = nil
}