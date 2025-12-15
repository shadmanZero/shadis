package logger

import (
	"github.com/shadman/shadis/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Init(cfg *config.Config) error {
	level := getZapLevel(cfg.LoggingConfig.LogLevel)
	isDev := cfg.LoggingConfig.LogLevel == string(config.LogLevelDebug)

	zapCfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: isDev,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeLevel:   zapcore.CapitalColorLevelEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := zapCfg.Build()
	if err != nil {
		return err
	}

	log = logger
	return nil
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}

func With(fields ...zap.Field) *zap.Logger {
	return log.With(fields...)
}

func Sync() error {
	if log != nil {
		return log.Sync()
	}
	return nil
}

func GetLogger() *zap.Logger {
	return log
}
