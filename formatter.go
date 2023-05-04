package zlog

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/bytedance/sonic"

	"github.com/sirupsen/logrus"
)

// selfFormatter 自定义日志格式
type selfFormatter struct {
	timeFormat string
}

type logStruct struct {
	Time        string            `json:"time"`
	Level       string            `json:"level"`
	Caller      string            `json:"caller,omitempty"`
	Msg         string            `json:"msg,omitempty"`
	Error       string            `json:"error,omitempty"`
	ExtraFields map[string]string `json:"extraFields,omitempty"`
}

var (
	logStructPool = sync.Pool{
		New: func() interface{} {
			return &logStruct{}
		},
	}
)

func (ls *logStruct) reset() {
	ls.Time = ""
	ls.Level = ""
	ls.Caller = ""
	ls.Msg = ""
	ls.Error = ""
	ls.ExtraFields = nil
}

// Format 实现 formatter 定义的接口，自定义日志格式
func (f *selfFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	ls := logStructPool.Get().(*logStruct)
	defer logStructPool.Put(ls)
	defer ls.reset()
	ls.Time = entry.Time.Format(f.timeFormat)
	ls.Level = entry.Level.String()
	ls.Caller = entry.Caller.Function + ":" + strconv.Itoa(entry.Caller.Line)
	ls.Msg = entry.Message
	ls.Error = f.exactErrorField(entry.Data)
	ls.ExtraFields = f.exactExtraFields(entry.Data)
	// marshal
	_ = sonic.ConfigDefault.NewEncoder(entry.Buffer).Encode(ls)
	return entry.Buffer.Bytes(), nil
}

func (f *selfFormatter) exactErrorField(data logrus.Fields) string {
	if len(data) == 0 {
		return ""
	}
	if errInfo, ok := data[EntityFieldNameError.String()]; ok {
		if _, ok := data[EntityFieldNameErrorStack.String()]; !ok {
			callersFramesStr := CallersFrames2Str(getCallerFramesFromError(errInfo.(error)))
			if callersFramesStr != "" {
				// 放入 entry.Data
				data[EntityFieldNameErrorStack.String()] = callersFramesStr
			}
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
