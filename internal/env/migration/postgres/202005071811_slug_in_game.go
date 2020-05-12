package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table games add column slug text not null;
			alter table game_revisions drop column slug;

			create unique index idx_unq_games_slug on games(slug) where trim(slug) <> '';
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions add column slug text not null;
			alter table games drop column slug;
			drop index idx_unq_games_slug;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("202005071811_slug_in_game", up, down, opts)
}
