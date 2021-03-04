// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: user-profile-service.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type SOURCE int32

const (
	SOURCE_UNKNOWN        SOURCE = 0
	SOURCE_GITHUB         SOURCE = 1
	SOURCE_GITLAB         SOURCE = 2
	SOURCE_BITBUCKET      SOURCE = 3
	SOURCE_EMAIL_PASSWORD SOURCE = 4
)

// Enum value maps for SOURCE.
var (
	SOURCE_name = map[int32]string{
		0: "UNKNOWN",
		1: "GITHUB",
		2: "GITLAB",
		3: "BITBUCKET",
		4: "EMAIL_PASSWORD",
	}
	SOURCE_value = map[string]int32{
		"UNKNOWN":        0,
		"GITHUB":         1,
		"GITLAB":         2,
		"BITBUCKET":      3,
		"EMAIL_PASSWORD": 4,
	}
)

func (x SOURCE) Enum() *SOURCE {
	p := new(SOURCE)
	*p = x
	return p
}

func (x SOURCE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SOURCE) Descriptor() protoreflect.EnumDescriptor {
	return file_user_profile_service_proto_enumTypes[0].Descriptor()
}

func (SOURCE) Type() protoreflect.EnumType {
	return &file_user_profile_service_proto_enumTypes[0]
}

func (x SOURCE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SOURCE.Descriptor instead.
func (SOURCE) EnumDescriptor() ([]byte, []int) {
	return file_user_profile_service_proto_rawDescGZIP(), []int{0}
}

type CreateUserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserProfile *UserProfile `protobuf:"bytes,1,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
}

func (x *CreateUserProfileRequest) Reset() {
	*x = CreateUserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserProfileRequest) ProtoMessage() {}

func (x *CreateUserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateUserProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_profile_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserProfileRequest) GetUserProfile() *UserProfile {
	if x != nil {
		return x.UserProfile
	}
	return nil
}

type DeleteUserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserProfileRequest) Reset() {
	*x = DeleteUserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserProfileRequest) ProtoMessage() {}

func (x *DeleteUserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserProfileRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_profile_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteUserProfileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateUserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserProfile *UserProfile `protobuf:"bytes,1,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
}

func (x *UpdateUserProfileRequest) Reset() {
	*x = UpdateUserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserProfileRequest) ProtoMessage() {}

func (x *UpdateUserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_profile_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserProfileRequest) GetUserProfile() *UserProfile {
	if x != nil {
		return x.UserProfile
	}
	return nil
}

type GetUserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserProfileRequest) Reset() {
	*x = GetUserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserProfileRequest) ProtoMessage() {}

func (x *GetUserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserProfileRequest.ProtoReflect.Descriptor instead.
func (*GetUserProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_profile_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserProfileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Source      SOURCE `protobuf:"varint,5,opt,name=source,proto3,enum=pb.SOURCE" json:"source,omitempty"`
	UserId      string `protobuf:"bytes,9,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_user_profile_service_proto_rawDescGZIP(), []int{4}
}

func (x *UserProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfile) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserProfile) GetSource() SOURCE {
	if x != nil {
		return x.Source
	}
	return SOURCE_UNKNOWN
}

func (x *UserProfile) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_user_profile_service_proto protoreflect.FileDescriptor

var file_user_profile_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4e, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a,
	0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x27, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x2a, 0x50, 0x0a, 0x06, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x48, 0x55,
	0x42, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x4c, 0x41, 0x42, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x42, 0x49, 0x54, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x03, 0x12, 0x12,
	0x0a, 0x0e, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44,
	0x10, 0x04, 0x32, 0xd3, 0x04, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x47,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d,
	0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0xba,
	0x43, 0x15, 0x08, 0x01, 0x12, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x47, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x29, 0x1a, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0xba, 0x43, 0x15,
	0x08, 0x01, 0x12, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x75, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x67, 0x65,
	0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0xba, 0x43, 0x06, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x92, 0x01, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b,
	0x2a, 0x29, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0xba, 0x43, 0x13, 0x12, 0x11,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x1a, 0x15, 0xba, 0x43, 0x12, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73,
	0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x10, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_profile_service_proto_rawDescOnce sync.Once
	file_user_profile_service_proto_rawDescData = file_user_profile_service_proto_rawDesc
)

func file_user_profile_service_proto_rawDescGZIP() []byte {
	file_user_profile_service_proto_rawDescOnce.Do(func() {
		file_user_profile_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_profile_service_proto_rawDescData)
	})
	return file_user_profile_service_proto_rawDescData
}

var file_user_profile_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_profile_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_profile_service_proto_goTypes = []interface{}{
	(SOURCE)(0),                      // 0: pb.SOURCE
	(*CreateUserProfileRequest)(nil), // 1: pb.CreateUserProfileRequest
	(*DeleteUserProfileRequest)(nil), // 2: pb.DeleteUserProfileRequest
	(*UpdateUserProfileRequest)(nil), // 3: pb.UpdateUserProfileRequest
	(*GetUserProfileRequest)(nil),    // 4: pb.GetUserProfileRequest
	(*UserProfile)(nil),              // 5: pb.UserProfile
	(*empty.Empty)(nil),              // 6: google.protobuf.Empty
}
var file_user_profile_service_proto_depIdxs = []int32{
	5, // 0: pb.CreateUserProfileRequest.user_profile:type_name -> pb.UserProfile
	5, // 1: pb.UpdateUserProfileRequest.user_profile:type_name -> pb.UserProfile
	0, // 2: pb.UserProfile.source:type_name -> pb.SOURCE
	1, // 3: pb.UserProfileService.CreateUserProfile:input_type -> pb.CreateUserProfileRequest
	3, // 4: pb.UserProfileService.UpdateUserProfile:input_type -> pb.UpdateUserProfileRequest
	4, // 5: pb.UserProfileService.GetUserProfile:input_type -> pb.GetUserProfileRequest
	2, // 6: pb.UserProfileService.DeleteUserProfile:input_type -> pb.DeleteUserProfileRequest
	5, // 7: pb.UserProfileService.CreateUserProfile:output_type -> pb.UserProfile
	5, // 8: pb.UserProfileService.UpdateUserProfile:output_type -> pb.UserProfile
	5, // 9: pb.UserProfileService.GetUserProfile:output_type -> pb.UserProfile
	6, // 10: pb.UserProfileService.DeleteUserProfile:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_profile_service_proto_init() }
func file_user_profile_service_proto_init() {
	if File_user_profile_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_profile_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserProfileRequest); i {
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
		file_user_profile_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserProfileRequest); i {
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
		file_user_profile_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserProfileRequest); i {
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
		file_user_profile_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserProfileRequest); i {
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
		file_user_profile_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
			RawDescriptor: file_user_profile_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_profile_service_proto_goTypes,
		DependencyIndexes: file_user_profile_service_proto_depIdxs,
		EnumInfos:         file_user_profile_service_proto_enumTypes,
		MessageInfos:      file_user_profile_service_proto_msgTypes,
	}.Build()
	File_user_profile_service_proto = out.File
	file_user_profile_service_proto_rawDesc = nil
	file_user_profile_service_proto_goTypes = nil
	file_user_profile_service_proto_depIdxs = nil
}
