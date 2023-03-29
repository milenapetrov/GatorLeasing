package handler

import (
	"encoding/json"
	"net/http"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/shared"
)

// respondJSON makes the response with payload as json format
func respondJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		errorJson, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errorJson))
		return
	}

	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, err error) {
	if err == nil {
		respondError(w, &shared.InternalServerError{Msg: "error response with no errors"})
	}
	errs := []error{}
	if _, ise := err.(*shared.InternalServerError); ise {
		errs = append(errs, &shared.InternalServerError{Msg: err.Error()})
		respondJson(w, http.StatusInternalServerError, errs)
	} else if _, br := err.(*shared.BadRequestError); br {
		errs = append(errs, &shared.BadRequestError{Msg: err.Error()})
		respondJson(w, http.StatusBadRequest, errs)
	} else {
		respondError(w, &shared.InternalServerError{Msg: "unknown error type"})
	}
	return
}

func respondErrors(w http.ResponseWriter, errs []error) {
	if errs == nil || len(errs) == 0 {
		respondError(w, &shared.InternalServerError{Msg: "errors response with no errors"})
		return
	}

	if _, ise := errs[0].(*shared.InternalServerError); ise {
		respondJson(w, http.StatusInternalServerError, errs)
	} else if _, br := errs[0].(*shared.BadRequestError); br {
		respondJson(w, http.StatusBadRequest, errs)
	} else {
		respondError(w, &shared.InternalServerError{Msg: "unknown error type"})
	}
	return
}
