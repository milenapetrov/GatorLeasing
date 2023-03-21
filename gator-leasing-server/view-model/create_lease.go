package viewModel

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateLease struct {
	Name          string          `json:"name" validate:"required"`
	Address       Address         `json:"address"`
	StartDate     time.Time       `json:"startDate"`
	EndDate       time.Time       `json:"endDate"`
	Rent          decimal.Decimal `json:"rent"`
	Utilities     decimal.Decimal `json:"utilities"`
	ParkingCost   decimal.Decimal `json:"parkingCost"`
	SquareFootage uint            `json:"squareFootage"`
	Furnished     bool            `json:"furnished"`
	Parking       bool            `json:"parking"`
	Beds          uint            `json:"beds"`
	Baths         decimal.Decimal `json:"baths"`
	Amenities     string          `json:"amenities"`
	Appliances    string          `json:"appliances"`
	Description   string          `json:"description"`
}
