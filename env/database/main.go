package database

import "Init/models"

type Database interface {
	CreateNote(models.Note) error
	UpdateNote(models.Note) error
	DeleteNote(string) error
	GetNote(string) (models.Note, error)
}
