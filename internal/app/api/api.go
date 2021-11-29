package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// API Base API server instance description
type API struct {
	// Unexported field
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server/ configure loggers, router, database connections, etc..
func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("Started API server at port ", api.config.BindAddr)
	api.configureRouterField()
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
