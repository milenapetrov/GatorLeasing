import { Component } from '@angular/core';
import { FormBuilder} from '@angular/forms';
import { LeaseService } from 'src/app/services/lease.service';
import { Post } from 'src/app/models/post';
import { Address } from 'src/app/models/address';
import {Lease} from 'src/app/models/lease';

@Component({
  selector: 'app-update',
  templateUrl: './update.component.html',
  styleUrls: ['./update.component.css']
})
export class UpdateComponent {
  date: Date = new Date();
  //lease = this.leaseService.getLease(13)

  addy: Address = {
    street: '',
    roomNumber: '',
    city: '',
    state: '',
    zipCode: ''
  };

  post: Post = {
    name: '', 
    address: this.addy,
    rent: 0.0,
    startDate: this.date,
    endDate: this.date,
    utilities: 0.0,
    parkingCost: 0.0,
    squareFootage: 0,
    furnished: false,
    parking: false,
    beds: 0,
    baths: 0.0,
    amenities: '',
    appliances: '',
    description: ''
  }

  /*lease2: Lease = {
    id: 0,
    name: '',
    ownerID: 0,
    address: this.addy,
    startDate: this.date,
    endDate: this.date,
    rent: 0,
    utilites: 0,
    parkingCost: 0,
    totalCost: 0,
    squareFootage: 0,
    furnished: false,
    parking: false,
    beds: 0,
    baths: 0,
    amenities: '',
    appliances: '',
    description: '',
    contacts: this.addy   //placeholder idk
  }*/

  
  constructor(private leaseService:LeaseService, private formBuilder: FormBuilder){
  }

  onSubmit(post:Post) {
    this.leaseService.updatePost(13, this.post);
  };

}
