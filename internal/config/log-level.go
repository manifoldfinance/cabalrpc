package config

import (
	"errors"
	"fmt"
	"go.uber.org/zap/zapcore"
	"strings"
)

const (
	LogLevelDebug = "DEBUG"
	LogLevelInfo  = "INFO"
	LogLevelWarn  = "WARN"
	LogLevelError = "ERROR"
)

func NewLogLevelFlag() LogLevelFlag {
	return LogLevelFlag{
		LevelString: "DEBUG",
		ZapLevel:    zapcore.DebugLevel,
	}
}

type LogLevelFlag struct {
	LevelString string
	ZapLevel    zapcore.Level
}

func (l LogLevelFlag) String() string {
	return l.LevelString
}

func (l *LogLevelFlag) Set(str string) error {
	levelString := strings.ToUpper(str)
	switch levelString {
	case LogLevelDebug:
		l.ZapLevel = zapcore.DebugLevel
		return nil
	case LogLevelInfo:
		l.ZapLevel = zapcore.InfoLevel
		return nil
	case LogLevelWarn:
		l.ZapLevel = zapcore.WarnLevel
		return nil
	case LogLevelError:
		l.ZapLevel = zapcore.ErrorLevel
		return nil
	}
	return errors.New(fmt.Sprintf("unknown log level: %s", str))
}

func (l LogLevelFlag) Type() string {
	return "logLevel"
}
