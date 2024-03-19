// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v1/machine.proto

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
	Machine_BootDevice_FullMethodName = "/github.com.tinkerbell.pbnj.api.v1.Machine/BootDevice"
	Machine_Power_FullMethodName      = "/github.com.tinkerbell.pbnj.api.v1.Machine/Power"
)

// MachineClient is the client API for Machine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MachineClient interface {
	BootDevice(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceResponse, error)
	Power(ctx context.Context, in *PowerRequest, opts ...grpc.CallOption) (*PowerResponse, error)
}

type machineClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineClient(cc grpc.ClientConnInterface) MachineClient {
	return &machineClient{cc}
}

func (c *machineClient) BootDevice(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceResponse, error) {
	out := new(DeviceResponse)
	err := c.cc.Invoke(ctx, Machine_BootDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) Power(ctx context.Context, in *PowerRequest, opts ...grpc.CallOption) (*PowerResponse, error) {
	out := new(PowerResponse)
	err := c.cc.Invoke(ctx, Machine_Power_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineServer is the server API for Machine service.
// All implementations must embed UnimplementedMachineServer
// for forward compatibility
type MachineServer interface {
	BootDevice(context.Context, *DeviceRequest) (*DeviceResponse, error)
	Power(context.Context, *PowerRequest) (*PowerResponse, error)
	mustEmbedUnimplementedMachineServer()
}

// UnimplementedMachineServer must be embedded to have forward compatible implementations.
type UnimplementedMachineServer struct {
}

func (UnimplementedMachineServer) BootDevice(context.Context, *DeviceRequest) (*DeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BootDevice not implemented")
}
func (UnimplementedMachineServer) Power(context.Context, *PowerRequest) (*PowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Power not implemented")
}
func (UnimplementedMachineServer) mustEmbedUnimplementedMachineServer() {}

// UnsafeMachineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MachineServer will
// result in compilation errors.
type UnsafeMachineServer interface {
	mustEmbedUnimplementedMachineServer()
}

func RegisterMachineServer(s grpc.ServiceRegistrar, srv MachineServer) {
	s.RegisterService(&Machine_ServiceDesc, srv)
}

func _Machine_BootDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).BootDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_BootDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).BootDevice(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_Power_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).Power(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_Power_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).Power(ctx, req.(*PowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Machine_ServiceDesc is the grpc.ServiceDesc for Machine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Machine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.tinkerbell.pbnj.api.v1.Machine",
	HandlerType: (*MachineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BootDevice",
			Handler:    _Machine_BootDevice_Handler,
		},
		{
			MethodName: "Power",
			Handler:    _Machine_Power_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/machine.proto",
}
