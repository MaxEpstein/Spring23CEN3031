describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:3000/')
    cy.contains('Search').click()
    cy.url().should('include', '/search')

    cy.get("input[type ='text']").type('AAPL')
    cy.get('button').contains("Search").click()

    cy.get('button').contains("Save to Dashboard").click()

    cy.contains('Log In').click()
    cy.url().should('include', '/LogIn')

    cy.get("input[type ='text']").type('harry')
    cy.get("input[type ='password']").type('password')

    cy.contains("Log in").click()
    cy.url().should('include', '/dashboard')

    cy.contains('AAPL')
  })
})