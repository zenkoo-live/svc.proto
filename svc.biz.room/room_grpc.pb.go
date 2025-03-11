// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Room_CreateRoom_FullMethodName                           = "/svc.biz.room.Room/CreateRoom"
	Room_UpdateRoom_FullMethodName                           = "/svc.biz.room.Room/UpdateRoom"
	Room_GetRoom_FullMethodName                              = "/svc.biz.room.Room/GetRoom"
	Room_GetRoomByStreamerID_FullMethodName                  = "/svc.biz.room.Room/GetRoomByStreamerID"
	Room_MGetRooms_FullMethodName                            = "/svc.biz.room.Room/MGetRooms"
	Room_MGetRoomsByStreamerIDs_FullMethodName               = "/svc.biz.room.Room/MGetRoomsByStreamerIDs"
	Room_MGetRoomsByStreamerIDsWithOnlineSort_FullMethodName = "/svc.biz.room.Room/MGetRoomsByStreamerIDsWithOnlineSort"
	Room_GetRoomList_FullMethodName                          = "/svc.biz.room.Room/GetRoomList"
	Room_GetOnlineRoomList_FullMethodName                    = "/svc.biz.room.Room/GetOnlineRoomList"
	Room_GetRandomRooms_FullMethodName                       = "/svc.biz.room.Room/GetRandomRooms"
	Room_KickoutUserInRoom_FullMethodName                    = "/svc.biz.room.Room/KickoutUserInRoom"
	Room_ForbidRoom_FullMethodName                           = "/svc.biz.room.Room/ForbidRoom"
	Room_ResumeRoom_FullMethodName                           = "/svc.biz.room.Room/ResumeRoom"
	Room_MuteUserInType_FullMethodName                       = "/svc.biz.room.Room/MuteUserInType"
	Room_UNMuteUser_FullMethodName                           = "/svc.biz.room.Room/UNMuteUser"
	Room_GetMuteUserInfo_FullMethodName                      = "/svc.biz.room.Room/GetMuteUserInfo"
	Room_StartLive_FullMethodName                            = "/svc.biz.room.Room/StartLive"
	Room_StopLive_FullMethodName                             = "/svc.biz.room.Room/StopLive"
)

