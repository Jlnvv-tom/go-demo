// package svc

// import (
// 	"api/internal/config"
// )

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
	"article/cmd/api/internal/config"
	"article/cmd/rpc/article"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ArticleRpc article.ArticleZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ArticleRpc: article.NewArticleZrpcClient(zrpc.MustNewClient(c.ArticleRpcConf)),
	}
}
