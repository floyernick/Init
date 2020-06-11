package notes

import (
	"Init/app/errors"
	handler "Init/controller/shared"
	"Init/database"
	"Init/tools/validator"
	"net/http"
)

type NotesUpdateRequest struct {
	Id    string  `json:"id" validate:"required,min=36,max=36"`
	Title *string `json:"title" validate:"omitempty,min=1"`
	Data  *string `json:"data" validate:"omitempty,min=1"`
}

type NotesUpdateResponse struct{}

type NotesUpdateController struct {
	db database.DB
}

func (controller NotesUpdateController) Handler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var req NotesUpdateRequest

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

func (controller NotesUpdateController) Usecase(params NotesUpdateRequest) (NotesUpdateResponse, error) {

	var result NotesUpdateResponse

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

	if params.Title != nil {
		note.Title = *params.Title
	}

	if params.Data != nil {
		note.Data = *params.Data
	}

	err = controller.db.UpdateNote(note)

	if err != nil {
		return result, errors.InternalError
	}

	return result, nil

}
