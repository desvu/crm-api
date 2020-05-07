package postgres

import (
	"github.com/go-pg/pg/v9"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

const Directory = "internal/env/migration/postgres"

func Migrate(db *pg.DB) error {
	return migrations.Run(db, Directory, []string{"app", "migrate"})
}
