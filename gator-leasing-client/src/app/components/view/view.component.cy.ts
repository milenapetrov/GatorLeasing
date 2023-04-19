import { ViewComponent } from './view.component'
import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpClientModule } from '@angular/common/http';

describe('ViewComponent', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    declarations: [ViewComponent],
    providers: [TestBed]
  }));

  it('can mount', () => {
    cy.mount(ViewComponent);
  });
}); 
