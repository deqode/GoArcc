package vcs_connection

import (
	"alfred/modules/vcs-connection/v1/external-svc"
	"go.uber.org/fx"
)

// Module will be
var Module = fx.Options(
	fx.Provide(
		external_svc.NewVCSConnectionServer,
	),
)
