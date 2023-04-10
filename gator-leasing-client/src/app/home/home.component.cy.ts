import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeComponent } from './home.component';
//import { AuthModule, AuthService } from '@auth0/auth0-angular';

describe('Home Component', () => {
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
