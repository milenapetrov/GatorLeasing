import { LoginComponent } from './login.component';

import { AuthModule, AuthService } from '@auth0/auth0-angular';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('LoginComponent', () => {
  it('can mount', () => {
    cy.mount(LoginComponent, {
      providers: [AuthService],
      imports: [
        AuthModule.forRoot({
          domain: 'auth0.domain',
          clientId: 'autho0.client',
        }),
      ],
    });
    cy.get('a');
  });
});
