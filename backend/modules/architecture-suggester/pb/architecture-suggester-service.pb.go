// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: architecture-suggester-git.proto

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

//todo when ever new provider will come we need to add here.
type ArchitectureType int32

const (
	ArchitectureType_AWS           ArchitectureType = 0
	ArchitectureType_GCP           ArchitectureType = 1
	ArchitectureType_AZURE         ArchitectureType = 2
	ArchitectureType_DIGITAL_OCEAN ArchitectureType = 3
)

// Enum value maps for ArchitectureType.
var (
	ArchitectureType_name = map[int32]string{
		0: "AWS",
		1: "GCP",
		2: "AZURE",
		3: "DIGITAL_OCEAN",
	}
	ArchitectureType_value = map[string]int32{
		"AWS":           0,
		"GCP":           1,
		"AZURE":         2,
		"DIGITAL_OCEAN": 3,
	}
)

func (x ArchitectureType) Enum() *ArchitectureType {
	p := new(ArchitectureType)
	*p = x
	return p
}

func (x ArchitectureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArchitectureType) Descriptor() protoreflect.EnumDescriptor {
	return file_architecture_suggester_service_proto_enumTypes[0].Descriptor()
}

func (ArchitectureType) Type() protoreflect.EnumType {
	return &file_architecture_suggester_service_proto_enumTypes[0]
}

func (x ArchitectureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArchitectureType.Descriptor instead.
func (ArchitectureType) EnumDescriptor() ([]byte, []int) {
	return file_architecture_suggester_service_proto_rawDescGZIP(), []int{0}
}

type GetArchitectureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Empty object because graphql does not read empty in request
	Nothing string `protobuf:"bytes,1,opt,name=nothing,proto3" json:"nothing,omitempty"`
}

func (x *GetArchitectureRequest) Reset() {
	*x = GetArchitectureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_architecture_suggester_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArchitectureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArchitectureRequest) ProtoMessage() {}

func (x *GetArchitectureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_architecture_suggester_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArchitectureRequest.ProtoReflect.Descriptor instead.
func (*GetArchitectureRequest) Descriptor() ([]byte, []int) {
	return file_architecture_suggester_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetArchitectureRequest) GetNothing() string {
	if x != nil {
		return x.Nothing
	}
	return ""
}

//Architecture Response
type GetArchitectureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArchitectureType []ArchitectureType `protobuf:"varint,1,rep,packed,name=architecture_type,json=architectureType,proto3,enum=pb.ArchitectureType" json:"architecture_type,omitempty"`
}

func (x *GetArchitectureResponse) Reset() {
	*x = GetArchitectureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_architecture_suggester_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArchitectureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArchitectureResponse) ProtoMessage() {}

func (x *GetArchitectureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_architecture_suggester_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArchitectureResponse.ProtoReflect.Descriptor instead.
func (*GetArchitectureResponse) Descriptor() ([]byte, []int) {
	return file_architecture_suggester_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetArchitectureResponse) GetArchitectureType() []ArchitectureType {
	if x != nil {
		return x.ArchitectureType
	}
	return nil
}

var File_architecture_suggester_service_proto protoreflect.FileDescriptor

var file_architecture_suggester_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x73,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x22, 0x5c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a,
	0x11, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x2a, 0x42, 0x0a, 0x10, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x57, 0x53, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x47, 0x43, 0x50, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x5a, 0x55, 0x52, 0x45, 0x10,
	0x02, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x49, 0x47, 0x49, 0x54, 0x41, 0x4c, 0x5f, 0x4f, 0x43, 0x45,
	0x41, 0x4e, 0x10, 0x03, 0x32, 0xb4, 0x01, 0x0a, 0x1c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x73, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0xba, 0x43, 0x11, 0x12, 0x0f, 0x67, 0x65, 0x74, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_architecture_suggester_service_proto_rawDescOnce sync.Once
	file_architecture_suggester_service_proto_rawDescData = file_architecture_suggester_service_proto_rawDesc
)

func file_architecture_suggester_service_proto_rawDescGZIP() []byte {
	file_architecture_suggester_service_proto_rawDescOnce.Do(func() {
		file_architecture_suggester_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_architecture_suggester_service_proto_rawDescData)
	})
	return file_architecture_suggester_service_proto_rawDescData
}

var file_architecture_suggester_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_architecture_suggester_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_architecture_suggester_service_proto_goTypes = []interface{}{
	(ArchitectureType)(0),           // 0: pb.ArchitectureType
	(*GetArchitectureRequest)(nil),  // 1: pb.GetArchitectureRequest
	(*GetArchitectureResponse)(nil), // 2: pb.GetArchitectureResponse
}
var file_architecture_suggester_service_proto_depIdxs = []int32{
	0, // 0: pb.GetArchitectureResponse.architecture_type:type_name -> pb.ArchitectureType
	1, // 1: pb.architecture-suggester.GetArchitecture:input_type -> pb.GetArchitectureRequest
	2, // 2: pb.architecture-suggester.GetArchitecture:output_type -> pb.GetArchitectureResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_architecture_suggester_service_proto_init() }
func file_architecture_suggester_service_proto_init() {
	if File_architecture_suggester_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_architecture_suggester_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArchitectureRequest); i {
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
		file_architecture_suggester_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArchitectureResponse); i {
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
			RawDescriptor: file_architecture_suggester_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_architecture_suggester_service_proto_goTypes,
		DependencyIndexes: file_architecture_suggester_service_proto_depIdxs,
		EnumInfos:         file_architecture_suggester_service_proto_enumTypes,
		MessageInfos:      file_architecture_suggester_service_proto_msgTypes,
	}.Build()
	File_architecture_suggester_service_proto = out.File
	file_architecture_suggester_service_proto_rawDesc = nil
	file_architecture_suggester_service_proto_goTypes = nil
	file_architecture_suggester_service_proto_depIdxs = nil
}