// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: svc.infra.link/mq.proto

package link

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActivityEventType int32

const (
	ActivityEventType_EventUnknown ActivityEventType = 0
	ActivityEventType_EventOnline  ActivityEventType = 1
	ActivityEventType_EventOffline ActivityEventType = 2
	ActivityEventType_EventBind    ActivityEventType = 3
	ActivityEventType_EventUnbind  ActivityEventType = 4
	ActivityEventType_EventActive  ActivityEventType = 5
)

// Enum value maps for ActivityEventType.
var (
	ActivityEventType_name = map[int32]string{
		0: "EventUnknown",
		1: "EventOnline",
		2: "EventOffline",
		3: "EventBind",
		4: "EventUnbind",
		5: "EventActive",
	}
	ActivityEventType_value = map[string]int32{
		"EventUnknown": 0,
		"EventOnline":  1,
		"EventOffline": 2,
		"EventBind":    3,
		"EventUnbind":  4,
		"EventActive":  5,
	}
)

func (x ActivityEventType) Enum() *ActivityEventType {
	p := new(ActivityEventType)
	*p = x
	return p
}

func (x ActivityEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActivityEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_infra_link_mq_proto_enumTypes[0].Descriptor()
}

func (ActivityEventType) Type() protoreflect.EnumType {
	return &file_svc_infra_link_mq_proto_enumTypes[0]
}

func (x ActivityEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActivityEventType.Descriptor instead.
func (ActivityEventType) EnumDescriptor() ([]byte, []int) {
	return file_svc_infra_link_mq_proto_rawDescGZIP(), []int{0}
}

type ScopeType int32

const (
	ScopeType_ScopeGlobal    ScopeType = 0
	ScopeType_ScopeGroup     ScopeType = 1
	ScopeType_ScopeAccount   ScopeType = 2
	ScopeType_ScopeDevice    ScopeType = 3
	ScopeType_ScopeSession   ScopeType = 4
	ScopeType_ScopeKick      ScopeType = 125
	ScopeType_ScopeBlackhole ScopeType = 126
	ScopeType_ScopeUnknown   ScopeType = 127
)

// Enum value maps for ScopeType.
var (
	ScopeType_name = map[int32]string{
		0:   "ScopeGlobal",
		1:   "ScopeGroup",
		2:   "ScopeAccount",
		3:   "ScopeDevice",
		4:   "ScopeSession",
		125: "ScopeKick",
		126: "ScopeBlackhole",
		127: "ScopeUnknown",
	}
	ScopeType_value = map[string]int32{
		"ScopeGlobal":    0,
		"ScopeGroup":     1,
		"ScopeAccount":   2,
		"ScopeDevice":    3,
		"ScopeSession":   4,
		"ScopeKick":      125,
		"ScopeBlackhole": 126,
		"ScopeUnknown":   127,
	}
)

func (x ScopeType) Enum() *ScopeType {
	p := new(ScopeType)
	*p = x
	return p
}

func (x ScopeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScopeType) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_infra_link_mq_proto_enumTypes[1].Descriptor()
}

func (ScopeType) Type() protoreflect.EnumType {
	return &file_svc_infra_link_mq_proto_enumTypes[1]
}

func (x ScopeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScopeType.Descriptor instead.
func (ScopeType) EnumDescriptor() ([]byte, []int) {
	return file_svc_infra_link_mq_proto_rawDescGZIP(), []int{1}
}

type UserEnterQuitAction int32

const (
	UserEnterQuitAction_Enter UserEnterQuitAction = 0
	UserEnterQuitAction_Quit  UserEnterQuitAction = 1
)

// Enum value maps for UserEnterQuitAction.
var (
	UserEnterQuitAction_name = map[int32]string{
		0: "Enter",
		1: "Quit",
	}
	UserEnterQuitAction_value = map[string]int32{
		"Enter": 0,
		"Quit":  1,
	}
)

func (x UserEnterQuitAction) Enum() *UserEnterQuitAction {
	p := new(UserEnterQuitAction)
	*p = x
	return p
}

func (x UserEnterQuitAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserEnterQuitAction) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_infra_link_mq_proto_enumTypes[2].Descriptor()
}

