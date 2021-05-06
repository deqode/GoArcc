package config

import "go.uber.org/fx"

// ConfigProviderFx : Provider of  GetConfig
var ConfigProviderFx = fx.Options(
	fx.Provide(
		GetConfig,
	),
)
