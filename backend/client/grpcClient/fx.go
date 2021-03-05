package grpcClient

import "go.uber.org/fx"

var GrpcClientFx = fx.Options(
	fx.Provide(
		GetGrpcClientConnectionContext,
	),
	fx.Invoke(
		GetGrpcClientConnection,
	),
)
