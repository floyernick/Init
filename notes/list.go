package notes

import (
	"Init/app/errors"
	handler "Init/controller/shared"
	"Init/database"
	"Init/tools/validator"
	"net/http"
)

type NotesListRequest struct {
	Offset int `json:"offset" validate:"min=0"`
	Limit  int `json:"limit" validate:"required,min=1,max=100"`
}

type NotesListResponse struct {
	Notes []NotesListResponseNote `json:"notes"`
}

type NotesListResponseNote struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Data  string `json:"data"`
}

type NotesListController struct {
	db database.DB
}

func (controller NotesListController) Handler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var req NotesListRequest

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

func (controller NotesListController) Usecase(params NotesListRequest) (NotesListResponse, error) {

	var result NotesListResponse

	if err := validator.Process(params); err != nil {
		return result, errors.InvalidParams
	}

	notesBuilder := controller.db.GetBuilder().Offset(params.Offset).Limit(params.Limit)

	notes, err := controller.db.GetNotes(notesBuilder).Fetch()

	if err != nil {
		return result, errors.InternalError
	}

	for _, note := range notes {
		resultNote := NotesListResponseNote{
			Id:    note.Id,
			Title: note.Title,
			Data:  note.Data,
		}

		result.Notes = append(result.Notes, resultNote)
	}

	return result, nil

}
