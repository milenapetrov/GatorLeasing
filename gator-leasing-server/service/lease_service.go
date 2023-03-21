package service

import (
	"net/http"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/dto"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/mapper"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/repository"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/shared"
)

//go:generate mockery --name ILeaseService
type ILeaseService interface {
	GetAllLeases() ([]*entity.Lease, error, int)
	CreateLease(leaseToCreate *entity.CreateLease) (uint, error, int)
	EditLease(leaseToEdit *entity.EditLease) (error, int)
	DeleteLease(id uint) (error, int)
	GetPaginatedLeases(paginatedLeasesRequest *entity.PaginatedLeasesRequest) ([]*entity.Lease, string, int64, error, int)
}

type LeaseService struct {
	userContext *shared.UserContext
	repository  repository.ILeaseRepository
}

func NewLeaseService(userContext *shared.UserContext, repository repository.ILeaseRepository) ILeaseService {
	return &LeaseService{
		userContext: userContext,
		repository:  repository,
	}
}

func (s *LeaseService) GetAllLeases() ([]*entity.Lease, error, int) {
	leaseDtos, err := s.repository.GetAllLeases()
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	mapper := mapper.NewMapper(&dto.Lease{}, &entity.Lease{})
	leaseEntities, err := mapper.MapSlice(leaseDtos)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return leaseEntities, nil, http.StatusOK
}

func (s *LeaseService) CreateLease(leaseToCreate *entity.CreateLease) (uint, error, int) {
	leaseToCreate.OwnerID = s.userContext.ID

	mapper := mapper.NewMapper(&entity.CreateLease{}, &dto.Lease{})
	lease, err := mapper.Map(leaseToCreate)
	if err != nil {
		return 0, err, http.StatusInternalServerError
	}

	id, err := s.repository.CreateLease(lease)
	if err != nil {
		return 0, err, http.StatusInternalServerError
	}

	return id, nil, http.StatusCreated
}

func (s *LeaseService) EditLease(leaseToEdit *entity.EditLease) (error, int) {
	mapper := mapper.NewMapper(&entity.EditLease{}, &dto.Lease{})
	lease, err := mapper.Map(leaseToEdit)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = s.repository.EditLease(lease)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusNoContent
}

func (s *LeaseService) DeleteLease(id uint) (error, int) {
	lease := &dto.Lease{ID: id}

	err := s.repository.DeleteLease(lease)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusNoContent
}

func (s *LeaseService) GetPaginatedLeases(paginatedLeasesRequest *entity.PaginatedLeasesRequest) ([]*entity.Lease, string, int64, error, int) {
	leaseDtos, paginationToken, count, err := s.repository.GetPaginatedLeases(paginatedLeasesRequest.PageSize, paginatedLeasesRequest.SortToken, paginatedLeasesRequest.PaginationToken, paginatedLeasesRequest.SortDirection)
	if err != nil {
		return nil, "", 0, err, http.StatusInternalServerError
	}

	mapper := mapper.NewMapper(&dto.Lease{}, &entity.Lease{})
	leaseEntities, err := mapper.MapSlice(leaseDtos)
	if err != nil {
		return nil, "", 0, err, http.StatusInternalServerError
	}

	return leaseEntities, paginationToken, count, nil, http.StatusOK
}
