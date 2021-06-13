package telemetry

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/semconv"

	"github.com/sirupsen/logrus"

	"go.opentelemetry.io/otel/exporters/trace/jaeger"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// Init configures an OpenTelemetry exporter and trace provider
func Init(le *logrus.Entry, opts ...Option) error {
	conf, err := getConfig(opts...)
	if err != nil {
		return fmt.Errorf("telemetry: failed to load config: %w", err)
	}

	if !conf.Enabled {
		le.Infof("telemetry is disabled")
		return nil
	}
	le.Infof("telemetry is enabled! Sending traces to: %s", conf.JaegerCollectorEndpoint)

	exp, err := jaeger.NewRawExporter(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.JaegerCollectorEndpoint)),
	)
	if err != nil {
		return fmt.Errorf("failed to create jaeger exporter: %w", err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSyncer(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.ServiceNameKey.String(conf.ServiceName),
			attribute.String("environment", conf.Environment),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return nil
}
