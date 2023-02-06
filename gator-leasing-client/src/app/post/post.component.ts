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
      <br>
      <div class="box">
       <br>
        1 bedroom in a 4x4 at the Standard.
        <br>
        <br> 
        <br>
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


