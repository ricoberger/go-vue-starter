package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ricoberger/go-vue-starter/pkg/server"
	"github.com/ricoberger/go-vue-starter/pkg/version"

	"github.com/sirupsen/logrus"
)

var (
	configFileFlag = flag.String("config.file", "config.yml", "Path to the configuration file.")
	versionFlag    = flag.Bool("version", false, "Show version information.")
	debugFlag      = flag.Bool("debug", false, "Show debug information.")
)

func init() {
	// Parse command-line flags
	flag.Parse()

	// Log settings
	if *debugFlag {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.TraceLevel)
	} else {
		logrus.SetReportCaller(false)
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	// Show version information
	if *versionFlag {
		fmt.Fprintln(os.Stdout, version.Print("starter"))
		os.Exit(0)
	}

	// Create server instance
	instance := server.NewInstance()

	// Interrupt handler
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		logrus.Infof("Received %s signal", <-c)
		instance.Shutdown()
	}()

	// Start server
	logrus.Infof("Starting starter %s", version.Info())
	logrus.Infof("Build context %s", version.BuildContext())
	instance.Start(*configFileFlag)
}
