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

// ValidateLoginServiceClient is the client API for ValidateLoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValidateLoginServiceClient interface {
	ValidateUserLogin(ctx context.Context, in *ValidateUserLoginRequest, opts ...grpc.CallOption) (*ValidateUserLoginResponse, error)
}

type validateLoginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValidateLoginServiceClient(cc grpc.ClientConnInterface) ValidateLoginServiceClient {
	return &validateLoginServiceClient{cc}
}

func (c *validateLoginServiceClient) ValidateUserLogin(ctx context.Context, in *ValidateUserLoginRequest, opts ...grpc.CallOption) (*ValidateUserLoginResponse, error) {
	out := new(ValidateUserLoginResponse)
	err := c.cc.Invoke(ctx, "/pb.ValidateLoginService/ValidateUserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidateLoginServiceServer is the server API for ValidateLoginService service.
// All implementations must embed UnimplementedValidateLoginServiceServer
// for forward compatibility
type ValidateLoginServiceServer interface {
	ValidateUserLogin(context.Context, *ValidateUserLoginRequest) (*ValidateUserLoginResponse, error)
	mustEmbedUnimplementedValidateLoginServiceServer()
}

// UnimplementedValidateLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedValidateLoginServiceServer struct {
}

func (UnimplementedValidateLoginServiceServer) ValidateUserLogin(context.Context, *ValidateUserLoginRequest) (*ValidateUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserLogin not implemented")
}
func (UnimplementedValidateLoginServiceServer) mustEmbedUnimplementedValidateLoginServiceServer() {}

// UnsafeValidateLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValidateLoginServiceServer will
// result in compilation errors.
type UnsafeValidateLoginServiceServer interface {
	mustEmbedUnimplementedValidateLoginServiceServer()
}

func RegisterValidateLoginServiceServer(s *grpc.Server, srv ValidateLoginServiceServer) {
	s.RegisterService(&_ValidateLoginService_serviceDesc, srv)
}

func _ValidateLoginService_ValidateUserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidateLoginServiceServer).ValidateUserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ValidateLoginService/ValidateUserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidateLoginServiceServer).ValidateUserLogin(ctx, req.(*ValidateUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ValidateLoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ValidateLoginService",
	HandlerType: (*ValidateLoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateUserLogin",
			Handler:    _ValidateLoginService_ValidateUserLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth-service.proto",
}

// UserLoginServiceClient is the client API for UserLoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserLoginServiceClient interface {
	// CreateUserLogin creates a new user
	UserLogin(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserLoginResponse, error)
	//callback endpoint
	UserLoginCallback(ctx context.Context, in *UserLoginCallbackRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UserLogout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// GetUserLogin returns the specified user by its id.
	GetUserLogin(ctx context.Context, in *GetUserLoginRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// DeleteUserLogin is used to delete a user from the system, this will delete all
	DeleteUserLogin(ctx context.Context, in *DeleteUserLoginRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateUserPassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// ResetUserPassword , if a user has forgot the password then
	// this method can be used to reset the password
	ResetUserPassword(ctx context.Context, in *ResetUserPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userLoginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLoginServiceClient(cc grpc.ClientConnInterface) UserLoginServiceClient {
	return &userLoginServiceClient{cc}
}

func (c *userLoginServiceClient) UserLogin(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/pb.UserLoginService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) UserLoginCallback(ctx context.Context, in *UserLoginCallbackRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserLoginService/UserLoginCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) UserLogout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserLoginService/UserLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) GetUserLogin(ctx context.Context, in *GetUserLoginRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserLoginService/GetUserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) DeleteUserLogin(ctx context.Context, in *DeleteUserLoginRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserLoginService/DeleteUserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) UpdateUserPassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserLoginService/UpdateUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) ResetUserPassword(ctx context.Context, in *ResetUserPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserLoginService/ResetUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLoginServiceServer is the server API for UserLoginService service.
// All implementations must embed UnimplementedUserLoginServiceServer
// for forward compatibility
type UserLoginServiceServer interface {
	// CreateUserLogin creates a new user
	UserLogin(context.Context, *empty.Empty) (*UserLoginResponse, error)
	//callback endpoint
	UserLoginCallback(context.Context, *UserLoginCallbackRequest) (*empty.Empty, error)
	UserLogout(context.Context, *empty.Empty) (*empty.Empty, error)
	// GetUserLogin returns the specified user by its id.
	GetUserLogin(context.Context, *GetUserLoginRequest) (*empty.Empty, error)
	// DeleteUserLogin is used to delete a user from the system, this will delete all
	DeleteUserLogin(context.Context, *DeleteUserLoginRequest) (*empty.Empty, error)
	UpdateUserPassword(context.Context, *UpdateUserPasswordRequest) (*empty.Empty, error)
	// ResetUserPassword , if a user has forgot the password then
	// this method can be used to reset the password
	ResetUserPassword(context.Context, *ResetUserPasswordRequest) (*empty.Empty, error)
	//mustEmbedUnimplementedUserLoginServiceServer()
}

// UnimplementedUserLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserLoginServiceServer struct {
}

func (UnimplementedUserLoginServiceServer) UserLogin(context.Context, *empty.Empty) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserLoginServiceServer) UserLoginCallback(context.Context, *UserLoginCallbackRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLoginCallback not implemented")
}
func (UnimplementedUserLoginServiceServer) UserLogout(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogout not implemented")
}
func (UnimplementedUserLoginServiceServer) GetUserLogin(context.Context, *GetUserLoginRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLogin not implemented")
}
func (UnimplementedUserLoginServiceServer) DeleteUserLogin(context.Context, *DeleteUserLoginRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserLogin not implemented")
}
func (UnimplementedUserLoginServiceServer) UpdateUserPassword(context.Context, *UpdateUserPasswordRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPassword not implemented")
}
func (UnimplementedUserLoginServiceServer) ResetUserPassword(context.Context, *ResetUserPasswordRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetUserPassword not implemented")
}
func (UnimplementedUserLoginServiceServer) mustEmbedUnimplementedUserLoginServiceServer() {}

// UnsafeUserLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLoginServiceServer will
// result in compilation errors.
type UnsafeUserLoginServiceServer interface {
	mustEmbedUnimplementedUserLoginServiceServer()
}

func RegisterUserLoginServiceServer(s *grpc.Server, srv UserLoginServiceServer) {
	s.RegisterService(&_UserLoginService_serviceDesc, srv)
}

func _UserLoginService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserLoginService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).UserLogin(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_UserLoginCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).UserLoginCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserLoginService/UserLoginCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).UserLoginCallback(ctx, req.(*UserLoginCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_UserLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).UserLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserLoginService/UserLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).UserLogout(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_GetUserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).GetUserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserLoginService/GetUserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).GetUserLogin(ctx, req.(*GetUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_DeleteUserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).DeleteUserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserLoginService/DeleteUserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).DeleteUserLogin(ctx, req.(*DeleteUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_UpdateUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).UpdateUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserLoginService/UpdateUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).UpdateUserPassword(ctx, req.(*UpdateUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_ResetUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).ResetUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserLoginService/ResetUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).ResetUserPassword(ctx, req.(*ResetUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserLoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserLoginService",
	HandlerType: (*UserLoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserLoginService_UserLogin_Handler,
		},
		{
			MethodName: "UserLoginCallback",
			Handler:    _UserLoginService_UserLoginCallback_Handler,
		},
		{
			MethodName: "UserLogout",
			Handler:    _UserLoginService_UserLogout_Handler,
		},
		{
			MethodName: "GetUserLogin",
			Handler:    _UserLoginService_GetUserLogin_Handler,
		},
		{
			MethodName: "DeleteUserLogin",
			Handler:    _UserLoginService_DeleteUserLogin_Handler,
		},
		{
			MethodName: "UpdateUserPassword",
			Handler:    _UserLoginService_UpdateUserPassword_Handler,
		},
		{
			MethodName: "ResetUserPassword",
			Handler:    _UserLoginService_ResetUserPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth-service.proto",
}