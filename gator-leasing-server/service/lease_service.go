package service

import (
	"GatorLeasing/gator-leasing-server/entity"
	"GatorLeasing/gator-leasing-server/repository"
)

type LeaseService struct {
	repository *repository.LeaseRepository
}

func NewLeaseService(repository *repository.LeaseRepository) *LeaseService {
	return &LeaseService{repository: repository}
}

func (s *LeaseService) GetAllLeases() ([]*entity.Lease, error) {
	leases, err := s.repository.GetAllLeases()

	var leaseEntities []*entity.Lease

	for _, l := range leases {
		leaseEntities = append(leaseEntities, entity.NewLease(&l))
	}

	return leaseEntities, err
}
