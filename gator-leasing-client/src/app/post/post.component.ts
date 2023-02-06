import { Component } from '@angular/core';
import { PostLeaseRequest } from '../models/PostLeaseRequest';

@Component({
  selector: 'app-post',
  template: `
  <h2> {{ post.name }} </h2>
  <div> 
      <label for="name"> Post name: </label>
      <input id="name" [(ngModel)]="post.name" placeholder="name">
   </div>
      <div class="box">
        This text is enclosed in a box.
      </div>
  
  `
  //templateUrl: './post.component.html',
  ,styleUrls: ['./post.component.css']
})
export class PostComponent {
  //name: 'Ur moms Penthouse'
  post: PostLeaseRequest ={
    name: 'urmom'
  };
}


