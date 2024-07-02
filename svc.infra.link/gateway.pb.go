// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: svc.infra.link/gateway.proto

package link

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

type Gateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                           // 网关ID
	AdvertiseAddress  string `protobuf:"bytes,2,opt,name=advertise_address,json=advertiseAddress,proto3" json:"advertise_address,omitempty"`       // 网关地址
	CurrentSessions   int64  `protobuf:"varint,3,opt,name=current_sessions,json=currentSessions,proto3" json:"current_sessions,omitempty"`         // 当前连接数
	CurrentGroups     int64  `protobuf:"varint,4,opt,name=current_groups,json=currentGroups,proto3" json:"current_groups,omitempty"`               // 当前群组数（包括空群组）
	CurrentAccounts   int64  `protobuf:"varint,5,opt,name=current_accounts,json=currentAccounts,proto3" json:"current_accounts,omitempty"`         // 当前账号数
	CurrentDevices    int64  `protobuf:"varint,6,opt,name=current_devices,json=currentDevices,proto3" json:"current_devices,omitempty"`            // 当前设备数
	TotalUpwardMsgs   int64  `protobuf:"varint,7,opt,name=total_upward_msgs,json=totalUpwardMsgs,proto3" json:"total_upward_msgs,omitempty"`       // 总上行消息条数
	TotalDownwardMsgs int64  `protobuf:"varint,8,opt,name=total_downward_msgs,json=totalDownwardMsgs,proto3" json:"total_downward_msgs,omitempty"` // 总下行消息条数
	TotalConnects     int64  `protobuf:"varint,9,opt,name=total_connects,json=totalConnects,proto3" json:"total_connects,omitempty"`               // 总创建连接次数
	TotalRegisters    int64  `protobuf:"varint,10,opt,name=total_registers,json=totalRegisters,proto3" json:"total_registers,omitempty"`           // 总注册成功次数
	TotalBinds        int64  `protobuf:"varint,11,opt,name=total_binds,json=totalBinds,proto3" json:"total_binds,omitempty"`                       // 总绑定群组次数
	MemAlloc          uint64 `protobuf:"varint,12,opt,name=mem_alloc,json=memAlloc,proto3" json:"mem_alloc,omitempty"`                             // 进程当前分配内存
	TotalMemAlloc     uint64 `protobuf:"varint,13,opt,name=total_mem_alloc,json=totalMemAlloc,proto3" json:"total_mem_alloc,omitempty"`            // 进程总分配内存
	Allocs            uint64 `protobuf:"varint,14,opt,name=allocs,proto3" json:"allocs,omitempty"`                                                 // 进程分配内存次数
	Frees             uint64 `protobuf:"varint,15,opt,name=frees,proto3" json:"frees,omitempty"`                                                   // 进程释放内存次数
	NumCpus           int32  `protobuf:"varint,16,opt,name=num_cpus,json=numCpus,proto3" json:"num_cpus,omitempty"`                                // 当前逻辑CPU个数
	NumGoroutines     int32  `protobuf:"varint,17,opt,name=num_goroutines,json=numGoroutines,proto3" json:"num_goroutines,omitempty"`              // 当前Goroutine数
	CpuIdle           uint64 `protobuf:"varint,18,opt,name=cpu_idle,json=cpuIdle,proto3" json:"cpu_idle,omitempty"`                                // CPU空闲时间片
	CpuNoIdle         uint64 `protobuf:"varint,19,opt,name=cpu_no_idle,json=cpuNoIdle,proto3" json:"cpu_no_idle,omitempty"`                        // CPU工作时间片
	MemTotal          uint64 `protobuf:"varint,20,opt,name=mem_total,json=memTotal,proto3" json:"mem_total,omitempty"`                             // 总内存数
	MemAvailable      uint64 `protobuf:"varint,21,opt,name=mem_available,json=memAvailable,proto3" json:"mem_available,omitempty"`                 // 可用内存数
	Svc               string `protobuf:"bytes,127,opt,name=svc,proto3" json:"svc,omitempty"`                                                       // 网关类型
}

func (x *Gateway) Reset() {
	*x = Gateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gateway) ProtoMessage() {}

