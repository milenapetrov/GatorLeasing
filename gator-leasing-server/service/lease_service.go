package service

import (
	"GatorLeasing/gator-leasing-server/model"
	"GatorLeasing/gator-leasing-server/repository"
)

type LeaseService struct {
	repository *repository.LeaseRepository
}

func NewLeaseService(repository *repository.LeaseRepository) *LeaseService {
	return &LeaseService{repository: repository}
}

func (s *LeaseService) GetAllLeases() (*[]model.Lease, error) {
	leases, err := s.repository.GetAllLeases()
	return leases, err
}
