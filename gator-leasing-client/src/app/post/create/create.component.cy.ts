import { CreateComponent } from './create.component';
import { LeaseService } from '../../services/lease.service';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

describe('CreateComponent', () => {
  it('should mount', () => {
    cy.mount(CreateComponent, {
      providers: [LeaseService],
      imports: [HttpClientTestingModule, ReactiveFormsModule, FormsModule],
    });
  });

});

/*
describe('Button', () => {
  it('should click', () => {
    cy.get('button[name="post"]').click();
  })
});
*/