package okr2go

import (
	"errors"
	"net/http"
	"strconv"

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
	router.HandleFunc("/api/objectives/{objectiveID}/{resultID}/plus", ResultPlusOne)
	router.HandleFunc("/api/objectives/{objectiveID}/{resultID}/minus", ResultMinusOne)
	router.PathPrefix("/").Handler(http.FileServer(box))

	return router
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
