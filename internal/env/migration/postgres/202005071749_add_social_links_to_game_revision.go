package postgres

import (
	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		_, err := db.Exec(`
			alter table game_revisions add column social_links jsonb not null default '[]'::jsonb;
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			alter table game_revisions drop column social_links;
		`)
		return err
	})
}
