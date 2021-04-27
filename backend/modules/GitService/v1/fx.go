package GitService

import (
	"alfred/modules/GitService/v1/pb"
	"go.uber.org/fx"
)

//Providing client to the Global fx pipe line, so other
//module can use our client
var Module = fx.Options(
	fx.Provide(
		pb.NewGitServiceClient,
	),
)
