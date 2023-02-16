import { Component , Input, Output, EventEmitter} from '@angular/core';

import { LeaseService } from '../services/lease.service';
import { Post } from '../models/post';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent {
  constructor(private leaseService:LeaseService){}

  //for displaying input read in this component in the parent app component
  @Output() newPost = new EventEmitter<string>(); 

  addPost(name: string) {
    this.newPost.emit(name);
    this.leaseService.createPost(name);
  };

}

  