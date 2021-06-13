package xlogger

import (
	"context"

	"go.uber.org/zap"
)

var defaultLogger *Logger

type Logger struct {
	Logger *zap.SugaredLogger
}

func New() *Logger {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	defaultLogger = &Logger{
		Logger: sugar,
	}
	return defaultLogger
}

func (l *Logger) Debugf(ctx context.Context) *Logger {
	return l
}
