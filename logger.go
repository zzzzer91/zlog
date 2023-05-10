package zlog

import (
	"context"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(config *LoggerConfig, hooks ...logrus.Hook) *Logger {
	var writers []io.Writer
	if config.TerminalOut == nil || *config.TerminalOut {
		writers = append(writers, os.Stderr)
	}
	if config.FileOut != nil && *config.FileOut {
		writers = append(writers, &lumberjack.Logger{
			Filename:   config.FileConfig.Path,
			MaxSize:    config.FileConfig.MaxSize,
			MaxBackups: config.FileConfig.MaxBackups,
			MaxAge:     config.FileConfig.MaxAge,
		})
	}
	l := logrus.New()
	l.SetNoLock()
	l.SetLevel(logrus.Level(config.Level))
	timeFormat := defaultTimeFormat
	if config.TimeFormat != "" {
		timeFormat = config.TimeFormat
	}
	l.SetFormatter(&selfFormatter{
		timeFormat: timeFormat,
	})
	l.SetReportCaller(true)
	l.SetOutput(io.MultiWriter(writers...))
	for _, hook := range hooks {
		l.AddHook(hook)
	}
	return &Logger{l}
}

func (l *Logger) Ctx(ctx context.Context) *logrus.Entry {
	entry := l.Logger.WithContext(ctx)
	fields := make(logrus.Fields)
	if v := ctx.Value(EntityFieldNameTraceId); v != nil {
		fields[EntityFieldNameTraceId.String()] = v
	}
	if v := ctx.Value(EntityFieldNameRequestId); v != nil {
		fields[EntityFieldNameRequestId.String()] = v
	}
	if v := ctx.Value(EntityFieldNameLogId); v != nil {
		fields[EntityFieldNameLogId.String()] = v
	}
	return entry.WithFields(fields)
}
