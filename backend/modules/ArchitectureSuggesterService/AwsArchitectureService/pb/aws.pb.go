// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: AwsArchitectureService.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetAwsMachineTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceType string `protobuf:"bytes,1,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	//global parent : like Ecs , Eks
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *GetAwsMachineTypesRequest) Reset() {
	*x = GetAwsMachineTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAwsMachineTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAwsMachineTypesRequest) ProtoMessage() {}

func (x *GetAwsMachineTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAwsMachineTypesRequest.ProtoReflect.Descriptor instead.
func (*GetAwsMachineTypesRequest) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{0}
}

func (x *GetAwsMachineTypesRequest) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *GetAwsMachineTypesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

type GetAwsInstanceTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent     string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	LaunchType string `protobuf:"bytes,1,opt,name=launch_type,json=launchType,proto3" json:"launch_type,omitempty"`
}

func (x *GetAwsInstanceTypesRequest) Reset() {
	*x = GetAwsInstanceTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAwsInstanceTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAwsInstanceTypesRequest) ProtoMessage() {}

func (x *GetAwsInstanceTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAwsInstanceTypesRequest.ProtoReflect.Descriptor instead.
func (*GetAwsInstanceTypesRequest) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{1}
}

func (x *GetAwsInstanceTypesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *GetAwsInstanceTypesRequest) GetLaunchType() string {
	if x != nil {
		return x.LaunchType
	}
	return ""
}

type GetAwsMachineTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineTypes  []string         `protobuf:"bytes,1,rep,name=machine_types,json=machineTypes,proto3" json:"machine_types,omitempty"`
	MachineConfig []*MachineConfig `protobuf:"bytes,2,rep,name=machine_config,json=machineConfig,proto3" json:"machine_config,omitempty"`
}

func (x *GetAwsMachineTypesResponse) Reset() {
	*x = GetAwsMachineTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAwsMachineTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAwsMachineTypesResponse) ProtoMessage() {}

func (x *GetAwsMachineTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAwsMachineTypesResponse.ProtoReflect.Descriptor instead.
func (*GetAwsMachineTypesResponse) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{2}
}

func (x *GetAwsMachineTypesResponse) GetMachineTypes() []string {
	if x != nil {
		return x.MachineTypes
	}
	return nil
}

func (x *GetAwsMachineTypesResponse) GetMachineConfig() []*MachineConfig {
	if x != nil {
		return x.MachineConfig
	}
	return nil
}

type GetAwsInstanceTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceTypes []string `protobuf:"bytes,1,rep,name=instance_types,json=instanceTypes,proto3" json:"instance_types,omitempty"`
}

func (x *GetAwsInstanceTypesResponse) Reset() {
	*x = GetAwsInstanceTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAwsInstanceTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAwsInstanceTypesResponse) ProtoMessage() {}

func (x *GetAwsInstanceTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAwsInstanceTypesResponse.ProtoReflect.Descriptor instead.
func (*GetAwsInstanceTypesResponse) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{3}
}

func (x *GetAwsInstanceTypesResponse) GetInstanceTypes() []string {
	if x != nil {
		return x.InstanceTypes
	}
	return nil
}

type GetAwsLaunchTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerService string `protobuf:"bytes,1,opt,name=container_service,json=containerService,proto3" json:"container_service,omitempty"`
}

func (x *GetAwsLaunchTypesRequest) Reset() {
	*x = GetAwsLaunchTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAwsLaunchTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAwsLaunchTypesRequest) ProtoMessage() {}

func (x *GetAwsLaunchTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAwsLaunchTypesRequest.ProtoReflect.Descriptor instead.
func (*GetAwsLaunchTypesRequest) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{4}
}

func (x *GetAwsLaunchTypesRequest) GetContainerService() string {
	if x != nil {
		return x.ContainerService
	}
	return ""
}

type GetAwsLaunchTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaunchType []string `protobuf:"bytes,1,rep,name=launch_type,json=launchType,proto3" json:"launch_type,omitempty"`
}

func (x *GetAwsLaunchTypesResponse) Reset() {
	*x = GetAwsLaunchTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAwsLaunchTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAwsLaunchTypesResponse) ProtoMessage() {}

func (x *GetAwsLaunchTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAwsLaunchTypesResponse.ProtoReflect.Descriptor instead.
func (*GetAwsLaunchTypesResponse) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{5}
}

