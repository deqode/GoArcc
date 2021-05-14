package main

import (
	"alfred/logger"
)

//todo : Alarm !!!!!! Do not touch the invocation sequence, either you might go through sleepless nights
// main : entry point
func main() {
	//logger initialize before app starts because in provider we need logger
	logger.InitLogger()
	GetApp().Run()
}
