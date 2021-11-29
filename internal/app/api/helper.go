package api

import (
	"github.com/drTragger/powerfulAPI/storage"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func (api *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(logLevel)
	return nil
}

func (api *API) configureRouterField() {
	api.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		Logger(writer.Write([]byte("Hello! This is REST API!")))
	})
}

func (api *API) configureStorageField() error {
	newStorage := storage.New(api.config.Storage)
	if err := newStorage.Open(); err != nil {
		return err
	}
	api.storage = newStorage
	return nil
}

func Logger(_ int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}
