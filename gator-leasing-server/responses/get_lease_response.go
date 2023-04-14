package responses

import (
	"time"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/shopspring/decimal"
)

// The response that is returned when the information of a lease is requested
// swagger:response GetLeaseResponse
type GetLeaseResponse struct {
	// in: body
	Body struct {
		// the lease ID
		// required: true
		ID uint

		// the lease name
		// required: true
		Name string

		// the lease owner id
		// required: true
		OwnerID uint

		// the lease address
		// required: true
		Address entity.Address

		// the lease start date
		// required: true
		StartDate time.Time

		// the lease end date
		// required: true
		EndDate time.Time

		// the lease rent cost
		// required: true
		Rent decimal.Decimal

		// the lease utilities cost
		// required: true
		Utilities decimal.Decimal

		// the lease parking cost
		// required: true
		ParkingCost decimal.Decimal

		// the lease total cost
		// required: true
		TotalCost decimal.Decimal

		// the lease square footage
		// required: true
		SquareFootage uint

		// the lease furnished value
		// required: true
		Furnished bool

		// the lease parking value
		// required: true
		Parking bool

		// the lease bedroom count
		// required: true
		Beds uint

		// the lease baths count
		// required: true
		Baths decimal.Decimal

		// the lease amenities
		// required: true
		Amenities string

		// the lease appliances
		// required: true
		Appliances string

		// the lease description
		// required: true
		Description string

		// the lease contacts
		// required: true
		Contacts []entity.Contact `faker:"contactsEntityFaker"`
	}
}
