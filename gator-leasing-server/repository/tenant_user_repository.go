package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/model"
)

//go:generate mockery - name IUserRepository
type ITenantUserRepository interface {
	GetTenantUserByUserID(userId string, tenantId uint) (*model.TenantUser, error)
	CreateTenantUser(tenantUser *model.TenantUser) (uint, error)
}

type TenantUserRepository struct {
	DB *gorm.DB
}

func NewTenantUserRepository(db *gorm.DB) ITenantUserRepository {
	return &TenantUserRepository{DB: db}
}

func (r *TenantUserRepository) GetTenantUserByUserID(userId string, tenantId uint) (*model.TenantUser, error) {
	tenantUser := model.TenantUser{}
	err := r.DB.Where("user_id = ? AND tenant_id = ?", userId, tenantId).First(&tenantUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &tenantUser, nil
}

func (r *TenantUserRepository) CreateTenantUser(tenantUser *model.TenantUser) (uint, error) {
	err := r.DB.Create(tenantUser).Error
	return tenantUser.ID, err
}
