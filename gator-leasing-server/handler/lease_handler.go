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
	leaseService service.ILeaseService
}

func NewLeaseHandler(leaseService service.ILeaseService) *LeaseHandler {
	return &LeaseHandler{leaseService: leaseService}
}

// GetAllLeases godoc
//	@Summary		List leases
//	@Description	get all leases
//	@Tags			leases
//	@Produce		json
//	@Success		200	{array}	entity.Lease
//	@Failure		500
//	@Router			/leases [get]
func (h *LeaseHandler) GetAllLeases(w http.ResponseWriter, r *http.Request) {
	leases, err := h.leaseService.GetAllLeases()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJson(w, http.StatusOK, leases)
	return
}

// PostLease godoc
//	@Summary		Create a lease
//	@Description	post a lease
//	@Tags			leases
//	@Accept			json
//	@Produce		plain
//	@Param			createLeaseRequest	body		entity.CreateLeaseRequest	true	"create lease request"
//	@Success		201					{object}	uint64						"id of created lease"
//	@Failure		500
//	@Failure		400
//	@Router			/leases [post]
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

// PutLease godoc
//	@Summary		Update a lease
//	@Description	update a lease by id
//	@Tags			leases
//	@Accept			json
//	@Param			id					path	uint64					true	"lease id"
//	@Param			editLeaseRequest	body	entity.EditLeaseRequest	true	"edit lease request"
//	@Success		204
//	@Failure		500
//	@Failure		400
//	@Router			/leases/{id} [put]
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

// DeleteLease godoc
//	@Summary		Delete a lease
//	@Description	delete a lease by id
//	@Tags			leases
//	@Param			id	path	uint64	true	"lease id"
//	@Success		204
//	@Failure		500
//	@Failure		400
//	@Router			/leases/{id} [delete]
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
