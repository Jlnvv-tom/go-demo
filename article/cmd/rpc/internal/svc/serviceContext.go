// package svc

// import "rpc/internal/config"

// type ServiceContext struct {
// 	Config config.Config
// }

// func NewServiceContext(c config.Config) *ServiceContext {
// 	return &ServiceContext{
// 		Config: c,
// 	}
// }

package svc

import (
	"article/cmd/rpc/internal/config"
	"article/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Redis
	ArticleModel model.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
