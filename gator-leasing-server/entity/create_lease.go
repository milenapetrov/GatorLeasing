package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateLease struct {
	Name          string
	OwnerID       uint
	Address       Address
	StartDate     time.Time
	EndDate       time.Time
	Rent          decimal.Decimal
	Utilities     decimal.Decimal
	ParkingCost   decimal.Decimal
	SquareFootage uint
	Furnished     bool
	Parking       bool
	Beds          uint
	Baths         decimal.Decimal
	Amenities     string
	Appliances    string
	Description   string
}
