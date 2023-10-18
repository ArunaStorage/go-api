// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: aruna/api/storage/services/v2/relations_service.proto

package v2

import (
	v2 "github.com/ArunaStorage/go-api/aruna/api/storage/models/v2"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ModifyRelationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId      string         `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	AddRelations    []*v2.Relation `protobuf:"bytes,2,rep,name=add_relations,json=addRelations,proto3" json:"add_relations,omitempty"`
	RemoveRelations []*v2.Relation `protobuf:"bytes,3,rep,name=remove_relations,json=removeRelations,proto3" json:"remove_relations,omitempty"`
}

func (x *ModifyRelationsRequest) Reset() {
	*x = ModifyRelationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyRelationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyRelationsRequest) ProtoMessage() {}

func (x *ModifyRelationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyRelationsRequest.ProtoReflect.Descriptor instead.
func (*ModifyRelationsRequest) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_relations_service_proto_rawDescGZIP(), []int{0}
}

func (x *ModifyRelationsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *ModifyRelationsRequest) GetAddRelations() []*v2.Relation {
	if x != nil {
		return x.AddRelations
	}
	return nil
}

func (x *ModifyRelationsRequest) GetRemoveRelations() []*v2.Relation {
	if x != nil {
		return x.RemoveRelations
	}
	return nil
}

type ModifyRelationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ModifyRelationsResponse) Reset() {
	*x = ModifyRelationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyRelationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyRelationsResponse) ProtoMessage() {}

func (x *ModifyRelationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyRelationsResponse.ProtoReflect.Descriptor instead.
func (*ModifyRelationsResponse) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_relations_service_proto_rawDescGZIP(), []int{1}
}

type GetHierarchyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *GetHierarchyRequest) Reset() {
	*x = GetHierarchyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHierarchyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHierarchyRequest) ProtoMessage() {}

func (x *GetHierarchyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHierarchyRequest.ProtoReflect.Descriptor instead.
func (*GetHierarchyRequest) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_relations_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetHierarchyRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type DatasetRelations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin         string   `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	ObjectChildren []string `protobuf:"bytes,2,rep,name=object_children,json=objectChildren,proto3" json:"object_children,omitempty"`
}

func (x *DatasetRelations) Reset() {
	*x = DatasetRelations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetRelations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetRelations) ProtoMessage() {}

func (x *DatasetRelations) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetRelations.ProtoReflect.Descriptor instead.
func (*DatasetRelations) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_relations_service_proto_rawDescGZIP(), []int{3}
}

func (x *DatasetRelations) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *DatasetRelations) GetObjectChildren() []string {
	if x != nil {
		return x.ObjectChildren
	}
	return nil
}

type CollectionRelations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin          string              `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	DatasetChildren []*DatasetRelations `protobuf:"bytes,2,rep,name=dataset_children,json=datasetChildren,proto3" json:"dataset_children,omitempty"`
	ObjectChildren  []string            `protobuf:"bytes,3,rep,name=object_children,json=objectChildren,proto3" json:"object_children,omitempty"`
}

func (x *CollectionRelations) Reset() {
	*x = CollectionRelations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionRelations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionRelations) ProtoMessage() {}

func (x *CollectionRelations) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionRelations.ProtoReflect.Descriptor instead.
func (*CollectionRelations) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_relations_service_proto_rawDescGZIP(), []int{4}
}

func (x *CollectionRelations) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *CollectionRelations) GetDatasetChildren() []*DatasetRelations {
	if x != nil {
		return x.DatasetChildren
	}
	return nil
}

func (x *CollectionRelations) GetObjectChildren() []string {
	if x != nil {
		return x.ObjectChildren
	}
	return nil
}

type ProjectRelations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin             string                 `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	CollectionChildren []*CollectionRelations `protobuf:"bytes,2,rep,name=collection_children,json=collectionChildren,proto3" json:"collection_children,omitempty"`
	DatasetChildren    []*DatasetRelations    `protobuf:"bytes,3,rep,name=dataset_children,json=datasetChildren,proto3" json:"dataset_children,omitempty"`
	ObjectChildren     []string               `protobuf:"bytes,4,rep,name=object_children,json=objectChildren,proto3" json:"object_children,omitempty"`
}

func (x *ProjectRelations) Reset() {
	*x = ProjectRelations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectRelations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectRelations) ProtoMessage() {}

