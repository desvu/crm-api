package postgres

import (
	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		_, err := db.Exec(`
			create table games
			(
				id           serial primary key,
				title        text         not null,
				summary      text         not null,
				description  text         not null,
				license      text         not null,
				ranking      text         not null,
				type         int          not null,
				platforms    int[]        not null,
				release_date timestamp(0) not null
			);
			
			create table tags
			(
				id   serial primary key,
				name text not null
			);
			
			create table game_tags
			(
				id      serial primary key,
				game_id int not null,
				tag_id  int not null,
				constraint fk_game_tags_games foreign key (game_id) references games (id) on delete cascade,
				constraint fk_game_tags_tags foreign key (tag_id) references tags (id) on delete cascade
			);
			
			create table developers
			(
				id   serial primary key,
				name text not null
			);
			
			create table game_developers
			(
				id           serial primary key,
				game_id      int not null,
				developer_id int not null,
				constraint fk_game_developers_games foreign key (game_id) references games (id) on delete cascade,
				constraint fk_game_developers_developers foreign key (developer_id) references developers (id) on delete cascade
			);
			
			create table publishers
			(
				id   serial primary key,
				name text not null
			);
			
			create table game_publishers
			(
				id           serial primary key,
				game_id      int not null,
				publisher_id int not null,
				constraint fk_game_publishers_games foreign key (game_id) references games (id) on delete cascade,
				constraint fk_game_publishers_publishers foreign key (publisher_id) references publishers (id) on delete cascade
			);
			
			create table features
			(
				id   serial primary key,
				name text not null
			);
			
			create table game_features
			(
				id         serial primary key,
				game_id    int not null,
				feature_id int not null,
				constraint fk_game_features_games foreign key (game_id) references games (id) on delete cascade,
				constraint fk_game_features_features foreign key (feature_id) references features (id) on delete cascade
			);
			
			create table genres
			(
				id   serial primary key,
				name text not null
			);
			
			create table game_genres
			(
				id       serial primary key,
				game_id  int not null,
				genre_id int not null,
				constraint fk_game_genres_games foreign key (game_id) references games (id) on delete cascade,
				constraint fk_game_genres_genres foreign key (genre_id) references genres (id) on delete cascade
			);

			create index idx_game_tags_game_id ON game_tags(game_id);
			create index idx_game_tags_tag_id ON game_tags(tag_id);

			create index idx_game_developers_game_id ON game_developers(game_id);
			create index idx_game_developers_developer_id ON game_developers(developer_id);

			create index idx_game_publishers_game_id ON game_publishers(game_id);
			create index idx_game_publishers_publisher_id ON game_publishers(publisher_id);

			create index idx_game_features_game_id ON game_features(game_id);
			create index idx_game_features_feature_id ON game_features(feature_id);

			create index idx_game_genres_game_id ON game_genres(game_id);
			create index idx_game_genres_genre_id ON game_genres(genre_id);
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			drop table games;
			drop table tags;
			drop table game_tags;
			drop table developers;
			drop table game_developers;
			drop table publishers;
			drop table game_publishers;
			drop table features;
			drop table game_features;
			drop table genres;
			drop table game_genres;
		`)
		return err
	})
}
