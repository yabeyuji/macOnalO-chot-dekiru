// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// DeliveryServiceClient is the client API for DeliveryService gateway.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryServiceClient interface {
	DeliveryRPC(ctx context.Context, in *DeliveryRequest, opts ...grpc.CallOption) (*DeliveryResponse, error)
}

type deliveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryServiceClient(cc grpc.ClientConnInterface) DeliveryServiceClient {
	return &deliveryServiceClient{cc}
}

func (c *deliveryServiceClient) DeliveryRPC(ctx context.Context, in *DeliveryRequest, opts ...grpc.CallOption) (*DeliveryResponse, error) {
	out := new(DeliveryResponse)
	err := c.cc.Invoke(ctx, "/delivery.DeliveryService/DeliveryRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryServiceServer is the server API for DeliveryService gateway.
// All implementations must embed UnimplementedDeliveryServiceServer
// for forward compatibility
type DeliveryServiceServer interface {
	DeliveryRPC(context.Context, *DeliveryRequest) (*DeliveryResponse, error)
	mustEmbedUnimplementedDeliveryServiceServer()
}

// UnimplementedDeliveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeliveryServiceServer struct {
}

func (UnimplementedDeliveryServiceServer) DeliveryRPC(context.Context, *DeliveryRequest) (*DeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliveryRPC not implemented")
}
func (UnimplementedDeliveryServiceServer) mustEmbedUnimplementedDeliveryServiceServer() {}

// UnsafeDeliveryServiceServer may be embedded to opt out of forward compatibility for this gateway.
// Use of this interface is not recommended, as added methods to DeliveryServiceServer will
// result in compilation errors.
type UnsafeDeliveryServiceServer interface {
	mustEmbedUnimplementedDeliveryServiceServer()
}

func RegisterDeliveryServiceServer(s grpc.ServiceRegistrar, srv DeliveryServiceServer) {
	s.RegisterService(&DeliveryService_ServiceDesc, srv)
}

func _DeliveryService_DeliveryRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).DeliveryRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/delivery.DeliveryService/DeliveryRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).DeliveryRPC(ctx, req.(*DeliveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryService_ServiceDesc is the grpc.ServiceDesc for DeliveryService gateway.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeliveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "delivery.DeliveryService",
	HandlerType: (*DeliveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeliveryRPC",
			Handler:    _DeliveryService_DeliveryRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "delivery.proto",
}
