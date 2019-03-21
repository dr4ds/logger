# Example
```go
package main

import (
	"github.com/dr4ds/logger"
)

func main() {
	Logger := logger.New(logger.LevelDebug, true)
	defer Logger.Close()

	Logger.Debugln("debug")
	Logger.Info("info\n")
	Logger.Successf("suc%s", "cess")
	Logger.Warning("warning\r\n")
	Logger.Errorln("error")
	Logger.Criticalf("critical %s", "exiting")
}
```