package zlog

import "context"

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
