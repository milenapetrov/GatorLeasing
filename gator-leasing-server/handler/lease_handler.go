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

// swagger:route GET /leases leases GetAllLeases
//
// List leases.
//
// get all leases.
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Deprecated: false
//
//
//	Responses:
//	  200: GetLeaseResponse[]
//	  500: ErrorResponse[]
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
//
// swagger:route POST /leases leases PostLease
//
// Create a lease.
//
// post a lease
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//
//		Deprecated: false
//
//		Security:
//		  oauth2:
//
//	Parameters:
//	  + name: limit
//	    in: body
//	    description: create a lease
//	    required: true
//	    type: viewModel.CreateLease
//
//
//	Responses:
//	  201: GetLeaseResponse
//	  400: ErrorResponse[]
//	  500: ErrorResponse[]
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

// swagger:route PUT /leases/{id} leases PutLease
//
// Update a lease.
//
// update a lease by id.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Deprecated: false
//
//	Security:
//	  oauth2:
//
//	Parameters:
//	  + name: id
//	    in: query
//	    description: lease id
//	    required: true
//	    type: integer
//	    format: uint64
//	  + name: editLeaseRequest
//	    in: body
//	    description: edit lease request
//	    required: true
//	    type: viewModel.EditLease
//
//
//	Responses:
//	  204: NoContentResponse
//	  400: ErrorResponse[]
//	  500: ErrprResponse[]
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

// swagger:route DELETE /leases/{id} leases DeleteLease
//
// Delete a lease.
//
// delete a lease by id.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Deprecated: false
//
//	Security:
//	  oauth2:
//
//	Parameters:
//	  + name: id
//	    in: path
//	    description: leaseid
//	    required: true
//	    type: integer
//	    format: uint64
//
//
//	Responses:
//	  204: NoContentResponse
//	  400: ErrorResponse[]
//	  500: ErrorResponse[]
func (h *LeaseHandler) DeleteLease(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id[0-9]+"], 10, 32)
	if err != nil {
		respondError(w, &errors.BadRequestError{Msg: err.Error()})
		return
	}

	if err := h.leaseService.DeleteLease(uint(id)); err != nil {
		respondError(w, err)
		return
	}

	respondJson(w, http.StatusNoContent, nil)
}

// swagger:route GET /leases/paged pets users listPets
//
// Get paged leases.
//
// get paged leases.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Deprecated: false
//
//	Security:
//	  oauth2:
//
//	Parameters:
//	  + name: getPaginatedLeasesRequest
//	    in: body
//	    description: page size, column to sort on, pagination token, sort direction, filter
//	    required: true
//	    type: viewModel.PaginatedLeasesRequest
//
//
//	Responses:
//	  200: GetPaginatedLeasesResponse
//	  400: ErrorResponse[]
//	  500: ErrorResponse[]
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
