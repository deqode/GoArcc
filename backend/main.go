package  main

import (
	"alfred/cmd"
	"alfred/config"
	"alfred/logger"
	"alfred/protocol/grpc"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		 config.Module,
		 grpc.Module,
		 fx.Invoke(
		 	 logger.InitLogger,
             cmd.RunServer,
             grpc.RunGRPCServer,
             grpc.RegisterGrpcModules,
		 	),
		).Run()
}
