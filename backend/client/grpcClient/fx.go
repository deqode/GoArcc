package grpcClient

import "go.uber.org/fx"

// ConnectionFx GrpcClientConnectionFx : Provider of GetGrpcClientConnection
var ConnectionFx = fx.Options(
	fx.Provide(
		GetGrpcClientConnection,
	),
)
