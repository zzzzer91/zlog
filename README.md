# zlog

## Usage

```go
import (
	"context"
	"github.com/zzzzer91/zlog/trace"
	"github.com/zzzzer91/zlog"
)

func main() {
	// 可选
	zlog.SetLogger(zlog.NewLogger(conf.App.Log, trace.NewTraceHook()))

	pwd, _ := os.Getwd()
	zlog.Ctx(context.Background()).Infof("%s starting, current env is %s, pwd is %s", conf.App.Name, pkg_conf.GetEnv(), pwd)
}
```