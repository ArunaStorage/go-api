// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aruna/api/storage/services/v2/authorization_service.proto

package servicesv2

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
	AuthorizationService_CreateAuthorization_FullMethodName = "/aruna.api.storage.services.v2.AuthorizationService/CreateAuthorization"
	AuthorizationService_GetAuthorizations_FullMethodName   = "/aruna.api.storage.services.v2.AuthorizationService/GetAuthorizations"
	AuthorizationService_DeleteAuthorization_FullMethodName = "/aruna.api.storage.services.v2.AuthorizationService/DeleteAuthorization"
	AuthorizationService_UpdateAuthorization_FullMethodName = "/aruna.api.storage.services.v2.AuthorizationService/UpdateAuthorization"
)

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	// CreateAuthorization
	//
	// Status: BETA
	//
	// This creates a user-specific attribute that handles permission for a
	// specific resource
	CreateAuthorization(ctx context.Context, in *CreateAuthorizationRequest, opts ...grpc.CallOption) (*CreateAuthorizationResponse, error)
	// GetAuthorization
	//
	// Status: BETA
	//
	// This gets resource specific user authorizations
	GetAuthorizations(ctx context.Context, in *GetAuthorizationsRequest, opts ...grpc.CallOption) (*GetAuthorizationsResponse, error)
	// DeleteAuthorization
	//
	// Status: BETA
	//
	// This creates a user-specific attribute that handles permission for a
	// specific resource
	DeleteAuthorization(ctx context.Context, in *DeleteAuthorizationRequest, opts ...grpc.CallOption) (*DeleteAuthorizationResponse, error)
	// UpdateAuthorization
	//
	// Status: BETA
	//
	// This creates a user-specific attribute that handles permission for a
	// specific resource
	UpdateAuthorization(ctx context.Context, in *UpdateAuthorizationRequest, opts ...grpc.CallOption) (*UpdateAuthorizationResponse, error)
}

type authorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) CreateAuthorization(ctx context.Context, in *CreateAuthorizationRequest, opts ...grpc.CallOption) (*CreateAuthorizationResponse, error) {
	out := new(CreateAuthorizationResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_CreateAuthorization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetAuthorizations(ctx context.Context, in *GetAuthorizationsRequest, opts ...grpc.CallOption) (*GetAuthorizationsResponse, error) {
	out := new(GetAuthorizationsResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_GetAuthorizations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) DeleteAuthorization(ctx context.Context, in *DeleteAuthorizationRequest, opts ...grpc.CallOption) (*DeleteAuthorizationResponse, error) {
	out := new(DeleteAuthorizationResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_DeleteAuthorization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) UpdateAuthorization(ctx context.Context, in *UpdateAuthorizationRequest, opts ...grpc.CallOption) (*UpdateAuthorizationResponse, error) {
	out := new(UpdateAuthorizationResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_UpdateAuthorization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
// All implementations should embed UnimplementedAuthorizationServiceServer
// for forward compatibility
type AuthorizationServiceServer interface {
	// CreateAuthorization
	//
	// Status: BETA
	//
	// This creates a user-specific attribute that handles permission for a
	// specific resource
	CreateAuthorization(context.Context, *CreateAuthorizationRequest) (*CreateAuthorizationResponse, error)
	// GetAuthorization
	//
	// Status: BETA
	//
	// This gets resource specific user authorizations
	GetAuthorizations(context.Context, *GetAuthorizationsRequest) (*GetAuthorizationsResponse, error)
	// DeleteAuthorization
	//
	// Status: BETA
	//
	// This creates a user-specific attribute that handles permission for a
	// specific resource
	DeleteAuthorization(context.Context, *DeleteAuthorizationRequest) (*DeleteAuthorizationResponse, error)
	// UpdateAuthorization
	//
	// Status: BETA
	//
	// This creates a user-specific attribute that handles permission for a
	// specific resource
	UpdateAuthorization(context.Context, *UpdateAuthorizationRequest) (*UpdateAuthorizationResponse, error)
}

// UnimplementedAuthorizationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (UnimplementedAuthorizationServiceServer) CreateAuthorization(context.Context, *CreateAuthorizationRequest) (*CreateAuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthorization not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetAuthorizations(context.Context, *GetAuthorizationsRequest) (*GetAuthorizationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorizations not implemented")
}
func (UnimplementedAuthorizationServiceServer) DeleteAuthorization(context.Context, *DeleteAuthorizationRequest) (*DeleteAuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthorization not implemented")
}
func (UnimplementedAuthorizationServiceServer) UpdateAuthorization(context.Context, *UpdateAuthorizationRequest) (*UpdateAuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthorization not implemented")
}

// UnsafeAuthorizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServiceServer will
// result in compilation errors.
type UnsafeAuthorizationServiceServer interface {
	mustEmbedUnimplementedAuthorizationServiceServer()
}

func RegisterAuthorizationServiceServer(s grpc.ServiceRegistrar, srv AuthorizationServiceServer) {
	s.RegisterService(&AuthorizationService_ServiceDesc, srv)
}

func _AuthorizationService_CreateAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CreateAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_CreateAuthorization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CreateAuthorization(ctx, req.(*CreateAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetAuthorizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorizationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetAuthorizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_GetAuthorizations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetAuthorizations(ctx, req.(*GetAuthorizationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_DeleteAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).DeleteAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_DeleteAuthorization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).DeleteAuthorization(ctx, req.(*DeleteAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_UpdateAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).UpdateAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_UpdateAuthorization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).UpdateAuthorization(ctx, req.(*UpdateAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorizationService_ServiceDesc is the grpc.ServiceDesc for AuthorizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aruna.api.storage.services.v2.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuthorization",
			Handler:    _AuthorizationService_CreateAuthorization_Handler,
		},
		{
			MethodName: "GetAuthorizations",
			Handler:    _AuthorizationService_GetAuthorizations_Handler,
		},
		{
			MethodName: "DeleteAuthorization",
			Handler:    _AuthorizationService_DeleteAuthorization_Handler,
		},
		{
			MethodName: "UpdateAuthorization",
			Handler:    _AuthorizationService_UpdateAuthorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aruna/api/storage/services/v2/authorization_service.proto",
}
