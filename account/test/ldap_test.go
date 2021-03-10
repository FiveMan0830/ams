package test

import (
	"fmt"
	"os"
	"testing"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

const adminUser string = "admin"
const adminPassword string = "admin"

// User test data
const null string = ""
const usernameNotExists string = "usernameNotExists"

const userID string = "c61965be-8176-4419-b289-4d52617728fb"
const username string = "testUser"
const givenName string = "testUser"
const surname string = "testUser"
const userPassword string = "testUser"
const userEmail string = "test@gmail.com"

const userID2 string = "c56515be-1654-7895-1564-1d52513528cf"
const username2 string = "testUser2"
const givenName2 string = "testUser2"
const surname2 string = "testUser2"
const userPassword2 string = "testUser2"
const userEmail2 string = "test2@gmail.com"

const leaderID string = "a34531da-8563-9517-3578-3e38754896dg"
const leaderUsername string = "testLeader"
const leaderGivenName string = "testLeader"
const leaderSurname string = "testLeader"
const leaderPassword string = "testLeader"
const leaderEmail string = "testLeader@gmail.com"

const leaderID2 string = "b96875kl-6842-7539-8549-2c56482648fa"
const leaderUsername2 string = "testLeader2"
const leaderGivenName2 string = "testLeader2"
const leaderSurname2 string = "testLeader2"
const leaderPassword2 string = "testLeader2"
const leaderEmail2 string = "testLeade2r@gmail.com"

// Group test data
const groupNameNotExists = "testGroupNotExists"
const groupLeaderNotExists string = "testLeaderNotExists"

const groupName string = "testGroup"
const groupLeader string = "testLeader"

const groupName2 string = "testGroup2"
const groupLeader2 string = "testLeader2"


// Ou test data
const ouName = "testOu"

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	accountManagement := account.NewLDAPManagement()

	accountManagement.CreateUser(adminUser, adminPassword, userID, username, givenName, surname, userPassword, userEmail)
	accountManagement.CreateUser(adminUser, adminPassword, userID2, username2, givenName2, surname2, userPassword2, userEmail2)

	accountManagement.CreateUser(adminUser, adminPassword, leaderID, leaderUsername, leaderGivenName, leaderSurname, leaderPassword, leaderEmail)
	accountManagement.CreateUser(adminUser, adminPassword, leaderID2, leaderUsername2, leaderGivenName2, leaderSurname2, leaderPassword2, leaderEmail2)

	accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeader)
	accountManagement.CreateGroup(adminUser, adminPassword, groupName2, groupLeader2)

	accountManagement.CreateOu(adminUser, adminPassword, ouName)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	accountManagement := account.NewLDAPManagement()

	accountManagement.DeleteUser(adminUser, adminPassword, username)
	accountManagement.DeleteUser(adminUser, adminPassword, username2)

	accountManagement.DeleteUser(adminUser, adminPassword, leaderUsername)
	accountManagement.DeleteUser(adminUser, adminPassword, leaderUsername2)

	accountManagement.DeleteGroup(adminUser, adminPassword, groupName)
	accountManagement.DeleteGroup(adminUser, adminPassword, groupName2)

	accountManagement.DeleteOu(adminUser, adminPassword, ouName)

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed\n")
}