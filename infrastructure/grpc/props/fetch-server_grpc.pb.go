// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: infrastructure/grpc/props/fetch-server.proto

package grpcapi

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
	StockConfigurator_UpdateStockConfigOne_FullMethodName  = "/grpcapi.StockConfigurator/UpdateStockConfigOne"
	StockConfigurator_UpdateStockConfigMany_FullMethodName = "/grpcapi.StockConfigurator/UpdateStockConfigMany"
	StockConfigurator_GetStockConfigOne_FullMethodName     = "/grpcapi.StockConfigurator/GetStockConfigOne"
	StockConfigurator_GetStockConfigAll_FullMethodName     = "/grpcapi.StockConfigurator/GetStockConfigAll"
)

// StockConfiguratorClient is the client API for StockConfigurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockConfiguratorClient interface {
	UpdateStockConfigOne(ctx context.Context, in *StockConfig, opts ...grpc.CallOption) (*ReturnMessage, error)
	UpdateStockConfigMany(ctx context.Context, in *StockConfigList, opts ...grpc.CallOption) (*ReturnMessageList, error)
	GetStockConfigOne(ctx context.Context, in *StockId, opts ...grpc.CallOption) (*StockConfig, error)
	GetStockConfigAll(ctx context.Context, in *Null, opts ...grpc.CallOption) (*StockConfigList, error)
}

type stockConfiguratorClient struct {
	cc grpc.ClientConnInterface
}

func NewStockConfiguratorClient(cc grpc.ClientConnInterface) StockConfiguratorClient {
	return &stockConfiguratorClient{cc}
}

func (c *stockConfiguratorClient) UpdateStockConfigOne(ctx context.Context, in *StockConfig, opts ...grpc.CallOption) (*ReturnMessage, error) {
	out := new(ReturnMessage)
	err := c.cc.Invoke(ctx, StockConfigurator_UpdateStockConfigOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockConfiguratorClient) UpdateStockConfigMany(ctx context.Context, in *StockConfigList, opts ...grpc.CallOption) (*ReturnMessageList, error) {
	out := new(ReturnMessageList)
	err := c.cc.Invoke(ctx, StockConfigurator_UpdateStockConfigMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockConfiguratorClient) GetStockConfigOne(ctx context.Context, in *StockId, opts ...grpc.CallOption) (*StockConfig, error) {
	out := new(StockConfig)
	err := c.cc.Invoke(ctx, StockConfigurator_GetStockConfigOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockConfiguratorClient) GetStockConfigAll(ctx context.Context, in *Null, opts ...grpc.CallOption) (*StockConfigList, error) {
	out := new(StockConfigList)
	err := c.cc.Invoke(ctx, StockConfigurator_GetStockConfigAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockConfiguratorServer is the server API for StockConfigurator service.
// All implementations must embed UnimplementedStockConfiguratorServer
// for forward compatibility
type StockConfiguratorServer interface {
	UpdateStockConfigOne(context.Context, *StockConfig) (*ReturnMessage, error)
	UpdateStockConfigMany(context.Context, *StockConfigList) (*ReturnMessageList, error)
	GetStockConfigOne(context.Context, *StockId) (*StockConfig, error)
	GetStockConfigAll(context.Context, *Null) (*StockConfigList, error)
	mustEmbedUnimplementedStockConfiguratorServer()
}

// UnimplementedStockConfiguratorServer must be embedded to have forward compatible implementations.
type UnimplementedStockConfiguratorServer struct {
}

func (UnimplementedStockConfiguratorServer) UpdateStockConfigOne(context.Context, *StockConfig) (*ReturnMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStockConfigOne not implemented")
}
func (UnimplementedStockConfiguratorServer) UpdateStockConfigMany(context.Context, *StockConfigList) (*ReturnMessageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStockConfigMany not implemented")
}
func (UnimplementedStockConfiguratorServer) GetStockConfigOne(context.Context, *StockId) (*StockConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockConfigOne not implemented")
}
func (UnimplementedStockConfiguratorServer) GetStockConfigAll(context.Context, *Null) (*StockConfigList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockConfigAll not implemented")
}
func (UnimplementedStockConfiguratorServer) mustEmbedUnimplementedStockConfiguratorServer() {}

// UnsafeStockConfiguratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockConfiguratorServer will
// result in compilation errors.
type UnsafeStockConfiguratorServer interface {
	mustEmbedUnimplementedStockConfiguratorServer()
}

func RegisterStockConfiguratorServer(s grpc.ServiceRegistrar, srv StockConfiguratorServer) {
	s.RegisterService(&StockConfigurator_ServiceDesc, srv)
}

func _StockConfigurator_UpdateStockConfigOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockConfiguratorServer).UpdateStockConfigOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockConfigurator_UpdateStockConfigOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockConfiguratorServer).UpdateStockConfigOne(ctx, req.(*StockConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockConfigurator_UpdateStockConfigMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockConfigList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockConfiguratorServer).UpdateStockConfigMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockConfigurator_UpdateStockConfigMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockConfiguratorServer).UpdateStockConfigMany(ctx, req.(*StockConfigList))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockConfigurator_GetStockConfigOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockConfiguratorServer).GetStockConfigOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockConfigurator_GetStockConfigOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockConfiguratorServer).GetStockConfigOne(ctx, req.(*StockId))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockConfigurator_GetStockConfigAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockConfiguratorServer).GetStockConfigAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockConfigurator_GetStockConfigAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockConfiguratorServer).GetStockConfigAll(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

// StockConfigurator_ServiceDesc is the grpc.ServiceDesc for StockConfigurator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockConfigurator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.StockConfigurator",
	HandlerType: (*StockConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateStockConfigOne",
			Handler:    _StockConfigurator_UpdateStockConfigOne_Handler,
		},
		{
			MethodName: "UpdateStockConfigMany",
			Handler:    _StockConfigurator_UpdateStockConfigMany_Handler,
		},
		{
			MethodName: "GetStockConfigOne",
			Handler:    _StockConfigurator_GetStockConfigOne_Handler,
		},
		{
			MethodName: "GetStockConfigAll",
			Handler:    _StockConfigurator_GetStockConfigAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infrastructure/grpc/props/fetch-server.proto",
}
