package userlogic

import (
	"context"

	proto "proto/user"
	"user/model"
	"user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------user-----------------------
func (l *CreateUserLogic) CreateUser(in *proto.CreateUserReq) (*proto.CreateUserResp, error) {
	// todo: add your logic here and delete this line
	user := model.User{}
	_ = user.Marshal(in)

	return &proto.CreateUserResp{}, nil
}
