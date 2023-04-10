import { TestBed } from '@angular/core/testing';

import { AuthHttpInterceptorExtendedService } from './auth-http-interceptor-extended.service';
import { LeaseService } from './lease.service';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('Http service', () => {
  it('playground', () => {
    cy.mount(AuthHttpInterceptorExtendedService);
  });
});

/*describe('AuthHttpInterceptorExtendedService', () => {
  let service: AuthHttpInterceptorExtendedService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [],
      imports: [HttpClientTestingModule]
    });
    service = TestBed.inject(AuthHttpInterceptorExtendedService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});*/
