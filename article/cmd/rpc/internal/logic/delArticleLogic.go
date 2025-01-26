package logic

import (
	"context"

	"article/cmd/rpc/internal/svc"
	"article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelArticleLogic {
	return &DelArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelArticleLogic) DelArticle(in *pb.DelArticleReq) (*pb.DelArticleResp, error) {
	err := l.svcCtx.ArticleModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DelArticleResp{}, nil
}
