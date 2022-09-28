package zlog

import (
	"runtime"
	"strconv"
	"strings"
)

// RecordStackTrace 获取当前调用栈，返回字符串
func RecordStackTrace(skip int) string {
	return CallersFrames2Str(GetCallersFrames(skip + 2))
}

// GetCallersFrames 获取当前调用栈
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
		sb.WriteString(f.Func.Name())
		sb.WriteByte('\n')
		sb.WriteByte('\t')
		sb.WriteString(f.File)
		sb.WriteByte(':')
		sb.WriteString(strconv.Itoa(f.Line))
		if again {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}
