package rest

import (
	"alfred/config"
	"alfred/logger"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"net/http"
)

func RunRestServer(lc fx.Lifecycle, config *config.Config) {
	Ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	if err := RegisterRESTModules(Ctx, mux, config); err != nil {
		panic(err)
	}
	srv := &http.Server{
		Addr: config.Rest.Host + ":" + config.Rest.Port,
		// add handler with middleware
		Handler:/* middleware.AddRequestID(
		middleware.AddLogger(logger.Log, mux))*/mux,
	}

	logger.Log.Info("starting HTTP/REST gateway...")

	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(ctx context.Context) error {

			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			// run HTTP gateway
			logger.Log.Info("starting HTTP/REST gateway...")
			go func() {
				_ = srv.ListenAndServe()
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Log.Info("Gracefully Shutting down REST server")
			_ = srv.Shutdown(Ctx)
			return nil
		},
	})
}