func (x *GetAwsLaunchTypesResponse) GetLaunchType() []string {
	if x != nil {
		return x.LaunchType
	}
	return nil
}

type GetAwsContainerServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nothing string `protobuf:"bytes,1,opt,name=nothing,proto3" json:"nothing,omitempty"`
}

func (x *GetAwsContainerServicesRequest) Reset() {
	*x = GetAwsContainerServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAwsContainerServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAwsContainerServicesRequest) ProtoMessage() {}

func (x *GetAwsContainerServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAwsContainerServicesRequest.ProtoReflect.Descriptor instead.
func (*GetAwsContainerServicesRequest) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{6}
}

func (x *GetAwsContainerServicesRequest) GetNothing() string {
	if x != nil {
		return x.Nothing
	}
	return ""
}

type GetAwsContainerServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmazonContainerServices []string `protobuf:"bytes,1,rep,name=amazon_container_services,json=amazonContainerServices,proto3" json:"amazon_container_services,omitempty"`
}

func (x *GetAwsContainerServicesResponse) Reset() {
	*x = GetAwsContainerServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAwsContainerServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAwsContainerServicesResponse) ProtoMessage() {}

func (x *GetAwsContainerServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAwsContainerServicesResponse.ProtoReflect.Descriptor instead.
func (*GetAwsContainerServicesResponse) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{7}
}

func (x *GetAwsContainerServicesResponse) GetAmazonContainerServices() []string {
	if x != nil {
		return x.AmazonContainerServices
	}
	return nil
}

type MachineConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent                    string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	InstanceSize              string `protobuf:"bytes,2,opt,name=instance_size,json=instanceSize,proto3" json:"instance_size,omitempty"`
	VCpu                      string `protobuf:"bytes,3,opt,name=v_cpu,json=vCpu,proto3" json:"v_cpu,omitempty"`
	MemoryInGib               string `protobuf:"bytes,4,opt,name=memory_in_gib,json=memoryInGib,proto3" json:"memory_in_gib,omitempty"`
	InstanceStorage           string `protobuf:"bytes,5,opt,name=instance_storage,json=instanceStorage,proto3" json:"instance_storage,omitempty"`
	InstanceStorageGib        string `protobuf:"bytes,18,opt,name=instance_storage_gib,json=instanceStorageGib,proto3" json:"instance_storage_gib,omitempty"`
	NetworkBandwidthGbpc      string `protobuf:"bytes,6,opt,name=network_bandwidth_gbpc,json=networkBandwidthGbpc,proto3" json:"network_bandwidth_gbpc,omitempty"`
	EbsBandwidthMbps          string `protobuf:"bytes,7,opt,name=ebs_bandwidth_mbps,json=ebsBandwidthMbps,proto3" json:"ebs_bandwidth_mbps,omitempty"`
	BaselinePerformance       string `protobuf:"bytes,8,opt,name=baseline_performance,json=baselinePerformance,proto3" json:"baseline_performance,omitempty"`
	CpuCreditsEarnedHr        string `protobuf:"bytes,9,opt,name=cpu_credits_earned_hr,json=cpuCreditsEarnedHr,proto3" json:"cpu_credits_earned_hr,omitempty"`
	NetworkBurstBandwidthGbps string `protobuf:"bytes,10,opt,name=network_burst_bandwidth_gbps,json=networkBurstBandwidthGbps,proto3" json:"network_burst_bandwidth_gbps,omitempty"`
	EbsBurstBandwidthMbps     string `protobuf:"bytes,11,opt,name=ebs_burst_bandwidth_mbps,json=ebsBurstBandwidthMbps,proto3" json:"ebs_burst_bandwidth_mbps,omitempty"`
	Mem                       string `protobuf:"bytes,12,opt,name=mem,proto3" json:"mem,omitempty"`
	MemGib                    string `protobuf:"bytes,19,opt,name=mem_gib,json=memGib,proto3" json:"mem_gib,omitempty"`
	NetworkPerformance        string `protobuf:"bytes,13,opt,name=network_performance,json=networkPerformance,proto3" json:"network_performance,omitempty"`
	Storage                   string `protobuf:"bytes,14,opt,name=storage,proto3" json:"storage,omitempty"`
	StorageGb                 string `protobuf:"bytes,23,opt,name=storage_gb,json=storageGb,proto3" json:"storage_gb,omitempty"`
	NetworkPerformanceGib     string `protobuf:"bytes,15,opt,name=network_performance_gib,json=networkPerformanceGib,proto3" json:"network_performance_gib,omitempty"`
	NetworkBandwidth          string `protobuf:"bytes,16,opt,name=network_bandwidth,json=networkBandwidth,proto3" json:"network_bandwidth,omitempty"`
	EbsBandwidth              string `protobuf:"bytes,17,opt,name=ebs_bandwidth,json=ebsBandwidth,proto3" json:"ebs_bandwidth,omitempty"`
	GpuMemGib                 string `protobuf:"bytes,20,opt,name=gpu_mem_gib,json=gpuMemGib,proto3" json:"gpu_mem_gib,omitempty"`
	GpuMem                    string `protobuf:"bytes,21,opt,name=gpu_mem,json=gpuMem,proto3" json:"gpu_mem,omitempty"`
	GpuP2P                    string `protobuf:"bytes,22,opt,name=gpu_p2p,json=gpuP2p,proto3" json:"gpu_p2p,omitempty"`
	DedicatedEbsBandwidth     string `protobuf:"bytes,24,opt,name=dedicated_ebs_bandwidth,json=dedicatedEbsBandwidth,proto3" json:"dedicated_ebs_bandwidth,omitempty"`
	LocalStorage              string `protobuf:"bytes,25,opt,name=local_storage,json=localStorage,proto3" json:"local_storage,omitempty"`
	LocalStorageGb            string `protobuf:"bytes,26,opt,name=local_storage_gb,json=localStorageGb,proto3" json:"local_storage_gb,omitempty"`
}

