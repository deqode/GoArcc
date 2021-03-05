package grpcClient

import (
	"alfred/config"
	"alfred/logger"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

type ClientContext struct {
	Ctx context.Context
}

func GetGrpcClientConnection(config *config.Config, clientContext *ClientContext) *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.DialContext(clientContext.Ctx, config.DatastoreDBHost+":"+config.GRPCPort, grpc.WithInsecure())
	if err != nil {
		logger.Log.Fatal("did not connect", zap.Error(err))
	}
	return conn
}

//getting the client connection context
func GetGrpcClientConnectionContext() *ClientContext {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	return &ClientContext{
		Ctx: ctx,
	}
}
