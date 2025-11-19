package pkg

import (
	"go.uber.org/zap"
)

type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	WithFields(fields map[string]interface{}) Logger
}
type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() Logger {
	zapLogger, _ := zap.NewProduction()
	return &ZapLogger{logger: zapLogger.Sugar()}
}

func (l *ZapLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Debugw(msg, fields...)
}

func (l *ZapLogger) Info(msg string, fields ...interface{}) {
	l.logger.Infow(msg, fields...)
}

func (l *ZapLogger) Warn(msg string, fields ...interface{}) {
	l.logger.Warnw(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...interface{}) {
	l.logger.Errorw(msg, fields...)
}

func (l *ZapLogger) WithFields(fields map[string]interface{}) Logger {
	var args []interface{}
	for k, v := range fields {
		args = append(args, k, v)
	}

	newLogger := l.logger.With(args...)
	return &ZapLogger{logger: newLogger}
}

// ✅ Добавляем метод Sync для закрытия логгера
func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}
