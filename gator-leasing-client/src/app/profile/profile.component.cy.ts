import { ProfileComponent } from './profile.component'
import { LeaseService } from '../services/lease.service';
import { HttpClientTestingModule} from '@angular/common/http/testing';

describe('ProfileComponent', () => {
  it('can mount', () => {
    cy.mount(ProfileComponent, {
      providers: [LeaseService],
      imports: [HttpClientTestingModule]
    })
    cy.get('input').type("Update Profile")
    cy.get('button').click()
  })
})