package grpcClient

import (
	"alfred/config"
	"alfred/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func GetGrpcClientConnection(config *config.Config) *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:"+config.GRPCPort, grpc.WithInsecure())
	if err != nil {
		logger.Log.Fatal("did not connect", zap.Error(err))
	}
	return conn
}
