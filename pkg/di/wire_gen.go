// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"database/sql"
	"github.com/gnulinuxindia/internet-chowkidar/api"
	"github.com/gnulinuxindia/internet-chowkidar/api/handler"
	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/internal/config"
	"github.com/gnulinuxindia/internet-chowkidar/internal/db"
	"github.com/gnulinuxindia/internet-chowkidar/internal/tracing"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/provider"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/service"
	"github.com/golang/mock/gomock"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// handlers
func InjectHandlers() (*api.Handlers, error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, err
	}
	client, err := db.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	counterRepository := repository.ProvideCounterRepository(client)
	emailService := service.ProvideEmailService()
	counterService := service.ProvideCounterService(counterRepository, emailService)
	counterHandler := handler.ProvideCounterHandler(counterService)
	handlers := &api.Handlers{
		CounterHandler: counterHandler,
	}
	return handlers, nil
}

func InjectMockHandlers(ctrl *gomock.Controller) (*api.Handlers, error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, err
	}
	client, err := db.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	counterRepository := repository.ProvideCounterRepository(client)
	emailService := di.ProvideMockEmailService(ctrl)
	counterService := service.ProvideCounterService(counterRepository, emailService)
	counterHandler := handler.ProvideCounterHandler(counterService)
	handlers := &api.Handlers{
		CounterHandler: counterHandler,
	}
	return handlers, nil
}

// services
func InjectServices() (*di.Services, error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, err
	}
	client, err := db.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	counterRepository := repository.ProvideCounterRepository(client)
	emailService := service.ProvideEmailService()
	counterService := service.ProvideCounterService(counterRepository, emailService)
	services := &di.Services{
		CounterService: counterService,
		EmailService:   emailService,
	}
	return services, nil
}

func InjectMockServices(ctrl *gomock.Controller) (*di.Services, error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, err
	}
	client, err := db.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	counterRepository := repository.ProvideCounterRepository(client)
	emailService := di.ProvideMockEmailService(ctrl)
	counterService := service.ProvideCounterService(counterRepository, emailService)
	services := &di.Services{
		CounterService: counterService,
		EmailService:   emailService,
	}
	return services, nil
}

// repositories
func InjectRepository() (*di.Repositories, error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, err
	}
	client, err := db.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	counterRepository := repository.ProvideCounterRepository(client)
	repositories := &di.Repositories{
		CounterRepository: counterRepository,
	}
	return repositories, nil
}

func InjectMockRepository(ctrl *gomock.Controller) (*di.Repositories, error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, err
	}
	client, err := db.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	counterRepository := repository.ProvideCounterRepository(client)
	repositories := &di.Repositories{
		CounterRepository: counterRepository,
	}
	return repositories, nil
}

// misc
func InjectDb() (*ent.Client, error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, err
	}
	client, err := db.ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InjectTracerProvider() (*trace.TracerProvider, error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, err
	}
	tracerProvider, err := tracing.ProvideTracerProvider(configConfig)
	if err != nil {
		return nil, err
	}
	return tracerProvider, nil
}

func InjectRawDb() (*sql.DB, error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.ProvideRawDB(configConfig)
	if err != nil {
		return nil, err
	}
	return sqlDB, nil
}

// wire.go:

var dbSet = wire.NewSet(db.ProvideDB, db.ProvideRawDB, config.ProvideConfig)

var concreteSet = wire.NewSet(api.HandlerSet, di.RepositorySet, di.ServiceSet)

var mockSet = wire.NewSet(api.MockHandlerSet, di.MockRepositorySet, di.MockServiceSet)

var tracingSet = wire.NewSet(config.ProvideConfig, tracing.ProvideTracerProvider)
