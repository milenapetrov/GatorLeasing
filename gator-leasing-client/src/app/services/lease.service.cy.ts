import { TestBed } from '@angular/core/testing';

import { LeaseService } from './lease.service';
import { HttpClientTestingModule} from '@angular/common/http/testing';

describe('Lease service', () => {
  it('playground', () => {
    cy.mount(LeaseService, {
      imports: [HttpClientTestingModule]
    })
  })
})

/*describe('LeaseService', () => {
  let service: LeaseService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [LeaseService],
      imports: [HttpClientTestingModule]
    });
    service = TestBed.inject(LeaseService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});*/
