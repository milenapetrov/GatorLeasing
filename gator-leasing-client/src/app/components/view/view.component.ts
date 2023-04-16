import { Component, Input } from '@angular/core';
import { LeaseService } from 'src/app/services/lease.service';
import { Lease } from 'src/app/models/lease';
import { Post } from 'src/app/models/post';
import { Address } from 'src/app/models/address';
import { DisplayLeasesComponent } from 'src/app/components/display-leases/display-leases.component';
import { GridCellComponent } from 'src/app/components/grid-cell/grid-cell.component';
import { LeaseListingsComponent } from '../lease-listings/lease-listings.component';
import { ActivatedRoute } from '@angular/router';
import { format, parseISO } from 'date-fns';

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
  id = 0;
  appliances: string[] = [];
  amenities: string[] = [];
  start: string = '';
  end: string = '';
  constructor ( private leaseService: LeaseService, private listing:LeaseListingsComponent, private grid: GridCellComponent, private route: ActivatedRoute) {}

  getApps(){
    var apps = this.post.appliances;
    this.appliances= apps.split(',');
    var amens = this.post.amenities;
    this.amenities = amens.split(',');
    this.start = format(parseISO(this.post.startDate.toLocaleString()), 'MM/dd/yyyy')
    this.end = format(parseISO(this.post.endDate.toLocaleString()), 'MM/dd/yyyy')
  }
  ngOnInit(){
    this.route.queryParams.subscribe(params => {
      this.id = params['fieldParam'];
      console.log(params['fieldParam'])
    })
    this.leaseService.getLease(this.id).subscribe(res => {
      console.log(res)
      this.post = res;
    })

  }

}
