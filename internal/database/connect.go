package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const schema = `
CREATE TABLE IF NOT EXISTS news (
    id INTEGER PRIMARY KEY
);
`

func NewDB(source string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", source)
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return db, nil
}
