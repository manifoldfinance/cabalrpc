package log

import (
	"github.com/manifoldfinance/cabalrpc/internal/config"
	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func GetLogger(cfg config.Config) *zap.Logger {
	if cfg.ApmEnabled {
		return getApmLogger(cfg.LogLevel.ZapLevel)
	} else {
		return getLogger(cfg.LogLevel.ZapLevel)
	}
}

func getLogger(level zapcore.Level) *zap.Logger {
	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	atom.SetLevel(level)
	return logger
}

func getApmLogger(level zapcore.Level) *zap.Logger {
	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	logger := zap.New(zapcore.NewTee(&apmzap.Core{}, zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	)))
	atom.SetLevel(level)
	return logger
}
