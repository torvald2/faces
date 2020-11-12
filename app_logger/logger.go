package app_logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var logLevel zap.AtomicLevel

func init() {
	switch configLevel := os.Getenv("LOG_LEVEL"); configLevel {
	case "DEBUG":
		logLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "ERROR":
		logLevel = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "WARN":
		logLevel = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	default:
		logLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	conf := zap.Config{
		Encoding:    "json",
		Level:       logLevel,
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}
	Logger, _ = conf.Build()
	zap.ReplaceGlobals(Logger)

}
