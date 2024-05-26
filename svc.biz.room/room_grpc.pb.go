// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: svc.biz.room/room.proto

package room

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
	Room_CreateRoom_FullMethodName             = "/svc.biz.room.Room/CreateRoom"
	Room_UpdateRoom_FullMethodName             = "/svc.biz.room.Room/UpdateRoom"
	Room_GetRoom_FullMethodName                = "/svc.biz.room.Room/GetRoom"
	Room_MGetRooms_FullMethodName              = "/svc.biz.room.Room/MGetRooms"
	Room_MGetRoomsByStreamerIDs_FullMethodName = "/svc.biz.room.Room/MGetRoomsByStreamerIDs"
	Room_GetRoomList_FullMethodName            = "/svc.biz.room.Room/GetRoomList"
	Room_GetOnlineRoomList_FullMethodName      = "/svc.biz.room.Room/GetOnlineRoomList"
	Room_StartLive_FullMethodName              = "/svc.biz.room.Room/StartLive"
	Room_StopLive_FullMethodName               = "/svc.biz.room.Room/StopLive"
)

// RoomClient is the client API for Room service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomClient interface {
	// CreateRoom 创建房间
	CreateRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*CreateRoomResp, error)
	// UpdateRoom 更新房间
	UpdateRoom(ctx context.Context, in *UpdateRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetRoom 查询房间
	GetRoom(ctx context.Context, in *GetRoomReq, opts ...grpc.CallOption) (*GetRoomResp, error)
	// MGetRooms 查询房间
	MGetRooms(ctx context.Context, in *MGetRoomsReq, opts ...grpc.CallOption) (*MGetRoomsResp, error)
	// MGetRoomByStreamerIDs 批量查询房间
	MGetRoomsByStreamerIDs(ctx context.Context, in *MGetRoomsByStreamerIDsReq, opts ...grpc.CallOption) (*MGetRoomsByStreamerIDsResp, error)
	// GetRoomList 查询房间列表（后台使用此接口）
	GetRoomList(ctx context.Context, in *GetRoomListReq, opts ...grpc.CallOption) (*GetRoomListResp, error)
	// GetOnlineRoomList 查询在线房间列表（用户端列表使用此接口）
	GetOnlineRoomList(ctx context.Context, in *GetOnlineRoomListReq, opts ...grpc.CallOption) (*GetOnlineRoomListResp, error)
	// StartLive 开始直播
	StartLive(ctx context.Context, in *StartLiveReq, opts ...grpc.CallOption) (*StartLiveResp, error)
	// StopLive 关闭直播
	StopLive(ctx context.Context, in *StopLiveReq, opts ...grpc.CallOption) (*StopLiveResp, error)
}

type roomClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomClient(cc grpc.ClientConnInterface) RoomClient {
	return &roomClient{cc}
}

