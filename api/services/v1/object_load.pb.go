// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/services/v1/object_load.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_services_v1_object_load_proto protoreflect.FileDescriptor

var file_api_services_v1_object_load_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc4, 0x09, 0x0a,
	0x11, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x97, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x6c, 0x6f, 0x61, 0x64,
	0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xaa,
	0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0xb6, 0x01, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69,
	0x6e, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x30, 0x01, 0x12, 0xa3, 0x01, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x28, 0x22, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xc5, 0x01, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x44, 0x12, 0x42,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x6c, 0x6f,
	0x61, 0x64, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x61, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x2f, 0x7b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x7d, 0x12, 0xae, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61,
	0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x6c, 0x6f, 0x61, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x61, 0x72, 0x74, 0x42, 0x80, 0x01, 0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x44, 0x42, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2f,
	0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_services_v1_object_load_proto_goTypes = []interface{}{
	(*CreateUploadLinkRequest)(nil),          // 0: api.services.v1.CreateUploadLinkRequest
	(*CreateDownloadLinkRequest)(nil),        // 1: api.services.v1.CreateDownloadLinkRequest
	(*CreateDownloadLinkBatchRequest)(nil),   // 2: api.services.v1.CreateDownloadLinkBatchRequest
	(*CreateDownloadLinkStreamRequest)(nil),  // 3: api.services.v1.CreateDownloadLinkStreamRequest
	(*StartMultipartUploadRequest)(nil),      // 4: api.services.v1.StartMultipartUploadRequest
	(*GetMultipartUploadLinkRequest)(nil),    // 5: api.services.v1.GetMultipartUploadLinkRequest
	(*CompleteMultipartUploadRequest)(nil),   // 6: api.services.v1.CompleteMultipartUploadRequest
	(*CreateUploadLinkResponse)(nil),         // 7: api.services.v1.CreateUploadLinkResponse
	(*CreateDownloadLinkResponse)(nil),       // 8: api.services.v1.CreateDownloadLinkResponse
	(*CreateDownloadLinkBatchResponse)(nil),  // 9: api.services.v1.CreateDownloadLinkBatchResponse
	(*CreateDownloadLinkStreamResponse)(nil), // 10: api.services.v1.CreateDownloadLinkStreamResponse
	(*StartMultipartUploadResponse)(nil),     // 11: api.services.v1.StartMultipartUploadResponse
	(*GetMultipartUploadLinkResponse)(nil),   // 12: api.services.v1.GetMultipartUploadLinkResponse
	(*CompleteMultipartUploadResponse)(nil),  // 13: api.services.v1.CompleteMultipartUploadResponse
}
var file_api_services_v1_object_load_proto_depIdxs = []int32{
	0,  // 0: api.services.v1.ObjectLoadService.CreateUploadLink:input_type -> api.services.v1.CreateUploadLinkRequest
	1,  // 1: api.services.v1.ObjectLoadService.CreateDownloadLink:input_type -> api.services.v1.CreateDownloadLinkRequest
	2,  // 2: api.services.v1.ObjectLoadService.CreateDownloadLinkBatch:input_type -> api.services.v1.CreateDownloadLinkBatchRequest
	3,  // 3: api.services.v1.ObjectLoadService.CreateDownloadLinkStream:input_type -> api.services.v1.CreateDownloadLinkStreamRequest
	4,  // 4: api.services.v1.ObjectLoadService.StartMultipartUpload:input_type -> api.services.v1.StartMultipartUploadRequest
	5,  // 5: api.services.v1.ObjectLoadService.GetMultipartUploadLink:input_type -> api.services.v1.GetMultipartUploadLinkRequest
	6,  // 6: api.services.v1.ObjectLoadService.CompleteMultipartUpload:input_type -> api.services.v1.CompleteMultipartUploadRequest
	7,  // 7: api.services.v1.ObjectLoadService.CreateUploadLink:output_type -> api.services.v1.CreateUploadLinkResponse
	8,  // 8: api.services.v1.ObjectLoadService.CreateDownloadLink:output_type -> api.services.v1.CreateDownloadLinkResponse
	9,  // 9: api.services.v1.ObjectLoadService.CreateDownloadLinkBatch:output_type -> api.services.v1.CreateDownloadLinkBatchResponse
	10, // 10: api.services.v1.ObjectLoadService.CreateDownloadLinkStream:output_type -> api.services.v1.CreateDownloadLinkStreamResponse
	11, // 11: api.services.v1.ObjectLoadService.StartMultipartUpload:output_type -> api.services.v1.StartMultipartUploadResponse
	12, // 12: api.services.v1.ObjectLoadService.GetMultipartUploadLink:output_type -> api.services.v1.GetMultipartUploadLinkResponse
	13, // 13: api.services.v1.ObjectLoadService.CompleteMultipartUpload:output_type -> api.services.v1.CompleteMultipartUploadResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_services_v1_object_load_proto_init() }
func file_api_services_v1_object_load_proto_init() {
	if File_api_services_v1_object_load_proto != nil {
		return
	}
	file_api_services_v1_object_load_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_services_v1_object_load_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_services_v1_object_load_proto_goTypes,
		DependencyIndexes: file_api_services_v1_object_load_proto_depIdxs,
	}.Build()
	File_api_services_v1_object_load_proto = out.File
	file_api_services_v1_object_load_proto_rawDesc = nil
	file_api_services_v1_object_load_proto_goTypes = nil
	file_api_services_v1_object_load_proto_depIdxs = nil
}
