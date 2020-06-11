package database

import (
	"Init/tools/logger"
)

func (db *DB) Transaction() error {
	if db.tx != nil {
		return nil
	}

	tx, err := db.pool.Begin()

	if err != nil {
		logger.Warning(err)
		return err
	}

	db.tx = tx

	return nil

}

func (db DB) Commit() error {

	if db.tx == nil {
		return nil
	}

	err := db.tx.Commit()

	if err != nil {
		logger.Warning(err)
		return err
	}

	return nil

}
