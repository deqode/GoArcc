package config

import "go.uber.org/fx"

var ConfigProvider = fx.Options(
	 fx.Provide(
	 	 GetConfig,
	 	),
	)
