// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: sciobjsdb/api/storage/services/v1/dataset_object_service.proto

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

var File_sciobjsdb_api_storage_services_v1_dataset_object_service_proto protoreflect.FileDescriptor

var file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x45, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfa, 0x08, 0x0a, 0x15, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xb5, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3b, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62,
	0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64,
	0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xc9, 0x01, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x40, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64,
	0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a,
	0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xa9, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x38, 0x2e, 0x73, 0x63, 0x69, 0x6f,
	0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67,
	0x65, 0x74, 0x12, 0xb3, 0x01, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3c, 0x2e, 0x73, 0x63, 0x69, 0x6f,
	0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a,
	0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01,
	0x2a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0xc7, 0x01, 0x0a, 0x17, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x41, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a,
	0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x12, 0xb0, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3b, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62,
	0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x73, 0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64,
	0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xa8, 0x01, 0x0a, 0x46, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x44, 0x42, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x16, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x63, 0x69, 0x6f, 0x62, 0x6a, 0x73, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_goTypes = []interface{}{
	(*CreateObjectGroupRequest)(nil),        // 0: sciobjsdb.api.storage.services.v1.CreateObjectGroupRequest
	(*CreateObjectGroupBatchRequest)(nil),   // 1: sciobjsdb.api.storage.services.v1.CreateObjectGroupBatchRequest
	(*GetObjectGroupRequest)(nil),           // 2: sciobjsdb.api.storage.services.v1.GetObjectGroupRequest
	(*FinishObjectUploadRequest)(nil),       // 3: sciobjsdb.api.storage.services.v1.FinishObjectUploadRequest
	(*FinishObjectGroupUploadRequest)(nil),  // 4: sciobjsdb.api.storage.services.v1.FinishObjectGroupUploadRequest
	(*DeleteObjectGroupRequest)(nil),        // 5: sciobjsdb.api.storage.services.v1.DeleteObjectGroupRequest
	(*CreateObjectGroupResponse)(nil),       // 6: sciobjsdb.api.storage.services.v1.CreateObjectGroupResponse
	(*CreateObjectGroupBatchResponse)(nil),  // 7: sciobjsdb.api.storage.services.v1.CreateObjectGroupBatchResponse
	(*GetObjectGroupResponse)(nil),          // 8: sciobjsdb.api.storage.services.v1.GetObjectGroupResponse
	(*FinishObjectUploadResponse)(nil),      // 9: sciobjsdb.api.storage.services.v1.FinishObjectUploadResponse
	(*FinishObjectGroupUploadResponse)(nil), // 10: sciobjsdb.api.storage.services.v1.FinishObjectGroupUploadResponse
	(*DeleteObjectGroupResponse)(nil),       // 11: sciobjsdb.api.storage.services.v1.DeleteObjectGroupResponse
}
var file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_depIdxs = []int32{
	0,  // 0: sciobjsdb.api.storage.services.v1.DatasetObjectsService.CreateObjectGroup:input_type -> sciobjsdb.api.storage.services.v1.CreateObjectGroupRequest
	1,  // 1: sciobjsdb.api.storage.services.v1.DatasetObjectsService.CreateObjectGroupBatch:input_type -> sciobjsdb.api.storage.services.v1.CreateObjectGroupBatchRequest
	2,  // 2: sciobjsdb.api.storage.services.v1.DatasetObjectsService.GetObjectGroup:input_type -> sciobjsdb.api.storage.services.v1.GetObjectGroupRequest
	3,  // 3: sciobjsdb.api.storage.services.v1.DatasetObjectsService.FinishObjectUpload:input_type -> sciobjsdb.api.storage.services.v1.FinishObjectUploadRequest
	4,  // 4: sciobjsdb.api.storage.services.v1.DatasetObjectsService.FinishObjectGroupUpload:input_type -> sciobjsdb.api.storage.services.v1.FinishObjectGroupUploadRequest
	5,  // 5: sciobjsdb.api.storage.services.v1.DatasetObjectsService.DeleteObjectGroup:input_type -> sciobjsdb.api.storage.services.v1.DeleteObjectGroupRequest
	6,  // 6: sciobjsdb.api.storage.services.v1.DatasetObjectsService.CreateObjectGroup:output_type -> sciobjsdb.api.storage.services.v1.CreateObjectGroupResponse
	7,  // 7: sciobjsdb.api.storage.services.v1.DatasetObjectsService.CreateObjectGroupBatch:output_type -> sciobjsdb.api.storage.services.v1.CreateObjectGroupBatchResponse
	8,  // 8: sciobjsdb.api.storage.services.v1.DatasetObjectsService.GetObjectGroup:output_type -> sciobjsdb.api.storage.services.v1.GetObjectGroupResponse
	9,  // 9: sciobjsdb.api.storage.services.v1.DatasetObjectsService.FinishObjectUpload:output_type -> sciobjsdb.api.storage.services.v1.FinishObjectUploadResponse
	10, // 10: sciobjsdb.api.storage.services.v1.DatasetObjectsService.FinishObjectGroupUpload:output_type -> sciobjsdb.api.storage.services.v1.FinishObjectGroupUploadResponse
	11, // 11: sciobjsdb.api.storage.services.v1.DatasetObjectsService.DeleteObjectGroup:output_type -> sciobjsdb.api.storage.services.v1.DeleteObjectGroupResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_init() }
func file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_init() {
	if File_sciobjsdb_api_storage_services_v1_dataset_object_service_proto != nil {
		return
	}
	file_sciobjsdb_api_storage_services_v1_dataset_object_service_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_goTypes,
		DependencyIndexes: file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_depIdxs,
	}.Build()
	File_sciobjsdb_api_storage_services_v1_dataset_object_service_proto = out.File
	file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_rawDesc = nil
	file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_goTypes = nil
	file_sciobjsdb_api_storage_services_v1_dataset_object_service_proto_depIdxs = nil
}
