package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

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