func (x *ProjectRelations) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectRelations.ProtoReflect.Descriptor instead.
func (*ProjectRelations) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_relations_service_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectRelations) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *ProjectRelations) GetCollectionChildren() []*CollectionRelations {
	if x != nil {
		return x.CollectionChildren
	}
	return nil
}

func (x *ProjectRelations) GetDatasetChildren() []*DatasetRelations {
	if x != nil {
		return x.DatasetChildren
	}
	return nil
}

func (x *ProjectRelations) GetObjectChildren() []string {
	if x != nil {
		return x.ObjectChildren
	}
	return nil
}

type GetHierarchyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Graph:
	//
	//	*GetHierarchyResponse_Project
	//	*GetHierarchyResponse_Collection
	//	*GetHierarchyResponse_Dataset
	Graph isGetHierarchyResponse_Graph `protobuf_oneof:"graph"`
}

func (x *GetHierarchyResponse) Reset() {
	*x = GetHierarchyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHierarchyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHierarchyResponse) ProtoMessage() {}

func (x *GetHierarchyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHierarchyResponse.ProtoReflect.Descriptor instead.
func (*GetHierarchyResponse) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_services_v2_relations_service_proto_rawDescGZIP(), []int{6}
}

func (m *GetHierarchyResponse) GetGraph() isGetHierarchyResponse_Graph {
	if m != nil {
		return m.Graph
	}
	return nil
}

func (x *GetHierarchyResponse) GetProject() *ProjectRelations {
	if x, ok := x.GetGraph().(*GetHierarchyResponse_Project); ok {
		return x.Project
	}
	return nil
}

func (x *GetHierarchyResponse) GetCollection() *CollectionRelations {
	if x, ok := x.GetGraph().(*GetHierarchyResponse_Collection); ok {
		return x.Collection
	}
	return nil
}

func (x *GetHierarchyResponse) GetDataset() *DatasetRelations {
	if x, ok := x.GetGraph().(*GetHierarchyResponse_Dataset); ok {
		return x.Dataset
	}
	return nil
}

type isGetHierarchyResponse_Graph interface {
	isGetHierarchyResponse_Graph()
}

type GetHierarchyResponse_Project struct {
	Project *ProjectRelations `protobuf:"bytes,1,opt,name=project,proto3,oneof"`
}

type GetHierarchyResponse_Collection struct {
	Collection *CollectionRelations `protobuf:"bytes,2,opt,name=collection,proto3,oneof"`
}

type GetHierarchyResponse_Dataset struct {
	Dataset *DatasetRelations `protobuf:"bytes,3,opt,name=dataset,proto3,oneof"`
}

func (*GetHierarchyResponse_Project) isGetHierarchyResponse_Graph() {}

func (*GetHierarchyResponse_Collection) isGetHierarchyResponse_Graph() {}

func (*GetHierarchyResponse_Dataset) isGetHierarchyResponse_Graph() {}

var File_aruna_api_storage_services_v2_relations_service_proto protoreflect.FileDescriptor

var file_aruna_api_storage_services_v2_relations_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x2f,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x28, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7,
	0x01, 0x0a, 0x16, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x0d, 0x61, 0x64,
	0x64, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x50, 0x0a, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x48, 0x69, 0x65, 0x72, 0x61, 0x72,
	0x63, 0x68, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x10, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x22, 0xb2, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x5a, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x72, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x72, 0x75,
	0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x94, 0x02, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x12, 0x63, 0x0a, 0x13, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x5a, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x72, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x8f, 0x02, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x48, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x54, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x72, 0x75, 0x6e,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x07, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x32, 0xc8,
	0x02, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36,
	0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01,
	0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x97, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79,
	0x12, 0x32, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x12, 0x16, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x68, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x42, 0x92, 0x01, 0x0a, 0x3e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x41, 0x72, 0x75, 0x6e, 0x61, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x10, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x75,
	0x6e, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aruna_api_storage_services_v2_relations_service_proto_rawDescOnce sync.Once
	file_aruna_api_storage_services_v2_relations_service_proto_rawDescData = file_aruna_api_storage_services_v2_relations_service_proto_rawDesc
)

func file_aruna_api_storage_services_v2_relations_service_proto_rawDescGZIP() []byte {
	file_aruna_api_storage_services_v2_relations_service_proto_rawDescOnce.Do(func() {
		file_aruna_api_storage_services_v2_relations_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_aruna_api_storage_services_v2_relations_service_proto_rawDescData)
	})
	return file_aruna_api_storage_services_v2_relations_service_proto_rawDescData
}

