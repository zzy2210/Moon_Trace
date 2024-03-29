// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppClient interface {
	HandleAppDomain(ctx context.Context, in *AppDomainRequest, opts ...grpc.CallOption) (*AppDomainResponse, error)
	HandleAppPort(ctx context.Context, in *AppPortRequest, opts ...grpc.CallOption) (*AppPortResponse, error)
	HandleAppUrl(ctx context.Context, in *AppUrlRequest, opts ...grpc.CallOption) (*AppUrlResponse, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) HandleAppDomain(ctx context.Context, in *AppDomainRequest, opts ...grpc.CallOption) (*AppDomainResponse, error) {
	out := new(AppDomainResponse)
	err := c.cc.Invoke(ctx, "/api.eng.v1.App/handleAppDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) HandleAppPort(ctx context.Context, in *AppPortRequest, opts ...grpc.CallOption) (*AppPortResponse, error) {
	out := new(AppPortResponse)
	err := c.cc.Invoke(ctx, "/api.eng.v1.App/handleAppPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) HandleAppUrl(ctx context.Context, in *AppUrlRequest, opts ...grpc.CallOption) (*AppUrlResponse, error) {
	out := new(AppUrlResponse)
	err := c.cc.Invoke(ctx, "/api.eng.v1.App/handleAppUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
// All implementations must embed UnimplementedAppServer
// for forward compatibility
type AppServer interface {
	HandleAppDomain(context.Context, *AppDomainRequest) (*AppDomainResponse, error)
	HandleAppPort(context.Context, *AppPortRequest) (*AppPortResponse, error)
	HandleAppUrl(context.Context, *AppUrlRequest) (*AppUrlResponse, error)
	mustEmbedUnimplementedAppServer()
}

// UnimplementedAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (UnimplementedAppServer) HandleAppDomain(context.Context, *AppDomainRequest) (*AppDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleAppDomain not implemented")
}
func (UnimplementedAppServer) HandleAppPort(context.Context, *AppPortRequest) (*AppPortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleAppPort not implemented")
}
func (UnimplementedAppServer) HandleAppUrl(context.Context, *AppUrlRequest) (*AppUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleAppUrl not implemented")
}
func (UnimplementedAppServer) mustEmbedUnimplementedAppServer() {}

// UnsafeAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServer will
// result in compilation errors.
type UnsafeAppServer interface {
	mustEmbedUnimplementedAppServer()
}

func RegisterAppServer(s grpc.ServiceRegistrar, srv AppServer) {
	s.RegisterService(&App_ServiceDesc, srv)
}

func _App_HandleAppDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).HandleAppDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.eng.v1.App/handleAppDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).HandleAppDomain(ctx, req.(*AppDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_HandleAppPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppPortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).HandleAppPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.eng.v1.App/handleAppPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).HandleAppPort(ctx, req.(*AppPortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_HandleAppUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).HandleAppUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.eng.v1.App/handleAppUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).HandleAppUrl(ctx, req.(*AppUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// App_ServiceDesc is the grpc.ServiceDesc for App service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.eng.v1.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "handleAppDomain",
			Handler:    _App_HandleAppDomain_Handler,
		},
		{
			MethodName: "handleAppPort",
			Handler:    _App_HandleAppPort_Handler,
		},
		{
			MethodName: "handleAppUrl",
			Handler:    _App_HandleAppUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/eng/v1/api.proto",
}
