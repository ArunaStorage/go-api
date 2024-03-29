// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aruna/api/storage/services/v2/license_service.proto

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
	LicenseService_CreateLicense_FullMethodName = "/aruna.api.storage.services.v2.LicenseService/CreateLicense"
	LicenseService_GetLicense_FullMethodName    = "/aruna.api.storage.services.v2.LicenseService/GetLicense"
	LicenseService_ListLicenses_FullMethodName  = "/aruna.api.storage.services.v2.LicenseService/ListLicenses"
)

// LicenseServiceClient is the client API for LicenseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LicenseServiceClient interface {
	// CreateLicense
	//
	// Status: BETA
	//
	// This creates a new license
	CreateLicense(ctx context.Context, in *CreateLicenseRequest, opts ...grpc.CallOption) (*CreateLicenseResponse, error)
	// GetLicense
	//
	// Status: BETA
	//
	// This returns the license for a given tag
	GetLicense(ctx context.Context, in *GetLicenseRequest, opts ...grpc.CallOption) (*GetLicenseResponse, error)
	// ListLicenses
	//
	// Status: BETA
	//
	// This returns a list of all licenses
	ListLicenses(ctx context.Context, in *ListLicensesRequest, opts ...grpc.CallOption) (*ListLicensesResponse, error)
}

type licenseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLicenseServiceClient(cc grpc.ClientConnInterface) LicenseServiceClient {
	return &licenseServiceClient{cc}
}

func (c *licenseServiceClient) CreateLicense(ctx context.Context, in *CreateLicenseRequest, opts ...grpc.CallOption) (*CreateLicenseResponse, error) {
	out := new(CreateLicenseResponse)
	err := c.cc.Invoke(ctx, LicenseService_CreateLicense_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseServiceClient) GetLicense(ctx context.Context, in *GetLicenseRequest, opts ...grpc.CallOption) (*GetLicenseResponse, error) {
	out := new(GetLicenseResponse)
	err := c.cc.Invoke(ctx, LicenseService_GetLicense_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseServiceClient) ListLicenses(ctx context.Context, in *ListLicensesRequest, opts ...grpc.CallOption) (*ListLicensesResponse, error) {
	out := new(ListLicensesResponse)
	err := c.cc.Invoke(ctx, LicenseService_ListLicenses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LicenseServiceServer is the server API for LicenseService service.
// All implementations should embed UnimplementedLicenseServiceServer
// for forward compatibility
type LicenseServiceServer interface {
	// CreateLicense
	//
	// Status: BETA
	//
	// This creates a new license
	CreateLicense(context.Context, *CreateLicenseRequest) (*CreateLicenseResponse, error)
	// GetLicense
	//
	// Status: BETA
	//
	// This returns the license for a given tag
	GetLicense(context.Context, *GetLicenseRequest) (*GetLicenseResponse, error)
	// ListLicenses
	//
	// Status: BETA
	//
	// This returns a list of all licenses
	ListLicenses(context.Context, *ListLicensesRequest) (*ListLicensesResponse, error)
}

// UnimplementedLicenseServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLicenseServiceServer struct {
}

func (UnimplementedLicenseServiceServer) CreateLicense(context.Context, *CreateLicenseRequest) (*CreateLicenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLicense not implemented")
}
func (UnimplementedLicenseServiceServer) GetLicense(context.Context, *GetLicenseRequest) (*GetLicenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLicense not implemented")
}
func (UnimplementedLicenseServiceServer) ListLicenses(context.Context, *ListLicensesRequest) (*ListLicensesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLicenses not implemented")
}

// UnsafeLicenseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LicenseServiceServer will
// result in compilation errors.
type UnsafeLicenseServiceServer interface {
	mustEmbedUnimplementedLicenseServiceServer()
}

func RegisterLicenseServiceServer(s grpc.ServiceRegistrar, srv LicenseServiceServer) {
	s.RegisterService(&LicenseService_ServiceDesc, srv)
}

func _LicenseService_CreateLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServiceServer).CreateLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LicenseService_CreateLicense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServiceServer).CreateLicense(ctx, req.(*CreateLicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LicenseService_GetLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServiceServer).GetLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LicenseService_GetLicense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServiceServer).GetLicense(ctx, req.(*GetLicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LicenseService_ListLicenses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLicensesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServiceServer).ListLicenses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LicenseService_ListLicenses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServiceServer).ListLicenses(ctx, req.(*ListLicensesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LicenseService_ServiceDesc is the grpc.ServiceDesc for LicenseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LicenseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aruna.api.storage.services.v2.LicenseService",
	HandlerType: (*LicenseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLicense",
			Handler:    _LicenseService_CreateLicense_Handler,
		},
		{
			MethodName: "GetLicense",
			Handler:    _LicenseService_GetLicense_Handler,
		},
		{
			MethodName: "ListLicenses",
			Handler:    _LicenseService_ListLicenses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aruna/api/storage/services/v2/license_service.proto",
}
