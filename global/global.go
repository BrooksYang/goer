package global

import (
	"goer/config"
	"goer/pkg/redis"

	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.RedisClient
)
