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
	Registry      *prometheus.Registry
	ServerMetrics *grpc_prometheus.ServerMetrics
}

func InitPromthesiusServer() *PromthesiusConfig {
	// Create a metrics registry.
	registry := prometheus.NewRegistry()
	// Create some standard server metrics.
	grpcServerMetrics := grpc_prometheus.NewServerMetrics()
	registry.MustRegister(grpcServerMetrics)
	//CreateMetrics()
	return &PromthesiusConfig{
		Registry:      registry,
		ServerMetrics: grpcServerMetrics,
	}
}

//Note: Promthesius will only work when all service will be registered with grpc already
func PromthesiusRunner(config *config.Config, grpcServer *grpc.Server, promthesusConfig *PromthesiusConfig) error {
	go func() {
		_ = RunPromthesiusServer(context.Background(), config, grpcServer, promthesusConfig)
	}()
	return nil
}

// Started the promthesius server
func RunPromthesiusServer(ctx context.Context, config *config.Config, grpcServer *grpc.Server, promthesusConfig *PromthesiusConfig) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	srv := &http.Server{
		Addr:    config.ServerHost + ":" + config.PromthesiusPort,
		Handler: promhttp.HandlerFor(promthesusConfig.Registry, promhttp.HandlerOpts{}),
	}
	//initializing metrics
	promthesusConfig.ServerMetrics.InitializeMetrics(grpcServer)
	grpc_prometheus.Register(grpcServer)
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}
		_, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		logger.Log.Info("graceful shutting down promthesius Server")
		_ = srv.Shutdown(ctx)
	}()

	logger.Log.Info("starting promthesius server...")
	return srv.ListenAndServe()
}

type PrometheusMetrics struct {
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
	Times     *prometheus.HistogramVec
}

func CreateMetrics() (*PrometheusMetrics, error) {
	var metr PrometheusMetrics
	metr.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "_hits_total",
	})
	if err := prometheus.Register(metr.HitsTotal); err != nil {
		return nil, err
	}
	metr.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "_hits",
		},
		[]string{"status", "method", "path"},
	)
	if err := prometheus.Register(metr.Hits); err != nil {
		return nil, err
	}
	metr.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "_times",
		},
		[]string{"status", "method", "path"},
	)
	if err := prometheus.Register(metr.Times); err != nil {
		return nil, err
	}
	if err := prometheus.Register(prometheus.NewBuildInfoCollector()); err != nil {
		return nil, err
	}
	return &metr, nil
}
