import { Lease } from "./lease";


export interface PaginatedLeasesResult {
    leases : Lease[]
    paginationToken : string
    count : number
}