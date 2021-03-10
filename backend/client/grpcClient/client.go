package grpcClient

import (
	"alfred/config"
	"alfred/logger"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ClientContext struct {
	Ctx context.Context
}

func GetGrpcClientConnection(config *config.Config) *grpc.ClientConn {
	/*	var opts []grpc.DialOption
		opts = append(
			opts, grpc.WithStreamInterceptor(
				otgrpc.OpenTracingStreamClientInterceptor(tracer, otgrpc.LogPayloads())))

		opts = append(opts, grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())))

		opts = append(opts, grpc.WithInsecure())*/
	// You must have some sort of OpenTracing Tracer instance on hand.
	conn, err := grpc.Dial(config.DatastoreDBHost+":"+config.GRPCPort, grpc.WithInsecure())
	if err != nil {
		logger.Log.Fatal("did not connect", zap.Error(err))
	}
	return conn
}
