package metrics

import (
	"alfred/servers/grpc/grpcErrors"
	"context"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

// Metrics Manager
type Manager struct {
	metrics Metrics
}

// InterceptorManager constructor
func NewMetricsManager(metrics Metrics) *Manager {
	return &Manager{metrics: metrics}
}

func (im *Manager) Metrics(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	var status = http.StatusOK
	if err != nil {
		status = grpcErrors.MapGRPCErrCodeToHttpStatus(grpcErrors.ParseGRPCErrStatusCode(err))
	}
	im.metrics.ObserveResponseTime(status, info.FullMethod, info.FullMethod, time.Since(start).Seconds())
	im.metrics.IncHits(status, info.FullMethod, info.FullMethod)

	return resp, err
}