var file_aruna_api_storage_services_v2_relations_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_aruna_api_storage_services_v2_relations_service_proto_goTypes = []interface{}{
	(*ModifyRelationsRequest)(nil),  // 0: aruna.api.storage.services.v2.ModifyRelationsRequest
	(*ModifyRelationsResponse)(nil), // 1: aruna.api.storage.services.v2.ModifyRelationsResponse
	(*GetHierarchyRequest)(nil),     // 2: aruna.api.storage.services.v2.GetHierarchyRequest
	(*DatasetRelations)(nil),        // 3: aruna.api.storage.services.v2.DatasetRelations
	(*CollectionRelations)(nil),     // 4: aruna.api.storage.services.v2.CollectionRelations
	(*ProjectRelations)(nil),        // 5: aruna.api.storage.services.v2.ProjectRelations
	(*GetHierarchyResponse)(nil),    // 6: aruna.api.storage.services.v2.GetHierarchyResponse
	(*v2.Relation)(nil),             // 7: aruna.api.storage.models.v2.Relation
}
var file_aruna_api_storage_services_v2_relations_service_proto_depIdxs = []int32{
	7,  // 0: aruna.api.storage.services.v2.ModifyRelationsRequest.add_relations:type_name -> aruna.api.storage.models.v2.Relation
	7,  // 1: aruna.api.storage.services.v2.ModifyRelationsRequest.remove_relations:type_name -> aruna.api.storage.models.v2.Relation
	3,  // 2: aruna.api.storage.services.v2.CollectionRelations.dataset_children:type_name -> aruna.api.storage.services.v2.DatasetRelations
	4,  // 3: aruna.api.storage.services.v2.ProjectRelations.collection_children:type_name -> aruna.api.storage.services.v2.CollectionRelations
	3,  // 4: aruna.api.storage.services.v2.ProjectRelations.dataset_children:type_name -> aruna.api.storage.services.v2.DatasetRelations
	5,  // 5: aruna.api.storage.services.v2.GetHierarchyResponse.project:type_name -> aruna.api.storage.services.v2.ProjectRelations
	4,  // 6: aruna.api.storage.services.v2.GetHierarchyResponse.collection:type_name -> aruna.api.storage.services.v2.CollectionRelations
	3,  // 7: aruna.api.storage.services.v2.GetHierarchyResponse.dataset:type_name -> aruna.api.storage.services.v2.DatasetRelations
	0,  // 8: aruna.api.storage.services.v2.RelationsService.ModifyRelations:input_type -> aruna.api.storage.services.v2.ModifyRelationsRequest
	2,  // 9: aruna.api.storage.services.v2.RelationsService.GetHierarchy:input_type -> aruna.api.storage.services.v2.GetHierarchyRequest
	1,  // 10: aruna.api.storage.services.v2.RelationsService.ModifyRelations:output_type -> aruna.api.storage.services.v2.ModifyRelationsResponse
	6,  // 11: aruna.api.storage.services.v2.RelationsService.GetHierarchy:output_type -> aruna.api.storage.services.v2.GetHierarchyResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_aruna_api_storage_services_v2_relations_service_proto_init() }
func file_aruna_api_storage_services_v2_relations_service_proto_init() {
	if File_aruna_api_storage_services_v2_relations_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyRelationsRequest); i {
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
		file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyRelationsResponse); i {
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
		file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHierarchyRequest); i {
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
		file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetRelations); i {
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
		file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionRelations); i {
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
		file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectRelations); i {
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
		file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHierarchyResponse); i {
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
	file_aruna_api_storage_services_v2_relations_service_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*GetHierarchyResponse_Project)(nil),
		(*GetHierarchyResponse_Collection)(nil),
		(*GetHierarchyResponse_Dataset)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aruna_api_storage_services_v2_relations_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aruna_api_storage_services_v2_relations_service_proto_goTypes,
		DependencyIndexes: file_aruna_api_storage_services_v2_relations_service_proto_depIdxs,
		MessageInfos:      file_aruna_api_storage_services_v2_relations_service_proto_msgTypes,
	}.Build()
	File_aruna_api_storage_services_v2_relations_service_proto = out.File
	file_aruna_api_storage_services_v2_relations_service_proto_rawDesc = nil
	file_aruna_api_storage_services_v2_relations_service_proto_goTypes = nil
	file_aruna_api_storage_services_v2_relations_service_proto_depIdxs = nil
}
