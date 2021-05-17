package cmd

import (
	"alfred.sh/common/logger"
	cadenceAdapter "alfred/background/adapters/cadence"
	"alfred/background/config"
	"alfred/background/workers/workflows"
	"alfred/config"
	"context"
	"go.uber.org/fx"
)

//CadenceSetup :
func CadenceSetup(config *config.Config) (*background.CadenceAppConfig, *cadenceAdapter.CadenceAdapter) {
	logger.Log.Info("Cadence Setup started....")
	cadenceAppConfig := background.CadenceAppConfig{
		Cadence: config.CadenceConfig,
	}
	cadenceAppConfig.Setup()
	var cadenceClient cadenceAdapter.CadenceAdapter
	cadenceClient.CadenceSetup(&cadenceAppConfig.Cadence)
	logger.Log.Info("Cadence Setup finished....")
	return &cadenceAppConfig, &cadenceClient
}

func StartBackground(lc fx.Lifecycle, config *background.CadenceAppConfig, adapter *cadenceAdapter.CadenceAdapter) {
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		logger.Log.Info("Starting Background....")
		workflows.StartWorkers(adapter)
		return nil
	}})
}
