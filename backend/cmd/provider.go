package cmd

import (
	"alfred/config"
	"alfred/db"
	"alfred/modules/HelloWorldService"
	"alfred/modules/UserProfileService"
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
		UserProfileService.Module,
		db.DatabaseConnectionFx,
		HelloWorldService.HelloServiceFx,
		cleanup.CleanupFx,
		jaeger.JaegerTracerFx,
	}
}
