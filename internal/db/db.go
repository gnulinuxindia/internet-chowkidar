package db

import (
	"database/sql"
	"log/slog"
	"strings"

	entsql "entgo.io/ent/dialect/sql"

	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/internal/config"
	"github.com/go-errors/errors"
	"github.com/pressly/goose/v3"
)

var (
	db    *ent.Client
	rawDb *sql.DB
)

func ProvideDB(raw *sql.DB, conf *config.Config) (*ent.Client, error) {
	if db == nil {
		drv := entsql.OpenDB(conf.DatabaseDriver, raw)
		db = ent.NewClient(ent.Driver(drv))
	}

	if strings.ToLower(conf.Env) == "debug" || strings.ToLower(conf.Env) == "CI" {
		return db.Debug(), nil
	} else {
		return db, nil
	}
}

func ProvideRawDB(conf *config.Config) (*sql.DB, error) {
	if rawDb == nil {
		slog.Debug("connecting to database", "driver", conf.DatabaseDriver)
		var err error
		rawDb, err = goose.OpenDBWithDriver(conf.DatabaseDriver, conf.DatabaseURL)
		if err != nil {
			return nil, errors.Wrap(err, 0)
		}
	}

	return rawDb, nil
}

func MigrateUp(db *sql.DB) error {
	slog.Debug("migrating up")
	return goose.Up(db, "migrations", goose.WithNoColor(true))
}
