package notes

import (
	"Init/app/errors"
	handler "Init/controller/shared"
	"Init/database"
	"Init/entities"
	"Init/tools/uuid"
	"Init/tools/validator"
	"net/http"
)

type NotesCreateRequest struct {
	Title string `json:"title" validate:"required,min=1"`
	Data  string `json:"data" validate:"required,min=1"`
}

type NotesCreateResponse struct {
	Id string `json:"id"`
}

type NotesCreateController struct {
	db database.DB
}

func (controller NotesCreateController) Handler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var req NotesCreateRequest

		if err := handler.ParseRequestBody(r, &req); err != nil {
			handler.RespondWithError(w, errors.BadRequest)
		}

		res, err := controller.Usecase(req)

		if err != nil {
			handler.RespondWithError(w, err)
		} else {
			handler.RespondWithSuccess(w, res)
		}
	}

}

func (controller NotesCreateController) Usecase(params NotesCreateRequest) (NotesCreateResponse, error) {

	var result NotesCreateResponse

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams
	}

	note := entities.Note{
		Id:    uuid.Generate(),
		Title: params.Title,
		Data:  params.Data,
	}

	err := controller.db.CreateNote(note)

	if err != nil {
		return result, errors.InternalError
	}

	result = NotesCreateResponse{
		Id: note.Id,
	}

	return result, nil

}
