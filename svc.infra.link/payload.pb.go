// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: svc.infra.link/payload.proto

package link

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PriorityType 消息权重
type PriorityType int32

const (
	PriorityType_Low    PriorityType = 0   // 低
	PriorityType_Middle PriorityType = 5   // 中
	PriorityType_High   PriorityType = 10  // 高
	PriorityType_Ultra  PriorityType = 100 // 超高
)

// Enum value maps for PriorityType.
var (
	PriorityType_name = map[int32]string{
		0:   "Low",
		5:   "Middle",
		10:  "High",
		100: "Ultra",
	}
	PriorityType_value = map[string]int32{
		"Low":    0,
		"Middle": 5,
		"High":   10,
		"Ultra":  100,
	}
)

func (x PriorityType) Enum() *PriorityType {
	p := new(PriorityType)
	*p = x
	return p
}

func (x PriorityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PriorityType) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_infra_link_payload_proto_enumTypes[0].Descriptor()
}

func (PriorityType) Type() protoreflect.EnumType {
	return &file_svc_infra_link_payload_proto_enumTypes[0]
}

func (x PriorityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PriorityType.Descriptor instead.
func (PriorityType) EnumDescriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{0}
}

// 消息命令类型
type CommandType int32

const (
	CommandType_UnKnow             CommandType = 0  // 未知
	CommandType_CommandMsgDownward CommandType = 10 // 通用下行消息
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0:  "UnKnow",
		10: "CommandMsgDownward",
	}
	CommandType_value = map[string]int32{
		"UnKnow":             0,
		"CommandMsgDownward": 10,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_infra_link_payload_proto_enumTypes[1].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_svc_infra_link_payload_proto_enumTypes[1]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{1}
}

// 业务消息类型
type PayloadType int32

const (
	PayloadType_StreamerDm      PayloadType = 0  // 主播弹幕
	PayloadType_UserCommDm      PayloadType = 1  // 普通弹幕
	PayloadType_UserVipDm       PayloadType = 2  // 付费弹幕
	PayloadType_UserGift        PayloadType = 3  // 用户送礼
	PayloadType_UserOpenFans    PayloadType = 4  // 开通粉丝团
	PayloadType_UserOpenGz      PayloadType = 5  // 开通贵族
	PayloadType_EnterRoom       PayloadType = 6  // 用户进入房间
	PayloadType_StreamerNotice  PayloadType = 7  // 主播公告通知
	PayloadType_StreamerCard    PayloadType = 8  // 主播名片通知
	PayloadType_UserKickRoom    PayloadType = 13 // 踢出房间
	PayloadType_StreamerOffline PayloadType = 14 // 主播下播
	PayloadType_StreamerOnline  PayloadType = 15 // 主播开播
	PayloadType_StreamerPause   PayloadType = 16 // 主播暂停
	PayloadType_StreamerResume  PayloadType = 17 // 主播恢复
)

// Enum value maps for PayloadType.
var (
	PayloadType_name = map[int32]string{
		0:  "StreamerDm",
		1:  "UserCommDm",
		2:  "UserVipDm",
		3:  "UserGift",
		4:  "UserOpenFans",
		5:  "UserOpenGz",
		6:  "EnterRoom",
		7:  "StreamerNotice",
		8:  "StreamerCard",
		13: "UserKickRoom",
		14: "StreamerOffline",
		15: "StreamerOnline",
		16: "StreamerPause",
		17: "StreamerResume",
	}
	PayloadType_value = map[string]int32{
		"StreamerDm":      0,
		"UserCommDm":      1,
		"UserVipDm":       2,
		"UserGift":        3,
		"UserOpenFans":    4,
		"UserOpenGz":      5,
		"EnterRoom":       6,
		"StreamerNotice":  7,
		"StreamerCard":    8,
		"UserKickRoom":    13,
		"StreamerOffline": 14,
		"StreamerOnline":  15,
		"StreamerPause":   16,
		"StreamerResume":  17,
	}
)

func (x PayloadType) Enum() *PayloadType {
	p := new(PayloadType)
	*p = x
	return p
}

func (x PayloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_infra_link_payload_proto_enumTypes[2].Descriptor()
}

func (PayloadType) Type() protoreflect.EnumType {
	return &file_svc_infra_link_payload_proto_enumTypes[2]
}

func (x PayloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayloadType.Descriptor instead.
func (PayloadType) EnumDescriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{2}
}

