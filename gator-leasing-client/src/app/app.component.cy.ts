import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';

describe('App Component', () => {
  it('Should mount', () => {
    cy.mount(AppComponent);
    cy.get('mat-toolbar');
    cy.get('span');
    cy.get('app-login').invoke('show');
  });
});

describe('Toolbar Buttons', () => {
  it('should click', () => {
    cy.mount(AppComponent);
    cy.get('button[name="mylease"]').click();
    cy.get('button[name="messages"]').click();
  });
});

describe('Non-buttons', () => {
  it('should click', () => {
    cy.mount(AppComponent);
    cy.get('span[name="home"]').click();
    cy.get('a[name="profile"]').click();
  });
});
/*
describe('AppComponent', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule
      ],
      declarations: [
        AppComponent
      ],
    }).compileComponents();
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });


  it('should render title', () => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('.content span')?.textContent).toContain('gator-leasing-client app is running!');
  });
});*/
