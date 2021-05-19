package cadenceAdapter

import (
	commonLogger "alfred.sh/common/logger"
	"alfred/config"
	"context"
	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/encoded"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

type (
	// CadenceAdapter Adapter class for workflow helper.
	CadenceAdapter struct {
		CadenceClient  client.Client
		ServiceClient  workflowserviceclient.Interface
		Scope          tally.Scope
		Logger         *zap.Logger
		Config         config.CadenceConfig
		Builder        *WorkflowClientBuilder
		DataConverter  encoded.DataConverter
		CtxPropagators []workflow.ContextPropagator
	}
)

// CadenceSetup Setup setup the config for the code run
func (h *CadenceAdapter) CadenceSetup(config *config.CadenceConfig) {
	if h.CadenceClient != nil {
		return
	}
	logger, _ := zap.NewDevelopment()
	h.Logger = logger
	h.Config = *config
	hostPort := h.Config.Host + ":" + h.Config.Port
	domainName := h.Config.Domain

	h.Builder = NewBuilder(logger, hostPort, domainName)

	// Cadence client used to perform CRUD operation.
	cadenceClient, err := h.Builder.BuildCadenceClient()
	if err != nil {
		panic(err)
	}
	h.CadenceClient = cadenceClient

	// Service client that communicates with cadence using the rpc.
	serviceClient, err := h.Builder.BuildServiceClient()
	if err != nil {
		panic(err)
	}
	h.ServiceClient = serviceClient

	domainClient, _ := h.Builder.BuildCadenceDomainClient()
	_, err = domainClient.Describe(context.Background(), domainName)
	if err != nil {
		commonLogger.Log.Info("Domain doesn't exist", zap.String("Domain", domainName), zap.String("HostPort", hostPort), zap.Error(err))
	} else {
		commonLogger.Log.Info("Domain successfully registered.", zap.String("Domain", domainName), zap.String("HostPort", hostPort))
	}
}
