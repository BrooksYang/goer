package config

import (
	"os"

	"go.uber.org/zap"
)

type Logger struct {
	Default *zap.Logger `mapstructure:"default" json:"default"`
}

type Logging struct {
	Default Channel `mapstructure:"default" json:"default"`
}

type Channel struct {
	Path  string `mapstructure:"path" json:"path"`
	Level string `mapstructure:"level" json:"level"`
	Days  int    `mapstructure:"days" json:"days"`
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
