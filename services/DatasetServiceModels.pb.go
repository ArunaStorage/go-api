// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: api/services/DatasetServiceModels.proto

package services

import (
	models "github.com/ScienceObjectsDB/go-api/models"
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

// Dataset related Models
type CreateDatasetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                            // Name of the dataset
	ProjectId string             `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"` //ProjectID of the project the dataset is associated with
	Labels    []*models.Label    `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	Metadata  []*models.Metadata `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *CreateDatasetRequest) Reset() {
	*x = CreateDatasetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_DatasetServiceModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatasetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatasetRequest) ProtoMessage() {}

func (x *CreateDatasetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_DatasetServiceModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatasetRequest.ProtoReflect.Descriptor instead.
func (*CreateDatasetRequest) Descriptor() ([]byte, []int) {
	return file_api_services_DatasetServiceModels_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDatasetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDatasetRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateDatasetRequest) GetLabels() []*models.Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateDatasetRequest) GetMetadata() []*models.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type DatasetVersionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetVersion []*models.DatasetVersion `protobuf:"bytes,1,rep,name=dataset_version,json=datasetVersion,proto3" json:"dataset_version,omitempty"`
}

func (x *DatasetVersionList) Reset() {
	*x = DatasetVersionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_DatasetServiceModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetVersionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetVersionList) ProtoMessage() {}

func (x *DatasetVersionList) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_DatasetServiceModels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetVersionList.ProtoReflect.Descriptor instead.
func (*DatasetVersionList) Descriptor() ([]byte, []int) {
	return file_api_services_DatasetServiceModels_proto_rawDescGZIP(), []int{1}
}

func (x *DatasetVersionList) GetDatasetVersion() []*models.DatasetVersion {
	if x != nil {
		return x.DatasetVersion
	}
	return nil
}

type ReleaseDatasetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DatasetId      string             `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	Version        *models.Version    `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	ObjectGroupIds []string           `protobuf:"bytes,4,rep,name=object_group_ids,json=objectGroupIds,proto3" json:"object_group_ids,omitempty"`
	Labels         []*models.Label    `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
	Metadata       []*models.Metadata `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ReleaseDatasetVersionRequest) Reset() {
	*x = ReleaseDatasetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_DatasetServiceModels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseDatasetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseDatasetVersionRequest) ProtoMessage() {}

func (x *ReleaseDatasetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_DatasetServiceModels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseDatasetVersionRequest.ProtoReflect.Descriptor instead.
func (*ReleaseDatasetVersionRequest) Descriptor() ([]byte, []int) {
	return file_api_services_DatasetServiceModels_proto_rawDescGZIP(), []int{2}
}

func (x *ReleaseDatasetVersionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReleaseDatasetVersionRequest) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *ReleaseDatasetVersionRequest) GetVersion() *models.Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *ReleaseDatasetVersionRequest) GetObjectGroupIds() []string {
	if x != nil {
		return x.ObjectGroupIds
	}
	return nil
}

func (x *ReleaseDatasetVersionRequest) GetLabels() []*models.Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ReleaseDatasetVersionRequest) GetMetadata() []*models.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ObjectGroupList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectGroups []*models.ObjectGroup `protobuf:"bytes,1,rep,name=object_groups,json=objectGroups,proto3" json:"object_groups,omitempty"`
}

func (x *ObjectGroupList) Reset() {
	*x = ObjectGroupList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_DatasetServiceModels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectGroupList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectGroupList) ProtoMessage() {}

func (x *ObjectGroupList) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_DatasetServiceModels_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectGroupList.ProtoReflect.Descriptor instead.
func (*ObjectGroupList) Descriptor() ([]byte, []int) {
	return file_api_services_DatasetServiceModels_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectGroupList) GetObjectGroups() []*models.ObjectGroup {
	if x != nil {
		return x.ObjectGroups
	}
	return nil
}

var File_api_services_DatasetServiceModels_proto protoreflect.FileDescriptor

var file_api_services_DatasetServiceModels_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2c,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x55, 0x0a, 0x12,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0xfb, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x4b, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2f, 0x67, 0x6f,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_services_DatasetServiceModels_proto_rawDescOnce sync.Once
	file_api_services_DatasetServiceModels_proto_rawDescData = file_api_services_DatasetServiceModels_proto_rawDesc
)

func file_api_services_DatasetServiceModels_proto_rawDescGZIP() []byte {
	file_api_services_DatasetServiceModels_proto_rawDescOnce.Do(func() {
		file_api_services_DatasetServiceModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_DatasetServiceModels_proto_rawDescData)
	})
	return file_api_services_DatasetServiceModels_proto_rawDescData
}

var file_api_services_DatasetServiceModels_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_services_DatasetServiceModels_proto_goTypes = []interface{}{
	(*CreateDatasetRequest)(nil),         // 0: services.CreateDatasetRequest
	(*DatasetVersionList)(nil),           // 1: services.DatasetVersionList
	(*ReleaseDatasetVersionRequest)(nil), // 2: services.ReleaseDatasetVersionRequest
	(*ObjectGroupList)(nil),              // 3: services.ObjectGroupList
	(*models.Label)(nil),                 // 4: models.Label
	(*models.Metadata)(nil),              // 5: models.Metadata
	(*models.DatasetVersion)(nil),        // 6: models.DatasetVersion
	(*models.Version)(nil),               // 7: models.Version
	(*models.ObjectGroup)(nil),           // 8: models.ObjectGroup
}
var file_api_services_DatasetServiceModels_proto_depIdxs = []int32{
	4, // 0: services.CreateDatasetRequest.labels:type_name -> models.Label
	5, // 1: services.CreateDatasetRequest.metadata:type_name -> models.Metadata
	6, // 2: services.DatasetVersionList.dataset_version:type_name -> models.DatasetVersion
	7, // 3: services.ReleaseDatasetVersionRequest.version:type_name -> models.Version
	4, // 4: services.ReleaseDatasetVersionRequest.labels:type_name -> models.Label
	5, // 5: services.ReleaseDatasetVersionRequest.metadata:type_name -> models.Metadata
	8, // 6: services.ObjectGroupList.object_groups:type_name -> models.ObjectGroup
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_services_DatasetServiceModels_proto_init() }
func file_api_services_DatasetServiceModels_proto_init() {
	if File_api_services_DatasetServiceModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_DatasetServiceModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatasetRequest); i {
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
		file_api_services_DatasetServiceModels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetVersionList); i {
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
		file_api_services_DatasetServiceModels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseDatasetVersionRequest); i {
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
		file_api_services_DatasetServiceModels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectGroupList); i {
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
			RawDescriptor: file_api_services_DatasetServiceModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_services_DatasetServiceModels_proto_goTypes,
		DependencyIndexes: file_api_services_DatasetServiceModels_proto_depIdxs,
		MessageInfos:      file_api_services_DatasetServiceModels_proto_msgTypes,
	}.Build()
	File_api_services_DatasetServiceModels_proto = out.File
	file_api_services_DatasetServiceModels_proto_rawDesc = nil
	file_api_services_DatasetServiceModels_proto_goTypes = nil
	file_api_services_DatasetServiceModels_proto_depIdxs = nil
}
