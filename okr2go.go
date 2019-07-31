package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/oxisto/go-httputil"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := handlers.LoggingHandler(&httputil.LogWriter{Level: log.InfoLevel, Component: "http"}, NewRouter())
	err := http.ListenAndServe("0.0.0.0:4300", router)

	log.Errorf("An error occurred: %v", err)
}

type Objective struct {
	Name       string
	KeyResults []KeyResult
}

type KeyResult struct {
	ID           string
	Name         string
	Current      int
	Max          int
	Contributors []string
	Comments     []string
}

// NewRouter returns a configured mux router containing all REST endpoints
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/objectives", GetObjectives)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./okr2go-ui/dist")))

	return router
}

func GetObjectives(w http.ResponseWriter, r *http.Request) {
	var err error

	objectives := []Objective{
		{
			Name: "Become an awesome open source programmer",
			KeyResults: []KeyResult{
				{ID: "KR-1",
					Name:         "Create open source projects on GitHub",
					Current:      30,
					Max:          50,
					Contributors: []string{"oxisto"},
				},
				{ID: "KR-2",
					Name:         "Learn programming languages",
					Current:      4,
					Max:          5,
					Contributors: []string{"oxisto"},
					Comments:     []string{"Java", "go", "C++", "Rust"},
				},
			},
		},
	}

	httputil.JSONResponse(w, r, objectives, err)
}
