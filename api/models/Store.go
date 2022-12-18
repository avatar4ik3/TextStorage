package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

var chema = `
	create table if not exists texts(
		id bigserial not null primary key,
		value text
	)
`

func NewStore(url string) (*Store, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	_, err2 := db.Exec(chema)
	if err2 != nil {
		return nil, err
	}
	return &Store{
		db: db,
	}, nil
}

func (this *Store) Close() error {
	return this.db.Close()
}
