package service

import (
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/model"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/repository"
)

//go:generate mockery --name ILeaseService
type ILeaseService interface {
	GetAllLeases() ([]*entity.Lease, error)
	CreateLease(request *entity.CreateLeaseRequest) (uint, error)
	EditLease(request *entity.EditLeaseRequest) error
	DeleteLease(request *entity.DeleteLeaseRequest) error
}

type LeaseService struct {
	userContext *entity.UserContext
	repository  repository.ILeaseRepository
}

func NewLeaseService(userContext *entity.UserContext, repository repository.ILeaseRepository) *LeaseService {
	return &LeaseService{
		userContext: userContext,
		repository:  repository,
	}
}

func (s *LeaseService) GetAllLeases() ([]*entity.Lease, error) {
	leaseModels, err := s.repository.GetAllLeases()

	if err != nil {
		return nil, err
	}

	leaseEntities := []*entity.Lease{}

	for _, l := range leaseModels {
		leaseEntities = append(leaseEntities, &entity.Lease{ID: l.ID, Name: l.Name, OwnerID: l.OwnerID})
	}

	return leaseEntities, nil
}

func (s *LeaseService) CreateLease(request *entity.CreateLeaseRequest) (uint, error) {
	lease := &model.Lease{Name: request.Name, OwnerID: s.userContext.ID}
	id, err := s.repository.CreateLease(lease)
	return id, err
}

func (s *LeaseService) EditLease(request *entity.EditLeaseRequest) error {
	lease := &model.Lease{ID: request.ID, Name: request.Name}
	err := s.repository.EditLease(lease)
	return err
}

func (s *LeaseService) DeleteLease(request *entity.DeleteLeaseRequest) error {
	lease := &model.Lease{ID: request.ID}
	err := s.repository.DeleteLease(lease)
	return err
}
