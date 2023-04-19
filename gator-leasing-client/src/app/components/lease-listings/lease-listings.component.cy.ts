import { LeaseListingsComponent } from './lease-listings.component';
import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('LeaseListingsComponent', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    providers: [LeaseListingsComponent]
  }));

  it('can mount', () => {
    cy.mount(LeaseListingsComponent);
  });
});
