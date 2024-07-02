// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: pb/protobuf/adminAccount.proto

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
	AdminAccountService_AdminLoginViaPassword_FullMethodName         = "/AdminAccountService/AdminLoginViaPassword"
	AdminAccountService_AdminGetProfile_FullMethodName               = "/AdminAccountService/AdminGetProfile"
	AdminAccountService_AdminUpdateProfile_FullMethodName            = "/AdminAccountService/AdminUpdateProfile"
	AdminAccountService_AdminUpdatePasswordUsingOldPw_FullMethodName = "/AdminAccountService/AdminUpdatePasswordUsingOldPw"
	AdminAccountService_AdminGetAdmins_FullMethodName                = "/AdminAccountService/AdminGetAdmins"
)

// AdminAccountServiceClient is the client API for AdminAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminAccountServiceClient interface {
	// Login
	AdminLoginViaPassword(ctx context.Context, in *AdminLoginViaPasswordRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	// profile
	AdminGetProfile(ctx context.Context, in *AdminGetProfileRequest, opts ...grpc.CallOption) (*AdminGetProfileResponse, error)
	AdminUpdateProfile(ctx context.Context, in *AdminUpdateProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AdminUpdatePasswordUsingOldPw(ctx context.Context, in *AdminUpdatePasswordUsingOldPwRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// manage account
	AdminGetAdmins(ctx context.Context, in *AdminGetAdminsRequest, opts ...grpc.CallOption) (*AdminGetAdminsResponse, error)
}

type adminAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminAccountServiceClient(cc grpc.ClientConnInterface) AdminAccountServiceClient {
	return &adminAccountServiceClient{cc}
}

func (c *adminAccountServiceClient) AdminLoginViaPassword(ctx context.Context, in *AdminLoginViaPasswordRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, AdminAccountService_AdminLoginViaPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAccountServiceClient) AdminGetProfile(ctx context.Context, in *AdminGetProfileRequest, opts ...grpc.CallOption) (*AdminGetProfileResponse, error) {
	out := new(AdminGetProfileResponse)
	err := c.cc.Invoke(ctx, AdminAccountService_AdminGetProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAccountServiceClient) AdminUpdateProfile(ctx context.Context, in *AdminUpdateProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AdminAccountService_AdminUpdateProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAccountServiceClient) AdminUpdatePasswordUsingOldPw(ctx context.Context, in *AdminUpdatePasswordUsingOldPwRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AdminAccountService_AdminUpdatePasswordUsingOldPw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAccountServiceClient) AdminGetAdmins(ctx context.Context, in *AdminGetAdminsRequest, opts ...grpc.CallOption) (*AdminGetAdminsResponse, error) {
	out := new(AdminGetAdminsResponse)
	err := c.cc.Invoke(ctx, AdminAccountService_AdminGetAdmins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminAccountServiceServer is the server API for AdminAccountService service.
// All implementations must embed UnimplementedAdminAccountServiceServer
// for forward compatibility
type AdminAccountServiceServer interface {
	// Login
	AdminLoginViaPassword(context.Context, *AdminLoginViaPasswordRequest) (*AdminLoginResponse, error)
	// profile
	AdminGetProfile(context.Context, *AdminGetProfileRequest) (*AdminGetProfileResponse, error)
	AdminUpdateProfile(context.Context, *AdminUpdateProfileRequest) (*emptypb.Empty, error)
	AdminUpdatePasswordUsingOldPw(context.Context, *AdminUpdatePasswordUsingOldPwRequest) (*emptypb.Empty, error)
	// manage account
	AdminGetAdmins(context.Context, *AdminGetAdminsRequest) (*AdminGetAdminsResponse, error)
	mustEmbedUnimplementedAdminAccountServiceServer()
}

// UnimplementedAdminAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminAccountServiceServer struct {
}

func (UnimplementedAdminAccountServiceServer) AdminLoginViaPassword(context.Context, *AdminLoginViaPasswordRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLoginViaPassword not implemented")
}
func (UnimplementedAdminAccountServiceServer) AdminGetProfile(context.Context, *AdminGetProfileRequest) (*AdminGetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetProfile not implemented")
}
func (UnimplementedAdminAccountServiceServer) AdminUpdateProfile(context.Context, *AdminUpdateProfileRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateProfile not implemented")
}
func (UnimplementedAdminAccountServiceServer) AdminUpdatePasswordUsingOldPw(context.Context, *AdminUpdatePasswordUsingOldPwRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdatePasswordUsingOldPw not implemented")
}
func (UnimplementedAdminAccountServiceServer) AdminGetAdmins(context.Context, *AdminGetAdminsRequest) (*AdminGetAdminsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetAdmins not implemented")
}
func (UnimplementedAdminAccountServiceServer) mustEmbedUnimplementedAdminAccountServiceServer() {}

// UnsafeAdminAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminAccountServiceServer will
// result in compilation errors.
type UnsafeAdminAccountServiceServer interface {
	mustEmbedUnimplementedAdminAccountServiceServer()
}

func RegisterAdminAccountServiceServer(s grpc.ServiceRegistrar, srv AdminAccountServiceServer) {
	s.RegisterService(&AdminAccountService_ServiceDesc, srv)
}

func _AdminAccountService_AdminLoginViaPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginViaPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAccountServiceServer).AdminLoginViaPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAccountService_AdminLoginViaPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAccountServiceServer).AdminLoginViaPassword(ctx, req.(*AdminLoginViaPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAccountService_AdminGetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAccountServiceServer).AdminGetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAccountService_AdminGetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAccountServiceServer).AdminGetProfile(ctx, req.(*AdminGetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAccountService_AdminUpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAccountServiceServer).AdminUpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAccountService_AdminUpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAccountServiceServer).AdminUpdateProfile(ctx, req.(*AdminUpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAccountService_AdminUpdatePasswordUsingOldPw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdatePasswordUsingOldPwRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAccountServiceServer).AdminUpdatePasswordUsingOldPw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAccountService_AdminUpdatePasswordUsingOldPw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAccountServiceServer).AdminUpdatePasswordUsingOldPw(ctx, req.(*AdminUpdatePasswordUsingOldPwRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAccountService_AdminGetAdmins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetAdminsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAccountServiceServer).AdminGetAdmins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAccountService_AdminGetAdmins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAccountServiceServer).AdminGetAdmins(ctx, req.(*AdminGetAdminsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminAccountService_ServiceDesc is the grpc.ServiceDesc for AdminAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AdminAccountService",
	HandlerType: (*AdminAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLoginViaPassword",
			Handler:    _AdminAccountService_AdminLoginViaPassword_Handler,
		},
		{
			MethodName: "AdminGetProfile",
			Handler:    _AdminAccountService_AdminGetProfile_Handler,
		},
		{
			MethodName: "AdminUpdateProfile",
			Handler:    _AdminAccountService_AdminUpdateProfile_Handler,
		},
		{
			MethodName: "AdminUpdatePasswordUsingOldPw",
			Handler:    _AdminAccountService_AdminUpdatePasswordUsingOldPw_Handler,
		},
		{
			MethodName: "AdminGetAdmins",
			Handler:    _AdminAccountService_AdminGetAdmins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/protobuf/adminAccount.proto",
}
