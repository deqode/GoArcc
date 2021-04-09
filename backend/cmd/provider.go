package cmd

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	user_profile "alfred/modules/UserProfileService/v1"
	"alfred/servers/cleanup"
	"alfred/servers/grpc"
	"alfred/servers/openTracing/tracer/jaeger"
	"go.uber.org/fx"
)

//ProviderOptions: Global Constructor. Sequence in fx does not matter.
func GetProviderOptions() []fx.Option {
	return []fx.Option{
		config.ConfigProviderFx,
		grpc.InitGrpcBeforeServingFx,
		user_profile.Module,
		db.DatabaseConnectionFx,
		cleanup.CleanupFx,
		jaeger.JaegerTracerFx,
		grpcClient.GrpcClientConnectionFx,
	}
}
