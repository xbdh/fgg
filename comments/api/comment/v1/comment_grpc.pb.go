// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	CreateComment(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
	ListTweetComment(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) CreateComment(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/comment.v1.Comment/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) ListTweetComment(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/comment.v1.Comment/ListTweetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	CreateComment(context.Context, *CreateRequest) (*CreateReply, error)
	ListTweetComment(context.Context, *ListRequest) (*ListReply, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) CreateComment(context.Context, *CreateRequest) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentServer) ListTweetComment(context.Context, *ListRequest) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTweetComment not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.v1.Comment/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).CreateComment(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_ListTweetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).ListTweetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.v1.Comment/ListTweetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).ListTweetComment(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.v1.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _Comment_CreateComment_Handler,
		},
		{
			MethodName: "ListTweetComment",
			Handler:    _Comment_ListTweetComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/comment/v1/comment.proto",
}