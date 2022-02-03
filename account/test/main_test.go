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
	accountManagement := account.NewLDAPManagement()

	accountManagement.CreateUser(adminUser, adminPassword, userId1, username1, givenName1, surname1, userPassword1, userEmail1)
	accountManagement.CreateUser(adminUser, adminPassword, userId2, username2, givenName2, surname2, userPassword2, userEmail2)
	accountManagement.CreateUser(adminUser, adminPassword, userId3, username3, givenName3, surname3, userPassword3, userEmail3)

	accountManagement.CreateUser(adminUser, adminPassword, leaderId1, leaderUsername1, leaderGivenName1, leaderSurname1, leaderPassword1, leaderEmail1)
	accountManagement.CreateUser(adminUser, adminPassword, leaderId2, leaderUsername2, leaderGivenName2, leaderSurname2, leaderPassword2, leaderEmail2)
	accountManagement.CreateUser(adminUser, adminPassword, leaderId3, leaderUsername3, leaderGivenName3, leaderSurname3, leaderPassword3, leaderEmail3)

	accountManagement.CreateOu(adminUser, adminPassword, ouName)

	accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeaderUsername, groupId1)
	accountManagement.CreateGroup(adminUser, adminPassword, groupName2, groupLeaderUsername2, groupId2)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	accountManagement := account.NewLDAPManagement()

	accountManagement.DeleteUserByUserId(adminUser, adminPassword, userId1)
	accountManagement.DeleteUserByUserId(adminUser, adminPassword, userId2)
	accountManagement.DeleteUserByUserId(adminUser, adminPassword, userId3)

	accountManagement.DeleteUserByUserId(adminUser, adminPassword, leaderId1)
	accountManagement.DeleteUserByUserId(adminUser, adminPassword, leaderId2)
	accountManagement.DeleteUserByUserId(adminUser, adminPassword, leaderId3)

	accountManagement.DeleteGroup(adminUser, adminPassword, groupName)
	accountManagement.DeleteGroup(adminUser, adminPassword, groupName2)

	accountManagement.DeleteOu(adminUser, adminPassword, ouName)

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed\n")
}
