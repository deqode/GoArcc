package graphql

import (
	"context"
	"goarcc/config"
	"goarcc/logger"
	"goarcc/servers/graphql/middleware"
	"net/http"
	"time"

	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RunGraphqlServer : Will start the graphql server.
func RunGraphqlServer(lc fx.Lifecycle, config *config.Config, conn *grpc.ClientConn) {
	Ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	if err := RegisterGraphqlModules(mux, conn); err != nil {
		logger.Log.Fatal("not able to register graphql modules", zap.Error(err))
	}
	http.Handle("/graphql", mux)
	srv := &http.Server{
		Addr:         config.Graphql.Host + ":" + config.Graphql.Port,
		WriteTimeout: time.Second * time.Duration(config.Graphql.RequestTimeout),
		// add handler with middleware
		Handler: http.TimeoutHandler(middleware.ChangeContext(middleware.AddCors(middleware.AddRequestID(
			middleware.AddLogger(logger.Log, mux)))), time.Second*time.Duration(config.Graphql.RequestTimeout), "Context deadline exceeded"),
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
