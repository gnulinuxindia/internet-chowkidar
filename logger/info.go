package logger

import (
	"context"
)

type ContextKey string

const (
	Host          ContextKey = "host"
	Method        ContextKey = "method"
	URI           ContextKey = "uri"
	OperationName ContextKey = "operation_name"
	UserAgent     ContextKey = "user_agent"
)

func GetInfo(ctx context.Context, key ContextKey) *string {
	if val, ok := ctx.Value(key).(string); ok {
		return &val
	}

	return nil
}
