package article

import (
	"context"

	"article/cmd/rpc/article"

	"article/cmd/api/internal/svc"
	"article/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改文章
func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.UpdateArticleReq) (resp *types.UpdateArticleResp, err error) {
	// 在这里调用 rpc 下的 AddArticle 方法
	_, err = l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, &article.UpdateArticleReq{
		Title:   req.Title,
		Content: req.Content,
		Id:      req.Id,
	})

	if err != nil {
		return nil, err
	}

	return
}
