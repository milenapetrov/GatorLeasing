import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpTestingController, HttpClientTestingModule } from '@angular/common/http/testing';
import { HomeComponent } from './home.component';
import { HttpClient } from '@angular/common/http';
import { LeaseService } from '../services/lease.service';
import { BreakpointObserver } from '@angular/cdk/layout';
import { Router } from '@angular/router';
import { AuthService } from '@auth0/auth0-angular';
//import { AUTH0_CLIENT } from 'auth0-client.token';

describe('HomeComponent', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    providers: [
      BreakpointObserver,
      Router,
      LeaseService,
      AuthService,
      { /*provide: AUTH0_CLIENT, useValue: your auth0.client value here */ }
    ]
  }));

  it('can mount', () => {
    cy.mount(HomeComponent);
  });
});

describe('Search', () => {
  it('has type text', () => {
    cy.get('input[name="search"]').should('have.attr', 'type', 'text');
  });
});

describe('Search Button', () => {
  it('can click', () => {
    cy.mount(HomeComponent);
    cy.get('button[name="s"]').click();
  });
});

describe('Post Button', () => {
  it('can click', () => {
    cy.mount(HomeComponent);
    cy.get('button[name="post"]').click();
  });
});


/*
describe('HomeComponent', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    providers: [BreakpointObserver, Router, HomeComponent]
  }));

  it('can mount', () => {
    cy.mount(HomeComponent);
  });
});

describe('Search', () => {
  it('is string', () => {
    cy.get('input[name="search"]').should(String);
  });
});

describe('Search Button', () => {
  it('can click', () => {
    cy.mount(HomeComponent);
    cy.get('button[name="s"]').click();
  });
});

describe('Post Button', () => {
  it('can click', () => {
    cy.mount(HomeComponent);
    cy.get('button[name="post"]').click();
  });
});
*/


/*
describe('HomeComponent', () => {
  let component: HomeComponent;
  let fixture: ComponentFixture<HomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HomeComponent ],
      providers: [LeaseService],
      imports: [HttpClientTestingModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});*/
