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
	r.DB.Find(&leases)
	return leases, nil
}

func (r *LeaseRepository) CreateLease(lease *model.Lease) (uint, error) {
	result := r.DB.Create(lease)
	return lease.ID, result.Error
}
