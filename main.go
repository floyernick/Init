package main

import (
	"Init/env/config"
	"Init/env/database/postgres"
	"Init/env/server"
	"Init/env/transfer/api"
	"Init/usecases"
	"flag"
)

func main() {

	envName := *flag.String("c", "default.cfg", "Environment config name")

	flag.Parse()

	cfg, err := config.NewConfig(envName)

	if err != nil {
		panic(err)
	}

	db, err := postgres.NewPostgresDatabase(cfg.Db)

	if err != nil {
		panic(err)
	}

	controller := usecases.NewController(db)

	APIHandler := api.NewAPIHandler(cfg.Handler, controller)

	err = server.RunServer(cfg.Server, APIHandler)

	if err != nil {
		panic(err)
	}

}
