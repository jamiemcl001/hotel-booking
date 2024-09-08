package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func setupDb() *sql.DB {
	dbConn, err := sql.Open("sqlite3", "hotelsdb.sqlite3")
	if err != nil {
		log.Fatalf("Failed to initialise the DB: %s", err.Error())
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatalf("Failed to connect to DB: %s", err.Error())
	}

	return dbConn
}

func run() {
	db := setupDb()
	defer db.Close()
}

func main() {
	run()
}
