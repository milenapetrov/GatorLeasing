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
        <div *ngFor="let lease of leases">
            {{ lease.id }} -  {{ lease.name }}
        </div>
      <app-post></app-post>
      
      </ul>
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