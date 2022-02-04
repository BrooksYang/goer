package logger

import (
	"fmt"
	"os"
	"time"

	"goer/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Channel struct {
	Path  string
	Level string
	Days  int
}

// NewChannel Log channel
func NewChannel(channel Channel) *zap.Logger {
	// Get log writer
	writeSyncer := getLogWriter(channel)

	// Get encoder
	encoder := getEncoder()

	// Log level
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(channel.Level)); err != nil {
		fmt.Println("init log level error")
	}

	// New core
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	// New logger
	logger := zap.New(core, zap.AddCaller())

	return logger
}

// Get encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// log format: NewJSONEncoder or NewConsoleEncoder
	if global.Config.App.IsLocal() {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

// Custom time encoder
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// Get log writer
func getLogWriter(channel Channel) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   channel.Path,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     channel.Days,
		Compress:   false,
	}

	// Additionally print to terminal for local environment.
	if global.Config.App.IsLocal() {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}

	return zapcore.AddSync(lumberJackLogger)
}
