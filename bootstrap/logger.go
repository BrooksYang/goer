package bootstrap

import (
	"goapp/config"
	"goapp/global"
)

func Logger() {
	global.Logger = config.NewLogger()
}