func (c *roomClient) CreateRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*CreateRoomResp, error) {
	out := new(CreateRoomResp)
	err := c.cc.Invoke(ctx, Room_CreateRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) UpdateRoom(ctx context.Context, in *UpdateRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Room_UpdateRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) GetRoom(ctx context.Context, in *GetRoomReq, opts ...grpc.CallOption) (*GetRoomResp, error) {
	out := new(GetRoomResp)
	err := c.cc.Invoke(ctx, Room_GetRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) MGetRooms(ctx context.Context, in *MGetRoomsReq, opts ...grpc.CallOption) (*MGetRoomsResp, error) {
	out := new(MGetRoomsResp)
	err := c.cc.Invoke(ctx, Room_MGetRooms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) MGetRoomsByStreamerIDs(ctx context.Context, in *MGetRoomsByStreamerIDsReq, opts ...grpc.CallOption) (*MGetRoomsByStreamerIDsResp, error) {
	out := new(MGetRoomsByStreamerIDsResp)
	err := c.cc.Invoke(ctx, Room_MGetRoomsByStreamerIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) GetRoomList(ctx context.Context, in *GetRoomListReq, opts ...grpc.CallOption) (*GetRoomListResp, error) {
	out := new(GetRoomListResp)
	err := c.cc.Invoke(ctx, Room_GetRoomList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) GetOnlineRoomList(ctx context.Context, in *GetOnlineRoomListReq, opts ...grpc.CallOption) (*GetOnlineRoomListResp, error) {
	out := new(GetOnlineRoomListResp)
	err := c.cc.Invoke(ctx, Room_GetOnlineRoomList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) StartLive(ctx context.Context, in *StartLiveReq, opts ...grpc.CallOption) (*StartLiveResp, error) {
	out := new(StartLiveResp)
	err := c.cc.Invoke(ctx, Room_StartLive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) StopLive(ctx context.Context, in *StopLiveReq, opts ...grpc.CallOption) (*StopLiveResp, error) {
	out := new(StopLiveResp)
	err := c.cc.Invoke(ctx, Room_StopLive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServer is the server API for Room service.
// All implementations must embed UnimplementedRoomServer
// for forward compatibility
type RoomServer interface {
	// CreateRoom 创建房间
	CreateRoom(context.Context, *CreateRoomReq) (*CreateRoomResp, error)
	// UpdateRoom 更新房间
	UpdateRoom(context.Context, *UpdateRoomReq) (*emptypb.Empty, error)
	// GetRoom 查询房间
	GetRoom(context.Context, *GetRoomReq) (*GetRoomResp, error)
	// MGetRooms 查询房间
	MGetRooms(context.Context, *MGetRoomsReq) (*MGetRoomsResp, error)
	// MGetRoomByStreamerIDs 批量查询房间
	MGetRoomsByStreamerIDs(context.Context, *MGetRoomsByStreamerIDsReq) (*MGetRoomsByStreamerIDsResp, error)
	// GetRoomList 查询房间列表（后台使用此接口）
	GetRoomList(context.Context, *GetRoomListReq) (*GetRoomListResp, error)
	// GetOnlineRoomList 查询在线房间列表（用户端列表使用此接口）
	GetOnlineRoomList(context.Context, *GetOnlineRoomListReq) (*GetOnlineRoomListResp, error)
	// StartLive 开始直播
	StartLive(context.Context, *StartLiveReq) (*StartLiveResp, error)
	// StopLive 关闭直播
	StopLive(context.Context, *StopLiveReq) (*StopLiveResp, error)
	mustEmbedUnimplementedRoomServer()
}

// UnimplementedRoomServer must be embedded to have forward compatible implementations.
type UnimplementedRoomServer struct {
}

func (UnimplementedRoomServer) CreateRoom(context.Context, *CreateRoomReq) (*CreateRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoomServer) UpdateRoom(context.Context, *UpdateRoomReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}
func (UnimplementedRoomServer) GetRoom(context.Context, *GetRoomReq) (*GetRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedRoomServer) MGetRooms(context.Context, *MGetRoomsReq) (*MGetRoomsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetRooms not implemented")
}
func (UnimplementedRoomServer) MGetRoomsByStreamerIDs(context.Context, *MGetRoomsByStreamerIDsReq) (*MGetRoomsByStreamerIDsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetRoomsByStreamerIDs not implemented")
}
func (UnimplementedRoomServer) GetRoomList(context.Context, *GetRoomListReq) (*GetRoomListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomList not implemented")
}
func (UnimplementedRoomServer) GetOnlineRoomList(context.Context, *GetOnlineRoomListReq) (*GetOnlineRoomListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnlineRoomList not implemented")
}
func (UnimplementedRoomServer) StartLive(context.Context, *StartLiveReq) (*StartLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartLive not implemented")
}
func (UnimplementedRoomServer) StopLive(context.Context, *StopLiveReq) (*StopLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopLive not implemented")
}
func (UnimplementedRoomServer) mustEmbedUnimplementedRoomServer() {}

// UnsafeRoomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServer will
// result in compilation errors.
type UnsafeRoomServer interface {
	mustEmbedUnimplementedRoomServer()
}

func RegisterRoomServer(s grpc.ServiceRegistrar, srv RoomServer) {
	s.RegisterService(&Room_ServiceDesc, srv)
}

func _Room_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).CreateRoom(ctx, req.(*CreateRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_UpdateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).UpdateRoom(ctx, req.(*UpdateRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_GetRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).GetRoom(ctx, req.(*GetRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_MGetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetRoomsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).MGetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_MGetRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).MGetRooms(ctx, req.(*MGetRoomsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_MGetRoomsByStreamerIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetRoomsByStreamerIDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).MGetRoomsByStreamerIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_MGetRoomsByStreamerIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).MGetRoomsByStreamerIDs(ctx, req.(*MGetRoomsByStreamerIDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_GetRoomList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).GetRoomList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_GetRoomList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).GetRoomList(ctx, req.(*GetRoomListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_GetOnlineRoomList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOnlineRoomListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).GetOnlineRoomList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_GetOnlineRoomList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).GetOnlineRoomList(ctx, req.(*GetOnlineRoomListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_StartLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).StartLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_StartLive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).StartLive(ctx, req.(*StartLiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_StopLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopLiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).StopLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_StopLive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).StopLive(ctx, req.(*StopLiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Room_ServiceDesc is the grpc.ServiceDesc for Room service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Room_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.room.Room",
	HandlerType: (*RoomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _Room_CreateRoom_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _Room_UpdateRoom_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _Room_GetRoom_Handler,
		},
		{
			MethodName: "MGetRooms",
			Handler:    _Room_MGetRooms_Handler,
		},
		{
			MethodName: "MGetRoomsByStreamerIDs",
			Handler:    _Room_MGetRoomsByStreamerIDs_Handler,
		},
		{
			MethodName: "GetRoomList",
			Handler:    _Room_GetRoomList_Handler,
		},
		{
			MethodName: "GetOnlineRoomList",
			Handler:    _Room_GetOnlineRoomList_Handler,
		},
		{
			MethodName: "StartLive",
			Handler:    _Room_StartLive_Handler,
		},
		{
			MethodName: "StopLive",
			Handler:    _Room_StopLive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.room/room.proto",
}