package main

import (
	"alfred.sh/common/logger"
	"go.uber.org/zap"
)

//todo : Alarm !!!!!! Do not touch the invocation sequence, either you might go through sleepless nights
// main : entry point
func main() {
	//logger initialize before app starts because in provider we need logger
	logger.Init(logger.Config{
		LogLevel: zap.DebugLevel,
		Development: false,
	})
	GetApp().Run()
}
