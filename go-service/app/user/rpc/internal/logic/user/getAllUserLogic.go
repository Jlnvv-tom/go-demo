package userlogic

import (
	"context"

	"proto/user"
	"user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUserLogic {
	return &GetAllUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllUserLogic) GetAllUser(in *user.GetAllUserReq) (*user.GetAllUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetAllUserResp{}, nil
}
