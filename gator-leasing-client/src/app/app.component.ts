import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { LeaseService } from './services/lease.service';
import { Lease } from './models/lease';
import { PostLeaseRequest } from './models/PostLeaseRequest';

@Component({
  selector: 'app-root',
  template: `
      <h1>{{ title }}</h1>
      <ul>
      <app-post></app-post>

      <br> <br>
      <h2> Current Listings: </h2>
      <body ng-app="myApp">
      
      <div *ngFor="let lease of leases">
            <p><a href="#leaseNumber{{lease.id}}"> {{lease.id}} - {{ lease.name }} </a></p>
        </div>

      <p><a href="#!3-bed-4-by-4-standard">3 bedroom in a 4x4 at the Standard</a> </p>

      <p><a href="#!4-bed-4-by-2-lark">4 bedrooms in a 4x2 at the Lark</a> </p>
      
      <div ng-view></div>
            
    `  
  //templateUrl: './app.component.html',    for writing code in .html file
  ,styleUrls: ['./app.component.css']
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

  /*addPost(post: String){
    if(post){
      this.posts.push(post);
    }
  }*/
}

// <app-post> </app-post>
/*<input #post
            (keyup.enter)= "addPost(post.value)"
            (blur)= "addPost(post.value); post.value='' ">

        <button type="button" (click)="addPost(post.value)"> post </button>

        <ul><div *ngFor="let p of posts">
            {{ p }}
        </div></ul>*/