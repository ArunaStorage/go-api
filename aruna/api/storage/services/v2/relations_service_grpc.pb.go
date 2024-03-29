// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aruna/api/storage/services/v2/relations_service.proto

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
	RelationsService_ModifyRelations_FullMethodName = "/aruna.api.storage.services.v2.RelationsService/ModifyRelations"
	RelationsService_GetHierarchy_FullMethodName    = "/aruna.api.storage.services.v2.RelationsService/GetHierarchy"
)

// RelationsServiceClient is the client API for RelationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationsServiceClient interface {
	// ModifyRelation
	//
	// Status: BETA
	//
	// Add/Remove/Modifies all relation types to / from a resource
	// Works for internal and external relations
	ModifyRelations(ctx context.Context, in *ModifyRelationsRequest, opts ...grpc.CallOption) (*ModifyRelationsResponse, error)
	// GetHierachy
	//
	// Status: BETA
	//
	// Gets all downstream hierarchy relations from a resource
	// results in a tree structure
	GetHierarchy(ctx context.Context, in *GetHierarchyRequest, opts ...grpc.CallOption) (*GetHierarchyResponse, error)
}

type relationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationsServiceClient(cc grpc.ClientConnInterface) RelationsServiceClient {
	return &relationsServiceClient{cc}
}

func (c *relationsServiceClient) ModifyRelations(ctx context.Context, in *ModifyRelationsRequest, opts ...grpc.CallOption) (*ModifyRelationsResponse, error) {
	out := new(ModifyRelationsResponse)
	err := c.cc.Invoke(ctx, RelationsService_ModifyRelations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationsServiceClient) GetHierarchy(ctx context.Context, in *GetHierarchyRequest, opts ...grpc.CallOption) (*GetHierarchyResponse, error) {
	out := new(GetHierarchyResponse)
	err := c.cc.Invoke(ctx, RelationsService_GetHierarchy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationsServiceServer is the server API for RelationsService service.
// All implementations should embed UnimplementedRelationsServiceServer
// for forward compatibility
type RelationsServiceServer interface {
	// ModifyRelation
	//
	// Status: BETA
	//
	// Add/Remove/Modifies all relation types to / from a resource
	// Works for internal and external relations
	ModifyRelations(context.Context, *ModifyRelationsRequest) (*ModifyRelationsResponse, error)
	// GetHierachy
	//
	// Status: BETA
	//
	// Gets all downstream hierarchy relations from a resource
	// results in a tree structure
	GetHierarchy(context.Context, *GetHierarchyRequest) (*GetHierarchyResponse, error)
}

// UnimplementedRelationsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRelationsServiceServer struct {
}

func (UnimplementedRelationsServiceServer) ModifyRelations(context.Context, *ModifyRelationsRequest) (*ModifyRelationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyRelations not implemented")
}
func (UnimplementedRelationsServiceServer) GetHierarchy(context.Context, *GetHierarchyRequest) (*GetHierarchyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHierarchy not implemented")
}

// UnsafeRelationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationsServiceServer will
// result in compilation errors.
type UnsafeRelationsServiceServer interface {
	mustEmbedUnimplementedRelationsServiceServer()
}

func RegisterRelationsServiceServer(s grpc.ServiceRegistrar, srv RelationsServiceServer) {
	s.RegisterService(&RelationsService_ServiceDesc, srv)
}

func _RelationsService_ModifyRelations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyRelationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationsServiceServer).ModifyRelations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationsService_ModifyRelations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationsServiceServer).ModifyRelations(ctx, req.(*ModifyRelationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationsService_GetHierarchy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHierarchyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationsServiceServer).GetHierarchy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationsService_GetHierarchy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationsServiceServer).GetHierarchy(ctx, req.(*GetHierarchyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationsService_ServiceDesc is the grpc.ServiceDesc for RelationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aruna.api.storage.services.v2.RelationsService",
	HandlerType: (*RelationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModifyRelations",
			Handler:    _RelationsService_ModifyRelations_Handler,
		},
		{
			MethodName: "GetHierarchy",
			Handler:    _RelationsService_GetHierarchy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aruna/api/storage/services/v2/relations_service.proto",
}
