package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revision_reviews alter column score type smallint using score::smallint;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table game_revision_reviews alter column score varchar(8) using score::varchar;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200521153922_alter_reviews_score", up, down, opts)
}
