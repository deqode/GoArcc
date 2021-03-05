package promthesiusServer

import (
	"alfred/config"
	"alfred/logger"
	"context"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"time"
)


type PromthesiusConfig struct {
	Registry *prometheus.Registry
	ServerMetrics *grpc_prometheus.ServerMetrics
}


func InitPromthesiusServer() *PromthesiusConfig {
	// Create a metrics registry.
	registry := prometheus.NewRegistry()
	// Create some standard server metrics.
	grpcServerMetrics := grpc_prometheus.NewServerMetrics()
	registry.MustRegister(grpcServerMetrics)
	grpc_prometheus.EnableHandlingTimeHistogram()
    return &PromthesiusConfig{
    	Registry: registry,
    	ServerMetrics: grpcServerMetrics,
	}
}

// Started the promthesius server
func RunPromthesiusServer(ctx context.Context, config *config.Config , grpcServer *grpc.Server) error {
	// Create a metrics registry.
	reg := prometheus.NewRegistry()
	// Create some standard server metrics.
	grpcMetrics := grpc_prometheus.NewServerMetrics()
	reg.MustRegister(grpcMetrics)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	srv := &http.Server{
		Addr:    config.ServerHost + ":" + config.PromthesiusPort,
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
	}
	//initializing metrics
	grpcMetrics.InitializeMetrics(grpcServer)
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	logger.Log.Info("starting promthesius server...")
	return srv.ListenAndServe()

}
