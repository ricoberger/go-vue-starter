package api

import (
	"net/http"

	"github.com/ricoberger/go-vue-starter/pkg/api/response"
	"github.com/ricoberger/go-vue-starter/pkg/db"
	"github.com/ricoberger/go-vue-starter/pkg/mail"

	"github.com/gorilla/mux"
)

// Config represents the API configuration
type Config struct {
	Domain        string `yaml:"domain"`
	SigningSecret string `yaml:"signing_secret"`
}

// API represents the structure of the API
type API struct {
	Router *mux.Router

	config *Config
	db     db.DB
	mail   *mail.Client
}

// New returns the api settings
func New(config *Config, db db.DB, mail *mail.Client, router *mux.Router) (*API, error) {
	api := &API{
		config: config,
		db:     db,
		mail:   mail,
		Router: router,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	// Endpoint for healtcheck
	api.Router.HandleFunc("/api/v1/health", api.corsMiddleware(api.logMiddleware(api.healthHandler))).Methods("GET")

	// Account related api endpoints
	api.Router.HandleFunc("/api/v1/auth", api.corsMiddleware(api.logMiddleware(api.userLoginHandler))).Methods("POST")
	api.Router.HandleFunc("/api/v1/account", api.corsMiddleware(api.logMiddleware(api.userSignupHandler))).Methods("POST")
	api.Router.HandleFunc("/api/v1/account", api.corsMiddleware(api.logMiddleware(api.jwtMiddleware(api.userUpdateProfileHandler)))).Methods("PUT")
	api.Router.HandleFunc("/api/v1/account", api.corsMiddleware(api.logMiddleware(api.jwtMiddleware(api.userProfileHandler)))).Methods("GET")
	api.Router.HandleFunc("/api/v1/account/email/{id}/{token}", api.corsMiddleware(api.logMiddleware(api.userVerifyHandler))).Methods("GET")
	api.Router.HandleFunc("/api/v1/account/email", api.corsMiddleware(api.logMiddleware(api.userResendVerificationMail))).Methods("POST")
	api.Router.HandleFunc("/api/v1/account/password", api.corsMiddleware(api.logMiddleware(api.forgotPasswordHandler))).Methods("POST")
	api.Router.HandleFunc("/api/v1/account/password", api.corsMiddleware(api.logMiddleware(api.resetPasswordHandler))).Methods("PUT")

	return api, nil
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	response.Write(w, r, struct {
		Status string `json:"status"`
	}{
		"ok",
	})

	return
}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
