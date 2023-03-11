# zlog

## 功能

基于 logrus 的日志组件封装，支持打印错误栈、输出到文件和链路追踪。

## Usage

执行：

```go
package main

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/zzzzer91/zlog/trace"
	"github.com/zzzzer91/zlog"
)

func main() {
	// 可选，加入链路追踪，然后可以在 jager 中看到日志
	// zlog.SetLogger(zlog.NewLogger(conf.App.Log, trace.NewTraceHook()))

	pwd, _ := os.Getwd()
	zlog.Ctx(context.Background()).Infof("%s starting, current env is %s, pwd is %s", "hello", "pro", pwd)
	ctx := context.WithValue(context.Background(), zlog.EntityFieldNameRequestId, "abcdefghijk")
	ctx = context.WithValue(ctx, zlog.EntityFieldNameTraceId, "123456789")
	err := f(ctx)
	zlog.Ctx(ctx).WithError(err).Error("failed to execute f()")
}

func f(ctx context.Context) error {
	return errors.New("new error")
}
```

输出：

```
{
  "time": "2023-03-11T11:22:26.801+08",
  "level": "info",
  "file": "main.main:16",
  "msg": "hello starting, current env is pro, pwd is /Users/zzzzer/go/src/github.com/zzzzer91/go-test"
}
{
  "time": "2023-03-11T11:22:26.801+08",
  "level": "error",
  "file": "main.main:20",
  "msg": "failed to execute f()",
  "error": "new error",
  "extraFields": {
    "errorStack": "/Users/zzzzer/go/src/github.com/zzzzer91/go-test/main.go:24\n\tf\n/Users/zzzzer/go/src/github.com/zzzzer91/go-test/main.go:19\n\tmain\n/usr/local/go/src/runtime/proc.go:250\n\tmain",
    "requestID": "abcdefghijk",
    "traceID": "123456789"
  }
}
```