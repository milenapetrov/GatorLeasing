// view.component.spec.ts

import { ViewComponent } from './view.component';
import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpClientModule } from '@angular/common/http';
import { LeaseListingsComponent } from '../lease-listings/lease-listings.component';

describe('ViewComponent', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      declarations: [ViewComponent],
      providers: []
    });
  });

  it('can mount', () => {
    cy.mount(ViewComponent);
    cy.get('app-view').should('exist');
  });
});


/*import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ViewComponent } from './view.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';
import { LeaseService } from 'src/app/services/lease.service';
import { Lease } from 'src/app/models/lease';
import { Address } from 'src/app/models/address';
import { GridCellComponent } from 'src/app/components/grid-cell/grid-cell.component';
import { LeaseListingsComponent } from '../lease-listings/lease-listings.component';
import { DisplayLeasesComponent } from 'src/app/components/display-leases/display-leases.component';
import { HttpClientModule } from '@angular/common/http';

describe('ViewComponent', () => {
  let component: ViewComponent;
  let fixture: ComponentFixture<ViewComponent>;
  let leaseService: jasmine.SpyObj<LeaseService>;
  let lease: Lease;
  let route: ActivatedRoute;

  beforeEach(async () => {
    leaseService = jasmine.createSpyObj('LeaseService', ['getLease']);
    lease = {
      id: 1,
      name: 'Test Lease',
      createdAt: new Date(),
      ownerID: 1,
      address: {} as Address,
      rent: 1000,
      startDate: new Date(),
      endDate: new Date(),
      parkingCost: 0,
      squareFootage: 0,
      furnished: true,
      parking: true,
      beds: 1,
      baths: 1,
      amenities: '',
      appliances: '',
      description: '',
      utilites: 0,
      totalCost: 1000,
      contacts: {} as Address,
    };
    leaseService.getLease.and.returnValue(of(lease));

    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, HttpClientModule],
      declarations: [ViewComponent, GridCellComponent, LeaseListingsComponent],
      providers: [
        {
          provide: ActivatedRoute,
          useValue: { queryParams: of({ data: 1 }) },
        },
        { provide: LeaseService, useValue: leaseService },
      ],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).to.exist;
  });

  it('should set the post correctly', () => {
    expect(component.post).equal(lease);
  });
});

*/

/*import { ViewComponent } from './view.component'
import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpClientModule } from '@angular/common/http';
import { LeaseListingsComponent } from '../lease-listings/lease-listings.component';

describe('ViewComponent', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    declarations: [ViewComponent, LeaseListingsComponent], // Add LeaseListingsComponent here
    providers: [TestBed]
  }));

  it('can mount', () => {
    cy.mount(ViewComponent);
  });
});


//
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
*/
