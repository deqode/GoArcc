package grpcClient

import (
	"alfred/config"
	"alfred/logger"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)


func GetGrpcClientConnection(config *config.Config) *grpc.ClientConn {
	// You must have some sort of OpenTracing Tracer instance on hand.
	conn , err := grpc.Dial(config.DatastoreDBHost+":"+config.GRPCPort ,
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(grpc_opentracing.StreamClientInterceptor()), grpc.WithInsecure())
	if err != nil {
		logger.Log.Fatal("did not connect", zap.Error(err))
	}
	return conn
}

