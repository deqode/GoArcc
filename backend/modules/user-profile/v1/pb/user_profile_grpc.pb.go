// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserProfilesClient is the client API for UserProfiles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProfilesClient interface {
	// GetUserProfile return a profile of a user
	GetUserProfile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserProfile, error)
	//GetUserProfileBySub return a user profile by its unique sub provided by vcs
	GetUserProfileBySub(ctx context.Context, in *GetUserProfileBySubRequest, opts ...grpc.CallOption) (*UserProfile, error)
}

type userProfilesClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProfilesClient(cc grpc.ClientConnInterface) UserProfilesClient {
	return &userProfilesClient{cc}
}

func (c *userProfilesClient) GetUserProfile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/goarcc.user_profile.v1.UserProfiles/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfilesClient) GetUserProfileBySub(ctx context.Context, in *GetUserProfileBySubRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/goarcc.user_profile.v1.UserProfiles/GetUserProfileBySub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProfilesServer is the server API for UserProfiles service.
// All implementations should embed UnimplementedUserProfilesServer
// for forward compatibility
type UserProfilesServer interface {
	// GetUserProfile return a profile of a user
	GetUserProfile(context.Context, *empty.Empty) (*UserProfile, error)
	//GetUserProfileBySub return a user profile by its unique sub provided by vcs
	GetUserProfileBySub(context.Context, *GetUserProfileBySubRequest) (*UserProfile, error)
}

// UnimplementedUserProfilesServer should be embedded to have forward compatible implementations.
type UnimplementedUserProfilesServer struct {
}

func (UnimplementedUserProfilesServer) GetUserProfile(context.Context, *empty.Empty) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserProfilesServer) GetUserProfileBySub(context.Context, *GetUserProfileBySubRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfileBySub not implemented")
}

// UnsafeUserProfilesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProfilesServer will
// result in compilation errors.
type UnsafeUserProfilesServer interface {
	mustEmbedUnimplementedUserProfilesServer()
}

func RegisterUserProfilesServer(s *grpc.Server, srv UserProfilesServer) {
	s.RegisterService(&_UserProfiles_serviceDesc, srv)
}

func _UserProfiles_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfilesServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goarcc.user_profile.v1.UserProfiles/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfilesServer).GetUserProfile(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfiles_GetUserProfileBySub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileBySubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfilesServer).GetUserProfileBySub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goarcc.user_profile.v1.UserProfiles/GetUserProfileBySub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfilesServer).GetUserProfileBySub(ctx, req.(*GetUserProfileBySubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserProfiles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goarcc.user_profile.v1.UserProfiles",
	HandlerType: (*UserProfilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserProfile",
			Handler:    _UserProfiles_GetUserProfile_Handler,
		},
		{
			MethodName: "GetUserProfileBySub",
			Handler:    _UserProfiles_GetUserProfileBySub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/user_profile.proto",
}
