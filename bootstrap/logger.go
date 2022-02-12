package bootstrap

import (
	"goer/config"
	"goer/global"
)

func Logger() {
	global.Logger = config.NewLogger()
}