func (UserEnterQuitAction) Type() protoreflect.EnumType {
	return &file_svc_infra_link_mq_proto_enumTypes[2]
}

func (x UserEnterQuitAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserEnterQuitAction.Descriptor instead.
func (UserEnterQuitAction) EnumDescriptor() ([]byte, []int) {
	return file_svc_infra_link_mq_proto_rawDescGZIP(), []int{2}
}

type ActivityMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64             `protobuf:"varint,1,opt,name=id,json=i,proto3" json:"id,omitempty"`
	Token       string            `protobuf:"bytes,2,opt,name=token,json=t,proto3" json:"token,omitempty"`
	Event       ActivityEventType `protobuf:"varint,3,opt,name=event,json=e,proto3,enum=svc.infra.link.ActivityEventType" json:"event,omitempty"`
	Svc         string            `protobuf:"bytes,4,opt,name=svc,json=c,proto3" json:"svc,omitempty"`
	GatewayId   string            `protobuf:"bytes,5,opt,name=gateway_id,json=n,proto3" json:"gateway_id,omitempty"`
	SessionId   string            `protobuf:"bytes,6,opt,name=session_id,json=s,proto3" json:"session_id,omitempty"`
	RemoteAddr  string            `protobuf:"bytes,7,opt,name=remote_addr,json=r,proto3" json:"remote_addr,omitempty"`
	Addition    string            `protobuf:"bytes,8,opt,name=addition,json=x,proto3" json:"addition,omitempty"`
	MerchantId  string            `protobuf:"bytes,9,opt,name=merchant_id,json=m,proto3" json:"merchant_id,omitempty"`
	AccountId   string            `protobuf:"bytes,10,opt,name=account_id,json=a,proto3" json:"account_id,omitempty"`
	GroupId     string            `protobuf:"bytes,11,opt,name=group_id,json=g,proto3" json:"group_id,omitempty"`
	LastGroupId string            `protobuf:"bytes,12,opt,name=last_group_id,json=l,proto3" json:"last_group_id,omitempty"`
	Device      string            `protobuf:"bytes,13,opt,name=device,json=d,proto3" json:"device,omitempty"`
	CloseCause  string            `protobuf:"bytes,14,opt,name=close_cause,json=o,proto3" json:"close_cause,omitempty"`
	Timestamp   int64             `protobuf:"varint,15,opt,name=timestamp,json=v,proto3" json:"timestamp,omitempty"`
}

func (x *ActivityMessage) Reset() {
	*x = ActivityMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_mq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityMessage) ProtoMessage() {}

func (x *ActivityMessage) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_mq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityMessage.ProtoReflect.Descriptor instead.
func (*ActivityMessage) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_mq_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ActivityMessage) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ActivityMessage) GetEvent() ActivityEventType {
	if x != nil {
		return x.Event
	}
	return ActivityEventType_EventUnknown
}

func (x *ActivityMessage) GetSvc() string {
	if x != nil {
		return x.Svc
	}
	return ""
}

func (x *ActivityMessage) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *ActivityMessage) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ActivityMessage) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *ActivityMessage) GetAddition() string {
	if x != nil {
		return x.Addition
	}
	return ""
}

func (x *ActivityMessage) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *ActivityMessage) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ActivityMessage) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ActivityMessage) GetLastGroupId() string {
	if x != nil {
		return x.LastGroupId
	}
	return ""
}

func (x *ActivityMessage) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *ActivityMessage) GetCloseCause() string {
	if x != nil {
		return x.CloseCause
	}
	return ""
}

