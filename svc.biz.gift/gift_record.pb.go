// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: svc.biz.gift/gift_record.proto

package gift

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

type GiftRecordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GiftId     string                 `protobuf:"bytes,1,opt,name=gift_id,json=giftId,proto3" json:"gift_id,omitempty"`              // 礼物id
	GiftName   string                 `protobuf:"bytes,2,opt,name=gift_name,json=giftName,proto3" json:"gift_name,omitempty"`        // 礼物名
	Num        int32                  `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`                                 // 数量
	TotalPrice int32                  `protobuf:"varint,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"` // 总价
	FromUid    string                 `protobuf:"bytes,5,opt,name=from_uid,json=fromUid,proto3" json:"from_uid,omitempty"`           // 送礼uid
	StreamerId string                 `protobuf:"bytes,6,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`  // 主播uid
	RoomId     string                 `protobuf:"bytes,7,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`              // 房间id
	LiveId     string                 `protobuf:"bytes,8,opt,name=live_id,json=liveId,proto3" json:"live_id,omitempty"`              // 直播id
	SendTime   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`        // 创建时间
}

func (x *GiftRecordInfo) Reset() {
	*x = GiftRecordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_gift_gift_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftRecordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftRecordInfo) ProtoMessage() {}

func (x *GiftRecordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_gift_gift_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftRecordInfo.ProtoReflect.Descriptor instead.
func (*GiftRecordInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_gift_gift_record_proto_rawDescGZIP(), []int{0}
}

func (x *GiftRecordInfo) GetGiftId() string {
	if x != nil {
		return x.GiftId
	}
	return ""
}

func (x *GiftRecordInfo) GetGiftName() string {
	if x != nil {
		return x.GiftName
	}
	return ""
}

func (x *GiftRecordInfo) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *GiftRecordInfo) GetTotalPrice() int32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *GiftRecordInfo) GetFromUid() string {
	if x != nil {
		return x.FromUid
	}
	return ""
}

func (x *GiftRecordInfo) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *GiftRecordInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *GiftRecordInfo) GetLiveId() string {
	if x != nil {
		return x.LiveId
	}
	return ""
}

func (x *GiftRecordInfo) GetSendTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SendTime
	}
	return nil
}

type GetSendRecordListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                           // 第几页
	Limit     int64                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`                         // 每页几条数据
	FromUid   string                 `protobuf:"bytes,3,opt,name=from_uid,json=fromUid,proto3" json:"from_uid,omitempty"`       // 用户id
	StartTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"` // 开始时间
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`       // 结束时间
}

func (x *GetSendRecordListReq) Reset() {
	*x = GetSendRecordListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_gift_gift_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSendRecordListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSendRecordListReq) ProtoMessage() {}

func (x *GetSendRecordListReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_gift_gift_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSendRecordListReq.ProtoReflect.Descriptor instead.
func (*GetSendRecordListReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_gift_gift_record_proto_rawDescGZIP(), []int{1}
}

func (x *GetSendRecordListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetSendRecordListReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetSendRecordListReq) GetFromUid() string {
	if x != nil {
		return x.FromUid
	}
	return ""
}

func (x *GetSendRecordListReq) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *GetSendRecordListReq) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type GetSendRecordListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GiftRecordInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetSendRecordListResp) Reset() {
	*x = GetSendRecordListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_gift_gift_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSendRecordListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSendRecordListResp) ProtoMessage() {}

func (x *GetSendRecordListResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_gift_gift_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSendRecordListResp.ProtoReflect.Descriptor instead.
func (*GetSendRecordListResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_gift_gift_record_proto_rawDescGZIP(), []int{2}
}

func (x *GetSendRecordListResp) GetItems() []*GiftRecordInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetGetRecordListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                              // 第几页
	Limit      int64                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`                            // 每页几条数据
	StreamerId string                 `protobuf:"bytes,3,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"` // 主播uid
	RoomId     string                 `protobuf:"bytes,4,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`             // 房间id
	LiveId     string                 `protobuf:"bytes,5,opt,name=live_id,json=liveId,proto3" json:"live_id,omitempty"`             // 直播id
	StartTime  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`    // 开始时间
	EndTime    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`          // 结束时间
}

