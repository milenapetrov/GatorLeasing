import { Component , Input, Output, EventEmitter, ChangeDetectorRef} from '@angular/core';

import { LeaseService } from '../services/lease.service';
import { Lease } from '../models/lease';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent {
  leases: Lease[] = [];
  
  constructor(private leaseService:LeaseService, private cd: ChangeDetectorRef){
    this.loadLeases();
  }

  loadLeases(){
    this.leaseService.getLeases().subscribe((leases) => { this.leases = leases});
  }
  //for displaying input read in this component in the parent app component
  @Output() newPost = new EventEmitter<string>(); 

  /*addPost(name: string) {
    this.newPost.emit(name);
    this.leaseService.createPost(name);
  };*/

}

  