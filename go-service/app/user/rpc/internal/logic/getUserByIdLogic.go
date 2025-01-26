package logic

import (
	"context"

	"user/D:/code/go-workspace/go-demo/go-service/proto/user"
	"user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.GetUserByIdReq) (*user.GetUserByIdResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserByIdResp{}, nil
}
