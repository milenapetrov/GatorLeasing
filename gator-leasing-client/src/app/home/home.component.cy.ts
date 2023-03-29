import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeComponent } from './home.component';


describe('Home Component', () => {
  it('playground', () => {
    cy.mount(HomeComponent)
    cy.get('input[name="search"]').should(String)
    cy.get('button[name="s"]').click()
    cy.get('button[name="post"]')
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
