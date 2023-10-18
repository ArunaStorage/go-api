// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aruna/api/storage/services/v2/workspace_service.proto

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
	WorkspaceService_CreateWorkspaceTemplate_FullMethodName     = "/aruna.api.storage.services.v2.WorkspaceService/CreateWorkspaceTemplate"
	WorkspaceService_GetWorkspaceTemplate_FullMethodName        = "/aruna.api.storage.services.v2.WorkspaceService/GetWorkspaceTemplate"
	WorkspaceService_ListOwnedWorkspaceTemplates_FullMethodName = "/aruna.api.storage.services.v2.WorkspaceService/ListOwnedWorkspaceTemplates"
	WorkspaceService_DeleteWorkspaceTemplate_FullMethodName     = "/aruna.api.storage.services.v2.WorkspaceService/DeleteWorkspaceTemplate"
	WorkspaceService_CreateWorkspace_FullMethodName             = "/aruna.api.storage.services.v2.WorkspaceService/CreateWorkspace"
	WorkspaceService_DeleteWorkspace_FullMethodName             = "/aruna.api.storage.services.v2.WorkspaceService/DeleteWorkspace"
	WorkspaceService_ClaimWorkspace_FullMethodName              = "/aruna.api.storage.services.v2.WorkspaceService/ClaimWorkspace"
)

// WorkspaceServiceClient is the client API for WorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceServiceClient interface {
	// CreatesNewWorkspaceTemplate
	//
	// Status: ALPHA
	//
	// This will create a new template for workspaces (admin only)
	CreateWorkspaceTemplate(ctx context.Context, in *CreateWorkspaceTemplateRequest, opts ...grpc.CallOption) (*CreateWorkspaceTemplateResponse, error)
	// GetWorkspaceTemplatesById
	//
	//	Status: ALPHA
	//
	// Gets workspace template by id
	GetWorkspaceTemplate(ctx context.Context, in *GetWorkspaceTemplateRequest, opts ...grpc.CallOption) (*GetWorkspaceTemplateResponse, error)
	// ListOwnedWorkspaceTemplates
	//
	//	Status: ALPHA
	//
	// Lists owned workspace templates
	ListOwnedWorkspaceTemplates(ctx context.Context, in *ListOwnedWorkspaceTemplatesRequest, opts ...grpc.CallOption) (*ListOwnedWorkspaceTemplatesResponse, error)
	// DeleteWorkspaceTemplates
	//
	//	Status: ALPHA
	//
	// Deletes specified workspace templates
	DeleteWorkspaceTemplate(ctx context.Context, in *DeleteWorkspaceTemplateRequest, opts ...grpc.CallOption) (*DeleteWorkspaceTemplateResponse, error)
	// CreateWorkspace
	//
	// Status: ALPHA
	//
	// A new request to create a personal anonymous workspace
	CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*CreateWorkspaceResponse, error)
	// DeleteWorkspace
	//
	// Status: ALPHA
	//
	// Delete a workspace
	DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error)
	// DeleteWorkspace
	//
	// Status: ALPHA
	//
	// Claims an anonymous workspace, and transfers the owner to a regular user account.
	ClaimWorkspace(ctx context.Context, in *ClaimWorkspaceRequest, opts ...grpc.CallOption) (*ClaimWorkspaceResponse, error)
}

type workspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceServiceClient(cc grpc.ClientConnInterface) WorkspaceServiceClient {
	return &workspaceServiceClient{cc}
}

