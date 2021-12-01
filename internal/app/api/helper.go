package api

import (
	"github.com/drTragger/powerfulAPI/internal/app/middleware"
	"github.com/drTragger/powerfulAPI/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	prefix = "/api/v1"
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
	api.router.HandleFunc(prefix+"/articles", api.GetAllArticles).Methods("GET")
	api.router.Handle(prefix+"/articles/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(api.GetArticleById),
	)).Methods("GET")
	api.router.HandleFunc(prefix+"/articles/{id}", api.DeleteArticleById).Methods("DELETE")
	api.router.HandleFunc(prefix+"/articles", api.PostArticle).Methods("POST")
	api.router.HandleFunc(prefix+"/users/register", api.PostUserRegister).Methods("POST")
	api.router.HandleFunc(prefix+"/users/auth", api.PostToAuth).Methods("POST")
}

func (api *API) configureStorageField() error {
	newStorage := storage.New(api.config.Storage)
	if err := newStorage.Open(); err != nil {
		return err
	}
	api.storage = newStorage
	return nil
}
