//go:generate generate-interfaces.sh

package logger

import (
	"context"
	"os"
	"sync"

	"go.elastic.co/apm/module/apmzap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger abstracts the zap logger to simplify logging.
type Logger struct {
	logger *zap.Logger
}

var (
	//nolint:gochecknoglobals // this global requires for singleton pattern
	loggerOnce sync.Once
	//nolint:gochecknoglobals // this global requires for singleton pattern
	logger InterfaceLogger
)

const (
	// LogLevelEnvKey is a finding key to getting log level from environment
	LogLevelEnvKey = "LOG_LEVEL"
	// DefaultLogLevel default log level, if log level not defined
	DefaultLogLevel = zapcore.ErrorLevel
)

// GetZapLogger returns the zap implementation used by the logger.
func (logger *Logger) GetZapLogger() *zap.Logger {
	return logger.logger
}

// Log wraps the zap logger with the needed context variables to match it to the apm trace.
func (logger *Logger) Log(ctx context.Context) InterfaceZapLogger {
	return logger.logger.With(apmzap.TraceContext(ctx)...)
}

// GetLogger returns logger instance
func GetLogger() InterfaceLogger {
	loggerOnce.Do(func() {
		logger = NewLogger()
	})

	return logger
}

// NewLogger initializes an apm wrapped zap logger.
func NewLogger() InterfaceLogger {
	return NewLoggerInstance()
}

// NewLoggerInstance initializes an apm wrapped zap logger.
// Deprecated: Use NewLogger instead
func NewLoggerInstance() *Logger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), os.Stdout, getLevel(os.Getenv(LogLevelEnvKey)))

	return &Logger{
		logger: zap.New(core),
	}
}

// getLevel returns the level based on the environment variable.
func getLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zapcore.DebugLevel

	case "info":
		return zapcore.InfoLevel

	case "warn":
		return zapcore.WarnLevel

	default:
		return DefaultLogLevel
	}
}
