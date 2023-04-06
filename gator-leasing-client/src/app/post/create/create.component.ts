import { Component } from '@angular/core';
import { FormBuilder} from '@angular/forms';
import { LeaseService } from 'src/app/services/lease.service';
import { Lease } from 'src/app/models/lease';
import { Post } from 'src/app/models/post';
import { Address } from 'src/app/models/address'
import { Router } from '@angular/router';

const today = new Date();
const month = today.getMonth();
const year = today.getFullYear();

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})

export class CreateComponent {
  date: Date = new Date();
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


  constructor(private leaseService:LeaseService, private formBuilder: FormBuilder, private router: Router){
  }

  onSubmit(post:Post) {
    this.leaseService.createPost(this.post);
    this.router.navigateByUrl('/display');
  };

}
