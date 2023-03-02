import { Component } from '@angular/core';

import { LeaseService } from '../services/lease.service';
 import { Lease } from '../models/lease';

@Component({
  selector: 'app-my-leases',
  templateUrl: './my-leases.component.html',
  styleUrls: ['./my-leases.component.css']
})
export class MyLeasesComponent {
  leases: Lease[] = [];

  showPost(name: string){
    
  }
}
