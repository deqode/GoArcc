package HelloWorldService

import (
	"alfred/modules/HelloWorldService/pb"
	"go.uber.org/fx"
)

var HelloServiceFx = fx.Options(

	fx.Provide(
		pb.NewHelloWorldServiceClient,
	),
)
