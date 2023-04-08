export interface PaginatedLeasesRequest {
    pageSize : number
    sortToken : string
    paginationToken : string
    sortDirection : number
    filter : string
}