package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"GatorLeasing/gator-leasing-server/entity"
	"GatorLeasing/gator-leasing-server/service"
)

type LeaseHandler struct {
	leaseService *service.LeaseService
}

func NewLeaseHandler(leaseService *service.LeaseService) *LeaseHandler {
	return &LeaseHandler{
		leaseService: leaseService,
	}
}

func (h *LeaseHandler) GetAllLeases(w http.ResponseWriter, r *http.Request) {
	leases, err := h.leaseService.GetAllLeases()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJson(w, http.StatusOK, leases)
	return
}

func (h *LeaseHandler) PostLease(w http.ResponseWriter, r *http.Request) {
	var request entity.CreateLeaseRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.leaseService.CreateLease(&request)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJson(w, http.StatusCreated, id)
	return
}

func (h *LeaseHandler) PutLease(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}

	request := entity.EditLeaseRequest{ID: uint(id)}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.leaseService.EditLease(&request)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJson(w, http.StatusNoContent, nil)
	return
}

func (h *LeaseHandler) DeleteLease(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}

	request := entity.DeleteLeaseRequest{ID: uint(id)}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.leaseService.DeleteLease(&request)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJson(w, http.StatusNoContent, nil)
	return
}
