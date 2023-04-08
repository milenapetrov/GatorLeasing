import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Lease } from '../models/lease'
import { PaginatedLeasesRequest } from '../models/paginated-leases-request';
import { PaginatedLeasesResult } from '../models/paginated-leases-result';
import { Post } from '../models/post'

@Injectable({
  providedIn: 'root'
})
export class LeaseService {
  constructor(private http : HttpClient) {}

  BASEURL: string = 'http://localhost:8080'
  LEASEURL: string = '/leases'
  PAGEDURL: string = '/leases/paged'

  getLeases(): Observable<Lease[]> {
    return this.http.get<Lease[]>(`${this.BASEURL}${this.LEASEURL}`);
  }

  getPagedLeases(pageSize : number, sortToken : string, paginationToken : string, sortDirection : number, filter : string): Observable<PaginatedLeasesResult> {
    const paginatedLeasesRequest : PaginatedLeasesRequest = {
      pageSize: pageSize,
      sortToken: sortToken,
      paginationToken: paginationToken,
      sortDirection: sortDirection,
      filter: filter
    }
    return this.http.post<PaginatedLeasesResult>(`${this.BASEURL}${this.PAGEDURL}`, paginatedLeasesRequest)
  }

  getLease(id: number): Observable<Lease> {
    console.log(`get lease ${id}`)
    return this.http.get<Lease>(`${this.BASEURL}${this.LEASEURL}/${id}`)
  }

  createPost(post: Post){
    this.http.post<Post>(`${this.BASEURL}${this.LEASEURL}`, post).subscribe(response =>{ 
      console.log(response);  })  
  }

  updatePost(post: Post): Observable<any> {
    return this.http.put(`${this.BASEURL}${this.LEASEURL}`, post);
  }
}
