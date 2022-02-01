package global

import (
	"goer/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
