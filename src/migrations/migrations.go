package migrations

import (
	"embed"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//go:embed *.sql
var migrationsFile embed.FS

func RunMigrations(dbUrl string) error {
	// d, err := iofs.New(migrationsFile, ".")
	// if err != nil {
	// 	return err
	// }

	// m, err := migrate.NewWithSourceInstance("iofs", d, dbUrl)
	// if err != nil {
	// 	return err
	// }

	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	return err
	// }

	// return nil

	m, err := migrate.New(
		"file://src/migrations",
		dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
