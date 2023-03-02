package service

import (
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/constants"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/model"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/repository"
)

type ITenantUserService interface {
	GetOrCreateUser() (*entity.TenantUser, error)
}

type TenantUserService struct {
	userContext *entity.UserContext
	repository  repository.ITenantUserRepository
}

func NewTenantUserService(userContext *entity.UserContext, repository repository.ITenantUserRepository) ITenantUserService {
	return &TenantUserService{
		userContext: userContext,
		repository:  repository,
	}
}

func (s *TenantUserService) GetOrCreateUser() (*entity.TenantUser, error) {
	tenantUserModel, err := s.repository.GetTenantUserByUserID(s.userContext.UserID, constants.TENANT_ID)
	if err != nil {
		return nil, err
	}
	if tenantUserModel == nil {
		tenantUserModel = &model.TenantUser{
			UserID:    s.userContext.UserID,
			TenantID:  constants.TENANT_ID,
			InvitedAs: enums.Member,
		}
		tenantUserModel.ID, err = s.repository.CreateTenantUser(tenantUserModel)
		if err != nil {
			return nil, err
		}
	}
	tenantUser := &entity.TenantUser{
		ID:        tenantUserModel.ID,
		UserID:    tenantUserModel.UserID,
		TenantID:  tenantUserModel.TenantID,
		InvitedAs: tenantUserModel.InvitedAs,
	}
	return tenantUser, nil
}
