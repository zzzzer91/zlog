package zlog

import (
	"github.com/pkg/errors"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

type errStackTracer interface {
	StackTrace() errors.StackTrace
}

type errCauser interface {
	Cause() error
}

// getCallerFramesFromError try to get error's stack.
func getCallerFramesFromError(err error) *runtime.Frames {
	stackTracer := tryFindErrStackTacker(err)
	if stackTracer == nil {
		return nil
	}
	st := stackTracer.StackTrace()
	return runtime.CallersFrames(*(*[]uintptr)(unsafe.Pointer(&st)))
}

// tryFindErrStackTacker try to find last err that implements errStackTracer.
func tryFindErrStackTacker(err error) errStackTracer {
	var st errStackTracer
	for err != nil {
		v, ok := err.(errStackTracer)
		if ok {
			st = v
		}
		cause, ok := err.(errCauser)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return st
}

// RecordStackTrace return current error stack string.
func RecordStackTrace(skip int) string {
	return CallersFrames2Str(GetCallersFrames(skip + 2))
}

// GetCallersFrames returns current error stack.
func GetCallersFrames(skip int) *runtime.Frames {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip+1, pcs[:])
	callersFrames := runtime.CallersFrames(pcs[:n])
	return callersFrames
}

func CallersFrames2Str(callersFrames *runtime.Frames) string {
	if callersFrames == nil {
		return ""
	}
	var sb strings.Builder
	for f, again := callersFrames.Next(); again; f, again = callersFrames.Next() {
		sb.WriteString(f.File)
		sb.WriteByte(':')
		sb.WriteString(strconv.Itoa(f.Line))
		sb.WriteByte('\n')
		sb.WriteByte('\t')
		sb.WriteString(function(f.Function))
		sb.WriteByte('\n')
	}
	s := sb.String()
	return s[:len(s)-1]
}

// function returns, if possible, the name of the function containing the PC.
func function(name string) string {
	const (
		centerDot = "·"
		dot       = "."
		slash     = "/"
	)
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contain dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastSlash := strings.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := strings.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	return strings.Replace(name, centerDot, dot, -1)
}
