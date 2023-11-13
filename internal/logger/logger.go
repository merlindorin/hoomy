package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(level zapcore.Level, development bool, opts ...zap.Option) (logger *zap.Logger, err error) {
	config := zap.NewProductionConfig()
	if development {
		config = zap.NewDevelopmentConfig()
	}

	config.Level = zap.NewAtomicLevelAt(level)

	return config.Build(opts...)
}
