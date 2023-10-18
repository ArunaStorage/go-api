// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aruna/api/notification/services/v2/notification_service.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EventNotificationService_CreateStreamConsumer_FullMethodName    = "/aruna.api.notification.services.v2.EventNotificationService/CreateStreamConsumer"
	EventNotificationService_GetEventMessageBatch_FullMethodName    = "/aruna.api.notification.services.v2.EventNotificationService/GetEventMessageBatch"
	EventNotificationService_GetEventMessageStream_FullMethodName   = "/aruna.api.notification.services.v2.EventNotificationService/GetEventMessageStream"
	EventNotificationService_AcknowledgeMessageBatch_FullMethodName = "/aruna.api.notification.services.v2.EventNotificationService/AcknowledgeMessageBatch"
	EventNotificationService_DeleteStreamConsumer_FullMethodName    = "/aruna.api.notification.services.v2.EventNotificationService/DeleteStreamConsumer"
)

// EventNotificationServiceClient is the client API for EventNotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventNotificationServiceClient interface {
	// CreateStreamConsumer
	//
	// Creates a new event stream consumer.
	CreateStreamConsumer(ctx context.Context, in *CreateStreamConsumerRequest, opts ...grpc.CallOption) (*CreateStreamConsumerResponse, error)
	// GetEventMessageBatch
	//
	// Reads a set of messages from a given stream group
	// Each message contains a separate acknowledgement message thatis protected by a salt and an hmac for verification.
	// The message can be send directly through the AcknowledgeMessageBatch call to acknowledge the message.
	GetEventMessageBatch(ctx context.Context, in *GetEventMessageBatchRequest, opts ...grpc.CallOption) (*GetEventMessageBatchResponse, error)
	// GetEventMessageBatch
	//
	// Opens a stream which pushes each received notification individual and just-in-time.
	// Each message contains a separate acknowledgement message that is protected by a salt and an hmac for verification.
	// The message can be send directly through the AcknowledgeMessageBatch call to acknowledge the message.
	GetEventMessageStream(ctx context.Context, in *GetEventMessageStreamRequest, opts ...grpc.CallOption) (EventNotificationService_GetEventMessageStreamClient, error)
	// AcknowledgeMessageBatch
	//
	// List of messages to acknowledge
	// Each reply is protected by a salt and and hmac that verifies the message
	AcknowledgeMessageBatch(ctx context.Context, in *AcknowledgeMessageBatchRequest, opts ...grpc.CallOption) (*AcknowledgeMessageBatchResponse, error)
	// DeleteEventStreamingGroup
	//
	// Deletes an existing event stream consumer by ID.
	DeleteStreamConsumer(ctx context.Context, in *DeleteStreamConsumerRequest, opts ...grpc.CallOption) (*DeleteStreamConsumerResponse, error)
}

type eventNotificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventNotificationServiceClient(cc grpc.ClientConnInterface) EventNotificationServiceClient {
	return &eventNotificationServiceClient{cc}
}

