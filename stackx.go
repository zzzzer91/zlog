package zlog

import (
	"bytes"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

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
	sb := new(strings.Builder)
	var lines [][]byte
	var lastFile string
	for f, again := callersFrames.Next(); again; f, again = callersFrames.Next() {
		sb.WriteString(f.File)
		sb.WriteByte(':')
		sb.WriteString(strconv.Itoa(f.Line))
		sb.WriteByte('\n')
		sb.WriteByte('\t')
		sb.Write(function(f.PC))
		if f.File != lastFile {
			data, err := os.ReadFile(f.File)
			if err == nil {
				lines = bytes.Split(data, []byte{'\n'})
				lastFile = f.File
				sb.WriteString(": ")
				sb.Write(source(lines, f.Line))
			}
		} else {
			if len(lines) > 0 {
				sb.WriteString(": ")
				sb.Write(source(lines, f.Line))
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contain dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}
