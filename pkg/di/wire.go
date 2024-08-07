//go:generate wire
//go:build wireinject

package di

import (
	"database/sql"

	"github.com/gnulinuxindia/internet-chowkidar/api"
	"github.com/gnulinuxindia/internet-chowkidar/api/handler"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/internal/config"
	"github.com/gnulinuxindia/internet-chowkidar/internal/db"
	"github.com/gnulinuxindia/internet-chowkidar/internal/tracing"
	domainProvider "github.com/gnulinuxindia/internet-chowkidar/pkg/domain/provider"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/mock/gomock"

	"github.com/google/wire"
)

var dbSet = wire.NewSet(
	db.ProvideDB,
	db.ProvideRawDB,
	config.ProvideConfig,
)

var mockDbSet = wire.NewSet(
	db.ProvideDB,
	db.ProvideRawDB,
	config.ProvideConfig,
)

var concreteSet = wire.NewSet(
	api.HandlerSet,
	domainProvider.RepositorySet,
	domainProvider.ServiceSet,
	repository.NewTxHandler,
)

var mockSet = wire.NewSet(
	api.MockHandlerSet,
	domainProvider.MockRepositorySet,
	domainProvider.MockServiceSet,
	repository.NewTxHandler,
)

var tracingSet = wire.NewSet(
	config.ProvideConfig,
	tracing.ProvideTracerProvider,
)

// handlers
func InjectHandlers() (*api.Handlers, error) {
	wire.Build(concreteSet, dbSet)
	return &api.Handlers{}, nil
}

func InjectSecurityHandler() (handler.SecurityHandler, error) {
	panic(wire.Build(config.ProvideConfig, handler.NewSecurityHandler))
}

func InjectMockHandlers(ctrl *gomock.Controller) (*api.Handlers, error) {
	wire.Build(mockSet, mockDbSet)
	return &api.Handlers{}, nil
}

// services
func InjectServices() (*domainProvider.Services, error) {
	wire.Build(concreteSet, dbSet)
	return &domainProvider.Services{}, nil
}

func InjectMockServices(ctrl *gomock.Controller) (*domainProvider.Services, error) {
	wire.Build(mockSet, mockDbSet)
	return &domainProvider.Services{}, nil
}

// repositories
func InjectRepository() (*domainProvider.Repositories, error) {
	wire.Build(concreteSet, dbSet)
	return &domainProvider.Repositories{}, nil
}

func InjectMockRepository(ctrl *gomock.Controller) (*domainProvider.Repositories, error) {
	wire.Build(mockSet, mockDbSet)
	return &domainProvider.Repositories{}, nil
}

// misc
func InjectDb() (*ent.Client, error) {
	wire.Build(dbSet)
	return &ent.Client{}, nil
}

func InjectMockDb() (*ent.Client, error) {
	panic(wire.Build(mockDbSet))
}

func InjectTracerProvider() (*trace.TracerProvider, error) {
	wire.Build(tracingSet)
	return &trace.TracerProvider{}, nil
}

func InjectRawDb() (*sql.DB, error) {
	wire.Build(dbSet)
	return &sql.DB{}, nil
}

func InjectConfig() (*config.Config, error) {
	wire.Build(config.ProvideConfig)
	return &config.Config{}, nil
}

func InjectTxHandler() (repository.TxHandler, error) {
	wire.Build(dbSet, repository.NewTxHandler)
	return repository.TxHandler{}, nil
}
