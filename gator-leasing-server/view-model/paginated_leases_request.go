package viewModel

import "github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"

type PaginatedLeasesRequest struct {
	PageSize        int                 `json:"pageSize"`
	SortToken       string              `json:"sortToken"`
	PaginationToken string              `json:"paginationToken"`
	SortDirection   enums.SortDirection `json:"sortDirection"`
	Filters         string              `json:"filters"`
}
