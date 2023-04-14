package viewModel

import (
	"time"

	"github.com/shopspring/decimal"
)

// swagger:model Lease
type Lease struct {
	// id for this lease
	ID int `json:"id"`
	// name for this lese
	Name string `json:"name"`
	// created at time for this lease
	CreatedAt time.Time `json:"createdAt"`
	// owner id for this lease
	OwnerID int `json:"ownerID"`
	// address for this lease
	Address Address `json:"address"`
	// start date for this lease
	StartDate time.Time `json:"startDate"`
	// end date for this lease
	EndDate time.Time `json:"endDate"`
	// rent for this lease
	Rent decimal.Decimal `json:"rent"`
	// utilities for this lease
	Utilities decimal.Decimal `json:"utilities"`
	// parking cost for this lease
	ParkingCost decimal.Decimal `json:"parkingCost"`
	// total cost for this lease
	TotalCost decimal.Decimal `json:"totalCost"`
	// square footage for this lease
	SquareFootage int `json:"squareFootage"`
	// furnished info for this lease
	Furnished bool `json:"furnished"`
	// parking info for this lease
	Parking bool `json:"parking"`
	// bedroom info for this lease
	Beds int `json:"beds"`
	// bathroom info for this lease
	Baths decimal.Decimal `json:"baths"`
	// amenities info for this lease
	Amenities string `json:"amenities"`
	// appliances info for this lease
	Appliances string `json:"appliances"`
	// description for this lease
	Description string `json:"description"`
	// contacts array for this lease
	Contacts []Contact `json:"contacts"`
}
