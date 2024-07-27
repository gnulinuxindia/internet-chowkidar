package db

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/go-errors/errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func ProvideMockDB() (*sql.DB, error) {
	ctx := context.Background()
	pgContainer, err := postgres.RunContainer(
		ctx,
		testcontainers.WithImage("postgres:16-alpine"),
		postgres.WithDatabase("kotozna-test"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(10*time.Second),
		),
	)
	if err != nil {
		slog.Error("failed to start postgres container", "error", err)
		return nil, errors.Wrap(err, 0)
	}

	rawDb, err := goose.OpenDBWithDriver("postgres", pgContainer.MustConnectionString(ctx))
	if err != nil {
		slog.Error("failed to open db", "error", err)
		return nil, errors.Wrap(err, 0)
	}

	err = MigrateUp(rawDb)
	if err != nil {
		slog.Error("failed to migrate up", "error", err)
		return nil, errors.Wrap(err, 0)
	}

	return rawDb, nil
}
