// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/services/v1/notification_service_models.proto

package v1

import (
	v1 "github.com/ScienceObjectsDB/go-api/api/models/v1"
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

type NotificationStreamRequest_EventResources int32

const (
	NotificationStreamRequest_EVENT_RESOURCES_UNSPECIFIED              NotificationStreamRequest_EventResources = 0
	NotificationStreamRequest_EVENT_RESOURCES_PROJECT_RESOURCE         NotificationStreamRequest_EventResources = 1
	NotificationStreamRequest_EVENT_RESOURCES_DATASET_RESOURCE         NotificationStreamRequest_EventResources = 2
	NotificationStreamRequest_EVENT_RESOURCES_DATASET_VERSION_RESOURCE NotificationStreamRequest_EventResources = 3
	NotificationStreamRequest_EVENT_RESOURCES_OBJECT_GROUP_RESOURCE    NotificationStreamRequest_EventResources = 4
)

// Enum value maps for NotificationStreamRequest_EventResources.
var (
	NotificationStreamRequest_EventResources_name = map[int32]string{
		0: "EVENT_RESOURCES_UNSPECIFIED",
		1: "EVENT_RESOURCES_PROJECT_RESOURCE",
		2: "EVENT_RESOURCES_DATASET_RESOURCE",
		3: "EVENT_RESOURCES_DATASET_VERSION_RESOURCE",
		4: "EVENT_RESOURCES_OBJECT_GROUP_RESOURCE",
	}
	NotificationStreamRequest_EventResources_value = map[string]int32{
		"EVENT_RESOURCES_UNSPECIFIED":              0,
		"EVENT_RESOURCES_PROJECT_RESOURCE":         1,
		"EVENT_RESOURCES_DATASET_RESOURCE":         2,
		"EVENT_RESOURCES_DATASET_VERSION_RESOURCE": 3,
		"EVENT_RESOURCES_OBJECT_GROUP_RESOURCE":    4,
	}
)

func (x NotificationStreamRequest_EventResources) Enum() *NotificationStreamRequest_EventResources {
	p := new(NotificationStreamRequest_EventResources)
	*p = x
	return p
}

func (x NotificationStreamRequest_EventResources) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationStreamRequest_EventResources) Descriptor() protoreflect.EnumDescriptor {
	return file_api_services_v1_notification_service_models_proto_enumTypes[0].Descriptor()
}

func (NotificationStreamRequest_EventResources) Type() protoreflect.EnumType {
	return &file_api_services_v1_notification_service_models_proto_enumTypes[0]
}

func (x NotificationStreamRequest_EventResources) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationStreamRequest_EventResources.Descriptor instead.
func (NotificationStreamRequest_EventResources) EnumDescriptor() ([]byte, []int) {
	return file_api_services_v1_notification_service_models_proto_rawDescGZIP(), []int{2, 0}
}

type EventNotificationMessage_UpdateType int32

const (
	EventNotificationMessage_UPDATE_TYPE_UNSPECIFIED      EventNotificationMessage_UpdateType = 0
	EventNotificationMessage_UPDATE_TYPE_CREATED          EventNotificationMessage_UpdateType = 1
	EventNotificationMessage_UPDATE_TYPE_AVAILABLE        EventNotificationMessage_UpdateType = 2
	EventNotificationMessage_UPDATE_TYPE_UPDATED          EventNotificationMessage_UpdateType = 3
	EventNotificationMessage_UPDATE_TYPE_METADATA_UPDATED EventNotificationMessage_UpdateType = 4
	EventNotificationMessage_UPDATE_TYPE_DELETED          EventNotificationMessage_UpdateType = 5
)

