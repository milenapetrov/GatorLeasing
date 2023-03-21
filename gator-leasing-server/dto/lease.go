package dto

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Lease struct {
	ID            uint `gorm:"primarykey" faker:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt  `gorm:"index"`
	Name          string          `gorm:"type:varchar(30); not null"`
	OwnerID       uint            `gorm:"not null"`
	Owner         TenantUser      `gorm:"foreignKey:OwnerID"`
	Address       Address         `gorm:"polymorphic:Owner"`
	StartDate     time.Time       `gorm:"not null"`
	EndDate       time.Time       `gorm:"not null"`
	Rent          decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	Utilities     decimal.Decimal `gorm:"type:decimal(10,2)"`
	ParkingCost   decimal.Decimal `gorm:"type:decimal(10,2)"`
	TotalCost     decimal.Decimal `gorm:"->;type:decimal(10,2) GENERATED ALWAYS AS (rent + utilities + parking_cost);not null"`
	SquareFootage uint            `gorm:"not null"`
	Furnished     bool            `gorm:"not null"`
	Parking       bool            `gorm:"not null"`
	Beds          uint            `gorm:"not null"`
	Baths         decimal.Decimal `gorm:"type:decimal(10,1)"`
	Amenities     string
	Appliances    string
	Description   string
	Contacts      []Contact
}
