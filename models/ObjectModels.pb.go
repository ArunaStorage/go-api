// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: api/models/ObjectModels.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type ObjectHeritage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string            `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"` // Name: Name of the dataset object heritage entity e.g.: mydata
	DatasetID string            `protobuf:"bytes,3,opt,name=DatasetID,proto3" json:"DatasetID,omitempty"`
	Labels    map[string]string `protobuf:"bytes,4,rep,name=Labels,proto3" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ObjectHeritage) Reset() {
	*x = ObjectHeritage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_ObjectModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectHeritage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectHeritage) ProtoMessage() {}

func (x *ObjectHeritage) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_ObjectModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectHeritage.ProtoReflect.Descriptor instead.
func (*ObjectHeritage) Descriptor() ([]byte, []int) {
	return file_api_models_ObjectModels_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectHeritage) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ObjectHeritage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectHeritage) GetDatasetID() string {
	if x != nil {
		return x.DatasetID
	}
	return ""
}

func (x *ObjectHeritage) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type DatasetObjectGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 string                     `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ObjectHeritageID   string                     `protobuf:"bytes,2,opt,name=ObjectHeritageID,proto3" json:"ObjectHeritageID,omitempty"`
	Name               string                     `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Version            *Version                   `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version,omitempty"`
	InitializedObjects int64                      `protobuf:"varint,5,opt,name=InitializedObjects,proto3" json:"InitializedObjects,omitempty"`
	UploadedObjects    int64                      `protobuf:"varint,6,opt,name=UploadedObjects,proto3" json:"UploadedObjects,omitempty"`
	Objects            []*DatasetObjectEntry      `protobuf:"bytes,7,rep,name=Objects,proto3" json:"Objects,omitempty"`
	AdditionalMetadata map[string]*_struct.Struct `protobuf:"bytes,8,rep,name=AdditionalMetadata,proto3" json:"AdditionalMetadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Additional metadata of the object
	Labels             map[string]string          `protobuf:"bytes,9,rep,name=Labels,proto3" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DatasetObjectGroup) Reset() {
	*x = DatasetObjectGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_ObjectModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetObjectGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetObjectGroup) ProtoMessage() {}

func (x *DatasetObjectGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_ObjectModels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetObjectGroup.ProtoReflect.Descriptor instead.
func (*DatasetObjectGroup) Descriptor() ([]byte, []int) {
	return file_api_models_ObjectModels_proto_rawDescGZIP(), []int{1}
}

func (x *DatasetObjectGroup) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DatasetObjectGroup) GetObjectHeritageID() string {
	if x != nil {
		return x.ObjectHeritageID
	}
	return ""
}

func (x *DatasetObjectGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DatasetObjectGroup) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *DatasetObjectGroup) GetInitializedObjects() int64 {
	if x != nil {
		return x.InitializedObjects
	}
	return 0
}

func (x *DatasetObjectGroup) GetUploadedObjects() int64 {
	if x != nil {
		return x.UploadedObjects
	}
	return 0
}

func (x *DatasetObjectGroup) GetObjects() []*DatasetObjectEntry {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *DatasetObjectGroup) GetAdditionalMetadata() map[string]*_struct.Struct {
	if x != nil {
		return x.AdditionalMetadata
	}
	return nil
}

