package grpc

import (
	"alfred/config"
	"alfred/logger"
	"alfred/protocol/grpc/middleware"
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

func InitGrpcBeforeServing(config *config.Config) (*grpc.Server, net.Listener, error) {
	listen, err := net.Listen("tcp", ":"+config.GRPCPort)
	if err != nil {
		return nil, nil, err
	}
	// gRPC server statup options
	opts := []grpc.ServerOption{}
	// add logging middleware
	opts = middleware.AddInterceptors(logger.Log, opts)
	// register service
	server := grpc.NewServer(opts...)
	return server, listen, nil
}

// RunServer runs gRPC service to publish ToDo service
func RunGRPCServer(lc fx.Lifecycle, server *grpc.Server, listener net.Listener) error {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// start gRPC server
				logger.Log.Info("starting gRPC server...")
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
