package postgres

import (
	"database/sql"

	"Init/models"
	"Init/tools/logger"
)

func (db Database) CreateNote(note models.Note) error {

	query := "INSERT INTO notes(id, title, data) VALUES($1, $2, $3)"

	_, err := db.db.Exec(query, note.Id, note.Title, note.Data)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (db Database) UpdateNote(note models.Note) error {

	query := "UPDATE notes SET title = $2, data = $3 WHERE id = $1"

	_, err := db.db.Exec(query, note.Id, note.Title, note.Data)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (db Database) DeleteNote(id string) error {

	query := "DELETE FROM notes WHERE id = $1"

	_, err := db.db.Exec(query, id)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (db Database) GetNote(id string) (models.Note, error) {

	var note models.Note

	query := "SELECT id, title, data FROM notes WHERE id = $1"

	err := db.db.QueryRow(query, id).Scan(&note.Id, &note.Title, &note.Data)

	if err != nil && err != sql.ErrNoRows {
		logger.Error(err.Error())
		return note, err
	}

	return note, nil
}
