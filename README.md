# zlog

## 功能

基于 logrus 的日志组件封装，支持打印错误栈、输出到文件和链路追踪。主要用于请求链路中的日志打印，基本的日志打印可以使用 [logx](https://github.com/zzzzer91/gopkg/tree/main/logx)。

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
    ctx := context.WithValue(context.Background(), zlog.EntityFieldNameRequestId, "abcdefghijk")
    ctx = context.WithValue(ctx, zlog.EntityFieldNameTraceId, "123456789")
    err := f(ctx)
    if err != nil {
        zlog.Ctx(ctx).WithError(err).Error("failed to execute f()")
    }
}

func f(ctx context.Context) error {
    return errors.New("new error")
}
```

输出：

```
{
  "time": "2023-03-11T11:22:26.801+08",
  "level": "error",
  "caller": "main.main:17",
  "msg": "failed to execute f()",
  "error": "new error",
  "extraFields": {
    "errorStack": ""main.f:22\nmain.main:15\nruntime.main:250"",
    "requestID": "abcdefghijk",
    "traceID": "123456789"
  }
}
```