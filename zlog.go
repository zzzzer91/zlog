package zlog

import (
	"context"

	"github.com/sirupsen/logrus"
)

var logger = NewLogger(
	&LoggerConfig{
		Level:       int(logrus.InfoLevel),
		TerminalOut: true,
	},
)

// DefaultLogger return the default logger.
func DefaultLogger() *Logger {
	return logger
}

// SetLogger sets the default logger.
// Note that this method is not concurrent-safe and must not be called
// after the use of DefaultLogger and global functions in this package.
func SetLogger(l *logrus.Logger) {
	logger = &Logger{l}
}

// SetLoggerLevel sets log level.
func SetLoggerLevel(level int) {
	logger.SetLevel(logrus.Level(level))
}

func Ctx(ctx context.Context) *logrus.Entry {
	return logger.Ctx(ctx)
}
