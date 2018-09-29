package database

import "Init/models"

type Database interface {
	GetUserById(int) (models.User, error)
}
