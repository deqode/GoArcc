package jaeger

import (
	"github.com/deqode/GoArcc/config"
	"github.com/deqode/GoArcc/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
	"io"
	"strconv"
)

// Jaeger struct for conversion
type Jaeger struct {
	Host        string
	ServiceName string
	LogSpans    bool
}

var (
	Tracer opentracing.Tracer
)

// InitJaeger todo close the connection
// InitJaeger is responsible for initialising the jaeger tracing instance.
func InitJaeger(config *config.Config) (io.Closer, opentracing.Tracer) {
	//conversion string val to bool
	logSpan, _ := strconv.ParseBool(config.Jaeger.LogSpans)
	//Jaeger configuration setup
	jaegerCfgInstance := jaegerconfig.Configuration{
		//Service name is the name of project for which tracing will be present
		ServiceName: config.Jaeger.ServiceName,
		Sampler: &jaegerconfig.SamplerConfig{
			// SamplerTypeConst is the type of sampler that always makes the same decision.
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		//Reporter configuration
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans: logSpan,
			//Local Agent Host port
			LocalAgentHostPort: config.Jaeger.Host + ":" + config.Jaeger.Port,
		},
	}
	//New Tracer instance
	tracer, closer, err := jaegerCfgInstance.NewTracer(
		jaegerconfig.Logger(jaegerlog.StdLogger),
		jaegerconfig.Metrics(metrics.NullFactory),
	)
	if err != nil {
		logger.Log.Fatal("cannot create tracer", zap.Error(err))
		panic(err)
	}

	logger.Log.Info("Jaeger tracer initiated")
	//global tracer setup , opentracing
	opentracing.SetGlobalTracer(tracer)
	return closer, tracer
}