func (x *Gateway) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gateway.ProtoReflect.Descriptor instead.
func (*Gateway) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *Gateway) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Gateway) GetAdvertiseAddress() string {
	if x != nil {
		return x.AdvertiseAddress
	}
	return ""
}

func (x *Gateway) GetCurrentSessions() int64 {
	if x != nil {
		return x.CurrentSessions
	}
	return 0
}

func (x *Gateway) GetCurrentGroups() int64 {
	if x != nil {
		return x.CurrentGroups
	}
	return 0
}

func (x *Gateway) GetCurrentAccounts() int64 {
	if x != nil {
		return x.CurrentAccounts
	}
	return 0
}

func (x *Gateway) GetCurrentDevices() int64 {
	if x != nil {
		return x.CurrentDevices
	}
	return 0
}

func (x *Gateway) GetTotalUpwardMsgs() int64 {
	if x != nil {
		return x.TotalUpwardMsgs
	}
	return 0
}

func (x *Gateway) GetTotalDownwardMsgs() int64 {
	if x != nil {
		return x.TotalDownwardMsgs
	}
	return 0
}

func (x *Gateway) GetTotalConnects() int64 {
	if x != nil {
		return x.TotalConnects
	}
	return 0
}

func (x *Gateway) GetTotalRegisters() int64 {
	if x != nil {
		return x.TotalRegisters
	}
	return 0
}

func (x *Gateway) GetTotalBinds() int64 {
	if x != nil {
		return x.TotalBinds
	}
	return 0
}

func (x *Gateway) GetMemAlloc() uint64 {
	if x != nil {
		return x.MemAlloc
	}
	return 0
}

func (x *Gateway) GetTotalMemAlloc() uint64 {
	if x != nil {
		return x.TotalMemAlloc
	}
	return 0
}

func (x *Gateway) GetAllocs() uint64 {
	if x != nil {
		return x.Allocs
	}
	return 0
}

func (x *Gateway) GetFrees() uint64 {
	if x != nil {
		return x.Frees
	}
	return 0
}

func (x *Gateway) GetNumCpus() int32 {
	if x != nil {
		return x.NumCpus
	}
	return 0
}

func (x *Gateway) GetNumGoroutines() int32 {
	if x != nil {
		return x.NumGoroutines
	}
	return 0
}

func (x *Gateway) GetCpuIdle() uint64 {
	if x != nil {
		return x.CpuIdle
	}
	return 0
}

func (x *Gateway) GetCpuNoIdle() uint64 {
	if x != nil {
		return x.CpuNoIdle
	}
	return 0
}

func (x *Gateway) GetMemTotal() uint64 {
	if x != nil {
		return x.MemTotal
	}
	return 0
}

func (x *Gateway) GetMemAvailable() uint64 {
	if x != nil {
		return x.MemAvailable
	}
	return 0
}

func (x *Gateway) GetSvc() string {
	if x != nil {
		return x.Svc
	}
	return ""
}

type SelectSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                     // 网关ID
	AdvertiseAddress string `protobuf:"bytes,2,opt,name=advertise_address,json=advertiseAddress,proto3" json:"advertise_address,omitempty"` // 网关地址
	Sign             string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`                                                 // 签名
	Svc              string `protobuf:"bytes,127,opt,name=svc,proto3" json:"svc,omitempty"`                                                 // 网关类型
}

func (x *SelectSummary) Reset() {
	*x = SelectSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectSummary) ProtoMessage() {}

func (x *SelectSummary) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectSummary.ProtoReflect.Descriptor instead.
func (*SelectSummary) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *SelectSummary) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SelectSummary) GetAdvertiseAddress() string {
	if x != nil {
		return x.AdvertiseAddress
	}
	return ""
}

func (x *SelectSummary) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *SelectSummary) GetSvc() string {
	if x != nil {
		return x.Svc
	}
	return ""
}

type ListGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Svc string `protobuf:"bytes,127,opt,name=svc,proto3" json:"svc,omitempty"` // 网关类型（可留空）
}

func (x *ListGatewayRequest) Reset() {
	*x = ListGatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGatewayRequest) ProtoMessage() {}

func (x *ListGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGatewayRequest.ProtoReflect.Descriptor instead.
func (*ListGatewayRequest) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *ListGatewayRequest) GetSvc() string {
	if x != nil {
		return x.Svc
	}
	return ""
}

type ListGatewayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gateways []*Gateway `protobuf:"bytes,1,rep,name=gateways,proto3" json:"gateways,omitempty"` // 网关列表
}

func (x *ListGatewayResponse) Reset() {
	*x = ListGatewayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGatewayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGatewayResponse) ProtoMessage() {}

func (x *ListGatewayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGatewayResponse.ProtoReflect.Descriptor instead.
func (*ListGatewayResponse) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *ListGatewayResponse) GetGateways() []*Gateway {
	if x != nil {
		return x.Gateways
	}
	return nil
}

type GetGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // 网关ID
}

func (x *GetGatewayRequest) Reset() {
	*x = GetGatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGatewayRequest) ProtoMessage() {}

func (x *GetGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGatewayRequest.ProtoReflect.Descriptor instead.
func (*GetGatewayRequest) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *GetGatewayRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetGatewayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gateway *Gateway `protobuf:"bytes,1,opt,name=gateway,proto3" json:"gateway,omitempty"` // 网关信息
}

func (x *GetGatewayResponse) Reset() {
	*x = GetGatewayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGatewayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGatewayResponse) ProtoMessage() {}

func (x *GetGatewayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGatewayResponse.ProtoReflect.Descriptor instead.
func (*GetGatewayResponse) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *GetGatewayResponse) GetGateway() *Gateway {
	if x != nil {
		return x.Gateway
	}
	return nil
}

type SelectGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` // 账号ID
	Platform  string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`                    // 平台
	Version   string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`                      // 客户端版本
	Device    string `protobuf:"bytes,4,opt,name=device,proto3" json:"device,omitempty"`                        // 设备标识（指纹）
	Svc       string `protobuf:"bytes,127,opt,name=svc,proto3" json:"svc,omitempty"`                            // 网关类型（默认ws）
	Total     int64  `protobuf:"varint,128,opt,name=total,proto3" json:"total,omitempty"`                       // 希望获取网关总数（默认5）
}

func (x *SelectGatewayRequest) Reset() {
	*x = SelectGatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectGatewayRequest) ProtoMessage() {}

func (x *SelectGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectGatewayRequest.ProtoReflect.Descriptor instead.
func (*SelectGatewayRequest) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *SelectGatewayRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *SelectGatewayRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *SelectGatewayRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SelectGatewayRequest) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *SelectGatewayRequest) GetSvc() string {
	if x != nil {
		return x.Svc
	}
	return ""
}

func (x *SelectGatewayRequest) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SelectGatewayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gateways []*SelectSummary `protobuf:"bytes,1,rep,name=gateways,proto3" json:"gateways,omitempty"` // 网关列表
}

func (x *SelectGatewayResponse) Reset() {
	*x = SelectGatewayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_gateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectGatewayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectGatewayResponse) ProtoMessage() {}

func (x *SelectGatewayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_gateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectGatewayResponse.ProtoReflect.Descriptor instead.
func (*SelectGatewayResponse) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_gateway_proto_rawDescGZIP(), []int{7}
}

func (x *SelectGatewayResponse) GetGateways() []*SelectSummary {
	if x != nil {
		return x.Gateways
	}
	return nil
}

var File_svc_infra_link_gateway_proto protoreflect.FileDescriptor

var file_svc_infra_link_gateway_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xfd,
	0x05, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2a, 0x0a,
	0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x70, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x73,
	0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55,
	0x70, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x6f, 0x77,
	0x6e, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x69, 0x6e, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6d, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x6d, 0x65, 0x6d, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x65, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x65, 0x65, 0x73,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x66, 0x72, 0x65, 0x65, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x70, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x6e, 0x75, 0x6d, 0x43, 0x70, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x5f,
	0x67, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x47, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x63, 0x70, 0x75, 0x49, 0x64, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x63, 0x70,
	0x75, 0x5f, 0x6e, 0x6f, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x63, 0x70, 0x75, 0x4e, 0x6f, 0x49, 0x64, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x6d, 0x5f, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x6d, 0x65, 0x6d, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x76, 0x63, 0x18, 0x7f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x76, 0x63, 0x22, 0x73,
	0x0a, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2b, 0x0a, 0x11, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x76, 0x63, 0x18, 0x7f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x76, 0x63, 0x22, 0x26, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x76, 0x63,
	0x18, 0x7f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x76, 0x63, 0x22, 0x4a, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x08, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x07, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0xac, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x76, 0x63, 0x18, 0x7f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x76, 0x63, 0x12, 0x15, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x80, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0x53, 0x0a, 0x15, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x32, 0x89, 0x02, 0x0a, 0x0b, 0x4c, 0x69,
	0x6e, 0x6b, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x51, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x22, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x06,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x3b, 0x6c, 0x69, 0x6e, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_infra_link_gateway_proto_rawDescOnce sync.Once
	file_svc_infra_link_gateway_proto_rawDescData = file_svc_infra_link_gateway_proto_rawDesc
)

