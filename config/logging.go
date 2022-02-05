package config

import (
	"os"

	"go.uber.org/zap"
)

type Logger struct {
	Default *zap.Logger
	Request *zap.Logger
}

type Logging struct {
	Default Channel
	Request Channel
}

type Channel struct {
	Path    string
	Level   string
	Days    int
	Console bool
}

func NewLogging() *Logging {
	return &Logging{
		Default: Channel{
			Path:    logPath("serve.log"),
			Level:   "debug",
			Days:    14,
			Console: true,
		},
		Request: Channel{
			Path:    logPath("request.log"),
			Level:   "debug",
			Days:    14,
			Console: false,
		},
	}
}

func logPath(path string) string {
	if path != "" {
		path = "/" + path
	}
	wd, _ := os.Getwd()

	return wd + "/storage/logs" + path
}
