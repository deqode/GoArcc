package jaeger

import (
	"alfred/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
	"io"
)

// Jaeger
type Jaeger struct {
	Host        string
	ServiceName string
	LogSpans    bool
}

var (
	Tracer opentracing.Tracer
)

//todo close the connection
func InitJaeger() (io.Closer, opentracing.Tracer) {
	jaegerCfgInstance := jaegerconfig.Configuration{
		ServiceName: "Demo check",
		Sampler: &jaegerconfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: "localhost:6831",
		},
	}

	tracer, closer, err := jaegerCfgInstance.NewTracer(
		jaegerconfig.Logger(jaegerlog.StdLogger),
		jaegerconfig.Metrics(metrics.NullFactory),
	)
	if err != nil {
		logger.Log.Fatal("cannot create tracer", zap.Error(err))
	}
	logger.Log.Info("Jaeger Connected successfully")

	opentracing.SetGlobalTracer(tracer)
	return closer, tracer
}
