package utils

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/ogen-go/ogen/otelogen"
	"go.opentelemetry.io/otel/trace"
)

func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "..")
}

func Filter[T any](data []T, f func(T) bool) []T {
	filtered := make([]T, 0, len(data))

	for _, e := range data {
		if f(e) {
			filtered = append(filtered, e)
		}
	}

	return filtered
}

func MapOver[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}

func Span(ctx context.Context, name string, f func(ctx context.Context) error) error {
	internalCtx, span := trace.SpanFromContext(ctx).TracerProvider().Tracer(otelogen.Name).Start(ctx, name)
	defer span.End()
	span.SetName(name)

	if err := f(internalCtx); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

func Contains[T comparable](data []T, elem T) bool {
	for _, e := range data {
		if e == elem {
			return true
		}
	}

	return false
}
