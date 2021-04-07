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

// VCSConnectionServiceClient is the client API for VCSConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VCSConnectionServiceClient interface {
	//ListAllSupportedVCSProviders returns all external providers list supported by alfred
	ListAllSupportedVCSProviders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListAllSupportedVCSProvidersResponse, error)
	//Authorize provide a URL of external OAuth application
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	//Callback called by external OAuth application
	Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*AccountVCSConnection, error)
	//ListVCSConnection list all connected service of user's account
	ListVCSConnection(ctx context.Context, in *ListVCSConnectionRequest, opts ...grpc.CallOption) (*ListVCSConnectionResponse, error)
}

type vCSConnectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVCSConnectionServiceClient(cc grpc.ClientConnInterface) VCSConnectionServiceClient {
	return &vCSConnectionServiceClient{cc}
}

func (c *vCSConnectionServiceClient) ListAllSupportedVCSProviders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListAllSupportedVCSProvidersResponse, error) {
	out := new(ListAllSupportedVCSProvidersResponse)
	err := c.cc.Invoke(ctx, "/alfred.vcs_connection.v1.VCSConnectionService/ListAllSupportedVCSProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectionServiceClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/alfred.vcs_connection.v1.VCSConnectionService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectionServiceClient) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*AccountVCSConnection, error) {
	out := new(AccountVCSConnection)
	err := c.cc.Invoke(ctx, "/alfred.vcs_connection.v1.VCSConnectionService/Callback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vCSConnectionServiceClient) ListVCSConnection(ctx context.Context, in *ListVCSConnectionRequest, opts ...grpc.CallOption) (*ListVCSConnectionResponse, error) {
	out := new(ListVCSConnectionResponse)
	err := c.cc.Invoke(ctx, "/alfred.vcs_connection.v1.VCSConnectionService/ListVCSConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VCSConnectionServiceServer is the server API for VCSConnectionService service.
// All implementations should embed UnimplementedVCSConnectionServiceServer
// for forward compatibility
type VCSConnectionServiceServer interface {
	//ListAllSupportedVCSProviders returns all external providers list supported by alfred
	ListAllSupportedVCSProviders(context.Context, *empty.Empty) (*ListAllSupportedVCSProvidersResponse, error)
	//Authorize provide a URL of external OAuth application
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	//Callback called 	by external OAuth application
	Callback(context.Context, *CallbackRequest) (*AccountVCSConnection, error)
	//ListVCSConnection list all connected service of user's account
	ListVCSConnection(context.Context, *ListVCSConnectionRequest) (*ListVCSConnectionResponse, error)
}

// UnimplementedVCSConnectionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVCSConnectionServiceServer struct {
}

func (UnimplementedVCSConnectionServiceServer) ListAllSupportedVCSProviders(context.Context, *empty.Empty) (*ListAllSupportedVCSProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllSupportedVCSProviders not implemented")
}
func (UnimplementedVCSConnectionServiceServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedVCSConnectionServiceServer) Callback(context.Context, *CallbackRequest) (*AccountVCSConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
}
func (UnimplementedVCSConnectionServiceServer) ListVCSConnection(context.Context, *ListVCSConnectionRequest) (*ListVCSConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVCSConnection not implemented")
}

// UnsafeVCSConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VCSConnectionServiceServer will
// result in compilation errors.
type UnsafeVCSConnectionServiceServer interface {
	mustEmbedUnimplementedVCSConnectionServiceServer()
}

func RegisterVCSConnectionServiceServer(s *grpc.Server, srv VCSConnectionServiceServer) {
	s.RegisterService(&_VCSConnectionService_serviceDesc, srv)
}

func _VCSConnectionService_ListAllSupportedVCSProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectionServiceServer).ListAllSupportedVCSProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.vcs_connection.v1.VCSConnectionService/ListAllSupportedVCSProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectionServiceServer).ListAllSupportedVCSProviders(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnectionService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectionServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.vcs_connection.v1.VCSConnectionService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectionServiceServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnectionService_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectionServiceServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.vcs_connection.v1.VCSConnectionService/Callback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectionServiceServer).Callback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VCSConnectionService_ListVCSConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVCSConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VCSConnectionServiceServer).ListVCSConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alfred.vcs_connection.v1.VCSConnectionService/ListVCSConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VCSConnectionServiceServer).ListVCSConnection(ctx, req.(*ListVCSConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VCSConnectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alfred.vcs_connection.v1.VCSConnectionService",
	HandlerType: (*VCSConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllSupportedVCSProviders",
			Handler:    _VCSConnectionService_ListAllSupportedVCSProviders_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _VCSConnectionService_Authorize_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _VCSConnectionService_Callback_Handler,
		},
		{
			MethodName: "ListVCSConnection",
			Handler:    _VCSConnectionService_ListVCSConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vcs-connection-service.proto",
}
