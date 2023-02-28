import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Lease } from '../models/lease'
import { Post } from '../models/post'

@Injectable({
  providedIn: 'root'
})
export class LeaseService {
  constructor(private http : HttpClient) { }

  getLeases(): Observable<Lease[]> {
    console.log("get leases")
    return this.http.get<Lease[]>("http://localhost:8080/leases");
  }

  createPost(name: string){
    const postData: Post = {name: name};
    this.http.post<Number>("http://localhost:8080/leases", postData).subscribe(response =>{ 
      console.log(response);  })  
  }
}
