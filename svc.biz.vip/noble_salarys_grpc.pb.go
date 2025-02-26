// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: svc.biz.vip/noble_salarys.proto

package vip

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
	NobleSalary_Distribute_FullMethodName     = "/svc.biz.vip.NobleSalary/Distribute"
	NobleSalary_Receive_FullMethodName        = "/svc.biz.vip.NobleSalary/Receive"
	NobleSalary_List_FullMethodName           = "/svc.biz.vip.NobleSalary/List"
	NobleSalary_GetReceiveInfo_FullMethodName = "/svc.biz.vip.NobleSalary/GetReceiveInfo"
)

// NobleSalaryClient is the client API for NobleSalary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NobleSalaryClient interface {
	// 发放俸禄(任务系统调用/每天)
	Distribute(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DistributeSalaryResp, error)
	// 领取俸禄
	Receive(ctx context.Context, in *ReceiveSalaryReq, opts ...grpc.CallOption) (*ReceiveSalaryResp, error)
	// 查询俸禄列表(按发放时间倒序)
	List(ctx context.Context, in *NobleSalaryListReq, opts ...grpc.CallOption) (*NobleSalaryListResp, error)
	// 查询俸禄领信息(金额、状态等)
	GetReceiveInfo(ctx context.Context, in *ReceiveSalaryReq, opts ...grpc.CallOption) (*NobleSalaryInfoResp, error)
}

type nobleSalaryClient struct {
	cc grpc.ClientConnInterface
}

func NewNobleSalaryClient(cc grpc.ClientConnInterface) NobleSalaryClient {
	return &nobleSalaryClient{cc}
}

func (c *nobleSalaryClient) Distribute(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DistributeSalaryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DistributeSalaryResp)
	err := c.cc.Invoke(ctx, NobleSalary_Distribute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleSalaryClient) Receive(ctx context.Context, in *ReceiveSalaryReq, opts ...grpc.CallOption) (*ReceiveSalaryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReceiveSalaryResp)
	err := c.cc.Invoke(ctx, NobleSalary_Receive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleSalaryClient) List(ctx context.Context, in *NobleSalaryListReq, opts ...grpc.CallOption) (*NobleSalaryListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NobleSalaryListResp)
	err := c.cc.Invoke(ctx, NobleSalary_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleSalaryClient) GetReceiveInfo(ctx context.Context, in *ReceiveSalaryReq, opts ...grpc.CallOption) (*NobleSalaryInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NobleSalaryInfoResp)
	err := c.cc.Invoke(ctx, NobleSalary_GetReceiveInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NobleSalaryServer is the server API for NobleSalary service.
// All implementations must embed UnimplementedNobleSalaryServer
// for forward compatibility.
type NobleSalaryServer interface {
	// 发放俸禄(任务系统调用/每天)
	Distribute(context.Context, *EmptyRequest) (*DistributeSalaryResp, error)
	// 领取俸禄
	Receive(context.Context, *ReceiveSalaryReq) (*ReceiveSalaryResp, error)
	// 查询俸禄列表(按发放时间倒序)
	List(context.Context, *NobleSalaryListReq) (*NobleSalaryListResp, error)
	// 查询俸禄领信息(金额、状态等)
	GetReceiveInfo(context.Context, *ReceiveSalaryReq) (*NobleSalaryInfoResp, error)
	mustEmbedUnimplementedNobleSalaryServer()
}

// UnimplementedNobleSalaryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNobleSalaryServer struct{}

func (UnimplementedNobleSalaryServer) Distribute(context.Context, *EmptyRequest) (*DistributeSalaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Distribute not implemented")
}
func (UnimplementedNobleSalaryServer) Receive(context.Context, *ReceiveSalaryReq) (*ReceiveSalaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedNobleSalaryServer) List(context.Context, *NobleSalaryListReq) (*NobleSalaryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNobleSalaryServer) GetReceiveInfo(context.Context, *ReceiveSalaryReq) (*NobleSalaryInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceiveInfo not implemented")
}
func (UnimplementedNobleSalaryServer) mustEmbedUnimplementedNobleSalaryServer() {}
func (UnimplementedNobleSalaryServer) testEmbeddedByValue()                     {}

// UnsafeNobleSalaryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NobleSalaryServer will
// result in compilation errors.
type UnsafeNobleSalaryServer interface {
	mustEmbedUnimplementedNobleSalaryServer()
}

func RegisterNobleSalaryServer(s grpc.ServiceRegistrar, srv NobleSalaryServer) {
	// If the following call pancis, it indicates UnimplementedNobleSalaryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NobleSalary_ServiceDesc, srv)
}

func _NobleSalary_Distribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleSalaryServer).Distribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleSalary_Distribute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleSalaryServer).Distribute(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleSalary_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveSalaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleSalaryServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleSalary_Receive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleSalaryServer).Receive(ctx, req.(*ReceiveSalaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleSalary_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NobleSalaryListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleSalaryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleSalary_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleSalaryServer).List(ctx, req.(*NobleSalaryListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleSalary_GetReceiveInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveSalaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleSalaryServer).GetReceiveInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleSalary_GetReceiveInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleSalaryServer).GetReceiveInfo(ctx, req.(*ReceiveSalaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NobleSalary_ServiceDesc is the grpc.ServiceDesc for NobleSalary service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NobleSalary_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.vip.NobleSalary",
	HandlerType: (*NobleSalaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Distribute",
			Handler:    _NobleSalary_Distribute_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _NobleSalary_Receive_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NobleSalary_List_Handler,
		},
		{
			MethodName: "GetReceiveInfo",
			Handler:    _NobleSalary_GetReceiveInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.vip/noble_salarys.proto",
}
