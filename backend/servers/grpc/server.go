package grpc

import (
	"alfred/config"
	"alfred/logger"
	"alfred/servers/grpc/middleware"
	"context"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func InitGrpcBeforeServing(config *config.Config, tracer opentracing.Tracer) (*grpc.Server, net.Listener) {
	listen, err := net.Listen("tcp", ":"+config.GRPCPort)
	if err != nil {
		logger.Log.Fatal("not able to initialize grpc server", zap.Error(err))
		return nil, nil
	}
	// gRPC server statup options
	opts := []grpc.ServerOption{}
	// add logging middleware
	opts = middleware.AddInterceptors(logger.Log, tracer, opts)
	// register service
	server := grpc.NewServer(opts...)
	return server, listen
}

// RunServer runs gRPC service to publish ToDo service
func RunGRPCServer(lc fx.Lifecycle, server *grpc.Server, listener net.Listener) error {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// start gRPC server
				go server.Serve(listener)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				// start gRPC server
				logger.Log.Info("stopping  gRPC server...")
				server.GracefulStop()
				return nil
			},
		},
	)

	return nil
}
