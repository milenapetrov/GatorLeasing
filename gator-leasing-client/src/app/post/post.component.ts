import { Component } from '@angular/core';
import { PostLeaseRequest } from '../models/PostLeaseRequest';

@Component({
  selector: 'app-post',
  template: `
    <div> 
      <label for="name"> Post name: </label>
      <input #p
            (keyup.enter)= "addPost(p.value)"
            
            placeholder="name">

      <button type="button" (click)="addPost(p.value)"> post </button>
    </div>
  

    <br>
    <br>
    <div class="box">
    <br>
      {{post.name}}
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
    name: ' '
  };

  addPost(name: string) {
    this.post.name = name;
  };
}


//<input id="name" [(ngModel)]="post.name" placeholder="name">         for dynamic input
//(blur)= "addPost(p.value); p.value='' "