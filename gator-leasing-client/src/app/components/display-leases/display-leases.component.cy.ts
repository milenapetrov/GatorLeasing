import { TestBed } from '@angular/core/testing';
import { DisplayLeasesComponent } from './display-leases.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpClientModule } from '@angular/common/http';

describe('DisplayLeasesComponent', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    providers: [DisplayLeasesComponent]
  }));

  it('can mount', () => {
    cy.mount(DisplayLeasesComponent);
  });
});
