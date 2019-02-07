package usecases

import (
	"errors"

	"Init/models"
	"Init/models/request"
	"Init/models/response"
	"Init/tools/uuid"
)

func (controller Controller) NotesCreate(req request.NotesCreate) (response.NotesCreate, error) {

	var res response.NotesCreate

	if err := controller.validator.Process(req); err != nil {
		return res, errors.New("invalid params")
	}

	note := models.Note{
		Id:    uuid.Generate(),
		Title: req.Title,
		Data:  req.Data,
	}

	err := controller.db.CreateNote(note)

	if err != nil {
		return res, errors.New("internal error")
	}

	res = response.NotesCreate{
		Id: note.Id,
	}

	return res, nil

}

func (controller Controller) NotesUpdate(req request.NotesUpdate) (response.NotesUpdate, error) {

	var res response.NotesUpdate

	if err := controller.validator.Process(req); err != nil {
		return res, errors.New("invalid params")
	}

	note, err := controller.db.GetNote(req.Id)

	if err != nil {
		return res, errors.New("internal error")
	}

	if note.IsEmpty() {
		return res, errors.New("invalid note id")
	}

	if req.Title != nil {
		note.Title = *req.Title
	}

	if req.Data != nil {
		note.Data = *req.Data
	}

	err = controller.db.UpdateNote(note)

	if err != nil {
		return res, errors.New("internal error")
	}

	res = response.NotesUpdate{}

	return res, nil

}

func (controller Controller) NotesDelete(req request.NotesDelete) (response.NotesDelete, error) {

	var res response.NotesDelete

	if err := controller.validator.Process(req); err != nil {
		return res, errors.New("invalid params")
	}

	note, err := controller.db.GetNote(req.Id)

	if err != nil {
		return res, errors.New("internal error")
	}

	if note.IsEmpty() {
		return res, errors.New("invalid note id")
	}

	err = controller.db.DeleteNote(note.Id)

	if err != nil {
		return res, errors.New("internal error")
	}

	res = response.NotesDelete{}

	return res, nil

}

func (controller Controller) NotesGet(req request.NotesGet) (response.NotesGet, error) {

	var res response.NotesGet

	if err := controller.validator.Process(req); err != nil {
		return res, errors.New("invalid params")
	}

	note, err := controller.db.GetNote(req.Id)

	if err != nil {
		return res, errors.New("internal error")
	}

	if note.IsEmpty() {
		return res, errors.New("invalid note id")
	}

	res = response.NotesGet{
		Id:    note.Id,
		Title: note.Title,
		Data:  note.Data,
	}

	return res, nil

}
