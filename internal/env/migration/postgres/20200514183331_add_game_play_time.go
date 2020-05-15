package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions add column play_time int not null default 0;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions drop column play_time;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200514183331_add_game_play_time", up, down, opts)
}
