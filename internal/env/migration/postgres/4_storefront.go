package postgres

import (
	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
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
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			drop table storefront_activations;
			drop table storefront_versions;
			drop table storefronts;
		`)
		return err
	})
}
