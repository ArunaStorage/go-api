// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: sciobjsdb/api/storage/services/v1/dataset_service.proto

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

var File_sciobjsdb_api_storage_services_v1_dataset_service_proto protoreflect.FileDescriptor

var file_sciobjsdb_api_storage_services_v1_dataset_service_proto_rawDesc = []byte{
	0x0a, 0x37, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x73, 0x63, 0x69, 0x6f, 0x62,
	0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x3e, 0x73, 0x63,
	0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8d, 0x12, 0x0a, 0x0e, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa5, 0x01,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12,
	0x37, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62,
	0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x12, 0x34, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x73, 0x63, 0x69,
	0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x67, 0x65,
	0x74, 0x12, 0xba, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62,
	0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73,
	0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a,
	0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xbe,
	0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x40, 0x2e, 0x73, 0x63, 0x69, 0x6f,
	0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x73, 0x63,
	0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0xcd, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x43, 0x2e,
	0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x44, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0xb4, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3c, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73,
	0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x37, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62,
	0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x38, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xf3, 0x01, 0x0a, 0x22, 0x47, 0x65,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x4c, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4d,
	0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x69, 0x6e,
	0x64, 0x61, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0xc4, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x2e, 0x73, 0x63, 0x69, 0x6f,
	0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x73, 0x63, 0x69,
	0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xb5, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x2e, 0x73,
	0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x73, 0x63, 0x69, 0x6f,
	0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a,
	0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x12, 0xda,
	0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x47, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x73, 0x63, 0x69, 0x6f,
	0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xbc, 0x01, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xa1, 0x01, 0x0a, 0x46, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x44, 0x42, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x69,
	0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_sciobjsdb_api_storage_services_v1_dataset_service_proto_goTypes = []interface{}{
	(*CreateDatasetRequest)(nil),                       // 0: sciobjsdb.api.storage.services.v1.CreateDatasetRequest
	(*GetDatasetRequest)(nil),                          // 1: sciobjsdb.api.storage.services.v1.GetDatasetRequest
	(*GetDatasetVersionsRequest)(nil),                  // 2: sciobjsdb.api.storage.services.v1.GetDatasetVersionsRequest
	(*GetDatasetObjectGroupsRequest)(nil),              // 3: sciobjsdb.api.storage.services.v1.GetDatasetObjectGroupsRequest
	(*GetObjectGroupsStreamLinkRequest)(nil),           // 4: sciobjsdb.api.storage.services.v1.GetObjectGroupsStreamLinkRequest
	(*UpdateDatasetFieldRequest)(nil),                  // 5: sciobjsdb.api.storage.services.v1.UpdateDatasetFieldRequest
	(*DeleteDatasetRequest)(nil),                       // 6: sciobjsdb.api.storage.services.v1.DeleteDatasetRequest
	(*GetObjectGroupRevisionsInDateRangeRequest)(nil),  // 7: sciobjsdb.api.storage.services.v1.GetObjectGroupRevisionsInDateRangeRequest
	(*ReleaseDatasetVersionRequest)(nil),               // 8: sciobjsdb.api.storage.services.v1.ReleaseDatasetVersionRequest
	(*GetDatasetVersionRequest)(nil),                   // 9: sciobjsdb.api.storage.services.v1.GetDatasetVersionRequest
	(*GetDatasetVersionObjectGroupsRequest)(nil),       // 10: sciobjsdb.api.storage.services.v1.GetDatasetVersionObjectGroupsRequest
	(*DeleteDatasetVersionRequest)(nil),                // 11: sciobjsdb.api.storage.services.v1.DeleteDatasetVersionRequest
	(*CreateDatasetResponse)(nil),                      // 12: sciobjsdb.api.storage.services.v1.CreateDatasetResponse
	(*GetDatasetResponse)(nil),                         // 13: sciobjsdb.api.storage.services.v1.GetDatasetResponse
	(*GetDatasetVersionsResponse)(nil),                 // 14: sciobjsdb.api.storage.services.v1.GetDatasetVersionsResponse
	(*GetDatasetObjectGroupsResponse)(nil),             // 15: sciobjsdb.api.storage.services.v1.GetDatasetObjectGroupsResponse
	(*GetObjectGroupsStreamLinkResponse)(nil),          // 16: sciobjsdb.api.storage.services.v1.GetObjectGroupsStreamLinkResponse
	(*UpdateDatasetFieldResponse)(nil),                 // 17: sciobjsdb.api.storage.services.v1.UpdateDatasetFieldResponse
	(*DeleteDatasetResponse)(nil),                      // 18: sciobjsdb.api.storage.services.v1.DeleteDatasetResponse
	(*GetObjectGroupRevisionsInDateRangeResponse)(nil), // 19: sciobjsdb.api.storage.services.v1.GetObjectGroupRevisionsInDateRangeResponse
	(*ReleaseDatasetVersionResponse)(nil),              // 20: sciobjsdb.api.storage.services.v1.ReleaseDatasetVersionResponse
	(*GetDatasetVersionResponse)(nil),                  // 21: sciobjsdb.api.storage.services.v1.GetDatasetVersionResponse
	(*GetDatasetVersionObjectGroupsResponse)(nil),      // 22: sciobjsdb.api.storage.services.v1.GetDatasetVersionObjectGroupsResponse
	(*DeleteDatasetVersionResponse)(nil),               // 23: sciobjsdb.api.storage.services.v1.DeleteDatasetVersionResponse
}
var file_sciobjsdb_api_storage_services_v1_dataset_service_proto_depIdxs = []int32{
	0,  // 0: sciobjsdb.api.storage.services.v1.DatasetService.CreateDataset:input_type -> sciobjsdb.api.storage.services.v1.CreateDatasetRequest
	1,  // 1: sciobjsdb.api.storage.services.v1.DatasetService.GetDataset:input_type -> sciobjsdb.api.storage.services.v1.GetDatasetRequest
	2,  // 2: sciobjsdb.api.storage.services.v1.DatasetService.GetDatasetVersions:input_type -> sciobjsdb.api.storage.services.v1.GetDatasetVersionsRequest
	3,  // 3: sciobjsdb.api.storage.services.v1.DatasetService.GetDatasetObjectGroups:input_type -> sciobjsdb.api.storage.services.v1.GetDatasetObjectGroupsRequest
	4,  // 4: sciobjsdb.api.storage.services.v1.DatasetService.GetObjectGroupsStreamLink:input_type -> sciobjsdb.api.storage.services.v1.GetObjectGroupsStreamLinkRequest
	5,  // 5: sciobjsdb.api.storage.services.v1.DatasetService.UpdateDatasetField:input_type -> sciobjsdb.api.storage.services.v1.UpdateDatasetFieldRequest
	6,  // 6: sciobjsdb.api.storage.services.v1.DatasetService.DeleteDataset:input_type -> sciobjsdb.api.storage.services.v1.DeleteDatasetRequest
	7,  // 7: sciobjsdb.api.storage.services.v1.DatasetService.GetObjectGroupRevisionsInDateRange:input_type -> sciobjsdb.api.storage.services.v1.GetObjectGroupRevisionsInDateRangeRequest
	8,  // 8: sciobjsdb.api.storage.services.v1.DatasetService.ReleaseDatasetVersion:input_type -> sciobjsdb.api.storage.services.v1.ReleaseDatasetVersionRequest
	9,  // 9: sciobjsdb.api.storage.services.v1.DatasetService.GetDatasetVersion:input_type -> sciobjsdb.api.storage.services.v1.GetDatasetVersionRequest
	10, // 10: sciobjsdb.api.storage.services.v1.DatasetService.GetDatasetVersionObjectGroups:input_type -> sciobjsdb.api.storage.services.v1.GetDatasetVersionObjectGroupsRequest
	11, // 11: sciobjsdb.api.storage.services.v1.DatasetService.DeleteDatasetVersion:input_type -> sciobjsdb.api.storage.services.v1.DeleteDatasetVersionRequest
	12, // 12: sciobjsdb.api.storage.services.v1.DatasetService.CreateDataset:output_type -> sciobjsdb.api.storage.services.v1.CreateDatasetResponse
	13, // 13: sciobjsdb.api.storage.services.v1.DatasetService.GetDataset:output_type -> sciobjsdb.api.storage.services.v1.GetDatasetResponse
	14, // 14: sciobjsdb.api.storage.services.v1.DatasetService.GetDatasetVersions:output_type -> sciobjsdb.api.storage.services.v1.GetDatasetVersionsResponse
	15, // 15: sciobjsdb.api.storage.services.v1.DatasetService.GetDatasetObjectGroups:output_type -> sciobjsdb.api.storage.services.v1.GetDatasetObjectGroupsResponse
	16, // 16: sciobjsdb.api.storage.services.v1.DatasetService.GetObjectGroupsStreamLink:output_type -> sciobjsdb.api.storage.services.v1.GetObjectGroupsStreamLinkResponse
	17, // 17: sciobjsdb.api.storage.services.v1.DatasetService.UpdateDatasetField:output_type -> sciobjsdb.api.storage.services.v1.UpdateDatasetFieldResponse
	18, // 18: sciobjsdb.api.storage.services.v1.DatasetService.DeleteDataset:output_type -> sciobjsdb.api.storage.services.v1.DeleteDatasetResponse
	19, // 19: sciobjsdb.api.storage.services.v1.DatasetService.GetObjectGroupRevisionsInDateRange:output_type -> sciobjsdb.api.storage.services.v1.GetObjectGroupRevisionsInDateRangeResponse
	20, // 20: sciobjsdb.api.storage.services.v1.DatasetService.ReleaseDatasetVersion:output_type -> sciobjsdb.api.storage.services.v1.ReleaseDatasetVersionResponse
	21, // 21: sciobjsdb.api.storage.services.v1.DatasetService.GetDatasetVersion:output_type -> sciobjsdb.api.storage.services.v1.GetDatasetVersionResponse
	22, // 22: sciobjsdb.api.storage.services.v1.DatasetService.GetDatasetVersionObjectGroups:output_type -> sciobjsdb.api.storage.services.v1.GetDatasetVersionObjectGroupsResponse
	23, // 23: sciobjsdb.api.storage.services.v1.DatasetService.DeleteDatasetVersion:output_type -> sciobjsdb.api.storage.services.v1.DeleteDatasetVersionResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_sciobjsdb_api_storage_services_v1_dataset_service_proto_init() }
func file_sciobjsdb_api_storage_services_v1_dataset_service_proto_init() {
	if File_sciobjsdb_api_storage_services_v1_dataset_service_proto != nil {
		return
	}
	file_sciobjsdb_api_storage_services_v1_dataset_service_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sciobjsdb_api_storage_services_v1_dataset_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sciobjsdb_api_storage_services_v1_dataset_service_proto_goTypes,
		DependencyIndexes: file_sciobjsdb_api_storage_services_v1_dataset_service_proto_depIdxs,
	}.Build()
	File_sciobjsdb_api_storage_services_v1_dataset_service_proto = out.File
	file_sciobjsdb_api_storage_services_v1_dataset_service_proto_rawDesc = nil
	file_sciobjsdb_api_storage_services_v1_dataset_service_proto_goTypes = nil
	file_sciobjsdb_api_storage_services_v1_dataset_service_proto_depIdxs = nil
}
