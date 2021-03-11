package cmd

import (
	"alfred/config"
	"alfred/logger"
	grpql "alfred/servers/graphql"
	"alfred/servers/rest"
	"context"
	"go.uber.org/fx"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer(lc fx.Lifecycle, cfg *config.Config) {

	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(ctx context.Context) error {

			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			// run HTTP gateway
			go func() {
				_ = rest.RunRESTServer(ctx, cfg)
			}()
			//run graphql gateway
			go func() {
				_ = grpql.RunGraphqlServer(ctx, cfg)
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Log.Info("Gracefully Shutting down graphql , rest ")
			return nil
		},
	})

}
