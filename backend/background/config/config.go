package background

import (
	"alfred/config"
	"go.uber.org/zap"
)

type CadenceAppConfig struct {
	Env            string
	WorkerTaskList string
	Cadence        config.CadenceConfig
	Logger         *zap.Logger
}

// Setup setup the config for the code run
func (h *CadenceAppConfig) Setup() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	h.Logger = logger
	logger.Debug("Finished loading Configuration!")
}

//mockgen -source=account_internal_grpc.pb.go -destination=account_internal_mock.go -package=pb
