describe("Create Team", () => {
    const inputTeam = 'OIS';
    const inputLeader = 'Patrick';

    it("Create Team Success", ()=>{
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
})

describe("Add Member To Team", () => {
    const inputTeam = 'OIS';
    const inputMember = 'Richard';
    const expectedMember = ['Patrick','Richard'];

    it("Add Member To Team", ()=>{
        cy.request({
        method : 'POST',
        url : "http://localhost:8080/add/member",
        body : {
            'GroupName': inputTeam, 
            'Username' : inputMember
            }
        }).then((response)=> {            
            expect(response.status).to.be.equal(200);
            cy.get(response.body).each((body,index)=>{
                cy.wrap(body).should('contain',expectedMember[index])
            })
        })
    })

    it("Tear down", () => {
        cy.request({
            method : 'POST',
            url : "http://localhost:8080/remove/member",
            body : {
                'GroupName': inputTeam, 
                'Username' : inputMember
                }
            }).then((response)=> {
                expect(response.status).to.be.equal(200);
            })
    });
});

describe("Add Member To Team with Not Registered User", () => {
    const inputTeam = 'OIS';
    const inputMember = 'Rebecca';
    it("Add Member To Team", ()=>{
        cy.request({
        method : 'POST',
        url : "http://localhost:8080/add/member",
        failOnStatusCode: false,
        body : {
            'GroupName': inputTeam, 
            'Username' : inputMember
            }
        }).then((response)=> {            
            expect(response.status).to.be.equal(500);
            expect(response).has.property("body","User does not exist");
        })
    })
});

describe("Add Member To Team with Not Registered Team", () => {
    const inputTeam = 'Sunbird';
    const inputMember = 'Ron';
    it("Add Member To Team", ()=>{
        cy.request({
        method : 'POST',
        url : "http://localhost:8080/add/member",
        failOnStatusCode: false,
        body : {
            'GroupName': inputTeam, 
            'Username' : inputMember
            }
        }).then((response)=> {            
            expect(response.status).to.be.equal(500);
            expect(response).has.property("body","Group does not exist");
        })
    })
});

describe("Add Duplicate Member to Team", () => {
    const inputTeam = 'OIS';
    const inputMember = 'Richard';
    const expectedMember = ['Patrick','Richard'];

    it("Add Member To Team for The First Time", ()=>{
        cy.request({
        method : 'POST',
        url : "http://localhost:8080/add/member",
        body : {
            'GroupName': inputTeam, 
            'Username' : inputMember
            }
        }).then((response)=> {            
            expect(response.status).to.be.equal(200);
            cy.get(response.body).each((body,index)=>{
                cy.wrap(body).should('contain',expectedMember[index])
            })
        })
    })

    it("Add Member To Team for The Second Time", ()=>{
        cy.request({
        method : 'POST',
        url : "http://localhost:8080/add/member",
        failOnStatusCode: false,
        body : {
            'GroupName': inputTeam, 
            'Username' : inputMember
            }
        }).then((response)=> {            
            expect(response.status).to.be.equal(500);
            expect(response).has.property("body","User is already a member");
        })
    })

    it("Tear down", () => {
        cy.request({
            method : 'POST',
            url : "http://localhost:8080/remove/member",
            body : {
                'GroupName': inputTeam, 
                'Username' : inputMember
                }
            }).then((response)=> {
                expect(response.status).to.be.equal(200);
            })
    });

});



describe("Tear down", () => {
    const inputTeam = 'OIS';

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

// describe("Create Team With Not Registered User as Leader", () => {
//     const inputTeam = 'TestCreate';
//     const notRegisterUser = 'noOne';

//     it("Create Team", ()=>{
//         cy.request({
//         method : 'POST',
//         url : "http://localhost:8080/create/team",
//         failOnStatusCode: false,
//         body : {
//             'GroupName': inputTeam, 
//             'Username' : notRegisterUser
//             }
//         }).then((response)=> {            
//             expect(response.status).to.be.equal(500);
//             expect(response).has.property("body","User does not exist");

//         })
//     })
// });


// describe("Create Team with Duplicate Team Name", () => {
//     const inputTeam = 'TestCreate';
//     const inputLeader = 'Test';

//     it("Create First Team", ()=>{
//         cy.request({
//             method : 'POST',
//             url : "http://localhost:8080/create/team",
//             body : {
//                 'GroupName': inputTeam, 
//                 'Username' : inputLeader
//                 }
//             }).then((response)=> {            
//                 expect(response.status).to.be.equal(200);
//                 expect(response).has.property("body",inputTeam);
//             })
//         })

//     it("Create Second Team with same Team Name", ()=>{
//         cy.request({
//             method : 'POST',
//             url : "http://localhost:8080/create/team",
//             failOnStatusCode: false,
//             body : {
//                 'GroupName': inputTeam, 
//                 'Username' : inputLeader
//                 }
//             }).then((response)=> {            
//                 expect(response.status).to.be.equal(500);
//                 expect(response).has.property("body","Duplicate Group Name");
//             })
//         })

//     it("Tear down", () => {
//         cy.request({
//             method : 'POST',
//             url : "http://localhost:8080/delete/team",
//             body : {
//                 'GroupName': inputTeam
//                 }
//             }).then((response)=> {
//                 expect(response.status).to.be.equal(200);
//             })
//         })
// });