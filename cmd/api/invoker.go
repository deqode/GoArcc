package main

import (
	"go.uber.org/fx"
	"goarcc/servers/cleanup"
	"goarcc/servers/graphql"
	"goarcc/servers/grpc"
	"goarcc/servers/healthcheck"
	"goarcc/servers/rest"
)

/*
   Todo Here you can new invokers.
   Todo Alarm !!!!!! Do not touch the invocation sequence, either you might go through sleepless nights
*/

// GetInvokersOptions GetInvokersOptions: Please do not change the sequence because it invoker is lifecycle based method .
// So changing the sequence will be harmful.
func GetInvokersOptions() fx.Option {
	return fx.Invoke(
		//run server will run Rest , Graphql , prometheus server
		//RunServer,
		//all service got registered
		grpc.RunGRPCServer,
		grpc.RegisterGrpcModules,
		rest.RunRestServer,
		graphql.RunGraphqlServer,
		//After Registering Grpc Modules then only we can use prometheus
		//prometheusServer.PrometheusRunner,
		//run cleanup code after closing the server
		//Add Health check
		healthcheck.HealthCheckRunner,
		cleanup.Cleanup,
	)
}