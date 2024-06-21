// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: svc.biz.vip/fanbase_member.proto

package vip

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type FanbaseLevel int32

const (
	FanbaseLevel_FanbaseLevelUnknown FanbaseLevel = 0 // 未知
	FanbaseLevel_FanbaseLevelPrimary FanbaseLevel = 1 // 初级
	FanbaseLevel_FanbaseLevelSuper   FanbaseLevel = 2 // 超级
	FanbaseLevel_FanbaseLevelExtreme FanbaseLevel = 3 // 至尊
)

// Enum value maps for FanbaseLevel.
var (
	FanbaseLevel_name = map[int32]string{
		0: "FanbaseLevelUnknown",
		1: "FanbaseLevelPrimary",
		2: "FanbaseLevelSuper",
		3: "FanbaseLevelExtreme",
	}
	FanbaseLevel_value = map[string]int32{
		"FanbaseLevelUnknown": 0,
		"FanbaseLevelPrimary": 1,
		"FanbaseLevelSuper":   2,
		"FanbaseLevelExtreme": 3,
	}
)

func (x FanbaseLevel) Enum() *FanbaseLevel {
	p := new(FanbaseLevel)
	*p = x
	return p
}

func (x FanbaseLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FanbaseLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_biz_vip_fanbase_member_proto_enumTypes[0].Descriptor()
}

func (FanbaseLevel) Type() protoreflect.EnumType {
	return &file_svc_biz_vip_fanbase_member_proto_enumTypes[0]
}

func (x FanbaseLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FanbaseLevel.Descriptor instead.
func (FanbaseLevel) EnumDescriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_member_proto_rawDescGZIP(), []int{0}
}

type JoinFanbaseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamerId string       `protobuf:"bytes,1,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`
	MemberId   string       `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Level      FanbaseLevel `protobuf:"varint,3,opt,name=level,proto3,enum=svc.biz.vip.FanbaseLevel" json:"level,omitempty"`
}

func (x *JoinFanbaseReq) Reset() {
	*x = JoinFanbaseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinFanbaseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinFanbaseReq) ProtoMessage() {}

func (x *JoinFanbaseReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinFanbaseReq.ProtoReflect.Descriptor instead.
func (*JoinFanbaseReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_member_proto_rawDescGZIP(), []int{0}
}

func (x *JoinFanbaseReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *JoinFanbaseReq) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *JoinFanbaseReq) GetLevel() FanbaseLevel {
	if x != nil {
		return x.Level
	}
	return FanbaseLevel_FanbaseLevelUnknown
}

type LeaveFanbaseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamerId string `protobuf:"bytes,1,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`
	MemberId   string `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *LeaveFanbaseReq) Reset() {
	*x = LeaveFanbaseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveFanbaseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveFanbaseReq) ProtoMessage() {}

func (x *LeaveFanbaseReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveFanbaseReq.ProtoReflect.Descriptor instead.
func (*LeaveFanbaseReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_member_proto_rawDescGZIP(), []int{1}
}

func (x *LeaveFanbaseReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *LeaveFanbaseReq) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type GetFanbaseMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamerId string `protobuf:"bytes,1,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`
	MemberId   string `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *GetFanbaseMemberReq) Reset() {
	*x = GetFanbaseMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFanbaseMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFanbaseMemberReq) ProtoMessage() {}

func (x *GetFanbaseMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFanbaseMemberReq.ProtoReflect.Descriptor instead.
func (*GetFanbaseMemberReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_member_proto_rawDescGZIP(), []int{2}
}

func (x *GetFanbaseMemberReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *GetFanbaseMemberReq) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type GetFanbaseMemberResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FanbaseMember *FanbaseMemberInfo `protobuf:"bytes,1,opt,name=fanbase_member,json=fanbaseMember,proto3" json:"fanbase_member,omitempty"`
}

func (x *GetFanbaseMemberResp) Reset() {
	*x = GetFanbaseMemberResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFanbaseMemberResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFanbaseMemberResp) ProtoMessage() {}

func (x *GetFanbaseMemberResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFanbaseMemberResp.ProtoReflect.Descriptor instead.
func (*GetFanbaseMemberResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_member_proto_rawDescGZIP(), []int{3}
}

func (x *GetFanbaseMemberResp) GetFanbaseMember() *FanbaseMemberInfo {
	if x != nil {
		return x.FanbaseMember
	}
	return nil
}

type FanbaseMemberInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FanbaseId  string                 `protobuf:"bytes,1,opt,name=fanbase_id,json=fanbaseId,proto3" json:"fanbase_id,omitempty"`
	StreamerId string                 `protobuf:"bytes,2,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`
	MemberId   string                 `protobuf:"bytes,3,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Level      FanbaseLevel           `protobuf:"varint,4,opt,name=level,proto3,enum=svc.biz.vip.FanbaseLevel" json:"level,omitempty"`
	JoinTime   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=join_time,json=joinTime,proto3" json:"join_time,omitempty"` // 加入时间
	Score      int32                  `protobuf:"varint,6,opt,name=score,proto3" json:"score,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 创建时间
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,102,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // 更新时间
}

func (x *FanbaseMemberInfo) Reset() {
	*x = FanbaseMemberInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanbaseMemberInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanbaseMemberInfo) ProtoMessage() {}

func (x *FanbaseMemberInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanbaseMemberInfo.ProtoReflect.Descriptor instead.
func (*FanbaseMemberInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_member_proto_rawDescGZIP(), []int{4}
}

func (x *FanbaseMemberInfo) GetFanbaseId() string {
	if x != nil {
		return x.FanbaseId
	}
	return ""
}

func (x *FanbaseMemberInfo) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *FanbaseMemberInfo) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *FanbaseMemberInfo) GetLevel() FanbaseLevel {
	if x != nil {
		return x.Level
	}
	return FanbaseLevel_FanbaseLevelUnknown
}

func (x *FanbaseMemberInfo) GetJoinTime() *timestamppb.Timestamp {
	if x != nil {
		return x.JoinTime
	}
	return nil
}

func (x *FanbaseMemberInfo) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *FanbaseMemberInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FanbaseMemberInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetFanbaseMemberByStreamerIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32        `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                                 // 页数
	Limit      int32        `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`                               // 条数
	StreamerId string       `protobuf:"bytes,3,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`    // 主播id
	Level      FanbaseLevel `protobuf:"varint,4,opt,name=level,proto3,enum=svc.biz.vip.FanbaseLevel" json:"level,omitempty"` // 等级
}

