// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: svc.infra.link/trace.proto

package link

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
	LinkTrace_Start_FullMethodName = "/svc.infra.link.LinkTrace/Start"
	LinkTrace_Stop_FullMethodName  = "/svc.infra.link.LinkTrace/Stop"
)

// LinkTraceClient is the client API for LinkTrace service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkTraceClient interface {
	// 设置Trace
	Start(ctx context.Context, in *StartTraceRequest, opts ...grpc.CallOption) (*StartTraceResponse, error)
	// 关闭Trace
	Stop(ctx context.Context, in *StopTraceRequest, opts ...grpc.CallOption) (*StopTraceResponse, error)
}

type linkTraceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkTraceClient(cc grpc.ClientConnInterface) LinkTraceClient {
	return &linkTraceClient{cc}
}

func (c *linkTraceClient) Start(ctx context.Context, in *StartTraceRequest, opts ...grpc.CallOption) (*StartTraceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartTraceResponse)
	err := c.cc.Invoke(ctx, LinkTrace_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkTraceClient) Stop(ctx context.Context, in *StopTraceRequest, opts ...grpc.CallOption) (*StopTraceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopTraceResponse)
	err := c.cc.Invoke(ctx, LinkTrace_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkTraceServer is the server API for LinkTrace service.
// All implementations must embed UnimplementedLinkTraceServer
// for forward compatibility.
type LinkTraceServer interface {
	// 设置Trace
	Start(context.Context, *StartTraceRequest) (*StartTraceResponse, error)
	// 关闭Trace
	Stop(context.Context, *StopTraceRequest) (*StopTraceResponse, error)
	mustEmbedUnimplementedLinkTraceServer()
}

// UnimplementedLinkTraceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLinkTraceServer struct{}

func (UnimplementedLinkTraceServer) Start(context.Context, *StartTraceRequest) (*StartTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedLinkTraceServer) Stop(context.Context, *StopTraceRequest) (*StopTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedLinkTraceServer) mustEmbedUnimplementedLinkTraceServer() {}
func (UnimplementedLinkTraceServer) testEmbeddedByValue()                   {}

// UnsafeLinkTraceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkTraceServer will
// result in compilation errors.
type UnsafeLinkTraceServer interface {
	mustEmbedUnimplementedLinkTraceServer()
}

func RegisterLinkTraceServer(s grpc.ServiceRegistrar, srv LinkTraceServer) {
	// If the following call pancis, it indicates UnimplementedLinkTraceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LinkTrace_ServiceDesc, srv)
}

func _LinkTrace_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkTraceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkTrace_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkTraceServer).Start(ctx, req.(*StartTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkTrace_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkTraceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkTrace_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkTraceServer).Stop(ctx, req.(*StopTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkTrace_ServiceDesc is the grpc.ServiceDesc for LinkTrace service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkTrace_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.infra.link.LinkTrace",
	HandlerType: (*LinkTraceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _LinkTrace_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _LinkTrace_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.infra.link/trace.proto",
}
