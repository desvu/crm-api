package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions add column social_links jsonb not null default '[]'::jsonb;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions drop column social_links;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200507181459_system_requirements", up, down, opts)
}
