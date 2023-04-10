import { ProfileComponent } from './profile.component';
import { LeaseService } from '../services/lease.service';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

describe('ProfileComponent', () => {
  it('can mount', () => {
    cy.mount(ProfileComponent, {
      providers: [LeaseService],
      imports: [HttpClientTestingModule, ReactiveFormsModule, FormsModule],
    });
  });
});

describe('First Name Input', () => {
  it('is string', () => {
    cy.get('input[name="firstName"]').should(String);
  });
});

describe('Last Name Input', () => {
  it('is string', () => {
    cy.get('input[name="lastName"]').should(String);
  });
});

describe('Phone Number Input', () => {
  it('is string', () => {
    cy.get('input[name="phoneNumber"]').should(String);
  });
});

describe('Email Input', () => {
  it('is string', () => {
    cy.get('input[name="email"]').should(String);
  });
});

describe('Address Input', () => {
  it('is address', () => {
    cy.get('input[name="street"]').should(String);
    cy.get('input[name="roomNumber"]').should(String);
    cy.get('input[name="city"]').should(String);
    cy.get('input[name="state"]').should(String);
    cy.get('input[name="zipCode"]').should(String);
  });
});

describe('Button', () => {
  it('can click', () => {
    cy.mount(ProfileComponent, {
      providers: [LeaseService],
      imports: [HttpClientTestingModule, ReactiveFormsModule, FormsModule],
    });
    cy.get('button').click();
  });
});
