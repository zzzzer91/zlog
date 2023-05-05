package trace

import (
	"github.com/sirupsen/logrus"
	"github.com/zzzzer91/gopkg/stackx"
	"github.com/zzzzer91/gopkg/typex"
	"github.com/zzzzer91/zlog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	exceptionErrorEventKey = "exception.error"
)

var _ logrus.Hook = (*TraceHook)(nil)

type TraceHook struct {
	cfg *Config
}

func NewTraceHook(opts ...Option) *TraceHook {
	cfg := &Config{
		EnableLevels:       logrus.AllLevels,
		ErrorSpanLevel:     logrus.ErrorLevel,
		IsRecordErrorStack: true,
	}
	for _, o := range opts {
		o(cfg)
	}
	return &TraceHook{cfg: cfg}
}

func (h *TraceHook) Levels() []logrus.Level {
	return h.cfg.EnableLevels
}

func (h *TraceHook) Fire(entry *logrus.Entry) error {
	if entry.Context == nil {
		return nil
	}

	span := trace.SpanFromContext(entry.Context)
	if !span.IsRecording() {
		return nil
	}

	// set span status
	if entry.Level <= h.cfg.ErrorSpanLevel {
		if err, ok := entry.Data[zlog.EntityFieldNameError.String()].(error); ok {
			span.SetStatus(codes.Error, err.Error())
			opts := []trace.EventOption{trace.WithAttributes(
				semconv.ExceptionTypeKey.String(typex.TypeStr(err)),
				semconv.ExceptionMessageKey.String(entry.Message),
				attribute.Key(exceptionErrorEventKey).String(err.Error()),
			)}
			if h.cfg.IsRecordErrorStack {
				if v, ok := entry.Data[zlog.EntityFieldNameErrorStack.String()].(string); ok {
					opts = append(opts, trace.WithAttributes(
						semconv.ExceptionStacktraceKey.String(v),
					))
				} else {
					opts = append(opts, trace.WithAttributes(
						semconv.ExceptionStacktraceKey.String(stackx.RecordStack(7)),
					))
				}
			}
			span.AddEvent(semconv.ExceptionEventName, opts...)
		} else {
			span.SetStatus(codes.Error, entry.Message)
			opts := []trace.EventOption{trace.WithAttributes(
				semconv.ExceptionMessageKey.String(entry.Message),
			)}
			if h.cfg.IsRecordErrorStack {
				opts = append(opts, trace.WithAttributes(
					semconv.ExceptionStacktraceKey.String(stackx.RecordStack(7)),
				))
			}
			span.AddEvent(semconv.ExceptionEventName, opts...)
		}
	}

	return nil
}
