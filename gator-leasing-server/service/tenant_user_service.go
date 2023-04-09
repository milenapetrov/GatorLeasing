package service

import (
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/constants"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/dto"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/mapper"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/repository"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/shared"
)

type ITenantUserService interface {
	GetOrCreateUser() (*entity.TenantUser, error)
}

type TenantUserService struct {
	userContext *shared.UserContext
	repository  repository.ITenantUserRepository
}

func NewTenantUserService(userContext *shared.UserContext, repository repository.ITenantUserRepository) ITenantUserService {
	return &TenantUserService{
		userContext: userContext,
		repository:  repository,
	}
}

func (s *TenantUserService) GetOrCreateUser() (*entity.TenantUser, error) {
	tenantUserDto, err := s.repository.GetTenantUserByUserId(s.userContext.UserID, constants.TENANT_ID)
	if err != nil {
		return nil, err
	}
	if tenantUserDto == nil {
		tenantUserDto = &dto.TenantUser{
			UserID:    s.userContext.UserID,
			TenantID:  constants.TENANT_ID,
			InvitedAs: enums.Member,
		}
		tenantUserDto.ID, err = s.repository.CreateTenantUser(tenantUserDto)
		if err != nil {
			return nil, err
		}
	}
	mapper := mapper.NewMapper(&dto.TenantUser{}, &entity.TenantUser{})
	tenantUser, err := mapper.Map(tenantUserDto)
	if err != nil {
		return nil, err
	}
	return tenantUser, nil
}
