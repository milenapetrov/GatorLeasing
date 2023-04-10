import { DisplayLeasesComponent } from './display-leases.component';
import { LeaseService } from '../../services/lease.service';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('DisplayLeasesComponent', () => {
  it('should mount', () => {
    cy.mount(DisplayLeasesComponent);
  });
});
