import { Component } from '@angular/core';
import { FormBuilder} from '@angular/forms';
import { LeaseService } from 'src/app/services/lease.service';
import { Post } from 'src/app/models/post';
import { Address } from 'src/app/models/address'

@Component({
  selector: 'app-update',
  templateUrl: './update.component.html',
  styleUrls: ['./update.component.css']
})
export class UpdateComponent {
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


  constructor(private leaseService:LeaseService, private formBuilder: FormBuilder){
  }

  onSubmit(post:Post) {
    this.leaseService.updatePost(this.post);
  };
}