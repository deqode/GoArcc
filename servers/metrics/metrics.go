package metrics

import (
	"github.com/deqode/GoArcc/config"
	"github.com/deqode/GoArcc/logger"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"strconv"
)

//Metrics details of hits , response time
type Metrics interface {
	IncHits(status int, method, path string)
	ObserveResponseTime(status int, method, path string, observeTime float64)
}

type PrometheusMetrics struct {
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
	Times     *prometheus.HistogramVec
}

// CreateMetrics : Create Metrics will create the metrics in prometheus server.
func CreateMetrics(config *config.Config) Metrics {
	var metrics PrometheusMetrics
	metrics.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: config.Metrics.ServiceName + "_hits_total",
	})
	if err := prometheus.Register(metrics.HitsTotal); err != nil {
		logger.Log.Fatal("error in registering prometheus hits", zap.Error(err))
		return nil
	}
	metrics.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: config.Metrics.ServiceName + "_hits",
		},
		[]string{"status", "method", "path"},
	)
	if err := prometheus.Register(metrics.Hits); err != nil {
		logger.Log.Fatal("error in registering prometheus hits", zap.Error(err))
		return nil
	}
	metrics.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: config.Metrics.ServiceName + "_times",
		},
		[]string{"status", "method", "path"},
	)
	if err := prometheus.Register(metrics.Times); err != nil {
		logger.Log.Fatal("error in registering prometheus", zap.Error(err))
		return nil
	}
	if err := prometheus.Register(prometheus.NewBuildInfoCollector()); err != nil {
		logger.Log.Fatal("error in registering prometheus", zap.Error(err))
		return nil
	}
	go func() {
		router := echo.New()
		router.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		if err := router.Start(config.Metrics.URL); err != nil {
			logger.Log.Fatal("unable to create metrics", zap.Error(err))
		}
	}()
	logger.Log.Info("Metrics available URL: ", zap.String("url", config.Metrics.URL),
		zap.String("service name:", config.Metrics.ServiceName))
	return &metrics
}

//IncHits : increment the total hits
func (metrics *PrometheusMetrics) IncHits(status int, method, path string) {
	metrics.HitsTotal.Inc()
	metrics.Hits.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}

//ObserveResponseTime :
func (metrics *PrometheusMetrics) ObserveResponseTime(status int, method, path string, observeTime float64) {
	metrics.Times.WithLabelValues(strconv.Itoa(status), method, path).Observe(observeTime)
}
