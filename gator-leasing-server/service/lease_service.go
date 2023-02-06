package service

import (
	"GatorLeasing/gator-leasing-server/entity"
	"GatorLeasing/gator-leasing-server/model"
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
	if err != nil {
		return nil, err
	}

	var leaseEntities []*entity.Lease

	for _, l := range leases {
		leaseEntities = append(leaseEntities, &entity.Lease{ID: l.ID, Name: l.Name})
	}

	return leaseEntities, nil
}

func (s *LeaseService) CreateLease(request entity.CreateLeaseRequest) (uint, error) {
	lease := &model.Lease{Name: request.Name}
	return s.repository.CreateLease(lease)
}

func (s *LeaseService) EditLease(request entity.EditLeaseRequest) error {
	lease := &model.Lease{Name: request.Name}
	return s.repository.EditLease(request.Id, lease)
}
