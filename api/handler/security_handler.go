package handler

import (
	"context"
	"errors"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/internal/config"
)

type SecurityHandler interface {
	HandleApiKeyAuth(ctx context.Context, operationName string, t genapi.ApiKeyAuth) (context.Context, error)
}

type securityHandlerImpl struct {
	conf *config.Config
}

func NewSecurityHandler(conf *config.Config) SecurityHandler {
	return &securityHandlerImpl{
		conf: conf,
	}
}

func (s *securityHandlerImpl) HandleApiKeyAuth(ctx context.Context, operationName string, t genapi.ApiKeyAuth) (context.Context, error) {
	if t.APIKey == s.conf.ApiKey {
		return ctx, nil
	}
	return nil, errors.New("invalid api key")
}
