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

// UserProfileServiceClient is the client API for UserProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProfileServiceClient interface {
	CreateUserProfile(ctx context.Context, in *CreateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
	GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
	DeleteUserProfile(ctx context.Context, in *DeleteUserProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUserMe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserProfile, error)
}

type userProfileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProfileServiceClient(cc grpc.ClientConnInterface) UserProfileServiceClient {
	return &userProfileServiceClient{cc}
}

func (c *userProfileServiceClient) CreateUserProfile(ctx context.Context, in *CreateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/pb.UserProfileService/CreateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileServiceClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/pb.UserProfileService/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileServiceClient) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/pb.UserProfileService/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileServiceClient) DeleteUserProfile(ctx context.Context, in *DeleteUserProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserProfileService/DeleteUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileServiceClient) GetUserMe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/pb.UserProfileService/GetUserMe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProfileServiceServer is the server API for UserProfileService service.
// All implementations must embed UnimplementedUserProfileServiceServer
// for forward compatibility
type UserProfileServiceServer interface {
	CreateUserProfile(context.Context, *CreateUserProfileRequest) (*UserProfile, error)
	UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserProfile, error)
	GetUserProfile(context.Context, *GetUserProfileRequest) (*UserProfile, error)
	DeleteUserProfile(context.Context, *DeleteUserProfileRequest) (*empty.Empty, error)
	GetUserMe(context.Context, *empty.Empty) (*UserProfile, error)
	//mustEmbedUnimplementedUserProfileServiceServer()
}

// UnimplementedUserProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserProfileServiceServer struct {
}

func (UnimplementedUserProfileServiceServer) CreateUserProfile(context.Context, *CreateUserProfileRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserProfile not implemented")
}
func (UnimplementedUserProfileServiceServer) UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (UnimplementedUserProfileServiceServer) GetUserProfile(context.Context, *GetUserProfileRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserProfileServiceServer) DeleteUserProfile(context.Context, *DeleteUserProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserProfile not implemented")
}
func (UnimplementedUserProfileServiceServer) GetUserMe(context.Context, *empty.Empty) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMe not implemented")
}
func (UnimplementedUserProfileServiceServer) mustEmbedUnimplementedUserProfileServiceServer() {}

// UnsafeUserProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProfileServiceServer will
// result in compilation errors.
type UnsafeUserProfileServiceServer interface {
	mustEmbedUnimplementedUserProfileServiceServer()
}

func RegisterUserProfileServiceServer(s *grpc.Server, srv UserProfileServiceServer) {
	s.RegisterService(&_UserProfileService_serviceDesc, srv)
}

func _UserProfileService_CreateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServiceServer).CreateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserProfileService/CreateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServiceServer).CreateUserProfile(ctx, req.(*CreateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfileService_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServiceServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserProfileService/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServiceServer).UpdateUserProfile(ctx, req.(*UpdateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfileService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserProfileService/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServiceServer).GetUserProfile(ctx, req.(*GetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfileService_DeleteUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServiceServer).DeleteUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserProfileService/DeleteUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServiceServer).DeleteUserProfile(ctx, req.(*DeleteUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfileService_GetUserMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServiceServer).GetUserMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserProfileService/GetUserMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServiceServer).GetUserMe(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserProfileService",
	HandlerType: (*UserProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserProfile",
			Handler:    _UserProfileService_CreateUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserProfileService_UpdateUserProfile_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserProfileService_GetUserProfile_Handler,
		},
		{
			MethodName: "DeleteUserProfile",
			Handler:    _UserProfileService_DeleteUserProfile_Handler,
		},
		{
			MethodName: "GetUserMe",
			Handler:    _UserProfileService_GetUserMe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user-profile-git-service.proto",
}