func (x *ActivityMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type MsgMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,json=i,proto3" json:"id,omitempty"`
	Type      PayloadType            `protobuf:"varint,2,opt,name=type,json=t,proto3,enum=svc.infra.link.PayloadType" json:"type,omitempty"`
	Priority  PriorityType           `protobuf:"varint,3,opt,name=priority,json=p,proto3,enum=svc.infra.link.PriorityType" json:"priority,omitempty"`
	Time      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=time,json=s,proto3" json:"time,omitempty"`
	Scope     ScopeType              `protobuf:"varint,5,opt,name=scope,json=e,proto3,enum=svc.infra.link.ScopeType" json:"scope,omitempty"`
	From      string                 `protobuf:"bytes,6,opt,name=from,json=f,proto3" json:"from,omitempty"`
	ToSession string                 `protobuf:"bytes,7,opt,name=to_session,json=ts,proto3" json:"to_session,omitempty"`
	ToAccount string                 `protobuf:"bytes,8,opt,name=to_account,json=ta,proto3" json:"to_account,omitempty"`
	ToGroup   string                 `protobuf:"bytes,9,opt,name=to_group,json=tg,proto3" json:"to_group,omitempty"`
	ToDevice  string                 `protobuf:"bytes,10,opt,name=to_device,json=td,proto3" json:"to_device,omitempty"`
	Payload   string                 `protobuf:"bytes,11,opt,name=payload,json=d,proto3" json:"payload,omitempty"`
}

func (x *MsgMessage) Reset() {
	*x = MsgMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_mq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgMessage) ProtoMessage() {}

func (x *MsgMessage) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_mq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgMessage.ProtoReflect.Descriptor instead.
func (*MsgMessage) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_mq_proto_rawDescGZIP(), []int{1}
}

func (x *MsgMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MsgMessage) GetType() PayloadType {
	if x != nil {
		return x.Type
	}
	return PayloadType_StreamerDm
}

func (x *MsgMessage) GetPriority() PriorityType {
	if x != nil {
		return x.Priority
	}
	return PriorityType_Low
}

func (x *MsgMessage) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *MsgMessage) GetScope() ScopeType {
	if x != nil {
		return x.Scope
	}
	return ScopeType_ScopeGlobal
}

func (x *MsgMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *MsgMessage) GetToSession() string {
	if x != nil {
		return x.ToSession
	}
	return ""
}

func (x *MsgMessage) GetToAccount() string {
	if x != nil {
		return x.ToAccount
	}
	return ""
}

func (x *MsgMessage) GetToGroup() string {
	if x != nil {
		return x.ToGroup
	}
	return ""
}

func (x *MsgMessage) GetToDevice() string {
	if x != nil {
		return x.ToDevice
	}
	return ""
}

func (x *MsgMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type MsgInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,json=i,proto3" json:"id,omitempty"`
	Time      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,json=s,proto3" json:"time,omitempty"`
	Scope     ScopeType              `protobuf:"varint,3,opt,name=scope,json=e,proto3,enum=svc.infra.link.ScopeType" json:"scope,omitempty"`
	From      string                 `protobuf:"bytes,4,opt,name=from,json=f,proto3" json:"from,omitempty"`
	ToSession string                 `protobuf:"bytes,5,opt,name=to_session,json=ts,proto3" json:"to_session,omitempty"`
	ToAccount string                 `protobuf:"bytes,6,opt,name=to_account,json=ta,proto3" json:"to_account,omitempty"`
	ToGroup   string                 `protobuf:"bytes,7,opt,name=to_group,json=tg,proto3" json:"to_group,omitempty"`
	ToDevice  string                 `protobuf:"bytes,8,opt,name=to_device,json=td,proto3" json:"to_device,omitempty"`
}

func (x *MsgInstruction) Reset() {
	*x = MsgInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_mq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgInstruction) ProtoMessage() {}

func (x *MsgInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_mq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgInstruction.ProtoReflect.Descriptor instead.
func (*MsgInstruction) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_mq_proto_rawDescGZIP(), []int{2}
}

func (x *MsgInstruction) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MsgInstruction) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *MsgInstruction) GetScope() ScopeType {
	if x != nil {
		return x.Scope
	}
	return ScopeType_ScopeGlobal
}

