package repository

import "gorm.io/gorm"

type LeaseRepository struct {
	DB *gorm.DB
}

func NewLeaseRepository(db *gorm.DB) *LeaseRepository {
	return &LeaseRepository{DB: db}
}
