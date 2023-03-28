import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeComponent } from './home.component';


describe('Home Component', () => {
  it('playground', () => {
    cy.mount(HomeComponent)

    cy.get('button')
    cy.get('button')
    cy.click()
    
  })

})
/*
describe('HomeComponent', () => {
  let component: HomeComponent;
  let fixture: ComponentFixture<HomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HomeComponent ],
      providers: [LeaseService],
      imports: [HttpClientTestingModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});*/
