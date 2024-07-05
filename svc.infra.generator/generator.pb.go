// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: svc.infra.generator/generator.proto

package generator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InitDBResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InitDBResp) Reset() {
	*x = InitDBResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_generator_generator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitDBResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitDBResp) ProtoMessage() {}

func (x *InitDBResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_generator_generator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitDBResp.ProtoReflect.Descriptor instead.
func (*InitDBResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_generator_generator_proto_rawDescGZIP(), []int{0}
}

func (x *InitDBResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type InitIDGeneratorReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generator string `protobuf:"bytes,1,opt,name=generator,proto3" json:"generator,omitempty"`
}

func (x *InitIDGeneratorReq) Reset() {
	*x = InitIDGeneratorReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_generator_generator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitIDGeneratorReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitIDGeneratorReq) ProtoMessage() {}

func (x *InitIDGeneratorReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_generator_generator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitIDGeneratorReq.ProtoReflect.Descriptor instead.
func (*InitIDGeneratorReq) Descriptor() ([]byte, []int) {
	return file_svc_infra_generator_generator_proto_rawDescGZIP(), []int{1}
}

func (x *InitIDGeneratorReq) GetGenerator() string {
	if x != nil {
		return x.Generator
	}
	return ""
}

type InitIDGeneratorResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InitIDGeneratorResp) Reset() {
	*x = InitIDGeneratorResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_generator_generator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitIDGeneratorResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitIDGeneratorResp) ProtoMessage() {}

func (x *InitIDGeneratorResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_generator_generator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitIDGeneratorResp.ProtoReflect.Descriptor instead.
func (*InitIDGeneratorResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_generator_generator_proto_rawDescGZIP(), []int{2}
}

func (x *InitIDGeneratorResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type OrdianIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generator string `protobuf:"bytes,1,opt,name=generator,proto3" json:"generator,omitempty"`
	Id        int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrdianIDReq) Reset() {
	*x = OrdianIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_generator_generator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdianIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdianIDReq) ProtoMessage() {}

func (x *OrdianIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_generator_generator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdianIDReq.ProtoReflect.Descriptor instead.
func (*OrdianIDReq) Descriptor() ([]byte, []int) {
	return file_svc_infra_generator_generator_proto_rawDescGZIP(), []int{3}
}

func (x *OrdianIDReq) GetGenerator() string {
	if x != nil {
		return x.Generator
	}
	return ""
}

func (x *OrdianIDReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrdianIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OrdianIDResp) Reset() {
	*x = OrdianIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_generator_generator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdianIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdianIDResp) ProtoMessage() {}

func (x *OrdianIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_generator_generator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdianIDResp.ProtoReflect.Descriptor instead.
func (*OrdianIDResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_generator_generator_proto_rawDescGZIP(), []int{4}
}

func (x *OrdianIDResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type NextIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generator string `protobuf:"bytes,1,opt,name=generator,proto3" json:"generator,omitempty"`
}

func (x *NextIDReq) Reset() {
	*x = NextIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_generator_generator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextIDReq) ProtoMessage() {}

func (x *NextIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_generator_generator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextIDReq.ProtoReflect.Descriptor instead.
func (*NextIDReq) Descriptor() ([]byte, []int) {
	return file_svc_infra_generator_generator_proto_rawDescGZIP(), []int{5}
}

func (x *NextIDReq) GetGenerator() string {
	if x != nil {
		return x.Generator
	}
	return ""
}

type NextIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NextIDResp) Reset() {
	*x = NextIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_generator_generator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextIDResp) ProtoMessage() {}

func (x *NextIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_generator_generator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextIDResp.ProtoReflect.Descriptor instead.
func (*NextIDResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_generator_generator_proto_rawDescGZIP(), []int{6}
}

func (x *NextIDResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IsPrettyIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IsPrettyIDReq) Reset() {
	*x = IsPrettyIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_generator_generator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsPrettyIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsPrettyIDReq) ProtoMessage() {}

func (x *IsPrettyIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_generator_generator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsPrettyIDReq.ProtoReflect.Descriptor instead.
func (*IsPrettyIDReq) Descriptor() ([]byte, []int) {
	return file_svc_infra_generator_generator_proto_rawDescGZIP(), []int{7}
}

func (x *IsPrettyIDReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IsPrettyIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *IsPrettyIDResp) Reset() {
	*x = IsPrettyIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_generator_generator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsPrettyIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsPrettyIDResp) ProtoMessage() {}

func (x *IsPrettyIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_generator_generator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsPrettyIDResp.ProtoReflect.Descriptor instead.
func (*IsPrettyIDResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_generator_generator_proto_rawDescGZIP(), []int{8}
}

func (x *IsPrettyIDResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_svc_infra_generator_generator_proto protoreflect.FileDescriptor

var file_svc_infra_generator_generator_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x44,
	0x42, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x32, 0x0a,
	0x12, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x44, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x22, 0x2d, 0x0a, 0x13, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x44, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x3b, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12,
	0x1c, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a,
	0x0c, 0x4f, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x29, 0x0a, 0x09, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x22, 0x1c, 0x0a, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f,
	0x0a, 0x0d, 0x49, 0x73, 0x50, 0x72, 0x65, 0x74, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x28, 0x0a, 0x0e, 0x49, 0x73, 0x50, 0x72, 0x65, 0x74, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xa7, 0x03, 0x0a, 0x09, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x06, 0x49, 0x6e, 0x69, 0x74, 0x44,
	0x42, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x12, 0x64, 0x0a, 0x0f, 0x49, 0x6e,
	0x69, 0x74, 0x49, 0x44, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x27, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x44, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x49, 0x44, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x4f, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x21,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x49, 0x0a, 0x06, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x55, 0x0a, 0x0a,
	0x49, 0x73, 0x50, 0x72, 0x65, 0x74, 0x74, 0x79, 0x49, 0x44, 0x12, 0x22, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x49, 0x73, 0x50, 0x72, 0x65, 0x74, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x23,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x73, 0x50, 0x72, 0x65, 0x74, 0x74, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x21, 0x5a, 0x1f, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x3b, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_infra_generator_generator_proto_rawDescOnce sync.Once
	file_svc_infra_generator_generator_proto_rawDescData = file_svc_infra_generator_generator_proto_rawDesc
)

func file_svc_infra_generator_generator_proto_rawDescGZIP() []byte {
	file_svc_infra_generator_generator_proto_rawDescOnce.Do(func() {
		file_svc_infra_generator_generator_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_infra_generator_generator_proto_rawDescData)
	})
	return file_svc_infra_generator_generator_proto_rawDescData
}

var file_svc_infra_generator_generator_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_svc_infra_generator_generator_proto_goTypes = []any{
	(*InitDBResp)(nil),          // 0: svc.infra.generator.InitDBResp
	(*InitIDGeneratorReq)(nil),  // 1: svc.infra.generator.InitIDGeneratorReq
	(*InitIDGeneratorResp)(nil), // 2: svc.infra.generator.InitIDGeneratorResp
	(*OrdianIDReq)(nil),         // 3: svc.infra.generator.OrdianIDReq
	(*OrdianIDResp)(nil),        // 4: svc.infra.generator.OrdianIDResp
	(*NextIDReq)(nil),           // 5: svc.infra.generator.NextIDReq
	(*NextIDResp)(nil),          // 6: svc.infra.generator.NextIDResp
	(*IsPrettyIDReq)(nil),       // 7: svc.infra.generator.IsPrettyIDReq
	(*IsPrettyIDResp)(nil),      // 8: svc.infra.generator.IsPrettyIDResp
	(*emptypb.Empty)(nil),       // 9: google.protobuf.Empty
}
var file_svc_infra_generator_generator_proto_depIdxs = []int32{
	9, // 0: svc.infra.generator.Generator.InitDB:input_type -> google.protobuf.Empty
	1, // 1: svc.infra.generator.Generator.InitIDGenerator:input_type -> svc.infra.generator.InitIDGeneratorReq
	3, // 2: svc.infra.generator.Generator.OrdianID:input_type -> svc.infra.generator.OrdianIDReq
	5, // 3: svc.infra.generator.Generator.NextID:input_type -> svc.infra.generator.NextIDReq
	7, // 4: svc.infra.generator.Generator.IsPrettyID:input_type -> svc.infra.generator.IsPrettyIDReq
	0, // 5: svc.infra.generator.Generator.InitDB:output_type -> svc.infra.generator.InitDBResp
	2, // 6: svc.infra.generator.Generator.InitIDGenerator:output_type -> svc.infra.generator.InitIDGeneratorResp
	4, // 7: svc.infra.generator.Generator.OrdianID:output_type -> svc.infra.generator.OrdianIDResp
	6, // 8: svc.infra.generator.Generator.NextID:output_type -> svc.infra.generator.NextIDResp
	8, // 9: svc.infra.generator.Generator.IsPrettyID:output_type -> svc.infra.generator.IsPrettyIDResp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_infra_generator_generator_proto_init() }
func file_svc_infra_generator_generator_proto_init() {
	if File_svc_infra_generator_generator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_infra_generator_generator_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InitDBResp); i {
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
		file_svc_infra_generator_generator_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InitIDGeneratorReq); i {
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
		file_svc_infra_generator_generator_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InitIDGeneratorResp); i {
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
		file_svc_infra_generator_generator_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*OrdianIDReq); i {
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
		file_svc_infra_generator_generator_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*OrdianIDResp); i {
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
		file_svc_infra_generator_generator_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*NextIDReq); i {
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
		file_svc_infra_generator_generator_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*NextIDResp); i {
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
		file_svc_infra_generator_generator_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*IsPrettyIDReq); i {
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
		file_svc_infra_generator_generator_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*IsPrettyIDResp); i {
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
			RawDescriptor: file_svc_infra_generator_generator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_infra_generator_generator_proto_goTypes,
		DependencyIndexes: file_svc_infra_generator_generator_proto_depIdxs,
		MessageInfos:      file_svc_infra_generator_generator_proto_msgTypes,
	}.Build()
	File_svc_infra_generator_generator_proto = out.File
	file_svc_infra_generator_generator_proto_rawDesc = nil
	file_svc_infra_generator_generator_proto_goTypes = nil
	file_svc_infra_generator_generator_proto_depIdxs = nil
}
