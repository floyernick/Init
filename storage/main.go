package storage

import (
	"database/sql"

	_ "github.com/lib/pq"

	"Init/config"
)

type Service struct {
	pool *sql.DB
}

func Init(config config.DatabaseConfig) (Storage, error) {
	pool, err := sql.Open("postgres", config.Url)

	if err != nil {
		return nil, err
	}

	pool.SetConnMaxLifetime(config.ConnLifetime)
	pool.SetMaxOpenConns(config.OpenConns)
	pool.SetMaxIdleConns(config.IdleConns)

	if err := pool.Ping(); err != nil {
		return nil, err
	}

	storage := Service{pool}

	return storage, nil

}
