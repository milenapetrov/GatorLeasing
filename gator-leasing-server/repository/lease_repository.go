package repository

import (
	stdErrors "errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/dto"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/errors"

	"gorm.io/gorm"
)

//go:generate mockery --name ILeaseRepository
type ILeaseRepository interface {
	GetLeasesByOwnerId(id uint) ([]*dto.Lease, error)
	GetAllLeases() ([]*dto.Lease, error)
	GetLeaseById(id uint) (*dto.Lease, error)
	CreateLease(lease *dto.Lease) (uint, error)
	EditLease(lease *dto.Lease) error
	DeleteLease(lease *dto.Lease) error
	GetPaginatedLeases(pageSize uint, sortToken string, paginationToken string, sortDirection enums.SortDirection, filters string) ([]*dto.Lease, string, int64, error)
}

type LeaseRepository struct {
	DB *gorm.DB
}

func NewLeaseRepository(db *gorm.DB) ILeaseRepository {
	return &LeaseRepository{DB: db}
}

func (r *LeaseRepository) GetAllLeases() ([]*dto.Lease, error) {
	leases := []*dto.Lease{}
	if err := r.DB.Preload("Address").Find(&leases).Error; err != nil {
		return nil, &errors.InternalServerError{Msg: err.Error()}
	}
	return leases, nil
}

func (r *LeaseRepository) GetLeaseById(id uint) (*dto.Lease, error) {
	lease := &dto.Lease{ID: id}
	if err := r.DB.Preload("Address").First(lease).Error; err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.BadRequestError{Msg: fmt.Sprintf("no lease with ID %d", id)}
		}
		return nil, &errors.InternalServerError{Msg: err.Error()}
	}

	return lease, nil
}

func (r *LeaseRepository) CreateLease(lease *dto.Lease) (uint, error) {
	if err := r.DB.Create(lease).Error; err != nil {
		return 0, &errors.InternalServerError{Msg: err.Error()}
	}
	return lease.ID, nil
}

func (r *LeaseRepository) EditLease(lease *dto.Lease) error {
	oldLease := &dto.Lease{ID: lease.ID}
	if err := r.DB.Preload("Address").First(oldLease).Error; err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return &errors.BadRequestError{Msg: fmt.Sprintf("no lease with ID %d", lease.ID)}
		}
		return &errors.InternalServerError{Msg: err.Error()}
	}

	oldLease.Name = lease.Name
	oldLease.Address = lease.Address
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

	if err := r.DB.Save(oldLease).Error; err != nil {
		return &errors.InternalServerError{Msg: err.Error()}
	}
	return nil
}

func (r *LeaseRepository) DeleteLease(lease *dto.Lease) error {
	if err := r.DB.First(lease).Error; err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return &errors.BadRequestError{Msg: fmt.Sprintf("no lease with ID %d", lease.ID)}
		}
		return &errors.InternalServerError{Msg: err.Error()}
	}

	if err := r.DB.Delete(lease).Error; err != nil {
		return &errors.InternalServerError{Msg: err.Error()}
	}
	return nil
}

func (r *LeaseRepository) GetPaginatedLeases(pageSize uint, sortToken string, paginationToken string, sortDirection enums.SortDirection, filters string) ([]*dto.Lease, string, int64, error) {
	if sortToken == "" {
		sortToken = "created_at"
	}

	leases := []*dto.Lease{}

	query := r.DB.Joins("Address")

	if sortToken == "created_at" {
		if sortDirection == enums.Ascending {
			query.Order("leases.created_at")
		} else {
			query.Order("leases.created_at desc")
		}
	} else {
		if sortDirection == enums.Ascending {
			query.Order(sortToken + ", leases.created_at desc")
		} else {
			query.Order(sortToken + " desc, leases.created_at desc")
		}
	}

	if paginationToken != "" {
		if sortToken == "created_at" {
			if sortDirection == enums.Ascending {
				query.Where("leases.created_at >= ?", paginationToken)
			} else {
				query.Where("leases.created_at <= ?", paginationToken)
			}
		} else {
			splitToken := strings.Split(paginationToken, "|")
			paginationValue := splitToken[0]
			paginationDate := splitToken[1]
			namedParams := map[string]interface{}{"paginationValue": paginationValue, "paginationDate": paginationDate}
			if sortDirection == enums.Ascending {
				query.Where(sortToken+" > @paginationValue OR "+sortToken+" = @paginationValue AND leases.created_at <= @paginationDate", namedParams)
			} else {
				query.Where(sortToken+" < @paginationValue OR "+sortToken+" = @paginationValue AND leases.created_at <= @paginationDate", namedParams)
			}
		}
	}

	query.Limit(int(pageSize) + 1)

	if filters != "" {
		filtersSplit := strings.Split(filters, ",")
		for _, filter := range filtersSplit {
			query.Where(filter)
		}
	}

	err := query.Find(&leases).Error
	if err != nil {
		return nil, "", 0, &errors.BadRequestError{Msg: err.Error()}
	}

	newPaginationToken := ""

	if len(leases) > int(pageSize) {
		lastLease := leases[len(leases)-1]
		leases = leases[:len(leases)-1]
		newPaginationToken = getPaginationToken(lastLease, sortToken)
	}

	var count int64
	query = r.DB.Model(&dto.Lease{})

	if filters != "" {
		filtersSplit := strings.Split(filters, ",")
		for _, filter := range filtersSplit {
			query.Where(filter)
		}
	}

	err = query.Count(&count).Error
	if err != nil {
		return nil, "", 0, &errors.InternalServerError{Msg: err.Error()}
	}

	return leases, newPaginationToken, count, nil
}

func getPaginationToken(lastLease *dto.Lease, sortToken string) string {
	paginationToken := ""
	switch sortToken {
	case "name":
		paginationToken = lastLease.Name + "|"

	case "start_date":
		paginationToken = lastLease.StartDate.String() + "|"

	case "end_date":
		paginationToken = lastLease.EndDate.String() + "|"

	case "rent":
		paginationToken = lastLease.Rent.String() + "|"

	case "utilities":
		paginationToken = lastLease.Utilities.String() + "|"

	case "parking_cost":
		paginationToken = lastLease.ParkingCost.String() + "|"

	case "total_cost":
		paginationToken = lastLease.TotalCost.String() + "|"

	case "beds":
		paginationToken = strconv.FormatUint(uint64(lastLease.Beds), 10) + "|"

	case "baths":
		paginationToken = lastLease.Baths.String() + "|"
	}

	paginationToken += lastLease.CreatedAt.String()

	return paginationToken
}

func (r *LeaseRepository) GetLeasesByOwnerId(id uint) ([]*dto.Lease, error) {
	leases := []*dto.Lease{}
	if err := r.DB.Preload("Address").Where("owner_id = ?", id).Find(&leases).Error; err != nil {
		return nil, &errors.InternalServerError{Msg: err.Error()}
	}
	return leases, nil
}
