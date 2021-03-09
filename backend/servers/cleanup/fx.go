package cleanup

import "go.uber.org/fx"

var CleanupFx = fx.Options(
	fx.Provide(
		GetCleanupConfig,
	),
)
