package entity

import "github.com/milenapetrov/GatorLeasing/gator-leasing-server/enums"

type PaginatedLeasesRequest struct {
	PageSize        uint
	SortToken       string
	PaginationToken string
	SortDirection   enums.SortDirection
}
