describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:3000/')

    cy.contains('Log In').click()
    cy.url().should('include', '/Login')

    cy.get("input[type ='text']").type('harry')
    cy.get("input[type ='password']").type('password')

    cy.contains("Log in").click()
    cy.url().should('include', '/dashboard')

    cy.contains('Search').click()
    cy.url().should('include', '/search')

    cy.get("input[type ='text']").type('AAPL')
    cy.get('button').contains("Search").click()

    cy.contains("Log In").click()

    cy.contains('Dashboard').click()
    cy.url().should('include', '/dashboard')
    cy.wait(10000)
    cy.contains('AAPL')
  })
})