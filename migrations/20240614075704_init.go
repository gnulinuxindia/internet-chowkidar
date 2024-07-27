package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInit, downInit)
}

func upInit(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec("CREATE TABLE `counters` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `count` integer NOT NULL DEFAULT 0);")
	return err
}

func downInit(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `counters`");
	return err
}
