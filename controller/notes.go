package controller

import (
	"Init/app/errors"
	"Init/models"
	"Init/models/request"
	"Init/models/response"
	"Init/tools/uuid"
	"Init/tools/validator"
)

func (controller Controller) NotesCreate(params request.NotesCreate) (response.NotesCreate, error) {

	var result response.NotesCreate

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams{}
	}

	note := models.Note{
		Id:    uuid.Generate(),
		Title: params.Title,
		Data:  params.Data,
	}

	err := controller.storage.CreateNote(note)

	if err != nil {
		return result, errors.InternalError{}
	}

	result = response.NotesCreate{
		Id: note.Id,
	}

	return result, nil

}

func (controller Controller) NotesUpdate(params request.NotesUpdate) (response.NotesUpdate, error) {

	var result response.NotesUpdate

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams{}
	}

	note, err := controller.storage.GetNote(params.Id)

	if err != nil {
		return result, errors.InternalError{}
	}

	if !note.Exists() {
		return result, errors.NoteNotFound{}
	}

	if params.Title != nil {
		note.Title = *params.Title
	}

	if params.Data != nil {
		note.Data = *params.Data
	}

	err = controller.storage.UpdateNote(note)

	if err != nil {
		return result, errors.InternalError{}
	}

	result = response.NotesUpdate{}

	return result, nil

}

func (controller Controller) NotesDelete(params request.NotesDelete) (response.NotesDelete, error) {

	var result response.NotesDelete

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams{}
	}

	note, err := controller.storage.GetNote(params.Id)

	if err != nil {
		return result, errors.InternalError{}
	}

	if !note.Exists() {
		return result, errors.NoteNotFound{}
	}

	err = controller.storage.DeleteNote(note.Id)

	if err != nil {
		return result, errors.InternalError{}
	}

	result = response.NotesDelete{}

	return result, nil

}

func (controller Controller) NotesGet(params request.NotesGet) (response.NotesGet, error) {

	var result response.NotesGet

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams{}
	}

	note, err := controller.storage.GetNote(params.Id)

	if err != nil {
		return result, errors.InternalError{}
	}

	if !note.Exists() {
		return result, errors.NoteNotFound{}
	}

	result = response.NotesGet{
		Id:    note.Id,
		Title: note.Title,
		Data:  note.Data,
	}

	return result, nil

}
