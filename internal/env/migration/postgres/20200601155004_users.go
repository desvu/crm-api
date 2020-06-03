package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			create table users
			(
				id serial primary key,
				external_id text
			);

			create table user_documents
			(
				id				serial primary key,
				user_id			int not null,
				document_id		int not null,
				created_at		timestamp(0) not null,
				constraint fk_user_document_users foreign key (user_id) references users (id) on delete cascade,
				constraint fk_user_document_documents foreign key (document_id) references documents (id) on delete cascade
			);

			create unique index idx_unq_user_documents_user_id_document_id on user_documents(user_id, document_id);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			drop table users;
			drop table user_documents;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200601155004_users", up, down, opts)
}
