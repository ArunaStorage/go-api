// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: aruna/api/storage/services/v1/objectgroup_service.proto

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

// ObjectGroupServiceClient is the client API for ObjectGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectGroupServiceClient interface {
	// CreateObjectGroup creates a new ObjectGroup in the collection
	CreateObjectGroup(ctx context.Context, in *CreateObjectGroupRequest, opts ...grpc.CallOption) (*CreateObjectGroupResponse, error)
	// UpdateObjectGroup creates an updated ObjectGroup
	// ObjectGroups are immutable
	// Updating an ObjectGroup will create a new Revision of the ObjectGroup
	UpdateObjectGroup(ctx context.Context, in *UpdateObjectGroupRequest, opts ...grpc.CallOption) (*UpdateObjectGroupResponse, error)
	// GetObjectGroupById gets a specific ObjectGroup by ID
	// By default the latest revision is always returned, older revisions need to
	// be specified separately
	GetObjectGroupById(ctx context.Context, in *GetObjectGroupByIdRequest, opts ...grpc.CallOption) (*GetObjectGroupByIdResponse, error)
	// GetObjectGroupsFromObject gets all ObjectGroups associated to a specific
	// Object Objects can be part of multiple ObjectGroups at once
	GetObjectGroupsFromObject(ctx context.Context, in *GetObjectGroupsFromObjectRequest, opts ...grpc.CallOption) (*GetObjectGroupsFromObjectResponse, error)
	// GetObjectGroups is a request that returns a (paginated) list of
	// ObjectGroups that contain a specific set of labels.
	GetObjectGroups(ctx context.Context, in *GetObjectGroupsRequest, opts ...grpc.CallOption) (*GetObjectGroupsResponse, error)
	GetObjectGroupHistory(ctx context.Context, in *GetObjectGroupHistoryRequest, opts ...grpc.CallOption) (*GetObjectGroupHistoryResponse, error)
	GetObjectGroupObjects(ctx context.Context, in *GetObjectGroupObjectsRequest, opts ...grpc.CallOption) (*GetObjectGroupObjectsResponse, error)
	// DeleteObjectGroup is a request that deletes a specified ObjectGroup
	// This does not delete the associated Objects
	DeleteObjectGroup(ctx context.Context, in *DeleteObjectGroupRequest, opts ...grpc.CallOption) (*DeleteObjectGroupResponse, error)
}

type objectGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectGroupServiceClient(cc grpc.ClientConnInterface) ObjectGroupServiceClient {
	return &objectGroupServiceClient{cc}
}

