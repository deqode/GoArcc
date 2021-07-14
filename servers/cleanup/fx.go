package cleanup

import "go.uber.org/fx"

// CleanupFx will provide the constructor method  GetCleanupConfig.
var CleanupFx = fx.Options(
	fx.Provide(
		GetCleanupConfig,
	),
)
