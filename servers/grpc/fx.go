package grpc

import "go.uber.org/fx"

// InitGrpcBeforeServingFx : constructor for InitGrpcBeforeServing method.
var InitGrpcBeforeServingFx = fx.Options(
	fx.Provide(
		InitGrpcBeforeServing,
	),
)
