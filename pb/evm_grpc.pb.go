// Copyright Rivtower Technologies LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: vm/evm.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RPCService_GetTransactionReceipt_FullMethodName = "/evm.RPCService/GetTransactionReceipt"
	RPCService_GetCode_FullMethodName               = "/evm.RPCService/GetCode"
	RPCService_GetBalance_FullMethodName            = "/evm.RPCService/GetBalance"
	RPCService_GetTransactionCount_FullMethodName   = "/evm.RPCService/GetTransactionCount"
	RPCService_GetAbi_FullMethodName                = "/evm.RPCService/GetAbi"
	RPCService_EstimateQuota_FullMethodName         = "/evm.RPCService/EstimateQuota"
	RPCService_GetReceiptProof_FullMethodName       = "/evm.RPCService/GetReceiptProof"
	RPCService_GetRootsInfo_FullMethodName          = "/evm.RPCService/GetRootsInfo"
	RPCService_GetStorageAt_FullMethodName          = "/evm.RPCService/GetStorageAt"
)

// RPCServiceClient is the client API for RPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCServiceClient interface {
	GetTransactionReceipt(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Receipt, error)
	GetCode(ctx context.Context, in *GetCodeRequest, opts ...grpc.CallOption) (*ByteCode, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*Balance, error)
	GetTransactionCount(ctx context.Context, in *GetTransactionCountRequest, opts ...grpc.CallOption) (*Nonce, error)
	GetAbi(ctx context.Context, in *GetAbiRequest, opts ...grpc.CallOption) (*ByteAbi, error)
	EstimateQuota(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*ByteQuota, error)
	GetReceiptProof(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*ReceiptProof, error)
	GetRootsInfo(ctx context.Context, in *BlockNumber, opts ...grpc.CallOption) (*RootsInfo, error)
	GetStorageAt(ctx context.Context, in *GetStorageAtRequest, opts ...grpc.CallOption) (*Hash, error)
}

type rPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCServiceClient(cc grpc.ClientConnInterface) RPCServiceClient {
	return &rPCServiceClient{cc}
}

func (c *rPCServiceClient) GetTransactionReceipt(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Receipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Receipt)
	err := c.cc.Invoke(ctx, RPCService_GetTransactionReceipt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) GetCode(ctx context.Context, in *GetCodeRequest, opts ...grpc.CallOption) (*ByteCode, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ByteCode)
	err := c.cc.Invoke(ctx, RPCService_GetCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*Balance, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Balance)
	err := c.cc.Invoke(ctx, RPCService_GetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) GetTransactionCount(ctx context.Context, in *GetTransactionCountRequest, opts ...grpc.CallOption) (*Nonce, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Nonce)
	err := c.cc.Invoke(ctx, RPCService_GetTransactionCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) GetAbi(ctx context.Context, in *GetAbiRequest, opts ...grpc.CallOption) (*ByteAbi, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ByteAbi)
	err := c.cc.Invoke(ctx, RPCService_GetAbi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) EstimateQuota(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*ByteQuota, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ByteQuota)
	err := c.cc.Invoke(ctx, RPCService_EstimateQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) GetReceiptProof(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*ReceiptProof, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReceiptProof)
	err := c.cc.Invoke(ctx, RPCService_GetReceiptProof_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) GetRootsInfo(ctx context.Context, in *BlockNumber, opts ...grpc.CallOption) (*RootsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RootsInfo)
	err := c.cc.Invoke(ctx, RPCService_GetRootsInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCServiceClient) GetStorageAt(ctx context.Context, in *GetStorageAtRequest, opts ...grpc.CallOption) (*Hash, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Hash)
	err := c.cc.Invoke(ctx, RPCService_GetStorageAt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServiceServer is the server API for RPCService service.
// All implementations must embed UnimplementedRPCServiceServer
// for forward compatibility.
type RPCServiceServer interface {
	GetTransactionReceipt(context.Context, *Hash) (*Receipt, error)
	GetCode(context.Context, *GetCodeRequest) (*ByteCode, error)
	GetBalance(context.Context, *GetBalanceRequest) (*Balance, error)
	GetTransactionCount(context.Context, *GetTransactionCountRequest) (*Nonce, error)
	GetAbi(context.Context, *GetAbiRequest) (*ByteAbi, error)
	EstimateQuota(context.Context, *CallRequest) (*ByteQuota, error)
	GetReceiptProof(context.Context, *Hash) (*ReceiptProof, error)
	GetRootsInfo(context.Context, *BlockNumber) (*RootsInfo, error)
	GetStorageAt(context.Context, *GetStorageAtRequest) (*Hash, error)
	mustEmbedUnimplementedRPCServiceServer()
}

// UnimplementedRPCServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRPCServiceServer struct{}

func (UnimplementedRPCServiceServer) GetTransactionReceipt(context.Context, *Hash) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionReceipt not implemented")
}
func (UnimplementedRPCServiceServer) GetCode(context.Context, *GetCodeRequest) (*ByteCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCode not implemented")
}
func (UnimplementedRPCServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*Balance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedRPCServiceServer) GetTransactionCount(context.Context, *GetTransactionCountRequest) (*Nonce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionCount not implemented")
}
func (UnimplementedRPCServiceServer) GetAbi(context.Context, *GetAbiRequest) (*ByteAbi, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAbi not implemented")
}
func (UnimplementedRPCServiceServer) EstimateQuota(context.Context, *CallRequest) (*ByteQuota, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateQuota not implemented")
}
func (UnimplementedRPCServiceServer) GetReceiptProof(context.Context, *Hash) (*ReceiptProof, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceiptProof not implemented")
}
func (UnimplementedRPCServiceServer) GetRootsInfo(context.Context, *BlockNumber) (*RootsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRootsInfo not implemented")
}
func (UnimplementedRPCServiceServer) GetStorageAt(context.Context, *GetStorageAtRequest) (*Hash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorageAt not implemented")
}
func (UnimplementedRPCServiceServer) mustEmbedUnimplementedRPCServiceServer() {}
func (UnimplementedRPCServiceServer) testEmbeddedByValue()                    {}

// UnsafeRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServiceServer will
// result in compilation errors.
type UnsafeRPCServiceServer interface {
	mustEmbedUnimplementedRPCServiceServer()
}

func RegisterRPCServiceServer(s grpc.ServiceRegistrar, srv RPCServiceServer) {
	// If the following call pancis, it indicates UnimplementedRPCServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RPCService_ServiceDesc, srv)
}

func _RPCService_GetTransactionReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).GetTransactionReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_GetTransactionReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).GetTransactionReceipt(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_GetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).GetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_GetCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).GetCode(ctx, req.(*GetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_GetTransactionCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).GetTransactionCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_GetTransactionCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).GetTransactionCount(ctx, req.(*GetTransactionCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_GetAbi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAbiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).GetAbi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_GetAbi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).GetAbi(ctx, req.(*GetAbiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_EstimateQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).EstimateQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_EstimateQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).EstimateQuota(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_GetReceiptProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).GetReceiptProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_GetReceiptProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).GetReceiptProof(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_GetRootsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).GetRootsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_GetRootsInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).GetRootsInfo(ctx, req.(*BlockNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCService_GetStorageAt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageAtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).GetStorageAt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCService_GetStorageAt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).GetStorageAt(ctx, req.(*GetStorageAtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCService_ServiceDesc is the grpc.ServiceDesc for RPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "evm.RPCService",
	HandlerType: (*RPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransactionReceipt",
			Handler:    _RPCService_GetTransactionReceipt_Handler,
		},
		{
			MethodName: "GetCode",
			Handler:    _RPCService_GetCode_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _RPCService_GetBalance_Handler,
		},
		{
			MethodName: "GetTransactionCount",
			Handler:    _RPCService_GetTransactionCount_Handler,
		},
		{
			MethodName: "GetAbi",
			Handler:    _RPCService_GetAbi_Handler,
		},
		{
			MethodName: "EstimateQuota",
			Handler:    _RPCService_EstimateQuota_Handler,
		},
		{
			MethodName: "GetReceiptProof",
			Handler:    _RPCService_GetReceiptProof_Handler,
		},
		{
			MethodName: "GetRootsInfo",
			Handler:    _RPCService_GetRootsInfo_Handler,
		},
		{
			MethodName: "GetStorageAt",
			Handler:    _RPCService_GetStorageAt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vm/evm.proto",
}
