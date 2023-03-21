package viewModel

type PaginatedLeasesResult struct {
	Leases          []*Lease
	Count           uint
	PaginationToken string
}
