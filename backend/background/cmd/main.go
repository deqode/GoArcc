package main

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	"alfred/background/config"
	"alfred/background/workers/workflows"
	"fmt"
)

func main() {
	fmt.Println("Starting Worker..")
	var cadenceAppConfig background.CadenceAppConfig
	cadenceAppConfig.Setup()
	var cadenceClient cadenceAdapter.CadenceAdapter
	cadenceClient.CadenceSetup(&cadenceAppConfig.Cadence)
    workflows.StartWorkers(&cadenceClient)
	// The workers are supposed to be long running process that should not exit.
	select {}
}

//