func (c *objectGroupServiceClient) CreateObjectGroup(ctx context.Context, in *CreateObjectGroupRequest, opts ...grpc.CallOption) (*CreateObjectGroupResponse, error) {
	out := new(CreateObjectGroupResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectGroupService/CreateObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectGroupServiceClient) UpdateObjectGroup(ctx context.Context, in *UpdateObjectGroupRequest, opts ...grpc.CallOption) (*UpdateObjectGroupResponse, error) {
	out := new(UpdateObjectGroupResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectGroupService/UpdateObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectGroupServiceClient) GetObjectGroupById(ctx context.Context, in *GetObjectGroupByIdRequest, opts ...grpc.CallOption) (*GetObjectGroupByIdResponse, error) {
	out := new(GetObjectGroupByIdResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectGroupService/GetObjectGroupById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectGroupServiceClient) GetObjectGroupsFromObject(ctx context.Context, in *GetObjectGroupsFromObjectRequest, opts ...grpc.CallOption) (*GetObjectGroupsFromObjectResponse, error) {
	out := new(GetObjectGroupsFromObjectResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectGroupService/GetObjectGroupsFromObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectGroupServiceClient) GetObjectGroups(ctx context.Context, in *GetObjectGroupsRequest, opts ...grpc.CallOption) (*GetObjectGroupsResponse, error) {
	out := new(GetObjectGroupsResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectGroupService/GetObjectGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectGroupServiceClient) GetObjectGroupHistory(ctx context.Context, in *GetObjectGroupHistoryRequest, opts ...grpc.CallOption) (*GetObjectGroupHistoryResponse, error) {
	out := new(GetObjectGroupHistoryResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectGroupService/GetObjectGroupHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectGroupServiceClient) GetObjectGroupObjects(ctx context.Context, in *GetObjectGroupObjectsRequest, opts ...grpc.CallOption) (*GetObjectGroupObjectsResponse, error) {
	out := new(GetObjectGroupObjectsResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectGroupService/GetObjectGroupObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectGroupServiceClient) DeleteObjectGroup(ctx context.Context, in *DeleteObjectGroupRequest, opts ...grpc.CallOption) (*DeleteObjectGroupResponse, error) {
	out := new(DeleteObjectGroupResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectGroupService/DeleteObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectGroupServiceServer is the server API for ObjectGroupService service.
// All implementations should embed UnimplementedObjectGroupServiceServer
// for forward compatibility
type ObjectGroupServiceServer interface {
	// CreateObjectGroup creates a new ObjectGroup in the collection
	CreateObjectGroup(context.Context, *CreateObjectGroupRequest) (*CreateObjectGroupResponse, error)
	// UpdateObjectGroup creates an updated ObjectGroup
	// ObjectGroups are immutable
	// Updating an ObjectGroup will create a new Revision of the ObjectGroup
	UpdateObjectGroup(context.Context, *UpdateObjectGroupRequest) (*UpdateObjectGroupResponse, error)
	// GetObjectGroupById gets a specific ObjectGroup by ID
	// By default the latest revision is always returned, older revisions need to
	// be specified separately
	GetObjectGroupById(context.Context, *GetObjectGroupByIdRequest) (*GetObjectGroupByIdResponse, error)
	// GetObjectGroupsFromObject gets all ObjectGroups associated to a specific
	// Object Objects can be part of multiple ObjectGroups at once
	GetObjectGroupsFromObject(context.Context, *GetObjectGroupsFromObjectRequest) (*GetObjectGroupsFromObjectResponse, error)
	// GetObjectGroups is a request that returns a (paginated) list of
	// ObjectGroups that contain a specific set of labels.
	GetObjectGroups(context.Context, *GetObjectGroupsRequest) (*GetObjectGroupsResponse, error)
	GetObjectGroupHistory(context.Context, *GetObjectGroupHistoryRequest) (*GetObjectGroupHistoryResponse, error)
	GetObjectGroupObjects(context.Context, *GetObjectGroupObjectsRequest) (*GetObjectGroupObjectsResponse, error)
	// DeleteObjectGroup is a request that deletes a specified ObjectGroup
	// This does not delete the associated Objects
	DeleteObjectGroup(context.Context, *DeleteObjectGroupRequest) (*DeleteObjectGroupResponse, error)
}

// UnimplementedObjectGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedObjectGroupServiceServer struct {
}

func (UnimplementedObjectGroupServiceServer) CreateObjectGroup(context.Context, *CreateObjectGroupRequest) (*CreateObjectGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectGroup not implemented")
}
func (UnimplementedObjectGroupServiceServer) UpdateObjectGroup(context.Context, *UpdateObjectGroupRequest) (*UpdateObjectGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObjectGroup not implemented")
}
func (UnimplementedObjectGroupServiceServer) GetObjectGroupById(context.Context, *GetObjectGroupByIdRequest) (*GetObjectGroupByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroupById not implemented")
}
func (UnimplementedObjectGroupServiceServer) GetObjectGroupsFromObject(context.Context, *GetObjectGroupsFromObjectRequest) (*GetObjectGroupsFromObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroupsFromObject not implemented")
}
func (UnimplementedObjectGroupServiceServer) GetObjectGroups(context.Context, *GetObjectGroupsRequest) (*GetObjectGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroups not implemented")
}
func (UnimplementedObjectGroupServiceServer) GetObjectGroupHistory(context.Context, *GetObjectGroupHistoryRequest) (*GetObjectGroupHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroupHistory not implemented")
}
func (UnimplementedObjectGroupServiceServer) GetObjectGroupObjects(context.Context, *GetObjectGroupObjectsRequest) (*GetObjectGroupObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroupObjects not implemented")
}
func (UnimplementedObjectGroupServiceServer) DeleteObjectGroup(context.Context, *DeleteObjectGroupRequest) (*DeleteObjectGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjectGroup not implemented")
}

// UnsafeObjectGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectGroupServiceServer will
// result in compilation errors.
type UnsafeObjectGroupServiceServer interface {
	mustEmbedUnimplementedObjectGroupServiceServer()
}

func RegisterObjectGroupServiceServer(s grpc.ServiceRegistrar, srv ObjectGroupServiceServer) {
	s.RegisterService(&ObjectGroupService_ServiceDesc, srv)
}

func _ObjectGroupService_CreateObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectGroupServiceServer).CreateObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectGroupService/CreateObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectGroupServiceServer).CreateObjectGroup(ctx, req.(*CreateObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectGroupService_UpdateObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectGroupServiceServer).UpdateObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectGroupService/UpdateObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectGroupServiceServer).UpdateObjectGroup(ctx, req.(*UpdateObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectGroupService_GetObjectGroupById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectGroupServiceServer).GetObjectGroupById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectGroupService/GetObjectGroupById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectGroupServiceServer).GetObjectGroupById(ctx, req.(*GetObjectGroupByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectGroupService_GetObjectGroupsFromObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupsFromObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectGroupServiceServer).GetObjectGroupsFromObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectGroupService/GetObjectGroupsFromObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectGroupServiceServer).GetObjectGroupsFromObject(ctx, req.(*GetObjectGroupsFromObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectGroupService_GetObjectGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectGroupServiceServer).GetObjectGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectGroupService/GetObjectGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectGroupServiceServer).GetObjectGroups(ctx, req.(*GetObjectGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectGroupService_GetObjectGroupHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectGroupServiceServer).GetObjectGroupHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectGroupService/GetObjectGroupHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectGroupServiceServer).GetObjectGroupHistory(ctx, req.(*GetObjectGroupHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectGroupService_GetObjectGroupObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectGroupServiceServer).GetObjectGroupObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectGroupService/GetObjectGroupObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectGroupServiceServer).GetObjectGroupObjects(ctx, req.(*GetObjectGroupObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectGroupService_DeleteObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectGroupServiceServer).DeleteObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectGroupService/DeleteObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectGroupServiceServer).DeleteObjectGroup(ctx, req.(*DeleteObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectGroupService_ServiceDesc is the grpc.ServiceDesc for ObjectGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aruna.api.storage.services.v1.ObjectGroupService",
	HandlerType: (*ObjectGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObjectGroup",
			Handler:    _ObjectGroupService_CreateObjectGroup_Handler,
		},
		{
			MethodName: "UpdateObjectGroup",
			Handler:    _ObjectGroupService_UpdateObjectGroup_Handler,
		},
		{
			MethodName: "GetObjectGroupById",
			Handler:    _ObjectGroupService_GetObjectGroupById_Handler,
		},
		{
			MethodName: "GetObjectGroupsFromObject",
			Handler:    _ObjectGroupService_GetObjectGroupsFromObject_Handler,
		},
		{
			MethodName: "GetObjectGroups",
			Handler:    _ObjectGroupService_GetObjectGroups_Handler,
		},
		{
			MethodName: "GetObjectGroupHistory",
			Handler:    _ObjectGroupService_GetObjectGroupHistory_Handler,
		},
		{
			MethodName: "GetObjectGroupObjects",
			Handler:    _ObjectGroupService_GetObjectGroupObjects_Handler,
		},
		{
			MethodName: "DeleteObjectGroup",
			Handler:    _ObjectGroupService_DeleteObjectGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aruna/api/storage/services/v1/objectgroup_service.proto",
}