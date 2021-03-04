package db

import "go.uber.org/fx"

//Database connection constructor
var DatabaseConnectionFx = fx.Options(
	fx.Provide(
		NewConnection,
	),
)
