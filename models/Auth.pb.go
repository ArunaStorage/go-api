// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: api/models/Auth.proto

package models

import (
	models "github.com/ScienceObjectsDB/go-api/models"
	proto "github.com/golang/protobuf/proto"
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

type TokenList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID string        `protobuf:"bytes,1,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	Token     []*TokenEntry `protobuf:"bytes,2,rep,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenList) Reset() {
	*x = TokenList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_Auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenList) ProtoMessage() {}

func (x *TokenList) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_Auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenList.ProtoReflect.Descriptor instead.
func (*TokenList) Descriptor() ([]byte, []int) {
	return file_api_models_Auth_proto_rawDescGZIP(), []int{0}
}

func (x *TokenList) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *TokenList) GetToken() []*TokenEntry {
	if x != nil {
		return x.Token
	}
	return nil
}

type TokenEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID   *models.User         `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Token    string               `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	Resource models.Resource      `protobuf:"varint,4,opt,name=Resource,proto3,enum=Resource" json:"Resource,omitempty"`
	Created  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=Created,proto3" json:"Created,omitempty"` // When the token was created
	Expires  *timestamp.Timestamp `protobuf:"bytes,6,opt,name=Expires,proto3" json:"Expires,omitempty"` // When the token expires
}

func (x *TokenEntry) Reset() {
	*x = TokenEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_Auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenEntry) ProtoMessage() {}

func (x *TokenEntry) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_Auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenEntry.ProtoReflect.Descriptor instead.
func (*TokenEntry) Descriptor() ([]byte, []int) {
	return file_api_models_Auth_proto_rawDescGZIP(), []int{1}
}

func (x *TokenEntry) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *TokenEntry) GetUserID() *models.User {
	if x != nil {
		return x.UserID
	}
	return nil
}

func (x *TokenEntry) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TokenEntry) GetResource() models.Resource {
	if x != nil {
		return x.Resource
	}
	return models.Resource_Project
}

func (x *TokenEntry) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *TokenEntry) GetExpires() *timestamp.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

type CreateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceID string               `protobuf:"bytes,1,opt,name=ResourceID,proto3" json:"ResourceID,omitempty"`
	Rights     []models.Right       `protobuf:"varint,2,rep,packed,name=Rights,proto3,enum=Right" json:"Rights,omitempty"`
	Resource   models.Resource      `protobuf:"varint,3,opt,name=Resource,proto3,enum=Resource" json:"Resource,omitempty"`
	Expires    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=Expires,proto3" json:"Expires,omitempty"` // When the token expires
}

func (x *CreateTokenRequest) Reset() {
	*x = CreateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_Auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRequest) ProtoMessage() {}

func (x *CreateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_Auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_models_Auth_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTokenRequest) GetResourceID() string {
	if x != nil {
		return x.ResourceID
	}
	return ""
}

func (x *CreateTokenRequest) GetRights() []models.Right {
	if x != nil {
		return x.Rights
	}
	return nil
}

func (x *CreateTokenRequest) GetResource() models.Resource {
	if x != nil {
		return x.Resource
	}
	return models.Resource_Project
}

func (x *CreateTokenRequest) GetExpires() *timestamp.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

var File_api_models_Auth_proto protoreflect.FileDescriptor

var file_api_models_Auth_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x41, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x89, 0x01, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x21, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x3a, 0x3b, 0x92, 0x41, 0x38, 0x0a, 0x36, 0x2a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x32, 0x29, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x22, 0x96, 0x02,
	0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x25, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34,
	0x0a, 0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x3a, 0x30, 0x92, 0x41, 0x2d, 0x0a, 0x2b, 0x2a, 0x0a, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x32, 0x1d, 0x41, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x70, 0x69,
	0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x84, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x06, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x06, 0x2e,
	0x52, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x25, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x09, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x3a, 0x51, 0x92, 0x41, 0x4e, 0x0a,
	0x4c, 0x2a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x36, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x64,
	0x61, 0x74, 0x61, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20,
	0x6e, 0x65, 0x77, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20,
	0x67, 0x69, 0x76, 0x65, 0x6e, 0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x42, 0x3b, 0x5a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x2d, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x62, 0x69, 0x6f,
	0x2f, 0x42, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_models_Auth_proto_rawDescOnce sync.Once
	file_api_models_Auth_proto_rawDescData = file_api_models_Auth_proto_rawDesc
)

func file_api_models_Auth_proto_rawDescGZIP() []byte {
	file_api_models_Auth_proto_rawDescOnce.Do(func() {
		file_api_models_Auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_models_Auth_proto_rawDescData)
	})
	return file_api_models_Auth_proto_rawDescData
}

var file_api_models_Auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_models_Auth_proto_goTypes = []interface{}{
	(*TokenList)(nil),           // 0: TokenList
	(*TokenEntry)(nil),          // 1: TokenEntry
	(*CreateTokenRequest)(nil),  // 2: CreateTokenRequest
	(*models.User)(nil),         // 3: User
	(models.Resource)(0),        // 4: Resource
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(models.Right)(0),           // 6: Right
}
var file_api_models_Auth_proto_depIdxs = []int32{
	1, // 0: TokenList.token:type_name -> TokenEntry
	3, // 1: TokenEntry.UserID:type_name -> User
	4, // 2: TokenEntry.Resource:type_name -> Resource
	5, // 3: TokenEntry.Created:type_name -> google.protobuf.Timestamp
	5, // 4: TokenEntry.Expires:type_name -> google.protobuf.Timestamp
	6, // 5: CreateTokenRequest.Rights:type_name -> Right
	4, // 6: CreateTokenRequest.Resource:type_name -> Resource
	5, // 7: CreateTokenRequest.Expires:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_api_models_Auth_proto_init() }
func file_api_models_Auth_proto_init() {
	if File_api_models_Auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_models_Auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenList); i {
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
		file_api_models_Auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenEntry); i {
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
		file_api_models_Auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenRequest); i {
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
			RawDescriptor: file_api_models_Auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_models_Auth_proto_goTypes,
		DependencyIndexes: file_api_models_Auth_proto_depIdxs,
		MessageInfos:      file_api_models_Auth_proto_msgTypes,
	}.Build()
	File_api_models_Auth_proto = out.File
	file_api_models_Auth_proto_rawDesc = nil
	file_api_models_Auth_proto_goTypes = nil
	file_api_models_Auth_proto_depIdxs = nil
}
