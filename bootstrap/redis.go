package bootstrap

import (
	"fmt"

	"goer/global"

	"github.com/goer-project/goer-core/redis"
)

func Redis() {
	// Redis config
	redisCfg := global.Config.Database.Redis
	addr := fmt.Sprintf("%v:%v", redisCfg.Host, redisCfg.Port)

	// New redis client
	global.Redis = redis.NewClient(addr, redisCfg.Password, redisCfg.Database)
}
