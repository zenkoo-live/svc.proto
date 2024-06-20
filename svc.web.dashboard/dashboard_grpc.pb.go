// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.3
// source: svc.web.dashboard/dashboard.proto

package dashboard

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

// DashboardClient is the client API for Dashboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DashboardClient interface {
}

type dashboardClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardClient(cc grpc.ClientConnInterface) DashboardClient {
	return &dashboardClient{cc}
}

// DashboardServer is the server API for Dashboard service.
// All implementations must embed UnimplementedDashboardServer
// for forward compatibility
type DashboardServer interface {
	mustEmbedUnimplementedDashboardServer()
}

// UnimplementedDashboardServer must be embedded to have forward compatible implementations.
type UnimplementedDashboardServer struct {
}

func (UnimplementedDashboardServer) mustEmbedUnimplementedDashboardServer() {}

// UnsafeDashboardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DashboardServer will
// result in compilation errors.
type UnsafeDashboardServer interface {
	mustEmbedUnimplementedDashboardServer()
}

func RegisterDashboardServer(s grpc.ServiceRegistrar, srv DashboardServer) {
	s.RegisterService(&Dashboard_ServiceDesc, srv)
}

// Dashboard_ServiceDesc is the grpc.ServiceDesc for Dashboard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dashboard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.web.dashboard.Dashboard",
	HandlerType: (*DashboardServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "svc.web.dashboard/dashboard.proto",
}
