package dto

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Lease struct {
	ID            uint            `gorm:"primarykey" faker:"-"`
	CreatedAt     time.Time       `faker:"-"`
	UpdatedAt     time.Time       `faker:"-"`
	DeletedAt     gorm.DeletedAt  `gorm:"index" faker:"-"`
	Name          string          `gorm:"type:varchar(30); not null" faker:"nameFaker"`
	OwnerID       uint            `gorm:"not null" faker:"-"`
	Owner         TenantUser      `gorm:"foreignKey:OwnerID" faker:"-"`
	Address       Address         `gorm:"polymorphic:Owner"`
	StartDate     time.Time       `gorm:"not null" faker:"startDateFaker"`
	EndDate       time.Time       `gorm:"not null" faker:"endDateFaker"`
	Rent          decimal.Decimal `gorm:"type:decimal(10,2);not null" faker:"rentFaker"`
	Utilities     decimal.Decimal `gorm:"type:decimal(10,2)" faker:"utilitiesFaker"`
	ParkingCost   decimal.Decimal `gorm:"type:decimal(10,2)" faker:"parkingCostFaker"`
	TotalCost     decimal.Decimal `gorm:"->;type:decimal(10,2) GENERATED ALWAYS AS (rent + utilities + parking_cost);not null"`
	SquareFootage uint            `gorm:"not null" faker:"boundary_start=200,boundary_end=5000"`
	Furnished     bool            `gorm:"not null"`
	Parking       bool            `gorm:"not null" faker:"parkingFaker"`
	Beds          uint            `gorm:"not null" faker:"boundary_start=1,boundary_end=8"`
	Baths         decimal.Decimal `gorm:"type:decimal(10,1)" faker:"bathsFaker"`
	Amenities     string          `faker:"amenitiesFaker"`
	Appliances    string          `faker:"appliancesFaker"`
	Description   string          `faker:"descriptionFaker"`
	Contacts      []Contact       `faker:"-"`
}
