package service

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/dto"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/faker"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/repository/mocks"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/shared"
)

func TestGetAllLeasesOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	leases := faker.FakeMany(&dto.Lease{}, 5)
	mockLeaseRepository.On("GetAllLeases").Return(leases, nil, http.StatusOK)

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	resultLeases, resultErr := leaseService.GetAllLeases()

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "GetAllLeases", 1)
	assert.NotEmpty(t, resultLeases)
	assert.Nil(t, resultErr)
}

func TestGetAllLeasesErr(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	leases := []*dto.Lease{}
	err := &shared.InternalServerError{}
	mockLeaseRepository.On("GetAllLeases").Return(leases, err)

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	resultLeases, resultErr := leaseService.GetAllLeases()

	mockLeaseRepository.AssertExpectations(t)
	assert.Nil(t, resultLeases)
	assert.NotNil(t, resultErr)
	_, ok := resultErr.(*shared.InternalServerError)
	assert.True(t, ok)
}

func TestCreateLeaseOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("CreateLease", mock.AnythingOfType("*dto.Lease")).Return(uint(1), nil)

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	request := &entity.CreateLease{}
	faker.FakeData(request)
	resultID, resultErr := leaseService.CreateLease(request)

	mockLeaseRepository.AssertExpectations(t)
	assert.NotNil(t, resultID)
	assert.Nil(t, resultErr)
}

func TestEditLeaseOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("EditLease", mock.AnythingOfType("*dto.Lease")).Return(nil)

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	request := &entity.EditLease{}
	faker.FakeData(request)
	resultErr := leaseService.EditLease(request)

	mockLeaseRepository.AssertExpectations(t)
	assert.Nil(t, resultErr)
}

func TestDeleteLeaseOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("DeleteLease", mock.AnythingOfType("*dto.Lease")).Return(nil)

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	resultErr := leaseService.DeleteLease(0)

	mockLeaseRepository.AssertExpectations(t)
	assert.Nil(t, resultErr)
}
