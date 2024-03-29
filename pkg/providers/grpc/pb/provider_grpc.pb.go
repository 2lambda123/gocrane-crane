// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.2
// source: provider.proto

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

// RealtimeClient is the client API for Realtime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealtimeClient interface {
	QueryLatestTimeSeries(ctx context.Context, in *QueryTimeSeriesRequest, opts ...grpc.CallOption) (*QueryTimeSeriesResponse, error)
}

type realtimeClient struct {
	cc grpc.ClientConnInterface
}

func NewRealtimeClient(cc grpc.ClientConnInterface) RealtimeClient {
	return &realtimeClient{cc}
}

func (c *realtimeClient) QueryLatestTimeSeries(ctx context.Context, in *QueryTimeSeriesRequest, opts ...grpc.CallOption) (*QueryTimeSeriesResponse, error) {
	out := new(QueryTimeSeriesResponse)
	err := c.cc.Invoke(ctx, "/Realtime/QueryLatestTimeSeries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealtimeServer is the server API for Realtime service.
// All implementations must embed UnimplementedRealtimeServer
// for forward compatibility
type RealtimeServer interface {
	QueryLatestTimeSeries(context.Context, *QueryTimeSeriesRequest) (*QueryTimeSeriesResponse, error)
	mustEmbedUnimplementedRealtimeServer()
}

// UnimplementedRealtimeServer must be embedded to have forward compatible implementations.
type UnimplementedRealtimeServer struct {
}

func (UnimplementedRealtimeServer) QueryLatestTimeSeries(context.Context, *QueryTimeSeriesRequest) (*QueryTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLatestTimeSeries not implemented")
}
func (UnimplementedRealtimeServer) mustEmbedUnimplementedRealtimeServer() {}

// UnsafeRealtimeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealtimeServer will
// result in compilation errors.
type UnsafeRealtimeServer interface {
	mustEmbedUnimplementedRealtimeServer()
}

func RegisterRealtimeServer(s grpc.ServiceRegistrar, srv RealtimeServer) {
	s.RegisterService(&Realtime_ServiceDesc, srv)
}

func _Realtime_QueryLatestTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealtimeServer).QueryLatestTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Realtime/QueryLatestTimeSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealtimeServer).QueryLatestTimeSeries(ctx, req.(*QueryTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Realtime_ServiceDesc is the grpc.ServiceDesc for Realtime service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Realtime_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Realtime",
	HandlerType: (*RealtimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryLatestTimeSeries",
			Handler:    _Realtime_QueryLatestTimeSeries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}

// HistoryClient is the client API for History service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoryClient interface {
	QueryTimeSeries(ctx context.Context, in *QueryTimeSeriesRequest, opts ...grpc.CallOption) (*QueryTimeSeriesResponse, error)
}

type historyClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryClient(cc grpc.ClientConnInterface) HistoryClient {
	return &historyClient{cc}
}

func (c *historyClient) QueryTimeSeries(ctx context.Context, in *QueryTimeSeriesRequest, opts ...grpc.CallOption) (*QueryTimeSeriesResponse, error) {
	out := new(QueryTimeSeriesResponse)
	err := c.cc.Invoke(ctx, "/History/QueryTimeSeries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryServer is the server API for History service.
// All implementations must embed UnimplementedHistoryServer
// for forward compatibility
type HistoryServer interface {
	QueryTimeSeries(context.Context, *QueryTimeSeriesRequest) (*QueryTimeSeriesResponse, error)
	mustEmbedUnimplementedHistoryServer()
}

// UnimplementedHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedHistoryServer struct {
}

func (UnimplementedHistoryServer) QueryTimeSeries(context.Context, *QueryTimeSeriesRequest) (*QueryTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTimeSeries not implemented")
}
func (UnimplementedHistoryServer) mustEmbedUnimplementedHistoryServer() {}

// UnsafeHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoryServer will
// result in compilation errors.
type UnsafeHistoryServer interface {
	mustEmbedUnimplementedHistoryServer()
}

func RegisterHistoryServer(s grpc.ServiceRegistrar, srv HistoryServer) {
	s.RegisterService(&History_ServiceDesc, srv)
}

func _History_QueryTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).QueryTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/History/QueryTimeSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).QueryTimeSeries(ctx, req.(*QueryTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// History_ServiceDesc is the grpc.ServiceDesc for History service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var History_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "History",
	HandlerType: (*HistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryTimeSeries",
			Handler:    _History_QueryTimeSeries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}
