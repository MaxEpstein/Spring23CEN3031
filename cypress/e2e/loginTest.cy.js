describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:3000/')

    cy.contains('Log In').click()
    cy.url().should('include', '/Login')

    cy.get("input[type ='text']").type('l')
    cy.get("input[type ='password']").type('o')
    cy.contains("Log in").click()

    cy.get("input[type ='text']").type('test')
    cy.get("input[type ='password']").type('pass')
    cy.contains("Sign up").click()

    cy.url().should('include', '/dashboard')
  })
})