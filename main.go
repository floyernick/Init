package main

import (
	"flag"

	"Init/config"
	"Init/controller"
	"Init/database"
	"Init/tools/logger"
)

func main() {

	environmentName := *flag.String("c", "default.yaml", "Environment config name")

	flag.Parse()

	cfg, err := config.LoadConfig(environmentName)

	if err != nil {
		logger.Error(err)
	}

	db, err := database.Init(cfg.Database)

	if err != nil {
		logger.Error(err)
	}

	err = controller.Init(cfg.Server, db)

	if err != nil {
		logger.Error(err)
	}

}
