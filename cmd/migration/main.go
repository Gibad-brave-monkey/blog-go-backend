package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// TODO: implement
	// var migrationPath, storagePath, databaseName string

	// flag.StringVar(&storagePath, "storage-path", "", "path to storage path")
	// flag.StringVar(&databaseName, "storage-name", "", "database name")
	// flag.StringVar(&migrationPath, "migration-path", "", "path to migrations")
	// flag.Parse()

	m, err := migrate.New(
		"file://./migrations",
		"postgres://admin:admin@localhost:5432/mydatabase?sslmode=disable",
	)
	if err != nil {
		m.Log.Printf("Error to new Migrate", err.Error())
	}

	if err := m.Up(); err != nil {
		m.Log.Printf("Error to up migrations", err.Error())
	}

	log.Print("Successfully migrations up")
}
