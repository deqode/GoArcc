package AuthService

import (
	vcs "alfred/modules/vcs-connection/v1"
	"go.uber.org/fx"
)

// Module Providing client to the Global fx pipe line, so other
//module can use our client
var Module = fx.Options(
	fx.Provide(
		vcs.NewVCSConnectionServer,
	),
)
