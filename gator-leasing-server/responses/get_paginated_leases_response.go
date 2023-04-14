package responses

import (
	viewModel "github.com/milenapetrov/GatorLeasing/gator-leasing-server/view-model"
)

// The response that is returned when paginated leases are requested
// swagger:response GetPaginatedLeasesResponse
type GetPaginatedLeasesResponse struct {
	// in: body
	Body struct {
		// the list of leases
		//
		// required: true
		Leases []*viewModel.Lease
		// the lease count
		// required: true
		Count uint
		// the pagination token
		// required: true
		PaginationToken string
	}
}
