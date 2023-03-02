package repository

import (
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/model"

	"gorm.io/gorm"
)

//go:generate mockery --name ILeaseRepository
type ILeaseRepository interface {
	GetAllLeases() ([]model.Lease, error)
	CreateLease(lease *model.Lease) (uint, error)
	EditLease(lease *model.Lease) error
	DeleteLease(lease *model.Lease) error
}

type LeaseRepository struct {
	DB *gorm.DB
}

func NewLeaseRepository(db *gorm.DB) ILeaseRepository {
	return &LeaseRepository{DB: db}
}

func (r *LeaseRepository) GetAllLeases() ([]model.Lease, error) {
	leases := []model.Lease{}
	err := r.DB.Find(&leases).Error
	return leases, err
}

func (r *LeaseRepository) CreateLease(lease *model.Lease) (uint, error) {
	err := r.DB.Create(lease).Error
	return lease.ID, err
}

func (r *LeaseRepository) EditLease(lease *model.Lease) error {
	oldLease := model.Lease{ID: lease.ID}
	err := r.DB.First(&oldLease).Error
	if err != nil {
		return err
	}

	oldLease.Name = lease.Name
	oldLease.OwnerID = lease.OwnerID

	err = r.DB.Save(&oldLease).Error
	return err
}

func (r *LeaseRepository) DeleteLease(lease *model.Lease) error {
	err := r.DB.First(&lease).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(lease).Error
	return err
}
