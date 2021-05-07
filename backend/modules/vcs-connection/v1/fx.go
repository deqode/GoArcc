package vcs_connection

import (
	"go.uber.org/fx"
)

// Module will be
var Module = fx.Options(
	fx.Provide(
		NewVCSConnectionServer,
	),
)
