package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/faker"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/service/mocks"
	viewModel "github.com/milenapetrov/GatorLeasing/gator-leasing-server/view-model"
)

func initialize() {
	faker.InitializeFaker()
}

func TestGetAllLeasesOK(t *testing.T) {
	initialize()
	mockLeaseService := mocks.NewILeaseService(t)
	leases := faker.FakeMany(&entity.Lease{}, 5)
	mockLeaseService.On("GetAllLeases").Return(leases, nil, http.StatusOK)

	leaseHandler := NewLeaseHandler(mockLeaseService)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/leases", nil)

	leaseHandler.GetAllLeases(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "GetAllLeases", 1)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestPostLeaseOK(t *testing.T) {
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("CreateLease", mock.AnythingOfType("*entity.CreateLease")).Return(uint(1), nil, http.StatusCreated)

	leaseHandler := NewLeaseHandler(mockLeaseService)

	createLease := viewModel.CreateLease{}
	faker.FakeData(&createLease)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/leases", createBody(createLease))

	leaseHandler.PostLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "CreateLease", 1)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestPutLeaseOK(t *testing.T) {
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("EditLease", mock.AnythingOfType("*entity.EditLease")).Return(nil, http.StatusNoContent)

	leaseHandler := NewLeaseHandler(mockLeaseService)

	editLease := viewModel.Lease{}
	faker.FakeData(&editLease)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/leases", createBody(editLease))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	leaseHandler.PutLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "EditLease", 1)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestDeleteLeaseOK(t *testing.T) {
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("DeleteLease", mock.AnythingOfType("uint")).Return(nil, http.StatusNoContent)

	leaseHandler := NewLeaseHandler(mockLeaseService)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/leases", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	leaseHandler.DeleteLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	mockLeaseService.AssertNumberOfCalls(t, "DeleteLease", 1)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
