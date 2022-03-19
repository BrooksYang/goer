package global

import (
	"goer/config"

	"github.com/goer-project/goer-core/cache"
	"github.com/goer-project/goer-core/redis"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.RedisClient
	Cache  *cache.CacheService
	Logger *config.Logger
)
