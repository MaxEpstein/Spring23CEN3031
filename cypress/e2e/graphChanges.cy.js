describe('test graph buttons', () => {
  it('passes', () => {
    cy.visit('http://localhost:3000/')
    cy.contains('Search').click()
    cy.url().should('include', '/search')

    cy.get("input[type ='text']").type('amzn')
    cy.get('button').contains("Search").click()

    cy.contains("5 Day").click()
    cy.contains("1 Year").click()
  })
})