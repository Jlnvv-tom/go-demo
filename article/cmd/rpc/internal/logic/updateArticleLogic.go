package logic

import (
	"article/cmd/rpc/internal/svc"
	"article/cmd/rpc/pb"
	"article/model"
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateArticleLogic) UpdateArticle(in *pb.UpdateArticleReq) (*pb.UpdateArticleResp, error) {

	article := new(model.Article)

	article.Id = in.Id
	article.Title = sql.NullString{String: in.Title, Valid: true}
	article.Content = sql.NullString{String: in.Content, Valid: true}

	err := l.svcCtx.ArticleModel.Update(l.ctx, article)
	l.Infof("=======================> title %v", article.Title)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateArticleResp{}, nil
}
