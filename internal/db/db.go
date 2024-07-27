package db

import (
	"database/sql"
	"fmt"

	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/internal/config"
	"github.com/pressly/goose/v3"
	_ "github.com/go-sql-driver/mysql"

	// required by ent if using sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var (
	db    *ent.Client
	rawDb *sql.DB
	err   error
)

func ProvideDB(conf *config.Config) (*ent.Client, error) {
	if db != nil {
		return db, nil
	}

	switch conf.DatabaseDriver {
	case "sqlite3":
		db, err = newSqliteDB(conf)
	default:
		err = fmt.Errorf("unknown db type: %s", conf.DatabaseDriver)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func ProvideRawDB(conf *config.Config) (*sql.DB, error) {
	if rawDb == nil {
		var err error
		rawDb, err = goose.OpenDBWithDriver(conf.DatabaseDriver, conf.DatabaseURL)
		if err != nil {
			return nil, err
		}
	}

	return rawDb, nil
}

func MigrateUp(db *sql.DB) error {
	return goose.Up(db, "migrations", goose.WithNoColor(true))
}
