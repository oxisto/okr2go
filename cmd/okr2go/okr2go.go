package main

import (
	"net/http"

	"github.com/pkg/browser"

	"github.com/gorilla/handlers"
	"github.com/kyokomi/emoji"
	"github.com/oxisto/go-httputil"
	"github.com/oxisto/okr2go"
	"github.com/sirupsen/logrus"
)

func main() {
	log := okr2go.ConfigureLogging()

	emoji.Printf("Welcome to okr2go! Your :books:tracker is ready at http://localhost:4300.\n")

	browser.OpenURL("http://localhost:4300")

	router := handlers.LoggingHandler(&httputil.LogWriter{Logger: log, Level: logrus.InfoLevel, Component: "http"}, okr2go.NewRouter())
	err := http.ListenAndServe("0.0.0.0:4300", router)

	log.Errorf("An error occurred: %v", err)
}
