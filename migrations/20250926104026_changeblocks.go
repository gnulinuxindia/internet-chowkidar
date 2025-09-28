package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upChangeblocks, downChangeblocks)
}

func upChangeblocks(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
		ALTER TABLE blocks ADD blocked boolean;
		ALTER TABLE blocks DROP COLUMN block_reports;
		ALTER TABLE blocks DROP COLUMN unblock_reports;
	`)
	return err
}

func downChangeblocks(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
		ALTER TABLE blocks ADD block_reports INT NOT NULL DEFAULT 0,;
		ALTER TABLE blocks ADD unblock_reports INT NOT NULL DEFAULT 0,;
		ALTER TABLE blocks DROP COLUMN blocked;
	`)
	return err
}
