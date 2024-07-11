// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: pb/protobuf/ksebUser.proto

package ksebpb

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
	KSEBUserService_AddConsumerNumber_FullMethodName = "/KSEBUserService/AddConsumerNumber"
)

// KSEBUserServiceClient is the client API for KSEBUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KSEBUserServiceClient interface {
	AddConsumerNumber(ctx context.Context, in *AddConsumerNumberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type kSEBUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKSEBUserServiceClient(cc grpc.ClientConnInterface) KSEBUserServiceClient {
	return &kSEBUserServiceClient{cc}
}

func (c *kSEBUserServiceClient) AddConsumerNumber(ctx context.Context, in *AddConsumerNumberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KSEBUserService_AddConsumerNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KSEBUserServiceServer is the server API for KSEBUserService service.
// All implementations must embed UnimplementedKSEBUserServiceServer
// for forward compatibility
type KSEBUserServiceServer interface {
	AddConsumerNumber(context.Context, *AddConsumerNumberRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedKSEBUserServiceServer()
}

// UnimplementedKSEBUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKSEBUserServiceServer struct {
}

func (UnimplementedKSEBUserServiceServer) AddConsumerNumber(context.Context, *AddConsumerNumberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConsumerNumber not implemented")
}
func (UnimplementedKSEBUserServiceServer) mustEmbedUnimplementedKSEBUserServiceServer() {}

// UnsafeKSEBUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KSEBUserServiceServer will
// result in compilation errors.
type UnsafeKSEBUserServiceServer interface {
	mustEmbedUnimplementedKSEBUserServiceServer()
}

func RegisterKSEBUserServiceServer(s grpc.ServiceRegistrar, srv KSEBUserServiceServer) {
	s.RegisterService(&KSEBUserService_ServiceDesc, srv)
}

func _KSEBUserService_AddConsumerNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddConsumerNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KSEBUserServiceServer).AddConsumerNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KSEBUserService_AddConsumerNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KSEBUserServiceServer).AddConsumerNumber(ctx, req.(*AddConsumerNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KSEBUserService_ServiceDesc is the grpc.ServiceDesc for KSEBUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KSEBUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KSEBUserService",
	HandlerType: (*KSEBUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddConsumerNumber",
			Handler:    _KSEBUserService_AddConsumerNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/protobuf/ksebUser.proto",
}
