// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: pb/protobuf/ksebUser.proto

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

const ()

// KSEBUserAccServiceClient is the client API for KSEBUserAccService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KSEBUserAccServiceClient interface {
}

type kSEBUserAccServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKSEBUserAccServiceClient(cc grpc.ClientConnInterface) KSEBUserAccServiceClient {
	return &kSEBUserAccServiceClient{cc}
}

// KSEBUserAccServiceServer is the server API for KSEBUserAccService service.
// All implementations must embed UnimplementedKSEBUserAccServiceServer
// for forward compatibility
type KSEBUserAccServiceServer interface {
	mustEmbedUnimplementedKSEBUserAccServiceServer()
}

// UnimplementedKSEBUserAccServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKSEBUserAccServiceServer struct {
}

func (UnimplementedKSEBUserAccServiceServer) mustEmbedUnimplementedKSEBUserAccServiceServer() {}

// UnsafeKSEBUserAccServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KSEBUserAccServiceServer will
// result in compilation errors.
type UnsafeKSEBUserAccServiceServer interface {
	mustEmbedUnimplementedKSEBUserAccServiceServer()
}

func RegisterKSEBUserAccServiceServer(s grpc.ServiceRegistrar, srv KSEBUserAccServiceServer) {
	s.RegisterService(&KSEBUserAccService_ServiceDesc, srv)
}

// KSEBUserAccService_ServiceDesc is the grpc.ServiceDesc for KSEBUserAccService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KSEBUserAccService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KSEBUserAccService",
	HandlerType: (*KSEBUserAccServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "pb/protobuf/ksebUser.proto",
}

const (
	KSEBAgencyUserService_AddConsumerNumber_FullMethodName             = "/KSEBAgencyUserService/AddConsumerNumber"
	KSEBAgencyUserService_GetUserConsumerNumbers_FullMethodName        = "/KSEBAgencyUserService/GetUserConsumerNumbers"
	KSEBAgencyUserService_RaiseComplaint_FullMethodName                = "/KSEBAgencyUserService/RaiseComplaint"
	KSEBAgencyUserService_CheckIfComplaintBelongsToUser_FullMethodName = "/KSEBAgencyUserService/CheckIfComplaintBelongsToUser"
)