// Enum value maps for EventNotificationMessage_UpdateType.
var (
	EventNotificationMessage_UpdateType_name = map[int32]string{
		0: "UPDATE_TYPE_UNSPECIFIED",
		1: "UPDATE_TYPE_CREATED",
		2: "UPDATE_TYPE_AVAILABLE",
		3: "UPDATE_TYPE_UPDATED",
		4: "UPDATE_TYPE_METADATA_UPDATED",
		5: "UPDATE_TYPE_DELETED",
	}
	EventNotificationMessage_UpdateType_value = map[string]int32{
		"UPDATE_TYPE_UNSPECIFIED":      0,
		"UPDATE_TYPE_CREATED":          1,
		"UPDATE_TYPE_AVAILABLE":        2,
		"UPDATE_TYPE_UPDATED":          3,
		"UPDATE_TYPE_METADATA_UPDATED": 4,
		"UPDATE_TYPE_DELETED":          5,
	}
)

func (x EventNotificationMessage_UpdateType) Enum() *EventNotificationMessage_UpdateType {
	p := new(EventNotificationMessage_UpdateType)
	*p = x
	return p
}

func (x EventNotificationMessage_UpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventNotificationMessage_UpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_services_v1_notification_service_models_proto_enumTypes[1].Descriptor()
}

func (EventNotificationMessage_UpdateType) Type() protoreflect.EnumType {
	return &file_api_services_v1_notification_service_models_proto_enumTypes[1]
}

func (x EventNotificationMessage_UpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventNotificationMessage_UpdateType.Descriptor instead.
func (EventNotificationMessage_UpdateType) EnumDescriptor() ([]byte, []int) {
	return file_api_services_v1_notification_service_models_proto_rawDescGZIP(), []int{7, 0}
}

type CreateEventStreamingGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateEventStreamingGroupRequest) Reset() {
	*x = CreateEventStreamingGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_notification_service_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventStreamingGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventStreamingGroupRequest) ProtoMessage() {}

func (x *CreateEventStreamingGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_notification_service_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventStreamingGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateEventStreamingGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_notification_service_models_proto_rawDescGZIP(), []int{0}
}

type CreateEventStreamingGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queuename string `protobuf:"bytes,1,opt,name=queuename,proto3" json:"queuename,omitempty"`
}

func (x *CreateEventStreamingGroupResponse) Reset() {
	*x = CreateEventStreamingGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_notification_service_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventStreamingGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventStreamingGroupResponse) ProtoMessage() {}

func (x *CreateEventStreamingGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_notification_service_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventStreamingGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateEventStreamingGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_notification_service_models_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventStreamingGroupResponse) GetQueuename() string {
	if x != nil {
		return x.Queuename
	}
	return ""
}

type NotificationStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource           NotificationStreamRequest_EventResources `protobuf:"varint,1,opt,name=resource,proto3,enum=api.services.v1.NotificationStreamRequest_EventResources" json:"resource,omitempty"`
	ResourceId         string                                   `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	IncludeSubresource bool                                     `protobuf:"varint,3,opt,name=include_subresource,json=includeSubresource,proto3" json:"include_subresource,omitempty"`
	// Types that are assignable to StreamType:
	//	*NotificationStreamRequest_StreamAll
	//	*NotificationStreamRequest_StreamFromDate
	//	*NotificationStreamRequest_StreamFromSequence
	StreamType  isNotificationStreamRequest_StreamType `protobuf_oneof:"StreamType"`
	StreamGroup string                                 `protobuf:"bytes,7,opt,name=stream_group,json=streamGroup,proto3" json:"stream_group,omitempty"`
}

func (x *NotificationStreamRequest) Reset() {
	*x = NotificationStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_notification_service_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationStreamRequest) ProtoMessage() {}

func (x *NotificationStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_notification_service_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationStreamRequest.ProtoReflect.Descriptor instead.
func (*NotificationStreamRequest) Descriptor() ([]byte, []int) {
	return file_api_services_v1_notification_service_models_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationStreamRequest) GetResource() NotificationStreamRequest_EventResources {
	if x != nil {
		return x.Resource
	}
	return NotificationStreamRequest_EVENT_RESOURCES_UNSPECIFIED
}

func (x *NotificationStreamRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *NotificationStreamRequest) GetIncludeSubresource() bool {
	if x != nil {
		return x.IncludeSubresource
	}
	return false
}

func (m *NotificationStreamRequest) GetStreamType() isNotificationStreamRequest_StreamType {
	if m != nil {
		return m.StreamType
	}
	return nil
}

func (x *NotificationStreamRequest) GetStreamAll() *StreamAll {
	if x, ok := x.GetStreamType().(*NotificationStreamRequest_StreamAll); ok {
		return x.StreamAll
	}
	return nil
}

func (x *NotificationStreamRequest) GetStreamFromDate() *StreamFromDate {
	if x, ok := x.GetStreamType().(*NotificationStreamRequest_StreamFromDate); ok {
		return x.StreamFromDate
	}
	return nil
}

func (x *NotificationStreamRequest) GetStreamFromSequence() *StreamFromSequence {
	if x, ok := x.GetStreamType().(*NotificationStreamRequest_StreamFromSequence); ok {
		return x.StreamFromSequence
	}
	return nil
}

func (x *NotificationStreamRequest) GetStreamGroup() string {
	if x != nil {
		return x.StreamGroup
	}
	return ""
}

type isNotificationStreamRequest_StreamType interface {
	isNotificationStreamRequest_StreamType()
}

type NotificationStreamRequest_StreamAll struct {
	StreamAll *StreamAll `protobuf:"bytes,4,opt,name=stream_all,json=streamAll,proto3,oneof"`
}

type NotificationStreamRequest_StreamFromDate struct {
	StreamFromDate *StreamFromDate `protobuf:"bytes,5,opt,name=stream_from_date,json=streamFromDate,proto3,oneof"`
}

type NotificationStreamRequest_StreamFromSequence struct {
	StreamFromSequence *StreamFromSequence `protobuf:"bytes,6,opt,name=stream_from_sequence,json=streamFromSequence,proto3,oneof"`
}

func (*NotificationStreamRequest_StreamAll) isNotificationStreamRequest_StreamType() {}

func (*NotificationStreamRequest_StreamFromDate) isNotificationStreamRequest_StreamType() {}

func (*NotificationStreamRequest_StreamFromSequence) isNotificationStreamRequest_StreamType() {}

type StreamFromSequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *StreamFromSequence) Reset() {
	*x = StreamFromSequence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_notification_service_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamFromSequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamFromSequence) ProtoMessage() {}

func (x *StreamFromSequence) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_notification_service_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamFromSequence.ProtoReflect.Descriptor instead.
func (*StreamFromSequence) Descriptor() ([]byte, []int) {
	return file_api_services_v1_notification_service_models_proto_rawDescGZIP(), []int{3}
}

func (x *StreamFromSequence) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type StreamFromDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *StreamFromDate) Reset() {
	*x = StreamFromDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_notification_service_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamFromDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamFromDate) ProtoMessage() {}

func (x *StreamFromDate) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_notification_service_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamFromDate.ProtoReflect.Descriptor instead.
func (*StreamFromDate) Descriptor() ([]byte, []int) {
	return file_api_services_v1_notification_service_models_proto_rawDescGZIP(), []int{4}
}

func (x *StreamFromDate) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type StreamAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamAll) Reset() {
	*x = StreamAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_notification_service_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAll) ProtoMessage() {}

func (x *StreamAll) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_notification_service_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAll.ProtoReflect.Descriptor instead.
func (*StreamAll) Descriptor() ([]byte, []int) {
	return file_api_services_v1_notification_service_models_proto_rawDescGZIP(), []int{5}
}

type NotificationStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *EventNotificationMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Sequence  uint64                    `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Timestamp *timestamppb.Timestamp    `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *NotificationStreamResponse) Reset() {
	*x = NotificationStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_notification_service_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationStreamResponse) ProtoMessage() {}

func (x *NotificationStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_notification_service_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationStreamResponse.ProtoReflect.Descriptor instead.
func (*NotificationStreamResponse) Descriptor() ([]byte, []int) {
	return file_api_services_v1_notification_service_models_proto_rawDescGZIP(), []int{6}
}

func (x *NotificationStreamResponse) GetMessage() *EventNotificationMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *NotificationStreamResponse) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *NotificationStreamResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type EventNotificationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource    v1.Resource                         `protobuf:"varint,1,opt,name=resource,proto3,enum=api.models.v1.Resource" json:"resource,omitempty"`
	ResourceId  string                              `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	UpdatedType EventNotificationMessage_UpdateType `protobuf:"varint,3,opt,name=updated_type,json=updatedType,proto3,enum=api.services.v1.EventNotificationMessage_UpdateType" json:"updated_type,omitempty"`
}

