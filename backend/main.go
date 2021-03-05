package main

import (
	"alfred/cmd"
	"alfred/config"
	"alfred/db"
	"alfred/logger"
	"alfred/modules/HelloWorldService"
	"alfred/modules/UserProfileService"
	"alfred/promthesiusServer"
	"alfred/protocol/grpc"
	"go.uber.org/fx"
)

//todo : Alarm !!!!!! Do not touch the invocation sequence, either you might go through sleepless nights
func main() {
	fx.New(
		config.ConfigProviderFx,
		grpc.InitGrpcBeforeServingFx,
		UserProfileService.Module,
		db.DatabaseConnectionFx,
		HelloWorldService.HelloServiceFx,
		promthesiusServer.InitPromthesiusServerFx,
		fx.Invoke(
			logger.InitLogger,
			//run server will run Rest , Graphql , Promthesius server
			cmd.RunServer,
			//all service got registered
			grpc.RunGRPCServer,
			grpc.RegisterGrpcModules,
			//After Registering Grpc Modules then only we can use promthesus
			promthesiusServer.PromthesiusRunner,
		),
	).Run()

}
