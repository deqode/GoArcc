package cleanup

import "go.uber.org/fx"

/*
  CleanupFx: Providing constructor for cleanup
*/
var CleanupFx = fx.Options(
	fx.Provide(
		GetCleanupConfig,
	),
)
