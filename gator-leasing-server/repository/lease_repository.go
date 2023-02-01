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

func (r *LeaseRepository) GetAllLeases() (*[]model.Lease, error) {
	return &[]model.Lease{}, nil
}
