// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: calculator.proto

package proto

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
	CalcualatorService_Add_FullMethodName                    = "/calcualator.CalcualatorService/Add"
	CalcualatorService_AddClientStream_FullMethodName        = "/calcualator.CalcualatorService/AddClientStream"
	CalcualatorService_AddServerStream_FullMethodName        = "/calcualator.CalcualatorService/AddServerStream"
	CalcualatorService_AddBidirectionalStream_FullMethodName = "/calcualator.CalcualatorService/AddBidirectionalStream"
)

// CalcualatorServiceClient is the client API for CalcualatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcualatorServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	AddClientStream(ctx context.Context, opts ...grpc.CallOption) (CalcualatorService_AddClientStreamClient, error)
	AddServerStream(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (CalcualatorService_AddServerStreamClient, error)
	AddBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (CalcualatorService_AddBidirectionalStreamClient, error)
}

type calcualatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcualatorServiceClient(cc grpc.ClientConnInterface) CalcualatorServiceClient {
	return &calcualatorServiceClient{cc}
}

func (c *calcualatorServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, CalcualatorService_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcualatorServiceClient) AddClientStream(ctx context.Context, opts ...grpc.CallOption) (CalcualatorService_AddClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcualatorService_ServiceDesc.Streams[0], CalcualatorService_AddClientStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calcualatorServiceAddClientStreamClient{stream}
	return x, nil
}

type CalcualatorService_AddClientStreamClient interface {
	Send(*AddRequest) error
	CloseAndRecv() (*AddResponse, error)
	grpc.ClientStream
}

type calcualatorServiceAddClientStreamClient struct {
	grpc.ClientStream
}

func (x *calcualatorServiceAddClientStreamClient) Send(m *AddRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcualatorServiceAddClientStreamClient) CloseAndRecv() (*AddResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcualatorServiceClient) AddServerStream(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (CalcualatorService_AddServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcualatorService_ServiceDesc.Streams[1], CalcualatorService_AddServerStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calcualatorServiceAddServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcualatorService_AddServerStreamClient interface {
	Recv() (*AddResponse, error)
	grpc.ClientStream
}

type calcualatorServiceAddServerStreamClient struct {
	grpc.ClientStream
}

func (x *calcualatorServiceAddServerStreamClient) Recv() (*AddResponse, error) {
	m := new(AddResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcualatorServiceClient) AddBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (CalcualatorService_AddBidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcualatorService_ServiceDesc.Streams[2], CalcualatorService_AddBidirectionalStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calcualatorServiceAddBidirectionalStreamClient{stream}
	return x, nil
}

type CalcualatorService_AddBidirectionalStreamClient interface {
	Send(*AddRequest) error
	Recv() (*AddResponse, error)
	grpc.ClientStream
}

type calcualatorServiceAddBidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *calcualatorServiceAddBidirectionalStreamClient) Send(m *AddRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcualatorServiceAddBidirectionalStreamClient) Recv() (*AddResponse, error) {
	m := new(AddResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcualatorServiceServer is the server API for CalcualatorService service.
// All implementations must embed UnimplementedCalcualatorServiceServer
// for forward compatibility
type CalcualatorServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	AddClientStream(CalcualatorService_AddClientStreamServer) error
	AddServerStream(*AddRequest, CalcualatorService_AddServerStreamServer) error
	AddBidirectionalStream(CalcualatorService_AddBidirectionalStreamServer) error
	mustEmbedUnimplementedCalcualatorServiceServer()
}

// UnimplementedCalcualatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalcualatorServiceServer struct {
}

func (UnimplementedCalcualatorServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalcualatorServiceServer) AddClientStream(CalcualatorService_AddClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AddClientStream not implemented")
}
func (UnimplementedCalcualatorServiceServer) AddServerStream(*AddRequest, CalcualatorService_AddServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AddServerStream not implemented")
}
func (UnimplementedCalcualatorServiceServer) AddBidirectionalStream(CalcualatorService_AddBidirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AddBidirectionalStream not implemented")
}
func (UnimplementedCalcualatorServiceServer) mustEmbedUnimplementedCalcualatorServiceServer() {}

// UnsafeCalcualatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcualatorServiceServer will
// result in compilation errors.
type UnsafeCalcualatorServiceServer interface {
	mustEmbedUnimplementedCalcualatorServiceServer()
}

func RegisterCalcualatorServiceServer(s grpc.ServiceRegistrar, srv CalcualatorServiceServer) {
	s.RegisterService(&CalcualatorService_ServiceDesc, srv)
}

func _CalcualatorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcualatorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalcualatorService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcualatorServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcualatorService_AddClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcualatorServiceServer).AddClientStream(&calcualatorServiceAddClientStreamServer{stream})
}

type CalcualatorService_AddClientStreamServer interface {
	SendAndClose(*AddResponse) error
	Recv() (*AddRequest, error)
	grpc.ServerStream
}

type calcualatorServiceAddClientStreamServer struct {
	grpc.ServerStream
}

func (x *calcualatorServiceAddClientStreamServer) SendAndClose(m *AddResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcualatorServiceAddClientStreamServer) Recv() (*AddRequest, error) {
	m := new(AddRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalcualatorService_AddServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AddRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcualatorServiceServer).AddServerStream(m, &calcualatorServiceAddServerStreamServer{stream})
}

type CalcualatorService_AddServerStreamServer interface {
	Send(*AddResponse) error
	grpc.ServerStream
}

type calcualatorServiceAddServerStreamServer struct {
	grpc.ServerStream
}

func (x *calcualatorServiceAddServerStreamServer) Send(m *AddResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalcualatorService_AddBidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcualatorServiceServer).AddBidirectionalStream(&calcualatorServiceAddBidirectionalStreamServer{stream})
}

type CalcualatorService_AddBidirectionalStreamServer interface {
	Send(*AddResponse) error
	Recv() (*AddRequest, error)
	grpc.ServerStream
}

type calcualatorServiceAddBidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *calcualatorServiceAddBidirectionalStreamServer) Send(m *AddResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcualatorServiceAddBidirectionalStreamServer) Recv() (*AddRequest, error) {
	m := new(AddRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcualatorService_ServiceDesc is the grpc.ServiceDesc for CalcualatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcualatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calcualator.CalcualatorService",
	HandlerType: (*CalcualatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalcualatorService_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddClientStream",
			Handler:       _CalcualatorService_AddClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AddServerStream",
			Handler:       _CalcualatorService_AddServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddBidirectionalStream",
			Handler:       _CalcualatorService_AddBidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
