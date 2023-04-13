import { Component, Input } from '@angular/core';
import { LeaseService } from 'src/app/services/lease.service';
import { Lease } from 'src/app/models/lease';
import { Post } from 'src/app/models/post';
import { Address } from 'src/app/models/address';
import { DisplayLeasesComponent } from 'src/app/components/display-leases/display-leases.component';
import { GridCellComponent } from 'src/app/components/grid-cell/grid-cell.component';
import { LeaseListingsComponent } from '../lease-listings/lease-listings.component';

@Component({
  selector: 'app-view',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.css'],
  
})
export class ViewComponent {
  date: Date = new Date();
  addy: Address = {
    street: '',
    roomNumber: '',
    city: '',
    state: '',
    zipCode: '',
  };

  post: Lease = {
    id: 0,
    name: '',
    createdAt: this.date,
    ownerID: 0,
    address: this.addy,
    rent: 0.0,
    startDate: this.date,
    endDate: this.date,
    parkingCost: 0.0,
    squareFootage: 0,
    furnished: false,
    parking: false,
    beds: 0,
    baths: 0.0,
    amenities: '',
    appliances: '',
    description: '',
    utilites: 0,
    totalCost: 0,
    contacts: this.addy  //add contacts
  };

  constructor ( private leaseService: LeaseService, private listing:LeaseListingsComponent, private grid: GridCellComponent) {
    console.log(this.listing.getID());
  }
  
  id = this.listing.getID();
  
  lease = this.leaseService.getLease(0).subscribe(res => {
      console.log(res)
      this.post = res;
  })
  
}
