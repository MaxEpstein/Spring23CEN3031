describe('login and get dashboard', () => {
  it('finds dashboard', () => {
    cy.visit('http://localhost:3000/')
    cy.contains('Log In').click()
    cy.url().should('include', '/LogIn')

    cy.get("input[type ='text']").type('harry')
    cy.get("input[type ='password']").type('password')

    cy.contains("Log in / Sign up").click()
    cy.url().should('include', '/dashboard')
  })
})