import { Component } from '@angular/core';

import { LeaseService } from '../services/lease.service';
import { Lease } from '../models/lease';
import { Post } from '../models/post';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  posts: Post[] = [];
  leases: Lease[] = [];

  constructor(private leaseService:LeaseService){
    this.loadLeases();
  }

  loadLeases(){
    this.leaseService.getLeases().subscribe((leases) => { this.leases = leases});
  }

}