// RoomClient is the client API for Room service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Room 房间
type RoomClient interface {
	// CreateRoom 创建房间
	CreateRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*CreateRoomResp, error)
	// UpdateRoom 更新房间
	UpdateRoom(ctx context.Context, in *UpdateRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetRoom 查询房间
	GetRoom(ctx context.Context, in *GetRoomReq, opts ...grpc.CallOption) (*GetRoomResp, error)
	// GetRoomByStreamerID 查询房间
	GetRoomByStreamerID(ctx context.Context, in *GetRoomByStreamerIDReq, opts ...grpc.CallOption) (*GetRoomByStreamerIDResp, error)
	// MGetRooms 查询房间
	MGetRooms(ctx context.Context, in *MGetRoomsReq, opts ...grpc.CallOption) (*MGetRoomsResp, error)
	// MGetRoomByStreamerIDs 批量查询房间
	MGetRoomsByStreamerIDs(ctx context.Context, in *MGetRoomsByStreamerIDsReq, opts ...grpc.CallOption) (*MGetRoomsByStreamerIDsResp, error)
	// MGetRoomsByStreamerIDsWithOnlineSort 批量查询房间（带在线分页，按照传入顺序获取，在线排在最前）
	MGetRoomsByStreamerIDsWithOnlineSort(ctx context.Context, in *MGetRoomsByStreamerIDsWithOnlineSortReq, opts ...grpc.CallOption) (*MGetRoomsByStreamerIDsWithOnlineSortResp, error)
	// GetRoomList 查询房间列表（前台的搜索或者后台使用此接口）
	GetRoomList(ctx context.Context, in *GetRoomListReq, opts ...grpc.CallOption) (*GetRoomListResp, error)
	// GetOnlineRoomList 查询在线房间列表（用户端列表使用此接口）
	GetOnlineRoomList(ctx context.Context, in *GetOnlineRoomListReq, opts ...grpc.CallOption) (*GetOnlineRoomListResp, error)
	// GetRandomRooms 随机获取房间(用户端上滑获取下一个房间用)
	GetRandomRooms(ctx context.Context, in *GetRandomRoomsReq, opts ...grpc.CallOption) (*GetRoomListResp, error)
	// KickoutUserInRoom 踢出直播间用
	KickoutUserInRoom(ctx context.Context, in *KickoutUserInRoomReq, opts ...grpc.CallOption) (*KickoutUserInRoomResp, error)
	// ForbidRoom 封禁直播间
	ForbidRoom(ctx context.Context, in *ForbidRoomReq, opts ...grpc.CallOption) (*ForbidRoomResp, error)
	// ResumeRoom 解封直播间
	ResumeRoom(ctx context.Context, in *ResumeRoomReq, opts ...grpc.CallOption) (*ResumeRoomResp, error)
	// 禁言
	MuteUserInType(ctx context.Context, in *MuteUserReq, opts ...grpc.CallOption) (*MuteUserCommResp, error)
	// 取消禁言
	UNMuteUser(ctx context.Context, in *MuteUserReq, opts ...grpc.CallOption) (*MuteUserCommResp, error)
	// 获取禁言信息
	GetMuteUserInfo(ctx context.Context, in *GetMuteInfoReq, opts ...grpc.CallOption) (*GetMuteInfoResp, error)
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
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRoomResp)
	err := c.cc.Invoke(ctx, Room_CreateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) UpdateRoom(ctx context.Context, in *UpdateRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Room_UpdateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) GetRoom(ctx context.Context, in *GetRoomReq, opts ...grpc.CallOption) (*GetRoomResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoomResp)
	err := c.cc.Invoke(ctx, Room_GetRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) GetRoomByStreamerID(ctx context.Context, in *GetRoomByStreamerIDReq, opts ...grpc.CallOption) (*GetRoomByStreamerIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoomByStreamerIDResp)
	err := c.cc.Invoke(ctx, Room_GetRoomByStreamerID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) MGetRooms(ctx context.Context, in *MGetRoomsReq, opts ...grpc.CallOption) (*MGetRoomsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MGetRoomsResp)
	err := c.cc.Invoke(ctx, Room_MGetRooms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) MGetRoomsByStreamerIDs(ctx context.Context, in *MGetRoomsByStreamerIDsReq, opts ...grpc.CallOption) (*MGetRoomsByStreamerIDsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MGetRoomsByStreamerIDsResp)
	err := c.cc.Invoke(ctx, Room_MGetRoomsByStreamerIDs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) MGetRoomsByStreamerIDsWithOnlineSort(ctx context.Context, in *MGetRoomsByStreamerIDsWithOnlineSortReq, opts ...grpc.CallOption) (*MGetRoomsByStreamerIDsWithOnlineSortResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MGetRoomsByStreamerIDsWithOnlineSortResp)
	err := c.cc.Invoke(ctx, Room_MGetRoomsByStreamerIDsWithOnlineSort_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) GetRoomList(ctx context.Context, in *GetRoomListReq, opts ...grpc.CallOption) (*GetRoomListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoomListResp)
	err := c.cc.Invoke(ctx, Room_GetRoomList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) GetOnlineRoomList(ctx context.Context, in *GetOnlineRoomListReq, opts ...grpc.CallOption) (*GetOnlineRoomListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOnlineRoomListResp)
	err := c.cc.Invoke(ctx, Room_GetOnlineRoomList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) GetRandomRooms(ctx context.Context, in *GetRandomRoomsReq, opts ...grpc.CallOption) (*GetRoomListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoomListResp)
	err := c.cc.Invoke(ctx, Room_GetRandomRooms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) KickoutUserInRoom(ctx context.Context, in *KickoutUserInRoomReq, opts ...grpc.CallOption) (*KickoutUserInRoomResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KickoutUserInRoomResp)
	err := c.cc.Invoke(ctx, Room_KickoutUserInRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) ForbidRoom(ctx context.Context, in *ForbidRoomReq, opts ...grpc.CallOption) (*ForbidRoomResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ForbidRoomResp)
	err := c.cc.Invoke(ctx, Room_ForbidRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) ResumeRoom(ctx context.Context, in *ResumeRoomReq, opts ...grpc.CallOption) (*ResumeRoomResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResumeRoomResp)
	err := c.cc.Invoke(ctx, Room_ResumeRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) MuteUserInType(ctx context.Context, in *MuteUserReq, opts ...grpc.CallOption) (*MuteUserCommResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MuteUserCommResp)
	err := c.cc.Invoke(ctx, Room_MuteUserInType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) UNMuteUser(ctx context.Context, in *MuteUserReq, opts ...grpc.CallOption) (*MuteUserCommResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MuteUserCommResp)
	err := c.cc.Invoke(ctx, Room_UNMuteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) GetMuteUserInfo(ctx context.Context, in *GetMuteInfoReq, opts ...grpc.CallOption) (*GetMuteInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMuteInfoResp)
	err := c.cc.Invoke(ctx, Room_GetMuteUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) StartLive(ctx context.Context, in *StartLiveReq, opts ...grpc.CallOption) (*StartLiveResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartLiveResp)
	err := c.cc.Invoke(ctx, Room_StartLive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) StopLive(ctx context.Context, in *StopLiveReq, opts ...grpc.CallOption) (*StopLiveResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopLiveResp)
	err := c.cc.Invoke(ctx, Room_StopLive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServer is the server API for Room service.
// All implementations must embed UnimplementedRoomServer
// for forward compatibility.
//
// Room 房间
type RoomServer interface {
	// CreateRoom 创建房间
	CreateRoom(context.Context, *CreateRoomReq) (*CreateRoomResp, error)
	// UpdateRoom 更新房间
	UpdateRoom(context.Context, *UpdateRoomReq) (*emptypb.Empty, error)
	// GetRoom 查询房间
	GetRoom(context.Context, *GetRoomReq) (*GetRoomResp, error)
	// GetRoomByStreamerID 查询房间
	GetRoomByStreamerID(context.Context, *GetRoomByStreamerIDReq) (*GetRoomByStreamerIDResp, error)
	// MGetRooms 查询房间
	MGetRooms(context.Context, *MGetRoomsReq) (*MGetRoomsResp, error)
	// MGetRoomByStreamerIDs 批量查询房间
	MGetRoomsByStreamerIDs(context.Context, *MGetRoomsByStreamerIDsReq) (*MGetRoomsByStreamerIDsResp, error)
	// MGetRoomsByStreamerIDsWithOnlineSort 批量查询房间（带在线分页，按照传入顺序获取，在线排在最前）
	MGetRoomsByStreamerIDsWithOnlineSort(context.Context, *MGetRoomsByStreamerIDsWithOnlineSortReq) (*MGetRoomsByStreamerIDsWithOnlineSortResp, error)
	// GetRoomList 查询房间列表（前台的搜索或者后台使用此接口）
	GetRoomList(context.Context, *GetRoomListReq) (*GetRoomListResp, error)
	// GetOnlineRoomList 查询在线房间列表（用户端列表使用此接口）
	GetOnlineRoomList(context.Context, *GetOnlineRoomListReq) (*GetOnlineRoomListResp, error)
	// GetRandomRooms 随机获取房间(用户端上滑获取下一个房间用)
	GetRandomRooms(context.Context, *GetRandomRoomsReq) (*GetRoomListResp, error)
	// KickoutUserInRoom 踢出直播间用
	KickoutUserInRoom(context.Context, *KickoutUserInRoomReq) (*KickoutUserInRoomResp, error)
	// ForbidRoom 封禁直播间
	ForbidRoom(context.Context, *ForbidRoomReq) (*ForbidRoomResp, error)
	// ResumeRoom 解封直播间
	ResumeRoom(context.Context, *ResumeRoomReq) (*ResumeRoomResp, error)
	// 禁言
	MuteUserInType(context.Context, *MuteUserReq) (*MuteUserCommResp, error)
	// 取消禁言
	UNMuteUser(context.Context, *MuteUserReq) (*MuteUserCommResp, error)
	// 获取禁言信息
	GetMuteUserInfo(context.Context, *GetMuteInfoReq) (*GetMuteInfoResp, error)
	// StartLive 开始直播
	StartLive(context.Context, *StartLiveReq) (*StartLiveResp, error)
	// StopLive 关闭直播
	StopLive(context.Context, *StopLiveReq) (*StopLiveResp, error)
	mustEmbedUnimplementedRoomServer()
}

// UnimplementedRoomServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoomServer struct{}

func (UnimplementedRoomServer) CreateRoom(context.Context, *CreateRoomReq) (*CreateRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoomServer) UpdateRoom(context.Context, *UpdateRoomReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}
func (UnimplementedRoomServer) GetRoom(context.Context, *GetRoomReq) (*GetRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedRoomServer) GetRoomByStreamerID(context.Context, *GetRoomByStreamerIDReq) (*GetRoomByStreamerIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomByStreamerID not implemented")
}
func (UnimplementedRoomServer) MGetRooms(context.Context, *MGetRoomsReq) (*MGetRoomsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetRooms not implemented")
}
func (UnimplementedRoomServer) MGetRoomsByStreamerIDs(context.Context, *MGetRoomsByStreamerIDsReq) (*MGetRoomsByStreamerIDsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetRoomsByStreamerIDs not implemented")
}
func (UnimplementedRoomServer) MGetRoomsByStreamerIDsWithOnlineSort(context.Context, *MGetRoomsByStreamerIDsWithOnlineSortReq) (*MGetRoomsByStreamerIDsWithOnlineSortResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetRoomsByStreamerIDsWithOnlineSort not implemented")
}
func (UnimplementedRoomServer) GetRoomList(context.Context, *GetRoomListReq) (*GetRoomListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomList not implemented")
}
func (UnimplementedRoomServer) GetOnlineRoomList(context.Context, *GetOnlineRoomListReq) (*GetOnlineRoomListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnlineRoomList not implemented")
}
func (UnimplementedRoomServer) GetRandomRooms(context.Context, *GetRandomRoomsReq) (*GetRoomListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomRooms not implemented")
}
func (UnimplementedRoomServer) KickoutUserInRoom(context.Context, *KickoutUserInRoomReq) (*KickoutUserInRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickoutUserInRoom not implemented")
}
func (UnimplementedRoomServer) ForbidRoom(context.Context, *ForbidRoomReq) (*ForbidRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForbidRoom not implemented")
}
func (UnimplementedRoomServer) ResumeRoom(context.Context, *ResumeRoomReq) (*ResumeRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeRoom not implemented")
}
func (UnimplementedRoomServer) MuteUserInType(context.Context, *MuteUserReq) (*MuteUserCommResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MuteUserInType not implemented")
}
func (UnimplementedRoomServer) UNMuteUser(context.Context, *MuteUserReq) (*MuteUserCommResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UNMuteUser not implemented")
}
func (UnimplementedRoomServer) GetMuteUserInfo(context.Context, *GetMuteInfoReq) (*GetMuteInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMuteUserInfo not implemented")
}
func (UnimplementedRoomServer) StartLive(context.Context, *StartLiveReq) (*StartLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartLive not implemented")
}
func (UnimplementedRoomServer) StopLive(context.Context, *StopLiveReq) (*StopLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopLive not implemented")
}
func (UnimplementedRoomServer) mustEmbedUnimplementedRoomServer() {}
func (UnimplementedRoomServer) testEmbeddedByValue()              {}

// UnsafeRoomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServer will
// result in compilation errors.
type UnsafeRoomServer interface {
	mustEmbedUnimplementedRoomServer()
}

func RegisterRoomServer(s grpc.ServiceRegistrar, srv RoomServer) {
	// If the following call pancis, it indicates UnimplementedRoomServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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

func _Room_GetRoomByStreamerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomByStreamerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).GetRoomByStreamerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_GetRoomByStreamerID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).GetRoomByStreamerID(ctx, req.(*GetRoomByStreamerIDReq))
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

