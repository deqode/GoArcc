package vcs_connection

import (
	"alfred/modules/AuthService/v1/pb"
	"go.uber.org/fx"
)

// Module will be
var Module = fx.Options(
	fx.Provide(
		pb.NewUserLoginServiceClient,
	),
)
