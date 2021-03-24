// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package weather

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

// WeatherClient is the client API for Weather service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherClient interface {
	// 获取某个城市的天气
	GetWeatherByCity(ctx context.Context, in *GetWeatherByCityReq, opts ...grpc.CallOption) (*GetWeatherByCityRes, error)
	// 获取某个城市的生活指数
	GetWeatherLifeByCity(ctx context.Context, in *GetWeatherLifeByCityReq, opts ...grpc.CallOption) (*GetWeatherLifeByCityRes, error)
}

type weatherClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherClient(cc grpc.ClientConnInterface) WeatherClient {
	return &weatherClient{cc}
}

func (c *weatherClient) GetWeatherByCity(ctx context.Context, in *GetWeatherByCityReq, opts ...grpc.CallOption) (*GetWeatherByCityRes, error) {
	out := new(GetWeatherByCityRes)
	err := c.cc.Invoke(ctx, "/proto.Weather/GetWeatherByCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherClient) GetWeatherLifeByCity(ctx context.Context, in *GetWeatherLifeByCityReq, opts ...grpc.CallOption) (*GetWeatherLifeByCityRes, error) {
	out := new(GetWeatherLifeByCityRes)
	err := c.cc.Invoke(ctx, "/proto.Weather/GetWeatherLifeByCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServer is the server API for Weather service.
// All implementations must embed UnimplementedWeatherServer
// for forward compatibility
type WeatherServer interface {
	// 获取某个城市的天气
	GetWeatherByCity(context.Context, *GetWeatherByCityReq) (*GetWeatherByCityRes, error)
	// 获取某个城市的生活指数
	GetWeatherLifeByCity(context.Context, *GetWeatherLifeByCityReq) (*GetWeatherLifeByCityRes, error)
	mustEmbedUnimplementedWeatherServer()
}

// UnimplementedWeatherServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServer struct {
}

func (UnimplementedWeatherServer) GetWeatherByCity(context.Context, *GetWeatherByCityReq) (*GetWeatherByCityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeatherByCity not implemented")
}
func (UnimplementedWeatherServer) GetWeatherLifeByCity(context.Context, *GetWeatherLifeByCityReq) (*GetWeatherLifeByCityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeatherLifeByCity not implemented")
}
func (UnimplementedWeatherServer) mustEmbedUnimplementedWeatherServer() {}

// UnsafeWeatherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServer will
// result in compilation errors.
type UnsafeWeatherServer interface {
	mustEmbedUnimplementedWeatherServer()
}

func RegisterWeatherServer(s grpc.ServiceRegistrar, srv WeatherServer) {
	s.RegisterService(&Weather_ServiceDesc, srv)
}

func _Weather_GetWeatherByCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeatherByCityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServer).GetWeatherByCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Weather/GetWeatherByCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServer).GetWeatherByCity(ctx, req.(*GetWeatherByCityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weather_GetWeatherLifeByCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeatherLifeByCityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServer).GetWeatherLifeByCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Weather/GetWeatherLifeByCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServer).GetWeatherLifeByCity(ctx, req.(*GetWeatherLifeByCityReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Weather_ServiceDesc is the grpc.ServiceDesc for Weather service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Weather_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Weather",
	HandlerType: (*WeatherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeatherByCity",
			Handler:    _Weather_GetWeatherByCity_Handler,
		},
		{
			MethodName: "GetWeatherLifeByCity",
			Handler:    _Weather_GetWeatherLifeByCity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "growth/weather/weather.proto",
}
