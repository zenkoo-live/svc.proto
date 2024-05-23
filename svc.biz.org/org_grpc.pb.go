// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: svc.biz.org/org.proto

package org

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
	Org_InitDB_FullMethodName = "/svc.biz.org.Org/InitDB"
)

// OrgClient is the client API for Org service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrgClient interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error)
}

type orgClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgClient(cc grpc.ClientConnInterface) OrgClient {
	return &orgClient{cc}
}

func (c *orgClient) InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error) {
	out := new(InitDBResp)
	err := c.cc.Invoke(ctx, Org_InitDB_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgServer is the server API for Org service.
// All implementations must embed UnimplementedOrgServer
// for forward compatibility
type OrgServer interface {
	InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error)
	mustEmbedUnimplementedOrgServer()
}

// UnimplementedOrgServer must be embedded to have forward compatible implementations.
type UnimplementedOrgServer struct {
}

func (UnimplementedOrgServer) InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDB not implemented")
}
func (UnimplementedOrgServer) mustEmbedUnimplementedOrgServer() {}

// UnsafeOrgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrgServer will
// result in compilation errors.
type UnsafeOrgServer interface {
	mustEmbedUnimplementedOrgServer()
}

func RegisterOrgServer(s grpc.ServiceRegistrar, srv OrgServer) {
	s.RegisterService(&Org_ServiceDesc, srv)
}

func _Org_InitDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).InitDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_InitDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).InitDB(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Org_ServiceDesc is the grpc.ServiceDesc for Org service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Org_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.org.Org",
	HandlerType: (*OrgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDB",
			Handler:    _Org_InitDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.org/org.proto",
}
