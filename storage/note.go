package storage

import (
	"database/sql"

	"Init/models"
	"Init/tools/logger"
)

type NoteQuery struct {
	*SelectQuery
}

func (builder *NoteQuery) FetchOne() (models.Note, error) {

	query := "SELECT id, title, data FROM notes"
	query = builder.formatQuery(query)

	var note models.Note

	err := builder.storage.performer().QueryRow(query, builder.params...).Scan(&note.Id, &note.Title, &note.Data)

	if err != nil && err != sql.ErrNoRows {
		logger.Error(err.Error())
		return note, err
	}

	return note, nil
}

func (builder *NoteQuery) Fetch() ([]models.Note, error) {

	query := "SELECT id, title, data FROM notes"
	query = builder.formatQuery(query)

	var notes []models.Note

	rows, err := builder.storage.performer().Query(query, builder.params...)

	if err != nil && err != sql.ErrNoRows {
		logger.Error(err.Error())
		return notes, err
	}

	defer rows.Close()

	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.Id, &note.Title, &note.Data)
		if err != nil {
			return notes, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (storage Service) CreateNote(note models.Note) error {

	query := "INSERT INTO notes(id, title, data) VALUES($1, $2, $3)"

	_, err := storage.performer().Exec(query, note.Id, note.Title, note.Data)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) UpdateNote(note models.Note) error {

	query := "UPDATE notes SET title = $2, data = $3 WHERE id = $1"

	_, err := storage.performer().Exec(query, note.Id, note.Title, note.Data)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) DeleteNote(id string) error {

	query := "DELETE FROM notes WHERE id = $1"

	_, err := storage.performer().Exec(query, id)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) GetNotes() NoteQuerier {
	builder := &NoteQuery{&SelectQuery{
		storage: storage,
	}}
	return builder
}

func (storage Service) GetNote(id string) (models.Note, error) {
	builder := storage.GetNotes()
	builder.Equals("id", id)
	return builder.FetchOne()
}
