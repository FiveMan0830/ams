
describe("Create a team and show the team list.", () => {
    const inputTeam = 'Team Test';
    const inputLeader = 'ellen';

    it("Visit the website", ()=>{

        cy.visit("http://localhost:8080/team.html");
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
        const groupList = '//table[@id="groups"]'
        const groupTR = "//td[contains(text(),'"+inputTeam+"')]"
        cy.xpath(groupList)
          .should(($tr)=>{
            expect($tr).to.contain(inputTeam)
          });
        cy.xpath(groupTR)
          .siblings()
          .should("contain", inputLeader);
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
