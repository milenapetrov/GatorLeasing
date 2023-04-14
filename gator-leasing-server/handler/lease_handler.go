package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/errors"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/mapper"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/service"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/validator"
	viewModel "github.com/milenapetrov/GatorLeasing/gator-leasing-server/view-model"
)

type LeaseHandler struct {
	leaseService service.ILeaseService
	validator    *validator.Validator
}

func NewLeaseHandler(leaseService service.ILeaseService, validator *validator.Validator) *LeaseHandler {
	return &LeaseHandler{leaseService: leaseService, validator: validator}
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
	leaseEntities, err := h.leaseService.GetAllLeases()
	if err != nil {
		respondError(w, err)
		return
	}

	mapper := mapper.NewMapper(&entity.Lease{}, &viewModel.Lease{})
	leaseViewModels, err := mapper.MapSlice(leaseEntities)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, http.StatusOK, leaseViewModels)
}

func (h *LeaseHandler) GetLeaseById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id[0-9]+"], 10, 32)
	if err != nil {
		respondError(w, &errors.BadRequestError{Msg: err.Error()})
		return
	}

	leaseEntity, err := h.leaseService.GetLeaseById(uint(id))
	if err != nil {
		respondError(w, err)
		return
	}

	mapper := mapper.NewMapper(&entity.Lease{}, &viewModel.Lease{})
	leaseViewModel, err := mapper.Map(leaseEntity)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, http.StatusOK, leaseViewModel)
}

// PostLease godoc
//
//	@Summary		Create a lease
//	@Description	post a lease
//	@Tags			leases
//	@Accept			json
//	@Produce		plain
//	@Param			createLeaseRequest	body		viewModel.CreateLease	true	"create lease"
//	@Success		201					{object}	uint64					"id of created lease"
//	@Failure		500
//	@Failure		400
//	@Router			/leases [post]
//	@Security		Auth0
func (h *LeaseHandler) PostLease(w http.ResponseWriter, r *http.Request) {
	createLeaseRequest := &viewModel.CreateLease{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(createLeaseRequest); err != nil {
		respondError(w, &errors.BadRequestError{Msg: err.Error()})
		return
	}

	if errs := h.validator.Struct(createLeaseRequest); errs != nil && len(errs) > 0 {
		respondErrors(w, errs)
		return
	}

	mapper := mapper.NewMapper(&viewModel.CreateLease{}, &entity.CreateLease{})
	leaseToCreate, err := mapper.Map(createLeaseRequest)
	if err != nil {
		respondError(w, err)
		return
	}

	id, err := h.leaseService.CreateLease(leaseToCreate)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, http.StatusCreated, id)
}

// PutLease godoc
//
//	@Summary		Update a lease
//	@Description	update a lease by id
//	@Tags			leases
//	@Accept			json
//	@Param			id					path	uint64				true	"lease id"
//	@Param			editLeaseRequest	body	viewModel.EditLease	true	"edit lease request"
//	@Success		204
//	@Failure		500
//	@Failure		400
//	@Router			/leases/{id} [put]
//	@Security		Auth0
func (h *LeaseHandler) PutLease(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id[0-9]+"], 10, 32)
	if err != nil {
		respondError(w, &errors.BadRequestError{Msg: err.Error()})
		return
	}

	editLeaseRequest := &viewModel.EditLease{ID: int(id)}
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(editLeaseRequest); err != nil {
		respondError(w, &errors.BadRequestError{Msg: err.Error()})
		return
	}

	mapper := mapper.NewMapper(&viewModel.EditLease{}, &entity.EditLease{})
	editLease, err := mapper.Map(editLeaseRequest)
	if err != nil {
		respondError(w, err)
		return
	}

	if err := h.leaseService.EditLease(editLease); err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, http.StatusNoContent, nil)
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
	id, err := strconv.ParseUint(params["id[0-9]+"], 10, 32)
	if err != nil {
		respondError(w, &errors.BadRequestError{Msg: err.Error()})
		return
	}

	if err := h.leaseService.DeleteLease(uint(id)); err != nil {
		respondError(w, err)
	}

	respondJson(w, http.StatusNoContent, nil)
}

// GetPaginatedLeases godoc
//
//	@Summary		Get paged leases
//	@Description	get paged leases
//	@Tags			leases
//	@Param			getPaginatedLeasesRequest	body		viewModel.PaginatedLeasesRequest	true	"page size, column to sort on, pagination token, sort direction, filter"
//	@Success		200							{object}	viewModel.PaginatedLeasesResult
//	@Failure		500
//	@Failure		400
//	@Router			/leases/paged [get]
func (h *LeaseHandler) GetPaginatedLeases(w http.ResponseWriter, r *http.Request) {
	paginatedLeasesViewModel := &viewModel.PaginatedLeasesRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(paginatedLeasesViewModel); err != nil {
		respondError(w, &errors.BadRequestError{Msg: err.Error()})
		return
	}

	requestMapper := mapper.NewMapper(&viewModel.PaginatedLeasesRequest{}, &entity.PaginatedLeasesRequest{})
	paginatedLeasesRequest, err := requestMapper.Map(paginatedLeasesViewModel)
	if err != nil {
		respondError(w, err)
		return
	}

	leases, paginationToken, count, err := h.leaseService.GetPaginatedLeases(paginatedLeasesRequest)
	if err != nil {
		respondError(w, err)
		return
	}

	leasesMapper := mapper.NewMapper(&entity.Lease{}, &viewModel.Lease{})
	leasesResult, err := leasesMapper.MapSlice(leases)
	if err != nil {
		respondError(w, err)
		return
	}

	result := viewModel.PaginatedLeasesResult{
		Leases:          leasesResult,
		PaginationToken: paginationToken,
		Count:           uint(count),
	}

	respondJson(w, http.StatusOK, result)
}

func (h *LeaseHandler) GetMyLeases(w http.ResponseWriter, r *http.Request) {
	myLeaseEntities, err := h.leaseService.GetMyLeases()
	if err != nil {
		respondError(w, err)
		return
	}

	mapper := mapper.NewMapper(&entity.Lease{}, &viewModel.Lease{})
	leaseViewModels, err := mapper.MapSlice(myLeaseEntities)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, http.StatusOK, leaseViewModels)
}

func (h *LeaseHandler) GetMyLeasesPaged(w http.ResponseWriter, r *http.Request) {
	myLeasesViewModel := &viewModel.PaginatedLeasesRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(myLeasesViewModel); err != nil {
		respondError(w, &errors.BadRequestError{Msg: err.Error()})
		return
	}

	requestMapper := mapper.NewMapper(&viewModel.PaginatedLeasesRequest{}, &entity.PaginatedLeasesRequest{})
	myLeasesRequest, err := requestMapper.Map(myLeasesViewModel)
	if err != nil {
		respondError(w, err)
		return
	}

	leases, paginationToken, count, err := h.leaseService.GetMyLeasesPaged(myLeasesRequest)
	if err != nil {
		respondError(w, err)
		return
	}

	leasesMapper := mapper.NewMapper(&entity.Lease{}, &viewModel.Lease{})
	leasesResult, err := leasesMapper.MapSlice(leases)
	if err != nil {
		respondError(w, err)
		return
	}

	result := viewModel.PaginatedLeasesResult{
		Leases:          leasesResult,
		PaginationToken: paginationToken,
		Count:           uint(count),
	}

	respondJson(w, http.StatusOK, result)
}