func (x *GetGetRecordListReq) Reset() {
	*x = GetGetRecordListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_gift_gift_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGetRecordListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGetRecordListReq) ProtoMessage() {}

func (x *GetGetRecordListReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_gift_gift_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGetRecordListReq.ProtoReflect.Descriptor instead.
func (*GetGetRecordListReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_gift_gift_record_proto_rawDescGZIP(), []int{3}
}

func (x *GetGetRecordListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetGetRecordListReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetGetRecordListReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *GetGetRecordListReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *GetGetRecordListReq) GetLiveId() string {
	if x != nil {
		return x.LiveId
	}
	return ""
}

func (x *GetGetRecordListReq) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *GetGetRecordListReq) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type GetGetRecordListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GiftRecordInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetGetRecordListResp) Reset() {
	*x = GetGetRecordListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_gift_gift_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGetRecordListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGetRecordListResp) ProtoMessage() {}

func (x *GetGetRecordListResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_gift_gift_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGetRecordListResp.ProtoReflect.Descriptor instead.
func (*GetGetRecordListResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_gift_gift_record_proto_rawDescGZIP(), []int{4}
}

func (x *GetGetRecordListResp) GetItems() []*GiftRecordInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetLiveStatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamerId string `protobuf:"bytes,1,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"` // 主播uid
	LiveId     string `protobuf:"bytes,2,opt,name=live_id,json=liveId,proto3" json:"live_id,omitempty"`             // 直播id
}

func (x *GetLiveStatReq) Reset() {
	*x = GetLiveStatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_gift_gift_record_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLiveStatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLiveStatReq) ProtoMessage() {}

func (x *GetLiveStatReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_gift_gift_record_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLiveStatReq.ProtoReflect.Descriptor instead.
func (*GetLiveStatReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_gift_gift_record_proto_rawDescGZIP(), []int{5}
}

func (x *GetLiveStatReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *GetLiveStatReq) GetLiveId() string {
	if x != nil {
		return x.LiveId
	}
	return ""
}

type GetLiveStatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalNum   int32 `protobuf:"varint,1,opt,name=total_num,json=totalNum,proto3" json:"total_num,omitempty"`       // 礼物总数
	TotalPrice int32 `protobuf:"varint,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"` // 礼物代币总数
	TotalUser  int32 `protobuf:"varint,3,opt,name=total_user,json=totalUser,proto3" json:"total_user,omitempty"`    // 礼物用户总数
}

func (x *GetLiveStatResp) Reset() {
	*x = GetLiveStatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_gift_gift_record_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLiveStatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLiveStatResp) ProtoMessage() {}

func (x *GetLiveStatResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_gift_gift_record_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLiveStatResp.ProtoReflect.Descriptor instead.
func (*GetLiveStatResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_gift_gift_record_proto_rawDescGZIP(), []int{6}
}

func (x *GetLiveStatResp) GetTotalNum() int32 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *GetLiveStatResp) GetTotalPrice() int32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *GetLiveStatResp) GetTotalUser() int32 {
	if x != nil {
		return x.TotalUser
	}
	return 0
}

var File_svc_biz_gift_gift_record_proto protoreflect.FileDescriptor

var file_svc_biz_gift_gift_record_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2f, 0x67,
	0x69, 0x66, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa0, 0x02, 0x0a, 0x0e, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x66, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x69, 0x66, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x69, 0x66, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x72, 0x6f, 0x6d, 0x55, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x69, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x84, 0x02, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69,
	0x76, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e, 0x47, 0x69, 0x66,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x64, 0x22, 0x6e,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x32, 0x97,
	0x02, 0x0a, 0x0a, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x5e, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x67, 0x69, 0x66,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x5b, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x21, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x67, 0x69, 0x66, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x67,
	0x69, 0x66, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x73, 0x76,
	0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x67, 0x69, 0x66, 0x74, 0x3b, 0x67, 0x69, 0x66, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_gift_gift_record_proto_rawDescOnce sync.Once
	file_svc_biz_gift_gift_record_proto_rawDescData = file_svc_biz_gift_gift_record_proto_rawDesc
)

