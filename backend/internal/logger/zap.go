package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	*zap.SugaredLogger
}

func NewLogger() *ZapLogger {
	logger, _ := zap.NewDevelopment()
	return &ZapLogger{logger.Sugar()}
}
