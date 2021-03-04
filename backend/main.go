package main

import (
	"alfred/cmd"
	"alfred/config"
	"alfred/db"
	"alfred/logger"
	"alfred/modules/HelloWorldService"
	"alfred/modules/UserProfileService"
	"alfred/protocol/grpc"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.ConfigProviderFx,
		grpc.InitGrpcBeforeServingFx,
		UserProfileService.Module,
		db.DatabaseConnectionFx,
		HelloWorldService.HelloServiceFx,
		fx.Invoke(
			logger.InitLogger,
			cmd.RunServer,
			grpc.RunGRPCServer,
			grpc.RegisterGrpcModules,
		),
	).Run()

}
