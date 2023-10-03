package zlog

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/zzzzer91/gopkg/uuidx"
)

var logger = NewLogger(
	&LoggerConfig{
		Level: int(logrus.InfoLevel),
	},
)

// DefaultLogger return the default logger.
func DefaultLogger() *Logger {
	return logger
}

// SetLogger sets the default logger.
// Note that this method is not concurrent-safe and must not be called
// after the use of DefaultLogger and global functions in this package.
func SetLogger(l *Logger) {
	logger = l
}

// SetLoggerLevel sets log level.
func SetLoggerLevel(level int) {
	logger.SetLevel(logrus.Level(level))
}

func Ctx(ctx context.Context) *logrus.Entry {
	return logger.Ctx(ctx)
}

func AddLogIDToCtx(ctx context.Context) context.Context {
	if v := ctx.Value(EntityFieldNameLogID); v == nil {
		ctx = context.WithValue(ctx, EntityFieldNameLogID, uuidx.New())
	}
	return ctx
}

// CopyContext 只保留原来 context 中指定字段
func CopyContext(ctx context.Context) context.Context {
	newCtx := context.Background()
	if v := ctx.Value(EntityFieldNameTraceID); v != nil {
		newCtx = context.WithValue(newCtx, EntityFieldNameTraceID, v)
	}
	if v := ctx.Value(EntityFieldNameRequestID); v != nil {
		newCtx = context.WithValue(newCtx, EntityFieldNameRequestID, v)
	}
	if v := ctx.Value(EntityFieldNameLogID); v != nil {
		newCtx = context.WithValue(newCtx, EntityFieldNameLogID, v)
	}
	return newCtx
}
