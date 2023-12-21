// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tide/v1/tide_service.proto

package tideV1

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
	TideService_GetDayPredictTide_FullMethodName      = "/tide.v1.TideService/GetDayPredictTide"
	TideService_GetLocationPredictTide_FullMethodName = "/tide.v1.TideService/GetLocationPredictTide"
	TideService_GetLocationList_FullMethodName        = "/tide.v1.TideService/GetLocationList"
)

// TideServiceClient is the client API for TideService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TideServiceClient interface {
	// GetPredictTide returns the predicted tide of the given day.
	GetDayPredictTide(ctx context.Context, in *GetDayPredictTideRequest, opts ...grpc.CallOption) (*GetDayPredictTideResponse, error)
	// GetPredictTide returns the predicted tide of the given location.
	GetLocationPredictTide(ctx context.Context, in *GetLocationPredictTideRequest, opts ...grpc.CallOption) (*GetLocationPredictTideResponse, error)
	// GetLocationList returns location list.
	GetLocationList(ctx context.Context, in *GetLocationListRequest, opts ...grpc.CallOption) (*GetLocationListResponse, error)
}

type tideServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTideServiceClient(cc grpc.ClientConnInterface) TideServiceClient {
	return &tideServiceClient{cc}
}

func (c *tideServiceClient) GetDayPredictTide(ctx context.Context, in *GetDayPredictTideRequest, opts ...grpc.CallOption) (*GetDayPredictTideResponse, error) {
	out := new(GetDayPredictTideResponse)
	err := c.cc.Invoke(ctx, TideService_GetDayPredictTide_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tideServiceClient) GetLocationPredictTide(ctx context.Context, in *GetLocationPredictTideRequest, opts ...grpc.CallOption) (*GetLocationPredictTideResponse, error) {
	out := new(GetLocationPredictTideResponse)
	err := c.cc.Invoke(ctx, TideService_GetLocationPredictTide_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tideServiceClient) GetLocationList(ctx context.Context, in *GetLocationListRequest, opts ...grpc.CallOption) (*GetLocationListResponse, error) {
	out := new(GetLocationListResponse)
	err := c.cc.Invoke(ctx, TideService_GetLocationList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TideServiceServer is the server API for TideService service.
// All implementations must embed UnimplementedTideServiceServer
// for forward compatibility
type TideServiceServer interface {
	// GetPredictTide returns the predicted tide of the given day.
	GetDayPredictTide(context.Context, *GetDayPredictTideRequest) (*GetDayPredictTideResponse, error)
	// GetPredictTide returns the predicted tide of the given location.
	GetLocationPredictTide(context.Context, *GetLocationPredictTideRequest) (*GetLocationPredictTideResponse, error)
	// GetLocationList returns location list.
	GetLocationList(context.Context, *GetLocationListRequest) (*GetLocationListResponse, error)
	mustEmbedUnimplementedTideServiceServer()
}

// UnimplementedTideServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTideServiceServer struct {
}

func (UnimplementedTideServiceServer) GetDayPredictTide(context.Context, *GetDayPredictTideRequest) (*GetDayPredictTideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDayPredictTide not implemented")
}
func (UnimplementedTideServiceServer) GetLocationPredictTide(context.Context, *GetLocationPredictTideRequest) (*GetLocationPredictTideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocationPredictTide not implemented")
}
func (UnimplementedTideServiceServer) GetLocationList(context.Context, *GetLocationListRequest) (*GetLocationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocationList not implemented")
}
func (UnimplementedTideServiceServer) mustEmbedUnimplementedTideServiceServer() {}

// UnsafeTideServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TideServiceServer will
// result in compilation errors.
type UnsafeTideServiceServer interface {
	mustEmbedUnimplementedTideServiceServer()
}

func RegisterTideServiceServer(s grpc.ServiceRegistrar, srv TideServiceServer) {
	s.RegisterService(&TideService_ServiceDesc, srv)
}

func _TideService_GetDayPredictTide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDayPredictTideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TideServiceServer).GetDayPredictTide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TideService_GetDayPredictTide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TideServiceServer).GetDayPredictTide(ctx, req.(*GetDayPredictTideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TideService_GetLocationPredictTide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationPredictTideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TideServiceServer).GetLocationPredictTide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TideService_GetLocationPredictTide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TideServiceServer).GetLocationPredictTide(ctx, req.(*GetLocationPredictTideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TideService_GetLocationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TideServiceServer).GetLocationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TideService_GetLocationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TideServiceServer).GetLocationList(ctx, req.(*GetLocationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TideService_ServiceDesc is the grpc.ServiceDesc for TideService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TideService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tide.v1.TideService",
	HandlerType: (*TideServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDayPredictTide",
			Handler:    _TideService_GetDayPredictTide_Handler,
		},
		{
			MethodName: "GetLocationPredictTide",
			Handler:    _TideService_GetLocationPredictTide_Handler,
		},
		{
			MethodName: "GetLocationList",
			Handler:    _TideService_GetLocationList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tide/v1/tide_service.proto",
}
