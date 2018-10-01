package postgres

import (
	"Init/models"
	"Init/tools/logger"
	"database/sql"
)

func (db *Database) GetUserById(id int) (models.User, error) {

	var user models.User

	query := "SELECT user_id, name FROM users WHERE user_id = $1"

	err := db.db.QueryRow(query, id).Scan(&user.Id, &user.Name)

	if err != nil && err != sql.ErrNoRows {
		logger.Error(err.Error())
		return user, err
	}

	return user, nil
}
