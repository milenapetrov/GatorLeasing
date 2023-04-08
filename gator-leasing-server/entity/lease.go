package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Lease struct {
	ID            uint
	Name          string
	CreatedAt     time.Time
	OwnerID       uint
	Address       Address
	StartDate     time.Time
	EndDate       time.Time
	Rent          decimal.Decimal
	Utilities     decimal.Decimal
	ParkingCost   decimal.Decimal
	TotalCost     decimal.Decimal
	SquareFootage uint
	Furnished     bool
	Parking       bool
	Beds          uint
	Baths         decimal.Decimal
	Amenities     string
	Appliances    string
	Description   string
	Contacts      []Contact `faker:"contactsEntityFaker"`
}
