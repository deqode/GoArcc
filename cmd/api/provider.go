package main

import (
	"github.com/deqode/GoArcc/client/grpcClient"
	"github.com/deqode/GoArcc/config"
	"github.com/deqode/GoArcc/db"
	"github.com/deqode/GoArcc/servers/cleanup"
	"github.com/deqode/GoArcc/servers/grpc"
	"github.com/deqode/GoArcc/servers/openTracing/tracer/jaeger"
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
