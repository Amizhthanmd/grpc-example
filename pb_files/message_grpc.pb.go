// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: protos/message.proto

package pb_files

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExampleService_SayHello_FullMethodName     = "/example.ExampleService/SayHello"
	ExampleService_SendMessages_FullMethodName = "/example.ExampleService/SendMessages"
	ExampleService_GetMessages_FullMethodName  = "/example.ExampleService/GetMessages"
	ExampleService_Chat_FullMethodName         = "/example.ExampleService/Chat"
)

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	// Unary
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// Client Streaming
	SendMessages(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[HelloRequest, HelloReply], error)
	// Server Streaming
	GetMessages(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloReply], error)
	// Bidirectional Streaming
	Chat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ChatMessage, ChatMessage], error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ExampleService_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) SendMessages(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[HelloRequest, HelloReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[0], ExampleService_SendMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloReply]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_SendMessagesClient = grpc.ClientStreamingClient[HelloRequest, HelloReply]

func (c *exampleServiceClient) GetMessages(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[1], ExampleService_GetMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_GetMessagesClient = grpc.ServerStreamingClient[HelloReply]

func (c *exampleServiceClient) Chat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ChatMessage, ChatMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[2], ExampleService_Chat_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ChatMessage, ChatMessage]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_ChatClient = grpc.BidiStreamingClient[ChatMessage, ChatMessage]

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility.
type ExampleServiceServer interface {
	// Unary
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// Client Streaming
	SendMessages(grpc.ClientStreamingServer[HelloRequest, HelloReply]) error
	// Server Streaming
	GetMessages(*HelloRequest, grpc.ServerStreamingServer[HelloReply]) error
	// Bidirectional Streaming
	Chat(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExampleServiceServer struct{}

func (UnimplementedExampleServiceServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedExampleServiceServer) SendMessages(grpc.ClientStreamingServer[HelloRequest, HelloReply]) error {
	return status.Errorf(codes.Unimplemented, "method SendMessages not implemented")
}
func (UnimplementedExampleServiceServer) GetMessages(*HelloRequest, grpc.ServerStreamingServer[HelloReply]) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedExampleServiceServer) Chat(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}
func (UnimplementedExampleServiceServer) testEmbeddedByValue()                        {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	// If the following call pancis, it indicates UnimplementedExampleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_SendMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).SendMessages(&grpc.GenericServerStream[HelloRequest, HelloReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_SendMessagesServer = grpc.ClientStreamingServer[HelloRequest, HelloReply]

func _ExampleService_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServiceServer).GetMessages(m, &grpc.GenericServerStream[HelloRequest, HelloReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_GetMessagesServer = grpc.ServerStreamingServer[HelloReply]

func _ExampleService_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).Chat(&grpc.GenericServerStream[ChatMessage, ChatMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_ChatServer = grpc.BidiStreamingServer[ChatMessage, ChatMessage]

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ExampleService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessages",
			Handler:       _ExampleService_SendMessages_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetMessages",
			Handler:       _ExampleService_GetMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Chat",
			Handler:       _ExampleService_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/message.proto",
}