// 进入房间
type PayloadEnterRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`                                    // 用户UID
	Nickname   string        `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`                          // 昵称
	Text       string        `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`                                  // 聊天内容
	Avatar     string        `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`                              // 头像地址
	IsRoomAdm  bool          `protobuf:"varint,5,opt,name=is_room_adm,json=isRoomAdm,proto3" json:"is_room_adm,omitempty"`    // 房管
	IsSuperAdm bool          `protobuf:"varint,6,opt,name=is_super_adm,json=isSuperAdm,proto3" json:"is_super_adm,omitempty"` // 是否超管
	Charm      *PayloadCharm `protobuf:"bytes,7,opt,name=charm,proto3" json:"charm,omitempty"`                                // 魅力值
	Fans       *PayloadFans  `protobuf:"bytes,8,opt,name=fans,proto3" json:"fans,omitempty"`                                  // 粉丝信息
	Guard      *PayloadGuard `protobuf:"bytes,9,opt,name=guard,proto3" json:"guard,omitempty"`                                // 守护信息
}

func (x *PayloadEnterRoom) Reset() {
	*x = PayloadEnterRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadEnterRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadEnterRoom) ProtoMessage() {}

func (x *PayloadEnterRoom) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadEnterRoom.ProtoReflect.Descriptor instead.
func (*PayloadEnterRoom) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{0}
}

func (x *PayloadEnterRoom) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *PayloadEnterRoom) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PayloadEnterRoom) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PayloadEnterRoom) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *PayloadEnterRoom) GetIsRoomAdm() bool {
	if x != nil {
		return x.IsRoomAdm
	}
	return false
}

func (x *PayloadEnterRoom) GetIsSuperAdm() bool {
	if x != nil {
		return x.IsSuperAdm
	}
	return false
}

func (x *PayloadEnterRoom) GetCharm() *PayloadCharm {
	if x != nil {
		return x.Charm
	}
	return nil
}

func (x *PayloadEnterRoom) GetFans() *PayloadFans {
	if x != nil {
		return x.Fans
	}
	return nil
}

func (x *PayloadEnterRoom) GetGuard() *PayloadGuard {
	if x != nil {
		return x.Guard
	}
	return nil
}

// 公告
type PayloadStreamNotice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"` // 公告内容
}

func (x *PayloadStreamNotice) Reset() {
	*x = PayloadStreamNotice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadStreamNotice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadStreamNotice) ProtoMessage() {}

func (x *PayloadStreamNotice) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadStreamNotice.ProtoReflect.Descriptor instead.
func (*PayloadStreamNotice) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{1}
}

func (x *PayloadStreamNotice) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// 名片包装
type PayloadStreamerCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PayloadStreamerCard) Reset() {
	*x = PayloadStreamerCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadStreamerCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadStreamerCard) ProtoMessage() {}

func (x *PayloadStreamerCard) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadStreamerCard.ProtoReflect.Descriptor instead.
func (*PayloadStreamerCard) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{2}
}

// Payload容器
type PayloadWrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I int64  `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
	T int32  `protobuf:"varint,2,opt,name=t,proto3" json:"t,omitempty"`
	P int32  `protobuf:"varint,3,opt,name=p,proto3" json:"p,omitempty"`
	D string `protobuf:"bytes,11,opt,name=d,proto3" json:"d,omitempty"`
}

func (x *PayloadWrap) Reset() {
	*x = PayloadWrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_payload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadWrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadWrap) ProtoMessage() {}

func (x *PayloadWrap) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_payload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadWrap.ProtoReflect.Descriptor instead.
func (*PayloadWrap) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{3}
}

func (x *PayloadWrap) GetI() int64 {
	if x != nil {
		return x.I
	}
	return 0
}

func (x *PayloadWrap) GetT() int32 {
	if x != nil {
		return x.T
	}
	return 0
}

func (x *PayloadWrap) GetP() int32 {
	if x != nil {
		return x.P
	}
	return 0
}

func (x *PayloadWrap) GetD() string {
	if x != nil {
		return x.D
	}
	return ""
}

// 守护消息
type PayloadGuard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PayloadGuard) Reset() {
	*x = PayloadGuard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_payload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadGuard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadGuard) ProtoMessage() {}

func (x *PayloadGuard) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_payload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadGuard.ProtoReflect.Descriptor instead.
func (*PayloadGuard) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{4}
}

// 粉丝消息
type PayloadFans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PayloadFans) Reset() {
	*x = PayloadFans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_payload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadFans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadFans) ProtoMessage() {}

func (x *PayloadFans) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_payload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadFans.ProtoReflect.Descriptor instead.
func (*PayloadFans) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{5}
}

// 魅力值
type PayloadCharm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PayloadCharm) Reset() {
	*x = PayloadCharm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_payload_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadCharm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadCharm) ProtoMessage() {}

func (x *PayloadCharm) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_payload_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadCharm.ProtoReflect.Descriptor instead.
func (*PayloadCharm) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{6}
}

// 主播聊天弹幕消息
type PayloadStreamerDm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`           // 用户UID
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"` // 昵称
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`         // 聊天内容
	Avatar   string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`     // 头像地址
}

