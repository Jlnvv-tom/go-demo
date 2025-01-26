package user

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据所有用户
func NewGetAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUserLogic {
	return &GetAllUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUserLogic) GetAllUser(req *types.GetListUserReq) (resp *types.GetListUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