func (x *EventNotificationMessage) Reset() {
	*x = EventNotificationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_v1_notification_service_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventNotificationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventNotificationMessage) ProtoMessage() {}

func (x *EventNotificationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_v1_notification_service_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventNotificationMessage.ProtoReflect.Descriptor instead.
func (*EventNotificationMessage) Descriptor() ([]byte, []int) {
	return file_api_services_v1_notification_service_models_proto_rawDescGZIP(), []int{7}
}

func (x *EventNotificationMessage) GetResource() v1.Resource {
	if x != nil {
		return x.Resource
	}
	return v1.Resource(0)
}

func (x *EventNotificationMessage) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *EventNotificationMessage) GetUpdatedType() EventNotificationMessage_UpdateType {
	if x != nil {
		return x.UpdatedType
	}
	return EventNotificationMessage_UPDATE_TYPE_UNSPECIFIED
}

var File_api_services_v1_notification_service_models_proto protoreflect.FileDescriptor

var file_api_services_v1_notification_service_models_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x21,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xb1, 0x05, 0x0a, 0x19, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x55, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x39, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x5f, 0x73, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x75, 0x62, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x41, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x41, 0x6c, 0x6c, 0x12, 0x4b, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x48, 0x00,
	0x52, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x57, 0x0a, 0x14, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x12, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x72, 0x6f,
	0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0xd6, 0x01, 0x0a,
	0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x1b, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x24, 0x0a, 0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x53, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x53, 0x45,
	0x54, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x02, 0x12, 0x2c, 0x0a, 0x28,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x53, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x53, 0x45, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x03, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x53, 0x5f, 0x4f, 0x42,
	0x4a, 0x45, 0x43, 0x54, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x10, 0x04, 0x42, 0x0c, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x30, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x72, 0x6f,
	0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46,
	0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x6c, 0x6c, 0x22, 0xb7,
	0x01, 0x0a, 0x1a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xfd, 0x02, 0x0a, 0x18, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x57, 0x0a, 0x0c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x17, 0x0a, 0x13, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x20, 0x0a,
	0x1c, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54,
	0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x17, 0x0a, 0x13, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x05, 0x42, 0x87, 0x01, 0x0a, 0x34, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x42, 0x19, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x50, 0x01, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x44, 0x42, 0x2f, 0x67, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_services_v1_notification_service_models_proto_rawDescOnce sync.Once
	file_api_services_v1_notification_service_models_proto_rawDescData = file_api_services_v1_notification_service_models_proto_rawDesc
)

func file_api_services_v1_notification_service_models_proto_rawDescGZIP() []byte {
	file_api_services_v1_notification_service_models_proto_rawDescOnce.Do(func() {
		file_api_services_v1_notification_service_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_v1_notification_service_models_proto_rawDescData)
	})
	return file_api_services_v1_notification_service_models_proto_rawDescData
}

