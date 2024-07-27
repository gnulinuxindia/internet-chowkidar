package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddTags, downAddTags)
}

func upAddTags(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
		create table categories
		(
			id         serial                              not null
				constraint tags_pk
					primary key,
			name       varchar(255)                        not null,
			created_at timestamp default current_timestamp not null,
			updated_at timestamp   default current_timestamp not null
		);

		create table sites_categories
		(
			sites_id      integer not null
				constraint sites_categories_sites_id_fk
					references sites,
			categories_id integer not null
				constraint sites_categories_categories_id_fk
					references categories,
			constraint sites_categories_pk
				primary key (sites_id, categories_id)
		);
	`)
	return err
}

func downAddTags(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
		drop table if exists sites_categories;
		drop table if exists categories;
	`)
	return err
}
