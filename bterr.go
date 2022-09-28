package zlog

import (
	"runtime"
	"unsafe"

	"github.com/pkg/errors"
)

type errStackTracer interface {
	StackTrace() errors.StackTrace
}

type errCauser interface {
	Cause() error
}

// getCallersFrames 尝试获取 err 的错误栈
func getCallersFrames(err error) *runtime.Frames {
	stackTracer := tryFindErrStackTacker(err)
	if stackTracer == nil {
		return nil
	}
	st := stackTracer.StackTrace()
	return runtime.CallersFrames(*(*[]uintptr)(unsafe.Pointer(&st)))
}

// tryFindErrStackTacker 递归寻找最后一个实现了 errStackTracer 接口的 err
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
