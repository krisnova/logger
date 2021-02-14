package main

import (
	"os"

	"github.com/kris-nova/novaarchive/logger"
)

func main() {

	logger.BitwiseLevel = logger.LogEverything
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Debug("Debug logging")

	logger.Writer = os.Stderr
	logger.Critical("Stderr logging")
}
