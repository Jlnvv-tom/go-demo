package article

import (
	"context"

	"article/cmd/api/internal/svc"
	"article/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获得文章列表
func NewGetArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesLogic {
	return &GetArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesLogic) GetArticles(req *types.GetArticleListReq) (resp *types.GetArticleListResp, err error) {
	// 调用 rpc 的GetArticleById 方法

	return &types.GetArticleListResp{Articles: []types.Article{{Id: 123, Title: "你好", Content: "内容"}}}, nil
}
