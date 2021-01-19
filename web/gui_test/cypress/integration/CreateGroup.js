
describe("Create a team and show the team list.", () => {
    const inputTeam = 'Team Test';
    const inputLeader = 'Test';

    it("Visit the website", ()=>{

        cy.visit("http://127.0.0.1:5500/web/team.html");
    });

    it("Create Team and enter leader's name", () => {
        const teaminputText = '//input[@id="groupname-field"]';
        const leadeinputText = '//input[@id="username-field"]';
        const createBtn = '//input[@id="create-button"]';

        cy.xpath(teaminputText)
          .type(inputTeam)
          .should("have.value",inputTeam);

        cy.xpath(leadeinputText)
          .type(inputLeader)
          .should("have.value",inputLeader);

        cy.xpath(createBtn).click();
    });


    it("Assert Team", () =>{
        const groupList = '//ul[@id="groups"]'
        cy.xpath(groupList)
          .should(($li)=>{
            expect($li).to.contain(inputTeam)
          });
    })


    it("Tear down", () => {
        const inputText = '//input[@id="groupname-field"]';
        const deleteBtn = '//input[@id="delete-button"]';

        cy.xpath(inputText)
          .type(inputTeam)
          .should("have.value",inputTeam);

        cy.xpath(deleteBtn).click();
    });



});
