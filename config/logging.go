package config

import (
	"os"

	"github.com/goer-project/goer-utils/core/logger"
	"go.uber.org/zap"
)

type Logger struct {
	Default *zap.Logger
	Request *zap.Logger
	Mail    *zap.Logger
	Open    *zap.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Default: logger.NewChannel(&logger.Channel{
			Path:    logPath("serve.log"),
			Level:   "debug",
			Days:    14,
			Console: true,
			Format:  "json",
		}),
		Request: logger.NewChannel(&logger.Channel{
			Path:    logPath("request.log"),
			Level:   "debug",
			Days:    14,
			Console: false,
			Format:  "json",
		}),
		Mail: logger.NewChannel(&logger.Channel{
			Path:    logPath("mail.log"),
			Level:   "debug",
			Days:    14,
			Console: true,
			Format:  "json",
		}),
		Open: logger.NewChannel(&logger.Channel{
			Path:    logPath("open.log"),
			Level:   "debug",
			Days:    14,
			Console: true,
			Format:  "json",
		}),
	}
}

func logPath(path string) string {
	if path != "" {
		path = "/" + path
	}
	wd, _ := os.Getwd()

	return wd + "/storage/logs" + path
}
