// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: meshops.proto

package meshes

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

// MeshServiceClient is the client API for MeshService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeshServiceClient interface {
	MeshName(ctx context.Context, in *MeshNameRequest, opts ...grpc.CallOption) (*MeshNameResponse, error)
	MeshVersions(ctx context.Context, in *MeshVersionsRequest, opts ...grpc.CallOption) (*MeshVersionsResponse, error)
	ApplyOperation(ctx context.Context, in *ApplyRuleRequest, opts ...grpc.CallOption) (*ApplyRuleResponse, error)
	SupportedOperations(ctx context.Context, in *SupportedOperationsRequest, opts ...grpc.CallOption) (*SupportedOperationsResponse, error)
	StreamEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (MeshService_StreamEventsClient, error)
	ProcessOAM(ctx context.Context, in *ProcessOAMRequest, opts ...grpc.CallOption) (*ProcessOAMResponse, error)
	ComponentInfo(ctx context.Context, in *ComponentInfoRequest, opts ...grpc.CallOption) (*ComponentInfoResponse, error)
}

type meshServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeshServiceClient(cc grpc.ClientConnInterface) MeshServiceClient {
	return &meshServiceClient{cc}
}

func (c *meshServiceClient) MeshName(ctx context.Context, in *MeshNameRequest, opts ...grpc.CallOption) (*MeshNameResponse, error) {
	out := new(MeshNameResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/MeshName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) MeshVersions(ctx context.Context, in *MeshVersionsRequest, opts ...grpc.CallOption) (*MeshVersionsResponse, error) {
	out := new(MeshVersionsResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/MeshVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) ApplyOperation(ctx context.Context, in *ApplyRuleRequest, opts ...grpc.CallOption) (*ApplyRuleResponse, error) {
	out := new(ApplyRuleResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/ApplyOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) SupportedOperations(ctx context.Context, in *SupportedOperationsRequest, opts ...grpc.CallOption) (*SupportedOperationsResponse, error) {
	out := new(SupportedOperationsResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/SupportedOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) StreamEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (MeshService_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MeshService_ServiceDesc.Streams[0], "/meshes.MeshService/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &meshServiceStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MeshService_StreamEventsClient interface {
	Recv() (*EventsResponse, error)
	grpc.ClientStream
}

type meshServiceStreamEventsClient struct {
	grpc.ClientStream
}

func (x *meshServiceStreamEventsClient) Recv() (*EventsResponse, error) {
	m := new(EventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *meshServiceClient) ProcessOAM(ctx context.Context, in *ProcessOAMRequest, opts ...grpc.CallOption) (*ProcessOAMResponse, error) {
	out := new(ProcessOAMResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/ProcessOAM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshServiceClient) ComponentInfo(ctx context.Context, in *ComponentInfoRequest, opts ...grpc.CallOption) (*ComponentInfoResponse, error) {
	out := new(ComponentInfoResponse)
	err := c.cc.Invoke(ctx, "/meshes.MeshService/ComponentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeshServiceServer is the server API for MeshService service.
// All implementations must embed UnimplementedMeshServiceServer
// for forward compatibility
type MeshServiceServer interface {
	MeshName(context.Context, *MeshNameRequest) (*MeshNameResponse, error)
	MeshVersions(context.Context, *MeshVersionsRequest) (*MeshVersionsResponse, error)
	ApplyOperation(context.Context, *ApplyRuleRequest) (*ApplyRuleResponse, error)
	SupportedOperations(context.Context, *SupportedOperationsRequest) (*SupportedOperationsResponse, error)
	StreamEvents(*EventsRequest, MeshService_StreamEventsServer) error
	ProcessOAM(context.Context, *ProcessOAMRequest) (*ProcessOAMResponse, error)
	ComponentInfo(context.Context, *ComponentInfoRequest) (*ComponentInfoResponse, error)
	mustEmbedUnimplementedMeshServiceServer()
}

// UnimplementedMeshServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMeshServiceServer struct {
}

func (UnimplementedMeshServiceServer) MeshName(context.Context, *MeshNameRequest) (*MeshNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MeshName not implemented")
}
func (UnimplementedMeshServiceServer) MeshVersions(context.Context, *MeshVersionsRequest) (*MeshVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MeshVersions not implemented")
}
func (UnimplementedMeshServiceServer) ApplyOperation(context.Context, *ApplyRuleRequest) (*ApplyRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyOperation not implemented")
}
func (UnimplementedMeshServiceServer) SupportedOperations(context.Context, *SupportedOperationsRequest) (*SupportedOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupportedOperations not implemented")
}
func (UnimplementedMeshServiceServer) StreamEvents(*EventsRequest, MeshService_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}
func (UnimplementedMeshServiceServer) ProcessOAM(context.Context, *ProcessOAMRequest) (*ProcessOAMResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessOAM not implemented")
}
func (UnimplementedMeshServiceServer) ComponentInfo(context.Context, *ComponentInfoRequest) (*ComponentInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComponentInfo not implemented")
}
func (UnimplementedMeshServiceServer) mustEmbedUnimplementedMeshServiceServer() {}

// UnsafeMeshServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeshServiceServer will
// result in compilation errors.
type UnsafeMeshServiceServer interface {
	mustEmbedUnimplementedMeshServiceServer()
}

func RegisterMeshServiceServer(s grpc.ServiceRegistrar, srv MeshServiceServer) {
	s.RegisterService(&MeshService_ServiceDesc, srv)
}

func _MeshService_MeshName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeshNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).MeshName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/MeshName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).MeshName(ctx, req.(*MeshNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_MeshVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeshVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).MeshVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/MeshVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).MeshVersions(ctx, req.(*MeshVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_ApplyOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).ApplyOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/ApplyOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).ApplyOperation(ctx, req.(*ApplyRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_SupportedOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportedOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).SupportedOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/SupportedOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).SupportedOperations(ctx, req.(*SupportedOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MeshServiceServer).StreamEvents(m, &meshServiceStreamEventsServer{stream})
}

type MeshService_StreamEventsServer interface {
	Send(*EventsResponse) error
	grpc.ServerStream
}

type meshServiceStreamEventsServer struct {
	grpc.ServerStream
}

func (x *meshServiceStreamEventsServer) Send(m *EventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MeshService_ProcessOAM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessOAMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).ProcessOAM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/ProcessOAM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).ProcessOAM(ctx, req.(*ProcessOAMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshService_ComponentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComponentInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServiceServer).ComponentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshes.MeshService/ComponentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServiceServer).ComponentInfo(ctx, req.(*ComponentInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeshService_ServiceDesc is the grpc.ServiceDesc for MeshService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeshService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meshes.MeshService",
	HandlerType: (*MeshServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MeshName",
			Handler:    _MeshService_MeshName_Handler,
		},
		{
			MethodName: "MeshVersions",
			Handler:    _MeshService_MeshVersions_Handler,
		},
		{
			MethodName: "ApplyOperation",
			Handler:    _MeshService_ApplyOperation_Handler,
		},
		{
			MethodName: "SupportedOperations",
			Handler:    _MeshService_SupportedOperations_Handler,
		},
		{
			MethodName: "ProcessOAM",
			Handler:    _MeshService_ProcessOAM_Handler,
		},
		{
			MethodName: "ComponentInfo",
			Handler:    _MeshService_ComponentInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _MeshService_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "meshops.proto",
}
