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
	_, err := tx.Exec(`
-- Sites table
CREATE TABLE sites (
    id SERIAL PRIMARY KEY,
    domain VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ISPs table
CREATE TABLE isps(
    id SERIAL PRIMARY KEY,
  latitude FLOAT NOT NULL,
  longitude FLOAT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- blocks table
CREATE TABLE blocks (
    id SERIAL PRIMARY KEY,
    site_id INT NOT NULL REFERENCES sites(id),
    isp_id INT NOT NULL REFERENCES isps(id),
    last_reported_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  block_reports INT NOT NULL DEFAULT 0,
  unblock_reports INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Abuse_reports table
CREATE TABLE abuse_reports (
    id SERIAL PRIMARY KEY,
    site_id INT NOT NULL REFERENCES sites(id),
    reason TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    resolved_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Site_suggestions table
CREATE TABLE site_suggestions (
    id SERIAL PRIMARY KEY,
    site_id INT NOT NULL REFERENCES sites(id),
    reason TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    resolved_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
	`)
	return err
}

func downInit(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `blocks`, `abuse_reports`, `site_suggestions`, `sites`, `isps`;");
	return err
}
