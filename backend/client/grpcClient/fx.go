package grpcClient

import "go.uber.org/fx"

var GrpcClientFx = fx.Options(
	fx.Provide(
		GetGrpcClientConnection,
	),
)