func _Room_MGetRoomsByStreamerIDsWithOnlineSort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetRoomsByStreamerIDsWithOnlineSortReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).MGetRoomsByStreamerIDsWithOnlineSort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_MGetRoomsByStreamerIDsWithOnlineSort_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).MGetRoomsByStreamerIDsWithOnlineSort(ctx, req.(*MGetRoomsByStreamerIDsWithOnlineSortReq))
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

func _Room_GetRandomRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandomRoomsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).GetRandomRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_GetRandomRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).GetRandomRooms(ctx, req.(*GetRandomRoomsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_KickoutUserInRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickoutUserInRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).KickoutUserInRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_KickoutUserInRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).KickoutUserInRoom(ctx, req.(*KickoutUserInRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_ForbidRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForbidRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).ForbidRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_ForbidRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).ForbidRoom(ctx, req.(*ForbidRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_ResumeRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).ResumeRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_ResumeRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).ResumeRoom(ctx, req.(*ResumeRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_MuteUserInType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MuteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).MuteUserInType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_MuteUserInType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).MuteUserInType(ctx, req.(*MuteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_UNMuteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MuteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).UNMuteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_UNMuteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).UNMuteUser(ctx, req.(*MuteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_GetMuteUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMuteInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).GetMuteUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_GetMuteUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).GetMuteUserInfo(ctx, req.(*GetMuteInfoReq))
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
			MethodName: "GetRoomByStreamerID",
			Handler:    _Room_GetRoomByStreamerID_Handler,
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
			MethodName: "MGetRoomsByStreamerIDsWithOnlineSort",
			Handler:    _Room_MGetRoomsByStreamerIDsWithOnlineSort_Handler,
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
			MethodName: "GetRandomRooms",
			Handler:    _Room_GetRandomRooms_Handler,
		},
		{
			MethodName: "KickoutUserInRoom",
			Handler:    _Room_KickoutUserInRoom_Handler,
		},
		{
			MethodName: "ForbidRoom",
			Handler:    _Room_ForbidRoom_Handler,
		},
		{
			MethodName: "ResumeRoom",
			Handler:    _Room_ResumeRoom_Handler,
		},
		{
			MethodName: "MuteUserInType",
			Handler:    _Room_MuteUserInType_Handler,
		},
		{
			MethodName: "UNMuteUser",
			Handler:    _Room_UNMuteUser_Handler,
		},
		{
			MethodName: "GetMuteUserInfo",
			Handler:    _Room_GetMuteUserInfo_Handler,
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
