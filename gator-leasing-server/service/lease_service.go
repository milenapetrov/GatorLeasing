package service

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/dto"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/mapper"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/repository"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/shared"
)

//go:generate mockery --name ILeaseService
type ILeaseService interface {
	GetAllLeases() ([]*entity.Lease, error)
	CreateLease(leaseToCreate *entity.CreateLease) (uint, error)
	EditLease(editLease *entity.EditLease) error
	DeleteLease(id uint) error
	GetPaginatedLeases(paginatedLeasesRequest *entity.PaginatedLeasesRequest) ([]*entity.Lease, string, int64, error)
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

func (s *LeaseService) GetAllLeases() ([]*entity.Lease, error) {
	leaseDtos, err := s.repository.GetAllLeases()
	if err != nil {
		return nil, err
	}

	mapper := mapper.NewMapper(&dto.Lease{}, &entity.Lease{})
	leaseEntities, err := mapper.MapSlice(leaseDtos)
	if err != nil {
		return nil, err
	}

	return leaseEntities, nil
}

func (s *LeaseService) CreateLease(leaseToCreate *entity.CreateLease) (uint, error) {
	leaseToCreate.OwnerID = s.userContext.ID

	mapper := mapper.NewMapper(&entity.CreateLease{}, &dto.Lease{})
	lease, err := mapper.Map(leaseToCreate)
	if err != nil {
		return 0, err
	}

	id, err := s.repository.CreateLease(lease)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *LeaseService) EditLease(editLease *entity.EditLease) error {
	if s.userContext.InvitedAs != enums.Administrator {
		lease, err := s.repository.GetLeaseById(editLease.ID)
		if err != nil {
			return err
		}
		if s.userContext.ID != lease.OwnerID {
			return &shared.BadRequestError{Msg: "you may only edit leases that belong to you"}
		}
	}

	mapper := mapper.NewMapper(&entity.EditLease{}, &dto.Lease{})
	lease, err := mapper.Map(editLease)
	if err != nil {
		return err
	}

	err = s.repository.EditLease(lease)
	return err
}

func (s *LeaseService) DeleteLease(id uint) error {
	lease := &dto.Lease{ID: id}
	if s.userContext.InvitedAs != enums.Administrator {
		lease, err := s.repository.GetLeaseById(id)
		if err != nil {
			return err
		}
		if s.userContext.ID != lease.OwnerID {
			return &shared.BadRequestError{Msg: "you may only delete leases that belong to you"}
		}
	}

	err := s.repository.DeleteLease(lease)
	return err
}

func (s *LeaseService) GetPaginatedLeases(paginatedLeasesRequest *entity.PaginatedLeasesRequest) ([]*entity.Lease, string, int64, error) {
	paginatedLeasesRequest.SortToken = strcase.ToSnake(paginatedLeasesRequest.SortToken)

	if paginatedLeasesRequest.Filter != "" {
		r := regexp.MustCompile("([a-zA-z_.]+) ([a-zA-z<>=]+) (.+)")
		match := r.FindStringSubmatch(paginatedLeasesRequest.Filter)
		if match == nil {
			return nil, "", 0, &shared.BadRequestError{Msg: "invalid filter"}
		}
		paginatedLeasesRequest.Filter = strcase.ToSnake(match[1]) + " " + strings.ToUpper(match[2]) + " " + match[3]
	}

	leaseDtos, paginationToken, count, err := s.repository.GetPaginatedLeases(paginatedLeasesRequest.PageSize, paginatedLeasesRequest.SortToken, paginatedLeasesRequest.PaginationToken, paginatedLeasesRequest.SortDirection, paginatedLeasesRequest.Filter)
	if err != nil {
		return nil, "", 0, err
	}

	mapper := mapper.NewMapper(&dto.Lease{}, &entity.Lease{})
	leaseEntities, err := mapper.MapSlice(leaseDtos)
	if err != nil {
		return nil, "", 0, err
	}

	return leaseEntities, paginationToken, count, nil
}
