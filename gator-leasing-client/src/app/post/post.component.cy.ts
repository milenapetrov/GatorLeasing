import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PostComponent } from './post.component';
import { LeaseService } from '../services/lease.service';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('PostComponent', () => {
  it('can mount', () => {
    cy.mount(PostComponent, {
      providers: [LeaseService],
      imports: [HttpClientTestingModule],
    });
  });
});

/*
describe('Post Title Input', () => {
  it('is string', () => {
    cy.get('input[name="postName"]').should(String)
  })
}) */

describe('Input', () => {
  it('is new post', () => {
    cy.mount(PostComponent, {
      providers: [LeaseService],
      imports: [HttpClientTestingModule],
    });
    //cy.get('input').type("New Post!")
  });
});

describe('Button', () => {
  it('can click', () => {
    cy.mount(PostComponent, {
      providers: [LeaseService],
      imports: [HttpClientTestingModule],
    });
    //cy.get('button').click()
  });
});
/*describe('PostComponent', () => {
  let component: PostComponent;
  let fixture: ComponentFixture<PostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      providers: [LeaseService],
      imports: [HttpClientTestingModule],
      declarations: [ PostComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });


});*/
