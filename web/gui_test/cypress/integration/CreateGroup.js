
describe("Create a team and show the team list.", () => {
    const inputTeam = 'Team Test';

    it("Visit the website", ()=>{

        cy.visit("http://127.0.0.1:5500/web/team.html");
    });

    it("Create Team", () => {
        const inputText = '//input[@id="groupname-field"]';
        const createBtn = '//input[@id="create-button"]';

        cy.xpath(inputText)
          .type(inputTeam)
          .should("have.value",inputTeam);

        cy.xpath(createBtn).click();
    });



    it("Tear down", () => {
        const inputText = '//input[@id="groupname-field"]';
        const deleteBtn = '//input[@id="delete-button"]';

        cy.xpath(inputText)
          .type(inputTeam)
          .should("have.value",inputTeam);

        cy.xpath(deleteBtn).click();
    });



});
