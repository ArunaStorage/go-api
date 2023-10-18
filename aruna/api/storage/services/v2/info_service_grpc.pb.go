// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aruna/api/storage/services/v2/info_service.proto

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
	StorageStatusService_GetStorageVersion_FullMethodName = "/aruna.api.storage.services.v2.StorageStatusService/GetStorageVersion"
	StorageStatusService_GetStorageStatus_FullMethodName  = "/aruna.api.storage.services.v2.StorageStatusService/GetStorageStatus"
	StorageStatusService_GetPubkeys_FullMethodName        = "/aruna.api.storage.services.v2.StorageStatusService/GetPubkeys"
	StorageStatusService_GetAnouncements_FullMethodName   = "/aruna.api.storage.services.v2.StorageStatusService/GetAnouncements"
	StorageStatusService_SetAnouncements_FullMethodName   = "/aruna.api.storage.services.v2.StorageStatusService/SetAnouncements"
)

// StorageStatusServiceClient is the client API for StorageStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageStatusServiceClient interface {
	// GetStorageVersion
	//
	// Status: BETA
	//
	// A request to get the current version of the server application
	// String representation and https://semver.org/
	GetStorageVersion(ctx context.Context, in *GetStorageVersionRequest, opts ...grpc.CallOption) (*GetStorageVersionResponse, error)
	// GetStorageStatus
	//
	// Status: ALPHA
	//
	// A request to get the current status of the storage components by location(s)
	GetStorageStatus(ctx context.Context, in *GetStorageStatusRequest, opts ...grpc.CallOption) (*GetStorageStatusResponse, error)
	GetPubkeys(ctx context.Context, in *GetPubkeysRequest, opts ...grpc.CallOption) (*GetPubkeysResponse, error)
	GetAnouncements(ctx context.Context, in *GetAnouncementsRequest, opts ...grpc.CallOption) (*GetAnouncementsResponse, error)
	SetAnouncements(ctx context.Context, in *SetAnouncementsRequest, opts ...grpc.CallOption) (*SetAnouncementsResponse, error)
}

type storageStatusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageStatusServiceClient(cc grpc.ClientConnInterface) StorageStatusServiceClient {
	return &storageStatusServiceClient{cc}
}

