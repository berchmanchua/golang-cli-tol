package db

import (
	"database/sql"

	sqlitevec "github.com/asg017/sqlite-vec-go-bindings/cgo"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(path string) (*Database, error) {
	sqlitevec.Auto()

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	database := &Database{db: db}
	if err := database.migrate(); err != nil {
		return nil, err
	}

	return database, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
