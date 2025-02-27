// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package user

import (
	"context"

	"proto/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateUserReq   = user.CreateUserReq
	CreateUserResp  = user.CreateUserResp
	DeleteUserReq   = user.DeleteUserReq
	DeleteUserResp  = user.DeleteUserResp
	GetAllUserReq   = user.GetAllUserReq
	GetAllUserResp  = user.GetAllUserResp
	GetUserByIdReq  = user.GetUserByIdReq
	GetUserByIdResp = user.GetUserByIdResp
	UpdateUserReq   = user.UpdateUserReq
	UpdateUserResp  = user.UpdateUserResp
	UserInfo        = user.UserInfo

	User interface {
		// -----------------------user-----------------------
		CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
		DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error)
		GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
		GetAllUser(ctx context.Context, in *GetAllUserReq, opts ...grpc.CallOption) (*GetAllUserResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// -----------------------user-----------------------
func (m *defaultUser) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUser) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

func (m *defaultUser) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultUser) GetAllUser(ctx context.Context, in *GetAllUserReq, opts ...grpc.CallOption) (*GetAllUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetAllUser(ctx, in, opts...)
}
