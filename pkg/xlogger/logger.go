package xlogger

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

func New() (logr.Logger, error) {
	zc := zap.NewProductionConfig()
	zc.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zc.DisableStacktrace = true
	z, err := zc.Build()
	if err != nil {
		return logr.Logger{}, err
	}
	return zapr.NewLogger(z), nil
}
