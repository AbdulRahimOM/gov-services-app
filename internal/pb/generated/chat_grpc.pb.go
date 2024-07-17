// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: pb/protobuf/chat.proto

package pb

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
	KsebChatService_SendMessage_FullMethodName = "/KsebChatService/SendMessage"
	KsebChatService_UserChat_FullMethodName    = "/KsebChatService/UserChat"
)

// KsebChatServiceClient is the client API for KsebChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KsebChatServiceClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	UserChat(ctx context.Context, in *UserChatRequest, opts ...grpc.CallOption) (KsebChatService_UserChatClient, error)
}

type ksebChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKsebChatServiceClient(cc grpc.ClientConnInterface) KsebChatServiceClient {
	return &ksebChatServiceClient{cc}
}

func (c *ksebChatServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, KsebChatService_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ksebChatServiceClient) UserChat(ctx context.Context, in *UserChatRequest, opts ...grpc.CallOption) (KsebChatService_UserChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &KsebChatService_ServiceDesc.Streams[0], KsebChatService_UserChat_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &ksebChatServiceUserChatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KsebChatService_UserChatClient interface {
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type ksebChatServiceUserChatClient struct {
	grpc.ClientStream
}

func (x *ksebChatServiceUserChatClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KsebChatServiceServer is the server API for KsebChatService service.
// All implementations must embed UnimplementedKsebChatServiceServer
// for forward compatibility
type KsebChatServiceServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	UserChat(*UserChatRequest, KsebChatService_UserChatServer) error
	mustEmbedUnimplementedKsebChatServiceServer()
}

// UnimplementedKsebChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKsebChatServiceServer struct {
}

func (UnimplementedKsebChatServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedKsebChatServiceServer) UserChat(*UserChatRequest, KsebChatService_UserChatServer) error {
	return status.Errorf(codes.Unimplemented, "method UserChat not implemented")
}
func (UnimplementedKsebChatServiceServer) mustEmbedUnimplementedKsebChatServiceServer() {}

// UnsafeKsebChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KsebChatServiceServer will
// result in compilation errors.
type UnsafeKsebChatServiceServer interface {
	mustEmbedUnimplementedKsebChatServiceServer()
}

func RegisterKsebChatServiceServer(s grpc.ServiceRegistrar, srv KsebChatServiceServer) {
	s.RegisterService(&KsebChatService_ServiceDesc, srv)
}

func _KsebChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KsebChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KsebChatService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KsebChatServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KsebChatService_UserChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserChatRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KsebChatServiceServer).UserChat(m, &ksebChatServiceUserChatServer{stream})
}

type KsebChatService_UserChatServer interface {
	Send(*ChatMessage) error
	grpc.ServerStream
}

type ksebChatServiceUserChatServer struct {
	grpc.ServerStream
}

func (x *ksebChatServiceUserChatServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

// KsebChatService_ServiceDesc is the grpc.ServiceDesc for KsebChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KsebChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KsebChatService",
	HandlerType: (*KsebChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _KsebChatService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UserChat",
			Handler:       _KsebChatService_UserChat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/protobuf/chat.proto",
}
