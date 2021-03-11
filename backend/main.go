package main

import (
	"alfred/cmd"
	"alfred/config"
	"alfred/db"
	"alfred/logger"
	"alfred/modules/HelloWorldService"
	"alfred/modules/UserProfileService"
	"alfred/servers/cleanup"
	"alfred/servers/grpc"
	"alfred/servers/healthcheck"
	"alfred/servers/openTracing/tracer/jaeger"
	"alfred/servers/promthesiusServer"
	"go.uber.org/fx"
)

//todo : Alarm !!!!!! Do not touch the invocation sequence, either you might go through sleepless nights
func main() {
	//logger initialize before app starts because in provider we need logger
	logger.InitLogger()
	fx.New(
		config.ConfigProviderFx,
		grpc.InitGrpcBeforeServingFx,
		UserProfileService.Module,
		db.DatabaseConnectionFx,
		HelloWorldService.HelloServiceFx,
		promthesiusServer.InitPromthesiusServerFx,
		cleanup.CleanupFx,
		jaeger.JaegerTracerFx,

		fx.Invoke(
			//run server will run Rest , Graphql , prometheus server
			cmd.RunServer,
			//all service got registered
			grpc.RunGRPCServer,
			grpc.RegisterGrpcModules,
			//After Registering Grpc Modules then only we can use prometheus
			promthesiusServer.PrometheusRunner,
			//run cleanup code after closing the server
			//Add Health check
			healthcheck.HealthCheckRunner,
			cleanup.Cleanup,
		),
	).Run()

}
