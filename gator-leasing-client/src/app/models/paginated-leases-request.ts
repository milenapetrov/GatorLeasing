import { SortDirection } from "src/enums/sort-direction";

export interface PaginatedLeasesRequest {
  pageSize: number;
  sortToken: string;
  paginationToken: string;
  sortDirection: SortDirection;
  filters: string;
}