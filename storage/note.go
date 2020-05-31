package storage

import (
	"database/sql"

	"Init/models"
	"Init/tools/logger"
)

func (storage Service) CreateNote(note models.Note) error {

	query := "INSERT INTO notes(id, title, data) VALUES($1, $2, $3)"

	_, err := storage.Performer().Exec(query, note.Id, note.Title, note.Data)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) UpdateNote(note models.Note) error {

	query := "UPDATE notes SET title = $2, data = $3 WHERE id = $1"

	_, err := storage.Performer().Exec(query, note.Id, note.Title, note.Data)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) DeleteNote(id string) error {

	query := "DELETE FROM notes WHERE id = $1"

	_, err := storage.Performer().Exec(query, id)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (storage Service) GetNote(id string) (models.Note, error) {

	var note models.Note

	query := "SELECT id, title, data FROM notes WHERE id = $1"

	err := storage.Performer().QueryRow(query, id).Scan(&note.Id, &note.Title, &note.Data)

	if err != nil && err != sql.ErrNoRows {
		logger.Error(err.Error())
		return note, err
	}

	return note, nil
}