func (x *PayloadStreamerDm) Reset() {
	*x = PayloadStreamerDm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_payload_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadStreamerDm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadStreamerDm) ProtoMessage() {}

func (x *PayloadStreamerDm) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_payload_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadStreamerDm.ProtoReflect.Descriptor instead.
func (*PayloadStreamerDm) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{7}
}

func (x *PayloadStreamerDm) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *PayloadStreamerDm) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PayloadStreamerDm) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PayloadStreamerDm) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

// 用户普通聊天消息
type PayloadUserCommDm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`                                    // 用户UID
	Nickname   string        `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`                          // 昵称
	Text       string        `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`                                  // 聊天内容
	Avatar     string        `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`                              // 头像地址
	IsRoomAdm  bool          `protobuf:"varint,5,opt,name=is_room_adm,json=isRoomAdm,proto3" json:"is_room_adm,omitempty"`    // 房管
	IsSuperAdm bool          `protobuf:"varint,6,opt,name=is_super_adm,json=isSuperAdm,proto3" json:"is_super_adm,omitempty"` // 是否超管
	Charm      *PayloadCharm `protobuf:"bytes,7,opt,name=charm,proto3" json:"charm,omitempty"`                                // 魅力值
	Fans       *PayloadFans  `protobuf:"bytes,8,opt,name=fans,proto3" json:"fans,omitempty"`                                  // 粉丝信息
	Guard      *PayloadGuard `protobuf:"bytes,9,opt,name=guard,proto3" json:"guard,omitempty"`                                // 守护信息
}

func (x *PayloadUserCommDm) Reset() {
	*x = PayloadUserCommDm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_payload_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadUserCommDm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadUserCommDm) ProtoMessage() {}

func (x *PayloadUserCommDm) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_payload_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadUserCommDm.ProtoReflect.Descriptor instead.
func (*PayloadUserCommDm) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_payload_proto_rawDescGZIP(), []int{8}
}

func (x *PayloadUserCommDm) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *PayloadUserCommDm) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PayloadUserCommDm) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PayloadUserCommDm) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *PayloadUserCommDm) GetIsRoomAdm() bool {
	if x != nil {
		return x.IsRoomAdm
	}
	return false
}

func (x *PayloadUserCommDm) GetIsSuperAdm() bool {
	if x != nil {
		return x.IsSuperAdm
	}
	return false
}

func (x *PayloadUserCommDm) GetCharm() *PayloadCharm {
	if x != nil {
		return x.Charm
	}
	return nil
}

func (x *PayloadUserCommDm) GetFans() *PayloadFans {
	if x != nil {
		return x.Fans
	}
	return nil
}

func (x *PayloadUserCommDm) GetGuard() *PayloadGuard {
	if x != nil {
		return x.Guard
	}
	return nil
}

var File_svc_infra_link_payload_proto protoreflect.FileDescriptor

var file_svc_infra_link_payload_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xc7,
	0x02, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1e, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x6d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x64, 0x6d, 0x12, 0x20, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x75, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x12,
	0x32, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x61, 0x72, 0x6d, 0x52, 0x05, 0x63, 0x68,
	0x61, 0x72, 0x6d, 0x12, 0x2f, 0x0a, 0x04, 0x66, 0x61, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x61, 0x6e, 0x73, 0x52, 0x04,
	0x66, 0x61, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x67, 0x75, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x47, 0x75, 0x61, 0x72,
	0x64, 0x52, 0x05, 0x67, 0x75, 0x61, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x22, 0x45, 0x0a, 0x0b, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x72, 0x61, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x69, 0x12, 0x0c, 0x0a, 0x01, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x47, 0x75, 0x61, 0x72,
	0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x61, 0x6e, 0x73,
	0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x61, 0x72, 0x6d,
	0x22, 0x6d, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x44, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22,
	0xc8, 0x02, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6d, 0x6d, 0x44, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x6d, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x64, 0x6d, 0x12,
	0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x75, 0x70, 0x65, 0x72, 0x41, 0x64,
	0x6d, 0x12, 0x32, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x61, 0x72, 0x6d, 0x52, 0x05,
	0x63, 0x68, 0x61, 0x72, 0x6d, 0x12, 0x2f, 0x0a, 0x04, 0x66, 0x61, 0x6e, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x61, 0x6e, 0x73,
	0x52, 0x04, 0x66, 0x61, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x67, 0x75, 0x61, 0x72, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x47, 0x75,
	0x61, 0x72, 0x64, 0x52, 0x05, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2a, 0x38, 0x0a, 0x0c, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x6f,
	0x77, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x10, 0x05, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x69, 0x67, 0x68, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x6c, 0x74,
	0x72, 0x61, 0x10, 0x64, 0x2a, 0x31, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x44, 0x6f, 0x77,
	0x6e, 0x77, 0x61, 0x72, 0x64, 0x10, 0x0a, 0x2a, 0x83, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x44, 0x6d, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6d, 0x6d, 0x44, 0x6d, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x56,
	0x69, 0x70, 0x44, 0x6d, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x47, 0x69,
	0x66, 0x74, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e,
	0x46, 0x61, 0x6e, 0x73, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x70,
	0x65, 0x6e, 0x47, 0x7a, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x6f, 0x6f, 0x6d, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x4b, 0x69, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x6d, 0x10, 0x0d, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x10, 0x0f, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x72, 0x50, 0x61, 0x75, 0x73, 0x65, 0x10, 0x10, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x10, 0x11, 0x42, 0x17, 0x5a,
	0x15, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x3b, 0x6c, 0x69, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_infra_link_payload_proto_rawDescOnce sync.Once
	file_svc_infra_link_payload_proto_rawDescData = file_svc_infra_link_payload_proto_rawDesc
)

func file_svc_infra_link_payload_proto_rawDescGZIP() []byte {
	file_svc_infra_link_payload_proto_rawDescOnce.Do(func() {
		file_svc_infra_link_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_infra_link_payload_proto_rawDescData)
	})
	return file_svc_infra_link_payload_proto_rawDescData
}

var file_svc_infra_link_payload_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_svc_infra_link_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_svc_infra_link_payload_proto_goTypes = []any{
	(PriorityType)(0),           // 0: svc.infra.link.PriorityType
	(CommandType)(0),            // 1: svc.infra.link.CommandType
	(PayloadType)(0),            // 2: svc.infra.link.PayloadType
	(*PayloadEnterRoom)(nil),    // 3: svc.infra.link.PayloadEnterRoom
	(*PayloadStreamNotice)(nil), // 4: svc.infra.link.PayloadStreamNotice
	(*PayloadStreamerCard)(nil), // 5: svc.infra.link.PayloadStreamerCard
	(*PayloadWrap)(nil),         // 6: svc.infra.link.PayloadWrap
	(*PayloadGuard)(nil),        // 7: svc.infra.link.PayloadGuard
	(*PayloadFans)(nil),         // 8: svc.infra.link.PayloadFans
	(*PayloadCharm)(nil),        // 9: svc.infra.link.PayloadCharm
	(*PayloadStreamerDm)(nil),   // 10: svc.infra.link.PayloadStreamerDm
	(*PayloadUserCommDm)(nil),   // 11: svc.infra.link.PayloadUserCommDm
}
var file_svc_infra_link_payload_proto_depIdxs = []int32{
	9, // 0: svc.infra.link.PayloadEnterRoom.charm:type_name -> svc.infra.link.PayloadCharm
	8, // 1: svc.infra.link.PayloadEnterRoom.fans:type_name -> svc.infra.link.PayloadFans
	7, // 2: svc.infra.link.PayloadEnterRoom.guard:type_name -> svc.infra.link.PayloadGuard
	9, // 3: svc.infra.link.PayloadUserCommDm.charm:type_name -> svc.infra.link.PayloadCharm
	8, // 4: svc.infra.link.PayloadUserCommDm.fans:type_name -> svc.infra.link.PayloadFans
	7, // 5: svc.infra.link.PayloadUserCommDm.guard:type_name -> svc.infra.link.PayloadGuard
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_svc_infra_link_payload_proto_init() }
func file_svc_infra_link_payload_proto_init() {
	if File_svc_infra_link_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_infra_link_payload_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadEnterRoom); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_link_payload_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadStreamNotice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_link_payload_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadStreamerCard); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_link_payload_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadWrap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_link_payload_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadGuard); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_link_payload_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadFans); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_link_payload_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadCharm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_link_payload_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadStreamerDm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_link_payload_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadUserCommDm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_infra_link_payload_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_svc_infra_link_payload_proto_goTypes,
		DependencyIndexes: file_svc_infra_link_payload_proto_depIdxs,
		EnumInfos:         file_svc_infra_link_payload_proto_enumTypes,
		MessageInfos:      file_svc_infra_link_payload_proto_msgTypes,
	}.Build()
	File_svc_infra_link_payload_proto = out.File
	file_svc_infra_link_payload_proto_rawDesc = nil
	file_svc_infra_link_payload_proto_goTypes = nil
	file_svc_infra_link_payload_proto_depIdxs = nil
}
