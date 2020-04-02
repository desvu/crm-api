package postgres

import (
	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		_, err := db.Exec(`
			create table games (
    			id             serial primary key,
    			title          text not null,
    			summary        text not null,
    			description    text not null,
    			license        text not null,
    			ranking        text not null,
    			type           int not null,
    			platforms      int[] not null,
    			release_date   timestamp(0) not null
			);
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`drop table games;`)
		return err
	})
}
