package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions add column system_requirements jsonb not null default '[]'::jsonb;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions drop column system_requirements;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("202005071814_system_requirements", up, down, opts)
}
