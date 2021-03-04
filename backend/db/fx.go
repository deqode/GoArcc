package db

import "go.uber.org/fx"

var DatabaseConnectionFx = fx.Options(
	fx.Provide(
		NewConnection,
	),
)
