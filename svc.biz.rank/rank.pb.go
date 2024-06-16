// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: svc.biz.rank/rank.proto

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
		mi := &file_svc_biz_rank_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankMemberInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankMemberInfo) ProtoMessage() {}

func (x *RankMemberInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_rank_rank_proto_msgTypes[0]
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
	return file_svc_biz_rank_rank_proto_rawDescGZIP(), []int{0}
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

type GetLiveOnlineRankMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLiveOnlineRankMemberReq) Reset() {
	*x = GetLiveOnlineRankMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_rank_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLiveOnlineRankMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLiveOnlineRankMemberReq) ProtoMessage() {}

func (x *GetLiveOnlineRankMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_rank_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLiveOnlineRankMemberReq.ProtoReflect.Descriptor instead.
func (*GetLiveOnlineRankMemberReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_rank_rank_proto_rawDescGZIP(), []int{1}
}

type GetLiveOnlineRankMemberResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*RankMemberInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetLiveOnlineRankMemberResp) Reset() {
	*x = GetLiveOnlineRankMemberResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_rank_rank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLiveOnlineRankMemberResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLiveOnlineRankMemberResp) ProtoMessage() {}

func (x *GetLiveOnlineRankMemberResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_rank_rank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLiveOnlineRankMemberResp.ProtoReflect.Descriptor instead.
func (*GetLiveOnlineRankMemberResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_rank_rank_proto_rawDescGZIP(), []int{2}
}

func (x *GetLiveOnlineRankMemberResp) GetItems() []*RankMemberInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_svc_biz_rank_rank_proto protoreflect.FileDescriptor

var file_svc_biz_rank_rank_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x72,
	0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x43, 0x0a, 0x0e, 0x52, 0x61, 0x6e, 0x6b, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x1c, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x61, 0x6e,
	0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x51, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x78, 0x0a,
	0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x70, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x28, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x61, 0x6e,
	0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76,
	0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x3b, 0x72, 0x61, 0x6e, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_rank_rank_proto_rawDescOnce sync.Once
	file_svc_biz_rank_rank_proto_rawDescData = file_svc_biz_rank_rank_proto_rawDesc
)

func file_svc_biz_rank_rank_proto_rawDescGZIP() []byte {
	file_svc_biz_rank_rank_proto_rawDescOnce.Do(func() {
		file_svc_biz_rank_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_rank_rank_proto_rawDescData)
	})
	return file_svc_biz_rank_rank_proto_rawDescData
}

var file_svc_biz_rank_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_svc_biz_rank_rank_proto_goTypes = []any{
	(*RankMemberInfo)(nil),              // 0: svc.biz.rank.RankMemberInfo
	(*GetLiveOnlineRankMemberReq)(nil),  // 1: svc.biz.rank.GetLiveOnlineRankMemberReq
	(*GetLiveOnlineRankMemberResp)(nil), // 2: svc.biz.rank.GetLiveOnlineRankMemberResp
}
var file_svc_biz_rank_rank_proto_depIdxs = []int32{
	0, // 0: svc.biz.rank.GetLiveOnlineRankMemberResp.items:type_name -> svc.biz.rank.RankMemberInfo
	1, // 1: svc.biz.rank.Rank.GetLiveOnlineRankMember:input_type -> svc.biz.rank.GetLiveOnlineRankMemberReq
	2, // 2: svc.biz.rank.Rank.GetLiveOnlineRankMember:output_type -> svc.biz.rank.GetLiveOnlineRankMemberResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_svc_biz_rank_rank_proto_init() }
func file_svc_biz_rank_rank_proto_init() {
	if File_svc_biz_rank_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_biz_rank_rank_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_svc_biz_rank_rank_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetLiveOnlineRankMemberReq); i {
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
		file_svc_biz_rank_rank_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetLiveOnlineRankMemberResp); i {
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
			RawDescriptor: file_svc_biz_rank_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_rank_rank_proto_goTypes,
		DependencyIndexes: file_svc_biz_rank_rank_proto_depIdxs,
		MessageInfos:      file_svc_biz_rank_rank_proto_msgTypes,
	}.Build()
	File_svc_biz_rank_rank_proto = out.File
	file_svc_biz_rank_rank_proto_rawDesc = nil
	file_svc_biz_rank_rank_proto_goTypes = nil
	file_svc_biz_rank_rank_proto_depIdxs = nil
}
