package zlog

import (
	"github.com/sirupsen/logrus"
)

var logger = NewLogger(&LoggerConfig{Level: int(logrus.InfoLevel)}, NewTraceHook(&TraceHookConfig{
	RecordStackTraceInSpan: true,
	EnableLevels:           logrus.AllLevels,
	ErrorSpanLevel:         logrus.ErrorLevel,
}))

func NewLogger(config *LoggerConfig, hooks ...logrus.Hook) *logrus.Logger {
	l := logrus.New()
	l.SetNoLock()
	l.SetFormatter(&selfFormatter{
		timeFormat: "2006-01-02 15:04:05.000",
	})
	l.SetReportCaller(true)
	l.SetLevel(logrus.Level(config.Level))
	for _, hook := range hooks {
		l.AddHook(hook)
	}
	return l
}

// DefaultLogger return the default logger for kitex.
func DefaultLogger() *logrus.Logger {
	return logger
}

// SetLogger sets the default logger.
// Note that this method is not concurrent-safe and must not be called
// after the use of DefaultLogger and global functions in this package.
func SetLogger(l *logrus.Logger) {
	logger = l
}

func SetLoggerLevel(level int) {
	logger.SetLevel(logrus.Level(level))
}
