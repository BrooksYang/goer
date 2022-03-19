package bootstrap

import (
	"fmt"

	"goer/global"

	"github.com/goer-project/goer-core/cache"
	"github.com/goer-project/goer-core/redis"
)

func Cache() {
	// Redis config
	redisCfg := global.Config.Database.Redis
	addr := fmt.Sprintf("%v:%v", redisCfg.Host, redisCfg.Port)

	// Redis store
	rs := &cache.RedisStore{}
	rs.RedisClient = redis.NewClient(addr, redisCfg.Password, redisCfg.DatabaseCache)
	rs.KeyPrefix = global.Config.App.Name + ":cache:"

	// New cache service
	global.Cache = cache.NewService(rs)
}
