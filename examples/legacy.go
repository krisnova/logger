package main

import "github.com/kris-nova/novaarchive/logger"

func main() {

	logger.Level = 4
	logger.Always("Always logging")
	logger.Info("Info logging")
	logger.Debug("Debug logging")

}
