package viewModel

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateLease struct {
	Name          string          `json:"name" validate:"required,min=3,max=20" faker:"len=10"`
	Address       Address         `json:"address" validate:"required"`
	StartDate     time.Time       `json:"startDate"`
	EndDate       time.Time       `json:"endDate" validate:"gtfield=StartDate"`
	Rent          decimal.Decimal `json:"rent" validate:"required,dmin=0.01"`
	Utilities     decimal.Decimal `json:"utilities"`
	ParkingCost   decimal.Decimal `json:"parkingCost"`
	SquareFootage int             `json:"squareFootage"`
	Furnished     bool            `json:"furnished"`
	Parking       bool            `json:"parking"`
	Beds          int             `json:"beds"`
	Baths         decimal.Decimal `json:"baths"`
	Amenities     string          `json:"amenities"`
	Appliances    string          `json:"appliances"`
	Description   string          `json:"description"`
}
