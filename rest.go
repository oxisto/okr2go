package okr2go

import (
	"embed"
	"encoding/json"
	"errors"
	"io/fs"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//go:embed ui/build/* ui/build/_app/pages/* ui/build/_app/assets/pages/*
var content embed.FS

// NewRouter returns a configured mux router containing all REST endpoints
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/objectives", GetObjectives)
	router.HandleFunc("/api/objectives/{objectiveID}/results/{resultID}/plus", ResultPlusOne)
	router.HandleFunc("/api/objectives/{objectiveID}/results/{resultID}/minus", ResultMinusOne)
	router.Methods("POST").Path("/api/objectives/{objectiveID}/results").HandlerFunc(PostKeyResult)

	fsys, _ := fs.Sub(content, "ui/build")

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
		JSONResponseWithStatus(w, r, nil, err, http.StatusBadRequest)
		return
	}

	if objective == nil {
		JSONResponseWithStatus(w, r, nil, err, http.StatusNotFound)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&result); err != nil {
		JSONResponse(w, r, nil, err)
	}

	objective.KeyResults = append(objective.KeyResults, &result)

	JSONResponse(w, r, result, nil)
}

func GetObjectives(w http.ResponseWriter, r *http.Request) {
	var err error

	JSONResponse(w, r, objectives, err)
}

// @todo Combine ResultPlusOne and ResultMinusOne functions
func ResultPlusOne(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		result *KeyResult
	)

	result, err = getResultFromRequest(w, r)
	if err != nil {
		JSONResponseWithStatus(w, r, nil, err, http.StatusBadRequest)
		return
	}

	if result == nil {
		JSONResponseWithStatus(w, r, nil, err, http.StatusNotFound)
		return
	}

	if result.Current == result.Target {
		JSONResponseWithStatus(w, r, result, nil, http.StatusNotModified)
		return
	}

	result.Current++

	err = SaveObjectives()
	if err != nil {
		JSONResponseWithStatus(w, r, nil, err, http.StatusBadRequest)
		return
	}

	JSONResponseWithStatus(w, r, result, nil, http.StatusOK)
}

func ResultMinusOne(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		result *KeyResult
	)

	result, err = getResultFromRequest(w, r)
	if err != nil {
		JSONResponseWithStatus(w, r, nil, err, http.StatusBadRequest)
		return
	}

	if result == nil {
		JSONResponseWithStatus(w, r, nil, err, http.StatusNotFound)
		return
	}

	if result.Current == 0 {
		JSONResponseWithStatus(w, r, result, nil, http.StatusNotModified)
		return
	}

	result.Current--

	err = SaveObjectives()
	if err != nil {
		JSONResponseWithStatus(w, r, nil, err, http.StatusBadRequest)
		return
	}

	JSONResponseWithStatus(w, r, result, nil, http.StatusOK)
}

func getObjectiveFromRequest(w http.ResponseWriter, r *http.Request) (objective *Objective, err error) {
	var (
		ok                bool
		objectiveID       int
		objectiveIDString string
	)

	if objectiveIDString, ok = mux.Vars(r)["objectiveID"]; !ok {
		return nil, errors.New("request did not contain a resultID")
	}

	if objectiveID, err = strconv.Atoi(objectiveIDString); err != nil {
		return nil, errors.New("could not parse objectiveID")
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
		return nil, errors.New("request did not contain a resultID")
	}

	result = objective.FindKeyResult(resultID)

	return result, nil
}

// JSONResponseWithStatus returns a JSON encoded object with statusCode, if error is nil.
// Otherwise the error is returned and status code is set to http.StatusInternalServerError
func JSONResponseWithStatus(w http.ResponseWriter, r *http.Request, object interface{}, err error, statusCode int) {
	// uh-uh, we have an error
	if err != nil {
		logrus.Errorf("An error occured during processing of a REST request: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return not found if object is nil
	if object == nil {
		http.NotFound(w, r)
		return
	}

	// otherwise, lets try to decode the JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(object); err != nil {
		// uh-uh we couldn't decode the JSON
		logrus.Errorf("An error occured during encoding of the JSON response: %v. Payload was: %+v", err, object)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// JSONResponse returns a JSON encoded object with http.StatusOK, if error is nil.
// Otherwise the error is returned and status code is set to http.StatusInternalServerError
func JSONResponse(w http.ResponseWriter, r *http.Request, object interface{}, err error) {
	JSONResponseWithStatus(w, r, object, err, http.StatusOK)
}
