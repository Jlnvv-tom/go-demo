package logic

import (
	"article/cmd/rpc/internal/svc"
	"article/cmd/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdLogic {
	return &GetArticleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleByIdLogic) GetArticleById(in *pb.GetArticleByIdReq) (*pb.GetArticleByIdResp, error) {

	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.Id)

	if err != nil {
		l.Errorf("query article by id error: %v", err)
		return nil, err
	}

	resp := &pb.GetArticleByIdResp{
		Article: &pb.Article{
			Id:      article.Id,
			Title:   "Title of the article",
			Content: "Content of the article",
		},
	}
	return resp, nil
}
