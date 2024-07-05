// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: svc.biz.rank/streamer_viewer_rank.proto

package rank

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

type StreamerViewerRankPeriod int32

const (
	StreamerViewerRankPeriod_StreamerViewerRankPeriodUnknown StreamerViewerRankPeriod = 0 // 未知
	StreamerViewerRankPeriod_StreamerViewerRankPeriodDay     StreamerViewerRankPeriod = 1 // 日榜
	StreamerViewerRankPeriod_StreamerViewerRankPeriodWeek    StreamerViewerRankPeriod = 2 // 周榜
	StreamerViewerRankPeriod_StreamerViewerRankPeriodMonth   StreamerViewerRankPeriod = 3 // 月榜
	StreamerViewerRankPeriod_StreamerViewerRankPeriodAll     StreamerViewerRankPeriod = 4 // 总榜
)

// Enum value maps for StreamerViewerRankPeriod.
var (
	StreamerViewerRankPeriod_name = map[int32]string{
		0: "StreamerViewerRankPeriodUnknown",
		1: "StreamerViewerRankPeriodDay",
		2: "StreamerViewerRankPeriodWeek",
		3: "StreamerViewerRankPeriodMonth",
		4: "StreamerViewerRankPeriodAll",
	}
	StreamerViewerRankPeriod_value = map[string]int32{
		"StreamerViewerRankPeriodUnknown": 0,
		"StreamerViewerRankPeriodDay":     1,
		"StreamerViewerRankPeriodWeek":    2,
		"StreamerViewerRankPeriodMonth":   3,
		"StreamerViewerRankPeriodAll":     4,
	}
)

func (x StreamerViewerRankPeriod) Enum() *StreamerViewerRankPeriod {
	p := new(StreamerViewerRankPeriod)
	*p = x
	return p
}

func (x StreamerViewerRankPeriod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamerViewerRankPeriod) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_biz_rank_streamer_viewer_rank_proto_enumTypes[0].Descriptor()
}

func (StreamerViewerRankPeriod) Type() protoreflect.EnumType {
	return &file_svc_biz_rank_streamer_viewer_rank_proto_enumTypes[0]
}

func (x StreamerViewerRankPeriod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamerViewerRankPeriod.Descriptor instead.
func (StreamerViewerRankPeriod) EnumDescriptor() ([]byte, []int) {
	return file_svc_biz_rank_streamer_viewer_rank_proto_rawDescGZIP(), []int{0}
}

// RankMemberInfo 排行榜成员信息
type RankMemberInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId string `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Score    int64  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *RankMemberInfo) Reset() {
	*x = RankMemberInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankMemberInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankMemberInfo) ProtoMessage() {}

func (x *RankMemberInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankMemberInfo.ProtoReflect.Descriptor instead.
func (*RankMemberInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_rank_streamer_viewer_rank_proto_rawDescGZIP(), []int{0}
}

func (x *RankMemberInfo) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *RankMemberInfo) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type GetStreamerViewerLiveMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit      int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	StreamerId string `protobuf:"bytes,3,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"` // 主播id
	LiveId     string `protobuf:"bytes,4,opt,name=live_id,json=liveId,proto3" json:"live_id,omitempty"`             // 直播id
}

func (x *GetStreamerViewerLiveMemberReq) Reset() {
	*x = GetStreamerViewerLiveMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamerViewerLiveMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamerViewerLiveMemberReq) ProtoMessage() {}

func (x *GetStreamerViewerLiveMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamerViewerLiveMemberReq.ProtoReflect.Descriptor instead.
func (*GetStreamerViewerLiveMemberReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_rank_streamer_viewer_rank_proto_rawDescGZIP(), []int{1}
}

func (x *GetStreamerViewerLiveMemberReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetStreamerViewerLiveMemberReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetStreamerViewerLiveMemberReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *GetStreamerViewerLiveMemberReq) GetLiveId() string {
	if x != nil {
		return x.LiveId
	}
	return ""
}

type GetStreamerViewerGiftMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32                    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit      int32                    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	RankPeriod StreamerViewerRankPeriod `protobuf:"varint,3,opt,name=rank_period,json=rankPeriod,proto3,enum=svc.biz.rank.StreamerViewerRankPeriod" json:"rank_period,omitempty"` // 排行榜周期
	StreamerId string                   `protobuf:"bytes,4,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`                                             // 主播id
	PeriodTime string                   `protobuf:"bytes,5,opt,name=period_time,json=periodTime,proto3" json:"period_time,omitempty"`                                             // 传空则按照当前时间计算（日2006-01-02，月2006-01，周2006-1、使用"-"拼接time.ISOWeek返回值）
}

func (x *GetStreamerViewerGiftMemberReq) Reset() {
	*x = GetStreamerViewerGiftMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamerViewerGiftMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamerViewerGiftMemberReq) ProtoMessage() {}

func (x *GetStreamerViewerGiftMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamerViewerGiftMemberReq.ProtoReflect.Descriptor instead.
func (*GetStreamerViewerGiftMemberReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_rank_streamer_viewer_rank_proto_rawDescGZIP(), []int{2}
}

func (x *GetStreamerViewerGiftMemberReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetStreamerViewerGiftMemberReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetStreamerViewerGiftMemberReq) GetRankPeriod() StreamerViewerRankPeriod {
	if x != nil {
		return x.RankPeriod
	}
	return StreamerViewerRankPeriod_StreamerViewerRankPeriodUnknown
}

func (x *GetStreamerViewerGiftMemberReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *GetStreamerViewerGiftMemberReq) GetPeriodTime() string {
	if x != nil {
		return x.PeriodTime
	}
	return ""
}

type GetStreamerViewerGlamourMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32                    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit      int32                    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	RankPeriod StreamerViewerRankPeriod `protobuf:"varint,3,opt,name=rank_period,json=rankPeriod,proto3,enum=svc.biz.rank.StreamerViewerRankPeriod" json:"rank_period,omitempty"` // 排行榜周期
	StreamerId string                   `protobuf:"bytes,4,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`                                             // 主播id
	PeriodTime string                   `protobuf:"bytes,5,opt,name=period_time,json=periodTime,proto3" json:"period_time,omitempty"`                                             // 传空则按照当前时间计算（日2006-01-02，月2006-01，周2006-1、使用"-"拼接time.ISOWeek返回值）
}

func (x *GetStreamerViewerGlamourMemberReq) Reset() {
	*x = GetStreamerViewerGlamourMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamerViewerGlamourMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamerViewerGlamourMemberReq) ProtoMessage() {}

func (x *GetStreamerViewerGlamourMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamerViewerGlamourMemberReq.ProtoReflect.Descriptor instead.
func (*GetStreamerViewerGlamourMemberReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_rank_streamer_viewer_rank_proto_rawDescGZIP(), []int{3}
}

func (x *GetStreamerViewerGlamourMemberReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetStreamerViewerGlamourMemberReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetStreamerViewerGlamourMemberReq) GetRankPeriod() StreamerViewerRankPeriod {
	if x != nil {
		return x.RankPeriod
	}
	return StreamerViewerRankPeriod_StreamerViewerRankPeriodUnknown
}

func (x *GetStreamerViewerGlamourMemberReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *GetStreamerViewerGlamourMemberReq) GetPeriodTime() string {
	if x != nil {
		return x.PeriodTime
	}
	return ""
}

type GetStreamerViewerRankResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*RankMemberInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetStreamerViewerRankResp) Reset() {
	*x = GetStreamerViewerRankResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamerViewerRankResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamerViewerRankResp) ProtoMessage() {}

func (x *GetStreamerViewerRankResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamerViewerRankResp.ProtoReflect.Descriptor instead.
func (*GetStreamerViewerRankResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_rank_streamer_viewer_rank_proto_rawDescGZIP(), []int{4}
}

func (x *GetStreamerViewerRankResp) GetItems() []*RankMemberInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_svc_biz_rank_streamer_viewer_rank_proto protoreflect.FileDescriptor

var file_svc_biz_rank_streamer_viewer_rank_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x5f, 0x72,
	0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x43, 0x0a, 0x0e, 0x52, 0x61, 0x6e, 0x6b, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x84, 0x01, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77,
	0x65, 0x72, 0x4c, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69,
	0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x76,
	0x65, 0x49, 0x64, 0x22, 0xd5, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x47, 0x69, 0x66, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x47, 0x0a, 0x0b, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e,
	0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65,
	0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x0a, 0x72,
	0x61, 0x6e, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x21,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65,
	0x72, 0x47, 0x6c, 0x61, 0x6d, 0x6f, 0x75, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x72,
	0x61, 0x6e, 0x6b, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x52, 0x61,
	0x6e, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0xc6, 0x01, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x44, 0x61, 0x79, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x57, 0x65, 0x65, 0x6b, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x52, 0x61,
	0x6e, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x10, 0x03, 0x12,
	0x1f, 0x0a, 0x1b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65,
	0x72, 0x52, 0x61, 0x6e, 0x6b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x41, 0x6c, 0x6c, 0x10, 0x04,
	0x32, 0x80, 0x04, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65,
	0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x76, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x76, 0x65,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65,
	0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56,
	0x69, 0x65, 0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x7c, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69,
	0x65, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56,
	0x69, 0x65, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65,
	0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x76, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77,
	0x65, 0x72, 0x47, 0x69, 0x66, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x47, 0x69, 0x66,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x47, 0x6c, 0x61, 0x6d, 0x6f, 0x75,
	0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2f, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x47, 0x6c, 0x61, 0x6d, 0x6f, 0x75, 0x72, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x3b, 0x72, 0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_svc_biz_rank_streamer_viewer_rank_proto_rawDescOnce sync.Once
	file_svc_biz_rank_streamer_viewer_rank_proto_rawDescData = file_svc_biz_rank_streamer_viewer_rank_proto_rawDesc
)