func (x *DatasetObjectGroup) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type DatasetObjectEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 string                     `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`                                                                                                                         //ID of the entity
	Filename           string                     `protobuf:"bytes,2,opt,name=Filename,proto3" json:"Filename,omitempty"`                                                                                                             // Filename: Name of the original file e.g.: mydata.json
	Filetype           string                     `protobuf:"bytes,3,opt,name=Filetype,proto3" json:"Filetype,omitempty"`                                                                                                             // Filetype: Type of the stored file, e.g.: json, gbff,...
	Origin             *Origin                    `protobuf:"bytes,4,opt,name=Origin,proto3" json:"Origin,omitempty"`                                                                                                                 // Origin: Source of the dataset
	ContentLen         int64                      `protobuf:"varint,5,opt,name=ContentLen,proto3" json:"ContentLen,omitempty"`                                                                                                        // ContentLen: Lenght of the stored dataset
	Location           *Location                  `protobuf:"bytes,6,opt,name=Location,proto3" json:"Location,omitempty"`                                                                                                             // Location: Location of the data
	Created            *timestamp.Timestamp       `protobuf:"bytes,7,opt,name=Created,proto3" json:"Created,omitempty"`                                                                                                               // When the data object was created
	AdditionalMetadata map[string]*_struct.Struct `protobuf:"bytes,8,rep,name=AdditionalMetadata,proto3" json:"AdditionalMetadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Additional metadata of the object
	UploadID           string                     `protobuf:"bytes,9,opt,name=UploadID,proto3" json:"UploadID,omitempty"`
	Status             Status                     `protobuf:"varint,10,opt,name=Status,proto3,enum=Status" json:"Status,omitempty"`
}

func (x *DatasetObjectEntry) Reset() {
	*x = DatasetObjectEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_ObjectModels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetObjectEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetObjectEntry) ProtoMessage() {}

func (x *DatasetObjectEntry) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_ObjectModels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetObjectEntry.ProtoReflect.Descriptor instead.
func (*DatasetObjectEntry) Descriptor() ([]byte, []int) {
	return file_api_models_ObjectModels_proto_rawDescGZIP(), []int{2}
}

func (x *DatasetObjectEntry) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DatasetObjectEntry) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DatasetObjectEntry) GetFiletype() string {
	if x != nil {
		return x.Filetype
	}
	return ""
}

func (x *DatasetObjectEntry) GetOrigin() *Origin {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *DatasetObjectEntry) GetContentLen() int64 {
	if x != nil {
		return x.ContentLen
	}
	return 0
}

func (x *DatasetObjectEntry) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *DatasetObjectEntry) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *DatasetObjectEntry) GetAdditionalMetadata() map[string]*_struct.Struct {
	if x != nil {
		return x.AdditionalMetadata
	}
	return nil
}

func (x *DatasetObjectEntry) GetUploadID() string {
	if x != nil {
		return x.UploadID
	}
	return ""
}

func (x *DatasetObjectEntry) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Initiating
}

var File_api_models_ObjectModels_proto protoreflect.FileDescriptor

var file_api_models_ObjectModels_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x0e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x48, 0x65, 0x72, 0x69, 0x74, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x06,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x65, 0x72, 0x69, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x44, 0x92, 0x41,
	0x41, 0x0a, 0x3f, 0x2a, 0x0e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x65, 0x72, 0x69, 0x74,
	0x61, 0x67, 0x65, 0x32, 0x2d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x20, 0x6f,
	0x66, 0x20, 0x61, 0x6e, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x22, 0x80, 0x06, 0x0a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x48, 0x65, 0x72, 0x69, 0x74, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x65, 0x72, 0x69, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x12, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x28, 0x0a,
	0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x5b, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x12, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x5e, 0x0a, 0x17,
	0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0xbb, 0x01, 0x92, 0x41, 0xb7, 0x01, 0x0a, 0xb4,
	0x01, 0x2a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x32, 0x9d, 0x01, 0x41, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x61, 0x72, 0x65, 0x20, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x64, 0x20, 0x74, 0x6f, 0x67, 0x65, 0x74, 0x68, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x65,
	0x61, 0x73, 0x69, 0x65, 0x72, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x20, 0x41, 0x6e, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x75, 0x73, 0x65,
	0x20, 0x63, 0x61, 0x73, 0x65, 0x20, 0x77, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x61,
	0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x20, 0x64,
	0x61, 0x74, 0x61, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x20, 0x61, 0x6c, 0x6f, 0x6e,
	0x67, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x20,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x22, 0xb3, 0x04, 0x0a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x4c, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x5b, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x41, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x5e, 0x0a, 0x17,
	0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x3d, 0x92, 0x41,
	0x3a, 0x0a, 0x38, 0x2a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x32, 0x22, 0x52, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x20, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_models_ObjectModels_proto_rawDescOnce sync.Once
	file_api_models_ObjectModels_proto_rawDescData = file_api_models_ObjectModels_proto_rawDesc
)

func file_api_models_ObjectModels_proto_rawDescGZIP() []byte {
	file_api_models_ObjectModels_proto_rawDescOnce.Do(func() {
		file_api_models_ObjectModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_models_ObjectModels_proto_rawDescData)
	})
	return file_api_models_ObjectModels_proto_rawDescData
}

var file_api_models_ObjectModels_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_models_ObjectModels_proto_goTypes = []interface{}{
	(*ObjectHeritage)(nil),      // 0: ObjectHeritage
	(*DatasetObjectGroup)(nil),  // 1: DatasetObjectGroup
	(*DatasetObjectEntry)(nil),  // 2: DatasetObjectEntry
	nil,                         // 3: ObjectHeritage.LabelsEntry
	nil,                         // 4: DatasetObjectGroup.AdditionalMetadataEntry
	nil,                         // 5: DatasetObjectGroup.LabelsEntry
	nil,                         // 6: DatasetObjectEntry.AdditionalMetadataEntry
	(*Version)(nil),             // 7: Version
	(*Origin)(nil),              // 8: Origin
	(*Location)(nil),            // 9: Location
	(*timestamp.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(Status)(0),                 // 11: Status
	(*_struct.Struct)(nil),      // 12: google.protobuf.Struct
}
var file_api_models_ObjectModels_proto_depIdxs = []int32{
	3,  // 0: ObjectHeritage.Labels:type_name -> ObjectHeritage.LabelsEntry
	7,  // 1: DatasetObjectGroup.Version:type_name -> Version
	2,  // 2: DatasetObjectGroup.Objects:type_name -> DatasetObjectEntry
	4,  // 3: DatasetObjectGroup.AdditionalMetadata:type_name -> DatasetObjectGroup.AdditionalMetadataEntry
	5,  // 4: DatasetObjectGroup.Labels:type_name -> DatasetObjectGroup.LabelsEntry
	8,  // 5: DatasetObjectEntry.Origin:type_name -> Origin
	9,  // 6: DatasetObjectEntry.Location:type_name -> Location
	10, // 7: DatasetObjectEntry.Created:type_name -> google.protobuf.Timestamp
	6,  // 8: DatasetObjectEntry.AdditionalMetadata:type_name -> DatasetObjectEntry.AdditionalMetadataEntry
	11, // 9: DatasetObjectEntry.Status:type_name -> Status
	12, // 10: DatasetObjectGroup.AdditionalMetadataEntry.value:type_name -> google.protobuf.Struct
	12, // 11: DatasetObjectEntry.AdditionalMetadataEntry.value:type_name -> google.protobuf.Struct
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_api_models_ObjectModels_proto_init() }
func file_api_models_ObjectModels_proto_init() {
	if File_api_models_ObjectModels_proto != nil {
		return
	}
	file_api_models_CommonModels_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_models_ObjectModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectHeritage); i {
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
		file_api_models_ObjectModels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetObjectGroup); i {
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
		file_api_models_ObjectModels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetObjectEntry); i {
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
			RawDescriptor: file_api_models_ObjectModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_models_ObjectModels_proto_goTypes,
		DependencyIndexes: file_api_models_ObjectModels_proto_depIdxs,
		MessageInfos:      file_api_models_ObjectModels_proto_msgTypes,
	}.Build()
	File_api_models_ObjectModels_proto = out.File
	file_api_models_ObjectModels_proto_rawDesc = nil
	file_api_models_ObjectModels_proto_goTypes = nil
	file_api_models_ObjectModels_proto_depIdxs = nil
}
