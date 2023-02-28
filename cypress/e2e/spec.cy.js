describe('Nav Bar tests', () => {
  it('Visits Search Page and Return Home', () => {
    cy.visit('http://localhost:3000/')
    cy.contains('Search').click()
    cy.url().should('include', '/search')

    cy.contains('Mind My Wallet').click()
    cy.url().should('include', '/')

  })
})