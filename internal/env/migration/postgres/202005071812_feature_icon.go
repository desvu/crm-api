package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table features add column icon text not null;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			alter table features drop column icon;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("202005071812_feature_icon", up, down, opts)
}
