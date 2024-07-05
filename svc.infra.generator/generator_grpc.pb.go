// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: svc.infra.generator/generator.proto

package generator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Generator_InitDB_FullMethodName          = "/svc.infra.generator.Generator/InitDB"
	Generator_InitIDGenerator_FullMethodName = "/svc.infra.generator.Generator/InitIDGenerator"
	Generator_OrdianID_FullMethodName        = "/svc.infra.generator.Generator/OrdianID"
	Generator_NextID_FullMethodName          = "/svc.infra.generator.Generator/NextID"
	Generator_IsPrettyID_FullMethodName      = "/svc.infra.generator.Generator/IsPrettyID"
)

// GeneratorClient is the client API for Generator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeneratorClient interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error)
	InitIDGenerator(ctx context.Context, in *InitIDGeneratorReq, opts ...grpc.CallOption) (*InitIDGeneratorResp, error)
	OrdianID(ctx context.Context, in *OrdianIDReq, opts ...grpc.CallOption) (*OrdianIDResp, error)
	NextID(ctx context.Context, in *NextIDReq, opts ...grpc.CallOption) (*NextIDResp, error)
	IsPrettyID(ctx context.Context, in *IsPrettyIDReq, opts ...grpc.CallOption) (*IsPrettyIDResp, error)
}

type generatorClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneratorClient(cc grpc.ClientConnInterface) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitDBResp)
	err := c.cc.Invoke(ctx, Generator_InitDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorClient) InitIDGenerator(ctx context.Context, in *InitIDGeneratorReq, opts ...grpc.CallOption) (*InitIDGeneratorResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitIDGeneratorResp)
	err := c.cc.Invoke(ctx, Generator_InitIDGenerator_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorClient) OrdianID(ctx context.Context, in *OrdianIDReq, opts ...grpc.CallOption) (*OrdianIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrdianIDResp)
	err := c.cc.Invoke(ctx, Generator_OrdianID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorClient) NextID(ctx context.Context, in *NextIDReq, opts ...grpc.CallOption) (*NextIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NextIDResp)
	err := c.cc.Invoke(ctx, Generator_NextID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorClient) IsPrettyID(ctx context.Context, in *IsPrettyIDReq, opts ...grpc.CallOption) (*IsPrettyIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsPrettyIDResp)
	err := c.cc.Invoke(ctx, Generator_IsPrettyID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneratorServer is the server API for Generator service.
// All implementations must embed UnimplementedGeneratorServer
// for forward compatibility
type GeneratorServer interface {
	InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error)
	InitIDGenerator(context.Context, *InitIDGeneratorReq) (*InitIDGeneratorResp, error)
	OrdianID(context.Context, *OrdianIDReq) (*OrdianIDResp, error)
	NextID(context.Context, *NextIDReq) (*NextIDResp, error)
	IsPrettyID(context.Context, *IsPrettyIDReq) (*IsPrettyIDResp, error)
	mustEmbedUnimplementedGeneratorServer()
}

// UnimplementedGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedGeneratorServer struct {
}

func (UnimplementedGeneratorServer) InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDB not implemented")
}
func (UnimplementedGeneratorServer) InitIDGenerator(context.Context, *InitIDGeneratorReq) (*InitIDGeneratorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitIDGenerator not implemented")
}
func (UnimplementedGeneratorServer) OrdianID(context.Context, *OrdianIDReq) (*OrdianIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrdianID not implemented")
}
func (UnimplementedGeneratorServer) NextID(context.Context, *NextIDReq) (*NextIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextID not implemented")
}
func (UnimplementedGeneratorServer) IsPrettyID(context.Context, *IsPrettyIDReq) (*IsPrettyIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPrettyID not implemented")
}
func (UnimplementedGeneratorServer) mustEmbedUnimplementedGeneratorServer() {}

// UnsafeGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeneratorServer will
// result in compilation errors.
type UnsafeGeneratorServer interface {
	mustEmbedUnimplementedGeneratorServer()
}

func RegisterGeneratorServer(s grpc.ServiceRegistrar, srv GeneratorServer) {
	s.RegisterService(&Generator_ServiceDesc, srv)
}

func _Generator_InitDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).InitDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Generator_InitDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).InitDB(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Generator_InitIDGenerator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitIDGeneratorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).InitIDGenerator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Generator_InitIDGenerator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).InitIDGenerator(ctx, req.(*InitIDGeneratorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Generator_OrdianID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdianIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).OrdianID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Generator_OrdianID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).OrdianID(ctx, req.(*OrdianIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Generator_NextID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).NextID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Generator_NextID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).NextID(ctx, req.(*NextIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Generator_IsPrettyID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsPrettyIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).IsPrettyID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Generator_IsPrettyID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).IsPrettyID(ctx, req.(*IsPrettyIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Generator_ServiceDesc is the grpc.ServiceDesc for Generator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Generator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.infra.generator.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDB",
			Handler:    _Generator_InitDB_Handler,
		},
		{
			MethodName: "InitIDGenerator",
			Handler:    _Generator_InitIDGenerator_Handler,
		},
		{
			MethodName: "OrdianID",
			Handler:    _Generator_OrdianID_Handler,
		},
		{
			MethodName: "NextID",
			Handler:    _Generator_NextID_Handler,
		},
		{
			MethodName: "IsPrettyID",
			Handler:    _Generator_IsPrettyID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.infra.generator/generator.proto",
}
