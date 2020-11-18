package app_logger

import (
	"atbmarket.comfaceapp/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var logLevel zap.AtomicLevel

func init() {
	app_conf := config.GetConfig()
	switch app_conf.LogLevel {
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
