package api

// API Base API server instance description
type API struct {
	// Unexported field
	config *Config
}

// New API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
	}
}

// Start http server/ configure loggers, router, database connections, etc..
func (api *API) Start() error {
	return nil
}
