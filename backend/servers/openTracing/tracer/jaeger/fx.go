package jaeger

import "go.uber.org/fx"

// JaegerTracerFx : Constructor method for InitJaeger.
var JaegerTracerFx = fx.Options(
	fx.Provide(
		InitJaeger,
	),
)
