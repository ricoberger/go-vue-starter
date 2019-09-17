package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ricoberger/go-vue-starter/pkg/static"

	"github.com/gorilla/mux"
)

// Config represents the app configuration
type Config struct{}

// App represents the structure of the app
type App struct {
	Router *mux.Router

	config *Config
}

// New returns the api settings
func New(config *Config, router *mux.Router) (*App, error) {
	app := &App{
		config: config,
		Router: router,
	}

	// Static assets handler
	staticHandler := http.StripPrefix("/", http.FileServer(static.Dir(false, "/web/vue.js/dist")))

	// Serve vue app
	vueApp, err := static.FSString(false, "/web/vue.js/dist/index.html")
	if err != nil {
		return nil, err
	}

	app.Router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_path := r.URL.Path
		// static files
		if strings.Contains(_path, ".") || _path == "/" {
			staticHandler.ServeHTTP(w, r)
			return
		}

		fmt.Fprintf(w, vueApp)
	})

	return app, nil
}