func file_svc_biz_gift_gift_record_proto_rawDescGZIP() []byte {
	file_svc_biz_gift_gift_record_proto_rawDescOnce.Do(func() {
		file_svc_biz_gift_gift_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_gift_gift_record_proto_rawDescData)
	})
	return file_svc_biz_gift_gift_record_proto_rawDescData
}

var file_svc_biz_gift_gift_record_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_svc_biz_gift_gift_record_proto_goTypes = []any{
	(*GiftRecordInfo)(nil),        // 0: svc.biz.gift.GiftRecordInfo
	(*GetSendRecordListReq)(nil),  // 1: svc.biz.gift.GetSendRecordListReq
	(*GetSendRecordListResp)(nil), // 2: svc.biz.gift.GetSendRecordListResp
	(*GetGetRecordListReq)(nil),   // 3: svc.biz.gift.GetGetRecordListReq
	(*GetGetRecordListResp)(nil),  // 4: svc.biz.gift.GetGetRecordListResp
	(*GetLiveStatReq)(nil),        // 5: svc.biz.gift.GetLiveStatReq
	(*GetLiveStatResp)(nil),       // 6: svc.biz.gift.GetLiveStatResp
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_svc_biz_gift_gift_record_proto_depIdxs = []int32{
	7,  // 0: svc.biz.gift.GiftRecordInfo.send_time:type_name -> google.protobuf.Timestamp
	7,  // 1: svc.biz.gift.GetSendRecordListReq.start_time:type_name -> google.protobuf.Timestamp
	7,  // 2: svc.biz.gift.GetSendRecordListReq.end_time:type_name -> google.protobuf.Timestamp
	0,  // 3: svc.biz.gift.GetSendRecordListResp.items:type_name -> svc.biz.gift.GiftRecordInfo
	7,  // 4: svc.biz.gift.GetGetRecordListReq.start_time:type_name -> google.protobuf.Timestamp
	7,  // 5: svc.biz.gift.GetGetRecordListReq.end_time:type_name -> google.protobuf.Timestamp
	0,  // 6: svc.biz.gift.GetGetRecordListResp.items:type_name -> svc.biz.gift.GiftRecordInfo
	1,  // 7: svc.biz.gift.GiftRecord.GetSendRecordList:input_type -> svc.biz.gift.GetSendRecordListReq
	3,  // 8: svc.biz.gift.GiftRecord.GetGetRecordList:input_type -> svc.biz.gift.GetGetRecordListReq
	5,  // 9: svc.biz.gift.GiftRecord.GetLiveStat:input_type -> svc.biz.gift.GetLiveStatReq
	2,  // 10: svc.biz.gift.GiftRecord.GetSendRecordList:output_type -> svc.biz.gift.GetSendRecordListResp
	4,  // 11: svc.biz.gift.GiftRecord.GetGetRecordList:output_type -> svc.biz.gift.GetGetRecordListResp
	6,  // 12: svc.biz.gift.GiftRecord.GetLiveStat:output_type -> svc.biz.gift.GetLiveStatResp
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_svc_biz_gift_gift_record_proto_init() }
func file_svc_biz_gift_gift_record_proto_init() {
	if File_svc_biz_gift_gift_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_biz_gift_gift_record_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GiftRecordInfo); i {
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
		file_svc_biz_gift_gift_record_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetSendRecordListReq); i {
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
		file_svc_biz_gift_gift_record_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetSendRecordListResp); i {
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
		file_svc_biz_gift_gift_record_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetGetRecordListReq); i {
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
		file_svc_biz_gift_gift_record_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetGetRecordListResp); i {
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
		file_svc_biz_gift_gift_record_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetLiveStatReq); i {
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
		file_svc_biz_gift_gift_record_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetLiveStatResp); i {
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
			RawDescriptor: file_svc_biz_gift_gift_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_gift_gift_record_proto_goTypes,
		DependencyIndexes: file_svc_biz_gift_gift_record_proto_depIdxs,
		MessageInfos:      file_svc_biz_gift_gift_record_proto_msgTypes,
	}.Build()
	File_svc_biz_gift_gift_record_proto = out.File
	file_svc_biz_gift_gift_record_proto_rawDesc = nil
	file_svc_biz_gift_gift_record_proto_goTypes = nil
	file_svc_biz_gift_gift_record_proto_depIdxs = nil
}
