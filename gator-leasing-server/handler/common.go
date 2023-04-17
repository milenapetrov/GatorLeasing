package handler

import (
	"encoding/json"
	stdErrors "errors"
	"net/http"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/errors"
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
		respondError(w, &errors.InternalServerError{Msg: "error response with no errors"})
	}
	errs := []error{}
	if stdErrors.As(err, &errors.INTERNAL_SERVER_ERROR) {
		errs = append(errs, err)
		respondJson(w, http.StatusInternalServerError, errs)
	} else if stdErrors.As(err, &errors.BAD_REQUEST_ERROR) {
		errs = append(errs, err)
		respondJson(w, http.StatusBadRequest, errs)
	} else {
		respondError(w, &errors.InternalServerError{Msg: "unknown error type"})
	}
	return
}

func respondErrors(w http.ResponseWriter, errs []error) {
	if errs == nil || len(errs) == 0 {
		respondError(w, &errors.InternalServerError{Msg: "errors response with no errors"})
		return
	}

	if stdErrors.As(errs[0], &errors.INTERNAL_SERVER_ERROR) {
		respondJson(w, http.StatusInternalServerError, errs)
	} else if stdErrors.As(errs[0], &errors.BAD_REQUEST_ERROR) {
		respondJson(w, http.StatusBadRequest, errs)
	} else {
		respondError(w, &errors.InternalServerError{Msg: "unknown error type"})
	}
	return
}
