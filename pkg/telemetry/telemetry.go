package telemetry

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/metric"

	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"

	"go.opentelemetry.io/otel/exporters/jaeger"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// Init configures an OpenTelemetry exporter and trace provider
func Init(lgr logr.Logger, opts ...Option) (trace.Tracer, func(ctx context.Context), error) {
	lgr = lgr.WithName("telemetry")
	conf, err := getConfig(opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("telemetry: failed to load config: %w", err)
	}

	if !conf.Enabled {
		lgr.Info("telemetry is disabled")
		return nil, nil, nil
	}
	lgr.Info("telemetry is enabled", "endpoint", conf.JaegerCollectorEndpoint)

	exp, err := jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.JaegerCollectorEndpoint)),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create jaeger exporter: %w", err)
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithSyncer(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(conf.ServiceName),
			attribute.String("environment", conf.Environment),
			attribute.String("job", "app/"+conf.ServiceName),
		)),
	)
	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
		b3.New(),
	))
	if conf.MetricsEnabled {
		lgr.Info("serving prometheus metrics")
		prometheusMetrics()
	}

	// TODO: this doesnt work and causes containers to hang if they
	// couldnt connect to the endpoint for jaeger
	stop := func(ctx context.Context) {
		lgr.Info("shutting down telemetry")
		_ = tp.Shutdown(ctx)
		_ = exp.Shutdown(ctx)
	}

	return tp.Tracer(""), stop, nil
}

func prometheusMetrics() {
	exporter := otelprom.New()
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	global.SetMeterProvider(provider)
	go serveMetrics(exporter.Collector)
}

func serveMetrics(collector prometheus.Collector) {
	registry := prometheus.NewRegistry()
	if err := registry.Register(collector); err != nil {
		fmt.Printf("error registering collector: %v", err)
		return
	}

	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	server := http.Server{
		Addr:         ":2222",
		Handler:      metricsMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("error serving http: %v", err)
		return
	}
}
