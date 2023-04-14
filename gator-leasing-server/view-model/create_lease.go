package viewModel

import (
	"time"

	"github.com/shopspring/decimal"
)

// swagger:parameters PostLease
type CreateLease struct {
	// the name for the new lease
	// required: true
	// in: body
	// min: 3
	// max: 20
	Name string `json:"name" validate:"required,min=3,max=20" faker:"len=10"`
	// the address for the new lease
	// required: true
	// in:body
	Address Address `json:"address" validate:"required"`
	// the start date for the new lease
	// required: true
	// in:body
	StartDate time.Time `json:"startDate" validate:"required" faker:"createLeaseStartDateFaker"`
	// the end date for the new lease
	// required: true
	// in:body
	EndDate time.Time `json:"endDate" validate:"required,gtfield=StartDate" faker:"createLeaseEndDateFaker"`
	// the rent cost for the new lease
	// required: true
	// in:body
	// min: 0.01
	Rent decimal.Decimal `json:"rent" validate:"required,dmin=0.01" faker:"createLeaseRentFaker"`
	// the utilites cost for the new lease
	// in:body
	Utilities decimal.Decimal `json:"utilities"`
	// the parking cost for the new lease
	// in:body
	ParkingCost decimal.Decimal `json:"parkingCost"`
	// the square footage for the new lease
	// in:body
	SquareFootage int `json:"squareFootage"`
	// the furnished data for the new lease
	// in:body
	Furnished bool `json:"furnished"`
	// the parking data for the new lease
	// in:body
	Parking bool `json:"parking"`
	// the bedroom info for the new lease
	// in:body
	Beds int `json:"beds"`
	// the bathroom info for the new lease
	// in:body
	Baths decimal.Decimal `json:"baths"`
	// the amenities info for the new lease
	// in:body
	Amenities string `json:"amenities"`
	// the appliances info for the new lease
	// in:body
	Appliances string `json:"appliances"`
	// the description for the new lease
	// in:body
	Description string `json:"description"`
}
