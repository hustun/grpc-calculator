// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: calculator.proto

package grpc_simple

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Add(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	Subtract(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	Multiply(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	Divide(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/grpc_simple.ChatService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Add(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/grpc_simple.ChatService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Subtract(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/grpc_simple.ChatService/Subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Multiply(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/grpc_simple.ChatService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Divide(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/grpc_simple.ChatService/Divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	Add(context.Context, *CalcRequest) (*CalcResponse, error)
	Subtract(context.Context, *CalcRequest) (*CalcResponse, error)
	Multiply(context.Context, *CalcRequest) (*CalcResponse, error)
	Divide(context.Context, *CalcRequest) (*CalcResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedChatServiceServer) Add(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedChatServiceServer) Subtract(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedChatServiceServer) Multiply(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedChatServiceServer) Divide(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_simple.ChatService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_simple.ChatService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Add(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_simple.ChatService/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Subtract(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_simple.ChatService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Multiply(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_simple.ChatService/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Divide(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_simple.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ChatService_SayHello_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _ChatService_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _ChatService_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _ChatService_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _ChatService_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}