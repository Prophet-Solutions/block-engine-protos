// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: block_engine.proto

package block_engine

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
	BlockEngineValidator_SubscribePackets_FullMethodName       = "/block_engine.BlockEngineValidator/SubscribePackets"
	BlockEngineValidator_SubscribeBundles_FullMethodName       = "/block_engine.BlockEngineValidator/SubscribeBundles"
	BlockEngineValidator_GetBlockBuilderFeeInfo_FullMethodName = "/block_engine.BlockEngineValidator/GetBlockBuilderFeeInfo"
)

// BlockEngineValidatorClient is the client API for BlockEngineValidator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// / Validators can connect to Block Engines to receive packets and bundles.
type BlockEngineValidatorClient interface {
	// / Validators can subscribe to the block engine to receive a stream of packets
	SubscribePackets(ctx context.Context, in *SubscribePacketsRequest, opts ...grpc.CallOption) (BlockEngineValidator_SubscribePacketsClient, error)
	// / Validators can subscribe to the block engine to receive a stream of simulated and profitable bundles
	SubscribeBundles(ctx context.Context, in *SubscribeBundlesRequest, opts ...grpc.CallOption) (BlockEngineValidator_SubscribeBundlesClient, error)
	// Block builders can optionally collect fees. This returns fee information if a block builder wants to
	// collect one.
	GetBlockBuilderFeeInfo(ctx context.Context, in *BlockBuilderFeeInfoRequest, opts ...grpc.CallOption) (*BlockBuilderFeeInfoResponse, error)
}

type blockEngineValidatorClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockEngineValidatorClient(cc grpc.ClientConnInterface) BlockEngineValidatorClient {
	return &blockEngineValidatorClient{cc}
}

