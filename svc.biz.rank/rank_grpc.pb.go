// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: svc.biz.rank/rank.proto

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
	Rank_GetLiveOnlineRankMember_FullMethodName = "/svc.biz.rank.Rank/GetLiveOnlineRankMember"
)

// RankClient is the client API for Rank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 排行
type RankClient interface {
	// 获取直播在线的排行榜成员
	GetLiveOnlineRankMember(ctx context.Context, in *GetLiveOnlineRankMemberReq, opts ...grpc.CallOption) (*GetLiveOnlineRankMemberResp, error)
}

type rankClient struct {
	cc grpc.ClientConnInterface
}

func NewRankClient(cc grpc.ClientConnInterface) RankClient {
	return &rankClient{cc}
}

func (c *rankClient) GetLiveOnlineRankMember(ctx context.Context, in *GetLiveOnlineRankMemberReq, opts ...grpc.CallOption) (*GetLiveOnlineRankMemberResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLiveOnlineRankMemberResp)
	err := c.cc.Invoke(ctx, Rank_GetLiveOnlineRankMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankServer is the server API for Rank service.
// All implementations must embed UnimplementedRankServer
// for forward compatibility
//
// 排行
type RankServer interface {
	// 获取直播在线的排行榜成员
	GetLiveOnlineRankMember(context.Context, *GetLiveOnlineRankMemberReq) (*GetLiveOnlineRankMemberResp, error)
	mustEmbedUnimplementedRankServer()
}

// UnimplementedRankServer must be embedded to have forward compatible implementations.
type UnimplementedRankServer struct {
}

func (UnimplementedRankServer) GetLiveOnlineRankMember(context.Context, *GetLiveOnlineRankMemberReq) (*GetLiveOnlineRankMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiveOnlineRankMember not implemented")
}
func (UnimplementedRankServer) mustEmbedUnimplementedRankServer() {}

// UnsafeRankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RankServer will
// result in compilation errors.
type UnsafeRankServer interface {
	mustEmbedUnimplementedRankServer()
}

func RegisterRankServer(s grpc.ServiceRegistrar, srv RankServer) {
	s.RegisterService(&Rank_ServiceDesc, srv)
}

func _Rank_GetLiveOnlineRankMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLiveOnlineRankMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).GetLiveOnlineRankMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rank_GetLiveOnlineRankMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).GetLiveOnlineRankMember(ctx, req.(*GetLiveOnlineRankMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Rank_ServiceDesc is the grpc.ServiceDesc for Rank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.rank.Rank",
	HandlerType: (*RankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLiveOnlineRankMember",
			Handler:    _Rank_GetLiveOnlineRankMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.rank/rank.proto",
}
