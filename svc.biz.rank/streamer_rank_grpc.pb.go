// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: svc.biz.rank/streamer_rank.proto

package rank

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	StreamerRank_GetStreamerRank_FullMethodName = "/svc.biz.rank.StreamerRank/GetStreamerRank"
)

// StreamerRankClient is the client API for StreamerRank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 主播排行
type StreamerRankClient interface {
	// 获取主播排行榜成员
	GetStreamerRank(ctx context.Context, in *GetStreamerRankReq, opts ...grpc.CallOption) (*GetStreamerRankResp, error)
}

type streamerRankClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamerRankClient(cc grpc.ClientConnInterface) StreamerRankClient {
	return &streamerRankClient{cc}
}

func (c *streamerRankClient) GetStreamerRank(ctx context.Context, in *GetStreamerRankReq, opts ...grpc.CallOption) (*GetStreamerRankResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStreamerRankResp)
	err := c.cc.Invoke(ctx, StreamerRank_GetStreamerRank_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamerRankServer is the server API for StreamerRank service.
// All implementations must embed UnimplementedStreamerRankServer
// for forward compatibility
//
// 主播排行
type StreamerRankServer interface {
	// 获取主播排行榜成员
	GetStreamerRank(context.Context, *GetStreamerRankReq) (*GetStreamerRankResp, error)
	mustEmbedUnimplementedStreamerRankServer()
}

// UnimplementedStreamerRankServer must be embedded to have forward compatible implementations.
type UnimplementedStreamerRankServer struct {
}

func (UnimplementedStreamerRankServer) GetStreamerRank(context.Context, *GetStreamerRankReq) (*GetStreamerRankResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamerRank not implemented")
}
func (UnimplementedStreamerRankServer) mustEmbedUnimplementedStreamerRankServer() {}

// UnsafeStreamerRankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamerRankServer will
// result in compilation errors.
type UnsafeStreamerRankServer interface {
	mustEmbedUnimplementedStreamerRankServer()
}

func RegisterStreamerRankServer(s grpc.ServiceRegistrar, srv StreamerRankServer) {
	s.RegisterService(&StreamerRank_ServiceDesc, srv)
}

func _StreamerRank_GetStreamerRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamerRankReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamerRankServer).GetStreamerRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamerRank_GetStreamerRank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamerRankServer).GetStreamerRank(ctx, req.(*GetStreamerRankReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamerRank_ServiceDesc is the grpc.ServiceDesc for StreamerRank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamerRank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.rank.StreamerRank",
	HandlerType: (*StreamerRankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStreamerRank",
			Handler:    _StreamerRank_GetStreamerRank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.rank/streamer_rank.proto",
}