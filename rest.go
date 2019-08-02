package okr2go

import (
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/oxisto/go-httputil"
)

// NewRouter returns a configured mux router containing all REST endpoints
func NewRouter() *mux.Router {
	// pack angular ui
	box := packr.NewBox("./okr2go-ui/dist/okr2go-ui")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/objectives", GetObjectives)
	router.PathPrefix("/").Handler(http.FileServer(box))

	return router
}

func GetObjectives(w http.ResponseWriter, r *http.Request) {
	var err error

	objectives, err := ParseMarkdown("example.md")

	if err != nil {
		httputil.JSONResponse(w, r, nil, err)
		return
	}

	httputil.JSONResponse(w, r, objectives, err)
}
