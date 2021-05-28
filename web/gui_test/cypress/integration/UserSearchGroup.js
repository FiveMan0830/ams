
describe("Create a team and show the team list.", () => {
    const inputUser = 'ellen';

    const team1 = 'Test';
    // const team2 = 'eleengroup';

    it("Visit the website", ()=>{

        cy.visit("http://localhost:8080/user.html");
    });

    it("Create Team and enter leader's name", () => {
        const inputText = '//input[@id="username-field"]';
        const searchBtn = '//input[@id="search-button"]';

        cy.xpath(inputText)
          .type(inputUser)
          .should("have.value",inputUser);

        cy.xpath(searchBtn).click();
    });

    it("Assert Team", () =>{
        const groupList = '//ul[@id="groups"]'
        cy.xpath(groupList)
          .should(($li)=>{
            expect($li).to.contain(team1)
          });
        // cy.xpath(groupList)
        //   .should(($li)=>{
        //     expect($li).to.contain(team2)
        //   });
    })

});