func (c *workspaceServiceClient) CreateWorkspaceTemplate(ctx context.Context, in *CreateWorkspaceTemplateRequest, opts ...grpc.CallOption) (*CreateWorkspaceTemplateResponse, error) {
	out := new(CreateWorkspaceTemplateResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_CreateWorkspaceTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetWorkspaceTemplate(ctx context.Context, in *GetWorkspaceTemplateRequest, opts ...grpc.CallOption) (*GetWorkspaceTemplateResponse, error) {
	out := new(GetWorkspaceTemplateResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_GetWorkspaceTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ListOwnedWorkspaceTemplates(ctx context.Context, in *ListOwnedWorkspaceTemplatesRequest, opts ...grpc.CallOption) (*ListOwnedWorkspaceTemplatesResponse, error) {
	out := new(ListOwnedWorkspaceTemplatesResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_ListOwnedWorkspaceTemplates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteWorkspaceTemplate(ctx context.Context, in *DeleteWorkspaceTemplateRequest, opts ...grpc.CallOption) (*DeleteWorkspaceTemplateResponse, error) {
	out := new(DeleteWorkspaceTemplateResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_DeleteWorkspaceTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*CreateWorkspaceResponse, error) {
	out := new(CreateWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_CreateWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error) {
	out := new(DeleteWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_DeleteWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ClaimWorkspace(ctx context.Context, in *ClaimWorkspaceRequest, opts ...grpc.CallOption) (*ClaimWorkspaceResponse, error) {
	out := new(ClaimWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_ClaimWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServiceServer is the server API for WorkspaceService service.
// All implementations should embed UnimplementedWorkspaceServiceServer
// for forward compatibility
type WorkspaceServiceServer interface {
	// CreatesNewWorkspaceTemplate
	//
	// Status: ALPHA
	//
	// This will create a new template for workspaces (admin only)
	CreateWorkspaceTemplate(context.Context, *CreateWorkspaceTemplateRequest) (*CreateWorkspaceTemplateResponse, error)
	// GetWorkspaceTemplatesById
	//
	//	Status: ALPHA
	//
	// Gets workspace template by id
	GetWorkspaceTemplate(context.Context, *GetWorkspaceTemplateRequest) (*GetWorkspaceTemplateResponse, error)
	// ListOwnedWorkspaceTemplates
	//
	//	Status: ALPHA
	//
	// Lists owned workspace templates
	ListOwnedWorkspaceTemplates(context.Context, *ListOwnedWorkspaceTemplatesRequest) (*ListOwnedWorkspaceTemplatesResponse, error)
	// DeleteWorkspaceTemplates
	//
	//	Status: ALPHA
	//
	// Deletes specified workspace templates
	DeleteWorkspaceTemplate(context.Context, *DeleteWorkspaceTemplateRequest) (*DeleteWorkspaceTemplateResponse, error)
	// CreateWorkspace
	//
	// Status: ALPHA
	//
	// A new request to create a personal anonymous workspace
	CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error)
	// DeleteWorkspace
	//
	// Status: ALPHA
	//
	// Delete a workspace
	DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error)
	// DeleteWorkspace
	//
	// Status: ALPHA
	//
	// Claims an anonymous workspace, and transfers the owner to a regular user account.
	ClaimWorkspace(context.Context, *ClaimWorkspaceRequest) (*ClaimWorkspaceResponse, error)
}

// UnimplementedWorkspaceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServiceServer struct {
}

func (UnimplementedWorkspaceServiceServer) CreateWorkspaceTemplate(context.Context, *CreateWorkspaceTemplateRequest) (*CreateWorkspaceTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspaceTemplate not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetWorkspaceTemplate(context.Context, *GetWorkspaceTemplateRequest) (*GetWorkspaceTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceTemplate not implemented")
}
func (UnimplementedWorkspaceServiceServer) ListOwnedWorkspaceTemplates(context.Context, *ListOwnedWorkspaceTemplatesRequest) (*ListOwnedWorkspaceTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOwnedWorkspaceTemplates not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteWorkspaceTemplate(context.Context, *DeleteWorkspaceTemplateRequest) (*DeleteWorkspaceTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspaceTemplate not implemented")
}
func (UnimplementedWorkspaceServiceServer) CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) ClaimWorkspace(context.Context, *ClaimWorkspaceRequest) (*ClaimWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimWorkspace not implemented")
}

// UnsafeWorkspaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServiceServer will
// result in compilation errors.
type UnsafeWorkspaceServiceServer interface {
	mustEmbedUnimplementedWorkspaceServiceServer()
}

func RegisterWorkspaceServiceServer(s grpc.ServiceRegistrar, srv WorkspaceServiceServer) {
	s.RegisterService(&WorkspaceService_ServiceDesc, srv)
}

func _WorkspaceService_CreateWorkspaceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateWorkspaceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_CreateWorkspaceTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateWorkspaceTemplate(ctx, req.(*CreateWorkspaceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetWorkspaceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspaceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_GetWorkspaceTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspaceTemplate(ctx, req.(*GetWorkspaceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ListOwnedWorkspaceTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOwnedWorkspaceTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ListOwnedWorkspaceTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_ListOwnedWorkspaceTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ListOwnedWorkspaceTemplates(ctx, req.(*ListOwnedWorkspaceTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteWorkspaceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteWorkspaceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_DeleteWorkspaceTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteWorkspaceTemplate(ctx, req.(*DeleteWorkspaceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_CreateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_CreateWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, req.(*CreateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_DeleteWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, req.(*DeleteWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ClaimWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ClaimWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_ClaimWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ClaimWorkspace(ctx, req.(*ClaimWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkspaceService_ServiceDesc is the grpc.ServiceDesc for WorkspaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aruna.api.storage.services.v2.WorkspaceService",
	HandlerType: (*WorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkspaceTemplate",
			Handler:    _WorkspaceService_CreateWorkspaceTemplate_Handler,
		},
		{
			MethodName: "GetWorkspaceTemplate",
			Handler:    _WorkspaceService_GetWorkspaceTemplate_Handler,
		},
		{
			MethodName: "ListOwnedWorkspaceTemplates",
			Handler:    _WorkspaceService_ListOwnedWorkspaceTemplates_Handler,
		},
		{
			MethodName: "DeleteWorkspaceTemplate",
			Handler:    _WorkspaceService_DeleteWorkspaceTemplate_Handler,
		},
		{
			MethodName: "CreateWorkspace",
			Handler:    _WorkspaceService_CreateWorkspace_Handler,
		},
		{
			MethodName: "DeleteWorkspace",
			Handler:    _WorkspaceService_DeleteWorkspace_Handler,
		},
		{
			MethodName: "ClaimWorkspace",
			Handler:    _WorkspaceService_ClaimWorkspace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aruna/api/storage/services/v2/workspace_service.proto",
}