package zlog

import (
	"fmt"
	"github.com/bytedance/sonic"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)

// selfFormatter 自定义日志格式
type selfFormatter struct {
	timeFormat string
}

type logStruct struct {
	Time        string            `json:"time"`
	Level       string            `json:"level"`
	File        string            `json:"file,omitempty"`
	Msg         string            `json:"msg,omitempty"`
	Error       string            `json:"error,omitempty"`
	ExtraFields map[string]string `json:"extraFields,omitempty"`
}

var (
	logStructPool = sync.Pool{
		New: func() any {
			return &logStruct{}
		},
	}
)

func (ls *logStruct) reset() {
	ls.Time = ""
	ls.Level = ""
	ls.File = ""
	ls.Msg = ""
	ls.Error = ""
	ls.ExtraFields = nil
}

// Format 实现 formatter 定义的接口，自定义日志格式
func (f *selfFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	ls := logStructPool.Get().(*logStruct)
	ls.Time = entry.Time.Format(f.timeFormat)
	ls.Level = entry.Level.String()
	ls.File = entry.Caller.Function + ":" + strconv.Itoa(entry.Caller.Line)
	ls.Msg = entry.Message
	ls.Error = f.exactErrorField(entry.Data)
	ls.ExtraFields = f.exactExtraFields(entry.Data)
	// marshal
	_ = sonic.ConfigDefault.NewEncoder(entry.Buffer).Encode(ls)
	entry.Buffer.WriteByte('\n')
	ls.reset()
	logStructPool.Put(ls)
	return entry.Buffer.Bytes(), nil
}

func (f *selfFormatter) exactErrorField(data logrus.Fields) string {
	if len(data) == 0 {
		return ""
	}
	errInfo, ok := data[EntityFieldNameError.String()]
	if ok {
		callersFramesStr := CallersFrames2Str(getCallersFrames(errInfo.(error)))
		if len(callersFramesStr) > 0 {
			// 放入 entry.Data
			data[EntityFieldNameErrorStack.String()] = callersFramesStr
		}
		return fmt.Sprintf("%s", errInfo)
	}
	return ""
}

func (f *selfFormatter) exactExtraFields(data logrus.Fields) map[string]string {
	if len(data) == 0 {
		return nil
	}
	m := make(map[string]string, len(data))
	for k, v := range data {
		if k == EntityFieldNameError.String() {
			continue
		}
		m[k] = fmt.Sprintf("%s", v)
	}
	return m
}
