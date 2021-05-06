package grpcClient

import "go.uber.org/fx"

//GrpcClientConnectionFx : Provider of GetGrpcClientConnection
var GrpcClientConnectionFx = fx.Options(
	fx.Provide(
		GetGrpcClientConnection,
	),
)
