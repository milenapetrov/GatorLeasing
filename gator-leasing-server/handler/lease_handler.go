package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/mapper"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/service"
	viewModel "github.com/milenapetrov/GatorLeasing/gator-leasing-server/view-model"
)

type LeaseHandler struct {
	leaseService service.ILeaseService
}

func NewLeaseHandler(leaseService service.ILeaseService) *LeaseHandler {
	return &LeaseHandler{leaseService: leaseService}
}

// GetAllLeases godoc
//
//	@Summary		List leases
//	@Description	get all leases
//	@Tags			leases
//	@Produce		json
//	@Success		200	{array}	viewModel.Lease
//	@Failure		500
//	@Router			/leases [get]
func (h *LeaseHandler) GetAllLeases(w http.ResponseWriter, r *http.Request) {
	leaseEntities, err, status := h.leaseService.GetAllLeases()
	if err != nil {
		respondError(w, status, err.Error())
		return
	}

	mapper := mapper.NewMapper(&entity.Lease{}, &viewModel.Lease{})
	leaseViewModels, err := mapper.MapSlice(leaseEntities)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJson(w, status, leaseViewModels)
}

// PostLease godoc
//
//	@Summary		Create a lease
//	@Description	post a lease
//	@Tags			leases
//	@Accept			json
//	@Produce		plain
//	@Param			createLeaseRequest	body		viewModel.CreateLease	true	"create lease"
//	@Success		201					{object}	uint64						"id of created lease"
//	@Failure		500
//	@Failure		400
//	@Router			/leases [post]
//	@Security		Auth0
func (h *LeaseHandler) PostLease(w http.ResponseWriter, r *http.Request) {
	createLeaseRequest := &viewModel.CreateLease{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(createLeaseRequest)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	mapper := mapper.NewMapper(&viewModel.CreateLease{}, &entity.CreateLease{})
	leaseToCreate, err := mapper.Map(createLeaseRequest)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err, status := h.leaseService.CreateLease(leaseToCreate)
	if err != nil {
		respondError(w, status, err.Error())
		return
	}

	respondJson(w, status, id)
}

// PutLease godoc
//
//	@Summary		Update a lease
//	@Description	update a lease by id
//	@Tags			leases
//	@Accept			json
//	@Param			id					path	uint64					true	"lease id"
//	@Param			editLeaseRequest	body	viewModel.EditLease	true	"edit lease request"
//	@Success		204
//	@Failure		500
//	@Failure		400
//	@Router			/leases/{id} [put]
//	@Security		Auth0
func (h *LeaseHandler) PutLease(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	editLeaseRequest := &viewModel.EditLease{ID: uint(id)}
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(editLeaseRequest); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	mapper := mapper.NewMapper(&viewModel.EditLease{}, &entity.EditLease{})
	leaseToEdit, err := mapper.Map(editLeaseRequest)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err, status := h.leaseService.EditLease(leaseToEdit)
	if err != nil {
		respondError(w, status, err.Error())
		return
	}

	respondJson(w, status, nil)
}

// DeleteLease godoc
//
//	@Summary		Delete a lease
//	@Description	delete a lease by id
//	@Tags			leases
//	@Param			id	path	uint64	true	"lease id"
//	@Success		204
//	@Failure		500
//	@Failure		400
//	@Router			/leases/{id} [delete]
//	@Security		Auth0
func (h *LeaseHandler) DeleteLease(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err, status := h.leaseService.DeleteLease(uint(id))
	if err != nil {
		respondError(w, status, err.Error())
		return
	}

	respondJson(w, status, nil)
}

func (h *LeaseHandler) GetPaginatedLeases(w http.ResponseWriter, r *http.Request) {
	paginatedLeasesViewModel := &viewModel.PaginatedLeasesRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(paginatedLeasesViewModel)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	requestMapper := mapper.NewMapper(&viewModel.PaginatedLeasesRequest{}, &entity.PaginatedLeasesRequest{})
	paginatedLeasesRequest, err := requestMapper.Map(paginatedLeasesViewModel)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	leases, paginationToken, count, err, status := h.leaseService.GetPaginatedLeases(paginatedLeasesRequest)
	if err != nil {
		respondError(w, status, err.Error())
		return
	}

	leasesMapper := mapper.NewMapper(&entity.Lease{}, &viewModel.Lease{})
	leasesResult, err := leasesMapper.MapSlice(leases)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	result := viewModel.PaginatedLeasesResult{
		Leases:          leasesResult,
		PaginationToken: paginationToken,
		Count:           uint(count),
	}

	respondJson(w, http.StatusOK, result)
}
