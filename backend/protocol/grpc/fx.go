package grpc

import "go.uber.org/fx"

var InitGrpcBeforeServingFx = fx.Options(
	fx.Provide(
		InitGrpcBeforeServing,
	),
)
