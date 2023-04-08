package viewModel

type PaginatedLeasesResult struct {
	Leases          []*Lease `json:"leases"`
	Count           uint     `json:"count"`
	PaginationToken string   `json:"paginationToken"`
}
