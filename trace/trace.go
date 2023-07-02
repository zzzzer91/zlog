package trace

import (
	"context"

	"github.com/zzzzer91/zlog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func StartTracing(ctx context.Context, spanName string) (context.Context, trace.Span) {
	var span trace.Span
	ctx, span = otel.Tracer("github.com/zzzzer91/zlog").Start(ctx, spanName)
	if span.IsRecording() {
		ctx = context.WithValue(ctx, zlog.EntityFieldNameTraceId, span.SpanContext().TraceID().String())
	}
	return ctx, span
}

// CopyContext 拷贝基本字段的同时，拷贝 trace 信息
func CopyContext(ctx context.Context) context.Context {
	return trace.ContextWithSpan(zlog.CopyContext(ctx), trace.SpanFromContext(ctx))
}
