package grpc

import (
	"alfred/config"
	"alfred/logger"
	"alfred/protocol/grpc/middleware"
	"context"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"sync"
)

var (
	server *grpc.Server
	// onceInit guarantee initialize logger only once
	onceInit sync.Once
)

func InitGrpcServer(srv *grpc.Server) {
	return

}

func InitGrpcBeforeServing( config *config.Config) (*grpc.Server , net.Listener , error) {
	listen, err := net.Listen("tcp", ":"+ config.GRPCPort)
	if err != nil {
		return nil , nil , err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	InitGrpcServer(server)
   return server , listen , nil
}

// RunServer runs gRPC service to publish ToDo service
func RunGRPCServer(server *grpc.Server , listener net.Listener ) error {
	ctx := context.Background()
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
	return server.Serve(listener)
}

