package storage

import "Init/models"

type Storage interface {
	noteStorage
}

type noteStorage interface {
	CreateNote(models.Note) error
	UpdateNote(models.Note) error
	DeleteNote(string) error
	GetNote(string) (models.Note, error)
}
