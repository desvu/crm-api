package postgres

import (
	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		_, err := db.Exec(`
			alter table game_revisions add column system_requirements jsonb not null default '[]'::jsonb;
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			alter table game_revisions drop column system_requirements;
		`)
		return err
	})
}
