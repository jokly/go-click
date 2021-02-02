package util

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitzap "github.com/go-kit/kit/log/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(minLevel int8) (*zap.Logger, log.Logger) {
	minLoggerLevel := zapcore.Level(minLevel)

	zLogger, err := initZapLogger(minLoggerLevel)
	if err != nil {
		panic(err)
	}

	logger := kitzap.NewZapSugarLogger(zLogger, minLoggerLevel)
	logger = level.NewFilter(logger, levelZapToGoKit(minLoggerLevel))
	logger = log.With(logger, "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)

	return zLogger, logger
}

func initZapLogger(minLevel zapcore.Level) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	config.Level = zap.NewAtomicLevelAt(minLevel)

	// Disable zap additional information
	config.EncoderConfig.TimeKey = ""
	config.EncoderConfig.MessageKey = ""
	config.EncoderConfig.CallerKey = ""
	config.EncoderConfig.LevelKey = ""

	return config.Build()
}

func levelZapToGoKit(lvl zapcore.Level) level.Option {
	switch lvl {
	case zapcore.DebugLevel:
		return level.AllowDebug()
	case zapcore.InfoLevel:
		return level.AllowInfo()
	case zapcore.WarnLevel:
		return level.AllowWarn()
	case zapcore.ErrorLevel:
		return level.AllowError()
	default:
		return level.AllowAll()
	}
}
