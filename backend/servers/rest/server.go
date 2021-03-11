package rest

import (
	"alfred/config"
	"alfred/logger"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// RunServer runs HTTP/REST gateway
func RunRESTServer(ctx context.Context, config *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	if err := RegisterRESTModules(ctx, mux, config.Grpc.Port); err != nil {
		panic(err)
	}
	srv := &http.Server{
		Addr: config.Rest.Host + ":" + config.Rest.Port,
		// add handler with middleware
		Handler:/* middleware.AddRequestID(
		middleware.AddLogger(logger.Log, mux))*/mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}
		_, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
		logger.Log.Info("graceful shuttting down Rest Server")
	}()

	logger.Log.Info("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
