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
		ctx = context.WithValue(ctx, zlog.EntityFieldNameTraceId, span.SpanContext().TraceID())
	}
	return ctx, span
}
