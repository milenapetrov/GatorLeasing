package viewModel

import (
	"time"

	"github.com/shopspring/decimal"
)

type Lease struct {
	ID            uint            `json:"id"`
	Name          string          `json:"name"`
	OwnerID       uint            `json:"ownerID"`
	Address       Address         `json:"address"`
	StartDate     time.Time       `json:"startDate"`
	EndDate       time.Time       `json:"endDate"`
	Rent          decimal.Decimal `json:"rent"`
	Utilities     decimal.Decimal `json:"utilities"`
	ParkingCost   decimal.Decimal `json:"parkingCost"`
	TotalCost     decimal.Decimal `json:"totalCost"`
	SquareFootage uint            `json:"squareFootage"`
	Furnished     bool            `json:"furnished"`
	Parking       bool            `json:"parking"`
	Beds          uint            `json:"beds"`
	Baths         decimal.Decimal `json:"baths"`
	Amenities     string          `json:"amenities"`
	Appliances    string          `json:"appliances"`
	Description   string          `json:"description"`
	Contacts      []Contact       `json:"contacts"`
}
