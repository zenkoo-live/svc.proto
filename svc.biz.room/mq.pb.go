// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: svc.biz.room/mq.proto

package room

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// topic: topic.room.start_live
type RoomStartLiveTopicInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Room          *RoomInfo              `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	StartLiveAt   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_live_at,json=startLiveAt,proto3" json:"start_live_at,omitempty"` // 开播时间
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomStartLiveTopicInfo) Reset() {
	*x = RoomStartLiveTopicInfo{}
	mi := &file_svc_biz_room_mq_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomStartLiveTopicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomStartLiveTopicInfo) ProtoMessage() {}

func (x *RoomStartLiveTopicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_mq_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomStartLiveTopicInfo.ProtoReflect.Descriptor instead.
func (*RoomStartLiveTopicInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_mq_proto_rawDescGZIP(), []int{0}
}

func (x *RoomStartLiveTopicInfo) GetRoom() *RoomInfo {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *RoomStartLiveTopicInfo) GetStartLiveAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartLiveAt
	}
	return nil
}

// topic: topic.room.stop_live
type RoomStopLiveTopicInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Room          *RoomInfo              `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	StopLiveAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=stop_live_at,json=stopLiveAt,proto3" json:"stop_live_at,omitempty"` // 关播时间
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomStopLiveTopicInfo) Reset() {
	*x = RoomStopLiveTopicInfo{}
	mi := &file_svc_biz_room_mq_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomStopLiveTopicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomStopLiveTopicInfo) ProtoMessage() {}

func (x *RoomStopLiveTopicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_mq_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomStopLiveTopicInfo.ProtoReflect.Descriptor instead.
func (*RoomStopLiveTopicInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_mq_proto_rawDescGZIP(), []int{1}
}

func (x *RoomStopLiveTopicInfo) GetRoom() *RoomInfo {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *RoomStopLiveTopicInfo) GetStopLiveAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StopLiveAt
	}
	return nil
}

var File_svc_biz_room_mq_proto protoreflect.FileDescriptor

var file_svc_biz_room_mq_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x6d,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e,
	0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x84, 0x01, 0x0a, 0x16, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x76,
	0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x3e, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x6c, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x4c, 0x69, 0x76, 0x65, 0x41, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x52, 0x6f, 0x6f, 0x6d, 0x53,
	0x74, 0x6f, 0x70, 0x4c, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2a, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x3c, 0x0a, 0x0c,
	0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x73, 0x74, 0x6f, 0x70, 0x4c, 0x69, 0x76, 0x65, 0x41, 0x74, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x3b, 0x72, 0x6f, 0x6f,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_svc_biz_room_mq_proto_rawDescOnce sync.Once
	file_svc_biz_room_mq_proto_rawDescData []byte
)

func file_svc_biz_room_mq_proto_rawDescGZIP() []byte {
	file_svc_biz_room_mq_proto_rawDescOnce.Do(func() {
		file_svc_biz_room_mq_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_svc_biz_room_mq_proto_rawDesc), len(file_svc_biz_room_mq_proto_rawDesc)))
	})
	return file_svc_biz_room_mq_proto_rawDescData
}

var file_svc_biz_room_mq_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_svc_biz_room_mq_proto_goTypes = []any{
	(*RoomStartLiveTopicInfo)(nil), // 0: svc.biz.room.RoomStartLiveTopicInfo
	(*RoomStopLiveTopicInfo)(nil),  // 1: svc.biz.room.RoomStopLiveTopicInfo
	(*RoomInfo)(nil),               // 2: svc.biz.room.RoomInfo
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_svc_biz_room_mq_proto_depIdxs = []int32{
	2, // 0: svc.biz.room.RoomStartLiveTopicInfo.room:type_name -> svc.biz.room.RoomInfo
	3, // 1: svc.biz.room.RoomStartLiveTopicInfo.start_live_at:type_name -> google.protobuf.Timestamp
	2, // 2: svc.biz.room.RoomStopLiveTopicInfo.room:type_name -> svc.biz.room.RoomInfo
	3, // 3: svc.biz.room.RoomStopLiveTopicInfo.stop_live_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_svc_biz_room_mq_proto_init() }
func file_svc_biz_room_mq_proto_init() {
	if File_svc_biz_room_mq_proto != nil {
		return
	}
	file_svc_biz_room_room_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_svc_biz_room_mq_proto_rawDesc), len(file_svc_biz_room_mq_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_svc_biz_room_mq_proto_goTypes,
		DependencyIndexes: file_svc_biz_room_mq_proto_depIdxs,
		MessageInfos:      file_svc_biz_room_mq_proto_msgTypes,
	}.Build()
	File_svc_biz_room_mq_proto = out.File
	file_svc_biz_room_mq_proto_goTypes = nil
	file_svc_biz_room_mq_proto_depIdxs = nil
}
