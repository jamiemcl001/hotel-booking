package main

import (
	_ "github.com/mattn/go-sqlite3"

	"github.com/jamiemcl001/hotel-booking/db"
)

func run() {
	db_conn, err := db.New("hotelsdb.sqlite3")
	if err != nil {
		panic(err)
	}

	m := db.NewMigrator(
		db.WithDB(db_conn),
		db.WithDBName("hotelsdb"),
	)
	err = m.RunMigrations()
	if err != nil {
		panic(err)
	}

	defer db_conn.Close()
}

func main() {
	run()
}