func (c *eventNotificationServiceClient) CreateStreamConsumer(ctx context.Context, in *CreateStreamConsumerRequest, opts ...grpc.CallOption) (*CreateStreamConsumerResponse, error) {
	out := new(CreateStreamConsumerResponse)
	err := c.cc.Invoke(ctx, EventNotificationService_CreateStreamConsumer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventNotificationServiceClient) GetEventMessageBatch(ctx context.Context, in *GetEventMessageBatchRequest, opts ...grpc.CallOption) (*GetEventMessageBatchResponse, error) {
	out := new(GetEventMessageBatchResponse)
	err := c.cc.Invoke(ctx, EventNotificationService_GetEventMessageBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventNotificationServiceClient) GetEventMessageStream(ctx context.Context, in *GetEventMessageStreamRequest, opts ...grpc.CallOption) (EventNotificationService_GetEventMessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventNotificationService_ServiceDesc.Streams[0], EventNotificationService_GetEventMessageStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &eventNotificationServiceGetEventMessageStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventNotificationService_GetEventMessageStreamClient interface {
	Recv() (*GetEventMessageStreamResponse, error)
	grpc.ClientStream
}

type eventNotificationServiceGetEventMessageStreamClient struct {
	grpc.ClientStream
}

func (x *eventNotificationServiceGetEventMessageStreamClient) Recv() (*GetEventMessageStreamResponse, error) {
	m := new(GetEventMessageStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventNotificationServiceClient) AcknowledgeMessageBatch(ctx context.Context, in *AcknowledgeMessageBatchRequest, opts ...grpc.CallOption) (*AcknowledgeMessageBatchResponse, error) {
	out := new(AcknowledgeMessageBatchResponse)
	err := c.cc.Invoke(ctx, EventNotificationService_AcknowledgeMessageBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventNotificationServiceClient) DeleteStreamConsumer(ctx context.Context, in *DeleteStreamConsumerRequest, opts ...grpc.CallOption) (*DeleteStreamConsumerResponse, error) {
	out := new(DeleteStreamConsumerResponse)
	err := c.cc.Invoke(ctx, EventNotificationService_DeleteStreamConsumer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventNotificationServiceServer is the server API for EventNotificationService service.
// All implementations should embed UnimplementedEventNotificationServiceServer
// for forward compatibility
type EventNotificationServiceServer interface {
	// CreateStreamConsumer
	//
	// Creates a new event stream consumer.
	CreateStreamConsumer(context.Context, *CreateStreamConsumerRequest) (*CreateStreamConsumerResponse, error)
	// GetEventMessageBatch
	//
	// Reads a set of messages from a given stream group
	// Each message contains a separate acknowledgement message thatis protected by a salt and an hmac for verification.
	// The message can be send directly through the AcknowledgeMessageBatch call to acknowledge the message.
	GetEventMessageBatch(context.Context, *GetEventMessageBatchRequest) (*GetEventMessageBatchResponse, error)
	// GetEventMessageBatch
	//
	// Opens a stream which pushes each received notification individual and just-in-time.
	// Each message contains a separate acknowledgement message that is protected by a salt and an hmac for verification.
	// The message can be send directly through the AcknowledgeMessageBatch call to acknowledge the message.
	GetEventMessageStream(*GetEventMessageStreamRequest, EventNotificationService_GetEventMessageStreamServer) error
	// AcknowledgeMessageBatch
	//
	// List of messages to acknowledge
	// Each reply is protected by a salt and and hmac that verifies the message
	AcknowledgeMessageBatch(context.Context, *AcknowledgeMessageBatchRequest) (*AcknowledgeMessageBatchResponse, error)
	// DeleteEventStreamingGroup
	//
	// Deletes an existing event stream consumer by ID.
	DeleteStreamConsumer(context.Context, *DeleteStreamConsumerRequest) (*DeleteStreamConsumerResponse, error)
}

// UnimplementedEventNotificationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEventNotificationServiceServer struct {
}

func (UnimplementedEventNotificationServiceServer) CreateStreamConsumer(context.Context, *CreateStreamConsumerRequest) (*CreateStreamConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStreamConsumer not implemented")
}
func (UnimplementedEventNotificationServiceServer) GetEventMessageBatch(context.Context, *GetEventMessageBatchRequest) (*GetEventMessageBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventMessageBatch not implemented")
}
func (UnimplementedEventNotificationServiceServer) GetEventMessageStream(*GetEventMessageStreamRequest, EventNotificationService_GetEventMessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEventMessageStream not implemented")
}
func (UnimplementedEventNotificationServiceServer) AcknowledgeMessageBatch(context.Context, *AcknowledgeMessageBatchRequest) (*AcknowledgeMessageBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcknowledgeMessageBatch not implemented")
}
func (UnimplementedEventNotificationServiceServer) DeleteStreamConsumer(context.Context, *DeleteStreamConsumerRequest) (*DeleteStreamConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStreamConsumer not implemented")
}

// UnsafeEventNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventNotificationServiceServer will
// result in compilation errors.
type UnsafeEventNotificationServiceServer interface {
	mustEmbedUnimplementedEventNotificationServiceServer()
}

func RegisterEventNotificationServiceServer(s grpc.ServiceRegistrar, srv EventNotificationServiceServer) {
	s.RegisterService(&EventNotificationService_ServiceDesc, srv)
}

func _EventNotificationService_CreateStreamConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStreamConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventNotificationServiceServer).CreateStreamConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventNotificationService_CreateStreamConsumer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventNotificationServiceServer).CreateStreamConsumer(ctx, req.(*CreateStreamConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventNotificationService_GetEventMessageBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventMessageBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventNotificationServiceServer).GetEventMessageBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventNotificationService_GetEventMessageBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventNotificationServiceServer).GetEventMessageBatch(ctx, req.(*GetEventMessageBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventNotificationService_GetEventMessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEventMessageStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventNotificationServiceServer).GetEventMessageStream(m, &eventNotificationServiceGetEventMessageStreamServer{stream})
}

type EventNotificationService_GetEventMessageStreamServer interface {
	Send(*GetEventMessageStreamResponse) error
	grpc.ServerStream
}

type eventNotificationServiceGetEventMessageStreamServer struct {
	grpc.ServerStream
}

func (x *eventNotificationServiceGetEventMessageStreamServer) Send(m *GetEventMessageStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EventNotificationService_AcknowledgeMessageBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcknowledgeMessageBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventNotificationServiceServer).AcknowledgeMessageBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventNotificationService_AcknowledgeMessageBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventNotificationServiceServer).AcknowledgeMessageBatch(ctx, req.(*AcknowledgeMessageBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventNotificationService_DeleteStreamConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStreamConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventNotificationServiceServer).DeleteStreamConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventNotificationService_DeleteStreamConsumer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventNotificationServiceServer).DeleteStreamConsumer(ctx, req.(*DeleteStreamConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventNotificationService_ServiceDesc is the grpc.ServiceDesc for EventNotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventNotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aruna.api.notification.services.v2.EventNotificationService",
	HandlerType: (*EventNotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStreamConsumer",
			Handler:    _EventNotificationService_CreateStreamConsumer_Handler,
		},
		{
			MethodName: "GetEventMessageBatch",
			Handler:    _EventNotificationService_GetEventMessageBatch_Handler,
		},
		{
			MethodName: "AcknowledgeMessageBatch",
			Handler:    _EventNotificationService_AcknowledgeMessageBatch_Handler,
		},
		{
			MethodName: "DeleteStreamConsumer",
			Handler:    _EventNotificationService_DeleteStreamConsumer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEventMessageStream",
			Handler:       _EventNotificationService_GetEventMessageStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "aruna/api/notification/services/v2/notification_service.proto",
}
