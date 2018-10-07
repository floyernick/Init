package main

import (
	"Init/env/config"
	"Init/env/database/postgres"
	"Init/env/server"
	"Init/env/transfer/api"
	"Init/env/validator/validator"
	"Init/tools/logger"
	"Init/usecases"
	"flag"
)

func main() {

	envName := *flag.String("c", "default.cfg", "Environment config name")

	flag.Parse()

	cfg, err := config.NewConfig(envName)

	if err != nil {
		logger.Error(err.Error())
	}

	db, err := postgres.NewPostgresDatabase(cfg.Db)

	if err != nil {
		logger.Error(err.Error())
	}

	validate := validator.NewValidator()

	controller := usecases.NewController(db, validate)

	APIHandler := api.NewAPIHandler(cfg.Handler, controller)

	err = server.RunServer(cfg.Server, APIHandler)

	if err != nil {
		logger.Error(err.Error())
	}

}
