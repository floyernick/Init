package presenter

import (
	"net/http"

	"Init/config"
	"Init/controller"
)

type Presenter struct {
	controller controller.Controller
}

func Init(config config.ServerConfig, controller controller.Controller) error {

	presenter := Presenter{controller}

	mux := http.NewServeMux()
	mux.HandleFunc("/notes.create", handleRequest(presenter.NotesCreate))
	mux.HandleFunc("/notes.update", handleRequest(presenter.NotesUpdate))
	mux.HandleFunc("/notes.delete", handleRequest(presenter.NotesDelete))
	mux.HandleFunc("/notes.get", handleRequest(presenter.NotesGet))

	server := &http.Server{
		Addr:         config.Port,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      mux,
	}

	return server.ListenAndServe()

}
