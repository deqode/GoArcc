package graphql

import (
	"alfred/config"
	"alfred/logger"
	"context"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

func RunGraphqlServer(lc fx.Lifecycle, config *config.Config) {
	Ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	if err := RegisterGraphqlModules(Ctx, mux); err != nil {
		logger.Log.Fatal("not able to register graphql modules", zap.Error(err))
	}
	http.Handle("/graphql", mux)
	srv := &http.Server{
		Addr: config.Graphql.Host + ":" + config.Graphql.Port,
		// add handler with middleware
		Handler:/*middleware.AddRequestID(
		middleware.AddLogger(logger.Log, mux))*/mux,
	}

	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(ctx context.Context) error {

			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			// run HTTP gateway
			logger.Log.Info("starting HTTP/GRAPHQL  gateway...")
			go func() {
				_ = srv.ListenAndServe()
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Log.Info("graceful shutting down Graphql Server")
			_ = srv.Shutdown(Ctx)
			return nil
		},
	})
}
