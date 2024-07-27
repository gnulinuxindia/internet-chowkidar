package tlog

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gnulinuxindia/internet-chowkidar/utils"

	"github.com/go-errors/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func addToSpan(ctx context.Context, msg string, severity string, attrs *map[string]any) {
	var otelAttrs []attribute.KeyValue
	otelAttrs = append(otelAttrs, attribute.String("message", msg))

	if attrs != nil {
		for k, v := range *attrs {
			switch value := v.(type) {
			case string:
				otelAttrs = append(otelAttrs, attribute.String(k, value))
			case bool:
				otelAttrs = append(otelAttrs, attribute.Bool(k, value))
			case int:
				otelAttrs = append(otelAttrs, attribute.Int(k, value))
			case int64:
				otelAttrs = append(otelAttrs, attribute.Int64(k, value))
			case float64:
				otelAttrs = append(otelAttrs, attribute.Float64(k, value))
			default:
				slog.Debug(fmt.Sprintf("unhandled attribute type for key '%s', using string representation", k))
				otelAttrs = append(otelAttrs, attribute.String(k, fmt.Sprint(value)))
			}
		}
	}

	span := trace.SpanFromContext(ctx)
	if span != nil && span.SpanContext().IsValid() {
		span.AddEvent(severity, trace.WithAttributes(otelAttrs...))
	}
}

// slog takes attributes as a list like [a, b, c, d] where
// a and c are keys and b and d are the values and so on
func makePairArray(attrs *map[string]any) []any {
	if attrs == nil {
		return []any{}
	}

	kvPairs := make([]any, 0, len(*attrs)*2)
	for k, v := range *attrs {
		kvPairs = append(kvPairs, k, v)
	}
	return kvPairs
}

func Debug(ctx context.Context, msg string, attrs *map[string]any) {
	addToSpan(ctx, msg, "debug", attrs)
	slog.Debug(msg, makePairArray(attrs)...)
}

func Info(ctx context.Context, msg string, attrs *map[string]any) {
	addToSpan(ctx, msg, "info", attrs)
	slog.Info(msg, makePairArray(attrs)...)
}

func Warn(ctx context.Context, msg string, attrs *map[string]any) {
	addToSpan(ctx, msg, "warn", attrs)
	slog.Warn(msg, makePairArray(attrs)...)
}

func Error(ctx context.Context, msg string, attrs *map[string]any) {
	addToSpan(ctx, msg, "error", attrs)
	slog.Error(msg, makePairArray(attrs)...)
}

func Exception(ctx context.Context, err error) {
	attrs := make(map[string]any)
	stack := utils.GetStack(errors.Wrap(err, 0), false)
	coloredStack := utils.GetStack(errors.Wrap(err, 0), true)
	addToSpan(ctx, "exception", "error", &map[string]any{
		"exception.message": err.Error(),
		"stack":             stack,
	})
	slog.Error(coloredStack, makePairArray(&attrs)...)
}
