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
const usernameNotExists string = "Eric"

const userID string = "c61965be-8176-4419-b289-4d52617728fb"
const username string = "fiveman123"
const givenName string = "Richard"
const surname string = "Zheng"
const userPassword string = "richard123"
const userEmail string = "richardEmail@gmail.com"

const userID2 string = "c56515be-1654-7895-1564-1d52513528cf"
const username2 string = "audi98"
const givenName2 string = "Audi"
const surname2 string = "Wu"
const userPassword2 string = "audipm984"
const userEmail2 string = "audiEmail@gmail.com"

const userID3 string = "a56515be-5783-8738-1564-1d52513528cz"
const username3 string = "stella83"
const givenName3 string = "Stella"
const surname3 string = "Chen"
const userPassword3 string = "stella423"
const userEmail3 string = "stellaEmail@gmail.com"

const leaderID string = "a34531da-8563-9517-3578-3e38754896dg"
const leaderUsername string = "george88"
const leaderGivenName string = "George"
const leaderSurname string = "Lim"
const leaderPassword string = "george189"
const leaderEmail string = "georgeEmail@gmail.com"

const leaderID2 string = "b96875kl-6842-7539-8549-2c56482648fa"
const leaderUsername2 string = "david93"
const leaderGivenName2 string = "David"
const leaderSurname2 string = "Wang"
const leaderPassword2 string = "david632"
const leaderEmail2 string = "davidEmail@gmail.com"

const leaderID3 string = "c16475kl-6762-7489-8359-2c56846648fm"
const leaderUsername3 string = "sherry99"
const leaderGivenName3 string = "Sherry"
const leaderSurname3 string = "Ye"
const leaderPassword3 string = "sherry753"
const leaderEmail3 string = "sherryEmail@gmail.com"

// Group test data
const groupNameNotExists = "Sunbird"
const groupLeaderNotExists string = "Ron"

const groupName string = "OIS"
const groupLeaderUsername string = "david93"
const groupID string = "d23475kl-4862-7456-8473-2c53916648fn"

const groupName2 string = "SSL"
const groupLeaderUsername2 string = "george88"
const groupID2 string = "e16987kl-9512-7424-9629-2c56884248lm"

const groupName3 string = "KanBan"
const groupLeaderUsername3 string = "sherry88"

// Ou test data
const ouName = "LabGroup"
const ouName2 = "NTUT"

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
	accountManagement.CreateUser(adminUser, adminPassword, leaderID3, leaderUsername3, leaderGivenName3, leaderSurname3, leaderPassword3, leaderEmail3)

	accountManagement.CreateGroup(adminUser, adminPassword, groupName, groupLeaderUsername, groupID)
	accountManagement.CreateGroup(adminUser, adminPassword, groupName2, groupLeaderUsername2, groupID2)

	accountManagement.CreateOu(adminUser, adminPassword, ouName)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	accountManagement := account.NewLDAPManagement()

	accountManagement.DeleteUser(adminUser, adminPassword, username)
	accountManagement.DeleteUser(adminUser, adminPassword, username2)

	accountManagement.DeleteUser(adminUser, adminPassword, leaderUsername)
	accountManagement.DeleteUser(adminUser, adminPassword, leaderUsername2)
	accountManagement.DeleteUser(adminUser, adminPassword, leaderUsername3)

	accountManagement.DeleteGroup(adminUser, adminPassword, groupName)
	accountManagement.DeleteGroup(adminUser, adminPassword, groupName2)

	accountManagement.DeleteOu(adminUser, adminPassword, ouName)

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed\n")
}