// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: svc.biz.room/live.proto

package room

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

type TimeBeginEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beigin *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=beigin,proto3" json:"beigin,omitempty"`
	End    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *TimeBeginEnd) Reset() {
	*x = TimeBeginEnd{}
	mi := &file_svc_biz_room_live_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeBeginEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeBeginEnd) ProtoMessage() {}

func (x *TimeBeginEnd) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_live_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeBeginEnd.ProtoReflect.Descriptor instead.
func (*TimeBeginEnd) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_live_proto_rawDescGZIP(), []int{0}
}

func (x *TimeBeginEnd) GetBeigin() *timestamppb.Timestamp {
	if x != nil {
		return x.Beigin
	}
	return nil
}

func (x *TimeBeginEnd) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type LiveInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LiveId          string                 `protobuf:"bytes,1,opt,name=live_id,json=liveId,proto3" json:"live_id,omitempty"`                                // id
	StreamerId      string                 `protobuf:"bytes,2,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`                    // 主播id
	RoomId          string                 `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`                                // 房间id
	CategoryId      string                 `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`                    // 分类id
	StartAt         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`                             // 开播时间
	EndAt           *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`                                   // 结束时间
	GiftTotalNum    int64                  `protobuf:"varint,21,opt,name=gift_total_num,json=giftTotalNum,proto3" json:"gift_total_num,omitempty"`          // 礼物总数
	GiftTotalPrice  int64                  `protobuf:"varint,22,opt,name=gift_total_price,json=giftTotalPrice,proto3" json:"gift_total_price,omitempty"`    // 礼物代币总数
	GiftTotalUser   int64                  `protobuf:"varint,23,opt,name=gift_total_user,json=giftTotalUser,proto3" json:"gift_total_user,omitempty"`       // 礼物用户总数
	TotalViewer     int64                  `protobuf:"varint,24,opt,name=total_viewer,json=totalViewer,proto3" json:"total_viewer,omitempty"`               // 观看人数
	NewFollow       int64                  `protobuf:"varint,25,opt,name=new_follow,json=newFollow,proto3" json:"new_follow,omitempty"`                     // 新增订阅
	NewFanbase      int64                  `protobuf:"varint,26,opt,name=new_fanbase,json=newFanbase,proto3" json:"new_fanbase,omitempty"`                  // 新增粉丝团
	NewFanbasePrice int64                  `protobuf:"varint,27,opt,name=new_fanbase_price,json=newFanbasePrice,proto3" json:"new_fanbase_price,omitempty"` // 新增粉丝团代币数
	NewNobble       int64                  `protobuf:"varint,28,opt,name=new_nobble,json=newNobble,proto3" json:"new_nobble,omitempty"`                     // 新增贵族
	NewNobblePrice  int64                  `protobuf:"varint,29,opt,name=new_nobble_price,json=newNobblePrice,proto3" json:"new_nobble_price,omitempty"`    // 新增贵族代币数
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,255,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                     // 创建时间
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,256,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`                     // 更新时间
}

func (x *LiveInfo) Reset() {
	*x = LiveInfo{}
	mi := &file_svc_biz_room_live_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LiveInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveInfo) ProtoMessage() {}

func (x *LiveInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_live_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveInfo.ProtoReflect.Descriptor instead.
func (*LiveInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_live_proto_rawDescGZIP(), []int{1}
}

func (x *LiveInfo) GetLiveId() string {
	if x != nil {
		return x.LiveId
	}
	return ""
}

func (x *LiveInfo) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *LiveInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *LiveInfo) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *LiveInfo) GetStartAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *LiveInfo) GetEndAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *LiveInfo) GetGiftTotalNum() int64 {
	if x != nil {
		return x.GiftTotalNum
	}
	return 0
}

func (x *LiveInfo) GetGiftTotalPrice() int64 {
	if x != nil {
		return x.GiftTotalPrice
	}
	return 0
}

func (x *LiveInfo) GetGiftTotalUser() int64 {
	if x != nil {
		return x.GiftTotalUser
	}
	return 0
}

func (x *LiveInfo) GetTotalViewer() int64 {
	if x != nil {
		return x.TotalViewer
	}
	return 0
}

func (x *LiveInfo) GetNewFollow() int64 {
	if x != nil {
		return x.NewFollow
	}
	return 0
}

func (x *LiveInfo) GetNewFanbase() int64 {
	if x != nil {
		return x.NewFanbase
	}
	return 0
}

func (x *LiveInfo) GetNewFanbasePrice() int64 {
	if x != nil {
		return x.NewFanbasePrice
	}
	return 0
}

func (x *LiveInfo) GetNewNobble() int64 {
	if x != nil {
		return x.NewNobble
	}
	return 0
}

func (x *LiveInfo) GetNewNobblePrice() int64 {
	if x != nil {
		return x.NewNobblePrice
	}
	return 0
}

func (x *LiveInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *LiveInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetLiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LiveId string `protobuf:"bytes,1,opt,name=live_id,json=liveId,proto3" json:"live_id,omitempty"` // 直播id
}

func (x *GetLiveReq) Reset() {
	*x = GetLiveReq{}
	mi := &file_svc_biz_room_live_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLiveReq) ProtoMessage() {}

func (x *GetLiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_live_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLiveReq.ProtoReflect.Descriptor instead.
func (*GetLiveReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_live_proto_rawDescGZIP(), []int{2}
}

func (x *GetLiveReq) GetLiveId() string {
	if x != nil {
		return x.LiveId
	}
	return ""
}

type GetLiveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Live *LiveInfo `protobuf:"bytes,1,opt,name=live,proto3" json:"live,omitempty"` // 直播信息
}

func (x *GetLiveResp) Reset() {
	*x = GetLiveResp{}
	mi := &file_svc_biz_room_live_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLiveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLiveResp) ProtoMessage() {}

func (x *GetLiveResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_live_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLiveResp.ProtoReflect.Descriptor instead.
func (*GetLiveResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_live_proto_rawDescGZIP(), []int{3}
}

func (x *GetLiveResp) GetLive() *LiveInfo {
	if x != nil {
		return x.Live
	}
	return nil
}

type MGetLiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LiveIds []string `protobuf:"bytes,1,rep,name=live_ids,json=liveIds,proto3" json:"live_ids,omitempty"` // 直播id列表
}

func (x *MGetLiveReq) Reset() {
	*x = MGetLiveReq{}
	mi := &file_svc_biz_room_live_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MGetLiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MGetLiveReq) ProtoMessage() {}

func (x *MGetLiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_live_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MGetLiveReq.ProtoReflect.Descriptor instead.
func (*MGetLiveReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_live_proto_rawDescGZIP(), []int{4}
}

func (x *MGetLiveReq) GetLiveIds() []string {
	if x != nil {
		return x.LiveIds
	}
	return nil
}

type MGetLiveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items map[string]*LiveInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 直播信息
}

func (x *MGetLiveResp) Reset() {
	*x = MGetLiveResp{}
	mi := &file_svc_biz_room_live_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MGetLiveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MGetLiveResp) ProtoMessage() {}

func (x *MGetLiveResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_live_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MGetLiveResp.ProtoReflect.Descriptor instead.
func (*MGetLiveResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_live_proto_rawDescGZIP(), []int{5}
}

func (x *MGetLiveResp) GetItems() map[string]*LiveInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListLiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32         `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                              // 页数
	Limit      int32         `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`                            // 条数
	StreamerId string        `protobuf:"bytes,3,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"` // 主播id
	StartAt    *TimeBeginEnd `protobuf:"bytes,4,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt      *TimeBeginEnd `protobuf:"bytes,5,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
}

