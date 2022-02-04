package config

import (
	"os"

	"go.uber.org/zap"
)

type Logger struct {
	Default *zap.Logger
}

type Logging struct {
	Default Channel
}

type Channel struct {
	Path  string
	Level string
	Days  int
}

func NewLogging() *Logging {
	return &Logging{
		Default: Channel{
			Path:  logPath("serve.log"),
			Level: "debug",
			Days:  14,
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
