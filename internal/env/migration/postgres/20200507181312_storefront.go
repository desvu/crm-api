package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`

			CREATE TABLE storefronts (
				id         bigserial, 
				name       text       not null, 
				created_at timestamp DEFAULT now(),
				PRIMARY KEY (id)
			);

			CREATE TABLE storefront_versions (
				storefront_id bigint     not null, 
				id            bigint     not null, 
				blocks        jsonb      not null,
				created_at    timestamp DEFAULT now(), 
				PRIMARY KEY (storefront_id, id)
			);

			CREATE TABLE storefront_activations (
				timestamp     timestamp DEFAULT now(),
				version_id    bigint     not null, 
				storefront_id bigint     not null, 
				PRIMARY KEY (timestamp)
			);

		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			drop table storefront_activations;
			drop table storefront_versions;
			drop table storefronts;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200507181312_storefront", up, down, opts)
}
