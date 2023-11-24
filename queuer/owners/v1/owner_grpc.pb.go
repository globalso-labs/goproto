// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: queuer/owners/v1/owner.proto

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

const (
	OwnerService_CreateOwner_FullMethodName = "/queuer.owners.v1.OwnerService/CreateOwner"
	OwnerService_GetOwner_FullMethodName    = "/queuer.owners.v1.OwnerService/GetOwner"
	OwnerService_ListQueues_FullMethodName  = "/queuer.owners.v1.OwnerService/ListQueues"
)

// OwnerServiceClient is the client API for OwnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OwnerServiceClient interface {
	CreateOwner(ctx context.Context, in *CreateOwnerRequest, opts ...grpc.CallOption) (*CreateOwnerResponse, error)
	GetOwner(ctx context.Context, in *GetOwnerRequest, opts ...grpc.CallOption) (*GetOwnerResponse, error)
	ListQueues(ctx context.Context, in *ListQueuesRequest, opts ...grpc.CallOption) (*ListQueuesResponse, error)
}

type ownerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOwnerServiceClient(cc grpc.ClientConnInterface) OwnerServiceClient {
	return &ownerServiceClient{cc}
}

func (c *ownerServiceClient) CreateOwner(ctx context.Context, in *CreateOwnerRequest, opts ...grpc.CallOption) (*CreateOwnerResponse, error) {
	out := new(CreateOwnerResponse)
	err := c.cc.Invoke(ctx, OwnerService_CreateOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownerServiceClient) GetOwner(ctx context.Context, in *GetOwnerRequest, opts ...grpc.CallOption) (*GetOwnerResponse, error) {
	out := new(GetOwnerResponse)
	err := c.cc.Invoke(ctx, OwnerService_GetOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownerServiceClient) ListQueues(ctx context.Context, in *ListQueuesRequest, opts ...grpc.CallOption) (*ListQueuesResponse, error) {
	out := new(ListQueuesResponse)
	err := c.cc.Invoke(ctx, OwnerService_ListQueues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerServiceServer is the server API for OwnerService service.
// All implementations must embed UnimplementedOwnerServiceServer
// for forward compatibility
type OwnerServiceServer interface {
	CreateOwner(context.Context, *CreateOwnerRequest) (*CreateOwnerResponse, error)
	GetOwner(context.Context, *GetOwnerRequest) (*GetOwnerResponse, error)
	ListQueues(context.Context, *ListQueuesRequest) (*ListQueuesResponse, error)
	mustEmbedUnimplementedOwnerServiceServer()
}

// UnimplementedOwnerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOwnerServiceServer struct {
}

func (UnimplementedOwnerServiceServer) CreateOwner(context.Context, *CreateOwnerRequest) (*CreateOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOwner not implemented")
}
func (UnimplementedOwnerServiceServer) GetOwner(context.Context, *GetOwnerRequest) (*GetOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwner not implemented")
}
func (UnimplementedOwnerServiceServer) ListQueues(context.Context, *ListQueuesRequest) (*ListQueuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQueues not implemented")
}
func (UnimplementedOwnerServiceServer) mustEmbedUnimplementedOwnerServiceServer() {}

// UnsafeOwnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OwnerServiceServer will
// result in compilation errors.
type UnsafeOwnerServiceServer interface {
	mustEmbedUnimplementedOwnerServiceServer()
}

func RegisterOwnerServiceServer(s grpc.ServiceRegistrar, srv OwnerServiceServer) {
	s.RegisterService(&OwnerService_ServiceDesc, srv)
}

func _OwnerService_CreateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).CreateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnerService_CreateOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).CreateOwner(ctx, req.(*CreateOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnerService_GetOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).GetOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnerService_GetOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).GetOwner(ctx, req.(*GetOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnerService_ListQueues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQueuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).ListQueues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnerService_ListQueues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).ListQueues(ctx, req.(*ListQueuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OwnerService_ServiceDesc is the grpc.ServiceDesc for OwnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OwnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "queuer.owners.v1.OwnerService",
	HandlerType: (*OwnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOwner",
			Handler:    _OwnerService_CreateOwner_Handler,
		},
		{
			MethodName: "GetOwner",
			Handler:    _OwnerService_GetOwner_Handler,
		},
		{
			MethodName: "ListQueues",
			Handler:    _OwnerService_ListQueues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queuer/owners/v1/owner.proto",
}
