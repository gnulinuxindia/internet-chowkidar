package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upPingurl, downPingurl)
}

func upPingurl(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
		ALTER TABLE sites ADD ping_url varchar(255);
	`)
	return err
}

func downPingurl(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
		ALTER TABLE sites DROP COLUMN ping_url;
	`)
	return err
}
