package trace

import (
	"context"

	"github.com/zzzzer91/zlog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func StartTracing(ctx context.Context, spanName string) (context.Context, trace.Span) {
	var span trace.Span
	ctx, span = otel.Tracer("github.com/zzzzer91/zlog").
		Start(ctx, spanName, trace.WithSpanKind(trace.SpanKindClient))
	if span.IsRecording() {
		ctx = context.WithValue(ctx,
			zlog.EntityFieldNameTraceID, span.SpanContext().TraceID().String())
	}
	return ctx, span
}
