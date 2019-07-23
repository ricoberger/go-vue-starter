package server

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/ricoberger/go-vue-starter/pkg/api"
	"github.com/ricoberger/go-vue-starter/pkg/app"
	"github.com/ricoberger/go-vue-starter/pkg/db"
	"github.com/ricoberger/go-vue-starter/pkg/mail"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Config represents the server configuration
type Config struct {
	ListenAddress string `yaml:"listen_address"`

	API  *api.Config  `yaml:"api"`
	App  *app.Config  `yaml:"app"`
	DB   *db.Config   `yaml:"database"`
	Mail *mail.Config `yaml:"mail"`
}

// Instance represents an instance of the server
type Instance struct {
	API    *api.API
	App    *app.App
	Config *Config
	DB     db.DB
	Mail   *mail.Client

	httpServer *http.Server
}

// NewInstance returns an new instance of our server
func NewInstance() *Instance {
	return &Instance{}
}

// Start starts the server
func (i *Instance) Start(file string) {
	var err error
	var router = mux.NewRouter()

	// Load configuration file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.WithError(err).Fatal("Could not load configuration")
	}

	err = yaml.Unmarshal(data, &i.Config)
	if err != nil {
		logrus.WithError(err).Fatal("Could not load configuration")
	}

	// Establish database connection
	i.DB, err = db.NewConnection(i.Config.DB)
	if err != nil {
		logrus.WithError(err).Fatal("Could not open database connection")
	}
	defer i.DB.CloseConnection()

	// Setup mailing client
	i.Mail = mail.NewClient(i.Config.Mail)

	// Initialize API
	i.API, err = api.New(i.Config.API, i.DB, i.Mail, router)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create API instance")
	}

	i.App, err = app.New(i.Config.App, router)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create app instance")
	}

	// Startup the HTTP Server in a way that we can gracefully shut it down again
	i.httpServer = &http.Server{
		Addr:    i.Config.ListenAddress,
		Handler: router,
	}

	err = i.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.WithError(err).Error("HTTP Server stopped unexpected")
		i.Shutdown()
	} else {
		logrus.WithError(err).Info("HTTP Server stopped")
	}
}

// Shutdown stops the server
func (i *Instance) Shutdown() {
	// Shutdown all dependencies
	i.DB.CloseConnection()

	// Shutdown HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := i.httpServer.Shutdown(ctx)
	if err != nil {
		logrus.WithError(err).Error("Failed to shutdown HTTP server gracefully")
		os.Exit(1)
	}

	logrus.Info("Shutdown HTTP server...")
	os.Exit(0)
}
