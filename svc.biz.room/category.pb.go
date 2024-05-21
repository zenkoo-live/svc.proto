// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: svc.biz.room/category.proto

package room

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// CategoryInfo 分类详情
type CategoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId   int64                  `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ParentId     int64                  `protobuf:"varint,4,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`            // 父级ID
	CategoryCode string                 `protobuf:"bytes,5,opt,name=category_code,json=categoryCode,proto3" json:"category_code,omitempty"` // 代号（唯一，预留）
	CategoryName string                 `protobuf:"bytes,6,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"` // 名称
	Sort         int32                  `protobuf:"varint,7,opt,name=sort,proto3" json:"sort,omitempty"`                                    // 排序
	Childrens    []*CategoryInfo        `protobuf:"bytes,8,rep,name=childrens,proto3" json:"childrens,omitempty"`
}

func (x *CategoryInfo) Reset() {
	*x = CategoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryInfo) ProtoMessage() {}

func (x *CategoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryInfo.ProtoReflect.Descriptor instead.
func (*CategoryInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryInfo) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CategoryInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CategoryInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CategoryInfo) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CategoryInfo) GetCategoryCode() string {
	if x != nil {
		return x.CategoryCode
	}
	return ""
}

func (x *CategoryInfo) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *CategoryInfo) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *CategoryInfo) GetChildrens() []*CategoryInfo {
	if x != nil {
		return x.Childrens
	}
	return nil
}

type CreateCategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *CategoryInfo `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CreateCategoryReq) Reset() {
	*x = CreateCategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryReq) ProtoMessage() {}

func (x *CreateCategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryReq.ProtoReflect.Descriptor instead.
func (*CreateCategoryReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCategoryReq) GetCategory() *CategoryInfo {
	if x != nil {
		return x.Category
	}
	return nil
}

type CreateCategoryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *CategoryInfo `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CreateCategoryResp) Reset() {
	*x = CreateCategoryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryResp) ProtoMessage() {}

func (x *CreateCategoryResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryResp.ProtoReflect.Descriptor instead.
func (*CreateCategoryResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCategoryResp) GetCategory() *CategoryInfo {
	if x != nil {
		return x.Category
	}
	return nil
}

type UpdateCategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId int64                  `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Category   *CategoryInfo          `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateCategoryReq) Reset() {
	*x = UpdateCategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryReq) ProtoMessage() {}

func (x *UpdateCategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryReq.ProtoReflect.Descriptor instead.
func (*UpdateCategoryReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCategoryReq) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *UpdateCategoryReq) GetCategory() *CategoryInfo {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *UpdateCategoryReq) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type GetCategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId int64 `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *GetCategoryReq) Reset() {
	*x = GetCategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryReq) ProtoMessage() {}

func (x *GetCategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryReq.ProtoReflect.Descriptor instead.
func (*GetCategoryReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{4}
}

func (x *GetCategoryReq) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

type GetCategoryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *CategoryInfo `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GetCategoryResp) Reset() {
	*x = GetCategoryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryResp) ProtoMessage() {}

func (x *GetCategoryResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryResp.ProtoReflect.Descriptor instead.
func (*GetCategoryResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{5}
}

func (x *GetCategoryResp) GetCategory() *CategoryInfo {
	if x != nil {
		return x.Category
	}
	return nil
}

type DeleteCategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId int64 `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *DeleteCategoryReq) Reset() {
	*x = DeleteCategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryReq) ProtoMessage() {}

func (x *DeleteCategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryReq.ProtoReflect.Descriptor instead.
func (*DeleteCategoryReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCategoryReq) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

type ListCategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId    int64 `protobuf:"varint,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`          // 父级ID，跟级分类传1
	ReturnCount bool  `protobuf:"varint,2,opt,name=return_count,json=returnCount,proto3" json:"return_count,omitempty"` // 是否返回总数
}

func (x *ListCategoryReq) Reset() {
	*x = ListCategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryReq) ProtoMessage() {}

func (x *ListCategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryReq.ProtoReflect.Descriptor instead.
func (*ListCategoryReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{7}
}

func (x *ListCategoryReq) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *ListCategoryReq) GetReturnCount() bool {
	if x != nil {
		return x.ReturnCount
	}
	return false
}

type ListCategoryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CategoryInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListCategoryResp) Reset() {
	*x = ListCategoryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryResp) ProtoMessage() {}

func (x *ListCategoryResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryResp.ProtoReflect.Descriptor instead.
func (*ListCategoryResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{8}
}

func (x *ListCategoryResp) GetItems() []*CategoryInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

// 获取全部板块分类（分类及子分类树结构）
type ListCategoryTreeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCategoryTreeReq) Reset() {
	*x = ListCategoryTreeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoryTreeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryTreeReq) ProtoMessage() {}

func (x *ListCategoryTreeReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryTreeReq.ProtoReflect.Descriptor instead.
func (*ListCategoryTreeReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{9}
}

type ListCategoryTreeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CategoryInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListCategoryTreeResp) Reset() {
	*x = ListCategoryTreeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_room_category_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoryTreeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryTreeResp) ProtoMessage() {}

func (x *ListCategoryTreeResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_category_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryTreeResp.ProtoReflect.Descriptor instead.
func (*ListCategoryTreeResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_category_proto_rawDescGZIP(), []int{10}
}

func (x *ListCategoryTreeResp) GetItems() []*CategoryInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_svc_biz_room_category_proto protoreflect.FileDescriptor

var file_svc_biz_room_category_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73,
	0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x02, 0x0a, 0x0c,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x73, 0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x36, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x22, 0xa9, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22,
	0x31, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x22, 0x49, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x34, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x15, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x65, 0x65,
	0x52, 0x65, 0x71, 0x22, 0x48, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xf7, 0x03,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x1f, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x65, 0x65, 0x12, 0x21,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x22, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x65,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x3b, 0x72, 0x6f, 0x6f, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_room_category_proto_rawDescOnce sync.Once
	file_svc_biz_room_category_proto_rawDescData = file_svc_biz_room_category_proto_rawDesc
)

