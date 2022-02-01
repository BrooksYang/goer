package bootstrap

import (
	"context"
	"fmt"
	"log"

	"goer/global"

	"github.com/go-redis/redis/v8"
)

func Redis() *redis.Client {
	redisCfg := global.Config.Database.Redis
	addr := fmt.Sprintf("%v:%v", redisCfg.Host, redisCfg.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisCfg.Password,
		DB:       redisCfg.Database,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("redis connect ping failed, err: %s", err)
		return nil
	}

	return client
}
