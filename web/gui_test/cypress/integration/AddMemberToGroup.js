
describe("Add a member to the group", () => {
    const inputTeam = 'Test';
    const inputUsername = 'audi';
    const teamText = '//input[@id="groupname-field"]';
    const usernameText = '//input[@id="username-field"]';


    it("Visit the website", ()=>{

        cy.visit("http://127.0.0.1:5500/web/member.html");
    });

    it("Enter Team's name and username", () => {
        const addBtn = '//input[@id="add-button"]';

        cy.xpath(teamText)
          .type(inputTeam)
          .should("have.value",inputTeam);

        cy.xpath(usernameText)
          .type(inputUsername)
          .should("have.value",inputUsername);

        cy.xpath(addBtn).click();
    });


    it("Assert Team", () =>{
        const groupList = '//ul[@id="member"]'
        const searchBtn = '//input[@id="get-member-button"]';

        cy.xpath(teamText).clear();
        cy.xpath(teamText)
          .type(inputTeam)
          .should("have.value",inputTeam);

        cy.xpath(searchBtn).click();

        cy.xpath(groupList)
          .should(($li)=>{
            expect($li).to.contain(inputUsername)
          });
    })


    it("Tear down", () => {
        const teamText = '//input[@id="groupname-field"]';
        const usernameText = '//input[@id="username-field"]';
        const removeBtn = '//input[@id="remove-button"]';

        cy.xpath(teamText).clear();
        cy.xpath(teamText)
          .type(inputTeam)
          .should("have.value",inputTeam);

        cy.xpath(usernameText).clear();
        cy.xpath(usernameText)
          .type(inputUsername)
          .should("have.value",inputUsername);

        cy.xpath(removeBtn).click();
    });



});