func (x *GetFanbaseMemberByStreamerIDReq) Reset() {
	*x = GetFanbaseMemberByStreamerIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFanbaseMemberByStreamerIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFanbaseMemberByStreamerIDReq) ProtoMessage() {}

func (x *GetFanbaseMemberByStreamerIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFanbaseMemberByStreamerIDReq.ProtoReflect.Descriptor instead.
func (*GetFanbaseMemberByStreamerIDReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_member_proto_rawDescGZIP(), []int{5}
}

func (x *GetFanbaseMemberByStreamerIDReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFanbaseMemberByStreamerIDReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetFanbaseMemberByStreamerIDReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *GetFanbaseMemberByStreamerIDReq) GetLevel() FanbaseLevel {
	if x != nil {
		return x.Level
	}
	return FanbaseLevel_FanbaseLevelUnknown
}

type GetFanbaseMembertByMemberIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                        // 页数
	Limit    int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`                      // 条数
	MemberId string `protobuf:"bytes,3,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"` // 成员id
}

func (x *GetFanbaseMembertByMemberIDResp) Reset() {
	*x = GetFanbaseMembertByMemberIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFanbaseMembertByMemberIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFanbaseMembertByMemberIDResp) ProtoMessage() {}

func (x *GetFanbaseMembertByMemberIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFanbaseMembertByMemberIDResp.ProtoReflect.Descriptor instead.
func (*GetFanbaseMembertByMemberIDResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_member_proto_rawDescGZIP(), []int{6}
}

func (x *GetFanbaseMembertByMemberIDResp) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFanbaseMembertByMemberIDResp) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetFanbaseMembertByMemberIDResp) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type GetOnlineFanbaseMemberByStreamerIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                              // 页数
	Limit      int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`                            // 条数
	StreamerId string `protobuf:"bytes,3,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"` // 主播id
}

func (x *GetOnlineFanbaseMemberByStreamerIDReq) Reset() {
	*x = GetOnlineFanbaseMemberByStreamerIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOnlineFanbaseMemberByStreamerIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOnlineFanbaseMemberByStreamerIDReq) ProtoMessage() {}

