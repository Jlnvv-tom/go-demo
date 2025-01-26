// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: article.proto

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

const (
	Article_AddArticle_FullMethodName     = "/pb.article/AddArticle"
	Article_UpdateArticle_FullMethodName  = "/pb.article/UpdateArticle"
	Article_DelArticle_FullMethodName     = "/pb.article/DelArticle"
	Article_GetArticleById_FullMethodName = "/pb.article/GetArticleById"
	Article_SearchArticle_FullMethodName  = "/pb.article/SearchArticle"
)

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleClient interface {
	// -----------------------article-----------------------
	AddArticle(ctx context.Context, in *AddArticleReq, opts ...grpc.CallOption) (*AddArticleResp, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleReq, opts ...grpc.CallOption) (*UpdateArticleResp, error)
	DelArticle(ctx context.Context, in *DelArticleReq, opts ...grpc.CallOption) (*DelArticleResp, error)
	GetArticleById(ctx context.Context, in *GetArticleByIdReq, opts ...grpc.CallOption) (*GetArticleByIdResp, error)
	SearchArticle(ctx context.Context, in *SearchArticleReq, opts ...grpc.CallOption) (*SearchArticleResp, error)
}

type articleClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleClient(cc grpc.ClientConnInterface) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) AddArticle(ctx context.Context, in *AddArticleReq, opts ...grpc.CallOption) (*AddArticleResp, error) {
	out := new(AddArticleResp)
	err := c.cc.Invoke(ctx, Article_AddArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) UpdateArticle(ctx context.Context, in *UpdateArticleReq, opts ...grpc.CallOption) (*UpdateArticleResp, error) {
	out := new(UpdateArticleResp)
	err := c.cc.Invoke(ctx, Article_UpdateArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) DelArticle(ctx context.Context, in *DelArticleReq, opts ...grpc.CallOption) (*DelArticleResp, error) {
	out := new(DelArticleResp)
	err := c.cc.Invoke(ctx, Article_DelArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArticleById(ctx context.Context, in *GetArticleByIdReq, opts ...grpc.CallOption) (*GetArticleByIdResp, error) {
	out := new(GetArticleByIdResp)
	err := c.cc.Invoke(ctx, Article_GetArticleById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) SearchArticle(ctx context.Context, in *SearchArticleReq, opts ...grpc.CallOption) (*SearchArticleResp, error) {
	out := new(SearchArticleResp)
	err := c.cc.Invoke(ctx, Article_SearchArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServer is the server API for Article service.
// All implementations must embed UnimplementedArticleServer
// for forward compatibility
type ArticleServer interface {
	// -----------------------article-----------------------
	AddArticle(context.Context, *AddArticleReq) (*AddArticleResp, error)
	UpdateArticle(context.Context, *UpdateArticleReq) (*UpdateArticleResp, error)
	DelArticle(context.Context, *DelArticleReq) (*DelArticleResp, error)
	GetArticleById(context.Context, *GetArticleByIdReq) (*GetArticleByIdResp, error)
	SearchArticle(context.Context, *SearchArticleReq) (*SearchArticleResp, error)
	mustEmbedUnimplementedArticleServer()
}

// UnimplementedArticleServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServer struct {
}

func (UnimplementedArticleServer) AddArticle(context.Context, *AddArticleReq) (*AddArticleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddArticle not implemented")
}
func (UnimplementedArticleServer) UpdateArticle(context.Context, *UpdateArticleReq) (*UpdateArticleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticleServer) DelArticle(context.Context, *DelArticleReq) (*DelArticleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelArticle not implemented")
}
func (UnimplementedArticleServer) GetArticleById(context.Context, *GetArticleByIdReq) (*GetArticleByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleById not implemented")
}
func (UnimplementedArticleServer) SearchArticle(context.Context, *SearchArticleReq) (*SearchArticleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArticle not implemented")
}
func (UnimplementedArticleServer) mustEmbedUnimplementedArticleServer() {}

// UnsafeArticleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServer will
// result in compilation errors.
type UnsafeArticleServer interface {
	mustEmbedUnimplementedArticleServer()
}

func RegisterArticleServer(s grpc.ServiceRegistrar, srv ArticleServer) {
	s.RegisterService(&Article_ServiceDesc, srv)
}

func _Article_AddArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).AddArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_AddArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).AddArticle(ctx, req.(*AddArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_UpdateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).UpdateArticle(ctx, req.(*UpdateArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_DelArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).DelArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_DelArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).DelArticle(ctx, req.(*DelArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArticleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArticleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_GetArticleById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArticleById(ctx, req.(*GetArticleByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_SearchArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).SearchArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_SearchArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).SearchArticle(ctx, req.(*SearchArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Article_ServiceDesc is the grpc.ServiceDesc for Article service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Article_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddArticle",
			Handler:    _Article_AddArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _Article_UpdateArticle_Handler,
		},
		{
			MethodName: "DelArticle",
			Handler:    _Article_DelArticle_Handler,
		},
		{
			MethodName: "GetArticleById",
			Handler:    _Article_GetArticleById_Handler,
		},
		{
			MethodName: "SearchArticle",
			Handler:    _Article_SearchArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}
