import { Component } from '@angular/core';
import { FormControl, FormBuilder} from '@angular/forms';
import { LeaseService } from 'src/app/services/lease.service';
import { Lease } from 'src/app/models/lease';
import { Post } from 'src/app/models/post';
import { timeInterval } from 'rxjs';

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

  post: Post = {
    name: '', 
    rent: 0.0,
    start_date: this.date,
    end_date: this.date,
    utilities: 0.0,
    parking_cost: 0.0,
    square_footage: 0,
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
    console.warn("date: ", this.post.start_date)
    this.leaseService.createPost(this.post);
  };
}
