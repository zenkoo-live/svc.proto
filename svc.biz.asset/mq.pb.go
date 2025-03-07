// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: svc.biz.asset/mq.proto

package asset

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

// topic: asset.user_money.incr
type AssetUserMoneyIncrTopicInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MerchantId    string                 `protobuf:"bytes,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	Uid           string                 `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	IncrValue     int64                  `protobuf:"varint,3,opt,name=incr_value,json=incrValue,proto3" json:"incr_value,omitempty"`
	NewValue      int64                  `protobuf:"varint,4,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	TransType     int64                  `protobuf:"varint,5,opt,name=trans_type,json=transType,proto3" json:"trans_type,omitempty"`
	TradeId       string                 `protobuf:"bytes,6,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	DetailId      string                 `protobuf:"bytes,7,opt,name=detail_id,json=detailId,proto3" json:"detail_id,omitempty"`
	SerialNumber  int64                  `protobuf:"varint,8,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Summary       string                 `protobuf:"bytes,9,opt,name=summary,proto3" json:"summary,omitempty"`
	TransTime     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=trans_time,json=transTime,proto3" json:"trans_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetUserMoneyIncrTopicInfo) Reset() {
	*x = AssetUserMoneyIncrTopicInfo{}
	mi := &file_svc_biz_asset_mq_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetUserMoneyIncrTopicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetUserMoneyIncrTopicInfo) ProtoMessage() {}

func (x *AssetUserMoneyIncrTopicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_asset_mq_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetUserMoneyIncrTopicInfo.ProtoReflect.Descriptor instead.
func (*AssetUserMoneyIncrTopicInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_asset_mq_proto_rawDescGZIP(), []int{0}
}

func (x *AssetUserMoneyIncrTopicInfo) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *AssetUserMoneyIncrTopicInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AssetUserMoneyIncrTopicInfo) GetIncrValue() int64 {
	if x != nil {
		return x.IncrValue
	}
	return 0
}

func (x *AssetUserMoneyIncrTopicInfo) GetNewValue() int64 {
	if x != nil {
		return x.NewValue
	}
	return 0
}

func (x *AssetUserMoneyIncrTopicInfo) GetTransType() int64 {
	if x != nil {
		return x.TransType
	}
	return 0
}

func (x *AssetUserMoneyIncrTopicInfo) GetTradeId() string {
	if x != nil {
		return x.TradeId
	}
	return ""
}

func (x *AssetUserMoneyIncrTopicInfo) GetDetailId() string {
	if x != nil {
		return x.DetailId
	}
	return ""
}

func (x *AssetUserMoneyIncrTopicInfo) GetSerialNumber() int64 {
	if x != nil {
		return x.SerialNumber
	}
	return 0
}

func (x *AssetUserMoneyIncrTopicInfo) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *AssetUserMoneyIncrTopicInfo) GetTransTime() *timestamppb.Timestamp {
	if x != nil {
		return x.TransTime
	}
	return nil
}

var File_svc_biz_asset_mq_proto protoreflect.FileDescriptor

var file_svc_biz_asset_mq_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f,
	0x6d, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a, 0x1b, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x49, 0x6e, 0x63, 0x72, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x6e, 0x63, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x69, 0x6e, 0x63, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65,
	0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e,
	0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x39, 0x0a,
	0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x73, 0x76,
	0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x3b, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_svc_biz_asset_mq_proto_rawDescOnce sync.Once
	file_svc_biz_asset_mq_proto_rawDescData []byte
)

func file_svc_biz_asset_mq_proto_rawDescGZIP() []byte {
	file_svc_biz_asset_mq_proto_rawDescOnce.Do(func() {
		file_svc_biz_asset_mq_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_svc_biz_asset_mq_proto_rawDesc), len(file_svc_biz_asset_mq_proto_rawDesc)))
	})
	return file_svc_biz_asset_mq_proto_rawDescData
}

var file_svc_biz_asset_mq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_svc_biz_asset_mq_proto_goTypes = []any{
	(*AssetUserMoneyIncrTopicInfo)(nil), // 0: svc.biz.asset.AssetUserMoneyIncrTopicInfo
	(*timestamppb.Timestamp)(nil),       // 1: google.protobuf.Timestamp
}
var file_svc_biz_asset_mq_proto_depIdxs = []int32{
	1, // 0: svc.biz.asset.AssetUserMoneyIncrTopicInfo.trans_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_svc_biz_asset_mq_proto_init() }
func file_svc_biz_asset_mq_proto_init() {
	if File_svc_biz_asset_mq_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_svc_biz_asset_mq_proto_rawDesc), len(file_svc_biz_asset_mq_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_svc_biz_asset_mq_proto_goTypes,
		DependencyIndexes: file_svc_biz_asset_mq_proto_depIdxs,
		MessageInfos:      file_svc_biz_asset_mq_proto_msgTypes,
	}.Build()
	File_svc_biz_asset_mq_proto = out.File
	file_svc_biz_asset_mq_proto_goTypes = nil
	file_svc_biz_asset_mq_proto_depIdxs = nil
}
