// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: chromadb/proto/logservice.proto

package logservicepb

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
	LogService_PushLogs_FullMethodName                      = "/chroma.LogService/PushLogs"
	LogService_PullLogs_FullMethodName                      = "/chroma.LogService/PullLogs"
	LogService_GetAllCollectionInfoToCompact_FullMethodName = "/chroma.LogService/GetAllCollectionInfoToCompact"
	LogService_UpdateCollectionLogOffset_FullMethodName     = "/chroma.LogService/UpdateCollectionLogOffset"
)

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	PushLogs(ctx context.Context, in *PushLogsRequest, opts ...grpc.CallOption) (*PushLogsResponse, error)
	PullLogs(ctx context.Context, in *PullLogsRequest, opts ...grpc.CallOption) (*PullLogsResponse, error)
	GetAllCollectionInfoToCompact(ctx context.Context, in *GetAllCollectionInfoToCompactRequest, opts ...grpc.CallOption) (*GetAllCollectionInfoToCompactResponse, error)
	UpdateCollectionLogOffset(ctx context.Context, in *UpdateCollectionLogOffsetRequest, opts ...grpc.CallOption) (*UpdateCollectionLogOffsetResponse, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) PushLogs(ctx context.Context, in *PushLogsRequest, opts ...grpc.CallOption) (*PushLogsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PushLogsResponse)
	err := c.cc.Invoke(ctx, LogService_PushLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) PullLogs(ctx context.Context, in *PullLogsRequest, opts ...grpc.CallOption) (*PullLogsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PullLogsResponse)
	err := c.cc.Invoke(ctx, LogService_PullLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) GetAllCollectionInfoToCompact(ctx context.Context, in *GetAllCollectionInfoToCompactRequest, opts ...grpc.CallOption) (*GetAllCollectionInfoToCompactResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllCollectionInfoToCompactResponse)
	err := c.cc.Invoke(ctx, LogService_GetAllCollectionInfoToCompact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) UpdateCollectionLogOffset(ctx context.Context, in *UpdateCollectionLogOffsetRequest, opts ...grpc.CallOption) (*UpdateCollectionLogOffsetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCollectionLogOffsetResponse)
	err := c.cc.Invoke(ctx, LogService_UpdateCollectionLogOffset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility.
type LogServiceServer interface {
	PushLogs(context.Context, *PushLogsRequest) (*PushLogsResponse, error)
	PullLogs(context.Context, *PullLogsRequest) (*PullLogsResponse, error)
	GetAllCollectionInfoToCompact(context.Context, *GetAllCollectionInfoToCompactRequest) (*GetAllCollectionInfoToCompactResponse, error)
	UpdateCollectionLogOffset(context.Context, *UpdateCollectionLogOffsetRequest) (*UpdateCollectionLogOffsetResponse, error)
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogServiceServer struct{}

func (UnimplementedLogServiceServer) PushLogs(context.Context, *PushLogsRequest) (*PushLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushLogs not implemented")
}
func (UnimplementedLogServiceServer) PullLogs(context.Context, *PullLogsRequest) (*PullLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullLogs not implemented")
}
func (UnimplementedLogServiceServer) GetAllCollectionInfoToCompact(context.Context, *GetAllCollectionInfoToCompactRequest) (*GetAllCollectionInfoToCompactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCollectionInfoToCompact not implemented")
}
func (UnimplementedLogServiceServer) UpdateCollectionLogOffset(context.Context, *UpdateCollectionLogOffsetRequest) (*UpdateCollectionLogOffsetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollectionLogOffset not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}
func (UnimplementedLogServiceServer) testEmbeddedByValue()                    {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	// If the following call pancis, it indicates UnimplementedLogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_PushLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).PushLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_PushLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).PushLogs(ctx, req.(*PushLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_PullLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).PullLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_PullLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).PullLogs(ctx, req.(*PullLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_GetAllCollectionInfoToCompact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCollectionInfoToCompactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).GetAllCollectionInfoToCompact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_GetAllCollectionInfoToCompact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).GetAllCollectionInfoToCompact(ctx, req.(*GetAllCollectionInfoToCompactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_UpdateCollectionLogOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollectionLogOffsetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).UpdateCollectionLogOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_UpdateCollectionLogOffset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).UpdateCollectionLogOffset(ctx, req.(*UpdateCollectionLogOffsetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chroma.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushLogs",
			Handler:    _LogService_PushLogs_Handler,
		},
		{
			MethodName: "PullLogs",
			Handler:    _LogService_PullLogs_Handler,
		},
		{
			MethodName: "GetAllCollectionInfoToCompact",
			Handler:    _LogService_GetAllCollectionInfoToCompact_Handler,
		},
		{
			MethodName: "UpdateCollectionLogOffset",
			Handler:    _LogService_UpdateCollectionLogOffset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chromadb/proto/logservice.proto",
}