func (c *blockEngineValidatorClient) SubscribePackets(ctx context.Context, in *SubscribePacketsRequest, opts ...grpc.CallOption) (BlockEngineValidator_SubscribePacketsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlockEngineValidator_ServiceDesc.Streams[0], BlockEngineValidator_SubscribePackets_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &blockEngineValidatorSubscribePacketsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockEngineValidator_SubscribePacketsClient interface {
	Recv() (*SubscribePacketsResponse, error)
	grpc.ClientStream
}

type blockEngineValidatorSubscribePacketsClient struct {
	grpc.ClientStream
}

func (x *blockEngineValidatorSubscribePacketsClient) Recv() (*SubscribePacketsResponse, error) {
	m := new(SubscribePacketsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockEngineValidatorClient) SubscribeBundles(ctx context.Context, in *SubscribeBundlesRequest, opts ...grpc.CallOption) (BlockEngineValidator_SubscribeBundlesClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlockEngineValidator_ServiceDesc.Streams[1], BlockEngineValidator_SubscribeBundles_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &blockEngineValidatorSubscribeBundlesClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockEngineValidator_SubscribeBundlesClient interface {
	Recv() (*SubscribeBundlesResponse, error)
	grpc.ClientStream
}

type blockEngineValidatorSubscribeBundlesClient struct {
	grpc.ClientStream
}

func (x *blockEngineValidatorSubscribeBundlesClient) Recv() (*SubscribeBundlesResponse, error) {
	m := new(SubscribeBundlesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockEngineValidatorClient) GetBlockBuilderFeeInfo(ctx context.Context, in *BlockBuilderFeeInfoRequest, opts ...grpc.CallOption) (*BlockBuilderFeeInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockBuilderFeeInfoResponse)
	err := c.cc.Invoke(ctx, BlockEngineValidator_GetBlockBuilderFeeInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockEngineValidatorServer is the server API for BlockEngineValidator service.
// All implementations must embed UnimplementedBlockEngineValidatorServer
// for forward compatibility
//
// / Validators can connect to Block Engines to receive packets and bundles.
type BlockEngineValidatorServer interface {
	// / Validators can subscribe to the block engine to receive a stream of packets
	SubscribePackets(*SubscribePacketsRequest, BlockEngineValidator_SubscribePacketsServer) error
	// / Validators can subscribe to the block engine to receive a stream of simulated and profitable bundles
	SubscribeBundles(*SubscribeBundlesRequest, BlockEngineValidator_SubscribeBundlesServer) error
	// Block builders can optionally collect fees. This returns fee information if a block builder wants to
	// collect one.
	GetBlockBuilderFeeInfo(context.Context, *BlockBuilderFeeInfoRequest) (*BlockBuilderFeeInfoResponse, error)
	mustEmbedUnimplementedBlockEngineValidatorServer()
}

// UnimplementedBlockEngineValidatorServer must be embedded to have forward compatible implementations.
type UnimplementedBlockEngineValidatorServer struct {
}

func (UnimplementedBlockEngineValidatorServer) SubscribePackets(*SubscribePacketsRequest, BlockEngineValidator_SubscribePacketsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribePackets not implemented")
}
func (UnimplementedBlockEngineValidatorServer) SubscribeBundles(*SubscribeBundlesRequest, BlockEngineValidator_SubscribeBundlesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeBundles not implemented")
}
func (UnimplementedBlockEngineValidatorServer) GetBlockBuilderFeeInfo(context.Context, *BlockBuilderFeeInfoRequest) (*BlockBuilderFeeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockBuilderFeeInfo not implemented")
}
func (UnimplementedBlockEngineValidatorServer) mustEmbedUnimplementedBlockEngineValidatorServer() {}

// UnsafeBlockEngineValidatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockEngineValidatorServer will
// result in compilation errors.
type UnsafeBlockEngineValidatorServer interface {
	mustEmbedUnimplementedBlockEngineValidatorServer()
}

func RegisterBlockEngineValidatorServer(s grpc.ServiceRegistrar, srv BlockEngineValidatorServer) {
	s.RegisterService(&BlockEngineValidator_ServiceDesc, srv)
}

func _BlockEngineValidator_SubscribePackets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribePacketsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockEngineValidatorServer).SubscribePackets(m, &blockEngineValidatorSubscribePacketsServer{ServerStream: stream})
}

type BlockEngineValidator_SubscribePacketsServer interface {
	Send(*SubscribePacketsResponse) error
	grpc.ServerStream
}

type blockEngineValidatorSubscribePacketsServer struct {
	grpc.ServerStream
}

func (x *blockEngineValidatorSubscribePacketsServer) Send(m *SubscribePacketsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BlockEngineValidator_SubscribeBundles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeBundlesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockEngineValidatorServer).SubscribeBundles(m, &blockEngineValidatorSubscribeBundlesServer{ServerStream: stream})
}

type BlockEngineValidator_SubscribeBundlesServer interface {
	Send(*SubscribeBundlesResponse) error
	grpc.ServerStream
}

type blockEngineValidatorSubscribeBundlesServer struct {
	grpc.ServerStream
}

func (x *blockEngineValidatorSubscribeBundlesServer) Send(m *SubscribeBundlesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BlockEngineValidator_GetBlockBuilderFeeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockBuilderFeeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockEngineValidatorServer).GetBlockBuilderFeeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockEngineValidator_GetBlockBuilderFeeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockEngineValidatorServer).GetBlockBuilderFeeInfo(ctx, req.(*BlockBuilderFeeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockEngineValidator_ServiceDesc is the grpc.ServiceDesc for BlockEngineValidator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockEngineValidator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "block_engine.BlockEngineValidator",
	HandlerType: (*BlockEngineValidatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockBuilderFeeInfo",
			Handler:    _BlockEngineValidator_GetBlockBuilderFeeInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribePackets",
			Handler:       _BlockEngineValidator_SubscribePackets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeBundles",
			Handler:       _BlockEngineValidator_SubscribeBundles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "block_engine.proto",
}

const (
	BlockEngineRelayer_SubscribeAccountsOfInterest_FullMethodName = "/block_engine.BlockEngineRelayer/SubscribeAccountsOfInterest"
	BlockEngineRelayer_SubscribeProgramsOfInterest_FullMethodName = "/block_engine.BlockEngineRelayer/SubscribeProgramsOfInterest"
	BlockEngineRelayer_StartExpiringPacketStream_FullMethodName   = "/block_engine.BlockEngineRelayer/StartExpiringPacketStream"
)

// BlockEngineRelayerClient is the client API for BlockEngineRelayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// / Relayers can forward packets to Block Engines.
// / Block Engines provide an AccountsOfInterest field to only send transactions that are of interest.
type BlockEngineRelayerClient interface {
	// / The block engine feeds accounts of interest (AOI) updates to the relayer periodically.
	// / For all transactions the relayer receives, it forwards transactions to the block engine which write-lock
	// / any of the accounts in the AOI.
	SubscribeAccountsOfInterest(ctx context.Context, in *AccountsOfInterestRequest, opts ...grpc.CallOption) (BlockEngineRelayer_SubscribeAccountsOfInterestClient, error)
	SubscribeProgramsOfInterest(ctx context.Context, in *ProgramsOfInterestRequest, opts ...grpc.CallOption) (BlockEngineRelayer_SubscribeProgramsOfInterestClient, error)
	// Validators can subscribe to packets from the relayer and receive a multiplexed signal that contains a mixture
	// of packets and heartbeats.
	// NOTE: This is a bi-directional stream due to a bug with how Envoy handles half closed client-side streams.
	// The issue is being tracked here: https://github.com/envoyproxy/envoy/issues/22748. In the meantime, the
	// server will stream heartbeats to clients at some reasonable cadence.
	StartExpiringPacketStream(ctx context.Context, opts ...grpc.CallOption) (BlockEngineRelayer_StartExpiringPacketStreamClient, error)
}

type blockEngineRelayerClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockEngineRelayerClient(cc grpc.ClientConnInterface) BlockEngineRelayerClient {
	return &blockEngineRelayerClient{cc}
}

func (c *blockEngineRelayerClient) SubscribeAccountsOfInterest(ctx context.Context, in *AccountsOfInterestRequest, opts ...grpc.CallOption) (BlockEngineRelayer_SubscribeAccountsOfInterestClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlockEngineRelayer_ServiceDesc.Streams[0], BlockEngineRelayer_SubscribeAccountsOfInterest_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &blockEngineRelayerSubscribeAccountsOfInterestClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockEngineRelayer_SubscribeAccountsOfInterestClient interface {
	Recv() (*AccountsOfInterestUpdate, error)
	grpc.ClientStream
}

type blockEngineRelayerSubscribeAccountsOfInterestClient struct {
	grpc.ClientStream
}

func (x *blockEngineRelayerSubscribeAccountsOfInterestClient) Recv() (*AccountsOfInterestUpdate, error) {
	m := new(AccountsOfInterestUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockEngineRelayerClient) SubscribeProgramsOfInterest(ctx context.Context, in *ProgramsOfInterestRequest, opts ...grpc.CallOption) (BlockEngineRelayer_SubscribeProgramsOfInterestClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlockEngineRelayer_ServiceDesc.Streams[1], BlockEngineRelayer_SubscribeProgramsOfInterest_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &blockEngineRelayerSubscribeProgramsOfInterestClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockEngineRelayer_SubscribeProgramsOfInterestClient interface {
	Recv() (*ProgramsOfInterestUpdate, error)
	grpc.ClientStream
}

type blockEngineRelayerSubscribeProgramsOfInterestClient struct {
	grpc.ClientStream
}

func (x *blockEngineRelayerSubscribeProgramsOfInterestClient) Recv() (*ProgramsOfInterestUpdate, error) {
	m := new(ProgramsOfInterestUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockEngineRelayerClient) StartExpiringPacketStream(ctx context.Context, opts ...grpc.CallOption) (BlockEngineRelayer_StartExpiringPacketStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlockEngineRelayer_ServiceDesc.Streams[2], BlockEngineRelayer_StartExpiringPacketStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &blockEngineRelayerStartExpiringPacketStreamClient{ClientStream: stream}
	return x, nil
}

type BlockEngineRelayer_StartExpiringPacketStreamClient interface {
	Send(*PacketBatchUpdate) error
	Recv() (*StartExpiringPacketStreamResponse, error)
	grpc.ClientStream
}

type blockEngineRelayerStartExpiringPacketStreamClient struct {
	grpc.ClientStream
}

func (x *blockEngineRelayerStartExpiringPacketStreamClient) Send(m *PacketBatchUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blockEngineRelayerStartExpiringPacketStreamClient) Recv() (*StartExpiringPacketStreamResponse, error) {
	m := new(StartExpiringPacketStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlockEngineRelayerServer is the server API for BlockEngineRelayer service.
// All implementations must embed UnimplementedBlockEngineRelayerServer
// for forward compatibility
//
// / Relayers can forward packets to Block Engines.
// / Block Engines provide an AccountsOfInterest field to only send transactions that are of interest.
type BlockEngineRelayerServer interface {
	// / The block engine feeds accounts of interest (AOI) updates to the relayer periodically.
	// / For all transactions the relayer receives, it forwards transactions to the block engine which write-lock
	// / any of the accounts in the AOI.
	SubscribeAccountsOfInterest(*AccountsOfInterestRequest, BlockEngineRelayer_SubscribeAccountsOfInterestServer) error
	SubscribeProgramsOfInterest(*ProgramsOfInterestRequest, BlockEngineRelayer_SubscribeProgramsOfInterestServer) error
	// Validators can subscribe to packets from the relayer and receive a multiplexed signal that contains a mixture
	// of packets and heartbeats.
	// NOTE: This is a bi-directional stream due to a bug with how Envoy handles half closed client-side streams.
	// The issue is being tracked here: https://github.com/envoyproxy/envoy/issues/22748. In the meantime, the
	// server will stream heartbeats to clients at some reasonable cadence.
	StartExpiringPacketStream(BlockEngineRelayer_StartExpiringPacketStreamServer) error
	mustEmbedUnimplementedBlockEngineRelayerServer()
}

// UnimplementedBlockEngineRelayerServer must be embedded to have forward compatible implementations.
type UnimplementedBlockEngineRelayerServer struct {
}

func (UnimplementedBlockEngineRelayerServer) SubscribeAccountsOfInterest(*AccountsOfInterestRequest, BlockEngineRelayer_SubscribeAccountsOfInterestServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeAccountsOfInterest not implemented")
}
func (UnimplementedBlockEngineRelayerServer) SubscribeProgramsOfInterest(*ProgramsOfInterestRequest, BlockEngineRelayer_SubscribeProgramsOfInterestServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeProgramsOfInterest not implemented")
}
func (UnimplementedBlockEngineRelayerServer) StartExpiringPacketStream(BlockEngineRelayer_StartExpiringPacketStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method StartExpiringPacketStream not implemented")
}
func (UnimplementedBlockEngineRelayerServer) mustEmbedUnimplementedBlockEngineRelayerServer() {}

// UnsafeBlockEngineRelayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockEngineRelayerServer will
// result in compilation errors.
type UnsafeBlockEngineRelayerServer interface {
	mustEmbedUnimplementedBlockEngineRelayerServer()
}

func RegisterBlockEngineRelayerServer(s grpc.ServiceRegistrar, srv BlockEngineRelayerServer) {
	s.RegisterService(&BlockEngineRelayer_ServiceDesc, srv)
}

func _BlockEngineRelayer_SubscribeAccountsOfInterest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AccountsOfInterestRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockEngineRelayerServer).SubscribeAccountsOfInterest(m, &blockEngineRelayerSubscribeAccountsOfInterestServer{ServerStream: stream})
}

type BlockEngineRelayer_SubscribeAccountsOfInterestServer interface {
	Send(*AccountsOfInterestUpdate) error
	grpc.ServerStream
}

type blockEngineRelayerSubscribeAccountsOfInterestServer struct {
	grpc.ServerStream
}

func (x *blockEngineRelayerSubscribeAccountsOfInterestServer) Send(m *AccountsOfInterestUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _BlockEngineRelayer_SubscribeProgramsOfInterest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProgramsOfInterestRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockEngineRelayerServer).SubscribeProgramsOfInterest(m, &blockEngineRelayerSubscribeProgramsOfInterestServer{ServerStream: stream})
}

type BlockEngineRelayer_SubscribeProgramsOfInterestServer interface {
	Send(*ProgramsOfInterestUpdate) error
	grpc.ServerStream
}

type blockEngineRelayerSubscribeProgramsOfInterestServer struct {
	grpc.ServerStream
}

func (x *blockEngineRelayerSubscribeProgramsOfInterestServer) Send(m *ProgramsOfInterestUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _BlockEngineRelayer_StartExpiringPacketStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlockEngineRelayerServer).StartExpiringPacketStream(&blockEngineRelayerStartExpiringPacketStreamServer{ServerStream: stream})
}

type BlockEngineRelayer_StartExpiringPacketStreamServer interface {
	Send(*StartExpiringPacketStreamResponse) error
	Recv() (*PacketBatchUpdate, error)
	grpc.ServerStream
}

type blockEngineRelayerStartExpiringPacketStreamServer struct {
	grpc.ServerStream
}

func (x *blockEngineRelayerStartExpiringPacketStreamServer) Send(m *StartExpiringPacketStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blockEngineRelayerStartExpiringPacketStreamServer) Recv() (*PacketBatchUpdate, error) {
	m := new(PacketBatchUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlockEngineRelayer_ServiceDesc is the grpc.ServiceDesc for BlockEngineRelayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockEngineRelayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "block_engine.BlockEngineRelayer",
	HandlerType: (*BlockEngineRelayerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeAccountsOfInterest",
			Handler:       _BlockEngineRelayer_SubscribeAccountsOfInterest_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeProgramsOfInterest",
			Handler:       _BlockEngineRelayer_SubscribeProgramsOfInterest_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StartExpiringPacketStream",
			Handler:       _BlockEngineRelayer_StartExpiringPacketStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "block_engine.proto",
}