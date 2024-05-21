// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: svc.biz.room/room.proto

package room

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

type RoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoomInfo) Reset() {
	*x = RoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomInfo) ProtoMessage() {}

func (x *RoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomInfo.ProtoReflect.Descriptor instead.
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_room_proto_rawDescGZIP(), []int{0}
}

type GetRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRoomReq) Reset() {
	*x = GetRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomReq) ProtoMessage() {}

func (x *GetRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomReq.ProtoReflect.Descriptor instead.
func (*GetRoomReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_room_proto_rawDescGZIP(), []int{1}
}

type GetRoomResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRoomResp) Reset() {
	*x = GetRoomResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomResp) ProtoMessage() {}

func (x *GetRoomResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomResp.ProtoReflect.Descriptor instead.
func (*GetRoomResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_room_proto_rawDescGZIP(), []int{2}
}

var File_svc_biz_room_room_proto protoreflect.FileDescriptor

var file_svc_biz_room_room_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x72,
	0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x22, 0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x32, 0x48, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x18, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x3b, 0x72, 0x6f, 0x6f,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_room_room_proto_rawDescOnce sync.Once
	file_svc_biz_room_room_proto_rawDescData = file_svc_biz_room_room_proto_rawDesc
)

func file_svc_biz_room_room_proto_rawDescGZIP() []byte {
	file_svc_biz_room_room_proto_rawDescOnce.Do(func() {
		file_svc_biz_room_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_room_room_proto_rawDescData)
	})
	return file_svc_biz_room_room_proto_rawDescData
}

var file_svc_biz_room_room_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_svc_biz_room_room_proto_goTypes = []interface{}{
	(*RoomInfo)(nil),    // 0: svc.biz.room.RoomInfo
	(*GetRoomReq)(nil),  // 1: svc.biz.room.GetRoomReq
	(*GetRoomResp)(nil), // 2: svc.biz.room.GetRoomResp
}
var file_svc_biz_room_room_proto_depIdxs = []int32{
	1, // 0: svc.biz.room.Room.GetRoom:input_type -> svc.biz.room.GetRoomReq
	2, // 1: svc.biz.room.Room.GetRoom:output_type -> svc.biz.room.GetRoomResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_biz_room_room_proto_init() }
func file_svc_biz_room_room_proto_init() {
	if File_svc_biz_room_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_biz_room_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomInfo); i {
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
		file_svc_biz_room_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomReq); i {
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
		file_svc_biz_room_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomResp); i {
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
			RawDescriptor: file_svc_biz_room_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_room_room_proto_goTypes,
		DependencyIndexes: file_svc_biz_room_room_proto_depIdxs,
		MessageInfos:      file_svc_biz_room_room_proto_msgTypes,
	}.Build()
	File_svc_biz_room_room_proto = out.File
	file_svc_biz_room_room_proto_rawDesc = nil
	file_svc_biz_room_room_proto_goTypes = nil
	file_svc_biz_room_room_proto_depIdxs = nil
}
