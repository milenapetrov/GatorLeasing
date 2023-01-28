import { Component } from '@angular/core';
import { LeaseService } from './services/lease.service';
import { Lease } from './models/lease';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  leases: Lease[] = [];
  title = 'gator-leasing-client';

  constructor(private leaseService:LeaseService){
    this.loadLeases();
  }

  loadLeases(){
    this.leaseService.getLeases().subscribe((leases) => this.leases = leases);
  }
}
