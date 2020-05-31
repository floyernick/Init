package storage

import (
	"database/sql"

	_ "github.com/lib/pq"

	"Init/config"
)

type Service struct {
	pool *sql.DB
	tx   *sql.Tx
}

func (service Service) Performer() Performer {
	if service.tx != nil {
		return service.tx
	}
	return service.pool
}

func (service Service) Transaction() (Storage, error) {

	tx, err := service.pool.Begin()

	if err != nil {
		return nil, err
	}

	storage := Service{nil, tx}

	return storage, nil

}

func (service Service) Commit() error {

	if service.tx == nil {
		return nil
	}

	err := service.tx.Commit()

	return err

}

func (service Service) Rollback() error {

	if service.tx == nil {
		return nil
	}

	err := service.tx.Rollback()

	return err

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

	storage := Service{pool, nil}

	return storage, nil

}