func (x *ListLiveReq) Reset() {
	*x = ListLiveReq{}
	mi := &file_svc_biz_room_live_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLiveReq) ProtoMessage() {}

func (x *ListLiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_live_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLiveReq.ProtoReflect.Descriptor instead.
func (*ListLiveReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_live_proto_rawDescGZIP(), []int{6}
}

func (x *ListLiveReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListLiveReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListLiveReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *ListLiveReq) GetStartAt() *TimeBeginEnd {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *ListLiveReq) GetEndAt() *TimeBeginEnd {
	if x != nil {
		return x.EndAt
	}
	return nil
}

type ListLiveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*LiveInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // 直播信息
}

func (x *ListLiveResp) Reset() {
	*x = ListLiveResp{}
	mi := &file_svc_biz_room_live_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLiveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLiveResp) ProtoMessage() {}

func (x *ListLiveResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_live_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLiveResp.ProtoReflect.Descriptor instead.
func (*ListLiveResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_live_proto_rawDescGZIP(), []int{7}
}

func (x *ListLiveResp) GetItems() []*LiveInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type StatLiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamerId string        `protobuf:"bytes,1,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"` // 主播id
	StartAt    *TimeBeginEnd `protobuf:"bytes,2,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt      *TimeBeginEnd `protobuf:"bytes,3,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
}

