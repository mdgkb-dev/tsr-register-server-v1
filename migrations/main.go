package migrations

import (
	"github.com/uptrace/bun/migrate"
	"log"
)

var Migrations = migrate.NewMigrations()

func init() {
	if err := Migrations.DiscoverCaller(); err != nil {
		log.Fatalln(err)
	}
}
