package cmd

import (
	"go.uber.org/fx"
)

// CadenceSetupFx ConfigProviderFx: config globally in the pipeline stream
var CadenceSetupFx = fx.Options(
	fx.Provide(
		CadenceSetup,
	),
)

