package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upSuggestionlinkedsite, downSuggestionlinkedsite)
}

func upSuggestionlinkedsite(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
		ALTER TABLE site_suggestions ADD linked_site INT REFERENCES sites(id);
	`)
	return err
}

func downSuggestionlinkedsite(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
		ALTER TABLE site_suggestions DROP COLUMN linked_site;
	`)
	return err
}
