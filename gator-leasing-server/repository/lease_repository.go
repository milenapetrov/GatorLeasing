package repository

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/dto"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/shared"

	"gorm.io/gorm"
)

//go:generate mockery --name ILeaseRepository
type ILeaseRepository interface {
	GetLeaseById(id uint) (*dto.Lease, error)
	GetAllLeases() ([]*dto.Lease, error)
	CreateLease(lease *dto.Lease) (uint, error)
	EditLease(lease *dto.Lease) error
	DeleteLease(lease *dto.Lease) error
	GetPaginatedLeases(pageSize uint, sortToken string, paginationToken string, sortDirection enums.SortDirection) ([]*dto.Lease, string, int64, error)
}

type LeaseRepository struct {
	DB *gorm.DB
}

func NewLeaseRepository(db *gorm.DB) ILeaseRepository {
	return &LeaseRepository{DB: db}
}

func (r *LeaseRepository) GetLeaseById(id uint) (*dto.Lease, error) {
	lease := &dto.Lease{ID: id}
	err := r.DB.Preload("Address").First(lease).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &shared.BadRequestError{Msg: fmt.Sprintf("no lease with ID %d", id)}
		}
		return nil, &shared.InternalServerError{Msg: err.Error()}
	}

	return lease, nil
}

func (r *LeaseRepository) GetAllLeases() ([]*dto.Lease, error) {
	leases := []*dto.Lease{}
	err := r.DB.Preload("Address").Find(&leases).Error
	if err != nil {
		return nil, &shared.InternalServerError{Msg: err.Error()}
	}
	return leases, nil
}

func (r *LeaseRepository) CreateLease(lease *dto.Lease) (uint, error) {
	err := r.DB.Create(lease).Error
	if err != nil {
		return 0, &shared.InternalServerError{Msg: err.Error()}
	}
	return lease.ID, nil
}

func (r *LeaseRepository) EditLease(lease *dto.Lease) error {
	oldLease := &dto.Lease{ID: lease.ID}
	err := r.DB.Preload("Address").First(oldLease).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &shared.BadRequestError{Msg: fmt.Sprintf("no lease with ID %d", lease.ID)}
		}
		return &shared.InternalServerError{Msg: err.Error()}
	}

	oldLease.Name = lease.Name
	oldLease.StartDate = lease.StartDate
	oldLease.EndDate = lease.EndDate
	oldLease.Rent = lease.Rent
	oldLease.Utilities = lease.Utilities
	oldLease.ParkingCost = lease.ParkingCost
	oldLease.SquareFootage = lease.SquareFootage
	oldLease.Furnished = lease.Furnished
	oldLease.Parking = lease.Parking
	oldLease.Beds = lease.Beds
	oldLease.Baths = lease.Baths
	oldLease.Amenities = lease.Amenities
	oldLease.Appliances = lease.Appliances
	oldLease.Description = lease.Description

	err = r.DB.Save(oldLease).Error
	if err != nil {
		return &shared.InternalServerError{Msg: err.Error()}
	}
	return nil
}

func (r *LeaseRepository) DeleteLease(lease *dto.Lease) error {
	err := r.DB.First(lease).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &shared.BadRequestError{Msg: fmt.Sprintf("no lease with ID %d", lease.ID)}
		}
		return &shared.InternalServerError{Msg: err.Error()}
	}

	err = r.DB.Delete(lease).Error
	if err != nil {
		return &shared.InternalServerError{Msg: err.Error()}
	}
	return nil
}

func (r *LeaseRepository) GetPaginatedLeases(pageSize uint, sortToken string, paginationToken string, sortDirection enums.SortDirection) ([]*dto.Lease, string, int64, error) {
	if sortToken == "" {
		sortToken = "created_at"
	}
	sortToken = strings.ToLower(sortToken)

	leases := []*dto.Lease{}

	query := r.DB.Preload("Address")

	if sortToken == "created_at" {
		if sortDirection == enums.Ascending {
			query.Order("created_at")
		} else {
			query.Order("created_at desc")
		}
	} else {
		if sortDirection == enums.Ascending {
			query.Order(sortToken + ", created_at desc")
		} else {
			query.Order(sortToken + " desc, created_at desc")
		}
	}

	if paginationToken != "" {
		if sortToken == "created_at" {
			if sortDirection == enums.Ascending {
				query.Where("created_at >= ?", paginationToken)
			} else {
				query.Where("created_at <= ?", paginationToken)
			}
		} else {
			splitToken := strings.Split(paginationToken, "|")
			paginationValue := splitToken[0]
			paginationDate := splitToken[1]
			namedParams := map[string]interface{}{"paginationValue": paginationValue, "paginationDate": paginationDate}
			if sortDirection == enums.Ascending {
				query.Where(sortToken+" > @paginationValue OR "+sortToken+" = @paginationValue AND created_at <= @paginationDate", namedParams)
			} else {
				query.Where(sortToken+" < @paginationValue OR "+sortToken+" = @paginationValue AND created_at <= @paginationDate", namedParams)
			}
		}
	}

	query.Limit(int(pageSize) + 1)

	err := query.Find(&leases).Error
	if err != nil {
		return nil, "", 0, &shared.InternalServerError{Msg: err.Error()}
	}

	newPaginationToken := ""

	if len(leases) > int(pageSize) {
		lastLease := leases[len(leases)-1]
		newPaginationToken = getPaginationToken(lastLease, sortToken)
	}

	var count int64
	err = r.DB.Model(&dto.Lease{}).Count(&count).Error
	if err != nil {
		return nil, "", 0, &shared.InternalServerError{Msg: err.Error()}
	}

	return leases[:len(leases)-1], newPaginationToken, count, nil
}

func getPaginationToken(lastLease *dto.Lease, sortToken string) string {
	paginationToken := ""
	switch sortToken {
	case "name":
		paginationToken = lastLease.Name + "|"

	case "startDate":
		paginationToken = lastLease.StartDate.String() + "|"

	case "endDate":
		paginationToken = lastLease.EndDate.String() + "|"

	case "rent":
		paginationToken = lastLease.Rent.String() + "|"

	case "utilities":
		paginationToken = lastLease.Utilities.String() + "|"

	case "parkingCost":
		paginationToken = lastLease.ParkingCost.String() + "|"

	case "beds":
		paginationToken = strconv.FormatUint(uint64(lastLease.Beds), 10) + "|"

	case "baths":
		paginationToken = lastLease.Baths.String() + "|"
	}

	paginationToken += lastLease.CreatedAt.String()

	return paginationToken
}