func file_svc_infra_link_gateway_proto_rawDescGZIP() []byte {
	file_svc_infra_link_gateway_proto_rawDescOnce.Do(func() {
		file_svc_infra_link_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_infra_link_gateway_proto_rawDescData)
	})
	return file_svc_infra_link_gateway_proto_rawDescData
}

var file_svc_infra_link_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_svc_infra_link_gateway_proto_goTypes = []any{
	(*Gateway)(nil),               // 0: svc.infra.link.gateway
	(*SelectSummary)(nil),         // 1: svc.infra.link.select_summary
	(*ListGatewayRequest)(nil),    // 2: svc.infra.link.ListGatewayRequest
	(*ListGatewayResponse)(nil),   // 3: svc.infra.link.ListGatewayResponse
	(*GetGatewayRequest)(nil),     // 4: svc.infra.link.GetGatewayRequest
	(*GetGatewayResponse)(nil),    // 5: svc.infra.link.GetGatewayResponse
	(*SelectGatewayRequest)(nil),  // 6: svc.infra.link.SelectGatewayRequest
	(*SelectGatewayResponse)(nil), // 7: svc.infra.link.SelectGatewayResponse
}
var file_svc_infra_link_gateway_proto_depIdxs = []int32{
	0, // 0: svc.infra.link.ListGatewayResponse.gateways:type_name -> svc.infra.link.gateway
	0, // 1: svc.infra.link.GetGatewayResponse.gateway:type_name -> svc.infra.link.gateway
	1, // 2: svc.infra.link.SelectGatewayResponse.gateways:type_name -> svc.infra.link.select_summary
	2, // 3: svc.infra.link.LinkGateway.List:input_type -> svc.infra.link.ListGatewayRequest
	4, // 4: svc.infra.link.LinkGateway.Get:input_type -> svc.infra.link.GetGatewayRequest
	6, // 5: svc.infra.link.LinkGateway.Select:input_type -> svc.infra.link.SelectGatewayRequest
	3, // 6: svc.infra.link.LinkGateway.List:output_type -> svc.infra.link.ListGatewayResponse
	5, // 7: svc.infra.link.LinkGateway.Get:output_type -> svc.infra.link.GetGatewayResponse
	7, // 8: svc.infra.link.LinkGateway.Select:output_type -> svc.infra.link.SelectGatewayResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_svc_infra_link_gateway_proto_init() }
func file_svc_infra_link_gateway_proto_init() {
	if File_svc_infra_link_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_infra_link_gateway_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Gateway); i {
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
		file_svc_infra_link_gateway_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SelectSummary); i {
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
		file_svc_infra_link_gateway_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListGatewayRequest); i {
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
		file_svc_infra_link_gateway_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListGatewayResponse); i {
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
		file_svc_infra_link_gateway_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetGatewayRequest); i {
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
		file_svc_infra_link_gateway_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetGatewayResponse); i {
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
		file_svc_infra_link_gateway_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SelectGatewayRequest); i {
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
		file_svc_infra_link_gateway_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SelectGatewayResponse); i {
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
			RawDescriptor: file_svc_infra_link_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_infra_link_gateway_proto_goTypes,
		DependencyIndexes: file_svc_infra_link_gateway_proto_depIdxs,
		MessageInfos:      file_svc_infra_link_gateway_proto_msgTypes,
	}.Build()
	File_svc_infra_link_gateway_proto = out.File
	file_svc_infra_link_gateway_proto_rawDesc = nil
	file_svc_infra_link_gateway_proto_goTypes = nil
	file_svc_infra_link_gateway_proto_depIdxs = nil
}
