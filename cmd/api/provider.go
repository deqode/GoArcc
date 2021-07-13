package main

import (
	"go.uber.org/fx"
	"goarcc/client/grpcClient"
	"goarcc/config"
	"goarcc/db"
	"goarcc/servers/cleanup"
	"goarcc/servers/grpc"
	"goarcc/servers/openTracing/tracer/jaeger"
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
