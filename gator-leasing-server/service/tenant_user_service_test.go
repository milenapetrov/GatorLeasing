package service

import (
	stdErrors "errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/errors"
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
	mockTenantUserRepository.On("GetTenantUserByUserId", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(nil, &errors.InternalServerError{})

	tenantUserService := NewTenantUserService(shared.NewUserContext(), mockTenantUserRepository)

	resultTenantUser, resultErr := tenantUserService.GetOrCreateUser()

	mockTenantUserRepository.AssertExpectations(t)
	mockTenantUserRepository.AssertNumberOfCalls(t, "GetTenantUserByUserId", 1)
	mockTenantUserRepository.AssertNotCalled(t, "CreateTenantUser")
	assert.Nil(t, resultTenantUser)
	assert.NotNil(t, resultErr)
	assert.True(t, stdErrors.As(resultErr, &errors.INTERNAL_SERVER_ERROR))
}

func TestGetOrCreateUserCreateTenantUserErr(t *testing.T) {
	mockTenantUserRepository := mocks.NewITenantUserRepository(t)
	mockTenantUserRepository.On("GetTenantUserByUserId", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(nil, nil)
	mockTenantUserRepository.On("CreateTenantUser", mock.AnythingOfType("*dto.TenantUser")).Return(uint(0), &errors.InternalServerError{})

	tenantUserService := NewTenantUserService(shared.NewUserContext(), mockTenantUserRepository)

	resultTenantUser, resultErr := tenantUserService.GetOrCreateUser()

	mockTenantUserRepository.AssertExpectations(t)
	mockTenantUserRepository.AssertNumberOfCalls(t, "GetTenantUserByUserId", 1)
	mockTenantUserRepository.AssertNumberOfCalls(t, "CreateTenantUser", 1)
	assert.Nil(t, resultTenantUser)
	assert.NotNil(t, resultErr)
	assert.True(t, stdErrors.As(resultErr, &errors.INTERNAL_SERVER_ERROR))
}
