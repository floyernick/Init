package usecases

import "Init/env/database"

type Controller struct {
	db database.Database
}

func NewController(db database.Database) Controller {
	return Controller{db}
}
