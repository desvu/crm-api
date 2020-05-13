package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			create table game_revision_reviews
			(
			  id serial primary key, 
			  game_revision_id int not null,
			  press_name text,
			  link text, 
			  score varchar(8), 
			  quote text,
			  constraint fk_game_revision_genres_games foreign key (game_revision_id) references game_revisions (id) on delete cascade
			);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			drop table game_revision_reviews;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200512215019_add_game_review", up, down, opts)
}
