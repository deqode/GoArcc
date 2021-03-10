package jaeger

import "go.uber.org/fx"

var JaegerTracerFx = fx.Options(
	fx.Provide(
		InitJaeger,
	),
)
