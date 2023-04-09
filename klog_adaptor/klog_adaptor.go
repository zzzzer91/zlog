package klog_adaptor

import (
	"context"
	"io"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/sirupsen/logrus"
	"github.com/zzzzer91/zlog"
)

type klogAdaptor struct {
	l *zlog.Logger
}

var _ klog.FullLogger = (*klogAdaptor)(nil)

func NewKlogAdaptor(l *zlog.Logger) *klogAdaptor {
	return &klogAdaptor{
		l: l,
	}
}

func (l *klogAdaptor) Trace(v ...interface{}) {
	l.l.Trace(v...)
}

func (l *klogAdaptor) Debug(v ...interface{}) {
	l.l.Debug(v...)
}

func (l *klogAdaptor) Info(v ...interface{}) {
	l.l.Info(v...)
}

func (l *klogAdaptor) Notice(v ...interface{}) {
	l.l.Warn(v...)
}

func (l *klogAdaptor) Warn(v ...interface{}) {
	l.l.Warn(v...)
}

func (l *klogAdaptor) Error(v ...interface{}) {
	l.l.Error(v...)
}

func (l *klogAdaptor) Fatal(v ...interface{}) {
	l.l.Fatal(v...)
}

func (l *klogAdaptor) Tracef(format string, v ...interface{}) {
	l.l.Tracef(format, v...)
}

func (l *klogAdaptor) Debugf(format string, v ...interface{}) {
	l.l.Debugf(format, v...)
}

func (l *klogAdaptor) Infof(format string, v ...interface{}) {
	l.l.Infof(format, v...)
}

func (l *klogAdaptor) Noticef(format string, v ...interface{}) {
	l.l.Warnf(format, v...)
}

func (l *klogAdaptor) Warnf(format string, v ...interface{}) {
	l.l.Warnf(format, v...)
}

func (l *klogAdaptor) Errorf(format string, v ...interface{}) {
	l.l.Errorf(format, v...)
}

func (l *klogAdaptor) Fatalf(format string, v ...interface{}) {
	l.l.Fatalf(format, v...)
}

func (l *klogAdaptor) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.l.Ctx(ctx).Tracef(format, v...)
}

func (l *klogAdaptor) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.l.Ctx(ctx).Debugf(format, v...)
}

func (l *klogAdaptor) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.l.Ctx(ctx).Infof(format, v...)
}

func (l *klogAdaptor) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.l.Ctx(ctx).Warnf(format, v...)
}

func (l *klogAdaptor) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.l.Ctx(ctx).Warnf(format, v...)
}

func (l *klogAdaptor) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.l.Ctx(ctx).Errorf(format, v...)
}

func (l *klogAdaptor) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.l.Ctx(ctx).Fatalf(format, v...)
}

func (l *klogAdaptor) SetLevel(level klog.Level) {
	var lv logrus.Level
	switch level {
	case klog.LevelTrace:
		lv = logrus.TraceLevel
	case klog.LevelDebug:
		lv = logrus.DebugLevel
	case klog.LevelInfo:
		lv = logrus.InfoLevel
	case klog.LevelWarn, klog.LevelNotice:
		lv = logrus.WarnLevel
	case klog.LevelError:
		lv = logrus.ErrorLevel
	case klog.LevelFatal:
		lv = logrus.FatalLevel
	default:
		lv = logrus.WarnLevel
	}
	l.l.SetLevel(lv)
}

func (l *klogAdaptor) SetOutput(writer io.Writer) {
	l.l.SetOutput(writer)
}
