
import { PostComponent } from '../src/app/post/post.component'

import { LeaseService } from '../src/app/services/lease.service';
import { HttpClientTestingModule} from '@angular/common/http/testing';

describe('PostComponent', () => {
  it('can mount', () => {
    cy.mount(PostComponent, {
      providers: [LeaseService],
      imports: [HttpClientTestingModule]
    })
  })

})