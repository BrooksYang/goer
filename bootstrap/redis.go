package bootstrap

import (
	"fmt"

	"goer/global"
	"goer/pkg/redis"
)

func Redis() {
	// Redis config
	redisCfg := global.Config.Database.Redis
	addr := fmt.Sprintf("%v:%v", redisCfg.Host, redisCfg.Port)

	// New redis client
	global.Redis = redis.NewClient(addr, redisCfg.Password, redisCfg.Database)
}
