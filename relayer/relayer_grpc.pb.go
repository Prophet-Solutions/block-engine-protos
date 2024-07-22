// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: relayer.proto

package relayer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Relayer_GetTpuConfigs_FullMethodName    = "/relayer.Relayer/GetTpuConfigs"
	Relayer_SubscribePackets_FullMethodName = "/relayer.Relayer/SubscribePackets"
)

// RelayerClient is the client API for Relayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// / Relayers offer a TPU and TPU forward proxy for Solana validators.
// / Validators can connect and fetch the TPU configuration for the relayer and start to advertise the
// / relayer's information in gossip.
// / They can also subscribe to packets which arrived on the TPU ports at the relayer
type RelayerClient interface {
	// The relayer has TPU and TPU forward sockets that validators can leverage.
	// A validator can fetch this config and change its TPU and TPU forward port in gossip.
	GetTpuConfigs(ctx context.Context, in *GetTpuConfigsRequest, opts ...grpc.CallOption) (*GetTpuConfigsResponse, error)
	// Validators can subscribe to packets from the relayer and receive a multiplexed signal that contains a mixture
	// of packets and heartbeats
	SubscribePackets(ctx context.Context, in *SubscribePacketsRequest, opts ...grpc.CallOption) (Relayer_SubscribePacketsClient, error)
}

type relayerClient struct {
	cc grpc.ClientConnInterface
}

func NewRelayerClient(cc grpc.ClientConnInterface) RelayerClient {
	return &relayerClient{cc}
}

func (c *relayerClient) GetTpuConfigs(ctx context.Context, in *GetTpuConfigsRequest, opts ...grpc.CallOption) (*GetTpuConfigsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTpuConfigsResponse)
	err := c.cc.Invoke(ctx, Relayer_GetTpuConfigs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relayerClient) SubscribePackets(ctx context.Context, in *SubscribePacketsRequest, opts ...grpc.CallOption) (Relayer_SubscribePacketsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Relayer_ServiceDesc.Streams[0], Relayer_SubscribePackets_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &relayerSubscribePacketsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Relayer_SubscribePacketsClient interface {
	Recv() (*SubscribePacketsResponse, error)
	grpc.ClientStream
}

type relayerSubscribePacketsClient struct {
	grpc.ClientStream
}

func (x *relayerSubscribePacketsClient) Recv() (*SubscribePacketsResponse, error) {
	m := new(SubscribePacketsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RelayerServer is the server API for Relayer service.
// All implementations must embed UnimplementedRelayerServer
// for forward compatibility
//
// / Relayers offer a TPU and TPU forward proxy for Solana validators.
// / Validators can connect and fetch the TPU configuration for the relayer and start to advertise the
// / relayer's information in gossip.
// / They can also subscribe to packets which arrived on the TPU ports at the relayer
type RelayerServer interface {
	// The relayer has TPU and TPU forward sockets that validators can leverage.
	// A validator can fetch this config and change its TPU and TPU forward port in gossip.
	GetTpuConfigs(context.Context, *GetTpuConfigsRequest) (*GetTpuConfigsResponse, error)
	// Validators can subscribe to packets from the relayer and receive a multiplexed signal that contains a mixture
	// of packets and heartbeats
	SubscribePackets(*SubscribePacketsRequest, Relayer_SubscribePacketsServer) error
	mustEmbedUnimplementedRelayerServer()
}

// UnimplementedRelayerServer must be embedded to have forward compatible implementations.
type UnimplementedRelayerServer struct {
}

func (UnimplementedRelayerServer) GetTpuConfigs(context.Context, *GetTpuConfigsRequest) (*GetTpuConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTpuConfigs not implemented")
}
func (UnimplementedRelayerServer) SubscribePackets(*SubscribePacketsRequest, Relayer_SubscribePacketsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribePackets not implemented")
}
func (UnimplementedRelayerServer) mustEmbedUnimplementedRelayerServer() {}

// UnsafeRelayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelayerServer will
// result in compilation errors.
type UnsafeRelayerServer interface {
	mustEmbedUnimplementedRelayerServer()
}

func RegisterRelayerServer(s grpc.ServiceRegistrar, srv RelayerServer) {
	s.RegisterService(&Relayer_ServiceDesc, srv)
}

func _Relayer_GetTpuConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTpuConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayerServer).GetTpuConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Relayer_GetTpuConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayerServer).GetTpuConfigs(ctx, req.(*GetTpuConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relayer_SubscribePackets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribePacketsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RelayerServer).SubscribePackets(m, &relayerSubscribePacketsServer{ServerStream: stream})
}

type Relayer_SubscribePacketsServer interface {
	Send(*SubscribePacketsResponse) error
	grpc.ServerStream
}

type relayerSubscribePacketsServer struct {
	grpc.ServerStream
}

func (x *relayerSubscribePacketsServer) Send(m *SubscribePacketsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Relayer_ServiceDesc is the grpc.ServiceDesc for Relayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Relayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relayer.Relayer",
	HandlerType: (*RelayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTpuConfigs",
			Handler:    _Relayer_GetTpuConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribePackets",
			Handler:       _Relayer_SubscribePackets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "relayer.proto",
}
