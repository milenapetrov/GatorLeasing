import { HomeComponent } from '../src/app/home/home.component'

describe('Home.cy.ts', () => {
  it('playground', () => {
    cy.mount(HomeComponent)
  })
})