package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			create table game_revision_localizations 
			(
			  id serial primary key, 
			  game_revision_id int not null,
			  language varchar(3),
			  interface bool, 
			  audio bool, 
			  subtitles bool,
			  constraint fk_game_revision_genres_games foreign key (game_revision_id) references game_revisions (id) on delete cascade
			);

			create unique index idx_unq_game_revision_id_language on game_revision_localizations(game_revision_id, language);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			drop table game_revision_localizations;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("202005081419_game_localizations", up, down, opts)
}
