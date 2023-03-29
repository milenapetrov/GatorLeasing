import { ProfileComponent } from './profile.component'
import { LeaseService } from '../services/lease.service';
import { HttpClientTestingModule} from '@angular/common/http/testing';
import { ReactiveFormsModule, FormsModule} from '@angular/forms';

describe('ProfileComponent', () => {
  it('can mount', () => {
    cy.mount(ProfileComponent, {
      providers: [LeaseService],
      imports: [HttpClientTestingModule, ReactiveFormsModule, FormsModule]
    })
    cy.get('input[name="firstName"]').should(String)
    cy.get('input[name="lastName"]').should(String)
    cy.get('button').click()
  })
})