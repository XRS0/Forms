// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: jwt.proto

package gen

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	CreateJWTToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
	VerifyJWTToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CreateJWTToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, "/jwt.AuthService/CreateJWTToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyJWTToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error) {
	out := new(VerifyTokenResponse)
	err := c.cc.Invoke(ctx, "/jwt.AuthService/VerifyJWTToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	CreateJWTToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	VerifyJWTToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) CreateJWTToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJWTToken not implemented")
}
func (UnimplementedAuthServiceServer) VerifyJWTToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyJWTToken not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_CreateJWTToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateJWTToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwt.AuthService/CreateJWTToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateJWTToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyJWTToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyJWTToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwt.AuthService/VerifyJWTToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyJWTToken(ctx, req.(*VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jwt.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJWTToken",
			Handler:    _AuthService_CreateJWTToken_Handler,
		},
		{
			MethodName: "VerifyJWTToken",
			Handler:    _AuthService_VerifyJWTToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jwt.proto",
}
