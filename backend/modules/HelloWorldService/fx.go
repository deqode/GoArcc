package HelloWorldService

import (
	"alfred/modules/HelloWorldService/pb"
	"go.uber.org/fx"
)

//Providing client to the Global fx pipe line, so other
//module can use our client
var HelloServiceFx = fx.Options(

	fx.Provide(
		pb.NewHelloWorldServiceClient,
	),
)