func (x *MsgInstruction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *MsgInstruction) GetToSession() string {
	if x != nil {
		return x.ToSession
	}
	return ""
}

func (x *MsgInstruction) GetToAccount() string {
	if x != nil {
		return x.ToAccount
	}
	return ""
}

func (x *MsgInstruction) GetToGroup() string {
	if x != nil {
		return x.ToGroup
	}
	return ""
}

func (x *MsgInstruction) GetToDevice() string {
	if x != nil {
		return x.ToDevice
	}
	return ""
}

// 用户进出房间topic: topic.user.room.enter_quit
type UserEnterQuitRoomTopic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 房间ID
	RoomId     string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MerchantId string `protobuf:"bytes,3,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	// 当前直播间的在人数数量
	OnlineCount int64 `protobuf:"varint,4,opt,name=online_count,json=onlineCount,proto3" json:"online_count,omitempty"`
	// 主播ID
	StreamerId string              `protobuf:"bytes,5,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`
	Action     UserEnterQuitAction `protobuf:"varint,255,opt,name=action,proto3,enum=svc.infra.link.UserEnterQuitAction" json:"action,omitempty"`
}

func (x *UserEnterQuitRoomTopic) Reset() {
	*x = UserEnterQuitRoomTopic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_mq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEnterQuitRoomTopic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEnterQuitRoomTopic) ProtoMessage() {}

func (x *UserEnterQuitRoomTopic) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_mq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEnterQuitRoomTopic.ProtoReflect.Descriptor instead.
func (*UserEnterQuitRoomTopic) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_mq_proto_rawDescGZIP(), []int{3}
}

func (x *UserEnterQuitRoomTopic) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *UserEnterQuitRoomTopic) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserEnterQuitRoomTopic) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *UserEnterQuitRoomTopic) GetOnlineCount() int64 {
	if x != nil {
		return x.OnlineCount
	}
	return 0
}

func (x *UserEnterQuitRoomTopic) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *UserEnterQuitRoomTopic) GetAction() UserEnterQuitAction {
	if x != nil {
		return x.Action
	}
	return UserEnterQuitAction_Enter
}

var File_svc_infra_link_mq_proto protoreflect.FileDescriptor

var file_svc_infra_link_mq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2f, 0x6d, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x76, 0x63, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x76, 0x63, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x02, 0x0a, 0x0f, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x69, 0x12, 0x10, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x74, 0x12, 0x33, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x01, 0x65, 0x12, 0x0e, 0x0a, 0x03, 0x73, 0x76, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x63, 0x12, 0x15, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x6e, 0x12, 0x15, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73,
	0x12, 0x16, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x72, 0x12, 0x13, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x12, 0x16, 0x0a,
	0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x6d, 0x12, 0x15, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x13, 0x0a, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x67, 0x12, 0x18, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x6c, 0x12, 0x11, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x64, 0x12, 0x16,
	0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x6f, 0x12, 0x14, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x76, 0x22, 0xd8, 0x02, 0x0a,
	0x0a, 0x4d, 0x73, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x69, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x01, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x01, 0x70, 0x12, 0x2b, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x01, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x01, 0x65, 0x12, 0x0f, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x66, 0x12, 0x16, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x73, 0x12, 0x16,
	0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x08, 0x74, 0x6f, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x67, 0x12, 0x15, 0x0a, 0x09,
	0x74, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x64, 0x12, 0x12, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x64, 0x22, 0xe7, 0x01, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x49,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x69, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x01, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x01, 0x65, 0x12, 0x0f, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x66, 0x12, 0x16, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x0a,
	0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x08, 0x74, 0x6f, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x67, 0x12, 0x15, 0x0a, 0x09, 0x74, 0x6f,
	0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x64, 0x22, 0xed, 0x01, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x51,
	0x75, 0x69, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xff, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x51,
	0x75, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0x79, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x05, 0x2a, 0x96, 0x01, 0x0a,
	0x09, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x10, 0x03, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x04,
	0x12, 0x0d, 0x0a, 0x09, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4b, 0x69, 0x63, 0x6b, 0x10, 0x7d, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x6c,
	0x65, 0x10, 0x7e, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x7f, 0x2a, 0x2a, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x51, 0x75, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x74, 0x10,
	0x01, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x3b, 0x6c, 0x69, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_svc_infra_link_mq_proto_rawDescOnce sync.Once
	file_svc_infra_link_mq_proto_rawDescData = file_svc_infra_link_mq_proto_rawDesc
)

