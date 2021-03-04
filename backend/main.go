package main

import (
	"alfred/client/grpcClient"
	"alfred/cmd"
	"alfred/config"
	"alfred/db"
	"alfred/logger"
	"alfred/protocol/grpc"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.ConfigProviderFx,
		grpc.InitGrpcBeforeServingFx,
		db.DatabaseConnectionFx,
		fx.Invoke(
			logger.InitLogger,
			cmd.RunServer,
			grpc.RunGRPCServer,
			grpc.RegisterGrpcModules,
			grpcClient.GetGrpcClientConnection,
		),
	).Run()

}
