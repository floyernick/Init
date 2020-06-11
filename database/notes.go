package database

import (
	"database/sql"

	"Init/entities"
	"Init/tools/logger"
)

type NotesQueryBuilder struct {
	queryBuilder
}

func (builder NotesQueryBuilder) FetchOne() (entities.Note, error) {

	query := "SELECT id, title, data FROM notes"
	query = builder.formatQuery(query)

	var note entities.Note

	err := builder.db.performer().QueryRow(query, builder.params...).Scan(&note.Id, &note.Title, &note.Data)

	if err != nil && err != sql.ErrNoRows {
		if builder.db.tx != nil {
			builder.db.tx.Rollback()
		}
		logger.Warning(err)
		return note, err
	}

	return note, nil
}

func (builder NotesQueryBuilder) Fetch() ([]entities.Note, error) {

	query := "SELECT id, title, data FROM notes"
	query = builder.formatQuery(query)

	var notes []entities.Note

	rows, err := builder.db.performer().Query(query, builder.params...)

	if err != nil && err != sql.ErrNoRows {
		if builder.db.tx != nil {
			builder.db.tx.Rollback()
		}
		logger.Warning(err)
		return notes, err
	}

	defer rows.Close()

	for rows.Next() {
		var note entities.Note
		err := rows.Scan(&note.Id, &note.Title, &note.Data)
		if err != nil {
			if builder.db.tx != nil {
				builder.db.tx.Rollback()
			}
			return notes, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (db DB) CreateNote(note entities.Note) error {

	query := "INSERT INTO notes(id, title, data) VALUES($1, $2, $3)"

	_, err := db.performer().Exec(query, note.Id, note.Title, note.Data)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return err
	}

	return nil
}

func (db DB) UpdateNote(note entities.Note) error {

	query := "UPDATE notes SET title = $2, data = $3 WHERE id = $1"

	_, err := db.performer().Exec(query, note.Id, note.Title, note.Data)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return err
	}

	return nil
}

func (db DB) DeleteNote(id string) error {

	query := "DELETE FROM notes WHERE id = $1"

	_, err := db.performer().Exec(query, id)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return err
	}

	return nil
}

func (db DB) GetNotes(builder queryBuilder) NotesQueryBuilder {
	return NotesQueryBuilder{builder}
}

func (db DB) GetNoteById(id string) (entities.Note, error) {
	builder := db.GetBuilder().Equals("id", id)
	return db.GetNotes(builder).FetchOne()
}
