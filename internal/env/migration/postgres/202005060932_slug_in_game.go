package postgres

import (
	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		_, err := db.Exec(`
			alter table games add column slug text not null;
			alter table game_revisions drop column slug;

			create unique index idx_unq_games_slug on games(slug) where trim(slug) <> '';
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			alter table game_revisions add column slug text not null;
			alter table games drop column slug;
			drop index idx_unq_games_slug;
		`)
		return err
	})
}
