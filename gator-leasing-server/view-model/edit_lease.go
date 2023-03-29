package viewModel

import (
	"time"

	"github.com/shopspring/decimal"
)

type EditLease struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Address       Address         `json:"address"`
	StartDate     time.Time       `json:"startDate"`
	EndDate       time.Time       `json:"endDate"`
	Rent          decimal.Decimal `json:"rent"`
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
