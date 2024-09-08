package db

import (
	"database/sql"
	"errors"
)

var (
	ErrFailedToCreateDB    = errors.New("failed to create an instance of the DB")
	ErrFailedToConnectToDB = errors.New("failed to connect to the instance of the DB")
)

func New(source_url string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", source_url)
	if err != nil {
		return nil, ErrFailedToCreateDB
	}

	err = db.Ping()
	if err != nil {
		return nil, ErrFailedToConnectToDB
	}

	return db, nil
}
