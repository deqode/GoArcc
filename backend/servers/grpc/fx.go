package grpc

import "go.uber.org/fx"

// InitGrpcBeforeServingFx InitGrpcBeforeServingFx: Constructors present in the pipeline.
var InitGrpcBeforeServingFx = fx.Options(
	fx.Provide(
		InitGrpcBeforeServing,
	),
)