func file_svc_biz_rank_streamer_viewer_rank_proto_rawDescGZIP() []byte {
	file_svc_biz_rank_streamer_viewer_rank_proto_rawDescOnce.Do(func() {
		file_svc_biz_rank_streamer_viewer_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_rank_streamer_viewer_rank_proto_rawDescData)
	})
	return file_svc_biz_rank_streamer_viewer_rank_proto_rawDescData
}

var file_svc_biz_rank_streamer_viewer_rank_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_svc_biz_rank_streamer_viewer_rank_proto_goTypes = []any{
	(StreamerViewerRankPeriod)(0),             // 0: svc.biz.rank.StreamerViewerRankPeriod
	(*RankMemberInfo)(nil),                    // 1: svc.biz.rank.RankMemberInfo
	(*GetStreamerViewerLiveMemberReq)(nil),    // 2: svc.biz.rank.GetStreamerViewerLiveMemberReq
	(*GetStreamerViewerGiftMemberReq)(nil),    // 3: svc.biz.rank.GetStreamerViewerGiftMemberReq
	(*GetStreamerViewerGlamourMemberReq)(nil), // 4: svc.biz.rank.GetStreamerViewerGlamourMemberReq
	(*GetStreamerViewerRankResp)(nil),         // 5: svc.biz.rank.GetStreamerViewerRankResp
}
var file_svc_biz_rank_streamer_viewer_rank_proto_depIdxs = []int32{
	0, // 0: svc.biz.rank.GetStreamerViewerGiftMemberReq.rank_period:type_name -> svc.biz.rank.StreamerViewerRankPeriod
	0, // 1: svc.biz.rank.GetStreamerViewerGlamourMemberReq.rank_period:type_name -> svc.biz.rank.StreamerViewerRankPeriod
	1, // 2: svc.biz.rank.GetStreamerViewerRankResp.items:type_name -> svc.biz.rank.RankMemberInfo
	2, // 3: svc.biz.rank.StreamerViewerRank.GetStreamerViewerLiveMember:input_type -> svc.biz.rank.GetStreamerViewerLiveMemberReq
	2, // 4: svc.biz.rank.StreamerViewerRank.GetStreamerViewerLiveOnlineMember:input_type -> svc.biz.rank.GetStreamerViewerLiveMemberReq
	3, // 5: svc.biz.rank.StreamerViewerRank.GetStreamerViewerGiftMember:input_type -> svc.biz.rank.GetStreamerViewerGiftMemberReq
	4, // 6: svc.biz.rank.StreamerViewerRank.GetStreamerViewerGlamourMember:input_type -> svc.biz.rank.GetStreamerViewerGlamourMemberReq
	5, // 7: svc.biz.rank.StreamerViewerRank.GetStreamerViewerLiveMember:output_type -> svc.biz.rank.GetStreamerViewerRankResp
	5, // 8: svc.biz.rank.StreamerViewerRank.GetStreamerViewerLiveOnlineMember:output_type -> svc.biz.rank.GetStreamerViewerRankResp
	5, // 9: svc.biz.rank.StreamerViewerRank.GetStreamerViewerGiftMember:output_type -> svc.biz.rank.GetStreamerViewerRankResp
	5, // 10: svc.biz.rank.StreamerViewerRank.GetStreamerViewerGlamourMember:output_type -> svc.biz.rank.GetStreamerViewerRankResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_svc_biz_rank_streamer_viewer_rank_proto_init() }
func file_svc_biz_rank_streamer_viewer_rank_proto_init() {
	if File_svc_biz_rank_streamer_viewer_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RankMemberInfo); i {
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
		file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetStreamerViewerLiveMemberReq); i {
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
		file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetStreamerViewerGiftMemberReq); i {
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
		file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetStreamerViewerGlamourMemberReq); i {
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
		file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetStreamerViewerRankResp); i {
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
			RawDescriptor: file_svc_biz_rank_streamer_viewer_rank_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_rank_streamer_viewer_rank_proto_goTypes,
		DependencyIndexes: file_svc_biz_rank_streamer_viewer_rank_proto_depIdxs,
		EnumInfos:         file_svc_biz_rank_streamer_viewer_rank_proto_enumTypes,
		MessageInfos:      file_svc_biz_rank_streamer_viewer_rank_proto_msgTypes,
	}.Build()
	File_svc_biz_rank_streamer_viewer_rank_proto = out.File
	file_svc_biz_rank_streamer_viewer_rank_proto_rawDesc = nil
	file_svc_biz_rank_streamer_viewer_rank_proto_goTypes = nil
	file_svc_biz_rank_streamer_viewer_rank_proto_depIdxs = nil
}
