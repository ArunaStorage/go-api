// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aruna/api/storage/services/v2/project_service.proto

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
	ProjectService_CreateProject_FullMethodName            = "/aruna.api.storage.services.v2.ProjectService/CreateProject"
	ProjectService_GetProject_FullMethodName               = "/aruna.api.storage.services.v2.ProjectService/GetProject"
	ProjectService_GetProjects_FullMethodName              = "/aruna.api.storage.services.v2.ProjectService/GetProjects"
	ProjectService_DeleteProject_FullMethodName            = "/aruna.api.storage.services.v2.ProjectService/DeleteProject"
	ProjectService_UpdateProjectName_FullMethodName        = "/aruna.api.storage.services.v2.ProjectService/UpdateProjectName"
	ProjectService_UpdateProjectDescription_FullMethodName = "/aruna.api.storage.services.v2.ProjectService/UpdateProjectDescription"
	ProjectService_UpdateProjectKeyValues_FullMethodName   = "/aruna.api.storage.services.v2.ProjectService/UpdateProjectKeyValues"
	ProjectService_UpdateProjectDataClass_FullMethodName   = "/aruna.api.storage.services.v2.ProjectService/UpdateProjectDataClass"
	ProjectService_ArchiveProject_FullMethodName           = "/aruna.api.storage.services.v2.ProjectService/ArchiveProject"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	// CreateProject
	//
	// Status: BETA
	//
	// Creates a new project. All subsequent resources are part of a project.
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	// GetProject
	//
	// Status: BETA
	//
	// Requests a project (by id)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error)
	// GetProjects
	//
	// Status: BETA
	//
	// Admin request to get all projects
	GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsResponse, error)
	// DeleteProject
	//
	// Status: BETA
	//
	// Deletes the project and all its associated data. Must be empty!
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error)
	// UpdateProjectName
	//
	// Status: BETA
	//
	// Updates the project name. Caveat! Will rename the "s3 bucket" for data proxies!
	UpdateProjectName(ctx context.Context, in *UpdateProjectNameRequest, opts ...grpc.CallOption) (*UpdateProjectNameResponse, error)
	// UpdateProjectDescription
	//
	// Status: BETA
	//
	// Updates the project name.
	UpdateProjectDescription(ctx context.Context, in *UpdateProjectDescriptionRequest, opts ...grpc.CallOption) (*UpdateProjectDescriptionResponse, error)
	// UpdateProjectKeyValues
	//
	// Status: BETA
	//
	// Updates the project key values.
	UpdateProjectKeyValues(ctx context.Context, in *UpdateProjectKeyValuesRequest, opts ...grpc.CallOption) (*UpdateProjectKeyValuesResponse, error)
	// UpdateProjectDataClass
	//
	// Status: BETA
	//
	// Updates the project name. All (meta) data will be overwritten.
	UpdateProjectDataClass(ctx context.Context, in *UpdateProjectDataClassRequest, opts ...grpc.CallOption) (*UpdateProjectDataClassResponse, error)
	// ArchiveProjectRequest
	//
	// Status: BETA
	//
	// Archives the full project, rendering all downstream relations immutable
	ArchiveProject(ctx context.Context, in *ArchiveProjectRequest, opts ...grpc.CallOption) (*ArchiveProjectResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_CreateProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error) {
	out := new(GetProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsResponse, error) {
	out := new(GetProjectsResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetProjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error) {
	out := new(DeleteProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_DeleteProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProjectName(ctx context.Context, in *UpdateProjectNameRequest, opts ...grpc.CallOption) (*UpdateProjectNameResponse, error) {
	out := new(UpdateProjectNameResponse)
	err := c.cc.Invoke(ctx, ProjectService_UpdateProjectName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProjectDescription(ctx context.Context, in *UpdateProjectDescriptionRequest, opts ...grpc.CallOption) (*UpdateProjectDescriptionResponse, error) {
	out := new(UpdateProjectDescriptionResponse)
	err := c.cc.Invoke(ctx, ProjectService_UpdateProjectDescription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProjectKeyValues(ctx context.Context, in *UpdateProjectKeyValuesRequest, opts ...grpc.CallOption) (*UpdateProjectKeyValuesResponse, error) {
	out := new(UpdateProjectKeyValuesResponse)
	err := c.cc.Invoke(ctx, ProjectService_UpdateProjectKeyValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProjectDataClass(ctx context.Context, in *UpdateProjectDataClassRequest, opts ...grpc.CallOption) (*UpdateProjectDataClassResponse, error) {
	out := new(UpdateProjectDataClassResponse)
	err := c.cc.Invoke(ctx, ProjectService_UpdateProjectDataClass_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) ArchiveProject(ctx context.Context, in *ArchiveProjectRequest, opts ...grpc.CallOption) (*ArchiveProjectResponse, error) {
	out := new(ArchiveProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_ArchiveProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations should embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	// CreateProject
	//
	// Status: BETA
	//
	// Creates a new project. All subsequent resources are part of a project.
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	// GetProject
	//
	// Status: BETA
	//
	// Requests a project (by id)
	GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error)
	// GetProjects
	//
	// Status: BETA
	//
	// Admin request to get all projects
	GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsResponse, error)
	// DeleteProject
	//
	// Status: BETA
	//
	// Deletes the project and all its associated data. Must be empty!
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error)
	// UpdateProjectName
	//
	// Status: BETA
	//
	// Updates the project name. Caveat! Will rename the "s3 bucket" for data proxies!
	UpdateProjectName(context.Context, *UpdateProjectNameRequest) (*UpdateProjectNameResponse, error)
	// UpdateProjectDescription
	//
	// Status: BETA
	//
	// Updates the project name.
	UpdateProjectDescription(context.Context, *UpdateProjectDescriptionRequest) (*UpdateProjectDescriptionResponse, error)
	// UpdateProjectKeyValues
	//
	// Status: BETA
	//
	// Updates the project key values.
	UpdateProjectKeyValues(context.Context, *UpdateProjectKeyValuesRequest) (*UpdateProjectKeyValuesResponse, error)
	// UpdateProjectDataClass
	//
	// Status: BETA
	//
	// Updates the project name. All (meta) data will be overwritten.
	UpdateProjectDataClass(context.Context, *UpdateProjectDataClassRequest) (*UpdateProjectDataClassResponse, error)
	// ArchiveProjectRequest
	//
	// Status: BETA
	//
	// Archives the full project, rendering all downstream relations immutable
	ArchiveProject(context.Context, *ArchiveProjectRequest) (*ArchiveProjectResponse, error)
}

// UnimplementedProjectServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectServiceServer) GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectServiceServer) GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjects not implemented")
}
func (UnimplementedProjectServiceServer) DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProjectName(context.Context, *UpdateProjectNameRequest) (*UpdateProjectNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectName not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProjectDescription(context.Context, *UpdateProjectDescriptionRequest) (*UpdateProjectDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectDescription not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProjectKeyValues(context.Context, *UpdateProjectKeyValuesRequest) (*UpdateProjectKeyValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectKeyValues not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProjectDataClass(context.Context, *UpdateProjectDataClassRequest) (*UpdateProjectDataClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectDataClass not implemented")
}
func (UnimplementedProjectServiceServer) ArchiveProject(context.Context, *ArchiveProjectRequest) (*ArchiveProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveProject not implemented")
}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProjects(ctx, req.(*GetProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProjectName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProjectName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_UpdateProjectName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProjectName(ctx, req.(*UpdateProjectNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProjectDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProjectDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_UpdateProjectDescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProjectDescription(ctx, req.(*UpdateProjectDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProjectKeyValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectKeyValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProjectKeyValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_UpdateProjectKeyValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProjectKeyValues(ctx, req.(*UpdateProjectKeyValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProjectDataClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectDataClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProjectDataClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_UpdateProjectDataClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProjectDataClass(ctx, req.(*UpdateProjectDataClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_ArchiveProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ArchiveProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_ArchiveProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ArchiveProject(ctx, req.(*ArchiveProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aruna.api.storage.services.v2.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectService_CreateProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _ProjectService_GetProject_Handler,
		},
		{
			MethodName: "GetProjects",
			Handler:    _ProjectService_GetProjects_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _ProjectService_DeleteProject_Handler,
		},
		{
			MethodName: "UpdateProjectName",
			Handler:    _ProjectService_UpdateProjectName_Handler,
		},
		{
			MethodName: "UpdateProjectDescription",
			Handler:    _ProjectService_UpdateProjectDescription_Handler,
		},
		{
			MethodName: "UpdateProjectKeyValues",
			Handler:    _ProjectService_UpdateProjectKeyValues_Handler,
		},
		{
			MethodName: "UpdateProjectDataClass",
			Handler:    _ProjectService_UpdateProjectDataClass_Handler,
		},
		{
			MethodName: "ArchiveProject",
			Handler:    _ProjectService_ArchiveProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aruna/api/storage/services/v2/project_service.proto",
}