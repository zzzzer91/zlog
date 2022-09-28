package zlog

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func StartTracing(ctx context.Context, spanName string) (context.Context, trace.Span) {
	var span trace.Span
	ctx, span = otel.Tracer("github.com/zzzzer91/zlog").Start(ctx, spanName)
	ctx = context.WithValue(ctx, EntityFieldNameTraceID, span.SpanContext().TraceID())
	if v, ok := ctx.Value(EntityFieldNameRequestID).(string); ok {
		span.SetAttributes(attribute.String("requestID", v))
	}
	return ctx, span
}
