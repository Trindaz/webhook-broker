// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/imyousuf/webhook-broker/config"
	"github.com/imyousuf/webhook-broker/controllers"
	"github.com/imyousuf/webhook-broker/storage"
)

// Injectors from wire.go:

func GetAppVersion() config.AppVersion {
	appVersion := config.GetVersion()
	return appVersion
}

func GetHTTPServer(cliConfig *config.CLIConfig) (*HTTPServiceContainer, error) {
	configConfig, err := config.GetConfigurationFromCLIConfig(cliConfig)
	if err != nil {
		return nil, err
	}
	serverLifecycleListenerImpl := NewServerListener()
	migrationConfig := GetMigrationConfig(cliConfig)
	dataAccessor, err := storage.GetNewDataAccessor(configConfig, migrationConfig, configConfig)
	if err != nil {
		return nil, err
	}
	appRepository := newAppRepository(dataAccessor)
	statusController := controllers.NewStatusController(appRepository)
	producerRepository := newProducerRepository(dataAccessor)
	producerController := controllers.NewProducerController(producerRepository)
	producersController := controllers.NewProducersController(producerRepository, producerController)
	channelRepository := newChannelRepository(dataAccessor)
	channelController := controllers.NewChannelController(channelRepository)
	controllersControllers := &controllers.Controllers{
		StatusController:    statusController,
		ProducersController: producersController,
		ProducerController:  producerController,
		ChannelController:   channelController,
	}
	router := controllers.NewRouter(controllersControllers)
	server := controllers.ConfigureAPI(configConfig, serverLifecycleListenerImpl, router)
	httpServiceContainer := NewHTTPServiceContainer(configConfig, serverLifecycleListenerImpl, server, dataAccessor)
	return httpServiceContainer, nil
}
