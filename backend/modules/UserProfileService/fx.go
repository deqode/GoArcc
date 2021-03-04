package UserProfileService

import (
	"alfred/modules/UserProfileService/pb"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		pb.NewUserProfileServiceClient,
	),
)
