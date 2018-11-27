# logger

Ported from it's [original location](https://github.com/kubicorn/kubicorn/tree/master/pkg/logger) in the Kubicorn code base.

Simple golang logger

```go
package main

import (
	"github.com/kris-nova/logger"
	"fmt"
	"os"
)


func main(){

	// Most Verbose
	//logger.Level = 4

	// Normal
	// No info or debug messages, only warnings and criticals
	logger.Level = 2

	// Off
	//logger.Level = 0

	err := fmt.Errorf("New error")
	logger.Info("we found an error: %v", err)

	logger.Debug("this is a useful message for software enigneers")

	logger.Warning("something bad happened but the software can still run")

	// Notice this does *NOT* exit!
	logger.Critical("the software should stop running, this is bad")

	// Now we have to exit
	os.Exit(123)
}

```