package config

import "go.uber.org/fx"

//Initializes config globally in the pipeline
var ConfigProviderFx = fx.Options(
	fx.Provide(
		GetConfig,
	),
)
