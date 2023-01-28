import { TestBed } from '@angular/core/testing';

import { LeaseService } from './lease.service';

describe('LeaseService', () => {
  let service: LeaseService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LeaseService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
