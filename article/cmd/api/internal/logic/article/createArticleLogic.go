package article

import (
	"article/cmd/rpc/article"
	"context"

	"article/cmd/api/internal/svc"
	"article/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建文章
func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateArticleLogic) CreateArticle(req *types.CreateArticleReq) (resp *types.CreateArticleResp, err error) {
	// 在这里调用 rpc 下的 AddArticle 方法
	_, err = l.svcCtx.ArticleRpc.AddArticle(l.ctx, &article.AddArticleReq{
		Title:   req.Title,
		Content: req.Content,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateArticleResp{Success: "添加成功"}, nil
}
