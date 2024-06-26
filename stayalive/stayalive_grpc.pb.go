// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: stayalive.proto

package stayalive

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

const (
	StayAliveService_StayAlive_FullMethodName = "/StayAliveService/StayAlive"
)

// StayAliveServiceClient is the client API for StayAliveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StayAliveServiceClient interface {
	StayAlive(ctx context.Context, in *StayAliveRequest, opts ...grpc.CallOption) (*StayAliveResponse, error)
}

type stayAliveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStayAliveServiceClient(cc grpc.ClientConnInterface) StayAliveServiceClient {
	return &stayAliveServiceClient{cc}
}

func (c *stayAliveServiceClient) StayAlive(ctx context.Context, in *StayAliveRequest, opts ...grpc.CallOption) (*StayAliveResponse, error) {
	out := new(StayAliveResponse)
	err := c.cc.Invoke(ctx, StayAliveService_StayAlive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StayAliveServiceServer is the server API for StayAliveService service.
// All implementations must embed UnimplementedStayAliveServiceServer
// for forward compatibility
type StayAliveServiceServer interface {
	StayAlive(context.Context, *StayAliveRequest) (*StayAliveResponse, error)
	mustEmbedUnimplementedStayAliveServiceServer()
}

// UnimplementedStayAliveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStayAliveServiceServer struct {
}

func (UnimplementedStayAliveServiceServer) StayAlive(context.Context, *StayAliveRequest) (*StayAliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StayAlive not implemented")
}
func (UnimplementedStayAliveServiceServer) mustEmbedUnimplementedStayAliveServiceServer() {}

// UnsafeStayAliveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StayAliveServiceServer will
// result in compilation errors.
type UnsafeStayAliveServiceServer interface {
	mustEmbedUnimplementedStayAliveServiceServer()
}

func RegisterStayAliveServiceServer(s grpc.ServiceRegistrar, srv StayAliveServiceServer) {
	s.RegisterService(&StayAliveService_ServiceDesc, srv)
}

func _StayAliveService_StayAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StayAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StayAliveServiceServer).StayAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StayAliveService_StayAlive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StayAliveServiceServer).StayAlive(ctx, req.(*StayAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StayAliveService_ServiceDesc is the grpc.ServiceDesc for StayAliveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StayAliveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StayAliveService",
	HandlerType: (*StayAliveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StayAlive",
			Handler:    _StayAliveService_StayAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stayalive.proto",
}
