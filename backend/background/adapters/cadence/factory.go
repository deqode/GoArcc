package cadenceAdapter

import (
	"errors"

	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
	"go.uber.org/zap"
)

const (
	_cadenceClientName      = "cadence-client"
	_cadenceFrontendService = "cadence-frontend"
)

// WorkflowClientBuilder build client to cadence service
type WorkflowClientBuilder struct {
	hostPort     string
	domain       string
	dispatcher   *yarpc.Dispatcher
	metricsScope tally.Scope
	Logger       *zap.Logger
}

// NewBuilder creates a new WorkflowClientBuilder
func NewBuilder(logger *zap.Logger, hostPort string, domain string) *WorkflowClientBuilder {
	return &WorkflowClientBuilder{
		hostPort:     hostPort,
		domain:       domain,
		metricsScope: tally.NoopScope,
		Logger:       logger,
	}
}

// BuildCadenceClient builds a client to cadence service
func (b *WorkflowClientBuilder) BuildCadenceClient() (client.Client, error) {
	service, err := b.BuildServiceClient()
	if err != nil {
		return nil, err
	}

	return client.NewClient(
		service, b.domain, &client.Options{MetricsScope: b.metricsScope}), nil
}

// BuildCadenceDomainClient builds a domain client to cadence service
func (b *WorkflowClientBuilder) BuildCadenceDomainClient() (client.DomainClient, error) {
	service, err := b.BuildServiceClient()
	if err != nil {
		return nil, err
	}

	return client.NewDomainClient(
		service, &client.Options{MetricsScope: b.metricsScope}), nil
}

// BuildServiceClient builds a rpc service client to cadence service
func (b *WorkflowClientBuilder) BuildServiceClient() (workflowserviceclient.Interface, error) {
	if err := b.build(); err != nil {
		return nil, err
	}

	if b.dispatcher == nil {
		b.Logger.Fatal("No RPC dispatcher provided to create a connection to Cadence Service")
	}

	return workflowserviceclient.New(b.dispatcher.ClientConfig(_cadenceFrontendService)), nil
}

func (b *WorkflowClientBuilder) build() error {
	if b.dispatcher != nil {
		return nil
	}

	if len(b.hostPort) == 0 {
		return errors.New("HostPort is empty")
	}

	ch, err := tchannel.NewChannelTransport(
		tchannel.ServiceName(_cadenceClientName))
	if err != nil {
		b.Logger.Fatal("Failed to create transport channel", zap.Error(err))
	}

	b.Logger.Debug("Creating RPC dispatcher outbound",
		zap.String("ServiceName", _cadenceFrontendService),
		zap.String("HostPort", b.hostPort))

	b.dispatcher = yarpc.NewDispatcher(yarpc.Config{
		Name: _cadenceClientName,
		Outbounds: yarpc.Outbounds{
			_cadenceFrontendService: {Unary: ch.NewSingleOutbound(b.hostPort)},
		},
	})

	if b.dispatcher != nil {
		if err := b.dispatcher.Start(); err != nil {
			b.Logger.Fatal("Failed to create outbound transport channel: %v", zap.Error(err))
		}
	}

	return nil
}
