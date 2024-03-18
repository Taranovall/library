// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: jwt.proto

package jwt

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

// JwtServiceClient is the client API for JwtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JwtServiceClient interface {
	GenerateToken(ctx context.Context, in *JwtRequest, opts ...grpc.CallOption) (*JwtResponse, error)
}

type jwtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJwtServiceClient(cc grpc.ClientConnInterface) JwtServiceClient {
	return &jwtServiceClient{cc}
}

func (c *jwtServiceClient) GenerateToken(ctx context.Context, in *JwtRequest, opts ...grpc.CallOption) (*JwtResponse, error) {
	out := new(JwtResponse)
	err := c.cc.Invoke(ctx, "/jwt.JwtService/generateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JwtServiceServer is the server API for JwtService service.
// All implementations must embed UnimplementedJwtServiceServer
// for forward compatibility
type JwtServiceServer interface {
	GenerateToken(context.Context, *JwtRequest) (*JwtResponse, error)
	mustEmbedUnimplementedJwtServiceServer()
}

// UnimplementedJwtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJwtServiceServer struct {
}

func (UnimplementedJwtServiceServer) GenerateToken(context.Context, *JwtRequest) (*JwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedJwtServiceServer) mustEmbedUnimplementedJwtServiceServer() {}

// UnsafeJwtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JwtServiceServer will
// result in compilation errors.
type UnsafeJwtServiceServer interface {
	mustEmbedUnimplementedJwtServiceServer()
}

func RegisterJwtServiceServer(s grpc.ServiceRegistrar, srv JwtServiceServer) {
	s.RegisterService(&JwtService_ServiceDesc, srv)
}

func _JwtService_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwt.JwtService/generateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).GenerateToken(ctx, req.(*JwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JwtService_ServiceDesc is the grpc.ServiceDesc for JwtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JwtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jwt.JwtService",
	HandlerType: (*JwtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "generateToken",
			Handler:    _JwtService_GenerateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jwt.proto",
}
