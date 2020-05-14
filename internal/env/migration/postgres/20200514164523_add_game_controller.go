package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions add column controller int not null default 0;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions drop column controller;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200514164523_add_game_controller", up, down, opts)
}