func (x *MachineConfig) Reset() {
	*x = MachineConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineConfig) ProtoMessage() {}

func (x *MachineConfig) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineConfig.ProtoReflect.Descriptor instead.
func (*MachineConfig) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{8}
}

func (x *MachineConfig) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *MachineConfig) GetInstanceSize() string {
	if x != nil {
		return x.InstanceSize
	}
	return ""
}

func (x *MachineConfig) GetVCpu() string {
	if x != nil {
		return x.VCpu
	}
	return ""
}

func (x *MachineConfig) GetMemoryInGib() string {
	if x != nil {
		return x.MemoryInGib
	}
	return ""
}

func (x *MachineConfig) GetInstanceStorage() string {
	if x != nil {
		return x.InstanceStorage
	}
	return ""
}

func (x *MachineConfig) GetInstanceStorageGib() string {
	if x != nil {
		return x.InstanceStorageGib
	}
	return ""
}

func (x *MachineConfig) GetNetworkBandwidthGbpc() string {
	if x != nil {
		return x.NetworkBandwidthGbpc
	}
	return ""
}

func (x *MachineConfig) GetEbsBandwidthMbps() string {
	if x != nil {
		return x.EbsBandwidthMbps
	}
	return ""
}

func (x *MachineConfig) GetBaselinePerformance() string {
	if x != nil {
		return x.BaselinePerformance
	}
	return ""
}

func (x *MachineConfig) GetCpuCreditsEarnedHr() string {
	if x != nil {
		return x.CpuCreditsEarnedHr
	}
	return ""
}

func (x *MachineConfig) GetNetworkBurstBandwidthGbps() string {
	if x != nil {
		return x.NetworkBurstBandwidthGbps
	}
	return ""
}

func (x *MachineConfig) GetEbsBurstBandwidthMbps() string {
	if x != nil {
		return x.EbsBurstBandwidthMbps
	}
	return ""
}

func (x *MachineConfig) GetMem() string {
	if x != nil {
		return x.Mem
	}
	return ""
}

func (x *MachineConfig) GetMemGib() string {
	if x != nil {
		return x.MemGib
	}
	return ""
}

func (x *MachineConfig) GetNetworkPerformance() string {
	if x != nil {
		return x.NetworkPerformance
	}
	return ""
}

func (x *MachineConfig) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *MachineConfig) GetStorageGb() string {
	if x != nil {
		return x.StorageGb
	}
	return ""
}

func (x *MachineConfig) GetNetworkPerformanceGib() string {
	if x != nil {
		return x.NetworkPerformanceGib
	}
	return ""
}

