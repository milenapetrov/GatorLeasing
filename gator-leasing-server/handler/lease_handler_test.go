package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"GatorLeasing/gator-leasing-server/entity"
	"GatorLeasing/gator-leasing-server/service/mocks"
)

func CreateBody(request interface{}) *bytes.Reader {
	data, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	return reader
}

func TestGetAllLeasesOK(t *testing.T) {
	mockLeaseService := mocks.NewILeaseService(t)
	leases := []*entity.Lease{
		{ID: 1, Name: "Lease", OwnerID: 1},
		{ID: 2, Name: "Lease2", OwnerID: 2},
	}
	mockLeaseService.On("GetAllLeases").Return(leases, nil)

	leaseHandler := NewLeaseHandler(mockLeaseService)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/leases", nil)

	leaseHandler.GetAllLeases(rr, req)

	mockLeaseService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, rr.Code)
	expected := `[{"id":1,"name":"Lease","ownerID":1},{"id":2,"name":"Lease2","ownerID":2}]`
	assert.Equal(t, expected, rr.Body.String())
}

func TestPostLeaseOK(t *testing.T) {
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("CreateLease", mock.AnythingOfType("*entity.CreateLeaseRequest")).Return(uint(1), nil)

	leaseHandler := NewLeaseHandler(mockLeaseService)

	request := entity.CreateLeaseRequest{Name: "new lease"}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/leases", CreateBody(request))

	leaseHandler.PostLease(rr, req)

	mockLeaseService.AssertExpectations(t)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, `1`, rr.Body.String())
}

func TestPutLeaseOK(t *testing.T) {
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("EditLease", mock.AnythingOfType("*entity.EditLeaseRequest")).Return(nil)

	leaseHandler := NewLeaseHandler(mockLeaseService)

	request := entity.EditLeaseRequest{ID: 1, Name: "edited lease"}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/leases/1", CreateBody(request))
	r := mux.NewRouter()
	r.HandleFunc("/leases/{id}", leaseHandler.PutLease)
	r.ServeHTTP(rr, req)

	mockLeaseService.AssertExpectations(t)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Equal(t, string("null"), rr.Body.String())
}

func TestDeleteLeaseOK(t *testing.T) {
	mockLeaseService := mocks.NewILeaseService(t)
	mockLeaseService.On("DeleteLease", mock.AnythingOfType("*entity.DeleteLeaseRequest")).Return(nil)

	leaseHandler := NewLeaseHandler(mockLeaseService)

	request := entity.DeleteLeaseRequest{ID: 1}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/leases/1", CreateBody(request))
	r := mux.NewRouter()
	r.HandleFunc("/leases/{id}", leaseHandler.DeleteLease)
	r.ServeHTTP(rr, req)

	mockLeaseService.AssertExpectations(t)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Equal(t, string("null"), rr.Body.String())
}
