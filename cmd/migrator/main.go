package main

import (
	"errors"
	"flag"
	"fmt"

	// library for migrations
	"github.com/golang-migrate/migrate/v4"
	// driver fo work sqlite
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	// driver for load migrations with file
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationPath, migrationTable string
	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationPath, "migration-path", "", "path to migrtion")
	flag.StringVar(&migrationTable, "migration-table", "migrations", "name of migrations")
	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}

	if migrationPath == "" {
		panic("migration-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration to apply")
		}
	}

	fmt.Println("All migrations to apply")
}