func file_svc_biz_room_category_proto_rawDescGZIP() []byte {
	file_svc_biz_room_category_proto_rawDescOnce.Do(func() {
		file_svc_biz_room_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_room_category_proto_rawDescData)
	})
	return file_svc_biz_room_category_proto_rawDescData
}

var file_svc_biz_room_category_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_svc_biz_room_category_proto_goTypes = []interface{}{
	(*CategoryInfo)(nil),          // 0: svc.biz.room.CategoryInfo
	(*CreateCategoryReq)(nil),     // 1: svc.biz.room.CreateCategoryReq
	(*CreateCategoryResp)(nil),    // 2: svc.biz.room.CreateCategoryResp
	(*UpdateCategoryReq)(nil),     // 3: svc.biz.room.UpdateCategoryReq
	(*GetCategoryReq)(nil),        // 4: svc.biz.room.GetCategoryReq
	(*GetCategoryResp)(nil),       // 5: svc.biz.room.GetCategoryResp
	(*DeleteCategoryReq)(nil),     // 6: svc.biz.room.DeleteCategoryReq
	(*ListCategoryReq)(nil),       // 7: svc.biz.room.ListCategoryReq
	(*ListCategoryResp)(nil),      // 8: svc.biz.room.ListCategoryResp
	(*ListCategoryTreeReq)(nil),   // 9: svc.biz.room.ListCategoryTreeReq
	(*ListCategoryTreeResp)(nil),  // 10: svc.biz.room.ListCategoryTreeResp
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 12: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),         // 13: google.protobuf.Empty
}
var file_svc_biz_room_category_proto_depIdxs = []int32{
	11, // 0: svc.biz.room.CategoryInfo.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: svc.biz.room.CategoryInfo.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: svc.biz.room.CategoryInfo.childrens:type_name -> svc.biz.room.CategoryInfo
	0,  // 3: svc.biz.room.CreateCategoryReq.category:type_name -> svc.biz.room.CategoryInfo
	0,  // 4: svc.biz.room.CreateCategoryResp.category:type_name -> svc.biz.room.CategoryInfo
	0,  // 5: svc.biz.room.UpdateCategoryReq.category:type_name -> svc.biz.room.CategoryInfo
	12, // 6: svc.biz.room.UpdateCategoryReq.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 7: svc.biz.room.GetCategoryResp.category:type_name -> svc.biz.room.CategoryInfo
	0,  // 8: svc.biz.room.ListCategoryResp.items:type_name -> svc.biz.room.CategoryInfo
	0,  // 9: svc.biz.room.ListCategoryTreeResp.items:type_name -> svc.biz.room.CategoryInfo
	4,  // 10: svc.biz.room.Category.GetCategory:input_type -> svc.biz.room.GetCategoryReq
	1,  // 11: svc.biz.room.Category.CreateCategory:input_type -> svc.biz.room.CreateCategoryReq
	3,  // 12: svc.biz.room.Category.UpdateCategory:input_type -> svc.biz.room.UpdateCategoryReq
	6,  // 13: svc.biz.room.Category.DeleteCategory:input_type -> svc.biz.room.DeleteCategoryReq
	7,  // 14: svc.biz.room.Category.ListCategory:input_type -> svc.biz.room.ListCategoryReq
	9,  // 15: svc.biz.room.Category.ListCategoryTree:input_type -> svc.biz.room.ListCategoryTreeReq
	5,  // 16: svc.biz.room.Category.GetCategory:output_type -> svc.biz.room.GetCategoryResp
	2,  // 17: svc.biz.room.Category.CreateCategory:output_type -> svc.biz.room.CreateCategoryResp
	13, // 18: svc.biz.room.Category.UpdateCategory:output_type -> google.protobuf.Empty
	13, // 19: svc.biz.room.Category.DeleteCategory:output_type -> google.protobuf.Empty
	8,  // 20: svc.biz.room.Category.ListCategory:output_type -> svc.biz.room.ListCategoryResp
	10, // 21: svc.biz.room.Category.ListCategoryTree:output_type -> svc.biz.room.ListCategoryTreeResp
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_svc_biz_room_category_proto_init() }
func file_svc_biz_room_category_proto_init() {
	if File_svc_biz_room_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_biz_room_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryInfo); i {
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
		file_svc_biz_room_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryReq); i {
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
		file_svc_biz_room_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryResp); i {
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
		file_svc_biz_room_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCategoryReq); i {
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
		file_svc_biz_room_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryReq); i {
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
		file_svc_biz_room_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryResp); i {
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
		file_svc_biz_room_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCategoryReq); i {
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
		file_svc_biz_room_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoryReq); i {
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
		file_svc_biz_room_category_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoryResp); i {
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
		file_svc_biz_room_category_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoryTreeReq); i {
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
		file_svc_biz_room_category_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoryTreeResp); i {
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
			RawDescriptor: file_svc_biz_room_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_room_category_proto_goTypes,
		DependencyIndexes: file_svc_biz_room_category_proto_depIdxs,
		MessageInfos:      file_svc_biz_room_category_proto_msgTypes,
	}.Build()
	File_svc_biz_room_category_proto = out.File
	file_svc_biz_room_category_proto_rawDesc = nil
	file_svc_biz_room_category_proto_goTypes = nil
	file_svc_biz_room_category_proto_depIdxs = nil
}
