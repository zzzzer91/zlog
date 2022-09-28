package zlog

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	exceptionErrorEventKey = "exception.error"
	logEventKey            = "log"
)

var (
	logSeverityTextKey = attribute.Key("otel.log.severity.text")
	logMessageKey      = attribute.Key("otel.log.message")
)

var _ logrus.Hook = (*TraceHook)(nil)

type TraceHookConfig struct {
	RecordStackTraceInSpan bool
	EnableLevels           []logrus.Level
	ErrorSpanLevel         logrus.Level
}

type TraceHook struct {
	cfg *TraceHookConfig
}

func NewTraceHook(cfg *TraceHookConfig) *TraceHook {
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

	// attach span context to log entry data fields
	entry.Data[EntityFieldNameTraceID.String()] = span.SpanContext().TraceID()

	// set span status
	if entry.Level <= h.cfg.ErrorSpanLevel {
		if err, ok := entry.Data[EntityFieldNameError.String()].(error); ok {
			span.SetStatus(codes.Error, err.Error())
			opts := []trace.EventOption{trace.WithAttributes(
				semconv.ExceptionTypeKey.String(TypeStr(err)),
				semconv.ExceptionMessageKey.String(entry.Message),
				attribute.Key(exceptionErrorEventKey).String(err.Error()),
			)}
			if v, ok := entry.Data[EntityFieldNameErrorStack.String()].(string); ok {
				opts = append(opts, trace.WithAttributes(
					semconv.ExceptionStacktraceKey.String(v),
				))
			} else {
				opts = append(opts, trace.WithAttributes(
					semconv.ExceptionStacktraceKey.String(RecordStackTrace(7)),
				))
			}
			span.AddEvent(semconv.ExceptionEventName, opts...)
		} else {
			span.SetStatus(codes.Error, entry.Message)
			span.RecordError(errors.New(entry.Message), trace.WithStackTrace(h.cfg.RecordStackTraceInSpan))
		}
	} else {
		// attach log to span event attributes
		attrs := []attribute.KeyValue{
			logMessageKey.String(entry.Message),
			logSeverityTextKey.String(otelSeverityText(entry.Level)),
		}
		span.AddEvent(logEventKey, trace.WithAttributes(attrs...))
	}

	return nil
}

// otelSeverityText convert logrus level to otel severityText
// ref to https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model.md#severity-fields
func otelSeverityText(lv logrus.Level) string {
	s := lv.String()
	if s == "warning" {
		s = "warn"
	}
	return strings.ToUpper(s)
}
