package main

import (
	"github.com/kris-nova/logger"
	lol "github.com/kris-nova/lolgopher"
)

func main() {
	//
	logger.Writer = lol.NewLolWriter()          // Sometimes this will work better
	logger.Writer = lol.NewTruecolorLolWriter() // Comment one of these out
	//

	logger.BitwiseLevel = logger.LogEverything
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Always(".....................")
	logger.Debug("Debug logging")
	logger.Critical("Stderr logging")
}
