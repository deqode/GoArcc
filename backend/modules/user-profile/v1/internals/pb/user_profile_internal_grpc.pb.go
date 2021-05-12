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

// UserProfileInternalClient is the client API for UserProfileInternal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProfileInternalClient interface {
	CreateUserProfile(ctx context.Context, in *CreateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
	DeleteUserProfile(ctx context.Context, in *DeleteUserProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUserMe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserProfile, error)
}

type userProfileInternalClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProfileInternalClient(cc grpc.ClientConnInterface) UserProfileInternalClient {
	return &userProfileInternalClient{cc}
}

func (c *userProfileInternalClient) CreateUserProfile(ctx context.Context, in *CreateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/alfred.user_profile.v1.internal.UserProfileInternal/CreateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileInternalClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/alfred.user_profile.v1.internal.UserProfileInternal/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileInternalClient) DeleteUserProfile(ctx context.Context, in *DeleteUserProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/alfred.user_profile.v1.internal.UserProfileInternal/DeleteUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileInternalClient) GetUserMe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/alfred.user_profile.v1.internal.UserProfileInternal/GetUserMe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProfileInternalServer is the server API for UserProfileInternal service.
// All implementations should embed UnimplementedUserProfileInternalServer
// for forward compatibility
type UserProfileInternalServer interface {
	CreateUserProfile(context.Context, *CreateUserProfileRequest) (*UserProfile, error)
	UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserProfile, error)
	DeleteUserProfile(context.Context, *DeleteUserProfileRequest) (*empty.Empty, error)
	GetUserMe(context.Context, *empty.Empty) (*UserProfile, error)
}

// UnimplementedUserProfileInternalServer should be embedded to have forward compatible implementations.
type UnimplementedUserProfileInternalServer struct {
}

func (UnimplementedUserProfileInternalServer) CreateUserProfile(context.Context, *CreateUserProfileRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserProfile not implemented")
}
func (UnimplementedUserProfileInternalServer) UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (UnimplementedUserProfileInternalServer) DeleteUserProfile(context.Context, *DeleteUserProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserProfile not implemented")
}
func (UnimplementedUserProfileInternalServer) GetUserMe(context.Context, *empty.Empty) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMe not implemented")
}

// UnsafeUserProfileInternalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProfileInternalServer will
// result in compilation errors.
type UnsafeUserProfileInternalServer interface {
	mustEmbedUnimplementedUserProfileInternalServer()
}

func RegisterUserProfileInternalServer(s *grpc.Server, srv UserProfileInternalServer) {
	s.RegisterService(&_UserProfileInternal_serviceDesc, srv)
}

func _UserProfileInternal_CreateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileInternalServer).CreateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.user_profile.v1.internal.UserProfileInternal/CreateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileInternalServer).CreateUserProfile(ctx, req.(*CreateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfileInternal_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileInternalServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.user_profile.v1.internal.UserProfileInternal/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileInternalServer).UpdateUserProfile(ctx, req.(*UpdateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfileInternal_DeleteUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileInternalServer).DeleteUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.user_profile.v1.internal.UserProfileInternal/DeleteUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileInternalServer).DeleteUserProfile(ctx, req.(*DeleteUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfileInternal_GetUserMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileInternalServer).GetUserMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.user_profile.v1.internal.UserProfileInternal/GetUserMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileInternalServer).GetUserMe(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserProfileInternal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alfred.user_profile.v1.internal.UserProfileInternal",
	HandlerType: (*UserProfileInternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserProfile",
			Handler:    _UserProfileInternal_CreateUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserProfileInternal_UpdateUserProfile_Handler,
		},
		{
			MethodName: "DeleteUserProfile",
			Handler:    _UserProfileInternal_DeleteUserProfile_Handler,
		},
		{
			MethodName: "GetUserMe",
			Handler:    _UserProfileInternal_GetUserMe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_profile_internal.proto",
}
