package tracing

import (
	"context"

	"github.com/gnulinuxindia/internet-chowkidar/internal/config"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var tp *sdktrace.TracerProvider

// ProvideTracerProvider provides a tracer provider
func ProvideTracerProvider(conf *config.Config) (*sdktrace.TracerProvider, error) {
	if tp == nil {
		ctx := context.Background()
		// ─── Configure Tracer Provider ────────────────────────────────────────
		exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
		if err != nil {
			return nil, err
		}

		resource, err := resource.New(ctx,
			resource.WithAttributes(
				semconv.ServiceNameKey.String(conf.ServiceName),
				semconv.ServiceVersionKey.String(conf.ServiceVersion),
				attribute.String("environment", conf.Env),
			),
		)
		if err != nil {
			return nil, err
		}

		tp = sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resource),
		)
	}

	return tp, nil
}
