package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			create table documents
			(
				id serial primary key,

				title text,
				text text not null,
				reason text not null,

				type  text not null,
				language varchar(8) not null,
				version text not null,

				created_by text,
				activated_by text,

				created_at timestamp(0),
				updated_at timestamp(0),
				activated_at timestamp(0)
			);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			drop table documents;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200525193918_documents", up, down, opts)
}
