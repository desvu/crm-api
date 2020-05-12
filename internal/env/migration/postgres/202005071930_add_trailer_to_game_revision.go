package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions add column trailer text not null default '';
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revisions drop column trailer;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("202005071930_add_trailer_to_game_revision", up, down, opts)
}
