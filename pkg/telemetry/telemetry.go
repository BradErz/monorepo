package telemetry

import (
	"fmt"

	"github.com/go-logr/logr"
	"go.opentelemetry.io/contrib/propagators/b3"

	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"

	"go.opentelemetry.io/otel/exporters/jaeger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// Init configures an OpenTelemetry exporter and trace provider
func Init(lgr logr.Logger, opts ...Option) (trace.Tracer, error) {
	lgr = lgr.WithName("telemetry")
	conf, err := getConfig(opts...)
	if err != nil {
		return nil, fmt.Errorf("telemetry: failed to load config: %w", err)
	}

	if !conf.Enabled {
		lgr.Info("telemetry is disabled")
		return nil, nil
	}
	lgr.Info("telemetry is enabled", "endpoint", conf.JaegerCollectorEndpoint)

	exp, err := jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.JaegerCollectorEndpoint)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create jaeger exporter: %w", err)
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithSyncer(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(conf.ServiceName),
			attribute.String("environment", conf.Environment),
		)),
	)
	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
		b3.New(),
	))
	return tp.Tracer(""), nil
}
