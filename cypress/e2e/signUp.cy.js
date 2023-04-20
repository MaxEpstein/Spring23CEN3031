describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:3000/')

    cy.contains('Log In').click()
    cy.url().should('include', '/Login')

    cy.get("input[type ='text']").type('harry')
    cy.get("input[type ='password']").type('password')

    cy.contains("Sign up").click()
    cy.url().should('include', '/dashboard')
  })
})