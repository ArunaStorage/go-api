// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: aruna/api/dataproxy/services/v2/bundler_service.proto

package v2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBundleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId []string               `protobuf:"bytes,1,rep,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Filename   string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`                    // .tar.gz / .zip
	ExpiresAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"` // Default 1 Month
}

func (x *CreateBundleRequest) Reset() {
	*x = CreateBundleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBundleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBundleRequest) ProtoMessage() {}

func (x *CreateBundleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBundleRequest.ProtoReflect.Descriptor instead.
func (*CreateBundleRequest) Descriptor() ([]byte, []int) {
	return file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBundleRequest) GetResourceId() []string {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

func (x *CreateBundleRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *CreateBundleRequest) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

type CreateBundleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BundleId  string `protobuf:"bytes,1,opt,name=bundle_id,json=bundleId,proto3" json:"bundle_id,omitempty"`
	BundleUrl string `protobuf:"bytes,2,opt,name=bundle_url,json=bundleUrl,proto3" json:"bundle_url,omitempty"`
}

func (x *CreateBundleResponse) Reset() {
	*x = CreateBundleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBundleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBundleResponse) ProtoMessage() {}

func (x *CreateBundleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBundleResponse.ProtoReflect.Descriptor instead.
func (*CreateBundleResponse) Descriptor() ([]byte, []int) {
	return file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBundleResponse) GetBundleId() string {
	if x != nil {
		return x.BundleId
	}
	return ""
}

func (x *CreateBundleResponse) GetBundleUrl() string {
	if x != nil {
		return x.BundleUrl
	}
	return ""
}

type DeleteBundleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BundleId string `protobuf:"bytes,1,opt,name=bundle_id,json=bundleId,proto3" json:"bundle_id,omitempty"`
}

func (x *DeleteBundleRequest) Reset() {
	*x = DeleteBundleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBundleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBundleRequest) ProtoMessage() {}

func (x *DeleteBundleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBundleRequest.ProtoReflect.Descriptor instead.
func (*DeleteBundleRequest) Descriptor() ([]byte, []int) {
	return file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBundleRequest) GetBundleId() string {
	if x != nil {
		return x.BundleId
	}
	return ""
}

type DeleteBundleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBundleResponse) Reset() {
	*x = DeleteBundleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBundleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBundleResponse) ProtoMessage() {}

func (x *DeleteBundleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBundleResponse.ProtoReflect.Descriptor instead.
func (*DeleteBundleResponse) Descriptor() ([]byte, []int) {
	return file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescGZIP(), []int{3}
}

var File_aruna_api_dataproxy_services_v2_bundler_service_proto protoreflect.FileDescriptor

var file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x41, 0x74, 0x22, 0x52, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd7, 0x02, 0x0a, 0x0e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x34, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35,
	0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a,
	0x22, 0x0b, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x9f, 0x01,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x34,
	0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x2a, 0x17, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x1a,
	0x0d, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x42, 0x9b,
	0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x13, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x75, 0x6e, 0x61, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72,
	0x75, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x32, 0xa2, 0x02, 0x04,
	0x41, 0x41, 0x44, 0x53, 0xaa, 0x02, 0x1f, 0x41, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x41, 0x70, 0x69,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x1f, 0x41, 0x72, 0x75, 0x6e, 0x61, 0x5c, 0x41,
	0x70, 0x69, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x2b, 0x41, 0x72, 0x75, 0x6e, 0x61,
	0x5c, 0x41, 0x70, 0x69, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x23, 0x41, 0x72, 0x75, 0x6e, 0x61, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x3a, 0x3a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescOnce sync.Once
	file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescData = file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDesc
)

func file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescGZIP() []byte {
	file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescOnce.Do(func() {
		file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescData)
	})
	return file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDescData
}

var file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aruna_api_dataproxy_services_v2_bundler_service_proto_goTypes = []interface{}{
	(*CreateBundleRequest)(nil),   // 0: aruna.api.dataproxy.services.v2.CreateBundleRequest
	(*CreateBundleResponse)(nil),  // 1: aruna.api.dataproxy.services.v2.CreateBundleResponse
	(*DeleteBundleRequest)(nil),   // 2: aruna.api.dataproxy.services.v2.DeleteBundleRequest
	(*DeleteBundleResponse)(nil),  // 3: aruna.api.dataproxy.services.v2.DeleteBundleResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_aruna_api_dataproxy_services_v2_bundler_service_proto_depIdxs = []int32{
	4, // 0: aruna.api.dataproxy.services.v2.CreateBundleRequest.expires_at:type_name -> google.protobuf.Timestamp
	0, // 1: aruna.api.dataproxy.services.v2.BundlerService.CreateBundle:input_type -> aruna.api.dataproxy.services.v2.CreateBundleRequest
	2, // 2: aruna.api.dataproxy.services.v2.BundlerService.DeleteBundle:input_type -> aruna.api.dataproxy.services.v2.DeleteBundleRequest
	1, // 3: aruna.api.dataproxy.services.v2.BundlerService.CreateBundle:output_type -> aruna.api.dataproxy.services.v2.CreateBundleResponse
	3, // 4: aruna.api.dataproxy.services.v2.BundlerService.DeleteBundle:output_type -> aruna.api.dataproxy.services.v2.DeleteBundleResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aruna_api_dataproxy_services_v2_bundler_service_proto_init() }
func file_aruna_api_dataproxy_services_v2_bundler_service_proto_init() {
	if File_aruna_api_dataproxy_services_v2_bundler_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBundleRequest); i {
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
		file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBundleResponse); i {
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
		file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBundleRequest); i {
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
		file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBundleResponse); i {
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
			RawDescriptor: file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aruna_api_dataproxy_services_v2_bundler_service_proto_goTypes,
		DependencyIndexes: file_aruna_api_dataproxy_services_v2_bundler_service_proto_depIdxs,
		MessageInfos:      file_aruna_api_dataproxy_services_v2_bundler_service_proto_msgTypes,
	}.Build()
	File_aruna_api_dataproxy_services_v2_bundler_service_proto = out.File
	file_aruna_api_dataproxy_services_v2_bundler_service_proto_rawDesc = nil
	file_aruna_api_dataproxy_services_v2_bundler_service_proto_goTypes = nil
	file_aruna_api_dataproxy_services_v2_bundler_service_proto_depIdxs = nil
}
