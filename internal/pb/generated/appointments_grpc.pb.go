// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: pb/protobuf/appointments.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AppointmentService_AppointAttender_FullMethodName              = "/AppointmentService/AppointAttender"
	AppointmentService_CreateChildOffice_FullMethodName            = "/AppointmentService/CreateChildOffice"
	AppointmentService_AppointChildOfficeHead_FullMethodName       = "/AppointmentService/AppointChildOfficeHead"
	AppointmentService_AppointChildOfficeDeputyHead_FullMethodName = "/AppointmentService/AppointChildOfficeDeputyHead"
)

// AppointmentServiceClient is the client API for AppointmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppointmentServiceClient interface {
	AppointAttender(ctx context.Context, in *AttenderAppointmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateChildOffice(ctx context.Context, in *CreateChildOfficeRequest, opts ...grpc.CallOption) (*CreateChildOfficeResponse, error)
	AppointChildOfficeHead(ctx context.Context, in *OfficeHeadAppointmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AppointChildOfficeDeputyHead(ctx context.Context, in *OfficeHeadAppointmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type appointmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppointmentServiceClient(cc grpc.ClientConnInterface) AppointmentServiceClient {
	return &appointmentServiceClient{cc}
}

func (c *appointmentServiceClient) AppointAttender(ctx context.Context, in *AttenderAppointmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AppointmentService_AppointAttender_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) CreateChildOffice(ctx context.Context, in *CreateChildOfficeRequest, opts ...grpc.CallOption) (*CreateChildOfficeResponse, error) {
	out := new(CreateChildOfficeResponse)
	err := c.cc.Invoke(ctx, AppointmentService_CreateChildOffice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) AppointChildOfficeHead(ctx context.Context, in *OfficeHeadAppointmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AppointmentService_AppointChildOfficeHead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) AppointChildOfficeDeputyHead(ctx context.Context, in *OfficeHeadAppointmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AppointmentService_AppointChildOfficeDeputyHead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppointmentServiceServer is the server API for AppointmentService service.
// All implementations must embed UnimplementedAppointmentServiceServer
// for forward compatibility
type AppointmentServiceServer interface {
	AppointAttender(context.Context, *AttenderAppointmentRequest) (*emptypb.Empty, error)
	CreateChildOffice(context.Context, *CreateChildOfficeRequest) (*CreateChildOfficeResponse, error)
	AppointChildOfficeHead(context.Context, *OfficeHeadAppointmentRequest) (*emptypb.Empty, error)
	AppointChildOfficeDeputyHead(context.Context, *OfficeHeadAppointmentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAppointmentServiceServer()
}

// UnimplementedAppointmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppointmentServiceServer struct {
}

func (UnimplementedAppointmentServiceServer) AppointAttender(context.Context, *AttenderAppointmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppointAttender not implemented")
}
func (UnimplementedAppointmentServiceServer) CreateChildOffice(context.Context, *CreateChildOfficeRequest) (*CreateChildOfficeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChildOffice not implemented")
}
func (UnimplementedAppointmentServiceServer) AppointChildOfficeHead(context.Context, *OfficeHeadAppointmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppointChildOfficeHead not implemented")
}
func (UnimplementedAppointmentServiceServer) AppointChildOfficeDeputyHead(context.Context, *OfficeHeadAppointmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppointChildOfficeDeputyHead not implemented")
}
func (UnimplementedAppointmentServiceServer) mustEmbedUnimplementedAppointmentServiceServer() {}

// UnsafeAppointmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppointmentServiceServer will
// result in compilation errors.
type UnsafeAppointmentServiceServer interface {
	mustEmbedUnimplementedAppointmentServiceServer()
}

func RegisterAppointmentServiceServer(s grpc.ServiceRegistrar, srv AppointmentServiceServer) {
	s.RegisterService(&AppointmentService_ServiceDesc, srv)
}

func _AppointmentService_AppointAttender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttenderAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).AppointAttender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_AppointAttender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).AppointAttender(ctx, req.(*AttenderAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_CreateChildOffice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChildOfficeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).CreateChildOffice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_CreateChildOffice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).CreateChildOffice(ctx, req.(*CreateChildOfficeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_AppointChildOfficeHead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfficeHeadAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).AppointChildOfficeHead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_AppointChildOfficeHead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).AppointChildOfficeHead(ctx, req.(*OfficeHeadAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_AppointChildOfficeDeputyHead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfficeHeadAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).AppointChildOfficeDeputyHead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_AppointChildOfficeDeputyHead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).AppointChildOfficeDeputyHead(ctx, req.(*OfficeHeadAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppointmentService_ServiceDesc is the grpc.ServiceDesc for AppointmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppointmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AppointmentService",
	HandlerType: (*AppointmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppointAttender",
			Handler:    _AppointmentService_AppointAttender_Handler,
		},
		{
			MethodName: "CreateChildOffice",
			Handler:    _AppointmentService_CreateChildOffice_Handler,
		},
		{
			MethodName: "AppointChildOfficeHead",
			Handler:    _AppointmentService_AppointChildOfficeHead_Handler,
		},
		{
			MethodName: "AppointChildOfficeDeputyHead",
			Handler:    _AppointmentService_AppointChildOfficeDeputyHead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/protobuf/appointments.proto",
}
