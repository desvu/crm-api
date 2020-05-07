package postgres

import (
	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		_, err := db.Exec(`
			alter table features add column icon text not null;
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			alter table features drop column icon;
		`)
		return err
	})
}
