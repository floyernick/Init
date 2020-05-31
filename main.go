package main

import (
	"flag"

	"Init/config"
	"Init/controller"
	"Init/presenter"
	"Init/storage"
	"Init/tools/logger"
)

func main() {

	environmentName := *flag.String("c", "default.yaml", "Environment config name")

	flag.Parse()

	configData, err := config.LoadConfig(environmentName)

	if err != nil {
		logger.Error(err.Error())
	}

	storageService, err := storage.Init(configData.Database)

	if err != nil {
		logger.Error(err.Error())
	}

	controllerService := controller.Init(storageService)

	err = presenter.Init(configData.Server, controllerService)

	if err != nil {
		logger.Error(err.Error())
	}

}