func (x *MachineConfig) GetNetworkBandwidth() string {
	if x != nil {
		return x.NetworkBandwidth
	}
	return ""
}

func (x *MachineConfig) GetEbsBandwidth() string {
	if x != nil {
		return x.EbsBandwidth
	}
	return ""
}

func (x *MachineConfig) GetGpuMemGib() string {
	if x != nil {
		return x.GpuMemGib
	}
	return ""
}

func (x *MachineConfig) GetGpuMem() string {
	if x != nil {
		return x.GpuMem
	}
	return ""
}

func (x *MachineConfig) GetGpuP2P() string {
	if x != nil {
		return x.GpuP2P
	}
	return ""
}

func (x *MachineConfig) GetDedicatedEbsBandwidth() string {
	if x != nil {
		return x.DedicatedEbsBandwidth
	}
	return ""
}

func (x *MachineConfig) GetLocalStorage() string {
	if x != nil {
		return x.LocalStorage
	}
	return ""
}

func (x *MachineConfig) GetLocalStorageGb() string {
	if x != nil {
		return x.LocalStorageGb
	}
	return ""
}

var File_aws_proto protoreflect.FileDescriptor

var file_aws_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x55, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x7b, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0e, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x44, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x77,
	0x73, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x3c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x5d, 0x0a, 0x1f, 0x47, 0x65,
	0x74, 0x41, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x19, 0x61, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x17, 0x61, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x9e, 0x08, 0x0a, 0x0d, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x76, 0x5f, 0x63, 0x70,
	0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x43, 0x70, 0x75, 0x12, 0x22, 0x0a,
	0x0d, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x5f, 0x67, 0x69, 0x62, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x47, 0x69,
	0x62, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x14,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x67, 0x69, 0x62, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x47, 0x69, 0x62, 0x12, 0x34,
	0x0a, 0x16, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x5f, 0x67, 0x62, 0x70, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x47, 0x62, 0x70, 0x63, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x62, 0x73, 0x5f, 0x62, 0x61, 0x6e, 0x64,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6d, 0x62, 0x70, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x65, 0x62, 0x73, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x4d, 0x62,
	0x70, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x73, 0x5f, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x5f, 0x68, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x70, 0x75, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73,
	0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x48, 0x72, 0x12, 0x3f, 0x0a, 0x1c, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x62, 0x75, 0x72, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x5f, 0x67, 0x62, 0x70, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x75, 0x72, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x64,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x47, 0x62, 0x70, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x65, 0x62, 0x73,
	0x5f, 0x62, 0x75, 0x72, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x5f, 0x6d, 0x62, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x65, 0x62, 0x73,
	0x42, 0x75, 0x72, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x4d, 0x62,
	0x70, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x65, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x5f, 0x67, 0x69, 0x62, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x47, 0x69, 0x62, 0x12, 0x2f, 0x0a,
	0x13, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x67, 0x62, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x47, 0x62, 0x12, 0x36, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x67,
	0x69, 0x62, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x47, 0x69, 0x62, 0x12,
	0x2b, 0x0a, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x62, 0x73, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x62, 0x73, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x12, 0x1e, 0x0a, 0x0b, 0x67, 0x70, 0x75, 0x5f, 0x6d, 0x65, 0x6d, 0x5f, 0x67, 0x69, 0x62,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x70, 0x75, 0x4d, 0x65, 0x6d, 0x47, 0x69,
	0x62, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x70, 0x75, 0x5f, 0x6d, 0x65, 0x6d, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x70, 0x75, 0x4d, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x70,
	0x75, 0x5f, 0x70, 0x32, 0x70, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x70, 0x75,
	0x50, 0x32, 0x70, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x65, 0x62, 0x73, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x45,
	0x62, 0x73, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x67, 0x62, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x47, 0x62, 0x32, 0xa7, 0x05, 0x0a, 0x0a, 0x41,
	0x77, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xaa, 0x01, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x41, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x77,
	0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f,
	0x67, 0x65, 0x74, 0x2d, 0x61, 0x77, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xba, 0x43, 0x19, 0x12, 0x17, 0x67,
	0x65, 0x74, 0x41, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0xa0, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x77,
	0x73, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x32, 0x12, 0x30, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x61,
	0x77, 0x73, 0x2d, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x7b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x7d, 0xba, 0x43, 0x13, 0x12, 0x11, 0x67, 0x65, 0x74, 0x41, 0x77, 0x73, 0x4c, 0x61,
	0x75, 0x6e, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0xa4, 0x01, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x41, 0x77, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12, 0x2c, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x77, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x61, 0x77, 0x73, 0x2d, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b, 0x6c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0xba, 0x43, 0x15, 0x12, 0x13, 0x67, 0x65, 0x74,
	0x41, 0x77, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0xa1, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x77, 0x73, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x77, 0x73, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x77, 0x73, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x61, 0x77, 0x73, 0x2d,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0xba, 0x43, 0x14,
	0x12, 0x12, 0x67, 0x65, 0x74, 0x41, 0x77, 0x73, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aws_proto_rawDescOnce sync.Once
	file_aws_proto_rawDescData = file_aws_proto_rawDesc
)

