package viewModel

import (
	"time"

	"github.com/shopspring/decimal"
)

// swagger:parameters PutLease
type EditLease struct {
	// the id for the lease to update
	// required: true
	// in: query
	ID int `json:"id"`
	// the updated name for the lease
	// in: body
	Name string `json:"name"`
	// the updated address for the lease
	// in: body
	Address Address `json:"address"`
	// the updated start date for the lease
	// in: body
	StartDate time.Time `json:"startDate"`
	// the updated end date for the lease
	// in: body
	EndDate time.Time `json:"endDate"`
	// the updated rent cost for the lease
	// in: body
	Rent decimal.Decimal `json:"rent"`
	// the updated utilities cost for the lease
	// in: body
	Utilities decimal.Decimal `json:"utilities"`
	// the updated parking cost for the lease
	// in: body
	ParkingCost decimal.Decimal `json:"parkingCost"`
	// the updated square footage for the lease
	// in: body
	SquareFootage int `json:"squareFootage"`
	// the updated furnished data for the lease
	// in: body
	Furnished bool `json:"furnished"`
	// the updated parking data for the lease
	// in: body
	Parking bool `json:"parking"`
	// the updated bedroom data for the lease
	// in: body
	Beds int `json:"beds"`
	// the updated bathroom data for the lease
	// in: body
	Baths decimal.Decimal `json:"baths"`
	// the updated amenities info for the lease
	// in: body
	Amenities string `json:"amenities"`
	// the updated appliances info for the lease
	// in: body
	Appliances string `json:"appliances"`
	// the updated description for the lease
	// in: body
	Description string `json:"description"`
}
