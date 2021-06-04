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

// StacksClient is the client API for Stacks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StacksClient interface {
	//CreateStack create a new stack correspond to an account
	CreateStack(ctx context.Context, in *CreateStackRequest, opts ...grpc.CallOption) (*Stack, error)
	// UpdateStack update existing fields in stack
	UpdateStack(ctx context.Context, in *UpdateStackRequest, opts ...grpc.CallOption) (*Stack, error)
	// GetStack get a existing stack
	GetStack(ctx context.Context, in *GetStackRequest, opts ...grpc.CallOption) (*Stack, error)
	// DeleteStack delete an existing stack
	DeleteStack(ctx context.Context, in *DeleteStackRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// ListStack list existing stacks
	ListStack(ctx context.Context, in *ListStackRequest, opts ...grpc.CallOption) (*ListStackResponse, error)
	// Stack Builds
	//CreateStackBuild create new stack build
	CreateStackBuild(ctx context.Context, in *CreateStackBuildRequest, opts ...grpc.CallOption) (*StackBuild, error)
	//GetStackBuild get information of existing stack build
	GetStackBuild(ctx context.Context, in *GetStackBuildRequest, opts ...grpc.CallOption) (*StackBuild, error)
	//DeleteStackBuild deletes existing stack build
	DeleteStackBuild(ctx context.Context, in *DeleteStackBuildRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// ListStackBuild list all the existing builds of a stack
	ListStackBuild(ctx context.Context, in *ListStackBuildRequest, opts ...grpc.CallOption) (*ListStackBuildResponse, error)
}

type stacksClient struct {
	cc grpc.ClientConnInterface
}

func NewStacksClient(cc grpc.ClientConnInterface) StacksClient {
	return &stacksClient{cc}
}

func (c *stacksClient) CreateStack(ctx context.Context, in *CreateStackRequest, opts ...grpc.CallOption) (*Stack, error) {
	out := new(Stack)
	err := c.cc.Invoke(ctx, "/alfred.stack.v1.Stacks/CreateStack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacksClient) UpdateStack(ctx context.Context, in *UpdateStackRequest, opts ...grpc.CallOption) (*Stack, error) {
	out := new(Stack)
	err := c.cc.Invoke(ctx, "/alfred.stack.v1.Stacks/UpdateStack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacksClient) GetStack(ctx context.Context, in *GetStackRequest, opts ...grpc.CallOption) (*Stack, error) {
	out := new(Stack)
	err := c.cc.Invoke(ctx, "/alfred.stack.v1.Stacks/GetStack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacksClient) DeleteStack(ctx context.Context, in *DeleteStackRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/alfred.stack.v1.Stacks/DeleteStack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacksClient) ListStack(ctx context.Context, in *ListStackRequest, opts ...grpc.CallOption) (*ListStackResponse, error) {
	out := new(ListStackResponse)
	err := c.cc.Invoke(ctx, "/alfred.stack.v1.Stacks/ListStack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacksClient) CreateStackBuild(ctx context.Context, in *CreateStackBuildRequest, opts ...grpc.CallOption) (*StackBuild, error) {
	out := new(StackBuild)
	err := c.cc.Invoke(ctx, "/alfred.stack.v1.Stacks/CreateStackBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacksClient) GetStackBuild(ctx context.Context, in *GetStackBuildRequest, opts ...grpc.CallOption) (*StackBuild, error) {
	out := new(StackBuild)
	err := c.cc.Invoke(ctx, "/alfred.stack.v1.Stacks/GetStackBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacksClient) DeleteStackBuild(ctx context.Context, in *DeleteStackBuildRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/alfred.stack.v1.Stacks/DeleteStackBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacksClient) ListStackBuild(ctx context.Context, in *ListStackBuildRequest, opts ...grpc.CallOption) (*ListStackBuildResponse, error) {
	out := new(ListStackBuildResponse)
	err := c.cc.Invoke(ctx, "/alfred.stack.v1.Stacks/ListStackBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StacksServer is the server API for Stacks service.
// All implementations should embed UnimplementedStacksServer
// for forward compatibility
type StacksServer interface {
	//CreateStack create a new stack correspond to an account
	CreateStack(context.Context, *CreateStackRequest) (*Stack, error)
	// UpdateStack update existing fields in stack
	UpdateStack(context.Context, *UpdateStackRequest) (*Stack, error)
	// GetStack get a existing stack
	GetStack(context.Context, *GetStackRequest) (*Stack, error)
	// DeleteStack delete an existing stack
	DeleteStack(context.Context, *DeleteStackRequest) (*empty.Empty, error)
	// ListStack list existing stacks
	ListStack(context.Context, *ListStackRequest) (*ListStackResponse, error)
	// Stack Builds
	//CreateStackBuild create new stack build
	CreateStackBuild(context.Context, *CreateStackBuildRequest) (*StackBuild, error)
	//GetStackBuild get information of existing stack build
	GetStackBuild(context.Context, *GetStackBuildRequest) (*StackBuild, error)
	//DeleteStackBuild deletes existing stack build
	DeleteStackBuild(context.Context, *DeleteStackBuildRequest) (*empty.Empty, error)
	// ListStackBuild list all the existing builds of a stack
	ListStackBuild(context.Context, *ListStackBuildRequest) (*ListStackBuildResponse, error)
}

// UnimplementedStacksServer should be embedded to have forward compatible implementations.
type UnimplementedStacksServer struct {
}

func (UnimplementedStacksServer) CreateStack(context.Context, *CreateStackRequest) (*Stack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStack not implemented")
}
func (UnimplementedStacksServer) UpdateStack(context.Context, *UpdateStackRequest) (*Stack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStack not implemented")
}
func (UnimplementedStacksServer) GetStack(context.Context, *GetStackRequest) (*Stack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStack not implemented")
}
func (UnimplementedStacksServer) DeleteStack(context.Context, *DeleteStackRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStack not implemented")
}
func (UnimplementedStacksServer) ListStack(context.Context, *ListStackRequest) (*ListStackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStack not implemented")
}
func (UnimplementedStacksServer) CreateStackBuild(context.Context, *CreateStackBuildRequest) (*StackBuild, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStackBuild not implemented")
}
func (UnimplementedStacksServer) GetStackBuild(context.Context, *GetStackBuildRequest) (*StackBuild, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackBuild not implemented")
}
func (UnimplementedStacksServer) DeleteStackBuild(context.Context, *DeleteStackBuildRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStackBuild not implemented")
}
func (UnimplementedStacksServer) ListStackBuild(context.Context, *ListStackBuildRequest) (*ListStackBuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStackBuild not implemented")
}

// UnsafeStacksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StacksServer will
// result in compilation errors.
type UnsafeStacksServer interface {
	mustEmbedUnimplementedStacksServer()
}

func RegisterStacksServer(s *grpc.Server, srv StacksServer) {
	s.RegisterService(&_Stacks_serviceDesc, srv)
}

func _Stacks_CreateStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacksServer).CreateStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.stack.v1.Stacks/CreateStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacksServer).CreateStack(ctx, req.(*CreateStackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stacks_UpdateStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacksServer).UpdateStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.stack.v1.Stacks/UpdateStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacksServer).UpdateStack(ctx, req.(*UpdateStackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stacks_GetStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacksServer).GetStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.stack.v1.Stacks/GetStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacksServer).GetStack(ctx, req.(*GetStackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stacks_DeleteStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacksServer).DeleteStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.stack.v1.Stacks/DeleteStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacksServer).DeleteStack(ctx, req.(*DeleteStackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stacks_ListStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacksServer).ListStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.stack.v1.Stacks/ListStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacksServer).ListStack(ctx, req.(*ListStackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stacks_CreateStackBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStackBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacksServer).CreateStackBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.stack.v1.Stacks/CreateStackBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacksServer).CreateStackBuild(ctx, req.(*CreateStackBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stacks_GetStackBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStackBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacksServer).GetStackBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.stack.v1.Stacks/GetStackBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacksServer).GetStackBuild(ctx, req.(*GetStackBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stacks_DeleteStackBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStackBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacksServer).DeleteStackBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.stack.v1.Stacks/DeleteStackBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacksServer).DeleteStackBuild(ctx, req.(*DeleteStackBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stacks_ListStackBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStackBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacksServer).ListStackBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.stack.v1.Stacks/ListStackBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacksServer).ListStackBuild(ctx, req.(*ListStackBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stacks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alfred.stack.v1.Stacks",
	HandlerType: (*StacksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStack",
			Handler:    _Stacks_CreateStack_Handler,
		},
		{
			MethodName: "UpdateStack",
			Handler:    _Stacks_UpdateStack_Handler,
		},
		{
			MethodName: "GetStack",
			Handler:    _Stacks_GetStack_Handler,
		},
		{
			MethodName: "DeleteStack",
			Handler:    _Stacks_DeleteStack_Handler,
		},
		{
			MethodName: "ListStack",
			Handler:    _Stacks_ListStack_Handler,
		},
		{
			MethodName: "CreateStackBuild",
			Handler:    _Stacks_CreateStackBuild_Handler,
		},
		{
			MethodName: "GetStackBuild",
			Handler:    _Stacks_GetStackBuild_Handler,
		},
		{
			MethodName: "DeleteStackBuild",
			Handler:    _Stacks_DeleteStackBuild_Handler,
		},
		{
			MethodName: "ListStackBuild",
			Handler:    _Stacks_ListStackBuild_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/stack.proto",
}
