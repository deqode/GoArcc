package config

import "go.uber.org/fx"

// ProviderFx ConfigProviderFx : Provider of  GetConfig
var ProviderFx = fx.Options(
	fx.Provide(
		GetConfig,
		GetFileInformation,
	),
)
