package presenter

import (
	"net/http"

	"Init/app/errors"
	"Init/models/request"
)

func (presenter Presenter) NotesCreate(r *http.Request) (interface{}, error) {

	var req request.NotesCreate

	if err := parseRequestBody(r, &req); err != nil {
		return nil, errors.BadRequest{}
	}

	res, err := presenter.controller.NotesCreate(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (presenter Presenter) NotesUpdate(r *http.Request) (interface{}, error) {

	var req request.NotesUpdate

	if err := parseRequestBody(r, &req); err != nil {
		return nil, errors.BadRequest{}
	}

	res, err := presenter.controller.NotesUpdate(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (presenter Presenter) NotesDelete(r *http.Request) (interface{}, error) {

	var req request.NotesDelete

	if err := parseRequestBody(r, &req); err != nil {
		return nil, errors.BadRequest{}
	}

	res, err := presenter.controller.NotesDelete(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (presenter Presenter) NotesGet(r *http.Request) (interface{}, error) {

	var req request.NotesGet

	if err := parseRequestBody(r, &req); err != nil {
		return nil, errors.BadRequest{}
	}

	res, err := presenter.controller.NotesGet(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}
