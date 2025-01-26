package article

import (
	"article/cmd/api/internal/svc"
	"article/cmd/api/internal/types"
	"article/cmd/rpc/article"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文章
func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleReq) (resp *types.DeleteArticleResp, err error) {
	_, err = l.svcCtx.ArticleRpc.DelArticle(l.ctx, &article.DelArticleReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return
}
