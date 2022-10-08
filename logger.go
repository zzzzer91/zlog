package zlog

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var logger = NewLogger(
	&LoggerConfig{
		Level:       int(logrus.InfoLevel),
		TerminalOut: true,
	},
	NewTraceHook(&TraceHookConfig{
		RecordStackTraceInSpan: true,
		EnableLevels:           logrus.AllLevels,
		ErrorSpanLevel:         logrus.ErrorLevel,
	}),
)

func NewLogger(config *LoggerConfig, hooks ...logrus.Hook) *logrus.Logger {
	var writers []io.Writer
	if config.TerminalOut {
		writers = append(writers, os.Stderr)
	}
	if config.FileOut {
		writers = append(writers, &lumberjack.Logger{
			Filename:   config.FileConfig.Path,
			MaxSize:    config.FileConfig.MaxSize,
			MaxBackups: config.FileConfig.MaxBackups,
			MaxAge:     config.FileConfig.MaxAge,
		})
	}
	l := &logrus.Logger{
		Out: io.MultiWriter(writers...),
		Formatter: &selfFormatter{
			timeFormat: "2006-01-02 15:04:05.000",
		},
		Level:        logrus.Level(config.Level),
		ReportCaller: true,
	}
	for _, hook := range hooks {
		l.AddHook(hook)
	}
	l.SetNoLock()
	return l
}

// DefaultLogger return the default logger.
func DefaultLogger() *logrus.Logger {
	return logger
}

// SetLogger sets the default logger.
// Note that this method is not concurrent-safe and must not be called
// after the use of DefaultLogger and global functions in this package.
func SetLogger(l *logrus.Logger) {
	logger = l
}

// SetLoggerLevel sets log level.
func SetLoggerLevel(level int) {
	logger.SetLevel(logrus.Level(level))
}
