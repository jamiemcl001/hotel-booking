package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
	db        *sql.DB
	db_driver database.Driver
	db_name   string
}

type MigratorOptFn func(m *Migrator)

var (
	ErrFailedToInitialiseMigrator = errors.New("failed to initialise migrator")
	ErrFailedToRunMigrations      = errors.New("failed to run migrations")
)

func WithDB(db *sql.DB) MigratorOptFn {
	return func(m *Migrator) {
		m.db = db
	}
}

func WithDBName(name string) MigratorOptFn {
	return func(m *Migrator) {
		m.db_name = name
	}
}

func NewMigrator(opts ...MigratorOptFn) *Migrator {
	m := new(Migrator)
	for _, opt := range opts {
		opt(m)
	}

	db, _ := sqlite3.WithInstance(m.db, &sqlite3.Config{})
	m.db_driver = db

	return m
}

func (m *Migrator) RunMigrations() error {
	mig, err := migrate.NewWithDatabaseInstance("file://migrations", "hotelsdb", m.db_driver)
	if err != nil {
		fmt.Println(err)
		return ErrFailedToInitialiseMigrator
	}

	return mig.Up()
}
