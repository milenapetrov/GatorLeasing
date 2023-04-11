package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/faker"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/service/mocks"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/shared"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/validator"
	viewModel "github.com/milenapetrov/GatorLeasing/gator-leasing-server/view-model"
)

func TestGetAllLeasesOK(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)
	leases := faker.FakeMany(&entity.Lease{}, 5)
	mockLeaseService.On("GetAllLeases").Return(leases, nil)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/leases", nil)

	leaseHandler.GetAllLeases(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "GetAllLeases", 1)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAllLeasesServiceErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("GetAllLeases").Return(nil, &shared.InternalServerError{})

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/leases", nil)

	leaseHandler.GetAllLeases(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "GetAllLeases", 1)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestPostLeaseOK(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("CreateLease", mock.AnythingOfType("*entity.CreateLease")).Return(uint(1), nil)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	createLease := viewModel.CreateLease{}
	faker.FakeData(&createLease)

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&createLease)
	req := httptest.NewRequest(http.MethodPost, "/leases", bytes.NewReader(body))

	leaseHandler.PostLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "CreateLease", 1)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestPostLeaseDecodeErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	badBody := struct {
		Name int `json:"name"`
	}{
		Name: 1,
	}

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&badBody)
	req := httptest.NewRequest(http.MethodPost, "/leases", bytes.NewReader(body))

	leaseHandler.PostLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNotCalled(t, "CreateLease")
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPostLeaseValidatorErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	createLease := viewModel.CreateLease{}
	faker.FakeData(&createLease)
	createLease.Name = "a"

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&createLease)
	req := httptest.NewRequest(http.MethodPost, "/leases", bytes.NewReader(body))

	leaseHandler.PostLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNotCalled(t, "CreateLease")
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPostLeaseServiceErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("CreateLease", mock.AnythingOfType("*entity.CreateLease")).Return(uint(0), &shared.InternalServerError{})

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	createLease := viewModel.CreateLease{}
	faker.FakeData(&createLease)

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&createLease)
	req := httptest.NewRequest(http.MethodPost, "/leases", bytes.NewReader(body))

	leaseHandler.PostLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "CreateLease", 1)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestPutLeaseOK(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("EditLease", mock.AnythingOfType("*entity.EditLease")).Return(nil)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	editLease := viewModel.EditLease{}
	faker.FakeData(&editLease)

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&editLease)
	req := httptest.NewRequest(http.MethodPut, "/leases", bytes.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id[0-9]+": "1"})

	leaseHandler.PutLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "EditLease", 1)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestPutLeaseBadPathParamErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	editLease := viewModel.EditLease{}
	faker.FakeData(&editLease)

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&editLease)
	req := httptest.NewRequest(http.MethodPut, "/leases", bytes.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id[0-9]+": "bad"})

	leaseHandler.PutLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNotCalled(t, "EditLease")
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPutLeaseDecodeErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	badBody := struct {
		Name int `json:"name"`
	}{
		Name: 1,
	}

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&badBody)
	req := httptest.NewRequest(http.MethodPut, "/leases", bytes.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id[0-9]+": "1"})

	leaseHandler.PutLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNotCalled(t, "EditLease")
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPutLeaseServiceErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("EditLease", mock.AnythingOfType("*entity.EditLease")).Return(&shared.InternalServerError{})

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	editLease := viewModel.EditLease{}
	faker.FakeData(&editLease)

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&editLease)
	req := httptest.NewRequest(http.MethodPut, "/leases", bytes.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id[0-9]+": "1"})

	leaseHandler.PutLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "EditLease", 1)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestDeleteLeaseOK(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("DeleteLease", mock.AnythingOfType("uint")).Return(nil)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/leases", nil)
	req = mux.SetURLVars(req, map[string]string{"id[0-9]+": "1"})

	leaseHandler.DeleteLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "DeleteLease", 1)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteLeaseBadPathParamErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/leases", nil)
	req = mux.SetURLVars(req, map[string]string{"id[0-9]+": "bad"})

	leaseHandler.DeleteLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNotCalled(t, "DeleteLease")
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteLeaseServiceErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("DeleteLease", mock.AnythingOfType("uint")).Return(&shared.InternalServerError{})

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/leases", nil)
	req = mux.SetURLVars(req, map[string]string{"id[0-9]+": "1"})

	leaseHandler.DeleteLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "DeleteLease", 1)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestGetPaginatedLeasesOK(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)
	leases := faker.FakeMany(&entity.Lease{}, 5)
	mockLeaseService.On("GetPaginatedLeases", mock.AnythingOfType("*entity.PaginatedLeasesRequest")).Return(leases, "paginationToken", int64(10), nil)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	paginatedLeasesRequest := viewModel.PaginatedLeasesRequest{}
	faker.FakeData(&paginatedLeasesRequest)

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&paginatedLeasesRequest)
	req := httptest.NewRequest(http.MethodGet, "/leases/paged", bytes.NewReader(body))

	leaseHandler.GetPaginatedLeases(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "GetPaginatedLeases", 1)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetPaginatedLeasesDecodeErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	badBody := struct {
		PageSize string `json:"pageSize"`
	}{
		PageSize: "size",
	}

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&badBody)
	req := httptest.NewRequest(http.MethodGet, "/leases/paged", bytes.NewReader(body))

	leaseHandler.GetPaginatedLeases(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNotCalled(t, "GetPaginatedLeases")
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetPaginateLeasesServiceErr(t *testing.T) {
	initializeTest()
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("GetPaginatedLeases", mock.AnythingOfType("*entity.PaginatedLeasesRequest")).Return(nil, "", int64(0), &shared.InternalServerError{})

	leaseHandler := NewLeaseHandler(mockLeaseService, validator.New())

	paginatedLeasesRequest := viewModel.PaginatedLeasesRequest{}
	faker.FakeData(&paginatedLeasesRequest)

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(&paginatedLeasesRequest)
	req := httptest.NewRequest(http.MethodGet, "/leases/paged", bytes.NewReader(body))

	leaseHandler.GetPaginatedLeases(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "GetPaginatedLeases", 1)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
