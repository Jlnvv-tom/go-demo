// package config

// import "github.com/zeromicro/go-zero/rest"

// type Config struct {
// 	rest.RestConf
// }

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	Cache          cache.CacheConf
	ArticleRpcConf zrpc.RpcClientConf
}
