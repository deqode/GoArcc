package grpcClient

import (
	"alfred/config"
	"alfred/logger"
	"alfred/servers/openTracing/tracer/jaeger"
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ClientContext struct {
	Ctx context.Context
}

func GetGrpcClientConnection(ctx context.Context, config *config.Config) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(jaeger.Tracer)),
		),
		grpc.WithStreamInterceptor(grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTracer(jaeger.Tracer))),
	)
	//append grpc insecure
	opts = append(opts,
		grpc.WithInsecure(),
	)
	// You must have some sort of OpenTracing Tracer instance on hand.
	conn, err := grpc.DialContext(ctx, config.Grpc.Host+":"+config.Grpc.Port, opts...)
	if err != nil {
		logger.Log.Fatal("did not connect", zap.Error(err))
	}
	return conn
}
