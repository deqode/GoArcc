package user_profile

import (
	"alfred/modules/user-profile/v1/internals"
	"go.uber.org/fx"
)

// Module Providing client to the Global fx pipe line, so other
//module can use our client
var Module = fx.Options(
	fx.Provide(
		internals.NewUserProfileInternalServer,
	),
)
