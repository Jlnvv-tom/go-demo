package logic

import (
	"context"
	"database/sql"

	"article/cmd/rpc/internal/svc"
	"article/cmd/rpc/pb"

	"article/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------article-----------------------
func (l *AddArticleLogic) AddArticle(in *pb.AddArticleReq) (*pb.AddArticleResp, error) {
	article := new(model.Article)
	article.Title = sql.NullString{String: in.Title, Valid: true}
	article.Content = sql.NullString{String: in.Content, Valid: true}
	// 调用model层的方法
	_, err := l.svcCtx.ArticleModel.Insert(l.ctx, article)

	if err != nil {
		return nil, err
	}

	return &pb.AddArticleResp{}, nil
}
