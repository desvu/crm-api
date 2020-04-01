package postgres

import (
	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
)

func Migrate(db *pg.DB) error {
	_, _, err := migrations.Run(db, "init")
	if err != nil {
		return err
	}

	_, _, err = migrations.Run(db, "up")
	if err != nil {
		return err
	}

	return nil
}
