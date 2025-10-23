package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upSitesuggestions, downSitesuggestions)
}

func upSitesuggestions(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.

	// Old sites_suggestions table was never used, so delete that and then create a new table
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS site_suggestions;
		CREATE TYPE suggestion_status AS ENUM ('pending', 'accepted', 'rejected');
		CREATE TABLE site_suggestions 
		(
			id serial primary key,
			domain varchar(255) unique not null,
			ping_url varchar(255) unique,
			categories varchar(255),
			reason varchar(255),
			status suggestion_status default 'pending',
			resolve_reason varchar(255),
			resolved_at timestamp,
			created_at timestamp default current_timestamp not null,
			updated_at timestamp   default current_timestamp not null
		);
	`)
	return err
}

func downSitesuggestions(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
		drop table if exists site_suggestions;
	`)
	return err
}
