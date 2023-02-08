import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { LeaseService } from './services/lease.service';
import { Lease } from './models/lease';
import { PostLeaseRequest } from './models/PostLeaseRequest';

@Component({
  selector: 'app-root',
  template: `
      <h1>{{ title }}</h1>
      <ul>

      <app-post (newPost)="addLease($event)"></app-post>

      <br> <br>
      <h2> Current Listings: </h2>
      <body ng-app="myApp">
      
      <div *ngFor="let lease of leases">
            <p><a href="#lease:{{lease.name}}"> {{ lease.name }} </a></p>
      </div>

      <p><a href="#!3-bed-4-by-4-standard">3 bedroom in a 4x4 at the Standard</a> </p>

      <p><a href="#!4-bed-4-by-2-lark">4 bedrooms in a 4x2 at the Lark</a> </p>
      
      <div ng-view></div>
            
    `  
  //templateUrl: './app.component.html',    for writing code in .html file
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

  addLease(newPost: string){
    var leas = <Lease>{};
    leas.name = newPost;
    this.leases.push(leas);
  }

}
