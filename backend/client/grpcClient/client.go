package grpcClient

import (
	"alfred/config"
	"alfred/logger"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func GetGrpcClientConnection(config *config.Config) *grpc.ClientConn {
	//context setup for inner client call
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var conn *grpc.ClientConn
	conn, err := grpc.DialContext(ctx, config.DatastoreDBHost+":"+config.GRPCPort, grpc.WithInsecure())
	if err != nil {
		logger.Log.Fatal("did not connect", zap.Error(err))
	}
	return conn
}
