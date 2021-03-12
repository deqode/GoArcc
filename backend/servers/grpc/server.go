package grpc

import (
	"alfred/config"
	"alfred/logger"
	"alfred/servers/grpc/middleware"
	"context"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

//InitGrpcBeforeServing: give the instance of server and listener.
func InitGrpcBeforeServing(config *config.Config, tracer opentracing.Tracer) (*grpc.Server, net.Listener) {
	listen, err := net.Listen("tcp", config.Grpc.Host+":"+config.Grpc.Port)
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

// RunGRPCServer runs gRPC service on the given port.
func RunGRPCServer(lc fx.Lifecycle, server *grpc.Server, listener net.Listener) error {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				reflection.Register(server)
				grpc_prometheus.EnableHandlingTimeHistogram()
				grpc_prometheus.Register(server)
				http.Handle("/metrics", promhttp.Handler())
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
