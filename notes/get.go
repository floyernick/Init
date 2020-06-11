package notes

import (
	"Init/app/errors"
	handler "Init/controller/shared"
	"Init/database"
	"Init/tools/validator"
	"net/http"
)

type NotesGetRequest struct {
	Id string `json:"id" validate:"required,min=36,max=36"`
}

type NotesGetResponse struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Data  string `json:"data"`
}

type NotesGetController struct {
	db database.DB
}

func (controller NotesGetController) Handler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var req NotesGetRequest

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

func (controller NotesGetController) Usecase(params NotesGetRequest) (NotesGetResponse, error) {

	var result NotesGetResponse

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams
	}

	note, err := controller.db.GetNoteById(params.Id)

	if err != nil {
		return result, errors.InternalError
	}

	if !note.Exists() {
		return result, errors.NoteNotFound
	}

	result = NotesGetResponse{
		Id:    note.Id,
		Title: note.Title,
		Data:  note.Data,
	}

	return result, nil

}
