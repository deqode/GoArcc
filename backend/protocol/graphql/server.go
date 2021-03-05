package graphql

import (
	"alfred/config"
	"alfred/logger"
	"alfred/protocol/graphql/middleware"
	"context"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// RunServer runs HTTP/REST gateway
func RunGraphqlServer(ctx context.Context, config *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	if err := RegisterGraphqlModules(ctx, mux); err != nil {
		logger.Log.Fatal("not able to register graphql modules", zap.Error(err))
	}
	http.Handle("/graphql", mux)
	srv := &http.Server{
		Addr: config.ServerHost + ":" + config.GraphqlPort,
		// add handler with middleware
		Handler: middleware.AddRequestID(
			middleware.AddLogger(logger.Log, mux)),
	}
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		logger.Log.Info("graceful shuttting down Graphql Server")
		_ = srv.Shutdown(ctx)
	}()
	logger.Log.Info("starting HTTP/GRAPHQL  gateway...")
	return srv.ListenAndServe()
}
