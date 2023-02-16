package handler

import (
	"encoding/json"
	"net/http"
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

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJson(w, code, map[string]string{"error": message})
}
