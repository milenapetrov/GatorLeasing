package service

import (
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
	mockLeaseRepository.On("GetAllLeases").Return(leases, nil)

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	resultLeases, resultErr := leaseService.GetAllLeases()

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "GetAllLeases", 1)
	assert.NotEmpty(t, resultLeases)
	assert.Nil(t, resultErr)
}

func TestGetAllLeasesRepositoryErr(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("GetAllLeases").Return(nil, &shared.InternalServerError{})

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	resultLeases, resultErr := leaseService.GetAllLeases()

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "GetAllLeases", 1)
	assert.Empty(t, resultLeases)
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
	mockLeaseRepository.AssertNumberOfCalls(t, "CreateLease", 1)
	assert.NotZero(t, resultID)
	assert.Nil(t, resultErr)
}

func TestCreateLeaseRepositoryErr(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("CreateLease", mock.AnythingOfType("*dto.Lease")).Return(uint(0), &shared.InternalServerError{})

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	request := &entity.CreateLease{}
	faker.FakeData(request)
	resultID, resultErr := leaseService.CreateLease(request)

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "CreateLease", 1)
	assert.Zero(t, resultID)
	assert.NotNil(t, resultErr)
	_, ok := resultErr.(*shared.InternalServerError)
	assert.True(t, ok)
}

func TestEditLeaseOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	lease := &dto.Lease{}
	faker.FakeData(lease)
	mockLeaseRepository.On("GetLeaseById", lease.ID).Return(lease, nil)
	mockLeaseRepository.On("EditLease", mock.AnythingOfType("*dto.Lease")).Return(nil)

	userContext := shared.NewUserContext()
	userContext.ID = lease.OwnerID
	leaseService := NewLeaseService(userContext, mockLeaseRepository)

	editLease := &entity.EditLease{}
	faker.FakeData(editLease)
	editLease.ID = lease.ID
	resultErr := leaseService.EditLease(editLease)

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "GetLeaseById", 1)
	mockLeaseRepository.AssertNumberOfCalls(t, "EditLease", 1)
	assert.Nil(t, resultErr)
}

func TestEditLeaseRepositoryErr(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("GetLeaseById", mock.AnythingOfType("uint")).Return(nil, &shared.InternalServerError{})

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	editLease := &entity.EditLease{}
	faker.FakeData(editLease)
	resultErr := leaseService.EditLease(editLease)

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "GetLeaseById", 1)
	mockLeaseRepository.AssertNotCalled(t, "EditLease")
	assert.NotNil(t, resultErr)
	_, ok := resultErr.(*shared.InternalServerError)
	assert.True(t, ok)
}

func TestEditLeaseMismatchedUserIdErr(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	lease := &dto.Lease{}
	faker.FakeData(lease)
	lease.OwnerID = 1
	mockLeaseRepository.On("GetLeaseById", lease.ID).Return(lease, nil)

	userContext := shared.NewUserContext()
	userContext.ID = 2
	leaseService := NewLeaseService(userContext, mockLeaseRepository)

	editLease := &entity.EditLease{}
	faker.FakeData(editLease)
	editLease.ID = lease.ID
	resultErr := leaseService.EditLease(editLease)

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "GetLeaseById", 1)
	mockLeaseRepository.AssertNotCalled(t, "EditLease")
	assert.NotNil(t, resultErr)
	_, ok := resultErr.(*shared.BadRequestError)
	assert.True(t, ok)
}

func TestDeleteLeaseOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	lease := &dto.Lease{}
	faker.FakeData(lease)
	mockLeaseRepository.On("GetLeaseById", lease.ID).Return(lease, nil)
	mockLeaseRepository.On("DeleteLease", mock.AnythingOfType("*dto.Lease")).Return(nil)

	userContext := shared.NewUserContext()
	userContext.ID = lease.OwnerID
	leaseService := NewLeaseService(userContext, mockLeaseRepository)

	resultErr := leaseService.DeleteLease(lease.ID)

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "GetLeaseById", 1)
	mockLeaseRepository.AssertNumberOfCalls(t, "DeleteLease", 1)
	assert.Nil(t, resultErr)
}

func TestDeleteLeaseRepositoryErr(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("GetLeaseById", mock.AnythingOfType("uint")).Return(nil, &shared.InternalServerError{})

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	resultErr := leaseService.DeleteLease(0)

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "GetLeaseById", 1)
	mockLeaseRepository.AssertNotCalled(t, "DeleteLease")
	assert.NotNil(t, resultErr)
	_, ok := resultErr.(*shared.InternalServerError)
	assert.True(t, ok)
}

func TestDeleteLeaseMismatchedUserIdErr(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	lease := &dto.Lease{}
	faker.FakeData(lease)
	lease.OwnerID = 1
	mockLeaseRepository.On("GetLeaseById", lease.ID).Return(lease, nil)

	userContext := shared.NewUserContext()
	userContext.ID = 2
	leaseService := NewLeaseService(userContext, mockLeaseRepository)

	resultErr := leaseService.DeleteLease(lease.ID)

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "GetLeaseById", 1)
	mockLeaseRepository.AssertNotCalled(t, "DeleteLease")
	assert.NotNil(t, resultErr)
	_, ok := resultErr.(*shared.BadRequestError)
	assert.True(t, ok)
}

func TestGetPaginatedLeasesOK(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	leases := faker.FakeMany(&dto.Lease{}, 10)
	mockLeaseRepository.On("GetPaginatedLeases",
		mock.AnythingOfType("uint"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("enums.SortDirection"),
		mock.AnythingOfType("string")).Return(leases, "paginationToken", int64(10), nil)

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	paginatedLeasesRequest := &entity.PaginatedLeasesRequest{}
	faker.FakeData(paginatedLeasesRequest)
	resultLeases, resultPaginationToken, resultCount, resultErr := leaseService.GetPaginatedLeases(paginatedLeasesRequest)

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNumberOfCalls(t, "GetPaginatedLeases", 1)
	assert.NotEmpty(t, resultLeases)
	assert.NotZero(t, resultPaginationToken)
	assert.NotZero(t, resultCount)
	assert.Nil(t, resultErr)
}

func TestGetPaginatedLeasesRepositoryErr(t *testing.T) {
	mockLeaseRepository := mocks.NewILeaseRepository(t)
	mockLeaseRepository.On("GetPaginatedLeases",
		mock.AnythingOfType("uint"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("enums.SortDirection"),
		mock.AnythingOfType("string")).Return(nil, "", int64(0), &shared.InternalServerError{})

	leaseService := NewLeaseService(shared.NewUserContext(), mockLeaseRepository)

	paginatedLeasesRequest := &entity.PaginatedLeasesRequest{}
	faker.FakeData(paginatedLeasesRequest)
	resultLeases, resultPaginationToken, resultCount, resultErr := leaseService.GetPaginatedLeases(paginatedLeasesRequest)

	mockLeaseRepository.AssertExpectations(t)
	mockLeaseRepository.AssertNotCalled(t, "GetPaginatedLeases", 1)
	assert.Empty(t, resultLeases)
	assert.Zero(t, resultPaginationToken)
	assert.Zero(t, resultCount)
	assert.NotNil(t, resultErr)
	_, ok := resultErr.(*shared.InternalServerError)
	assert.True(t, ok)
}
