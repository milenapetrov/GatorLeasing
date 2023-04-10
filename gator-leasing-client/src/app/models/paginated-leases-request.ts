export interface PaginatedLeasesRequest {
  pageSize: number;
  sortToken: string;
  paginationToken: string;
  sortDirection: number;
  filters: string;
}
