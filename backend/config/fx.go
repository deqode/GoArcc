package config

import "go.uber.org/fx"

var ConfigProviderFx = fx.Options(
	fx.Provide(
		GetConfig,
	),
)
