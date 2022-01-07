package test

import (
	"fmt"
	"os"
	"testing"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

type user struct {
	ID        string
	name      string
	givenName string
	surname   string
	password  string
	email     string
}

type group struct {
	name                string
	groupLeaderUsername string
	ID                  string
}

type member struct {
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func setup() {
	user1 := user{userID, username, givenName, surname, userPassword, userEmail}
	user2 := user{userID2, username2, givenName2, surname2, userPassword2, userEmail2}

	leader1 := user{leaderID, leaderUsername, leaderGivenName, leaderSurname, leaderPassword, leaderEmail}
	leader2 := user{leaderID2, leaderUsername2, leaderGivenName2, leaderSurname2, leaderPassword2, leaderEmail2}
	leader3 := user{leaderID3, leaderUsername3, leaderGivenName3, leaderSurname3, leaderPassword3, leaderEmail3}

	group1 := group{groupName, groupLeaderUsername, groupID}
	group2 := group{groupName2, groupLeaderUsername2, groupID2}

	accountManagement := account.NewLDAPManagement()

	accountManagement.CreateUser(adminUser, adminPassword, user1.ID, user1.name, user1.givenName, user1.surname, user1.password, user1.email)
	accountManagement.CreateUser(adminUser, adminPassword, user2.ID, user2.name, user2.givenName, user2.surname, user2.password, user2.email)
	accountManagement.CreateUser(adminUser, adminPassword, userID3, username3, givenName3, surname3, userPassword3, userEmail3)

	accountManagement.CreateUser(adminUser, adminPassword, leader1.ID, leader1.name, leader1.givenName, leader1.surname, leader1.password, leader1.email)
	accountManagement.CreateUser(adminUser, adminPassword, leader2.ID, leader2.name, leader2.givenName, leader2.surname, leader2.password, leader2.email)
	accountManagement.CreateUser(adminUser, adminPassword, leader3.ID, leader3.name, leader3.givenName, leader3.surname, leader3.password, leader3.email)

	accountManagement.CreateOu(adminUser, adminPassword, ouName)

	accountManagement.CreateGroup(adminUser, adminPassword, group1.name, group1.groupLeaderUsername, group1.ID)
	accountManagement.CreateGroup(adminUser, adminPassword, group2.name, group2.groupLeaderUsername, group2.ID)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	user1 := user{userID, username, givenName, surname, userPassword, userEmail}
	user2 := user{userID2, username2, givenName2, surname2, userPassword2, userEmail2}

	leader1 := user{leaderID, leaderUsername, leaderGivenName, leaderSurname, leaderPassword, leaderEmail}
	leader2 := user{leaderID2, leaderUsername2, leaderGivenName2, leaderSurname2, leaderPassword2, leaderEmail2}
	leader3 := user{leaderID3, leaderUsername3, leaderGivenName3, leaderSurname3, leaderPassword3, leaderEmail3}

	group1 := group{groupName, groupLeaderUsername, groupID}
	group2 := group{groupName2, groupLeaderUsername2, groupID2}

	accountManagement := account.NewLDAPManagement()

	accountManagement.DeleteUser(adminUser, adminPassword, user1.name)
	accountManagement.DeleteUser(adminUser, adminPassword, user2.name)
	accountManagement.DeleteUser(adminUser, adminPassword, username3)

	accountManagement.DeleteUser(adminUser, adminPassword, leader1.name)
	accountManagement.DeleteUser(adminUser, adminPassword, leader2.name)
	accountManagement.DeleteUser(adminUser, adminPassword, leader3.name)

	accountManagement.DeleteGroup(adminUser, adminPassword, group1.name)
	accountManagement.DeleteGroup(adminUser, adminPassword, group2.name)

	accountManagement.DeleteOu(adminUser, adminPassword, ouName)

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed\n")
}
