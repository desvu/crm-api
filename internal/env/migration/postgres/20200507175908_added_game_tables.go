package postgres

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			create table games
			(
				id    uuid primary key,
				title text not null,
				type  int  not null
			);
			
			create table game_revisions
			(
				id           serial primary key,
				game_id      uuid         not null,
				status       int          not null,
				summary      text         not null,
				description  text         not null,
				slug 		 text         not null,
				license      text         not null,
				platforms    int[]        not null,
				release_date timestamp(0) not null,
				published_at timestamp(0),
				constraint fk_game_revisions_games foreign key (game_id) references games (id) on delete cascade
			);
			
			create table tags
			(
				id   serial primary key,
				name text not null
			);
			
			create table game_revision_tags
			(
				id               serial primary key,
				game_revision_id int not null,
				tag_id           int not null,
				constraint fk_game_revision_tags_game_revisions foreign key (game_revision_id) references game_revisions (id) on delete cascade,
				constraint fk_game_revision_tags_tags foreign key (tag_id) references tags (id) on delete cascade
			);
			
			create table developers
			(
				id   serial primary key,
				name text not null
			);
			
			create table game_revision_developers
			(
				id               serial primary key,
				game_revision_id int not null,
				developer_id     int not null,
				constraint fk_game_revision_developers_game_revisions foreign key (game_revision_id) references game_revisions (id) on delete cascade,
				constraint fk_game_revision_developers_developers foreign key (developer_id) references developers (id) on delete cascade
			);
			
			create table publishers
			(
				id   serial primary key,
				name text not null
			);
			
			create table game_revision_publishers
			(
				id               serial primary key,
				game_revision_id int not null,
				publisher_id     int not null,
				constraint fk_game_revision_publishers_games foreign key (game_revision_id) references game_revisions (id) on delete cascade,
				constraint fk_game_revision_publishers_publishers foreign key (publisher_id) references publishers (id) on delete cascade
			);
			
			create table features
			(
				id   serial primary key,
				name text not null
			);
			
			create table game_revision_features
			(
				id               serial primary key,
				game_revision_id int not null,
				feature_id       int not null,
				constraint fk_game_revision_features_games foreign key (game_revision_id) references game_revisions (id) on delete cascade,
				constraint fk_game_revision_features_features foreign key (feature_id) references features (id) on delete cascade
			);
			
			create table genres
			(
				id   serial primary key,
				name text not null
			);
			
			create table game_revision_genres
			(
				id               serial primary key,
				game_revision_id int not null,
				genre_id         int not null,
				constraint fk_game_revision_genres_games foreign key (game_revision_id) references game_revisions (id) on delete cascade,
				constraint fk_game_revision_genres_genres foreign key (genre_id) references genres (id) on delete cascade
			);
			
			create index idx_game_revision_tags_game_revision_id ON game_revision_tags (game_revision_id);
			create index idx_game_revision_tags_tag_id ON game_revision_tags (tag_id);
			
			create index idx_game_revision_developers_game_revision_id ON game_revision_developers (game_revision_id);
			create index idx_game_revision_developers_developer_id ON game_revision_developers (developer_id);
			
			create index idx_game_revision_publishers_game_revision_id ON game_revision_publishers (game_revision_id);
			create index idx_game_revision_publishers_publisher_id ON game_revision_publishers (publisher_id);
			
			create index idx_game_revision_features_game_revision_id ON game_revision_features (game_revision_id);
			create index idx_game_revision_features_feature_id ON game_revision_features (feature_id);
			
			create index idx_game_revision_genres_game_revision_id ON game_revision_genres (game_revision_id);
			create index idx_game_revision_genres_genre_id ON game_revision_genres (genre_id);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			drop table games;
			drop table game_revisions;
			drop table tags;
			drop table game_revision_tags;
			drop table developers;
			drop table game_revision_developers;
			drop table publishers;
			drop table game_revision_publishers;
			drop table features;
			drop table game_revision_features;
			drop table genres;
			drop table game_revision_genres;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200507175908_added_game_tables", up, down, opts)
}
