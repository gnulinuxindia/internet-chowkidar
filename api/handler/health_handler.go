package handler

import "context"

type HealthHandler interface {
	HealthCheck(ctx context.Context) (string, error)
}

type healthHandlerImpl struct {
}

func (h *healthHandlerImpl) HealthCheck(ctx context.Context) (string, error) {
	return "ok", nil
}
