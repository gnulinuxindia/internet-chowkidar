package db

import (
	"log/slog"

	"github.com/gnulinuxindia/internet-chowkidar/ent"

	"github.com/gnulinuxindia/internet-chowkidar/internal/config"

	// required by ent if using mysql
	"github.com/go-errors/errors"
	_ "github.com/go-sql-driver/mysql"

	// required by ent if using sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func newSqliteDB(conf *config.Config) (*ent.Client, error) {
	client, err := ent.Open(conf.DatabaseDriver, conf.DatabaseURL)
	if err != nil {
		slog.Error("error connecting to sqlite database", "err", err)
		return nil, errors.Wrap(err, 0)
	}

	return client, nil
}
