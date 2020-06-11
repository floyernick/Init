package notes

import (
	"Init/database"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db database.DB) {
	mux.HandleFunc("/notes.create", NotesCreateController{db}.Handler())
	mux.HandleFunc("/notes.update", NotesUpdateController{db}.Handler())
	mux.HandleFunc("/notes.delete", NotesDeleteController{db}.Handler())
	mux.HandleFunc("/notes.get", NotesGetController{db}.Handler())
	mux.HandleFunc("/notes.list", NotesListController{db}.Handler())
}
