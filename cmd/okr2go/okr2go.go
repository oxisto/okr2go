package main

import (
	"net/http"
	"strings"

	"github.com/pkg/browser"

	"github.com/gorilla/handlers"
	"github.com/kyokomi/emoji"
	"github.com/oxisto/okr2go"
	"github.com/sirupsen/logrus"
)

func main() {
	log := okr2go.ConfigureLogging()

	emoji.Printf("Welcome to okr2go! Your :books:tracker is ready at http://localhost:4300.\n")

	_ = browser.OpenURL("http://localhost:4300")

	if err := okr2go.LoadObjectives(); err != nil {
		log.Errorf("An error occurred: %v", err)
		return
	}

	router := handlers.LoggingHandler(&LogWriter{Logger: log, Level: logrus.InfoLevel, Component: "http"}, okr2go.NewRouter())
	err := http.ListenAndServe("0.0.0.0:4300", router)

	log.Errorf("An error occurred: %v", err)
}

// LogWriter implements io.Writer and writes all incoming text out to the specified log level.
type LogWriter struct {
	Logger    *logrus.Logger
	Level     logrus.Level
	Component string
}

func (d LogWriter) Write(p []byte) (n int, err error) {
	var entry *logrus.Entry

	if d.Logger == nil {
		entry = logrus.WithField("component", d.Component)
	} else {
		entry = d.Logger.WithField("component", d.Component)
	}

	entry.Log(d.Level, strings.TrimRight(string(p), "\n"))

	return len(p), nil
}
