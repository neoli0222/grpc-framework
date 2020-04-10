// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/echo/echo.proto

package echo // import "github.com/lpy-neo/grpc-framework/api/proto/echo"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/lpy-neo/grpc-framework/api/proto/shared/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	Echo(ctx context.Context, in *types.StringMessage, opts ...grpc.CallOption) (*types.StringMessage, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *types.StringMessage, opts ...grpc.CallOption) (*types.StringMessage, error) {
	out := new(types.StringMessage)
	err := c.cc.Invoke(ctx, "/echo.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	Echo(context.Context, *types.StringMessage) (*types.StringMessage, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*types.StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/echo/echo.proto",
}

func init() { proto.RegisterFile("api/proto/echo/echo.proto", fileDescriptor_echo_57a37a224ca38d8c) }

var fileDescriptor_echo_57a37a224ca38d8c = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x4d, 0xce, 0x80, 0x10, 0x7a, 0x60, 0xbe, 0x10, 0x0b,
	0x88, 0x2d, 0x25, 0x93, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x0f, 0x52, 0x97, 0x98, 0x97, 0x97,
	0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x51, 0x23, 0xa5, 0x8d, 0xd0, 0x5e, 0x9c, 0x91,
	0x58, 0x94, 0x9a, 0xa2, 0x5f, 0x52, 0x59, 0x90, 0x5a, 0xac, 0x5f, 0x5c, 0x52, 0x94, 0x99, 0x97,
	0x1e, 0x9f, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x0a, 0x51, 0x6c, 0x14, 0xcc, 0xc5, 0xed, 0x9a,
	0x9c, 0x91, 0x1f, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xc2, 0xc5, 0x02, 0xe2, 0x0a,
	0x89, 0xe8, 0x81, 0xf5, 0xe8, 0x05, 0x83, 0xf5, 0xf8, 0x42, 0xb4, 0x48, 0x61, 0x15, 0x55, 0x12,
	0x68, 0xba, 0xfc, 0x64, 0x32, 0x13, 0x97, 0x10, 0x87, 0x7e, 0x99, 0x21, 0xd8, 0xad, 0x4e, 0x46,
	0x51, 0x06, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x39, 0x05, 0x95,
	0xba, 0x79, 0xa9, 0xf9, 0xfa, 0xe9, 0x45, 0x05, 0xc9, 0xba, 0x69, 0x45, 0x89, 0xb9, 0xa9, 0xe5,
	0xf9, 0x45, 0xd9, 0xfa, 0xa8, 0x9e, 0x4c, 0x62, 0x03, 0xb3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x64, 0x47, 0xdb, 0x5d, 0xfd, 0x00, 0x00, 0x00,
}
