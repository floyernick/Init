package postgres

import (
	"Init/env/config"
	"Init/env/database"
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewPostgresDatabase(config config.DatabaseConfig) (database.Database, error) {

	pool, err := sql.Open("postgres", config.Url)

	if err != nil {
		return nil, err
	}

	pool.SetConnMaxLifetime(config.ConnLifetime)
	pool.SetMaxOpenConns(config.OpenConns)
	pool.SetMaxIdleConns(config.IdleConns)

	if pool.Ping() != nil {
		return nil, err
	}

	db := &Database{pool}

	return db, nil

}
