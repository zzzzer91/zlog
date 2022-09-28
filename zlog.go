package zlog

import (
	"context"

	"github.com/sirupsen/logrus"
)

func Ctx(ctx context.Context) *logrus.Entry {
	entry := logger.WithContext(ctx)
	fields := make(logrus.Fields)
	if v := ctx.Value(EntityFieldNameTraceID); v != nil {
		fields[EntityFieldNameTraceID.String()] = v
	}
	if v := ctx.Value(EntityFieldNameRequestID); v != nil {
		fields[EntityFieldNameRequestID.String()] = v
	}
	return entry.WithFields(fields)
}
