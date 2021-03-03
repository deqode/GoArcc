package graphql

import (
	"alfred/logger"
	"alfred/protocol/graphql/middleware"
	"context"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// RunServer runs HTTP/REST gateway
func RunGraphqlServer(ctx context.Context, graphqlPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	if err := RegisterGraphqlModules(ctx, mux); err != nil {
		panic(err)
	}
	http.Handle("/graphql", mux)
	srv := &http.Server{
		Addr: ":" + graphqlPort,
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

		_ = srv.Shutdown(ctx)
	}()
	logger.Log.Info("starting HTTP/GRAPHQL  gateway...")
	return srv.ListenAndServe()
}
