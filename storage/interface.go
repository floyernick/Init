package storage

import (
	"Init/models"
	"database/sql"
)

type Storage interface {
	Transaction() (Storage, error)
	Commit() error
	Rollback() error
	noteStorage
}

type noteStorage interface {
	CreateNote(models.Note) error
	UpdateNote(models.Note) error
	DeleteNote(string) error
	GetNote(string) (models.Note, error)
	GetNotes() NoteQuerier
}

type Performer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type SelectQuerier interface {
	Equals(field string, value interface{}) SelectQuerier
	NotEquals(field string, value interface{}) SelectQuerier
	Greater(field string, value interface{}) SelectQuerier
	GreaterOrEquals(field string, value interface{}) SelectQuerier
	Less(field string, value interface{}) SelectQuerier
	LessOrEquals(field string, value interface{}) SelectQuerier
	Like(field string, value interface{}) SelectQuerier
	Contains(field string, value interface{}) SelectQuerier
	And() SelectQuerier
	Or() SelectQuerier
	Group() SelectQuerier
	EndGroup() SelectQuerier
	Paginate(offset int, limit int) SelectQuerier
	Order(field string, value string) SelectQuerier
}

type NoteQuerier interface {
	SelectQuerier
	Fetch() ([]models.Note, error)
	FetchOne() (models.Note, error)
}
