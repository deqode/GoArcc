package prometheusServer

/*
import (
	"alfred/config"
	"alfred.sh/common/logger"
	"context"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type PrometheusConfig struct {
	Registry      *prometheus.Registry
	ServerMetrics *grpc_prometheus.ServerMetrics
}

func InitPromthesiusServer() *PrometheusConfig {
	// Create a metrics registry.
	registry := prometheus.NewRegistry()
	// Create some standard server metrics.
	grpcServerMetrics := grpc_prometheus.NewServerMetrics()
	registry.MustRegister(grpcServerMetrics)
	//CreateMetrics()
	return &PrometheusConfig{
		Registry:      registry,
		ServerMetrics: grpcServerMetrics,
	}
}

//Note: Prometheus will only work when all service will be registered with grpc already
func PrometheusRunner(config *config.Config, grpcServer *grpc.Server, prometheusConfig *PrometheusConfig) error {
	go func() {
		_ = RunPrometheusServer(context.Background(), config, grpcServer, prometheusConfig)
	}()
	return nil
}

// Started the prometheus server
func RunPrometheusServer(ctxhelper context.Context, config *config.Config, grpcServer *grpc.Server, prometheusConfig *PrometheusConfig) error {
	ctxhelper, cancel := context.WithCancel(ctxhelper)
	defer cancel()

	srv := &http.Server{
		Addr:    config.Promthesius.Host + ":" + config.Promthesius.Port,
		Handler: promhttp.HandlerFor(prometheusConfig.Registry, promhttp.HandlerOpts{}),
	}

	reflection.Register(grpcServer)
	//initializing metrics
	prometheusConfig.ServerMetrics.InitializeMetrics(grpcServer)
	grpc_prometheus.Register(grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()
	http.Handle("/metrics", promhttp.Handler())
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}
		_, cancel := context.WithTimeout(ctxhelper, 30*time.Second)
		defer cancel()
		logger.Log.Info("graceful shutting down promthesius Server")
		_ = srv.Shutdown(ctxhelper)
	}()

	logger.Log.Info("starting prometheus server...")
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
*/
