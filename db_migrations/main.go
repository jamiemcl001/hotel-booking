package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrFailedToCreateDB    = errors.New("failed to create an instance of the DB")
	ErrFailedToConnectToDB = errors.New("failed to connect to the instance of the DB")
)

const DB_SOURCE = "../hotelsdb.sqlite3"

func newDB(source_url string) (*sql.DB, error) {
	f, _ := filepath.Abs(source_url)
	db_path := fmt.Sprintf("file://%s", f)
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		return nil, ErrFailedToCreateDB
	}

	if err = db.Ping(); err != nil {
		return nil, ErrFailedToConnectToDB
	}

	return db, nil
}

func main() {
	db, err := newDB(DB_SOURCE)
	if err != nil {
		log.Fatalf("Failed to initialise DB: %s", err.Error())
	}

	m := newMigrator(withDB(db), withDBName("hotelsdb"))
	if err = m.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %s", err.Error())
	}
}