func (x *GetOnlineFanbaseMemberByStreamerIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOnlineFanbaseMemberByStreamerIDReq.ProtoReflect.Descriptor instead.
func (*GetOnlineFanbaseMemberByStreamerIDReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_member_proto_rawDescGZIP(), []int{7}
}

func (x *GetOnlineFanbaseMemberByStreamerIDReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetOnlineFanbaseMemberByStreamerIDReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetOnlineFanbaseMemberByStreamerIDReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

type GetListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*FanbaseMemberInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetListResp) Reset() {
	*x = GetListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResp) ProtoMessage() {}

func (x *GetListResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_member_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResp.ProtoReflect.Descriptor instead.
func (*GetListResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_member_proto_rawDescGZIP(), []int{8}
}

func (x *GetListResp) GetItems() []*FanbaseMemberInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_svc_biz_vip_fanbase_member_proto protoreflect.FileDescriptor

var file_svc_biz_vip_fanbase_member_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2f, 0x66, 0x61,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a,
	0x0e, 0x4a, 0x6f, 0x69, 0x6e, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x46, 0x61, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x4f,
	0x0a, 0x0f, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x0e,
	0x66, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76,
	0x69, 0x70, 0x2e, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x66, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0xe6, 0x02, 0x0a, 0x11, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e,
	0x76, 0x69, 0x70, 0x2e, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x37, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9d, 0x01, 0x0a,
	0x1f, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x68, 0x0a, 0x1f,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x74, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x25, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a,
	0x70, 0x0a, 0x0c, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x17, 0x0a, 0x13, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x61, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x53, 0x75, 0x70, 0x65, 0x72, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x61, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x65, 0x6d, 0x65, 0x10,
	0x03, 0x32, 0xc1, 0x04, 0x0a, 0x0d, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x46, 0x61, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0c, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x46, 0x61, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e,
	0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2c, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x32, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x1a, 0x18, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x74, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2c, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x74, 0x42, 0x79, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x1a, 0x18, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x3b, 0x76, 0x69, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_svc_biz_vip_fanbase_member_proto_rawDescOnce sync.Once
	file_svc_biz_vip_fanbase_member_proto_rawDescData = file_svc_biz_vip_fanbase_member_proto_rawDesc
)

func file_svc_biz_vip_fanbase_member_proto_rawDescGZIP() []byte {
	file_svc_biz_vip_fanbase_member_proto_rawDescOnce.Do(func() {
		file_svc_biz_vip_fanbase_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_vip_fanbase_member_proto_rawDescData)
	})
	return file_svc_biz_vip_fanbase_member_proto_rawDescData
}

