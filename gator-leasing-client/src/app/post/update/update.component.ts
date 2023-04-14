import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { LeaseService } from 'src/app/services/lease.service';
import { Post } from 'src/app/models/post';
import { Address } from 'src/app/models/address';
import { Lease } from 'src/app/models/lease';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-update',
  templateUrl: './update.component.html',
  styleUrls: ['./update.component.css'],
})
export class UpdateComponent {

  date: Date = new Date();
  addy: Address = {
    street: '',
    roomNumber: '',
    city: '',
    state: '',
    zipCode: '',
  };

  post: Lease = {
    name: '',
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
    id: 0,
    createdAt: this.date,
    ownerID: 0,
    utilites: 0,
    totalCost: 0,
    contacts: this.addy   //add contacts
  };

  constructor(
    private leaseService: LeaseService,
    private formBuilder: FormBuilder,
    private route: ActivatedRoute
  ) {}

  id = 0;

  ngOnInit(){
    this.route.queryParams.subscribe(params => {
      this.id = params['fieldParam'];
    })
    this.leaseService.getLease(this.id).subscribe(res => {
      console.log(res)
      this.post = res;
    })
  }
  onSubmit(post: Post) {
    this.leaseService.updatePost(this.id, this.post).subscribe(res => {
      console.log(res)
    });
  }
}