func (x *StatLiveReq) Reset() {
	*x = StatLiveReq{}
	mi := &file_svc_biz_room_live_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatLiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatLiveReq) ProtoMessage() {}

func (x *StatLiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_live_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatLiveReq.ProtoReflect.Descriptor instead.
func (*StatLiveReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_live_proto_rawDescGZIP(), []int{8}
}

func (x *StatLiveReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *StatLiveReq) GetStartAt() *TimeBeginEnd {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *StatLiveReq) GetEndAt() *TimeBeginEnd {
	if x != nil {
		return x.EndAt
	}
	return nil
}

type StatLiveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count           int64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`                                              // 直播次数
	Duration        float32 `protobuf:"fixed32,2,opt,name=duration,proto3" json:"duration,omitempty"`                                       // 直播时长
	GiftTotalNum    int64   `protobuf:"varint,3,opt,name=gift_total_num,json=giftTotalNum,proto3" json:"gift_total_num,omitempty"`          // 礼物总数
	GiftTotalPrice  int64   `protobuf:"varint,4,opt,name=gift_total_price,json=giftTotalPrice,proto3" json:"gift_total_price,omitempty"`    // 礼物代币总数
	GiftTotalUser   int64   `protobuf:"varint,5,opt,name=gift_total_user,json=giftTotalUser,proto3" json:"gift_total_user,omitempty"`       // 礼物用户总数（每场直播礼物用户总数相加，不去重）
	TotalViewer     int64   `protobuf:"varint,6,opt,name=total_viewer,json=totalViewer,proto3" json:"total_viewer,omitempty"`               // 观看人数（每场直播观看人数相加，不去重）
	NewFollow       int64   `protobuf:"varint,7,opt,name=new_follow,json=newFollow,proto3" json:"new_follow,omitempty"`                     // 新增订阅（每场新增订阅人数相加，不去重）
	NewFanbase      int64   `protobuf:"varint,8,opt,name=new_fanbase,json=newFanbase,proto3" json:"new_fanbase,omitempty"`                  // 新增粉丝团（每场新增粉丝团相加，不去重）
	NewFanbasePrice int64   `protobuf:"varint,9,opt,name=new_fanbase_price,json=newFanbasePrice,proto3" json:"new_fanbase_price,omitempty"` // 新增粉丝团代币数（每场直播粉丝团付费数相加）
	NewNobble       int64   `protobuf:"varint,10,opt,name=new_nobble,json=newNobble,proto3" json:"new_nobble,omitempty"`                    // 新增贵族（每场直播贵族付费用户数相加）
	NewNobblePrice  int64   `protobuf:"varint,11,opt,name=new_nobble_price,json=newNobblePrice,proto3" json:"new_nobble_price,omitempty"`   // 新增贵族代币数（每场直播贵族付费数相加）
}

func (x *StatLiveResp) Reset() {
	*x = StatLiveResp{}
	mi := &file_svc_biz_room_live_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatLiveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatLiveResp) ProtoMessage() {}

func (x *StatLiveResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_live_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatLiveResp.ProtoReflect.Descriptor instead.
func (*StatLiveResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_live_proto_rawDescGZIP(), []int{9}
}

func (x *StatLiveResp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *StatLiveResp) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *StatLiveResp) GetGiftTotalNum() int64 {
	if x != nil {
		return x.GiftTotalNum
	}
	return 0
}

func (x *StatLiveResp) GetGiftTotalPrice() int64 {
	if x != nil {
		return x.GiftTotalPrice
	}
	return 0
}

func (x *StatLiveResp) GetGiftTotalUser() int64 {
	if x != nil {
		return x.GiftTotalUser
	}
	return 0
}

func (x *StatLiveResp) GetTotalViewer() int64 {
	if x != nil {
		return x.TotalViewer
	}
	return 0
}

func (x *StatLiveResp) GetNewFollow() int64 {
	if x != nil {
		return x.NewFollow
	}
	return 0
}

func (x *StatLiveResp) GetNewFanbase() int64 {
	if x != nil {
		return x.NewFanbase
	}
	return 0
}

func (x *StatLiveResp) GetNewFanbasePrice() int64 {
	if x != nil {
		return x.NewFanbasePrice
	}
	return 0
}

func (x *StatLiveResp) GetNewNobble() int64 {
	if x != nil {
		return x.NewNobble
	}
	return 0
}

func (x *StatLiveResp) GetNewNobblePrice() int64 {
	if x != nil {
		return x.NewNobblePrice
	}
	return 0
}

var File_svc_biz_room_live_proto protoreflect.FileDescriptor

var file_svc_biz_room_live_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x6c,
	0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x62, 0x65, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x62, 0x65, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0xb0, 0x05, 0x0a, 0x08, 0x4c,
	0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x41, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05,
	0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x67,
	0x69, 0x66, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x67,
	0x69, 0x66, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x67, 0x69, 0x66, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x67, 0x69, 0x66, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x66, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x77, 0x5f, 0x66, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x65, 0x77,
	0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x6f, 0x62, 0x62, 0x6c, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x6e, 0x65, 0x77, 0x4e, 0x6f, 0x62, 0x62, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6e,
	0x65, 0x77, 0x5f, 0x6e, 0x6f, 0x62, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x4e, 0x6f, 0x62, 0x62, 0x6c, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x80, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x25, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69,
	0x76, 0x65, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x04, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x4c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x76, 0x65, 0x22,
	0x28, 0x0a, 0x0b, 0x4d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x64, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x0c, 0x4d, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x50, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc2, 0x01, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45,
	0x6e, 0x64, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x65,
	0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x45, 0x6e, 0x64, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x22, 0x3c,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x76,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x98, 0x01, 0x0a,
	0x0b, 0x53, 0x74, 0x61, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x6e, 0x64, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72,
	0x6f, 0x6f, 0x6d, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x6e, 0x64,
	0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x22, 0x90, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74,
	0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x69,
	0x66, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x67, 0x69, 0x66, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d,
	0x12, 0x28, 0x0a, 0x10, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x67, 0x69, 0x66, 0x74,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x69,
	0x66, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x67, 0x69, 0x66, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x69, 0x65, 0x77,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56,
	0x69, 0x65, 0x77, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x66, 0x61, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x46, 0x61,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x77, 0x5f, 0x66, 0x61, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x6e, 0x65, 0x77, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x6f, 0x62, 0x62, 0x6c, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x4e, 0x6f, 0x62, 0x62, 0x6c, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x6f, 0x62, 0x62, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x4e,
	0x6f, 0x62, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0x8f, 0x02, 0x0a, 0x04, 0x4c,
	0x69, 0x76, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x18,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x08, 0x4d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x12,
	0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4d,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4d, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f,
	0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x08, 0x53, 0x74, 0x61,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e,
	0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x15, 0x5a, 0x13,
	0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x3b, 0x72,
	0x6f, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_room_live_proto_rawDescOnce sync.Once
	file_svc_biz_room_live_proto_rawDescData = file_svc_biz_room_live_proto_rawDesc
)

func file_svc_biz_room_live_proto_rawDescGZIP() []byte {
	file_svc_biz_room_live_proto_rawDescOnce.Do(func() {
		file_svc_biz_room_live_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_room_live_proto_rawDescData)
	})
	return file_svc_biz_room_live_proto_rawDescData
}

var file_svc_biz_room_live_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_svc_biz_room_live_proto_goTypes = []any{
	(*TimeBeginEnd)(nil),          // 0: svc.biz.room.TimeBeginEnd
	(*LiveInfo)(nil),              // 1: svc.biz.room.LiveInfo
	(*GetLiveReq)(nil),            // 2: svc.biz.room.GetLiveReq
	(*GetLiveResp)(nil),           // 3: svc.biz.room.GetLiveResp
	(*MGetLiveReq)(nil),           // 4: svc.biz.room.MGetLiveReq
	(*MGetLiveResp)(nil),          // 5: svc.biz.room.MGetLiveResp
	(*ListLiveReq)(nil),           // 6: svc.biz.room.ListLiveReq
	(*ListLiveResp)(nil),          // 7: svc.biz.room.ListLiveResp
	(*StatLiveReq)(nil),           // 8: svc.biz.room.StatLiveReq
	(*StatLiveResp)(nil),          // 9: svc.biz.room.StatLiveResp
	nil,                           // 10: svc.biz.room.MGetLiveResp.ItemsEntry
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_svc_biz_room_live_proto_depIdxs = []int32{
	11, // 0: svc.biz.room.TimeBeginEnd.beigin:type_name -> google.protobuf.Timestamp
	11, // 1: svc.biz.room.TimeBeginEnd.end:type_name -> google.protobuf.Timestamp
	11, // 2: svc.biz.room.LiveInfo.start_at:type_name -> google.protobuf.Timestamp
	11, // 3: svc.biz.room.LiveInfo.end_at:type_name -> google.protobuf.Timestamp
	11, // 4: svc.biz.room.LiveInfo.created_at:type_name -> google.protobuf.Timestamp
	11, // 5: svc.biz.room.LiveInfo.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 6: svc.biz.room.GetLiveResp.live:type_name -> svc.biz.room.LiveInfo
	10, // 7: svc.biz.room.MGetLiveResp.items:type_name -> svc.biz.room.MGetLiveResp.ItemsEntry
	0,  // 8: svc.biz.room.ListLiveReq.start_at:type_name -> svc.biz.room.TimeBeginEnd
	0,  // 9: svc.biz.room.ListLiveReq.end_at:type_name -> svc.biz.room.TimeBeginEnd
	1,  // 10: svc.biz.room.ListLiveResp.items:type_name -> svc.biz.room.LiveInfo
	0,  // 11: svc.biz.room.StatLiveReq.start_at:type_name -> svc.biz.room.TimeBeginEnd
	0,  // 12: svc.biz.room.StatLiveReq.end_at:type_name -> svc.biz.room.TimeBeginEnd
	1,  // 13: svc.biz.room.MGetLiveResp.ItemsEntry.value:type_name -> svc.biz.room.LiveInfo
	2,  // 14: svc.biz.room.Live.GetLive:input_type -> svc.biz.room.GetLiveReq
	4,  // 15: svc.biz.room.Live.MGetLive:input_type -> svc.biz.room.MGetLiveReq
	6,  // 16: svc.biz.room.Live.ListLive:input_type -> svc.biz.room.ListLiveReq
	8,  // 17: svc.biz.room.Live.StatLive:input_type -> svc.biz.room.StatLiveReq
	3,  // 18: svc.biz.room.Live.GetLive:output_type -> svc.biz.room.GetLiveResp
	5,  // 19: svc.biz.room.Live.MGetLive:output_type -> svc.biz.room.MGetLiveResp
	7,  // 20: svc.biz.room.Live.ListLive:output_type -> svc.biz.room.ListLiveResp
	9,  // 21: svc.biz.room.Live.StatLive:output_type -> svc.biz.room.StatLiveResp
	18, // [18:22] is the sub-list for method output_type
	14, // [14:18] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_svc_biz_room_live_proto_init() }
func file_svc_biz_room_live_proto_init() {
	if File_svc_biz_room_live_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_biz_room_live_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_room_live_proto_goTypes,
		DependencyIndexes: file_svc_biz_room_live_proto_depIdxs,
		MessageInfos:      file_svc_biz_room_live_proto_msgTypes,
	}.Build()
	File_svc_biz_room_live_proto = out.File
	file_svc_biz_room_live_proto_rawDesc = nil
	file_svc_biz_room_live_proto_goTypes = nil
	file_svc_biz_room_live_proto_depIdxs = nil
}
