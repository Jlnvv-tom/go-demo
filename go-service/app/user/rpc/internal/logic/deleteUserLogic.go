package logic

import (
	"context"

	"user/D:/code/go-workspace/go-demo/go-service/proto/user"
	"user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.DeleteUserReq) (*user.DeleteUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteUserResp{}, nil
}
