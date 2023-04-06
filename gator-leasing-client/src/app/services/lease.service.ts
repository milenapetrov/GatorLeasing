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

  getLease(id: number): Observable<Lease> {
    console.log(`get lease ${id}`)
    return this.http.get<Lease>(`${this.BASEURL}${this.LEASEURL}/${id}`)
  }

  createPost(post: Post){
    this.http.post<Post>(`${this.BASEURL}${this.LEASEURL}`, post).subscribe(response =>{ 
      console.log(response);  })  
  }

  updatePost(id: number, post:Post): Observable<any> {
    console.log(`update lease ${id}`)
    return this.http.put(`${this.BASEURL}${this.LEASEURL}/${id}`, post);
  }
}
