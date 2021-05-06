package logger

import "go.uber.org/fx"

// LoggerProviderFx : provide the zap logger instance with all the zap level.
var LoggerProviderFx = fx.Options(
	fx.Invoke(
		InitLogger,
	),
)
