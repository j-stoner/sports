// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package imageboard

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ImageBoardClient is the client API for ImageBoard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageBoardClient interface {
	SetStatus(ctx context.Context, in *SetStatusReq, opts ...grpc.CallOption) (*empty.Empty, error)
	GetStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResp, error)
	Jump(ctx context.Context, in *JumpReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type imageBoardClient struct {
	cc grpc.ClientConnInterface
}

func NewImageBoardClient(cc grpc.ClientConnInterface) ImageBoardClient {
	return &imageBoardClient{cc}
}

func (c *imageBoardClient) SetStatus(ctx context.Context, in *SetStatusReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/imageboard.v1.ImageBoard/SetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageBoardClient) GetStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResp, error) {
	out := new(StatusResp)
	err := c.cc.Invoke(ctx, "/imageboard.v1.ImageBoard/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageBoardClient) Jump(ctx context.Context, in *JumpReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/imageboard.v1.ImageBoard/Jump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageBoardServer is the server API for ImageBoard service.
// All implementations must embed UnimplementedImageBoardServer
// for forward compatibility
type ImageBoardServer interface {
	SetStatus(context.Context, *SetStatusReq) (*empty.Empty, error)
	GetStatus(context.Context, *empty.Empty) (*StatusResp, error)
	Jump(context.Context, *JumpReq) (*empty.Empty, error)
	mustEmbedUnimplementedImageBoardServer()
}

// UnimplementedImageBoardServer must be embedded to have forward compatible implementations.
type UnimplementedImageBoardServer struct {
}

func (UnimplementedImageBoardServer) SetStatus(context.Context, *SetStatusReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatus not implemented")
}
func (UnimplementedImageBoardServer) GetStatus(context.Context, *empty.Empty) (*StatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedImageBoardServer) Jump(context.Context, *JumpReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Jump not implemented")
}
func (UnimplementedImageBoardServer) mustEmbedUnimplementedImageBoardServer() {}

// UnsafeImageBoardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageBoardServer will
// result in compilation errors.
type UnsafeImageBoardServer interface {
	mustEmbedUnimplementedImageBoardServer()
}

func RegisterImageBoardServer(s grpc.ServiceRegistrar, srv ImageBoardServer) {
	s.RegisterService(&ImageBoard_ServiceDesc, srv)
}

func _ImageBoard_SetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageBoardServer).SetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imageboard.v1.ImageBoard/SetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageBoardServer).SetStatus(ctx, req.(*SetStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageBoard_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageBoardServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imageboard.v1.ImageBoard/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageBoardServer).GetStatus(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageBoard_Jump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JumpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageBoardServer).Jump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imageboard.v1.ImageBoard/Jump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageBoardServer).Jump(ctx, req.(*JumpReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageBoard_ServiceDesc is the grpc.ServiceDesc for ImageBoard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageBoard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "imageboard.v1.ImageBoard",
	HandlerType: (*ImageBoardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetStatus",
			Handler:    _ImageBoard_SetStatus_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _ImageBoard_GetStatus_Handler,
		},
		{
			MethodName: "Jump",
			Handler:    _ImageBoard_Jump_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imageboard/imageboard.proto",
}
