package viewModel

import (
	"encoding/json"
	"errors"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"
)

// swagger:parameters GetPaginatedLeases GetMyLeasesPaged
type PaginatedLeasesRequest struct {
	// the page size
	// in: body
	PageSize int
	// the sort token
	// in: body
	SortToken string
	// the paginated token
	// in: body
	PaginationToken string
	// the sort direction {asc, desc}
	// required: true
	// in: body
	SortDirection enums.SortDirection `faker:"sortDirectionFaker"`
	// the filters
	// in: body
	Filters string `faker:"filtersFake"`
}

func (r *PaginatedLeasesRequest) MarshalJSON() ([]byte, error) {
	enumToString := map[enums.SortDirection]string{
		enums.Ascending:  "asc",
		enums.Descending: "desc",
	}
	sortDirection, ok := enumToString[r.SortDirection]
	if !ok {
		return nil, errors.New("unknown sort direction")
	}

	request := struct {
		PageSize        int    `json:"pageSize"`
		SortToken       string `json:"sortToken"`
		PaginationToken string `json:"paginationToken"`
		SortDirection   string `json:"sortDirection"`
		Filters         string `json:"filters"`
	}{
		PageSize:        r.PageSize,
		SortToken:       r.SortToken,
		PaginationToken: r.PaginationToken,
		SortDirection:   sortDirection,
		Filters:         r.Filters,
	}

	return json.Marshal(&request)
}

func (r *PaginatedLeasesRequest) UnmarshalJSON(data []byte) error {
	request := struct {
		PageSize        int    `json:"pageSize"`
		SortToken       string `json:"sortToken"`
		PaginationToken string `json:"paginationToken"`
		SortDirection   string `json:"sortDirection"`
		Filters         string `json:"filters"`
	}{}

	if err := json.Unmarshal(data, &request); err != nil {
		return err
	}

	stringToEnum := map[string]enums.SortDirection{
		"asc":  enums.Ascending,
		"desc": enums.Descending,
	}

	r.PageSize = request.PageSize
	r.SortToken = request.SortToken
	r.PaginationToken = request.PaginationToken
	enum, ok := stringToEnum[request.SortDirection]
	if !ok {
		return errors.New("sort direction must be 'asc' or 'desc'")
	}
	r.SortDirection = enum
	r.Filters = request.Filters
	return nil
}
