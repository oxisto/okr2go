package okr2go

import (
	"embed"
	"encoding/json"
	"errors"
	"io/fs"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oxisto/go-httputil"
)

//go:embed okr2go-ui/dist/okr2go-ui/*
var content embed.FS

// NewRouter returns a configured mux router containing all REST endpoints
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/objectives", GetObjectives)
	router.HandleFunc("/api/objectives/{objectiveID}/results/{resultID}/plus", ResultPlusOne)
	router.HandleFunc("/api/objectives/{objectiveID}/results/{resultID}/minus", ResultMinusOne)
	router.Methods("POST").Path("/api/objectives/{objectiveID}/results").HandlerFunc(PostKeyResult)

	fsys, _ := fs.Sub(content, "okr2go-ui/dist/okr2go-ui")

	router.PathPrefix("/").Handler(http.FileServer(http.FS(fsys)))

	return router
}

func PostKeyResult(w http.ResponseWriter, r *http.Request) {
	var (
		result    KeyResult
		objective *Objective
		err       error
	)

	// @todo Create utility function for the NotFound/BadRequest construct
	// @body Return StatusNotFound if object is nil and StatusBadRequest if error
	objective, err = getObjectiveFromRequest(w, r)
	if err != nil {
		httputil.JSONResponseWithStatus(w, r, nil, err, http.StatusBadRequest)
		return
	}

	if objective == nil {
		httputil.JSONResponseWithStatus(w, r, nil, err, http.StatusNotFound)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&result); err != nil {
		httputil.JSONResponse(w, r, nil, err)
	}

	objective.KeyResults = append(objective.KeyResults, &result)

	httputil.JSONResponse(w, r, result, nil)
}

func GetObjectives(w http.ResponseWriter, r *http.Request) {
	var err error

	httputil.JSONResponse(w, r, objectives, err)
}

// @todo Combine ResultPlusOne and ResultMinusOne functions
func ResultPlusOne(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		result *KeyResult
	)

	result, err = getResultFromRequest(w, r)
	if err != nil {
		httputil.JSONResponseWithStatus(w, r, nil, err, http.StatusBadRequest)
		return
	}

	if result == nil {
		httputil.JSONResponseWithStatus(w, r, nil, err, http.StatusNotFound)
		return
	}

	if result.Current == result.Target {
		httputil.JSONResponseWithStatus(w, r, result, nil, http.StatusNotModified)
		return
	}

	result.Current++

	SaveObjectives()

	httputil.JSONResponseWithStatus(w, r, result, nil, http.StatusOK)
	return
}

func ResultMinusOne(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		result *KeyResult
	)

	result, err = getResultFromRequest(w, r)
	if err != nil {
		httputil.JSONResponseWithStatus(w, r, nil, err, http.StatusBadRequest)
		return
	}

	if result == nil {
		httputil.JSONResponseWithStatus(w, r, nil, err, http.StatusNotFound)
		return
	}

	if result.Current == 0 {
		httputil.JSONResponseWithStatus(w, r, result, nil, http.StatusNotModified)
		return
	}

	result.Current--

	SaveObjectives()

	httputil.JSONResponseWithStatus(w, r, result, nil, http.StatusOK)
	return
}

func getObjectiveFromRequest(w http.ResponseWriter, r *http.Request) (objective *Objective, err error) {
	var (
		ok                bool
		objectiveID       int
		objectiveIDString string
	)

	if objectiveIDString, ok = mux.Vars(r)["objectiveID"]; !ok {
		return nil, errors.New("Request did not contain a resultID")
	}

	if objectiveID, err = strconv.Atoi(objectiveIDString); err != nil {
		return nil, errors.New("Could not parse objectiveID")
	}

	objective = objectives.FindObjective(objectiveID)

	return objective, nil
}

func getResultFromRequest(w http.ResponseWriter, r *http.Request) (result *KeyResult, err error) {
	var (
		ok        bool
		objective *Objective
		resultID  string
	)

	// return if we either have an error (which we pass down, resulting in a BadRequest)
	// or if the objective was not found (resulting in a NotFound)
	if objective, err = getObjectiveFromRequest(w, r); err != nil || objective == nil {
		return nil, err
	}

	if resultID, ok = mux.Vars(r)["resultID"]; !ok {
		return nil, errors.New("Request did not contain a resultID")
	}

	result = objective.FindKeyResult(resultID)

	return result, nil
}
