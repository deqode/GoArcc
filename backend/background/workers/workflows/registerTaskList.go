package workflows

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

// GitCloneGroup TaskListName This is registration process where you register all your workflows
// and activity function handlers.
const (
	GitCloneGroup = "gitCloneGroup"
)

//add string constant here
func getAllWorkers() []string {
	return []string{
		GitCloneGroup,
	}
}



func StartWorkers(h *cadenceAdapter.CadenceAdapter) {
	taskList := getAllWorkers()
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}
	for _, taskName := range taskList {
			cadenceWorker := worker.New(h.ServiceClient, h.Config.Domain, taskName, workerOptions)
			err := cadenceWorker.Start()
			if err != nil {
				h.Logger.Error("Failed to start worker:"+taskName, zap.Error(err))
			}

	}
}
