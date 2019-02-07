package api

import (
	"net/http"

	"Init/models/request"
)

func (api API) NotesCreate(w http.ResponseWriter, r *http.Request) {

	var req request.NotesCreate

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.NotesCreate(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) NotesUpdate(w http.ResponseWriter, r *http.Request) {

	var req request.NotesUpdate

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.NotesUpdate(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) NotesDelete(w http.ResponseWriter, r *http.Request) {

	var req request.NotesDelete

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.NotesDelete(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) NotesGet(w http.ResponseWriter, r *http.Request) {

	var req request.NotesGet

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.NotesGet(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}
