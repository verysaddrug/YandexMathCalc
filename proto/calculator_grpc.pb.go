// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1

// source: proto/calculator.proto

package distr_calc

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

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculateServiceClient interface {
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
}

type calculateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateServiceClient(cc grpc.ClientConnInterface) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServiceServer is the server API for CalculateService service.
// All implementations must embed UnimplementedCalculateServiceServer
// for forward compatibility
type CalculateServiceServer interface {
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
	mustEmbedUnimplementedCalculateServiceServer()
}

// UnimplementedCalculateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (UnimplementedCalculateServiceServer) Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalculateServiceServer) mustEmbedUnimplementedCalculateServiceServer() {}

// UnsafeCalculateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculateServiceServer will
// result in compilation errors.
type UnsafeCalculateServiceServer interface {
	mustEmbedUnimplementedCalculateServiceServer()
}

func RegisterCalculateServiceServer(s grpc.ServiceRegistrar, srv CalculateServiceServer) {
	s.RegisterService(&CalculateService_ServiceDesc, srv)
}

func _CalculateService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculateService_ServiceDesc is the grpc.ServiceDesc for CalculateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _CalculateService_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/calculator.proto",
}