func file_svc_infra_link_mq_proto_rawDescGZIP() []byte {
	file_svc_infra_link_mq_proto_rawDescOnce.Do(func() {
		file_svc_infra_link_mq_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_infra_link_mq_proto_rawDescData)
	})
	return file_svc_infra_link_mq_proto_rawDescData
}

var file_svc_infra_link_mq_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_svc_infra_link_mq_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_svc_infra_link_mq_proto_goTypes = []any{
	(ActivityEventType)(0),         // 0: svc.infra.link.ActivityEventType
	(ScopeType)(0),                 // 1: svc.infra.link.ScopeType
	(UserEnterQuitAction)(0),       // 2: svc.infra.link.UserEnterQuitAction
	(*ActivityMessage)(nil),        // 3: svc.infra.link.ActivityMessage
	(*MsgMessage)(nil),             // 4: svc.infra.link.MsgMessage
	(*MsgInstruction)(nil),         // 5: svc.infra.link.MsgInstruction
	(*UserEnterQuitRoomTopic)(nil), // 6: svc.infra.link.UserEnterQuitRoomTopic
	(PayloadType)(0),               // 7: svc.infra.link.PayloadType
	(PriorityType)(0),              // 8: svc.infra.link.PriorityType
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_svc_infra_link_mq_proto_depIdxs = []int32{
	0, // 0: svc.infra.link.ActivityMessage.event:type_name -> svc.infra.link.ActivityEventType
	7, // 1: svc.infra.link.MsgMessage.type:type_name -> svc.infra.link.PayloadType
	8, // 2: svc.infra.link.MsgMessage.priority:type_name -> svc.infra.link.PriorityType
	9, // 3: svc.infra.link.MsgMessage.time:type_name -> google.protobuf.Timestamp
	1, // 4: svc.infra.link.MsgMessage.scope:type_name -> svc.infra.link.ScopeType
	9, // 5: svc.infra.link.MsgInstruction.time:type_name -> google.protobuf.Timestamp
	1, // 6: svc.infra.link.MsgInstruction.scope:type_name -> svc.infra.link.ScopeType
	2, // 7: svc.infra.link.UserEnterQuitRoomTopic.action:type_name -> svc.infra.link.UserEnterQuitAction
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_svc_infra_link_mq_proto_init() }
func file_svc_infra_link_mq_proto_init() {
	if File_svc_infra_link_mq_proto != nil {
		return
	}
	file_svc_infra_link_payload_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_svc_infra_link_mq_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityMessage); i {
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
		file_svc_infra_link_mq_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MsgMessage); i {
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
		file_svc_infra_link_mq_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MsgInstruction); i {
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
		file_svc_infra_link_mq_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UserEnterQuitRoomTopic); i {
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
			RawDescriptor: file_svc_infra_link_mq_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_svc_infra_link_mq_proto_goTypes,
		DependencyIndexes: file_svc_infra_link_mq_proto_depIdxs,
		EnumInfos:         file_svc_infra_link_mq_proto_enumTypes,
		MessageInfos:      file_svc_infra_link_mq_proto_msgTypes,
	}.Build()
	File_svc_infra_link_mq_proto = out.File
	file_svc_infra_link_mq_proto_rawDesc = nil
	file_svc_infra_link_mq_proto_goTypes = nil
	file_svc_infra_link_mq_proto_depIdxs = nil
}