func (c *storageStatusServiceClient) GetStorageVersion(ctx context.Context, in *GetStorageVersionRequest, opts ...grpc.CallOption) (*GetStorageVersionResponse, error) {
	out := new(GetStorageVersionResponse)
	err := c.cc.Invoke(ctx, StorageStatusService_GetStorageVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageStatusServiceClient) GetStorageStatus(ctx context.Context, in *GetStorageStatusRequest, opts ...grpc.CallOption) (*GetStorageStatusResponse, error) {
	out := new(GetStorageStatusResponse)
	err := c.cc.Invoke(ctx, StorageStatusService_GetStorageStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageStatusServiceClient) GetPubkeys(ctx context.Context, in *GetPubkeysRequest, opts ...grpc.CallOption) (*GetPubkeysResponse, error) {
	out := new(GetPubkeysResponse)
	err := c.cc.Invoke(ctx, StorageStatusService_GetPubkeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageStatusServiceClient) GetAnouncements(ctx context.Context, in *GetAnouncementsRequest, opts ...grpc.CallOption) (*GetAnouncementsResponse, error) {
	out := new(GetAnouncementsResponse)
	err := c.cc.Invoke(ctx, StorageStatusService_GetAnouncements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageStatusServiceClient) SetAnouncements(ctx context.Context, in *SetAnouncementsRequest, opts ...grpc.CallOption) (*SetAnouncementsResponse, error) {
	out := new(SetAnouncementsResponse)
	err := c.cc.Invoke(ctx, StorageStatusService_SetAnouncements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageStatusServiceServer is the server API for StorageStatusService service.
// All implementations should embed UnimplementedStorageStatusServiceServer
// for forward compatibility
type StorageStatusServiceServer interface {
	// GetStorageVersion
	//
	// Status: BETA
	//
	// A request to get the current version of the server application
	// String representation and https://semver.org/
	GetStorageVersion(context.Context, *GetStorageVersionRequest) (*GetStorageVersionResponse, error)
	// GetStorageStatus
	//
	// Status: ALPHA
	//
	// A request to get the current status of the storage components by location(s)
	GetStorageStatus(context.Context, *GetStorageStatusRequest) (*GetStorageStatusResponse, error)
	GetPubkeys(context.Context, *GetPubkeysRequest) (*GetPubkeysResponse, error)
	GetAnouncements(context.Context, *GetAnouncementsRequest) (*GetAnouncementsResponse, error)
	SetAnouncements(context.Context, *SetAnouncementsRequest) (*SetAnouncementsResponse, error)
}

// UnimplementedStorageStatusServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStorageStatusServiceServer struct {
}

func (UnimplementedStorageStatusServiceServer) GetStorageVersion(context.Context, *GetStorageVersionRequest) (*GetStorageVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorageVersion not implemented")
}
func (UnimplementedStorageStatusServiceServer) GetStorageStatus(context.Context, *GetStorageStatusRequest) (*GetStorageStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorageStatus not implemented")
}
func (UnimplementedStorageStatusServiceServer) GetPubkeys(context.Context, *GetPubkeysRequest) (*GetPubkeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPubkeys not implemented")
}
func (UnimplementedStorageStatusServiceServer) GetAnouncements(context.Context, *GetAnouncementsRequest) (*GetAnouncementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnouncements not implemented")
}
func (UnimplementedStorageStatusServiceServer) SetAnouncements(context.Context, *SetAnouncementsRequest) (*SetAnouncementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAnouncements not implemented")
}

// UnsafeStorageStatusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageStatusServiceServer will
// result in compilation errors.
type UnsafeStorageStatusServiceServer interface {
	mustEmbedUnimplementedStorageStatusServiceServer()
}

func RegisterStorageStatusServiceServer(s grpc.ServiceRegistrar, srv StorageStatusServiceServer) {
	s.RegisterService(&StorageStatusService_ServiceDesc, srv)
}

func _StorageStatusService_GetStorageVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageStatusServiceServer).GetStorageVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageStatusService_GetStorageVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageStatusServiceServer).GetStorageVersion(ctx, req.(*GetStorageVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageStatusService_GetStorageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageStatusServiceServer).GetStorageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageStatusService_GetStorageStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageStatusServiceServer).GetStorageStatus(ctx, req.(*GetStorageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageStatusService_GetPubkeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPubkeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageStatusServiceServer).GetPubkeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageStatusService_GetPubkeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageStatusServiceServer).GetPubkeys(ctx, req.(*GetPubkeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageStatusService_GetAnouncements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnouncementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageStatusServiceServer).GetAnouncements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageStatusService_GetAnouncements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageStatusServiceServer).GetAnouncements(ctx, req.(*GetAnouncementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageStatusService_SetAnouncements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAnouncementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageStatusServiceServer).SetAnouncements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageStatusService_SetAnouncements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageStatusServiceServer).SetAnouncements(ctx, req.(*SetAnouncementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageStatusService_ServiceDesc is the grpc.ServiceDesc for StorageStatusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageStatusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aruna.api.storage.services.v2.StorageStatusService",
	HandlerType: (*StorageStatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStorageVersion",
			Handler:    _StorageStatusService_GetStorageVersion_Handler,
		},
		{
			MethodName: "GetStorageStatus",
			Handler:    _StorageStatusService_GetStorageStatus_Handler,
		},
		{
			MethodName: "GetPubkeys",
			Handler:    _StorageStatusService_GetPubkeys_Handler,
		},
		{
			MethodName: "GetAnouncements",
			Handler:    _StorageStatusService_GetAnouncements_Handler,
		},
		{
			MethodName: "SetAnouncements",
			Handler:    _StorageStatusService_SetAnouncements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aruna/api/storage/services/v2/info_service.proto",
}