// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: pb/cloud-connection-service.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateCloudConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudConnection *CloudConnection `protobuf:"bytes,1,opt,name=cloud_connection,json=cloudConnection,proto3" json:"cloud_connection,omitempty"`
}

func (x *CreateCloudConnectionRequest) Reset() {
	*x = CreateCloudConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cloud_connection_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCloudConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCloudConnectionRequest) ProtoMessage() {}

func (x *CreateCloudConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cloud_connection_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCloudConnectionRequest.ProtoReflect.Descriptor instead.
func (*CreateCloudConnectionRequest) Descriptor() ([]byte, []int) {
	return file_pb_cloud_connection_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCloudConnectionRequest) GetCloudConnection() *CloudConnection {
	if x != nil {
		return x.CloudConnection
	}
	return nil
}

type DeleteCloudConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCloudConnectionRequest) Reset() {
	*x = DeleteCloudConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cloud_connection_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCloudConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCloudConnectionRequest) ProtoMessage() {}

func (x *DeleteCloudConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cloud_connection_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCloudConnectionRequest.ProtoReflect.Descriptor instead.
func (*DeleteCloudConnectionRequest) Descriptor() ([]byte, []int) {
	return file_pb_cloud_connection_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteCloudConnectionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateCloudConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudConnection *CloudConnection `protobuf:"bytes,1,opt,name=cloud_connection,json=cloudConnection,proto3" json:"cloud_connection,omitempty"`
}

func (x *UpdateCloudConnectionRequest) Reset() {
	*x = UpdateCloudConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cloud_connection_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCloudConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCloudConnectionRequest) ProtoMessage() {}

func (x *UpdateCloudConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cloud_connection_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCloudConnectionRequest.ProtoReflect.Descriptor instead.
func (*UpdateCloudConnectionRequest) Descriptor() ([]byte, []int) {
	return file_pb_cloud_connection_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCloudConnectionRequest) GetCloudConnection() *CloudConnection {
	if x != nil {
		return x.CloudConnection
	}
	return nil
}

type GetCloudConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCloudConnectionRequest) Reset() {
	*x = GetCloudConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cloud_connection_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCloudConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCloudConnectionRequest) ProtoMessage() {}

func (x *GetCloudConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cloud_connection_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCloudConnectionRequest.ProtoReflect.Descriptor instead.
func (*GetCloudConnectionRequest) Descriptor() ([]byte, []int) {
	return file_pb_cloud_connection_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCloudConnectionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CloudConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//id generated by uuid
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// stored in encrypted form
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// stored in encrypted form
	ClientSecret string   `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	AccountId    string   `protobuf:"bytes,4,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Scope        []string `protobuf:"bytes,5,rep,name=scope,proto3" json:"scope,omitempty"`
	Provider     string   `protobuf:"bytes,6,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *CloudConnection) Reset() {
	*x = CloudConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cloud_connection_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudConnection) ProtoMessage() {}

func (x *CloudConnection) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cloud_connection_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudConnection.ProtoReflect.Descriptor instead.
func (*CloudConnection) Descriptor() ([]byte, []int) {
	return file_pb_cloud_connection_service_proto_rawDescGZIP(), []int{4}
}

func (x *CloudConnection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CloudConnection) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CloudConnection) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *CloudConnection) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CloudConnection) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *CloudConnection) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

var File_pb_cloud_connection_service_proto protoreflect.FileDescriptor

var file_pb_cloud_connection_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x1c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x10, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x1c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x1c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x10, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb4, 0x01, 0x0a, 0x0f, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x32, 0xa4, 0x05,
	0x0a, 0x16, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa3, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x31, 0x22, 0x2c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x01, 0x2a, 0xba, 0x43, 0x19, 0x08, 0x01, 0x12, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xa3,
	0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x53, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x1a, 0x2c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0xba, 0x43, 0x19, 0x08, 0x01, 0x12, 0x15, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x94, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65,
	0x74, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0xba, 0x43, 0x11, 0x12, 0x0f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xa6, 0x01, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x53, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x2a, 0x31, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0xba, 0x43, 0x17, 0x12, 0x15, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_cloud_connection_service_proto_rawDescOnce sync.Once
	file_pb_cloud_connection_service_proto_rawDescData = file_pb_cloud_connection_service_proto_rawDesc
)

func file_pb_cloud_connection_service_proto_rawDescGZIP() []byte {
	file_pb_cloud_connection_service_proto_rawDescOnce.Do(func() {
		file_pb_cloud_connection_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_cloud_connection_service_proto_rawDescData)
	})
	return file_pb_cloud_connection_service_proto_rawDescData
}

var file_pb_cloud_connection_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_cloud_connection_service_proto_goTypes = []interface{}{
	(*CreateCloudConnectionRequest)(nil), // 0: pb.CreateCloudConnectionRequest
	(*DeleteCloudConnectionRequest)(nil), // 1: pb.DeleteCloudConnectionRequest
	(*UpdateCloudConnectionRequest)(nil), // 2: pb.UpdateCloudConnectionRequest
	(*GetCloudConnectionRequest)(nil),    // 3: pb.GetCloudConnectionRequest
	(*CloudConnection)(nil),              // 4: pb.CloudConnection
	(*empty.Empty)(nil),                  // 5: google.protobuf.Empty
}
var file_pb_cloud_connection_service_proto_depIdxs = []int32{
	4, // 0: pb.CreateCloudConnectionRequest.cloud_connection:type_name -> pb.CloudConnection
	4, // 1: pb.UpdateCloudConnectionRequest.cloud_connection:type_name -> pb.CloudConnection
	0, // 2: pb.CloudConnectionService.CreateCloudConnection:input_type -> pb.CreateCloudConnectionRequest
	2, // 3: pb.CloudConnectionService.UpdateCloudConnection:input_type -> pb.UpdateCloudConnectionRequest
	3, // 4: pb.CloudConnectionService.GetCloudConnection:input_type -> pb.GetCloudConnectionRequest
	1, // 5: pb.CloudConnectionService.DeleteCloudConnection:input_type -> pb.DeleteCloudConnectionRequest
	4, // 6: pb.CloudConnectionService.CreateCloudConnection:output_type -> pb.CloudConnection
	4, // 7: pb.CloudConnectionService.UpdateCloudConnection:output_type -> pb.CloudConnection
	4, // 8: pb.CloudConnectionService.GetCloudConnection:output_type -> pb.CloudConnection
	5, // 9: pb.CloudConnectionService.DeleteCloudConnection:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_cloud_connection_service_proto_init() }
func file_pb_cloud_connection_service_proto_init() {
	if File_pb_cloud_connection_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_cloud_connection_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCloudConnectionRequest); i {
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
		file_pb_cloud_connection_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCloudConnectionRequest); i {
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
		file_pb_cloud_connection_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCloudConnectionRequest); i {
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
		file_pb_cloud_connection_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCloudConnectionRequest); i {
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
		file_pb_cloud_connection_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudConnection); i {
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
			RawDescriptor: file_pb_cloud_connection_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_cloud_connection_service_proto_goTypes,
		DependencyIndexes: file_pb_cloud_connection_service_proto_depIdxs,
		MessageInfos:      file_pb_cloud_connection_service_proto_msgTypes,
	}.Build()
	File_pb_cloud_connection_service_proto = out.File
	file_pb_cloud_connection_service_proto_rawDesc = nil
	file_pb_cloud_connection_service_proto_goTypes = nil
	file_pb_cloud_connection_service_proto_depIdxs = nil
}