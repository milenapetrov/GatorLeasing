package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/dto"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/shared"
)

//go:generate mockery --name ITenantUserRepository
type ITenantUserRepository interface {
	GetTenantUserByUserId(userId string, tenantId uint) (*dto.TenantUser, error)
	CreateTenantUser(tenantUser *dto.TenantUser) (uint, error)
}

type TenantUserRepository struct {
	DB *gorm.DB
}

func NewTenantUserRepository(db *gorm.DB) ITenantUserRepository {
	return &TenantUserRepository{DB: db}
}

func (r *TenantUserRepository) GetTenantUserByUserId(userId string, tenantId uint) (*dto.TenantUser, error) {
	tenantUser := dto.TenantUser{}
	err := r.DB.Where("user_id = ? AND tenant_id = ?", userId, tenantId).First(&tenantUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, &shared.InternalServerError{Msg: err.Error()}
	}
	return &tenantUser, nil
}

func (r *TenantUserRepository) CreateTenantUser(tenantUser *dto.TenantUser) (uint, error) {
	err := r.DB.Create(tenantUser).Error
	if err != nil {
		return 0, err
	}
	return tenantUser.ID, nil
}
