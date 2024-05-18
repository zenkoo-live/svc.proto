// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: svc.web.streamer/streamer.proto

package streamer

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// StreamerClient is the client API for Streamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamerClient interface {
}

type streamerClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamerClient(cc grpc.ClientConnInterface) StreamerClient {
	return &streamerClient{cc}
}

// StreamerServer is the server API for Streamer service.
// All implementations must embed UnimplementedStreamerServer
// for forward compatibility
type StreamerServer interface {
	mustEmbedUnimplementedStreamerServer()
}

// UnimplementedStreamerServer must be embedded to have forward compatible implementations.
type UnimplementedStreamerServer struct {
}

func (UnimplementedStreamerServer) mustEmbedUnimplementedStreamerServer() {}

// UnsafeStreamerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamerServer will
// result in compilation errors.
type UnsafeStreamerServer interface {
	mustEmbedUnimplementedStreamerServer()
}

func RegisterStreamerServer(s grpc.ServiceRegistrar, srv StreamerServer) {
	s.RegisterService(&Streamer_ServiceDesc, srv)
}

// Streamer_ServiceDesc is the grpc.ServiceDesc for Streamer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Streamer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.web.streamer.Streamer",
	HandlerType: (*StreamerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "svc.web.streamer/streamer.proto",
}