var file_api_services_v1_notification_service_models_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_services_v1_notification_service_models_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_services_v1_notification_service_models_proto_goTypes = []interface{}{
	(NotificationStreamRequest_EventResources)(0), // 0: api.services.v1.NotificationStreamRequest.EventResources
	(EventNotificationMessage_UpdateType)(0),      // 1: api.services.v1.EventNotificationMessage.UpdateType
	(*CreateEventStreamingGroupRequest)(nil),      // 2: api.services.v1.CreateEventStreamingGroupRequest
	(*CreateEventStreamingGroupResponse)(nil),     // 3: api.services.v1.CreateEventStreamingGroupResponse
	(*NotificationStreamRequest)(nil),             // 4: api.services.v1.NotificationStreamRequest
	(*StreamFromSequence)(nil),                    // 5: api.services.v1.StreamFromSequence
	(*StreamFromDate)(nil),                        // 6: api.services.v1.StreamFromDate
	(*StreamAll)(nil),                             // 7: api.services.v1.StreamAll
	(*NotificationStreamResponse)(nil),            // 8: api.services.v1.NotificationStreamResponse
	(*EventNotificationMessage)(nil),              // 9: api.services.v1.EventNotificationMessage
	(*timestamppb.Timestamp)(nil),                 // 10: google.protobuf.Timestamp
	(v1.Resource)(0),                              // 11: api.models.v1.Resource
}
var file_api_services_v1_notification_service_models_proto_depIdxs = []int32{
	0,  // 0: api.services.v1.NotificationStreamRequest.resource:type_name -> api.services.v1.NotificationStreamRequest.EventResources
	7,  // 1: api.services.v1.NotificationStreamRequest.stream_all:type_name -> api.services.v1.StreamAll
	6,  // 2: api.services.v1.NotificationStreamRequest.stream_from_date:type_name -> api.services.v1.StreamFromDate
	5,  // 3: api.services.v1.NotificationStreamRequest.stream_from_sequence:type_name -> api.services.v1.StreamFromSequence
	10, // 4: api.services.v1.StreamFromDate.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 5: api.services.v1.NotificationStreamResponse.message:type_name -> api.services.v1.EventNotificationMessage
	10, // 6: api.services.v1.NotificationStreamResponse.timestamp:type_name -> google.protobuf.Timestamp
	11, // 7: api.services.v1.EventNotificationMessage.resource:type_name -> api.models.v1.Resource
	1,  // 8: api.services.v1.EventNotificationMessage.updated_type:type_name -> api.services.v1.EventNotificationMessage.UpdateType
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_services_v1_notification_service_models_proto_init() }
func file_api_services_v1_notification_service_models_proto_init() {
	if File_api_services_v1_notification_service_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_v1_notification_service_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventStreamingGroupRequest); i {
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
		file_api_services_v1_notification_service_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventStreamingGroupResponse); i {
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
		file_api_services_v1_notification_service_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationStreamRequest); i {
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
		file_api_services_v1_notification_service_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamFromSequence); i {
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
		file_api_services_v1_notification_service_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamFromDate); i {
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
		file_api_services_v1_notification_service_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAll); i {
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
		file_api_services_v1_notification_service_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationStreamResponse); i {
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
		file_api_services_v1_notification_service_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventNotificationMessage); i {
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
	file_api_services_v1_notification_service_models_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*NotificationStreamRequest_StreamAll)(nil),
		(*NotificationStreamRequest_StreamFromDate)(nil),
		(*NotificationStreamRequest_StreamFromSequence)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_services_v1_notification_service_models_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_services_v1_notification_service_models_proto_goTypes,
		DependencyIndexes: file_api_services_v1_notification_service_models_proto_depIdxs,
		EnumInfos:         file_api_services_v1_notification_service_models_proto_enumTypes,
		MessageInfos:      file_api_services_v1_notification_service_models_proto_msgTypes,
	}.Build()
	File_api_services_v1_notification_service_models_proto = out.File
	file_api_services_v1_notification_service_models_proto_rawDesc = nil
	file_api_services_v1_notification_service_models_proto_goTypes = nil
	file_api_services_v1_notification_service_models_proto_depIdxs = nil
}
