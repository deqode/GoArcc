package main

import (
	"go.uber.org/zap"
	"goarcc/logger"
)

//todo : Alarm !!!!!! Do not touch the invocation sequence, either you might go through sleepless nights
// main : entry point
func main() {
	//logger initialize before app starts because in provider we need logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Take this level from config
		Development: false,
	})
	GetApp().Run()
}