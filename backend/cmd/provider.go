package cmd

import (
	"alfred/background/cmd"
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	git_service "alfred/modules/GitService/v1"
	user_profile "alfred/modules/UserProfileService/v1"
	"alfred/servers/cleanup"
	"alfred/servers/grpc"
	"alfred/servers/openTracing/tracer/jaeger"
	"go.uber.org/fx"
)

// GetProviderOptions ProviderOptions: Global Constructor. Sequence in fx does not matter.
func GetProviderOptions() []fx.Option {
	return []fx.Option{
		config.ConfigProviderFx,
		grpc.InitGrpcBeforeServingFx,
		user_profile.Module,
		git_service.Module,
		db.DatabaseConnectionFx,
		cleanup.CleanupFx,
		jaeger.JaegerTracerFx,
		grpcClient.GrpcClientConnectionFx,
		cmd.CadenceSetupFx,
	}
}
