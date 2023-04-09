package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/repository/mocks"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/shared"
)

func TestGetOrCreateUserOK(t *testing.T) {
	mockTenantUserRepository := mocks.NewITenantUserRepository(t)
	mockTenantUserRepository.On("GetTenantUserByUserId", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(nil, nil)
	mockTenantUserRepository.On("CreateTenantUser", mock.AnythingOfType("*dto.TenantUser")).Return(uint(1), nil)

	tenantUserService := NewTenantUserService(shared.NewUserContext(), mockTenantUserRepository)

	resultTenantUser, resultErr := tenantUserService.GetOrCreateUser()

	mockTenantUserRepository.AssertExpectations(t)
	mockTenantUserRepository.AssertNumberOfCalls(t, "GetTenantUserByUserId", 1)
	mockTenantUserRepository.AssertNumberOfCalls(t, "CreateTenantUser", 1)
	assert.NotNil(t, resultTenantUser)
	assert.Nil(t, resultErr)
}

func TestGetOrCreateUserGetTenantUserErr(t *testing.T) {
	mockTenantUserRepository := mocks.NewITenantUserRepository(t)
	mockTenantUserRepository.On("GetTenantUserByUserId", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(nil, &shared.InternalServerError{})

	tenantUserService := NewTenantUserService(shared.NewUserContext(), mockTenantUserRepository)

	resultTenantUser, resultErr := tenantUserService.GetOrCreateUser()

	mockTenantUserRepository.AssertExpectations(t)
	mockTenantUserRepository.AssertNumberOfCalls(t, "GetTenantUserByUserId", 1)
	mockTenantUserRepository.AssertNotCalled(t, "CreateTenantUser")
	assert.Nil(t, resultTenantUser)
	assert.NotNil(t, resultErr)
	_, ok := resultErr.(*shared.InternalServerError)
	assert.True(t, ok)
}

func TestGetOrCreateUserCreateTenantUserErr(t *testing.T) {
	mockTenantUserRepository := mocks.NewITenantUserRepository(t)
	mockTenantUserRepository.On("GetTenantUserByUserId", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(nil, nil)
	mockTenantUserRepository.On("CreateTenantUser", mock.AnythingOfType("*dto.TenantUser")).Return(uint(0), &shared.InternalServerError{})

	tenantUserService := NewTenantUserService(shared.NewUserContext(), mockTenantUserRepository)

	resultTenantUser, resultErr := tenantUserService.GetOrCreateUser()

	mockTenantUserRepository.AssertExpectations(t)
	mockTenantUserRepository.AssertNumberOfCalls(t, "GetTenantUserByUserId", 1)
	mockTenantUserRepository.AssertNumberOfCalls(t, "CreateTenantUser", 1)
	assert.Nil(t, resultTenantUser)
	assert.NotNil(t, resultErr)
	_, ok := resultErr.(*shared.InternalServerError)
	assert.True(t, ok)
}
