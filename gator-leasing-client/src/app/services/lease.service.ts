import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Lease } from '../models/lease'

@Injectable({
  providedIn: 'root'
})
export class LeaseService {
  constructor(private http : HttpClient) { }

  getLeases(): Observable<Lease[]> {
    return this.http.get<Lease[]>("http://localhost:8080/leases");
  }
}
