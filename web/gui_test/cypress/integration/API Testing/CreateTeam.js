describe("Create Team Success", () => {
    const inputTeam = 'OIS';
    const inputLeader = 'Patrick';

    it("Create Team", ()=>{
        cy.request({
        method : 'POST',
        url : "http://localhost:8080/create/team",
        body : {
            'GroupName': inputTeam, 
            'Username' : inputLeader
            }
        }).then((response)=> {            
            expect(response.status).to.be.equal(200);
            expect(response).has.property("body",inputTeam);

        })
    })

    it("Tear down", () => {
        cy.request({
            method : 'POST',
            url : "http://localhost:8080/delete/team",
            body : {
                'GroupName': inputTeam
                }
            }).then((response)=> {
                expect(response.status).to.be.equal(200);
            })
    });
});

describe("Create Team With Not Registered User as Leader", () => {
    const inputTeam = 'OIS';
    const notRegisterUser = 'Rebecca';

    it("Create Team", ()=>{
        cy.request({
        method : 'POST',
        url : "http://localhost:8080/create/team",
        failOnStatusCode: false,
        body : {
            'GroupName': inputTeam, 
            'Username' : notRegisterUser
            }
        }).then((response)=> {            
            expect(response.status).to.be.equal(500);
            expect(response).has.property("body","User does not exist");

        })
    })
});


describe("Create Team with Duplicate Team Name", () => {
    const inputTeam = 'OIS';
    const inputLeader = 'Richard';

    it("Create First Team", ()=>{
        cy.request({
            method : 'POST',
            url : "http://localhost:8080/create/team",
            body : {
                'GroupName': inputTeam, 
                'Username' : inputLeader
                }
            }).then((response)=> {            
                expect(response.status).to.be.equal(200);
                expect(response).has.property("body",inputTeam);
            })
        })

    it("Create Second Team with same Team Name", ()=>{
        cy.request({
            method : 'POST',
            url : "http://localhost:8080/create/team",
            failOnStatusCode: false,
            body : {
                'GroupName': inputTeam, 
                'Username' : inputLeader
                }
            }).then((response)=> {            
                expect(response.status).to.be.equal(500);
                expect(response).has.property("body","Duplicate Group Name");
            })
        })

    it("Tear down", () => {
        cy.request({
            method : 'POST',
            url : "http://localhost:8080/delete/team",
            body : {
                'GroupName': inputTeam
                }
            }).then((response)=> {
                expect(response.status).to.be.equal(200);
            })
        })
});