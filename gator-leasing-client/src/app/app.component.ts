import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { LeaseService } from './services/lease.service';
import { Lease } from './models/lease';
import { Post } from './models/post';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html', 
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  leases: Lease[] = [];
  title = 'Leases';

  constructor(private leaseService:LeaseService){
    this.loadLeases();
  }

  loadLeases(){
    this.leaseService.getLeases().subscribe((leases) => { this.leases = leases});
  }

  //sends input to server and adds it to lease array
  postLease(newPost: string){
    var leas = <Lease>{};
    leas.name = newPost;
    this.leases.push(leas);
    //this.leaseService.createPost(newPost);
  }

}

