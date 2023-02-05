import { Component } from '@angular/core';
import { LeaseService } from './services/lease.service';
import { Lease } from './models/lease';

@Component({
  selector: 'app-root',
  template: `
      <h2>{{ title }}</h2>
      <ul>
        <li *ngFor="let lease of leases">
            {{ lease.id + '-' + lease.name }}
        </li>
      </ul>
    `  
  //templateUrl: './app.component.html',
  ,styleUrls: ['./app.component.css']
})
export class AppComponent {
  leases: Lease[] = [];
  title = 'Leases';

  constructor(private leaseService:LeaseService){
    this.loadLeases();
  }

  loadLeases(){
    this.leaseService.getLeases().subscribe((leases) => { this.leases = leases});
  }
}
