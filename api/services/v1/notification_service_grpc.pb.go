// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// UpdateNotificationServiceClient is the client API for UpdateNotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateNotificationServiceClient interface {
	CreateEventStreamingGroup(ctx context.Context, in *CreateEventStreamingGroupRequest, opts ...grpc.CallOption) (*CreateEventStreamingGroupResponse, error)
	NotificationStream(ctx context.Context, in *NotificationStreamRequest, opts ...grpc.CallOption) (UpdateNotificationService_NotificationStreamClient, error)
}

type updateNotificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateNotificationServiceClient(cc grpc.ClientConnInterface) UpdateNotificationServiceClient {
	return &updateNotificationServiceClient{cc}
}

func (c *updateNotificationServiceClient) CreateEventStreamingGroup(ctx context.Context, in *CreateEventStreamingGroupRequest, opts ...grpc.CallOption) (*CreateEventStreamingGroupResponse, error) {
	out := new(CreateEventStreamingGroupResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.UpdateNotificationService/CreateEventStreamingGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateNotificationServiceClient) NotificationStream(ctx context.Context, in *NotificationStreamRequest, opts ...grpc.CallOption) (UpdateNotificationService_NotificationStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UpdateNotificationService_ServiceDesc.Streams[0], "/api.services.v1.UpdateNotificationService/NotificationStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateNotificationServiceNotificationStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpdateNotificationService_NotificationStreamClient interface {
	Recv() (*NotificationStreamResponse, error)
	grpc.ClientStream
}

type updateNotificationServiceNotificationStreamClient struct {
	grpc.ClientStream
}

func (x *updateNotificationServiceNotificationStreamClient) Recv() (*NotificationStreamResponse, error) {
	m := new(NotificationStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UpdateNotificationServiceServer is the server API for UpdateNotificationService service.
// All implementations should embed UnimplementedUpdateNotificationServiceServer
// for forward compatibility
type UpdateNotificationServiceServer interface {
	CreateEventStreamingGroup(context.Context, *CreateEventStreamingGroupRequest) (*CreateEventStreamingGroupResponse, error)
	NotificationStream(*NotificationStreamRequest, UpdateNotificationService_NotificationStreamServer) error
}

// UnimplementedUpdateNotificationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUpdateNotificationServiceServer struct {
}

func (UnimplementedUpdateNotificationServiceServer) CreateEventStreamingGroup(context.Context, *CreateEventStreamingGroupRequest) (*CreateEventStreamingGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEventStreamingGroup not implemented")
}
func (UnimplementedUpdateNotificationServiceServer) NotificationStream(*NotificationStreamRequest, UpdateNotificationService_NotificationStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method NotificationStream not implemented")
}

// UnsafeUpdateNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateNotificationServiceServer will
// result in compilation errors.
type UnsafeUpdateNotificationServiceServer interface {
	mustEmbedUnimplementedUpdateNotificationServiceServer()
}

func RegisterUpdateNotificationServiceServer(s grpc.ServiceRegistrar, srv UpdateNotificationServiceServer) {
	s.RegisterService(&UpdateNotificationService_ServiceDesc, srv)
}

func _UpdateNotificationService_CreateEventStreamingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventStreamingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateNotificationServiceServer).CreateEventStreamingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.UpdateNotificationService/CreateEventStreamingGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateNotificationServiceServer).CreateEventStreamingGroup(ctx, req.(*CreateEventStreamingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpdateNotificationService_NotificationStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdateNotificationServiceServer).NotificationStream(m, &updateNotificationServiceNotificationStreamServer{stream})
}

type UpdateNotificationService_NotificationStreamServer interface {
	Send(*NotificationStreamResponse) error
	grpc.ServerStream
}

type updateNotificationServiceNotificationStreamServer struct {
	grpc.ServerStream
}

func (x *updateNotificationServiceNotificationStreamServer) Send(m *NotificationStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// UpdateNotificationService_ServiceDesc is the grpc.ServiceDesc for UpdateNotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdateNotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.services.v1.UpdateNotificationService",
	HandlerType: (*UpdateNotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEventStreamingGroup",
			Handler:    _UpdateNotificationService_CreateEventStreamingGroup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NotificationStream",
			Handler:       _UpdateNotificationService_NotificationStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/services/v1/notification_service.proto",
}
