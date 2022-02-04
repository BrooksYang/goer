package bootstrap

import (
	"goer/config"
	"goer/global"
	"goer/pkg/logger"
)

func Logger() {
	logging := config.NewLogging()

	global.Logger = &config.Logger{
		Default: logger.NewChannel(logger.Channel(logging.Default)),
	}
}
