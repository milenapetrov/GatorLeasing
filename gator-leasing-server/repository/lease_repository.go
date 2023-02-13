package repository

import (
	"GatorLeasing/gator-leasing-server/model"

	"gorm.io/gorm"
)

type LeaseRepository struct {
	DB *gorm.DB
}

func NewLeaseRepository(db *gorm.DB) *LeaseRepository {
	return &LeaseRepository{DB: db}
}

func (r *LeaseRepository) GetAllLeases() ([]model.Lease, error) {
	var leases []model.Lease
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
