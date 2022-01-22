package global

import (
	"goer/config"

	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
)
