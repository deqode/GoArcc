package grpc

import (
	"alfred/logger"
	"alfred/protocol/grpc/middleware"
	"context"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

// RunServer runs gRPC service to publish ToDo service
func RunGRPCServer(ctx context.Context ,port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	//Register grpc modules
	RegisterGrpcModules(server)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}

