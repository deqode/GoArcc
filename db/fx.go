package db

import "go.uber.org/fx"

// DatabaseConnectionFx : constructor provide a initialised db connection with the given port.
// No need to close the connection of the db because we have a single pool with single db instance.
var DatabaseConnectionFx = fx.Options(
	fx.Provide(
		NewConnection,
	),
)
