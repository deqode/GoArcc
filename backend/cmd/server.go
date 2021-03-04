package cmd

import (
	"alfred/config"
	grpql "alfred/protocol/graphql"
	"alfred/protocol/rest"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer(lc fx.Lifecycle, cfg *config.Config) {

	ctx := context.Background()
	cfg.GRPCPort = "8080"
	cfg.HTTPPort = "8081"
	cfg.GraphqlPort = "8082"
	cfg.LogLevel = int(zap.DebugLevel)

	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			// run HTTP gateway
			go func() {
				_ = rest.RunRESTServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
			}()
			//run graphql gateway
			go func() {
				_ = grpql.RunGraphqlServer(ctx, cfg.GraphqlPort)
			}()

			return nil
		},
	})

}
