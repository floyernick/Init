package notes

import (
	"Init/app/errors"
	handler "Init/controller/shared"
	"Init/database"
	"Init/tools/validator"
	"net/http"
)

type NotesDeleteRequest struct {
	Id string `json:"id" validate:"required,min=36,max=36"`
}

type NotesDeleteResponse struct{}

type NotesDeleteController struct {
	db database.DB
}

func (controller NotesDeleteController) Handler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var req NotesDeleteRequest

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

func (controller NotesDeleteController) Usecase(params NotesDeleteRequest) (NotesDeleteResponse, error) {

	var result NotesDeleteResponse

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

	err = controller.db.DeleteNote(note.Id)

	if err != nil {
		return result, errors.InternalError
	}

	return result, nil

}