// KSEBAgencyUserServiceClient is the client API for KSEBAgencyUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KSEBAgencyUserServiceClient interface {
	// user
	AddConsumerNumber(ctx context.Context, in *AddConsumerNumberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserConsumerNumbers(ctx context.Context, in *GetUserConsumerNumbersRequest, opts ...grpc.CallOption) (*GetUserConsumerNumbersResponse, error)
	RaiseComplaint(ctx context.Context, in *RaiseComplaintRequest, opts ...grpc.CallOption) (*RaiseComplaintResponse, error)
	CheckIfComplaintBelongsToUser(ctx context.Context, in *CheckIfComplaintBelongsToUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type kSEBAgencyUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKSEBAgencyUserServiceClient(cc grpc.ClientConnInterface) KSEBAgencyUserServiceClient {
	return &kSEBAgencyUserServiceClient{cc}
}

func (c *kSEBAgencyUserServiceClient) AddConsumerNumber(ctx context.Context, in *AddConsumerNumberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KSEBAgencyUserService_AddConsumerNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kSEBAgencyUserServiceClient) GetUserConsumerNumbers(ctx context.Context, in *GetUserConsumerNumbersRequest, opts ...grpc.CallOption) (*GetUserConsumerNumbersResponse, error) {
	out := new(GetUserConsumerNumbersResponse)
	err := c.cc.Invoke(ctx, KSEBAgencyUserService_GetUserConsumerNumbers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kSEBAgencyUserServiceClient) RaiseComplaint(ctx context.Context, in *RaiseComplaintRequest, opts ...grpc.CallOption) (*RaiseComplaintResponse, error) {
	out := new(RaiseComplaintResponse)
	err := c.cc.Invoke(ctx, KSEBAgencyUserService_RaiseComplaint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kSEBAgencyUserServiceClient) CheckIfComplaintBelongsToUser(ctx context.Context, in *CheckIfComplaintBelongsToUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KSEBAgencyUserService_CheckIfComplaintBelongsToUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KSEBAgencyUserServiceServer is the server API for KSEBAgencyUserService service.
// All implementations must embed UnimplementedKSEBAgencyUserServiceServer
// for forward compatibility
type KSEBAgencyUserServiceServer interface {
	// user
	AddConsumerNumber(context.Context, *AddConsumerNumberRequest) (*emptypb.Empty, error)
	GetUserConsumerNumbers(context.Context, *GetUserConsumerNumbersRequest) (*GetUserConsumerNumbersResponse, error)
	RaiseComplaint(context.Context, *RaiseComplaintRequest) (*RaiseComplaintResponse, error)
	CheckIfComplaintBelongsToUser(context.Context, *CheckIfComplaintBelongsToUserRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedKSEBAgencyUserServiceServer()
}

// UnimplementedKSEBAgencyUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKSEBAgencyUserServiceServer struct {
}

func (UnimplementedKSEBAgencyUserServiceServer) AddConsumerNumber(context.Context, *AddConsumerNumberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConsumerNumber not implemented")
}
func (UnimplementedKSEBAgencyUserServiceServer) GetUserConsumerNumbers(context.Context, *GetUserConsumerNumbersRequest) (*GetUserConsumerNumbersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserConsumerNumbers not implemented")
}
func (UnimplementedKSEBAgencyUserServiceServer) RaiseComplaint(context.Context, *RaiseComplaintRequest) (*RaiseComplaintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RaiseComplaint not implemented")
}
func (UnimplementedKSEBAgencyUserServiceServer) CheckIfComplaintBelongsToUser(context.Context, *CheckIfComplaintBelongsToUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfComplaintBelongsToUser not implemented")
}
func (UnimplementedKSEBAgencyUserServiceServer) mustEmbedUnimplementedKSEBAgencyUserServiceServer() {}

// UnsafeKSEBAgencyUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KSEBAgencyUserServiceServer will
// result in compilation errors.
type UnsafeKSEBAgencyUserServiceServer interface {
	mustEmbedUnimplementedKSEBAgencyUserServiceServer()
}

func RegisterKSEBAgencyUserServiceServer(s grpc.ServiceRegistrar, srv KSEBAgencyUserServiceServer) {
	s.RegisterService(&KSEBAgencyUserService_ServiceDesc, srv)
}

func _KSEBAgencyUserService_AddConsumerNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddConsumerNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KSEBAgencyUserServiceServer).AddConsumerNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KSEBAgencyUserService_AddConsumerNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KSEBAgencyUserServiceServer).AddConsumerNumber(ctx, req.(*AddConsumerNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KSEBAgencyUserService_GetUserConsumerNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserConsumerNumbersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KSEBAgencyUserServiceServer).GetUserConsumerNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KSEBAgencyUserService_GetUserConsumerNumbers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KSEBAgencyUserServiceServer).GetUserConsumerNumbers(ctx, req.(*GetUserConsumerNumbersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KSEBAgencyUserService_RaiseComplaint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaiseComplaintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KSEBAgencyUserServiceServer).RaiseComplaint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KSEBAgencyUserService_RaiseComplaint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KSEBAgencyUserServiceServer).RaiseComplaint(ctx, req.(*RaiseComplaintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KSEBAgencyUserService_CheckIfComplaintBelongsToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIfComplaintBelongsToUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KSEBAgencyUserServiceServer).CheckIfComplaintBelongsToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KSEBAgencyUserService_CheckIfComplaintBelongsToUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KSEBAgencyUserServiceServer).CheckIfComplaintBelongsToUser(ctx, req.(*CheckIfComplaintBelongsToUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KSEBAgencyUserService_ServiceDesc is the grpc.ServiceDesc for KSEBAgencyUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KSEBAgencyUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KSEBAgencyUserService",
	HandlerType: (*KSEBAgencyUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddConsumerNumber",
			Handler:    _KSEBAgencyUserService_AddConsumerNumber_Handler,
		},
		{
			MethodName: "GetUserConsumerNumbers",
			Handler:    _KSEBAgencyUserService_GetUserConsumerNumbers_Handler,
		},
		{
			MethodName: "RaiseComplaint",
			Handler:    _KSEBAgencyUserService_RaiseComplaint_Handler,
		},
		{
			MethodName: "CheckIfComplaintBelongsToUser",
			Handler:    _KSEBAgencyUserService_CheckIfComplaintBelongsToUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/protobuf/ksebUser.proto",
}