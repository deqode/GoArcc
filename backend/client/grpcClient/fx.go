package grpcClient

import "go.uber.org/fx"

var GrpcClientConnectionFx = fx.Options(
	fx.Provide(
		GetGrpcClientConnection,
	),
)
