package main

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/servers/cleanup"
	"alfred/servers/grpc"
	"alfred/servers/openTracing/tracer/jaeger"
	"go.uber.org/fx"
)

// GetProviderOptions ProviderOptions: Global Constructor.
// Sequence in fx does not matter,  So you can write in any order you want.
func GetProviderOptions() []fx.Option {
	return []fx.Option{
		config.ProviderFx,
		grpc.InitGrpcBeforeServingFx,
		db.DatabaseConnectionFx,
		cleanup.CleanupFx,
		jaeger.JaegerTracerFx,
		grpcClient.ConnectionFx,
	}
}
