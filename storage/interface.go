package storage

import (
	"Init/models"
	"database/sql"
)

type Performer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Storage interface {
	Performer() Performer
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
}

type QueryBuilder interface {
	Equals(field string, value interface{}) string
	NotEquals(field string, value interface{}) string
	Greater(field string, value interface{}) string
	GreaterOrEquals(field string, value interface{}) string
	Less(field string, value interface{}) string
	LessOrEquals(field string, value interface{}) string
	Like(field string, value interface{}) string
	Contains(field string, value interface{}) string
	And(conditions ...string) string
	Or(conditions ...string) string
	Add(condition string)
	Paginate(offset int, limit int)
	Order(field string, value string)
	Count() int
}

type NoteQuery interface {
	QueryBuilder
	Fetch() []models.Note
	FetchOne() models.Note
}
