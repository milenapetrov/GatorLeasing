import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Lease } from '../models/lease'
import { Post } from '../models/post'

@Injectable({
  providedIn: 'root'
})
export class LeaseService {
  constructor(private http : HttpClient) {}

  BASEURL: string = 'http://localhost:8080'
  LEASEURL: string = '/leases'

  getLeases(): Observable<Lease[]> {
    console.log("get leases")
    return this.http.get<Lease[]>(`${this.BASEURL}${this.LEASEURL}`);
  }

  createPost(post: Post){
    this.http.post<Post>(`${this.BASEURL}${this.LEASEURL}`, post).subscribe(response =>{ 
      console.log(response);  })  
  }

  updatePost(post: Post): Observable<any> {
    return this.http.put(`${this.BASEURL}${this.LEASEURL}`, post);
  }
}