func file_aws_proto_rawDescGZIP() []byte {
	file_aws_proto_rawDescOnce.Do(func() {
		file_aws_proto_rawDescData = protoimpl.X.CompressGZIP(file_aws_proto_rawDescData)
	})
	return file_aws_proto_rawDescData
}

var file_aws_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_aws_proto_goTypes = []interface{}{
	(*GetAwsMachineTypesRequest)(nil),       // 0: pb.GetAwsMachineTypesRequest
	(*GetAwsInstanceTypesRequest)(nil),      // 1: pb.GetAwsInstanceTypesRequest
	(*GetAwsMachineTypesResponse)(nil),      // 2: pb.GetAwsMachineTypesResponse
	(*GetAwsInstanceTypesResponse)(nil),     // 3: pb.GetAwsInstanceTypesResponse
	(*GetAwsLaunchTypesRequest)(nil),        // 4: pb.GetAwsLaunchTypesRequest
	(*GetAwsLaunchTypesResponse)(nil),       // 5: pb.GetAwsLaunchTypesResponse
	(*GetAwsContainerServicesRequest)(nil),  // 6: pb.GetAwsContainerServicesRequest
	(*GetAwsContainerServicesResponse)(nil), // 7: pb.GetAwsContainerServicesResponse
	(*MachineConfig)(nil),                   // 8: pb.MachineConfig
}
var file_aws_proto_depIdxs = []int32{
	8, // 0: pb.GetAwsMachineTypesResponse.machine_config:type_name -> pb.MachineConfig
	6, // 1: pb.AwsService.GetAwsContainerServices:input_type -> pb.GetAwsContainerServicesRequest
	4, // 2: pb.AwsService.GetAwsLaunchTypes:input_type -> pb.GetAwsLaunchTypesRequest
	1, // 3: pb.AwsService.GetAwsInstanceTypes:input_type -> pb.GetAwsInstanceTypesRequest
	0, // 4: pb.AwsService.GetAwsMachineTypes:input_type -> pb.GetAwsMachineTypesRequest
	7, // 5: pb.AwsService.GetAwsContainerServices:output_type -> pb.GetAwsContainerServicesResponse
	5, // 6: pb.AwsService.GetAwsLaunchTypes:output_type -> pb.GetAwsLaunchTypesResponse
	3, // 7: pb.AwsService.GetAwsInstanceTypes:output_type -> pb.GetAwsInstanceTypesResponse
	2, // 8: pb.AwsService.GetAwsMachineTypes:output_type -> pb.GetAwsMachineTypesResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aws_proto_init() }
func file_aws_proto_init() {
	if File_aws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAwsMachineTypesRequest); i {
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
		file_aws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAwsInstanceTypesRequest); i {
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
		file_aws_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAwsMachineTypesResponse); i {
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
		file_aws_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAwsInstanceTypesResponse); i {
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
		file_aws_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAwsLaunchTypesRequest); i {
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
		file_aws_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAwsLaunchTypesResponse); i {
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
		file_aws_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAwsContainerServicesRequest); i {
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
		file_aws_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAwsContainerServicesResponse); i {
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
		file_aws_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineConfig); i {
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
			RawDescriptor: file_aws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aws_proto_goTypes,
		DependencyIndexes: file_aws_proto_depIdxs,
		MessageInfos:      file_aws_proto_msgTypes,
	}.Build()
	File_aws_proto = out.File
	file_aws_proto_rawDesc = nil
	file_aws_proto_goTypes = nil
	file_aws_proto_depIdxs = nil
}