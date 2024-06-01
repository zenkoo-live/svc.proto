// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: svc.biz.log/log.proto

package log

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
	Log_AddLog_FullMethodName      = "/svc.biz.log.Log/AddLog"
	Log_MGetLastLog_FullMethodName = "/svc.biz.log.Log/MGetLastLog"
)

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	// AddLog 记录日志
	AddLog(ctx context.Context, in *AddLogReq, opts ...grpc.CallOption) (*AddLogResp, error)
	// MGetLastLog 批量获取最近一次操作
	MGetLastLog(ctx context.Context, in *MGetLastLogReq, opts ...grpc.CallOption) (*MGetLastLogResp, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) AddLog(ctx context.Context, in *AddLogReq, opts ...grpc.CallOption) (*AddLogResp, error) {
	out := new(AddLogResp)
	err := c.cc.Invoke(ctx, Log_AddLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) MGetLastLog(ctx context.Context, in *MGetLastLogReq, opts ...grpc.CallOption) (*MGetLastLogResp, error) {
	out := new(MGetLastLogResp)
	err := c.cc.Invoke(ctx, Log_MGetLastLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
// All implementations must embed UnimplementedLogServer
// for forward compatibility
type LogServer interface {
	// AddLog 记录日志
	AddLog(context.Context, *AddLogReq) (*AddLogResp, error)
	// MGetLastLog 批量获取最近一次操作
	MGetLastLog(context.Context, *MGetLastLogReq) (*MGetLastLogResp, error)
	mustEmbedUnimplementedLogServer()
}

// UnimplementedLogServer must be embedded to have forward compatible implementations.
type UnimplementedLogServer struct {
}

func (UnimplementedLogServer) AddLog(context.Context, *AddLogReq) (*AddLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLog not implemented")
}
func (UnimplementedLogServer) MGetLastLog(context.Context, *MGetLastLogReq) (*MGetLastLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetLastLog not implemented")
}
func (UnimplementedLogServer) mustEmbedUnimplementedLogServer() {}

// UnsafeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServer will
// result in compilation errors.
type UnsafeLogServer interface {
	mustEmbedUnimplementedLogServer()
}

func RegisterLogServer(s grpc.ServiceRegistrar, srv LogServer) {
	s.RegisterService(&Log_ServiceDesc, srv)
}

func _Log_AddLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).AddLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_AddLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).AddLog(ctx, req.(*AddLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_MGetLastLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetLastLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).MGetLastLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_MGetLastLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).MGetLastLog(ctx, req.(*MGetLastLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Log_ServiceDesc is the grpc.ServiceDesc for Log service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Log_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.log.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddLog",
			Handler:    _Log_AddLog_Handler,
		},
		{
			MethodName: "MGetLastLog",
			Handler:    _Log_MGetLastLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.log/log.proto",
}
