package main

import (
	"flag"
	"log"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/env/migration/postgres"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func main() {
	var uri = flag.String("uri", "postgres://user:password@localhost:5434/crm?sslmode=disable", "postgres uri")
	flag.Parse()

	opts, err := pg.ParseURL(*uri)
	if err != nil {
		log.Fatalln(err)
	}

	var db = pg.Connect(opts)

	err = migrations.Run(db, postgres.Directory, append([]string{"migrations"}, flag.Args()...))
	if err != nil {
		log.Fatalln(err)
	}
}
