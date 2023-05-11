package zlog

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

func AddLogIdToCtx(ctx context.Context) context.Context {
	if v := ctx.Value(EntityFieldNameLogId); v == nil {
		ctx = context.WithValue(ctx, EntityFieldNameLogId, uuid.New().String())
	}
	return ctx
}

// CopyContext 只保留原来 context 中指定字段
func CopyContext(ctx context.Context) context.Context {
	newCtx := context.Background()
	if v := ctx.Value(EntityFieldNameTraceId); v != nil {
		newCtx = context.WithValue(newCtx, EntityFieldNameTraceId, v)
	}
	if v := ctx.Value(EntityFieldNameRequestId); v != nil {
		newCtx = context.WithValue(newCtx, EntityFieldNameRequestId, v)
	}
	if v := ctx.Value(EntityFieldNameLogId); v != nil {
		newCtx = context.WithValue(newCtx, EntityFieldNameLogId, v)
	}
	return newCtx
}
