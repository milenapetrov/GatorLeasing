package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"GatorLeasing/gator-leasing-server/entity"
	"GatorLeasing/gator-leasing-server/model"
	"GatorLeasing/gator-leasing-server/repository/mocks"
)

func TestGetAllLeasesOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	leases := []model.Lease{
		{Name: "Lease"},
	}
	mockLeaseRepository.On("GetAllLeases").Return(leases, nil)

	leaseService := NewLeaseService(mockLeaseRepository)

	resultLeases, err := leaseService.GetAllLeases()

	mockLeaseRepository.AssertExpectations(t)
	assert.NotEmpty(t, resultLeases)
	assert.Nil(t, err)
}

func TestGetAllLeasesErr(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	leases := []model.Lease{
		{Name: "Lease"},
	}
	err := errors.New("error")
	mockLeaseRepository.On("GetAllLeases").Return(leases, err)

	leaseService := NewLeaseService(mockLeaseRepository)

	resultLeases, resultErr := leaseService.GetAllLeases()

	mockLeaseRepository.AssertExpectations(t)
	assert.Nil(t, resultLeases)
	assert.NotNil(t, resultErr)
}

func TestCreateLeaseOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("CreateLease", mock.AnythingOfType("*model.Lease")).Return(uint(1), nil)

	leaseService := NewLeaseService(mockLeaseRepository)

	request := entity.CreateLeaseRequest{Name: "lease"}
	resultID, resultErr := leaseService.CreateLease(&request)

	mockLeaseRepository.AssertExpectations(t)
	assert.NotNil(t, resultID)
	assert.Nil(t, resultErr)
}

func TestEditLeaseOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("EditLease", mock.AnythingOfType("*model.Lease")).Return(nil)

	leaseService := NewLeaseService(mockLeaseRepository)

	request := entity.EditLeaseRequest{ID: 0, Name: "lease"}
	resultErr := leaseService.EditLease(&request)

	mockLeaseRepository.AssertExpectations(t)
	assert.Nil(t, resultErr)
}

func TestDeleteLeaseOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("DeleteLease", mock.AnythingOfType("*model.Lease")).Return(nil)

	leaseService := NewLeaseService(mockLeaseRepository)

	request := entity.DeleteLeaseRequest{ID: 0, Name: "lease"}
	resultErr := leaseService.DeleteLease(&request)

	mockLeaseRepository.AssertExpectations(t)
	assert.Nil(t, resultErr)
	assert.True(t, false)
}