var file_svc_biz_vip_fanbase_member_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_svc_biz_vip_fanbase_member_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_svc_biz_vip_fanbase_member_proto_goTypes = []interface{}{
	(FanbaseLevel)(0),                             // 0: svc.biz.vip.FanbaseLevel
	(*JoinFanbaseReq)(nil),                        // 1: svc.biz.vip.JoinFanbaseReq
	(*LeaveFanbaseReq)(nil),                       // 2: svc.biz.vip.LeaveFanbaseReq
	(*GetFanbaseMemberReq)(nil),                   // 3: svc.biz.vip.GetFanbaseMemberReq
	(*GetFanbaseMemberResp)(nil),                  // 4: svc.biz.vip.GetFanbaseMemberResp
	(*FanbaseMemberInfo)(nil),                     // 5: svc.biz.vip.FanbaseMemberInfo
	(*GetFanbaseMemberByStreamerIDReq)(nil),       // 6: svc.biz.vip.GetFanbaseMemberByStreamerIDReq
	(*GetFanbaseMembertByMemberIDResp)(nil),       // 7: svc.biz.vip.GetFanbaseMembertByMemberIDResp
	(*GetOnlineFanbaseMemberByStreamerIDReq)(nil), // 8: svc.biz.vip.GetOnlineFanbaseMemberByStreamerIDReq
	(*GetListResp)(nil),                           // 9: svc.biz.vip.GetListResp
	(*timestamppb.Timestamp)(nil),                 // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                         // 11: google.protobuf.Empty
}
var file_svc_biz_vip_fanbase_member_proto_depIdxs = []int32{
	0,  // 0: svc.biz.vip.JoinFanbaseReq.level:type_name -> svc.biz.vip.FanbaseLevel
	5,  // 1: svc.biz.vip.GetFanbaseMemberResp.fanbase_member:type_name -> svc.biz.vip.FanbaseMemberInfo
	0,  // 2: svc.biz.vip.FanbaseMemberInfo.level:type_name -> svc.biz.vip.FanbaseLevel
	10, // 3: svc.biz.vip.FanbaseMemberInfo.join_time:type_name -> google.protobuf.Timestamp
	10, // 4: svc.biz.vip.FanbaseMemberInfo.created_at:type_name -> google.protobuf.Timestamp
	10, // 5: svc.biz.vip.FanbaseMemberInfo.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 6: svc.biz.vip.GetFanbaseMemberByStreamerIDReq.level:type_name -> svc.biz.vip.FanbaseLevel
	5,  // 7: svc.biz.vip.GetListResp.items:type_name -> svc.biz.vip.FanbaseMemberInfo
	1,  // 8: svc.biz.vip.FanbaseMember.JoinFanbase:input_type -> svc.biz.vip.JoinFanbaseReq
	2,  // 9: svc.biz.vip.FanbaseMember.LeaveFanbase:input_type -> svc.biz.vip.LeaveFanbaseReq
	3,  // 10: svc.biz.vip.FanbaseMember.GetFanbaseMember:input_type -> svc.biz.vip.GetFanbaseMemberReq
	6,  // 11: svc.biz.vip.FanbaseMember.GetFanbaseMemberByStreamerID:input_type -> svc.biz.vip.GetFanbaseMemberByStreamerIDReq
	8,  // 12: svc.biz.vip.FanbaseMember.GetOnlineFanbaseMemberByStreamerID:input_type -> svc.biz.vip.GetOnlineFanbaseMemberByStreamerIDReq
	7,  // 13: svc.biz.vip.FanbaseMember.GetFanbaseMembertByMemberID:input_type -> svc.biz.vip.GetFanbaseMembertByMemberIDResp
	11, // 14: svc.biz.vip.FanbaseMember.JoinFanbase:output_type -> google.protobuf.Empty
	11, // 15: svc.biz.vip.FanbaseMember.LeaveFanbase:output_type -> google.protobuf.Empty
	4,  // 16: svc.biz.vip.FanbaseMember.GetFanbaseMember:output_type -> svc.biz.vip.GetFanbaseMemberResp
	9,  // 17: svc.biz.vip.FanbaseMember.GetFanbaseMemberByStreamerID:output_type -> svc.biz.vip.GetListResp
	9,  // 18: svc.biz.vip.FanbaseMember.GetOnlineFanbaseMemberByStreamerID:output_type -> svc.biz.vip.GetListResp
	9,  // 19: svc.biz.vip.FanbaseMember.GetFanbaseMembertByMemberID:output_type -> svc.biz.vip.GetListResp
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_svc_biz_vip_fanbase_member_proto_init() }
func file_svc_biz_vip_fanbase_member_proto_init() {
	if File_svc_biz_vip_fanbase_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_biz_vip_fanbase_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinFanbaseReq); i {
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
		file_svc_biz_vip_fanbase_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveFanbaseReq); i {
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
		file_svc_biz_vip_fanbase_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFanbaseMemberReq); i {
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
		file_svc_biz_vip_fanbase_member_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFanbaseMemberResp); i {
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
		file_svc_biz_vip_fanbase_member_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanbaseMemberInfo); i {
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
		file_svc_biz_vip_fanbase_member_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFanbaseMemberByStreamerIDReq); i {
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
		file_svc_biz_vip_fanbase_member_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFanbaseMembertByMemberIDResp); i {
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
		file_svc_biz_vip_fanbase_member_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOnlineFanbaseMemberByStreamerIDReq); i {
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
		file_svc_biz_vip_fanbase_member_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResp); i {
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
			RawDescriptor: file_svc_biz_vip_fanbase_member_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_vip_fanbase_member_proto_goTypes,
		DependencyIndexes: file_svc_biz_vip_fanbase_member_proto_depIdxs,
		EnumInfos:         file_svc_biz_vip_fanbase_member_proto_enumTypes,
		MessageInfos:      file_svc_biz_vip_fanbase_member_proto_msgTypes,
	}.Build()
	File_svc_biz_vip_fanbase_member_proto = out.File
	file_svc_biz_vip_fanbase_member_proto_rawDesc = nil
	file_svc_biz_vip_fanbase_member_proto_goTypes = nil
	file_svc_biz_vip_fanbase_member_proto_depIdxs = nil
}
