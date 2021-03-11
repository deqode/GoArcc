package cmd

import (
	"alfred/servers/cleanup"
	"alfred/servers/grpc"
	"alfred/servers/healthcheck"
	"go.uber.org/fx"
)

/*
   Todo Here you can new invokers.
   Todo Alarm !!!!!! Do not touch the invocation sequence, either you might go through sleepless nights
*/

//GetInvokersOptions: Please do not change the sequence because it invoker is lifecycle based method .
// So changing the sequence will be harmful.
func GetInvokersOptions() fx.Option {
	return fx.Invoke(
		//run server will run Rest , Graphql , prometheus server
		RunServer,
		//all service got registered
		grpc.RunGRPCServer,
		grpc.RegisterGrpcModules,
		//After Registering Grpc Modules then only we can use prometheus
		//prometheusServer.PrometheusRunner,
		//run cleanup code after closing the server
		//Add Health check
		healthcheck.HealthCheckRunner,
		cleanup.Cleanup,
	)
}
