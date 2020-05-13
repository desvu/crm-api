package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			create table game_media
			(
				id           serial primary key,
				game_id      uuid         not null,
				type         int          not null,
				file_path    text         not null,
				is_uploaded  bool         not null,
				created_at   timestamp(0) not null default now(),
				constraint fk_game_media_games foreign key (game_id) references games (id) on delete restrict
			);

			CREATE TABLE game_revision_media (
				id            serial primary key, 
				revision_id   int    not null, 
				game_media_id int    not null,
				constraint fk_game_revision_media_game_revisions foreign key (revision_id) references game_revisions (id) on delete cascade,
				constraint fk_game_revision_media_game_media foreign key (game_media_id) references game_media (id) on delete cascade
			);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			drop table game_media;
			drop table game_revision_media;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("202005071814_game_media", up, down, opts)
}
