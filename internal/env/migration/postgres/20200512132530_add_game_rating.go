package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			create table game_revision_ratings
			(
			  id serial primary key, 
			  game_revision_id int not null,
			  agency int not null,
			  rating int not null, 
			  display_online_notice bool, 
			  show_age_restrict bool,
			  constraint fk_game_revision_genres_games foreign key (game_revision_id) references game_revisions (id) on delete cascade
			);

			create unique index idx_unq_game_revision_id_organization on game_revision_ratings(game_revision_id, agency);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			drop table game_revision_ratings;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200512132530_add_age_restrictions", up, down, opts)
}
