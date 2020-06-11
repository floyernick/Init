package controller

import (
	"net/http"

	"Init/config"
	"Init/database"
	"Init/notes"
)

func Init(config config.ServerConfig, db database.DB) error {

	mux := http.NewServeMux()
	notes.RegisterRoutes(mux, db)

	server := &http.Server{
		Addr:         config.Port,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      mux,
	}

	return server.ListenAndServe()

}
