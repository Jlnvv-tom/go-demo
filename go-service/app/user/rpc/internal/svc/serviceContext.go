package svc

import (
	"user/model"
	"user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		// ArticleModel: model.NewArticleModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
