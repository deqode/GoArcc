package middleware

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddInterceptors AddInterceptors: Add interceptors to the grpc server
func AddInterceptors(logger *zap.Logger, tracer opentracing.Tracer, opts []grpc.ServerOption) []grpc.ServerOption {
	// Make sure that log statements internals to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	//grpc recovery options
	recoveryOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(grpcPanicsRecovery),
	}

	// Add unary interceptor
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		//turns grpc panics into unknown error
		grpc_recovery.UnaryServerInterceptor(recoveryOptions...),
		//for context tags

		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),

		grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		//Adding prothesis monitoring
		grpc_prometheus.UnaryServerInterceptor,
		//zap logger implementation
		grpc_zap.UnaryServerInterceptor(logger),

		grpc_auth.UnaryServerInterceptor(AuthMiddleware),

		//validate the incoming request - inbound in proto file
		//If request is not correct the error will be sent to client
		grpc_validator.UnaryServerInterceptor(),
	))

	// Add stream interceptor (added as an example here)
	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_recovery.StreamServerInterceptor(recoveryOptions...),
		//context tag implementation
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		//open tracing implementation
		grpc_opentracing.StreamServerInterceptor(),
		//prom implementation
		grpc_prometheus.StreamServerInterceptor,
		grpc_auth.StreamServerInterceptor(AuthMiddleware),
		//zap implementation
		grpc_zap.StreamServerInterceptor(logger),

		//validate the incoming request - inbound in proto file
		//If request is not correct the error will be sent to client
		grpc_validator.StreamServerInterceptor(),

	))

	return opts
}

//grpcPanicsRecovery: is responsible to convert panic to the custom message
func grpcPanicsRecovery(in interface{}) error {
	return status.Errorf(codes.Unknown, "Unknown error")
}
